package main

import (
	"api/gateway"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func HandleRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	err := godotenv.Load()
	if err != nil {
		return gateway.NewResponse(500, "Internal server error"), err
	}

	stage := os.Getenv("STAGE")
	message := fmt.Sprintf("Hello from %s!", strings.ToUpper(stage))

	return gateway.NewResponse(200, message), nil
}

func main() {
	lambda.Start(HandleRequest)
}
