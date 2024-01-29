package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
	"github.com/stretchr/testify/suite"
)

type ComponentAlocationSuite struct {
	APISuite
}

func TestComponentSuite(t *testing.T) {
	suite.Run(t, new(ComponentAlocationSuite))
}

func (s *ComponentAlocationSuite) TestComponentAllocations() {
	ctx := context.Background()

	customer := s.generateCustomer(ctx)
	pf := s.generateProductFamily(ctx)
	product := s.generateProduct(ctx, *pf.Id)
	subscription := s.generateSubscription(ctx, customer, product, "", []models.CreateSubscriptionComponent{})

	onOffResp, err := s.client.ComponentsController().CreateOnOffComponent(
		ctx,
		*pf.Id,
		&models.CreateOnOffComponent{
			OnOffComponent: models.OnOffComponent{
				Name: "OnOff component",
				PricePoints: []models.ComponentPricePointItem{
					{
						Name: strPtr(s.fkr.RandomStringWithLength(20)),
						Prices: []models.Price{
							{
								StartingQuantity: 1,
								UnitPrice:        1,
							},
						},
					},
				},
				UnitPrice: interfacePtr("100"),
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
	s.Equal(http.StatusCreated, onOffResp.Response.StatusCode)

	quantityResp, err := s.client.ComponentsController().CreateQuantityBasedComponent(
		ctx,
		*pf.Id,
		&models.CreateQuantityBasedComponent{
			QuantityBasedComponent: models.QuantityBasedComponent{
				Name:                      s.fkr.RandomStringWithLength(20),
				UnitName:                  s.fkr.RandomStringWithLength(10),
				PricingScheme:             models.PricingScheme_PERUNIT,
				AllowFractionalQuantities: boolPtr(true),
				PricePoints: []models.ComponentPricePointItem{
					{
						Name:          strPtr(s.fkr.RandomStringWithLength(10)),
						PricingScheme: toPtr[models.PricingScheme](models.PricingScheme_PERUNIT),
						Prices: []models.Price{
							{
								StartingQuantity: 1,
								UnitPrice:        1,
							},
						},
					},
				},
				UnitPrice: interfacePtr("100"),
			},
		},
	)
	s.NoError(err)
	s.Equal(http.StatusCreated, quantityResp.Response.StatusCode)

	cases := []struct {
		name        string
		allocations []models.CreateAllocation
		assert      func(t *testing.T, response models.ApiResponse[models.AllocationPreviewResponse], input []models.CreateAllocation, err error)
	}{
		{
			name: "valid",
			allocations: []models.CreateAllocation{
				{
					Quantity:    1,
					ComponentId: onOffResp.Data.Component.Id,
					Memo:        strPtr("foo"),
				},
				{
					Quantity:    10.3,
					ComponentId: quantityResp.Data.Component.Id,
					Memo:        strPtr("bar"),
				},
			},
			assert: func(t *testing.T, response models.ApiResponse[models.AllocationPreviewResponse], expected []models.CreateAllocation, err error) {
				s.NoError(err)
				s.Equal(http.StatusOK, response.Response.StatusCode)

				prevAllocations := response.Data.AllocationPreview.Allocations
				s.Equal(float64(1), *prevAllocations[0].Quantity)
				s.Equal(fmt.Sprintf("%.1f", expected[1].Quantity), *prevAllocations[1].Quantity)

				allocationsResp, err := s.client.SubscriptionComponentsController().AllocateComponents(
					ctx,
					*subscription.Id,
					&models.AllocateComponents{
						Allocations: expected,
					},
				)
				s.NoError(err)
				s.Equal(http.StatusCreated, allocationsResp.Response.StatusCode)

				responseAllocations := allocationsResp.Data
				s.Equal(float64(1), *responseAllocations[0].Allocation.Quantity)
				s.Equal(fmt.Sprintf("%.1f", expected[1].Quantity), *responseAllocations[1].Allocation.Quantity)
			},
		},
		{
			name: "unprocessable entity",
			allocations: []models.CreateAllocation{
				{
					Quantity:    50,
					ComponentId: onOffResp.Data.Component.Id,
				},
			},
			assert: func(t *testing.T, response models.ApiResponse[models.AllocationPreviewResponse], input []models.CreateAllocation, err error) {
				s.Equal(errors.NewComponentAllocationError(422, "Unprocessable Entity (WebDAV)").Error(), err.Error())
				// cant check errors map here because it comes empty
				s.Equal(http.StatusUnprocessableEntity, response.Response.StatusCode)
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			previewResp, err := s.client.SubscriptionComponentsController().PreviewAllocations(
				ctx,
				*subscription.Id,
				&models.PreviewAllocationsRequest{
					Allocations: c.allocations,
				},
			)

			c.assert(t, previewResp, c.allocations, err)
		})
	}
}
