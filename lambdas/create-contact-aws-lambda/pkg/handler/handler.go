package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"uala/go-workshop/internal/processor"
	"uala/go-workshop/internal/repository"
	"uala/go-workshop/pkg/dto"
)

// Responsibilidad: manejar el evento que dispara el lambda
// Validar el input
// Si es necesario, dar una respuesta
type Response = events.APIGatewayProxyResponse

type Handler interface {
	Create(ctx context.Context, req dto.Request) (Response, error)
}

type LambdaHandler struct {
	ContactProcessor processor.Processor
}

func NewHandler() Handler {
	r := &repository.LambdaRepository{}
	r.Setup()
	p := processor.NewProcessor(r)
	return &LambdaHandler{
		ContactProcessor: p,
	}
}

func (h *LambdaHandler) Create(ctx context.Context, req dto.Request) (Response, error) {
	// TODO: Do some dummy validation
	if req.FirstName == "" || req.LastName == "" {
		// error de validacion
		return Response{
			StatusCode: http.StatusBadRequest,
			Body: dto.WrongRequestError.Error(),
		}, nil
	}

	// TODO: Process
	_, err := h.ContactProcessor.Process(req)
	if err != nil {

	}

	return Response{}, nil
}
