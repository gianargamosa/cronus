package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
	cronus "github.com/gianargamosa/booky-middleware/core"
	"github.com/go-playground/validator"
)

func SchemaValidatorMiddleware(v *validator.Validate, bodyType interface{}) cronus.Middleware {
	return func(next cronus.Handler) cronus.Handler {
		return func(ctx context.Context, req cronus.Request) (events.APIGatewayProxyResponse, error) {
			bodyInstance := reflect.New(reflect.TypeOf(bodyType)).Interface()

			if err := json.Unmarshal([]byte(req.APIGatewayProxyRequest.Body), bodyInstance); err != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Body:       `{"error": "Invalid JSON body"}`,
				}, nil
			}

			if err := v.Struct(bodyInstance); err != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Body:       fmt.Sprintf(`{"error": "Validation failed", "details": "%s"}`, err.Error()),
				}, nil
			}

			req.Body = bodyInstance

			return next(ctx, req)
		}
	}
}
