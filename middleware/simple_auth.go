package middleware

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	cronus "github.com/gianargamosa/booky-middleware/core"
)

func SimpleAuthMiddleware(validToken string) cronus.Middleware {
	return func(next cronus.Handler) cronus.Handler {
		return func(ctx context.Context, req cronus.Request) (events.APIGatewayProxyResponse, error) {
			token := req.Headers["Authorization"]
			if token != validToken {
				return events.APIGatewayProxyResponse{
					StatusCode: 401,
					Headers:    map[string]string{"Content-Type": "application/json"},
					Body:       `{"error": "Unauthorized"}`,
				}, nil
			}

			return next(ctx, req)
		}
	}
}
