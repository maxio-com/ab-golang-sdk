package test

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/maxio-com/ab-golang-sdk/models"
)

func (s *APISuite) createCustomer(ctx context.Context) models.Customer {
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

	s.NoErrorf(err, "create customer err")
	s.Equalf(http.StatusCreated, resp.Response.StatusCode, "create customer status code")

	return resp.Data.Customer
}

func (s *APISuite) createProductFamily(ctx context.Context) models.ProductFamily {
	resp, err := s.client.ProductFamiliesController().CreateProductFamily(ctx, &models.CreateProductFamilyRequest{
		ProductFamily: models.CreateProductFamily{
			Name: strPtr(s.fkr.Company().Name()),
		},
	})

	s.NoErrorf(err, "create product family err")
	s.Equalf(http.StatusCreated, resp.Response.StatusCode, "create product family status code")

	return *resp.Data.ProductFamily
}

func (s *APISuite) createProduct(ctx context.Context, productFamilyID int) models.Product {
	resp, err := s.client.ProductsController().CreateProduct(ctx, productFamilyID, &models.CreateOrUpdateProductRequest{
		Product: models.CreateOrUpdateProduct{
			Name:         s.fkr.RandomStringWithLength(30),
			Description:  "Testable product",
			PriceInCents: 50,
			Interval:     1,
			IntervalUnit: models.IntervalUnit_MONTH,
		},
	})

	s.NoErrorf(err, "create product err")
	s.Equalf(http.StatusCreated, resp.Response.StatusCode, "create product status code")

	return resp.Data.Product
}

func (s *APISuite) createCoupon(ctx context.Context, productFamilyID int) models.Coupon {
	coupon := &models.Coupon{
		Name:                        strPtr("100\\% off first month of usage"),
		Code:                        strPtr("100OFF" + s.fkr.RandomStringWithLength(30)),
		Description:                 strPtr("100\\% off one-time"),
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

	s.NoErrorf(err, "create coupon err")
	s.Equalf(http.StatusCreated, resp.Response.StatusCode, "create coupon err")

	return *resp.Data.Coupon
}

type MeteredComponent struct {
	Name          string  `json:"name"`
	UnitName      string  `json:"unit_name"`
	Taxable       bool    `json:"taxable"`
	PricingScheme string  `json:"pricing_scheme"`
	Prices        []Price `json:"prices"`
}

type Price struct {
	StartingQuantity int `json:"starting_quantity"`
	UnitPrice        int `json:"unit_price"`
}

func (s *APISuite) createMeteredComponent(ctx context.Context, productFamilyID int) models.Component {
	component := struct {
		MeteredComponent MeteredComponent `json:"metered_component"`
	}{
		MeteredComponent: MeteredComponent{
			Name:          "test 2",
			UnitName:      "test message",
			Taxable:       false,
			PricingScheme: "stairstep",
			Prices: []Price{{
				StartingQuantity: 1,
				UnitPrice:        1,
			},
			},
		},
	}

	resp, err := s.client.ComponentsController().CreateComponent(
		ctx,
		productFamilyID,
		models.ComponentKindPath_METEREDCOMPONENTS,
		interfacePtr(component),
	)

	s.NoErrorf(err, "create component err")
	s.Equalf(http.StatusCreated, resp.Response.StatusCode, "create component err")

	return resp.Data.Component
}

func strPtr(v string) *string {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}

func intPtr(v int) *int {
	return &v
}

func interfacePtr(v interface{}) *interface{} {
	return &v
}

func toPtr[T any](v T) *T {
	return &v
}

func newDate() string {
	t := time.Now().Add(time.Hour)

	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}

func timePtr(v time.Time) *time.Time {
	return &v
}
