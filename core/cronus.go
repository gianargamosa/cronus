package core

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Request struct {
	events.APIGatewayProxyRequest
	Body interface{}
}

type Handler func(context.Context, Request) (events.APIGatewayProxyResponse, error)

type Middleware func(next Handler) Handler

type LambdaHandler struct {
	handler     Handler
	middlewares []Middleware
}

func Create(handler Handler) *LambdaHandler {
	return &LambdaHandler{
		handler: handler,
	}
}

func (lh *LambdaHandler) Use(middleware Middleware) {
	lh.middlewares = append(lh.middlewares, middleware)
}

func (lh *LambdaHandler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	handler := lh.handler
	for _, middleware := range lh.middlewares {
		handler = middleware(handler)
	}

	requestData := Request{
		APIGatewayProxyRequest: req,
		Body:                   nil,
	}

	return handler(ctx, requestData)
}
