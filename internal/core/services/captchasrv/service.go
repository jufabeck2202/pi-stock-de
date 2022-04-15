package captchasrv

import (
	"os"
	"time"

	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

type service struct {
	captchaClient *recaptcha.ReCAPTCHA
}

func New() (*service, error) {
	Capcha, err := recaptcha.NewReCAPTCHA(os.Getenv("RECAPTCHA_SECRET"), recaptcha.V3, 10*time.Second) // for v3 API use https://g.co/recaptcha/v3 (apperently the same admin UI at the time of writing)

	return &service{
		captchaClient: &Capcha,
	}, err
}

func (s *service) Verify(captcha string) error {
	return s.captchaClient.Verify(captcha)
}
