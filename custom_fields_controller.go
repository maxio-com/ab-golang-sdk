package advancedbilling

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
	"net/http"
	"time"
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSingleErrorResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result []models.Metafield
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Metafield](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListMetafieldsInput represents the input of the ListMetafields endpoint.
type ListMetafieldsInput struct {
	// the resource type to which the metafields belong
	ResourceType models.ResourceType
	// filter by the name of the metafield
	Name *string
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
	// Controls the order in which results are returned.
	// Use in query `direction=asc`.
	Direction *models.SortingDirection
}

// ListMetafields takes context, resourceType, name, page, perPage, direction as parameters and
// returns an models.ApiResponse with models.ListMetafieldsResponse data and
// an error if there was an issue with the request or response.
// This endpoint lists metafields associated with a site. The metafield description and usage is contained in the response.
func (c *CustomFieldsController) ListMetafields(
	ctx context.Context,
	input ListMetafieldsInput) (
	models.ApiResponse[models.ListMetafieldsResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/%v/metafields.json", input.ResourceType),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	if input.Name != nil {
		req.QueryParam("name", *input.Name)
	}
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.Direction != nil {
		req.QueryParam("direction", *input.Direction)
	}

	var result models.ListMetafieldsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListMetafieldsResponse](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSingleErrorResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result []models.Metafield
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Metafield](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})
	if name != nil {
		req.QueryParam("name", *name)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
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
	resourceId int,
	body *models.CreateMetadataRequest) (
	models.ApiResponse[[]models.Metadata],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/%v/%v/metadata.json", resourceType, resourceId),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSingleErrorResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result []models.Metadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Metadata](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListMetadataInput represents the input of the ListMetadata endpoint.
type ListMetadataInput struct {
	// the resource type to which the metafields belong
	ResourceType models.ResourceType
	// The Chargify id of the customer or the subscription for which the metadata applies
	ResourceId int
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
}

// ListMetadata takes context, resourceType, resourceId, page, perPage as parameters and
// returns an models.ApiResponse with models.PaginatedMetadata data and
// an error if there was an issue with the request or response.
// This request will list all of the metadata belonging to a particular resource (ie. subscription, customer) that is specified.
// ## Metadata Data
// This endpoint will also display the current stats of your metadata to use as a tool for pagination.
func (c *CustomFieldsController) ListMetadata(
	ctx context.Context,
	input ListMetadataInput) (
	models.ApiResponse[models.PaginatedMetadata],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/%v/%v/metadata.json", input.ResourceType, input.ResourceId),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	if input.Page != nil {
		req.QueryParamWithArraySerializationOption("page", *input.Page, https.Plain)
	}
	if input.PerPage != nil {
		req.QueryParamWithArraySerializationOption("per_page", *input.PerPage, https.Plain)
	}

	var result models.PaginatedMetadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.PaginatedMetadata](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateMetadata takes context, resourceType, resourceId, body as parameters and
// returns an models.ApiResponse with []models.Metadata data and
// an error if there was an issue with the request or response.
// This method allows you to update the existing metadata associated with a subscription or customer.
func (c *CustomFieldsController) UpdateMetadata(
	ctx context.Context,
	resourceType models.ResourceType,
	resourceId int,
	body *models.UpdateMetadataRequest) (
	models.ApiResponse[[]models.Metadata],
	error) {
	req := c.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/%v/%v/metadata.json", resourceType, resourceId),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSingleErrorResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result []models.Metadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Metadata](decoder)
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
	resourceId int,
	name *string,
	names []string) (
	*http.Response,
	error) {
	req := c.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/%v/%v/metadata.json", resourceType, resourceId),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})
	if name != nil {
		req.QueryParamWithArraySerializationOption("name", *name, https.Plain)
	}
	if names != nil {
		req.QueryParamWithArraySerializationOption("names[]", names, https.Plain)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// ListMetadataForResourceTypeInput represents the input of the ListMetadataForResourceType endpoint.
type ListMetadataForResourceTypeInput struct {
	// the resource type to which the metafields belong
	ResourceType models.ResourceType
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
	// The type of filter you would like to apply to your search.
	DateField *models.BasicDateField
	// The start date (format YYYY-MM-DD) with which to filter the date_field. Returns metadata with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
	StartDate *time.Time
	// The end date (format YYYY-MM-DD) with which to filter the date_field. Returns metadata with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
	EndDate *time.Time
	// The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns metadata with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
	StartDatetime *time.Time
	// The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns metadata with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.
	EndDatetime *time.Time
	// Allow to fetch deleted metadata.
	WithDeleted *bool
	// Allow to fetch metadata for multiple records based on provided ids. Use in query: `resource_ids[]=122&resource_ids[]=123&resource_ids[]=124`.
	ResourceIds []int
	// Controls the order in which results are returned.
	// Use in query `direction=asc`.
	Direction *models.SortingDirection
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
	input ListMetadataForResourceTypeInput) (
	models.ApiResponse[models.PaginatedMetadata],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/%v/metadata.json", input.ResourceType),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.DateField != nil {
		req.QueryParam("date_field", *input.DateField)
	}
	if input.StartDate != nil {
		req.QueryParam("start_date", input.StartDate.Format(models.DEFAULT_DATE))
	}
	if input.EndDate != nil {
		req.QueryParam("end_date", input.EndDate.Format(models.DEFAULT_DATE))
	}
	if input.StartDatetime != nil {
		req.QueryParam("start_datetime", input.StartDatetime.Format(time.RFC3339))
	}
	if input.EndDatetime != nil {
		req.QueryParam("end_datetime", input.EndDatetime.Format(time.RFC3339))
	}
	if input.WithDeleted != nil {
		req.QueryParam("with_deleted", *input.WithDeleted)
	}
	if input.ResourceIds != nil {
		req.QueryParam("resource_ids[]", input.ResourceIds)
	}
	if input.Direction != nil {
		req.QueryParam("direction", *input.Direction)
	}

	var result models.PaginatedMetadata
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.PaginatedMetadata](decoder)
	return models.NewApiResponse(result, resp), err
}
