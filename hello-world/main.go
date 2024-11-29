package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	cronus "github.com/gianargamosa/booky-middleware/core"
	middleware "github.com/gianargamosa/booky-middleware/middleware"
	"github.com/go-playground/validator"
)

type RequestBody struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required,min=18"`
}

func handler(ctx context.Context, request cronus.Request) (events.APIGatewayProxyResponse, error) {
	req, _ := request.Body.(*RequestBody)

	greeting := fmt.Sprintf("Hello, %s! Age: %d\n", req.Name, req.Age)

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

func main() {
	const sampleValidToken = "Bearer my-secret-token"

	responseCache := make(map[string]events.APIGatewayProxyResponse)
	validate := validator.New()

	handler := cronus.Create(handler)
	handler.Use(middleware.ErrorHandlerMiddleware())
	handler.Use(middleware.SchemaValidatorMiddleware(validate, RequestBody{}))
	handler.Use(middleware.IdempotencyMiddleware(responseCache))
	handler.Use(middleware.LogMiddleware())
	handler.Use(middleware.SimpleAuthMiddleware(sampleValidToken))

	lambda.Start(handler.Handle)
}
