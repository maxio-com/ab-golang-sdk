package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/maxio-com/ab-golang-sdk/models"
)

// this structure has to be defined at this point since default json
// marshaller does not work with marshalling enum as optional interface
// for single metafield requests.
type metafield struct {
	Id        *int                   `json:"id,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Scope     *models.MetafieldScope `json:"scope,omitempty"`
	DataCount *int                   `json:"data_count,omitempty"`
	InputType *string                `json:"input_type,omitempty"`
	Enum      []string               `json:"enum"`
}

func (s *APISuite) TestMetafields() {
	ctx := context.Background()

	customer := s.createCustomer(ctx)
	productFamilyID := s.createProductFamily(ctx)
	product := s.createProduct(ctx, *productFamilyID.Id)
	subscription := s.newSubscription(customer, product, "", []models.CreateSubscriptionComponent{})

	radioFieldName := s.fkr.RandomStringWithLength(20)
	dropdownFieldName := s.fkr.RandomStringWithLength(20)
	textFieldName := s.fkr.RandomStringWithLength(20)

	resp, err := s.client.SubscriptionsController().CreateSubscription(ctx, &models.CreateSubscriptionRequest{
		Subscription: subscription,
	})
	s.NoError(err)
	s.Equal(http.StatusCreated, resp.Response.StatusCode)

	subs := resp.Data.Subscription

	cases := []struct {
		name         string
		resourceType models.ResourceType
		metafields   interface{}
		assert       func(t *testing.T, resp models.ApiResponse[[]models.Metafield], err error)
	}{
		{
			name:         "subscriptions",
			resourceType: models.ResourceType_SUBSCRIPTIONS,
			metafields: []models.Metafield{
				{
					Name:      &dropdownFieldName,
					InputType: strPtr("dropdown"),
					Scope: &models.MetafieldScope{
						PublicShow: toPtr[models.IncludeOption](models.IncludeOption_INCLUDE),
						PublicEdit: toPtr[models.IncludeOption](models.IncludeOption_INCLUDE),
					},
					Enum: models.NewOptional[interface{}](interfacePtr([]string{
						"option 1",
						"option 2",
					})),
				},
				{
					Name: &textFieldName,
				},
			},
			assert: func(t *testing.T, resp models.ApiResponse[[]models.Metafield], err error) {
				s.NoError(err)
				s.Contains([]int{http.StatusOK, http.StatusCreated}, resp.Response.StatusCode)

				dropdownField := resp.Data[0]
				s.Equal(dropdownFieldName, *dropdownField.Name)
				s.Equal("dropdown", *dropdownField.InputType)
				// s.Len(dropdownField.Enum.Value(), 2) unusable at this stage
				s.Equal(models.IncludeOption_INCLUDE, *dropdownField.Scope.PublicShow)
				s.Equal(models.IncludeOption_INCLUDE, *dropdownField.Scope.PublicEdit)

				textField := resp.Data[1]
				s.Equal(textFieldName, *textField.Name)
				s.Equal("text", *textField.InputType)

				r, err := s.client.CustomFieldsController().CreateMetadata(
					ctx,
					models.ResourceType_SUBSCRIPTIONS,
					fmt.Sprintf("%d", *subs.Id),
					&models.CreateMetadataRequest{
						Metadata: []models.CreateMetadata{
							{
								Name:  dropdownField.Name,
								Value: strPtr("option 1"),
							},
							{
								Name:  textField.Name,
								Value: strPtr("something"),
							},
						},
					})
				s.NoError(err)
				s.Equal(http.StatusOK, r.Response.StatusCode)

				rSubs, err := s.client.SubscriptionsController().ListSubscriptions(
					ctx,
					nil,
					nil,
					nil,
					product.Id,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					map[string]string{
						dropdownFieldName: "option 1",
					},
					nil,
					nil,
					[]models.SubscriptionListInclude{},
				)
				s.NoError(err)
				s.Equal(http.StatusOK, r.Response.StatusCode)

				s.Len(rSubs.Data, 1)
				s.Equal(subs.Id, rSubs.Data[0].Subscription.Id)
			},
		},
		{
			name:         "customers",
			resourceType: models.ResourceType_CUSTOMERS,
			metafields: metafield{
				Name:      &radioFieldName,
				InputType: strPtr("radio"),
				Enum: []string{
					"option 1",
					"option 2",
				},
				Scope: &models.MetafieldScope{
					Csv:      toPtr[models.IncludeOption](models.IncludeOption_INCLUDE),
					Invoices: toPtr[models.IncludeOption](models.IncludeOption_INCLUDE),
					Portal:   toPtr[models.IncludeOption](models.IncludeOption_INCLUDE),
				},
			},
			assert: func(t *testing.T, resp models.ApiResponse[[]models.Metafield], err error) {
				s.NoError(err)
				s.Contains([]int{http.StatusOK, http.StatusCreated}, resp.Response.StatusCode)

				radioField := resp.Data[0]
				s.Equal(radioFieldName, *radioField.Name)
				s.Equal("radio", *radioField.InputType)

				s.Equal(models.IncludeOption_INCLUDE, *radioField.Scope.Csv)
				s.Equal(models.IncludeOption_INCLUDE, *radioField.Scope.Invoices)
				s.Equal(models.IncludeOption_INCLUDE, *radioField.Scope.Portal)

				customerResp, err := s.client.CustomFieldsController().CreateMetadata(
					ctx,
					models.ResourceType_CUSTOMERS,
					fmt.Sprintf("%d", *customer.Id),
					&models.CreateMetadataRequest{
						Metadata: []models.CreateMetadata{
							{
								Name:  radioField.Name,
								Value: strPtr("option 2"),
							},
						},
					},
				)
				s.NoError(err)
				s.Equal(http.StatusOK, customerResp.Response.StatusCode)

				s.Len(customerResp.Data, 1)
				s.Equal(customer.Id, customerResp.Data[0].ResourceId)
				s.Equal(radioField.Name, customerResp.Data[0].Name)
				s.Equal("option 2", *customerResp.Data[0].Value)
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			resp, err := s.client.CustomFieldsController().CreateMetafields(
				ctx,
				c.resourceType,
				&models.CreateMetafieldsRequest{
					Metafields: c.metafields,
				},
			)

			c.assert(t, resp, err)
		})
	}
}
