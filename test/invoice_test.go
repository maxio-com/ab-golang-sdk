package test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
)

func (s *APISuite) TestInvoice() {
	ctx := context.Background()

	customer := s.createCustomer(ctx)
	productFamily := s.createProductFamily(ctx)
	product := s.createProduct(ctx, *productFamily.Id)
	coupon := s.createCoupon(ctx, *productFamily.Id)

	subscription := s.newSubscription(customer, product, "", []models.CreateSubscriptionComponent{})

	sub, err := s.client.SubscriptionsController().CreateSubscription(ctx, &models.CreateSubscriptionRequest{Subscription: subscription})
	s.NoError(err)
	s.Equal(http.StatusCreated, sub.Response.StatusCode)

	cases := []struct {
		name    string
		assert  func(t *testing.T, resp models.ApiResponse[models.InvoiceResponse], invoice models.CreateInvoice, err error)
		invoice models.CreateInvoice
	}{
		{
			name: "valid",
			invoice: models.CreateInvoice{
				Coupons: []models.CreateInvoiceCoupon{
					{
						Code: coupon.Code,
					},
				},
				LineItems: []models.CreateInvoiceItem{
					{
						Title:     strPtr("A product"),
						Quantity:  interfacePtr(12),
						UnitPrice: interfacePtr("150"),
					},
				},
				IssueDate: timePtr(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
				BillingAddress: &models.CreateInvoiceAddress{
					FirstName: &s.baseBillingAddress.firstName,
					LastName:  &s.baseBillingAddress.lastName,
					Phone:     &s.baseBillingAddress.phone,
					Address:   &s.baseBillingAddress.address,
					Country:   &s.baseBillingAddress.country,
					Zip:       &s.baseBillingAddress.zip,
					City:      &s.baseBillingAddress.city,
				},
			},
			assert: func(t *testing.T, resp models.ApiResponse[models.InvoiceResponse], invoice models.CreateInvoice, err error) {
				s.NoError(err)
				s.Equal(http.StatusCreated, resp.Response.StatusCode)

				inv := resp.Data.Invoice

				s.Equal(s.baseBillingAddress.address, *inv.BillingAddress.Street.Value())
				s.Equal(s.baseBillingAddress.city, *inv.BillingAddress.City.Value())
				s.Equal(s.baseBillingAddress.country, *inv.BillingAddress.Country.Value())

				s.Equal(customer.Id, inv.CustomerId)
				s.Equal(customer.Email, inv.Customer.Email)
				s.Equal(customer.FirstName, inv.Customer.FirstName)
				s.Equal(customer.LastName, inv.Customer.LastName)

				s.Len(inv.Discounts, 1)
				s.Equal(inv.Discounts[0].Code, coupon.Code)

				s.Len(inv.LineItems, 1)

				void, err := s.client.InvoicesController().VoidInvoice(ctx, *resp.Data.Invoice.Uid, &models.VoidInvoiceRequest{
					Void: models.VoidInvoice{Reason: "Duplicate invoice"},
				})

				s.NoError(err)
				s.Equal(http.StatusCreated, void.Response.StatusCode)

				events, err := s.client.InvoicesController().ListInvoiceEvents(
					ctx,
					nil,
					nil,
					nil,
					nil,
					resp.Data.Invoice.Uid,
					nil,
					nil, // cant pass event types here. Server throws 500
				)

				s.NoError(err)
				s.Equal(http.StatusOK, events.Response.StatusCode)
				s.Greater(len(events.Data.Events), 0)
			},
		},
		{
			name: "invalid invoice",
			invoice: models.CreateInvoice{
				Coupons: []models.CreateInvoiceCoupon{
					{
						Code: coupon.Code,
					},
				},
				LineItems: []models.CreateInvoiceItem{
					{
						Title:            strPtr("A product"),
						Quantity:         interfacePtr(12),
						UnitPrice:        interfacePtr("150"),
						PeriodRangeStart: strPtr(dateFromTime(time.Now())),
						PeriodRangeEnd:   strPtr(dateFromTime(time.Now().AddDate(-1, 0, 0))),
					},
				},
				IssueDate: timePtr(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
				BillingAddress: &models.CreateInvoiceAddress{
					FirstName: &s.baseBillingAddress.firstName,
					LastName:  &s.baseBillingAddress.lastName,
					Phone:     &s.baseBillingAddress.phone,
					Address:   &s.baseBillingAddress.address,
					Country:   &s.baseBillingAddress.country,
					Zip:       &s.baseBillingAddress.zip,
					City:      &s.baseBillingAddress.city,
				},
			},
			assert: func(t *testing.T, resp models.ApiResponse[models.InvoiceResponse], invoice models.CreateInvoice, err error) {
				s.Equal(err.Error(), errors.NewErrorArrayMapResponse(422, "Unprocessable Entity (WebDAV)").Error())
				s.Equal(http.StatusUnprocessableEntity, resp.Response.StatusCode)

				_, ok := err.(*errors.ErrorArrayMapResponse)
				s.True(ok)
				// cant check list of errors because map is nil
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			resp, err := s.client.InvoicesController().CreateInvoice(
				ctx,
				*sub.Data.Subscription.Id,
				&models.CreateInvoiceRequest{
					Invoice: c.invoice,
				})

			c.assert(t, resp, c.invoice, err)
		})
	}
}
