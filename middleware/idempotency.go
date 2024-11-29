package middleware

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	cronus "github.com/gianargamosa/booky-middleware/core"
)

func IdempotencyMiddleware(responseCache map[string]events.APIGatewayProxyResponse) cronus.Middleware {
	return func(next cronus.Handler) cronus.Handler {
		return func(ctx context.Context, req cronus.Request) (events.APIGatewayProxyResponse, error) {
			idempotencyKey := req.Headers["X-Idempotency-Key"]
			if idempotencyKey == "" {
				return next(ctx, req)
			}

			if cachedResponse, exists := responseCache[idempotencyKey]; exists {
				fmt.Println("Returning cached response for Idempotency-Key:", idempotencyKey)
				return cachedResponse, nil
			}

			resp, err := next(ctx, req)
			if err != nil {
				return events.APIGatewayProxyResponse{}, err
			}

			responseCache[idempotencyKey] = resp
			fmt.Println("Caching response for Idempotency-Key:", idempotencyKey)

			return resp, nil
		}
	}
}
