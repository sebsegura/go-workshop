package handler

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
	"seb7887/create-contact/internal/processor"
	"seb7887/create-contact/pkg/dto"
)

type Response = events.APIGatewayProxyResponse

type Handler interface {
	Create(ctx context.Context, req dto.Request) (Response, error)
}

type LambdaHandler struct {
	Processor processor.Processor
}

func (h *LambdaHandler) setProcessor() {
	if h.Processor == nil {
		h.Processor = &processor.LambdaProcessor{}
	}
	// Setup processor
	h.Processor.Setup()
}

func (h *LambdaHandler) Create(ctx context.Context, req dto.Request) (Response, error) {
	log.Println("Setting processor...")
	h.setProcessor()
	log.Println("Running processor...")

	// equals to:
	// var contact dto.Contact
	// var error Error
	contact, err := h.Processor.Process(req)
	if err != nil {
		return Response{}, err
	}

	body, err := json.Marshal(contact)
	if err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: http.StatusOK,
		Body: string(body),
	}, nil
}