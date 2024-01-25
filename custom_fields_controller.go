package advancedbilling

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
	"net/http"
)

// CustomFieldsController represents a controller struct.
type CustomFieldsController struct {
	baseController
}

// NewCustomFieldsController creates a new instance of CustomFieldsController.
// It takes a baseController as a parameter and returns a pointer to the CustomFieldsController.
func NewCustomFieldsController(baseController baseController) *CustomFieldsController {
	customFieldsController := CustomFieldsController{baseController: baseController}
	return &customFieldsController
}

// CreateMetafields takes context, resourceType, body as parameters and
// returns an models.ApiResponse with []models.Metafield data and
// an error if there was an issue with the request or response.
// ## Custom Fields: Metafield Intro
// **Chargify refers to Custom Fields in the API documentation as metafields and metadata.** Within the Chargify UI, metadata and metafields are grouped together under the umbrella of "Custom Fields." All of our UI-based documentation that references custom fields will not cite the terminology metafields or metadata.
// + **Metafield is the custom field**
// + **Metadata is the data populating the custom field.**
// Chargify Metafields are used to add meaningful attributes to subscription and customer resources. Full documentation on how to create Custom Fields in the Chargify UI can be located [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405332553613-Custom-Fields-Reference). For additional documentation on how to record data within custom fields, please see our subscription-based documentation [here.](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404434903181-Subscription-Summary#custom-fields)
// Metafield are the place where you will set up your resource to accept additional data. It is scoped to the site instead of a specific customer or subscription. Think of it as the key, and Metadata as the value on every record.
// ## Create Metafields
// Use this endpoint to create metafields for your Site. Metafields can be populated with metadata after the fact.
// Each site is limited to 100 unique Metafields (i.e. keys, or names) per resource. This means you can have 100 Metafields for Subscription and another 100 for Customer.
// ### Metafields "On-the-Fly"
// It is possible to create Metafields “on the fly” when you create your Metadata – if a non-existant name is passed when creating Metadata, a Metafield for that key will be automatically created. The Metafield API, however, gives you more control over your “keys”.
// ### Metafield Scope Warning
// If configuring metafields in the Admin UI or via the API, be careful sending updates to metafields with the scope attribute – **if a partial update is sent it will overwrite the current configuration**.
func (c *CustomFieldsController) CreateMetafields(
	ctx context.Context,
	resourceType models.ResourceType,
	body *models.CreateMetafieldsRequest) (
	models.ApiResponse[[]models.Metafield],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/%v/metafields.json", resourceType),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result []models.Metafield
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Metafield](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 422 {
		err = errors.NewSingleErrorResponse(422, "Unprocessable Entity (WebDAV)")
	}
	return models.NewApiResponse(result, resp), err
}

