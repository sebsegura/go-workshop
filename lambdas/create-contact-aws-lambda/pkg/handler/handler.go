package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
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
	// todas las dependencias que se le quieran inyectar a esta entidad
	// TODO: processor
}

func (h *LambdaHandler) Create(ctx context.Context, req dto.Request) (Response, error) {
	return Response{}, nil
}
