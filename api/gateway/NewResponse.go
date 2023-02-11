package gateway

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func NewResponse(code int, message string) events.APIGatewayProxyResponse {
	if code != 200 {
		fmt.Printf("[ERROR: %d] %s\n\n", code, message)
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      code,
		IsBase64Encoded: false,
		Body:            message,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}
}
