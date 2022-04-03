package contactform_test

import (
	cf "contactform"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type mockSendgrid struct {
}

func (m *mockSendgrid) Send(email *mail.SGMailV3) (*rest.Response, error) {
	return &rest.Response{StatusCode: 200, Body: "success", Headers: nil}, nil
}

var _ = Describe("Contact form tests", func() {
	form_input_test := &cf.CreateEmailInput{Email: "testing@test.com", FirstName: "nick", LastName: "fig", RequestType: "tutoring", Message: "yo"}
	res, err := cf.GetContactFormEmail(form_input_test)

	Context("Valid input test", func() {
		It("Successfully processes the input", func() {
			Expect(err).To(BeNil())
			Expect(res.To.Address).To(BeEquivalentTo(cf.TutoringEmail.String()))
			Expect(res.From.Address).To(BeEquivalentTo(cf.NoReplyEmail.String()))
			Expect(res.Subject).To(BeEquivalentTo("Contact Form Submission from nick fig"))
		})

		It("Has the correct message body", func () {
			Expect(res.PlainTextContent).To(BeEquivalentTo(cf.FormatPlainTextFormSubmission(form_input_test)))
			Expect(res.PlainTextContent).To(ContainSubstring("Name: nick fig"))
			Expect(res.PlainTextContent).To(ContainSubstring("Message: yo"))
		})

		client := cf.NewEmailClient(&mockSendgrid{})
		It("Sends the email", func() {
			response, err := client.SendEmail(res)
			Expect(response).To(BeEquivalentTo("success"))
			Expect(err).To(BeNil())
		})
	})
})


var _ = Describe("CORS Requests", func() {
	Context("Valid CORS request", func() {
		client := cf.NewEmailClient(&mockSendgrid{})
		ctx := context.Background()
		req := &events.APIGatewayProxyRequest{}
		req.Headers = map[string]string{
			"Access-Control-Request-Method": "POST",
			"Access-Control-Request-Headers": "Content-Type, Accept",
			"Origin": "nickfiggins.com",
		}
		req.HTTPMethod = "OPTIONS"
		resp, err := client.Handle(ctx, req)
		It("Successfully gave a response to the preflight request", func() {
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(BeEquivalentTo(204))
		})

		It("Has the correct headers", func() {
			headers := resp.Headers
			Expect(headers["Access-Control-Allow-Origin"]).To(BeEquivalentTo("*"))
			Expect(headers["Access-Control-Allow-Headers"]).To(BeEquivalentTo("*"))
			Expect(headers["Access-Control-Allow-Methods"]).To(BeEquivalentTo("GET, POST, OPTIONS"))
			Expect(headers["Access-Control-Allow-Credentials"]).To(BeEquivalentTo("true"))
		})
	})
})

type LambdaId string

const AwsRequestId LambdaId = "AwsRequestId"


var _ = Describe("End-to-end handler requests", func() {
	client := cf.NewEmailClient(&mockSendgrid{})
	form_input := &cf.CreateEmailInput{Email: "testing@test.com", FirstName: "nick", LastName: "fig", RequestType: "tutoring", Message: "yo"}
	json_body, _ := json.Marshal(form_input)
	req := &events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: string(json_body)}
	ctx := context.WithValue(context.Background(), AwsRequestId, "test req id")
	Context("Happy path", func() {
		resp, err := client.Handle(ctx, req)
		It("Successfully handles the request", func() {
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(BeEquivalentTo(200))
		})

		It("Gives a successful response", func() {
			Expect(resp.Body).To(BeEquivalentTo("success"))
		})
	})

	Context("Invalid JSON Body", func() {
		req := &events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: "{"}
		resp, _ := client.Handle(ctx, req)
		It("Returns an error", func() {
			Expect(resp.StatusCode).To(BeEquivalentTo(400))
			Expect(resp.Body).ToNot(BeEmpty())
		})
	})

	Context("Empty message", func() {
		form_input.Message = ""
		json_body, _ := json.Marshal(form_input)
		req := &events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: string(json_body)}
		resp, _ := client.Handle(ctx, req)
		It("Returns successfully", func() {
			Expect(resp.StatusCode).To(BeEquivalentTo(200))
			Expect(resp.Body).To(BeEquivalentTo("Empty request was received"))
		})
	})
})