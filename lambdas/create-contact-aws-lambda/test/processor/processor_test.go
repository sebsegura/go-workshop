package processor_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"uala/go-workshop/internal/processor"
	"uala/go-workshop/test/common"
	"uala/go-workshop/test/mocks"
)

var _ = Describe("Test", func() {
	Describe("Procesor", func() {
		Context("Happy path", happyPath)
		Context("Internal error path", internalErrorPath)
	})
})

func happyPath() {
	It("should process a request", func() {
		var (
			mockRepository = mocks.MockContactRepository{}
		)

		mockRepository.On("PutItem", common.MockContact).Return(nil)

		p := processor.New(&mockRepository)

		contact, err := p.Process(common.MockRequest)

		// Assertions
		Ω(err).Should(BeNil())
		Ω(contact).Should(BeEquivalentTo(common.MockContact))
	})
}

func internalErrorPath() {
	It("should return error if db operations fail", func() {
		var (
			mockRepository = mocks.MockContactRepository{}
		)

		mockRepository.On("PutItem", common.MockContact).Return(&ddb.InternalError{
			Op:  "PutItem",
			Err: errors.New("something bad happened"),
		})

		p := processor.New(&mockRepository)

		_, err := p.Process(common.MockRequest)

		// Assertions
		Ω(err).ShouldNot(BeNil())
		Ω(errors.As(err, &common.MockInternalError)).Should(BeTrue())
	})
}
