package contactform

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


type EmailClient struct {
	sendgrid SendgridAPI
}

func is2xxSuccessful(code int) bool {
	return (code >= 200) && (code < 300)
}

func (e *EmailClient) SendEmail(email *Email) (string, error) {
	res, err := e.sendgrid.Send(
		mail.NewSingleEmailPlainText(email.From, email.Subject, email.To, email.PlainTextContent),
		); if err != nil || !is2xxSuccessful(res.StatusCode) {
			errorMsg := fmt.Errorf("unable to send email, err: %v response: %v", err, res)
			return "", errorMsg
	}
	return res.Body, nil
}

func (e *EmailClient) Cors() (*events.APIGatewayProxyResponse) {
	headers := map[string]string{}
	headers["Access-Control-Allow-Origin"] = "*"
	headers["Access-Control-Allow-Headers"] = "*"
	headers["Access-Control-Allow-Methods"] = "GET, POST, OPTIONS"
	headers["Access-Control-Allow-Credentials"] = "true"
	return &events.APIGatewayProxyResponse{StatusCode: 204, Headers: headers}
}

func NewEmailClient(sg SendgridAPI) *EmailClient {
	return &EmailClient{sendgrid: sg}
}

func (e *EmailClient) Handle(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "OPTIONS" {
		return e.Cors(), nil
	}

	var input CreateEmailInput
	err := json.Unmarshal([]byte(request.Body), &input); if err != nil {
		return &events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
	}

	if input.Email == "" || input.Message == "" {
		return &events.APIGatewayProxyResponse{Body: "Empty request was received", StatusCode: 200}, nil
	}

	input.GwId = request.RequestContext.RequestID
	lambdaId, ok := ctx.Value("AwsRequestId").(string); if ok {
		input.LambdaId = lambdaId
	} else {
		log.Println("AwsRequestId was not set", ctx)
	}

	email, err := GetContactFormEmail(&input); if err != nil {
		return &events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	res, err := e.SendEmail(email); if err != nil {
		return &events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}


	return &events.APIGatewayProxyResponse{Body: res, StatusCode: 200}, nil
}