package test

import (
	"context"
	"net/http"
	"testing"

	"advancedbilling/models"

	"github.com/stretchr/testify/suite"
)

type PaymentProfilesSuite struct {
	APISuite
}

func TestPaymentProfilesSuite(t *testing.T) {
	suite.Run(t, new(PaymentProfilesSuite))
}

func (s *PaymentProfilesSuite) TestPaymentProfiles() {
	ctx := context.Background()

	customer := s.generateCustomer(ctx)
	pf := s.generateProductFamily(ctx)
	product := s.generateProduct(ctx, *pf.Id)

	paymentProfileRes, err := s.client.PaymentProfilesController().CreatePaymentProfile(
		ctx,
		&models.CreatePaymentProfileRequest{
			PaymentProfile: models.CreatePaymentProfile{
				CustomerId:      customer.Id,
				ExpirationMonth: models.ToPointer(models.CreatePaymentProfileExpirationMonthContainer.FromNumber(12)),
				ExpirationYear:  models.ToPointer(models.CreatePaymentProfileExpirationYearContainer.FromNumber(2030)),
				FullNumber:      models.ToPointer("4111111111111111"),
			},
		},
	)
	s.NoError(err)
	s.Equal(http.StatusCreated, paymentProfileRes.Response.StatusCode)

	payCred, ok := paymentProfileRes.Data.PaymentProfile.AsCreditCardPaymentProfile()

	s.Equal(true, ok)

	subscriptionRes, err := s.client.SubscriptionsController().CreateSubscription(ctx, &models.CreateSubscriptionRequest{
		Subscription: models.CreateSubscription{
			ProductId:  product.Id,
			CustomerId: customer.Id,
		},
	})
	s.NoError(err)
	s.Equal(http.StatusCreated, paymentProfileRes.Response.StatusCode)

	paymentProfileNewRes, err := s.client.PaymentProfilesController().ChangeSubscriptionDefaultPaymentProfile(ctx, *subscriptionRes.Data.Subscription.Id, *payCred.Id)
	s.NoError(err)
	payCredNew, ok := paymentProfileNewRes.Data.PaymentProfile.AsCreditCardPaymentProfile()
	s.NotNil(payCredNew)
	s.Equal(true, ok)

	paymentProfileNewNewRes, err := s.client.PaymentProfilesController().ChangeSubscriptionDefaultPaymentProfile(ctx, *subscriptionRes.Data.Subscription.Id, *payCred.Id)

	s.Error(err)
	s.NotNil(paymentProfileNewNewRes)
	s.Equal(http.StatusUnprocessableEntity, paymentProfileRes.Response.StatusCode)
}
