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

type OnOffComponent struct {
	Name          string       `json:"name"`
	Prices        []Price      `json:"prices"`
	PricingScheme string       `json:"pricing_scheme"`
	PricePoint    []PricePoint `json:"price_points"`
	UnitPrice     string       `json:"unit_price"`
}

type PricePoint struct {
	Name          string  `json:"name"`
	Prices        []Price `json:"prices"`
	PricingScheme string  `json:"pricing_scheme,omitempty"`
}

type OnOffBody struct {
	OnOffComponent OnOffComponent `json:"on_off_component"`
}

type QuantityBasedComponent struct {
	Name                      string       `json:"name"`
	UnitName                  string       `json:"unit_name"`
	PricingScheme             string       `json:"pricing_scheme"`
	AllowFractionalQuantities bool         `json:"allow_fractional_quantities"`
	PricePoint                []PricePoint `json:"price_points"`
	UnitPrice                 string       `json:"unit_price"`
}

type QuantityBody struct {
	QuantityBasedComponent QuantityBasedComponent `json:"quantity_based_component"`
}

func (s *ComponentAlocationSuite) TestComponentAllocations() {
	ctx := context.Background()

	customer := s.createCustomer(ctx)
	pf := s.createProductFamily(ctx)
	product := s.createProduct(ctx, *pf.Id)
	subscription := s.newSubscription(customer, product, "", []models.CreateSubscriptionComponent{})
	subResp, err := s.client.SubscriptionsController().CreateSubscription(ctx, &models.CreateSubscriptionRequest{Subscription: subscription})
	s.NoError(err)
	s.Equal(http.StatusCreated, subResp.Response.StatusCode)

	onOffResp, err := s.client.ComponentsController().CreateComponent(
		ctx,
		*pf.Id,
		models.ComponentKindPath_ONOFFCOMPONENTS,
		interfacePtr(OnOffBody{
			OnOffComponent: OnOffComponent{
				Name: "One of component",
				Prices: []Price{
					{
						StartingQuantity: 1,
						UnitPrice:        1,
					},
				},
				PricePoint: []PricePoint{
					{
						Name: "test price point",
						Prices: []Price{
							{
								StartingQuantity: 1,
								UnitPrice:        1,
							},
						},
					},
				},
				UnitPrice: "100",
			},
		}),
	)
	s.NoError(err)
	s.Equal(http.StatusCreated, onOffResp.Response.StatusCode)

	quantityResp, err := s.client.ComponentsController().CreateComponent(
		ctx,
		*pf.Id,
		models.ComponentKindPath_QUANTITYBASEDCOMPONENTS,
		interfacePtr(QuantityBody{
			QuantityBasedComponent: QuantityBasedComponent{
				Name:                      "Quantity based",
				UnitName:                  "test unit",
				PricingScheme:             string(models.PricingScheme_PERUNIT),
				AllowFractionalQuantities: true,
				PricePoint: []PricePoint{
					{
						Name: "Other price point",
						Prices: []Price{
							{
								StartingQuantity: 1,
								UnitPrice:        1,
							},
						},
						PricingScheme: string(models.PricingScheme_PERUNIT),
					},
				},
				UnitPrice: "100",
			},
		}),
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
					*subResp.Data.Subscription.Id,
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
				*subResp.Data.Subscription.Id,
				&models.PreviewAllocationsRequest{
					Allocations: c.allocations,
				},
			)

			c.assert(t, previewResp, c.allocations, err)
		})
	}
}
