package test

import (
	"testing"

	"github.com/caarlos0/env/v10"
	"github.com/jaswdr/faker"
	advancedbilling "github.com/maxio-com/ab-golang-sdk"
	"github.com/stretchr/testify/suite"
)

type Config struct {
	APIKey    string `env:"TEST_API_KEY,required"`
	Password  string `env:"TEST_API_PASSWORD,required"`
	Domain    string `env:"TEST_DOMAIN,required"`
	Subdomain string `env:"TEST_SUBDOMAIN,required"`
}

type APISuite struct {
	suite.Suite
	fkr faker.Faker

	client             advancedbilling.ClientInterface
	unauthorizedClient advancedbilling.ClientInterface
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APISuite))
}

func (s *APISuite) SetupTest() {
	s.fkr = faker.New()

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	config := advancedbilling.CreateConfiguration(
		advancedbilling.WithBasicAuthUserName(cfg.APIKey),
		advancedbilling.WithBasicAuthPassword(cfg.Password),
		advancedbilling.WithDomain(cfg.Domain),
		advancedbilling.WithSubdomain(cfg.Subdomain),
	)

	configUnauthorized := advancedbilling.CreateConfiguration(
		advancedbilling.WithDomain(cfg.Domain),
		advancedbilling.WithSubdomain(cfg.Subdomain),
	)

	s.client = advancedbilling.NewClient(config)
	s.unauthorizedClient = advancedbilling.NewClient(configUnauthorized)
}
