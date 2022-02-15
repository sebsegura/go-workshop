package handler_test

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambdacontext"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"net/http"
	"uala/go-workshop/pkg/dto"
	"uala/go-workshop/pkg/handler"
	"uala/go-workshop/test/common"
	"uala/go-workshop/test/mocks"
)

var (
	ct = &lambdacontext.LambdaContext{
		AwsRequestID:       "awsRequestId1234",
		InvokedFunctionArn: "arn:aws:lambda:xxx",
		Identity:           lambdacontext.CognitoIdentity{},
		ClientContext:      lambdacontext.ClientContext{},
	}
	ctx             = lambdacontext.NewContext(context.TODO(), ct)
)

var _ = Describe("Handler", func() {
	Describe("Handler", func() {
		Context("Happy path", happyPath)
		Context("Validation error", validationError)
		Context("Internal server error", internalServerError)
	})
})

func happyPath() {
	It("should handle a request", func() {
		var (
			mockProcessor = mocks.MockProcessor{}
		)

		mockProcessor.On("Process", common.MockRequest).Return(common.MockContact, nil)

		h := handler.New(&mockProcessor)

		res, err := h.Create(ctx, common.MockRequest)

		// Assertions
		Ω(err).Should(BeNil())
		Ω(res.StatusCode).Should(BeEquivalentTo(http.StatusOK))
	})
}

func validationError() {
	It("should return error if cannot validate request", func() {
		var (
			mockProcessor = mocks.MockProcessor{}
			badRequest    = dto.Request{
				FirstName: "",
				LastName:  "",
			}
		)

		h := handler.New(&mockProcessor)

		res, err := h.Create(ctx, badRequest)

		// Assertions
		Ω(err).Should(BeNil())
		Ω(res.StatusCode).Should(BeEquivalentTo(http.StatusBadRequest))
	})
}

func internalServerError() {
	It("should return error if an internal error happens", func() {
		var (
			mockProcessor = mocks.MockProcessor{}
		)

		mockProcessor.On("Process", common.MockRequest).Return(ddb.Contact{}, &ddb.InternalError{
			Op:  "PutItem",
			Err: errors.New("something bad happened"),
		})

		h := handler.New(&mockProcessor)

		res, err := h.Create(ctx, common.MockRequest)

		// Assertions
		Ω(err).Should(BeNil())
		Ω(res.StatusCode).Should(BeEquivalentTo(http.StatusInternalServerError))
	})
}
