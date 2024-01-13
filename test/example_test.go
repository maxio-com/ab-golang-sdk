package test

import (
	"testing"

	"github.com/caarlos0/env/v10"
	advancedbilling "github.com/maxio-com/ab-golang-sdk"
	"github.com/stretchr/testify/suite"
)

type Config struct {
	APIKey    string `env:"TEST_API_KEY"`
	Password  string `env:"TEST_API_PASSWORD"`
	Domain    string `env:"TEST_DOMAIN"`
	Subdomain string `env:"TEST_SUBDOMAIN"`
}

type APITestSuite struct {
	suite.Suite
	client advancedbilling.ClientInterface
}

func (s *APITestSuite) SetupTest() {
	cfg := Config{}
	s.Assert().NoErrorf(env.Parse(&cfg), "could not load test config")

	config := advancedbilling.CreateConfiguration(
		advancedbilling.WithBasicAuthUserName(cfg.APIKey),
		advancedbilling.WithBasicAuthPassword(cfg.Password),
		advancedbilling.WithDomain(cfg.Domain),
		advancedbilling.WithSubdomain(cfg.Subdomain),
	)

	s.client = advancedbilling.NewClient(config)
}

func (s *APITestSuite) TestExample() {
	s.T().Log(s.client.Configuration().Domain())
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
