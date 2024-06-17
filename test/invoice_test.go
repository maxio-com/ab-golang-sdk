package test

import (
	"context"
	"net/http"
	"testing"
	"time"

	advancedbilling "github.com/maxio-com/ab-golang-sdk"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
	"github.com/stretchr/testify/suite"
)

type InvoiceSuite struct {
	APISuite
}

func TestInvoiceSuite(t *testing.T) {
	suite.Run(t, new(InvoiceSuite))
}

func (s *InvoiceSuite) TestInvoice() {
	ctx := context.Background()

	customer := s.generateCustomer(ctx)
	productFamily := s.generateProductFamily(ctx)
	product := s.generateProduct(ctx, *productFamily.Id)
	coupon := s.generateCoupon(ctx, *productFamily.Id)
	subscription := s.generateSubscription(ctx, customer, product, "", []models.CreateSubscriptionComponent{})

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
						Quantity:  models.ToPointer(models.CreateInvoiceItemQuantityContainer.FromPrecision(12)),
						UnitPrice: models.ToPointer(models.CreateInvoiceItemUnitPriceContainer.FromString("150")),
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

				s.Equal(models.InvoiceStatus_OPEN, *inv.Status)

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
				s.Equal(models.InvoiceStatus_VOIDED, *void.Data.Status)

				// sometimes some events are missing
				time.Sleep(2500 * time.Millisecond)

                events, err := s.client.InvoicesController().ListInvoiceEvents(
                    ctx,
                    advancedbilling.ListInvoiceEventsInput{
                        EventTypes: []models.InvoiceEventType{models.InvoiceEventType_ISSUEINVOICE, models.InvoiceEventType_VOIDINVOICE},
                        InvoiceUid: resp.Data.Invoice.Uid,
                    },
                )

                s.NoError(err)
                s.Equal(http.StatusOK, events.Response.StatusCode)
                s.Equal(len(events.Data.Events), 2)

                event1, ok := events.Data.Events[0].AsIssueInvoiceEvent()
                s.True(ok)
                s.Equal(event1.EventType, models.InvoiceEventType_ISSUEINVOICE)
                s.Equal(event1.EventData.TotalAmount, "900.0")

                event2, ok := events.Data.Events[1].AsVoidInvoiceEvent()
                s.True(ok)
                s.Equal(event2.EventType, models.InvoiceEventType_VOIDINVOICE)
                s.Equal(event2.EventData.Reason, "Duplicate invoice")
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
						Quantity:         models.ToPointer(models.CreateInvoiceItemQuantityContainer.FromPrecision(12)),
						UnitPrice:        models.ToPointer(models.CreateInvoiceItemUnitPriceContainer.FromString("150")),
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
				s.Equal(err.Error(), "ErrorArrayMapResponse occured: HTTP Response Not OK. Status code: 422. Response: "+
					"'{\"errors\":{\"line_items[0].period_range_end\":[\"Must be greater or equal to period_range_start.\"]}}'.")

				actualErr, ok := err.(*errors.ErrorArrayMapResponse)
				s.Equal(http.StatusUnprocessableEntity, actualErr.StatusCode)
				s.True(ok)
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			resp, err := s.client.InvoicesController().CreateInvoice(
				ctx,
				*subscription.Id,
				&models.CreateInvoiceRequest{
					Invoice: c.invoice,
				})

			c.assert(t, resp, c.invoice, err)
		})
	}
}
