package service

import (
	"context"
	"net/http"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/hrvadl/studdy-buddy/email/pkg/config"
)

func InitEmail(c *config.Config) *EmailService {
	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", c.EmailApiKey)
	cl := brevo.NewAPIClient(cfg)
	return &EmailService{cl}
}

type EmailService struct {
	client *brevo.APIClient
}

func (s *EmailService) SendEmail(emailPayload *brevo.SendSmtpEmail) (*http.Response, error) {
	ctx := context.Background()
	_, res, err := s.client.TransactionalEmailsApi.SendTransacEmail(ctx, *emailPayload)
	return res, err
}
