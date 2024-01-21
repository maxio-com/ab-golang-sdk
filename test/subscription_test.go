package test

import (
	"context"
	"net/http"
	"testing"
	"time"

	advancedbilling "github.com/maxio-com/ab-golang-sdk"
	"github.com/maxio-com/ab-golang-sdk/models"
)

func (s *APISuite) TestSubscriptionCreate() {
	ctx := context.Background()

	customer := s.createCustomer(ctx)
	productFamily := s.createProductFamily(ctx)
	product := s.createProduct(ctx, *productFamily.Id)
	coupon := s.createCoupon(ctx, *productFamily.Id)
	component := s.createMeteredComponent(ctx, *productFamily.Id)

	cases := []struct {
		name         string
		client       advancedbilling.ClientInterface
		subscription models.CreateSubscription
		assert       func(*testing.T, models.ApiResponse[models.SubscriptionResponse], models.CreateSubscription, error)
	}{
			subscription: s.newSubscription(
				customer,
				product,
				*coupon.Code,
				[]models.CreateSubscriptionComponent{
					{
						ComponentId: interfacePtr(*component.Id),
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
					*createdSubscription.Id,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
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
				s.Equal(http.StatusUnprocessableEntity, ar.Response.StatusCode)
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
						ComponentId: interfacePtr(*component.Id),
						Enabled:     boolPtr(true),
						UnitBalance: intPtr(1),
					},
				},
			),
			assert: func(t *testing.T, ar models.ApiResponse[models.SubscriptionResponse], cs models.CreateSubscription, err error) {
				s.Equal(http.StatusUnauthorized, ar.Response.StatusCode)
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
		PaymentCollectionMethod: toPtr[models.PaymentCollectionMethod](models.PaymentCollectionMethod_AUTOMATIC),
		CustomerId:              customer.Id,
		Currency:                strPtr("USD"),
		InitialBillingAt:        strPtr("2029-08-29T12:00:00-04:00"),
		CouponCode:              &couponCode,
		Components:              components,
		PaymentProfileAttributes: &models.PaymentProfileAttributes{
			FirstName:       strPtr("Joe"),
			LastName:        strPtr("Smith"),
			FullNumber:      strPtr("4111111111111111"),
			CardType:        toPtr[models.CardType](models.CardType_VISA),
			ExpirationMonth: interfacePtr(1),
			ExpirationYear:  interfacePtr(time.Now().Year() + 1),
			BillingAddress:  strPtr("123 Mass Ave."),
			BillingCity:     strPtr("Boston"),
			BillingState:    strPtr("MA"),
			BillingCountry:  strPtr("US"),
			BillingZip:      strPtr("02120"),
		},
	}
}
