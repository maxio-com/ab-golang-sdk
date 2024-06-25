package test

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v10"
	"math/rand"
	"net/http"
	"time"

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
		advancedbilling.WithBasicAuthCredentials(
			advancedbilling.NewBasicAuthCredentials(
				cfg.APIKey,
				cfg.Password,
			),
		),
		advancedbilling.WithDomain(cfg.Domain),
		advancedbilling.WithSubdomain(cfg.Subdomain),
	)

	configUnauthorized := advancedbilling.CreateConfiguration(
		advancedbilling.WithBasicAuthCredentials(
			advancedbilling.NewBasicAuthCredentials(
				"abc",
				"abc",
			),
		),
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
			Name: s.fkr.RandomStringWithLength(20),
		},
	})

	s.NoError(err)
	s.Equal(http.StatusCreated, resp.Response.StatusCode)

	return *resp.Data.ProductFamily
}

func (s *APISuite) generateProduct(ctx context.Context, productFamilyID int) models.Product {
	resp, err := s.client.ProductsController().CreateProduct(ctx, fmt.Sprint(productFamilyID), &models.CreateOrUpdateProductRequest{
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
	coupon := &models.CreateOrUpdatePercentageCoupon{
		Name:                        s.fkr.RandomStringWithLength(20),
		Code:                        "100OFF" + s.fkr.RandomStringWithLength(30),
		Description:                 strPtr(s.fkr.RandomStringWithLength(20)),
		Percentage:                  models.CreateOrUpdatePercentageCouponPercentageContainer.FromString("50"),
		AllowNegativeBalance:        boolPtr(false),
		Recurring:                   boolPtr(false),
		EndDate:                     timePtr(time.Now().AddDate(1, 0, 0)),
		ProductFamilyId:             models.ToPointer(fmt.Sprint(productFamilyID)),
		Stackable:                   boolPtr(false),
		ExcludeMidPeriodAllocations: boolPtr(true),
		ApplyOnCancelAtEndOfPeriod:  boolPtr(true),
	}

	resp, err := s.client.CouponsController().CreateCoupon(ctx, productFamilyID, &models.CreateOrUpdateCoupon{
		Coupon: models.ToPointer(models.CreateOrUpdateCouponCouponContainer.FromCreateOrUpdatePercentageCoupon(*coupon)),
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
						StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(1),
						UnitPrice:        models.PriceUnitPriceContainer.FromPrecision(1),
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
				InitialBillingAt:        timePtr(time.Now().AddDate(1, 0, 0)),
				CouponCode:              &couponCode,
				Components:              components,
				PaymentProfileAttributes: &models.PaymentProfileAttributes{
					FirstName:       &s.baseBillingAddress.firstName,
					LastName:        &s.baseBillingAddress.lastName,
					FullNumber:      &s.baseBillingAddress.phone,
					CardType:        toPtr[models.CardType](models.CardType_VISA),
					ExpirationMonth: models.ToPointer(models.PaymentProfileAttributesExpirationMonthContainer.FromNumber(1)),
					ExpirationYear:  models.ToPointer(models.PaymentProfileAttributesExpirationYearContainer.FromNumber(time.Now().Year() + 1)),
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
