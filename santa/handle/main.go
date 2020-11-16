package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("bod", req.Body)
	fmt.Println("meth", req.HTTPMethod)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Helloooo, %v", "hi"),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
