package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/maxio-com/ab-golang-sdk/models"
	"github.com/stretchr/testify/suite"
)

type MetafieldsSuite struct {
	APISuite
}

func TestMetafieldsSuite(t *testing.T) {
	suite.Run(t, new(MetafieldsSuite))
}

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

func (s *MetafieldsSuite) TestMetafields() {
	ctx := context.Background()

	customer := s.generateCustomer(ctx)
	productFamilyID := s.generateProductFamily(ctx)
	product := s.generateProduct(ctx, *productFamilyID.Id)
	subscription := s.generateSubscription(ctx, customer, product, "", []models.CreateSubscriptionComponent{})

	radioFieldName := s.fkr.RandomStringWithLength(20)
	dropdownFieldName := s.fkr.RandomStringWithLength(20)
	textFieldName := s.fkr.RandomStringWithLength(20)

	dropdownInputType := "dropdown"
	radioInputType := "radio"
	textInputType := "text"

	enums := []string{
		"option 1",
		"option 2",
	}

	cases := []struct {
		name         string
		resourceType models.ResourceType
		metafields   interface{}
		assert       func(t *testing.T, resp models.ApiResponse[[]models.Metafield], err error)
		afterTest    func(t *testing.T, metafields []models.Metafield)
	}{
		{
			name:         "subscriptions",
			resourceType: models.ResourceType_SUBSCRIPTIONS,
			metafields: []models.Metafield{
				{
					Name:      &dropdownFieldName,
					InputType: &dropdownInputType,
					Scope: &models.MetafieldScope{
						PublicShow: toPtr[models.IncludeOption](models.IncludeOption_INCLUDE),
						PublicEdit: toPtr[models.IncludeOption](models.IncludeOption_INCLUDE),
					},
					Enum: models.NewOptional[interface{}](interfacePtr(enums)),
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
				s.Equal(dropdownInputType, *dropdownField.InputType)
				// s.Len(dropdownField.Enum.Value(), 2) unusable at this stage
				s.Equal(models.IncludeOption_INCLUDE, *dropdownField.Scope.PublicShow)
				s.Equal(models.IncludeOption_INCLUDE, *dropdownField.Scope.PublicEdit)

				textField := resp.Data[1]
				s.Equal(textFieldName, *textField.Name)
				s.Equal(textInputType, *textField.InputType)

				metadata := []models.CreateMetadata{
					{
						Name:  dropdownField.Name,
						Value: strPtr(enums[0]),
					},
					{
						Name:  textField.Name,
						Value: strPtr("something"),
					},
				}

				r, err := s.client.CustomFieldsController().CreateMetadata(
					ctx,
					models.ResourceType_SUBSCRIPTIONS,
					fmt.Sprintf("%d", *subscription.Id),
					&models.CreateMetadataRequest{
						Metadata: metadata,
					})
				s.NoError(err)
				s.Equal(http.StatusOK, r.Response.StatusCode)

				s.Equal(subscription.Id, r.Data[0].ResourceId)
				s.Len(r.Data, 2)

				s.Equal(metadata[0].Name, r.Data[0].Name)
				s.Equal(*metadata[0].Value, *r.Data[0].Value)
				s.Equal(metadata[1].Name, r.Data[1].Name)
				s.Equal(*metadata[1].Value, *r.Data[1].Value)

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
						dropdownFieldName: enums[0],
					},
					nil,
					nil,
					[]models.SubscriptionListInclude{},
				)
				s.NoError(err)
				s.Equal(http.StatusOK, r.Response.StatusCode)

				s.Len(rSubs.Data, 1)
				s.Equal(subscription.Id, rSubs.Data[0].Subscription.Id)
			},
			afterTest: func(t *testing.T, metafields []models.Metafield) {
				for _, metafield := range metafields {
					resp, err := s.client.CustomFieldsController().DeleteMetafield(
						ctx,
						models.ResourceType_SUBSCRIPTIONS,
						metafield.Name,
					)

					s.NoError(err)
					s.Equal(http.StatusOK, resp.StatusCode)
				}
			},
		},
		{
			name:         "customers",
			resourceType: models.ResourceType_CUSTOMERS,
			metafields: metafield{
				Name:      &radioFieldName,
				InputType: &radioInputType,
				Enum:      enums,
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
				s.Equal(radioInputType, *radioField.InputType)

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
								Value: &enums[1],
							},
						},
					},
				)
				s.NoError(err)
				s.Equal(http.StatusOK, customerResp.Response.StatusCode)

				s.Len(customerResp.Data, 1)
				s.Equal(customer.Id, customerResp.Data[0].ResourceId)
				s.Equal(radioField.Name, customerResp.Data[0].Name)
				s.Equal(enums[1], *customerResp.Data[0].Value)
			},
			afterTest: func(t *testing.T, metafields []models.Metafield) {
				for _, metafield := range metafields {
					resp, err := s.client.CustomFieldsController().DeleteMetafield(
						ctx,
						models.ResourceType_CUSTOMERS,
						metafield.Name,
					)

					s.NoError(err)
					s.Equal(http.StatusOK, resp.StatusCode)
				}
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
			c.afterTest(t, resp.Data)
		})
	}
}
