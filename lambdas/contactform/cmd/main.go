package main

import (
	"contactform"
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sendgrid/sendgrid-go"
)

var Sg *sendgrid.Client

func handler(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return contactform.NewEmailClient(Sg).Handle(ctx, request)
}

func init() {
	apiKey := os.Getenv("SENDGRID_API_KEY"); if len(apiKey) == 0 {
		panic("Sendgrid API key has not been set")
	}
	Sg = sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
}

func main() {
	lambda.Start(handler)
}