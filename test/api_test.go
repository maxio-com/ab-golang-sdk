package test

import (
	"math/rand"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/jaswdr/faker"
	advancedbilling "github.com/maxio-com/ab-golang-sdk"
	"github.com/maxio-com/ab-golang-sdk/models"
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
	baseBillingAddress baseBillingAddress
}

type baseBillingAddress struct {
	firstName string
	lastName  string
	phone     string
	address   string
	city      string
	state     string
	country   string
	zip       string
	card      card
}

type card struct {
	expirationMonth int
	expirationYear  int
	cardType        models.CardType
}

func (s *APISuite) SetupTest() {
	s.fkr = faker.NewWithSeed(rand.NewSource(time.Now().Unix()))

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

	s.baseBillingAddress = baseBillingAddress{
		firstName: "Joe",
		lastName:  "Smith",
		phone:     "4111111111111111",
		card: card{
			expirationMonth: int(time.Now().Month()),
			expirationYear:  time.Now().Year() + 1,
			cardType:        models.CardType_VISA,
		},
		city:    "Boston",
		state:   "MA",
		address: "123 Mass Ave.",
		country: "US",
		zip:     "02120",
	}
}
