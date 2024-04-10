package test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/apimatic/go-core-runtime/https"
	advancedbilling "github.com/maxio-com/ab-golang-sdk"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
	"github.com/stretchr/testify/suite"
)

type SubscriptionSuite struct {
	APISuite
}

func TestSubscriptionSuite(t *testing.T) {
	suite.Run(t, new(SubscriptionSuite))
}

func (s *SubscriptionSuite) TestSubscriptionCreate() {
	ctx := context.Background()

	customer := s.generateCustomer(ctx)
	productFamily := s.generateProductFamily(ctx)
	product := s.generateProduct(ctx, *productFamily.Id)
	coupon := s.generateCoupon(ctx, *productFamily.Id)
	component := s.generateMeteredComponent(ctx, *productFamily.Id)

	cases := []struct {
		name         string
		client       advancedbilling.ClientInterface
		subscription models.CreateSubscription
		assert       func(*testing.T, models.ApiResponse[models.SubscriptionResponse], models.CreateSubscription, error)
	}{
		{
			name:   "valid",
			client: s.client,
			subscription: s.newSubscription(
				customer,
				product,
				*coupon.Code,
				[]models.CreateSubscriptionComponent{
					{
						ComponentId: models.ToPointer(models.CreateSubscriptionComponentComponentIdContainer.FromNumber(*component.Id)),
						Enabled:     boolPtr(true),
						UnitBalance: intPtr(1),
					},
				},
			),
			assert: func(t *testing.T, resp models.ApiResponse[models.SubscriptionResponse], subscription models.CreateSubscription, err error) {
				s.NoError(err)
				s.Equal(http.StatusCreated, resp.Response.StatusCode)

				createdSubscription := resp.Data.Subscription
				s.Equal(models.SubscriptionState_AWAITINGSIGNUP, *createdSubscription.State)

				s.Equal(subscription.ProductId, createdSubscription.Product.Id)
				s.Equal(subscription.PaymentCollectionMethod, createdSubscription.PaymentCollectionMethod)
				s.Equal(subscription.CustomerId, createdSubscription.Customer.Id)
				s.Equal(subscription.Currency, createdSubscription.Currency)
				s.Equal(subscription.CouponCode, createdSubscription.CouponCode.Value())
				s.Equal(subscription.CustomerId, createdSubscription.Customer.Id)
				s.Equal(subscription.ProductId, createdSubscription.Product.Id)

				listResp, err := s.client.SubscriptionComponentsController().ListSubscriptionComponents(
					ctx,
					advancedbilling.ListSubscriptionComponentsInput{
						SubscriptionId: *createdSubscription.Id,
					},
				)

				s.NoError(err)
				s.Equal(http.StatusOK, listResp.Response.StatusCode)

				prices := *component.Prices.Value()
				s.Equal(
					*prices[0].StartingQuantity,
					*listResp.Data[0].Component.UnitBalance,
				)

				readSubResp, err := s.client.SubscriptionsController().ReadSubscription(
					ctx,
					*createdSubscription.Id,
					[]models.SubscriptionInclude{
						models.SubscriptionInclude_COUPONS,
					},
				)

				s.NoError(err)
				s.Equal(http.StatusOK, readSubResp.Response.StatusCode)
				s.Len(readSubResp.Data.Subscription.CouponCodes, 1)
				s.Equal(*coupon.Code, readSubResp.Data.Subscription.CouponCodes[0])

				listComponentsFilteredResp1, err := s.client.SubscriptionComponentsController().ListSubscriptionComponentsForSite(
					ctx,
					advancedbilling.ListSubscriptionComponentsForSiteInput{
						SubscriptionIds: []int{*createdSubscription.Id},
						Include: models.ToPointer(models.ListSubscriptionComponentsInclude_SUBSCRIPTION),
						Filter: models.ToPointer(models.ListSubscriptionComponentsForSiteFilter{
						    Subscription: models.ToPointer(models.SubscriptionFilter{
                                DateField: models.ToPointer(models.SubscriptionListDateField_UPDATEDAT),
                                StartDate: timePtr(time.Now().AddDate(-1, 0, 0)),
						    }),
						}),
					},
				)

				s.NoError(err)
                s.Equal(http.StatusOK, listComponentsFilteredResp1.Response.StatusCode)
                s.Len(listComponentsFilteredResp1.Data.SubscriptionsComponents, 1)
                s.Equal(component.Id, listComponentsFilteredResp1.Data.SubscriptionsComponents[0].ComponentId)

                listComponentsFilteredResp2, err := s.client.SubscriptionComponentsController().ListSubscriptionComponentsForSite(
					ctx,
					advancedbilling.ListSubscriptionComponentsForSiteInput{
						SubscriptionIds: []int{*createdSubscription.Id},
						Include: models.ToPointer(models.ListSubscriptionComponentsInclude_SUBSCRIPTION),
						Filter: models.ToPointer(models.ListSubscriptionComponentsForSiteFilter{
						    Subscription: models.ToPointer(models.SubscriptionFilter{
                                DateField: models.ToPointer(models.SubscriptionListDateField_UPDATEDAT),
                                EndDate: timePtr(time.Now().AddDate(-1, 0, 0)),
						    }),
						}),
					},
				)

				s.NoError(err)
                s.Equal(http.StatusOK, listComponentsFilteredResp2.Response.StatusCode)
                s.Len(listComponentsFilteredResp2.Data.SubscriptionsComponents, 0)
			},
		},
		{
			name:   "invalid coupon code",
			client: s.client,
			subscription: s.newSubscription(
				customer,
				product,
				"invalid coupon code",
				[]models.CreateSubscriptionComponent{},
			),
			assert: func(t *testing.T, ar models.ApiResponse[models.SubscriptionResponse], subscription models.CreateSubscription, err error) {
				s.Equal(string("ErrorListResponse occured: HTTP Response Not OK. Status code: 422. Response: "+
					"'{\"errors\":[\"Coupon Code: 'invalid coupon code' - Coupon code could not be found.\"]}'."),
					err.Error())

				actualErr, ok := err.(*errors.ErrorListResponse)
				s.Equal(http.StatusUnprocessableEntity, actualErr.StatusCode)
				s.True(ok)
			},
		},
		{
			name:   "unauthorized",
			client: s.unauthorizedClient,
			subscription: s.newSubscription(
				customer,
				product,
				*coupon.Code,
				[]models.CreateSubscriptionComponent{
					{
						ComponentId: models.ToPointer(models.CreateSubscriptionComponentComponentIdContainer.FromNumber(*component.Id)),
						Enabled:     boolPtr(true),
						UnitBalance: intPtr(1),
					},
				},
			),
			assert: func(t *testing.T, ar models.ApiResponse[models.SubscriptionResponse], cs models.CreateSubscription, err error) {
				s.Equal("ApiError occured: HTTP Response Not OK. Status code: 401. Response: 'HTTP Basic: Access denied.\n'.",
					err.Error())

				actualErr, ok := err.(https.ApiError)
				s.Equal(http.StatusUnauthorized, actualErr.StatusCode)
				s.True(ok)
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			resp, err := c.client.SubscriptionsController().CreateSubscription(
				ctx,
				&models.CreateSubscriptionRequest{
					Subscription: c.subscription,
				},
			)

			c.assert(t, resp, c.subscription, err)
		})
	}
}

func (s *APISuite) newSubscription(
	customer models.Customer,
	product models.Product,
	couponCode string,
	components []models.CreateSubscriptionComponent,
) models.CreateSubscription {
	return models.CreateSubscription{
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
	}
}
