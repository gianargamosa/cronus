package middleware

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	cronus "github.com/gianargamosa/booky-middleware/core"
)

func LogMiddleware() cronus.Middleware {
	return func(next cronus.Handler) cronus.Handler {
		return func(ctx context.Context, req cronus.Request) (events.APIGatewayProxyResponse, error) {
			fmt.Printf("Received request: %s %s\n", req.HTTPMethod, req.Path)
			return next(ctx, req)
		}
	}
}
