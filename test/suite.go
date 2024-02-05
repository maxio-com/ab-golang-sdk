package test

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
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

func (s *APISuite) generateCustomer(ctx context.Context) models.Customer {
	person := s.fkr.Person()

	resp, err := s.client.CustomersController().CreateCustomer(ctx, &models.CreateCustomerRequest{
		Customer: models.CreateCustomer{
			FirstName:    person.FirstName(),
			LastName:     person.LastName(),
			Email:        person.Contact().Email,
			Organization: strPtr(s.fkr.Company().Name()),
			Reference:    strPtr(fmt.Sprintf("%d", s.fkr.RandomNumber(10))),
			Address:      strPtr(person.Faker.Address().StreetAddress()),
			City:         strPtr(person.Faker.Address().City()),
			State:        strPtr(person.Faker.Address().State()),
			Zip:          strPtr(person.Faker.Address().PostCode()),
			Country:      strPtr(person.Faker.Address().Country()),
			Phone:        strPtr(person.Faker.Address().Country()),
		},
	})

	s.NoError(err)
	s.Equal(http.StatusCreated, resp.Response.StatusCode)

	return resp.Data.Customer
}

func (s *APISuite) generateProductFamily(ctx context.Context) models.ProductFamily {
	resp, err := s.client.ProductFamiliesController().CreateProductFamily(ctx, &models.CreateProductFamilyRequest{
		ProductFamily: models.CreateProductFamily{
			Name: strPtr(s.fkr.RandomStringWithLength(20)),
		},
	})

	s.NoError(err)
	s.Equal(http.StatusCreated, resp.Response.StatusCode)

	return *resp.Data.ProductFamily
}

func (s *APISuite) generateProduct(ctx context.Context, productFamilyID int) models.Product {
	resp, err := s.client.ProductsController().CreateProduct(ctx, productFamilyID, &models.CreateOrUpdateProductRequest{
		Product: models.CreateOrUpdateProduct{
			Name:         s.fkr.RandomStringWithLength(20),
			PriceInCents: 50,
			Interval:     1,
			IntervalUnit: models.IntervalUnit_MONTH,
		},
	})

	s.NoError(err)
	s.Equal(http.StatusCreated, resp.Response.StatusCode)

	return resp.Data.Product
}

func (s *APISuite) generateCoupon(ctx context.Context, productFamilyID int) models.Coupon {
	coupon := &models.Coupon{
		Name:                        strPtr(s.fkr.RandomStringWithLength(20)),
		Code:                        strPtr("100OFF" + s.fkr.RandomStringWithLength(30)),
		Description:                 strPtr(s.fkr.RandomStringWithLength(20)),
		Percentage:                  models.NewOptional[string](strPtr("50")),
		AllowNegativeBalance:        boolPtr(false),
		Recurring:                   boolPtr(false),
		EndDate:                     models.NewOptional[string](strPtr(newDate())),
		ProductFamilyId:             &productFamilyID,
		Stackable:                   boolPtr(false),
		ExcludeMidPeriodAllocations: boolPtr(true),
		ApplyOnCancelAtEndOfPeriod:  boolPtr(true),
	}

	resp, err := s.client.CouponsController().CreateCoupon(ctx, productFamilyID, &models.CreateOrUpdateCoupon{
		Coupon: interfacePtr(coupon),
	})

	s.NoError(err)
	s.Equal(http.StatusCreated, resp.Response.StatusCode)

	return *resp.Data.Coupon
}

func (s *APISuite) generateMeteredComponent(ctx context.Context, productFamilyID int) models.Component {
	resp, err := s.client.ComponentsController().CreateMeteredComponent(
		ctx,
		productFamilyID,
		&models.CreateMeteredComponent{
			MeteredComponent: models.MeteredComponent{
				Name:          s.fkr.RandomStringWithLength(20),
				UnitName:      s.fkr.RandomStringWithLength(10),
				Taxable:       boolPtr(false),
				PricingScheme: models.PricingScheme_STAIRSTEP,
				Prices: []models.Price{
					{
						StartingQuantity: 1,
						UnitPrice:        1,
					},
				},
			},
		},
	)

	s.NoError(err)
	s.Equal(http.StatusCreated, resp.Response.StatusCode)

	return resp.Data.Component
}

func (s *APISuite) generateSubscription(
	ctx context.Context,
	customer models.Customer,
	product models.Product,
	couponCode string,
	components []models.CreateSubscriptionComponent,
) models.Subscription {
	resp, err := s.client.SubscriptionsController().CreateSubscription(
		ctx,
		&models.CreateSubscriptionRequest{
			Subscription: models.CreateSubscription{
				ProductId:               product.Id,
				PaymentCollectionMethod: toPtr[models.CollectionMethod](models.CollectionMethod_AUTOMATIC),
				CustomerId:              customer.Id,
				Currency:                strPtr("USD"),
				InitialBillingAt:        timePtr(time.Date(2029, 8, 29, 12, 0, 0, 0, time.UTC)),
				CouponCode:              &couponCode,
				Components:              components,
				PaymentProfileAttributes: &models.PaymentProfileAttributes{
					FirstName:       &s.baseBillingAddress.firstName,
					LastName:        &s.baseBillingAddress.lastName,
					FullNumber:      &s.baseBillingAddress.phone,
					CardType:        toPtr[models.CardType](models.CardType_VISA),
					ExpirationMonth: interfacePtr(1),
					ExpirationYear:  interfacePtr(time.Now().Year() + 1),
					BillingAddress:  &s.baseBillingAddress.address,
					BillingCity:     &s.baseBillingAddress.city,
					BillingState:    &s.baseBillingAddress.state,
					BillingCountry:  &s.baseBillingAddress.country,
					BillingZip:      &s.baseBillingAddress.zip,
				},
			},
		},
	)

	s.NoError(err)
	s.Equal(http.StatusCreated, resp.Response.StatusCode)

	return *resp.Data.Subscription
}