// ListMetafields takes context, resourceType, name, page, perPage, direction as parameters and
// returns an models.ApiResponse with models.ListMetafieldsResponse data and
// an error if there was an issue with the request or response.
// This endpoint lists metafields associated with a site. The metafield description and usage is contained in the response.
func (c *CustomFieldsController) ListMetafields(
	ctx context.Context,
	resourceType models.ResourceType,
	name *string,
	page *int,
	perPage *int,
	direction *models.SortingDirection) (
	models.ApiResponse[models.ListMetafieldsResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/%v/metafields.json", resourceType),
	)
	req.Authenticate(true)
	if name != nil {
		req.QueryParam("name", *name)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if direction != nil {
		req.QueryParam("direction", *direction)
	}

	var result models.ListMetafieldsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListMetafieldsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateMetafield takes context, resourceType, body as parameters and
// returns an models.ApiResponse with []models.Metafield data and
// an error if there was an issue with the request or response.
// Use the following method to update metafields for your Site. Metafields can be populated with metadata after the fact.
func (c *CustomFieldsController) UpdateMetafield(
	ctx context.Context,
	resourceType models.ResourceType,
	body *models.UpdateMetafieldsRequest) (
	models.ApiResponse[[]models.Metafield],
	error) {
	req := c.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/%v/metafields.json", resourceType),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result []models.Metafield
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Metafield](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteMetafield takes context, resourceType, name as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Use the following method to delete a metafield. This will remove the metafield from the Site.
// Additionally, this will remove the metafield and associated metadata with all Subscriptions on the Site.
func (c *CustomFieldsController) DeleteMetafield(
	ctx context.Context,
	resourceType models.ResourceType,
	name *string) (
	*http.Response,
	error) {
	req := c.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/%v/metafields.json", resourceType),
	)
	req.Authenticate(true)
	if name != nil {
		req.QueryParam("name", *name)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	if context.Response.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return context.Response, err
}

// CreateMetadata takes context, resourceType, resourceId, body as parameters and
// returns an models.ApiResponse with []models.Metadata data and
// an error if there was an issue with the request or response.
// ## Custom Fields: Metadata Intro
// **Chargify refers to Custom Fields in the API documentation as metafields and metadata.** Within the Chargify UI, metadata and metafields are grouped together under the umbrella of "Custom Fields." All of our UI-based documentation that references custom fields will not cite the terminology metafields or metadata.
// + **Metafield is the custom field**
// + **Metadata is the data populating the custom field.**
// Chargify Metafields are used to add meaningful attributes to subscription and customer resources. Full documentation on how to create Custom Fields in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407659856411). For additional documentation on how to record data within custom fields, please see our subscription-based documentation [here.](https://chargify.zendesk.com/hc/en-us/articles/4407884887835#custom-fields)
// Metadata is associated to a customer or subscription, and corresponds to a Metafield. When creating a new metadata object for a given record, **if the metafield is not present it will be created**.
// ## Metadata limits
// Metadata values are limited to 2kB in size. Additonally, there are limits on the number of unique metafields available per resource.
// ## Create Metadata
// This method will create a metafield for the site on the fly if it does not already exist, and populate the metadata value.
// ### Subscription or Customer Resource
// Please pay special attention to the resource you use when creating metadata.
func (c *CustomFieldsController) CreateMetadata(
	ctx context.Context,
	resourceType models.ResourceType,
	resourceId string,
	body *models.CreateMetadataRequest) (
	models.ApiResponse[[]models.Metadata],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/%v/%v/metadata.json", resourceType, resourceId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result []models.Metadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Metadata](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 422 {
		err = errors.NewSingleErrorResponse(422, "Unprocessable Entity (WebDAV)")
	}
	return models.NewApiResponse(result, resp), err
}

// ListMetadata takes context, resourceType, resourceId, page, perPage as parameters and
// returns an models.ApiResponse with models.PaginatedMetadata data and
// an error if there was an issue with the request or response.
// This request will list all of the metadata belonging to a particular resource (ie. subscription, customer) that is specified.
// ## Metadata Data
// This endpoint will also display the current stats of your metadata to use as a tool for pagination.
func (c *CustomFieldsController) ListMetadata(
	ctx context.Context,
	resourceType models.ResourceType,
	resourceId string,
	page *int,
	perPage *int) (
	models.ApiResponse[models.PaginatedMetadata],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/%v/%v/metadata.json", resourceType, resourceId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}

	var result models.PaginatedMetadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.PaginatedMetadata](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateMetadata takes context, resourceType, resourceId, body as parameters and
// returns an models.ApiResponse with []models.Metadata data and
// an error if there was an issue with the request or response.
// This method allows you to update the existing metadata associated with a subscription or customer.
func (c *CustomFieldsController) UpdateMetadata(
	ctx context.Context,
	resourceType models.ResourceType,
	resourceId string,
	body *models.UpdateMetadataRequest) (
	models.ApiResponse[[]models.Metadata],
	error) {
	req := c.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/%v/%v/metadata.json", resourceType, resourceId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result []models.Metadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Metadata](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteMetadata takes context, resourceType, resourceId, name, names as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This method removes the metadata from the subscriber/customer cited.
// ## Query String Usage
// For instance if you wanted to delete the metadata for customer 99 named weight you would request:
// ```
// https://acme.chargify.com/customers/99/metadata.json?name=weight
// ```
// If you want to delete multiple metadata fields for a customer 99 named: `weight` and `age` you wrould request:
// ```
// https://acme.chargify.com/customers/99/metadata.json?names[]=weight&names[]=age
// ```
// ## Successful Response
// For a success, there will be a code `200` and the plain text response `true`.
// ## Unsuccessful Response
// When a failed response is encountered, you will receive a `404` response and the plain text response of `true`.
func (c *CustomFieldsController) DeleteMetadata(
	ctx context.Context,
	resourceType models.ResourceType,
	resourceId string,
	name *string,
	names []string) (
	*http.Response,
	error) {
	req := c.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/%v/%v/metadata.json", resourceType, resourceId),
	)
	req.Authenticate(true)
	if name != nil {
		req.QueryParam("name", *name)
	}
	if names != nil {
		req.QueryParam("names[]", names)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	if context.Response.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return context.Response, err
}

// ListMetadataForResourceType takes context, resourceType, page, perPage, dateField, startDate, endDate, startDatetime, endDatetime, withDeleted, resourceIds, direction as parameters and
// returns an models.ApiResponse with models.PaginatedMetadata data and
// an error if there was an issue with the request or response.
// This method will provide you information on usage of metadata across your selected resource (ie. subscriptions, customers)
// ## Metadata Data
// This endpoint will also display the current stats of your metadata to use as a tool for pagination.
// ### Metadata for multiple records
// `https://acme.chargify.com/subscriptions/metadata.json?resource_ids[]=1&resource_ids[]=2`
// ## Read Metadata for a Site
// This endpoint will list the number of pages of metadata information that are contained within a site.
func (c *CustomFieldsController) ListMetadataForResourceType(
	ctx context.Context,
	resourceType models.ResourceType,
	page *int,
	perPage *int,
	dateField *models.BasicDateField,
	startDate *string,
	endDate *string,
	startDatetime *string,
	endDatetime *string,
	withDeleted *bool,
	resourceIds []int,
	direction *models.SortingDirection) (
	models.ApiResponse[models.PaginatedMetadata],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/%v/metadata.json", resourceType),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if dateField != nil {
		req.QueryParam("date_field", *dateField)
	}
	if startDate != nil {
		req.QueryParam("start_date", *startDate)
	}
	if endDate != nil {
		req.QueryParam("end_date", *endDate)
	}
	if startDatetime != nil {
		req.QueryParam("start_datetime", *startDatetime)
	}
	if endDatetime != nil {
		req.QueryParam("end_datetime", *endDatetime)
	}
	if withDeleted != nil {
		req.QueryParam("with_deleted", *withDeleted)
	}
	if resourceIds != nil {
		req.QueryParam("resource_ids[]", resourceIds)
	}
	if direction != nil {
		req.QueryParam("direction", *direction)
	}

	var result models.PaginatedMetadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.PaginatedMetadata](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
