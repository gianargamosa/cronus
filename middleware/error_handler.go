package middleware

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	cronus "github.com/gianargamosa/booky-middleware/core"
)

func ErrorHandlerMiddleware() cronus.Middleware {
	return func(next cronus.Handler) cronus.Handler {
		return func(ctx context.Context, req cronus.Request) (events.APIGatewayProxyResponse, error) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovered from panic: %v\n", r)
				}
			}()

			resp, err := next(ctx, req)
			if err != nil {
				fmt.Printf("Error: %v\n", err)

				return events.APIGatewayProxyResponse{
					StatusCode: 500,
					Body:       `{"error": "Internal Server Error"}`,
				}, nil
			}

			return resp, nil
		}
	}
}
