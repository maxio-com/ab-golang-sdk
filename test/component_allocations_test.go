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
		fmt.Sprint(*pf.Id),
		&models.CreateOnOffComponent{
			OnOffComponent: models.OnOffComponent{
				Name: "OnOff component",
				PricePoints: []models.ComponentPricePointItem{
					{
						Name: strPtr(s.fkr.RandomStringWithLength(20)),
						Prices: []models.Price{
							{
								StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(1),
								UnitPrice:        models.PriceUnitPriceContainer.FromPrecision(1),
							},
						},
					},
				},
				UnitPrice: models.ToPointer(models.OnOffComponentUnitPriceContainer.FromString("100")),
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
	s.Equal(http.StatusCreated, onOffResp.Response.StatusCode)

	quantityResp, err := s.client.ComponentsController().CreateQuantityBasedComponent(
		ctx,
		fmt.Sprint(*pf.Id),
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
								StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(1),
								UnitPrice:        models.PriceUnitPriceContainer.FromPrecision(1),
							},
						},
					},
				},
				UnitPrice: models.ToPointer(models.QuantityBasedComponentUnitPriceContainer.FromString("100")),
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
				q0, ok := prevAllocations[0].Quantity.AsNumber()
				s.Equal(true, ok)
				s.Equal(int(1), *q0)
				q1, ok := prevAllocations[1].Quantity.AsString()
				s.Equal(true, ok)
				s.Equal(fmt.Sprintf("%.1f", expected[1].Quantity), *q1)

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
				r0, ok := responseAllocations[0].Allocation.Quantity.AsNumber()
				s.Equal(true, ok)
				s.Equal(int(1), *r0)
				r1, ok := responseAllocations[1].Allocation.Quantity.AsString()
				s.Equal(true, ok)
				s.Equal(fmt.Sprintf("%.1f", expected[1].Quantity), *r1)
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
				s.Equal("ComponentAllocationError occured: HTTP Response Not OK. Status code: 422. Response: "+
					fmt.Sprintf("'{\"errors\":[{\"kind\":\"allocation\",\"component_id\":%d,\"on\":\"quantity\",", *onOffResp.Data.Component.Id)+
					"\"message\":\"Quantity: must be either 1 (on) or 0 (off).\"}]}'.",
					err.Error())

				actualErr, ok := err.(*errors.ComponentAllocationError)
				s.Equal(http.StatusUnprocessableEntity, actualErr.StatusCode)
				s.True(ok)
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
