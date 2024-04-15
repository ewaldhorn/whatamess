package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Lambda Handler Called")

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("%s", request.Body),
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
