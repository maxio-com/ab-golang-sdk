// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "context"
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
// Creates metafields on a Site for either the Subscriptions or Customers resource. 
// Metafields and their metadata are created in the Custom Fields configuration page on your Site. Metafields can be populated with metadata when you create them or later with the [Update Metafield]($e/Custom%20Fields/updateMetafield), [Create Metadata]($e/Custom%20Fields/createMetadata), or [Update Metadata]($e/Custom%20Fields/updateMetadata) endpoints. The Create Metadata and Update Metadata endpoints allow you to add metafields and metadata values to a specific subscription or customer.
// Each site is limited to 100 unique metafields per resource. This means you can have 100 metafields for Subscriptions and another 100 for Customers.
// > Note: After creating a metafield, the resource type cannot be modified.
// In the UI and product documentation, metafields and metadata are called Custom Fields. 
// - Metafield is the custom field
// - Metadata is the data populating the custom field.
// See [Custom Fields Reference](https://docs.maxio.com/hc/en-us/articles/24266140850573-Custom-Fields-Reference) and [Custom Fields Tab](https://maxio.zendesk.com/hc/en-us/articles/24251701302925-Subscription-Summary-Custom-Fields-Tab) for information on using Custom Fields in the Advanced Billing UI.
func (c *CustomFieldsController) CreateMetafields(
    ctx context.Context,
    resourceType models.ResourceType,
    body *models.CreateMetafieldsRequest) (
    models.ApiResponse[[]models.Metafield],
    error) {
    req := c.prepareRequest(ctx, "POST", "/%v/metafields.json")
    req.AppendTemplateParams(resourceType)
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
    // The resource type to which the metafields belong.
    ResourceType models.ResourceType      
    // Filter by the name of the metafield.
    Name         *string                  
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page         *int                     
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage      *int                     
    // Controls the order in which results are returned.
    // Use in query `direction=asc`.
    Direction    *models.SortingDirection 
}

// ListMetafields takes context, resourceType, name, page, perPage, direction as parameters and
// returns an models.ApiResponse with models.ListMetafieldsResponse data and
// an error if there was an issue with the request or response.
// Lists the metafields and their associated details for a Site and resource type. You can filter the request to a specific metafield.
func (c *CustomFieldsController) ListMetafields(
    ctx context.Context,
    input ListMetafieldsInput) (
    models.ApiResponse[models.ListMetafieldsResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/%v/metafields.json")
    req.AppendTemplateParams(input.ResourceType)
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
// Updates metafields on your Site for a resource type.  Depending on the request structure, you can update or add metafields and metadata to the Subscriptions or Customers resource.
// With this endpoint, you can: 
// - Add metafields. If the metafield specified in current_name does not exist, a new metafield is added. 
// >Note: Each site is limited to 100 unique metafields per resource. This means you can have 100 metafields for Subscriptions and another 100 for Customers.
// - Change the name of a metafield. 
// >Note: To keep the metafield name the same and only update the metadata for the metafield, you must use the current metafield name in both the `current_name` and `name` parameters.
// - Change the input type for the metafield. For example, you can change a metafield input type from text to a dropdown. If you change the input type from text to a dropdown or radio, you must update the specific subscriptions or customers where the metafield was used to reflect the updated metafield and metadata. 
// - Add metadata values to the existing metadata for a dropdown or radio metafield. 
// >Note: Updates to metadata overwrite. To add one or more values, you must specify all metadata values including the new value you want to add.
// - Add new metadata to a dropdown or radio for a metafield that was created without metadata.
// - Remove  metadata for a dropdown or radio for a metafield.  
// >Note: Updates to metadata overwrite existing values. To remove one or more values, specify all metadata values except those you want to remove.
// - Add or update scope settings for a metafield.
// >Note: Scope changes overwrite existing settings. You must specify the complete scope, including the changes you want to make.
func (c *CustomFieldsController) UpdateMetafield(
    ctx context.Context,
    resourceType models.ResourceType,
    body *models.UpdateMetafieldsRequest) (
    models.ApiResponse[[]models.Metafield],
    error) {
    req := c.prepareRequest(ctx, "PUT", "/%v/metafields.json")
    req.AppendTemplateParams(resourceType)
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
// returns an *Response and
// an error if there was an issue with the request or response.
// Deletes a metafield from your Site. Removes the metafield and associated metadata from all Subscriptions or Customers resources on the Site.
func (c *CustomFieldsController) DeleteMetafield(
    ctx context.Context,
    resourceType models.ResourceType,
    name *string) (
    *http.Response,
    error) {
    req := c.prepareRequest(ctx, "DELETE", "/%v/metafields.json")
    req.AppendTemplateParams(resourceType)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
    })
    if name != nil {
        req.QueryParam("name", *name)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// CreateMetadata takes context, resourceType, resourceId, body as parameters and
// returns an models.ApiResponse with []models.Metadata data and
// an error if there was an issue with the request or response.
// Creates metadata and metafields for a specific subscription or customer, or updates metadata values of existing metafields for a subscription or customer. Metadata values are limited to 2 KB in size.
// If you create metadata on a subscription or customer with a metafield that does not already exist, the metafield is created with the metadata you specify and it is always added as a text field. You can update the input_type for the metafield with the [Update Metafield]($e/Custom%20Fields/updateMetafield) endpoint. 
// >Note: Each site is limited to 100 unique metafields per resource. This means you can have 100 metafields for Subscriptions and another 100 for Customers.
func (c *CustomFieldsController) CreateMetadata(
    ctx context.Context,
    resourceType models.ResourceType,
    resourceId int,
    body *models.CreateMetadataRequest) (
    models.ApiResponse[[]models.Metadata],
    error) {
    req := c.prepareRequest(ctx, "POST", "/%v/%v/metadata.json")
    req.AppendTemplateParams(resourceType, resourceId)
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
    // The resource type to which the metafields belong.
    ResourceType models.ResourceType 
    // The Advanced Billing id of the customer or the subscription for which the metadata applies
    ResourceId   int                 
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page         *int                
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage      *int                
}

// ListMetadata takes context, resourceType, resourceId, page, perPage as parameters and
// returns an models.ApiResponse with models.PaginatedMetadata data and
// an error if there was an issue with the request or response.
// Lists metadata and metafields for a specific customer or subscription.
func (c *CustomFieldsController) ListMetadata(
    ctx context.Context,
    input ListMetadataInput) (
    models.ApiResponse[models.PaginatedMetadata],
    error) {
    req := c.prepareRequest(ctx, "GET", "/%v/%v/metadata.json")
    req.AppendTemplateParams(input.ResourceType, input.ResourceId)
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
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
// Updates metadata and metafields on the Site and the customer or subscription specified, and updates the metadata value on a subscription or customer.
// If you update metadata on a subscription or customer with a metafield that does not already exist, the metafield is created with the metadata you specify and it is always added as a text field to the Site and to the subscription or customer you specify. You can update the input_type for the metafield with the Update Metafield endpoint. 
// Each site is limited to 100 unique metafields per resource. This means you can have 100 metafields for Subscription and another 100 for Customer.
func (c *CustomFieldsController) UpdateMetadata(
    ctx context.Context,
    resourceType models.ResourceType,
    resourceId int,
    body *models.UpdateMetadataRequest) (
    models.ApiResponse[[]models.Metadata],
    error) {
    req := c.prepareRequest(ctx, "PUT", "/%v/%v/metadata.json")
    req.AppendTemplateParams(resourceType, resourceId)
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
// returns an *Response and
// an error if there was an issue with the request or response.
// Deletes one or more metafields (and associated metadata) from the specified subscription or customer.
func (c *CustomFieldsController) DeleteMetadata(
    ctx context.Context,
    resourceType models.ResourceType,
    resourceId int,
    name *string,
    names []string) (
    *http.Response,
    error) {
    req := c.prepareRequest(ctx, "DELETE", "/%v/%v/metadata.json")
    req.AppendTemplateParams(resourceType, resourceId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
    })
    if name != nil {
        req.QueryParamWithArraySerializationOption("name", *name, https.UnIndexed)
    }
    if names != nil {
        req.QueryParamWithArraySerializationOption("names", names, https.UnIndexed)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// ListMetadataForResourceTypeInput represents the input of the ListMetadataForResourceType endpoint.
type ListMetadataForResourceTypeInput struct {
    // The resource type to which the metafields belong.
    ResourceType  models.ResourceType      
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page          *int                     
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage       *int                     
    // The type of filter you would like to apply to your search.
    DateField     *models.BasicDateField   
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns metadata with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
    StartDate     *time.Time               
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns metadata with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
    EndDate       *time.Time               
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns metadata with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
    StartDatetime *time.Time               
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns metadata with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.
    EndDatetime   *time.Time               
    // Allow to fetch deleted metadata.
    WithDeleted   *bool                    
    // Allow to fetch metadata for multiple records based on provided ids. Use in query: `resource_ids[]=122&resource_ids[]=123&resource_ids[]=124`.
    ResourceIds   []int                    
    // Controls the order in which results are returned.
    // Use in query `direction=asc`.
    Direction     *models.SortingDirection 
}

// ListMetadataForResourceType takes context, resourceType, page, perPage, dateField, startDate, endDate, startDatetime, endDatetime, withDeleted, resourceIds, direction as parameters and
// returns an models.ApiResponse with models.PaginatedMetadata data and
// an error if there was an issue with the request or response.
// Lists  metadata for a specified array of subscriptions or customers.
func (c *CustomFieldsController) ListMetadataForResourceType(
    ctx context.Context,
    input ListMetadataForResourceTypeInput) (
    models.ApiResponse[models.PaginatedMetadata],
    error) {
    req := c.prepareRequest(ctx, "GET", "/%v/metadata.json")
    req.AppendTemplateParams(input.ResourceType)
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParamWithArraySerializationOption("page", *input.Page, https.UnIndexed)
    }
    if input.PerPage != nil {
        req.QueryParamWithArraySerializationOption("per_page", *input.PerPage, https.UnIndexed)
    }
    if input.DateField != nil {
        req.QueryParamWithArraySerializationOption("date_field", *input.DateField, https.UnIndexed)
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
        req.QueryParamWithArraySerializationOption("with_deleted", *input.WithDeleted, https.UnIndexed)
    }
    if input.ResourceIds != nil {
        req.QueryParamWithArraySerializationOption("resource_ids", input.ResourceIds, https.UnIndexed)
    }
    if input.Direction != nil {
        req.QueryParamWithArraySerializationOption("direction", *input.Direction, https.UnIndexed)
    }
    
    var result models.PaginatedMetadata
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PaginatedMetadata](decoder)
    return models.NewApiResponse(result, resp), err
}
