// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// ComponentsController represents a controller struct.
type ComponentsController struct {
    baseController
}

// NewComponentsController creates a new instance of ComponentsController.
// It takes a baseController as a parameter and returns a pointer to the ComponentsController.
func NewComponentsController(baseController baseController) *ComponentsController {
    componentsController := ComponentsController{baseController: baseController}
    return &componentsController
}

// CreateMeteredComponent takes context, productFamilyId, body as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will create a component definition of kind **metered_component** under the specified product family. Metered component can then be added and “allocated” for a subscription.
// Metered components are used to bill for any type of unit that resets to 0 at the end of the billing period (think daily Google Adwords clicks or monthly cell phone minutes). This is most commonly associated with usage-based billing and many other pricing schemes.
// Note that this is different from recurring quantity-based components, which DO NOT reset to zero at the start of every billing period. If you want to bill for a quantity of something that does not change unless you change it, then you want quantity components, instead.
// For more information on components, see our documentation [here](https://maxio.zendesk.com/hc/en-us/articles/24261141522189-Components-Overview).
func (c *ComponentsController) CreateMeteredComponent(
    ctx context.Context,
    productFamilyId string,
    body *models.CreateMeteredComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      "/product_families/%v/metered_components.json",
    )
    req.AppendTemplateParams(productFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateQuantityBasedComponent takes context, productFamilyId, body as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will create a component definition of kind **quantity_based_component** under the specified product family. Quantity Based component can then be added and “allocated” for a subscription.
// When defining Quantity Based component, You can choose one of 2 types:
// #### Recurring
// Recurring quantity-based components are used to bill for the number of some unit (think monthly software user licenses or the number of pairs of socks in a box-a-month club). This is most commonly associated with billing for user licenses, number of users, number of employees, etc.
// #### One-time
// One-time quantity-based components are used to create ad hoc usage charges that do not recur. For example, at the time of signup, you might want to charge your customer a one-time fee for onboarding or other services.
// The allocated quantity for one-time quantity-based components immediately gets reset back to zero after the allocation is made.
// For more information on components, see our documentation [here](https://maxio.zendesk.com/hc/en-us/articles/24261141522189-Components-Overview).
func (c *ComponentsController) CreateQuantityBasedComponent(
    ctx context.Context,
    productFamilyId string,
    body *models.CreateQuantityBasedComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      "/product_families/%v/quantity_based_components.json",
    )
    req.AppendTemplateParams(productFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOnOffComponent takes context, productFamilyId, body as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will create a component definition of kind **on_off_component** under the specified product family. On/Off component can then be added and “allocated” for a subscription.
// On/off components are used for any flat fee, recurring add on (think $99/month for tech support or a flat add on shipping fee).
// For more information on components, see our documentation [here](https://maxio.zendesk.com/hc/en-us/articles/24261141522189-Components-Overview).
func (c *ComponentsController) CreateOnOffComponent(
    ctx context.Context,
    productFamilyId string,
    body *models.CreateOnOffComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      "/product_families/%v/on_off_components.json",
    )
    req.AppendTemplateParams(productFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreatePrepaidUsageComponent takes context, productFamilyId, body as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will create a component definition of kind **prepaid_usage_component** under the specified product family. Prepaid component can then be added and “allocated” for a subscription.
// Prepaid components allow customers to pre-purchase units that can be used up over time on their subscription. In a sense, they are the mirror image of metered components; while metered components charge at the end of the period for the amount of units used, prepaid components are charged for at the time of purchase, and we subsequently keep track of the usage against the amount purchased.
// For more information on components, see our documentation [here](https://maxio.zendesk.com/hc/en-us/articles/24261141522189-Components-Overview).
func (c *ComponentsController) CreatePrepaidUsageComponent(
    ctx context.Context,
    productFamilyId string,
    body *models.CreatePrepaidComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      "/product_families/%v/prepaid_usage_components.json",
    )
    req.AppendTemplateParams(productFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateEventBasedComponent takes context, productFamilyId, body as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will create a component definition of kind **event_based_component** under the specified product family. Event-based component can then be added and “allocated” for a subscription.
// Event-based components are similar to other component types, in that you define the component parameters (such as name and taxability) and the pricing. A key difference for the event-based component is that it must be attached to a metric. This is because the metric provides the component with the actual quantity used in computing what and how much will be billed each period for each subscription.
// So, instead of reporting usage directly for each component (as you would with metered components), the usage is derived from analysis of your events.
// For more information on components, see our documentation [here](https://maxio.zendesk.com/hc/en-us/articles/24261141522189-Components-Overview).
func (c *ComponentsController) CreateEventBasedComponent(
    ctx context.Context,
    productFamilyId string,
    body *models.CreateEBBComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      "/product_families/%v/event_based_components.json",
    )
    req.AppendTemplateParams(productFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// FindComponent takes context, handle as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will return information regarding a component having the handle you provide. You can identify your components with a handle so you don't have to save or reference the IDs we generate.
func (c *ComponentsController) FindComponent(
    ctx context.Context,
    handle string) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/components/lookup.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.QueryParam("handle", handle)
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadComponent takes context, productFamilyId, componentId as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// Returns information regarding a component from a specific product family.
// You can read the component by either the component's id or handle. When using the handle, it must be prefixed with `handle:`.
func (c *ComponentsController) ReadComponent(
    ctx context.Context,
    productFamilyId int,
    componentId string) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/product_families/%v/components/%v.json")
    req.AppendTemplateParams(productFamilyId, componentId)
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateProductFamilyComponent takes context, productFamilyId, componentId, body as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will update a component from a specific product family.
// You may read the component by either the component's id or handle. When using the handle, it must be prefixed with `handle:`.
func (c *ComponentsController) UpdateProductFamilyComponent(
    ctx context.Context,
    productFamilyId int,
    componentId string,
    body *models.UpdateComponentRequest) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(ctx, "PUT", "/product_families/%v/components/%v.json")
    req.AppendTemplateParams(productFamilyId, componentId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ArchiveComponent takes context, productFamilyId, componentId as parameters and
// returns an models.ApiResponse with models.Component data and
// an error if there was an issue with the request or response.
// Sending a DELETE request to this endpoint will archive the component. All current subscribers will be unffected; their subscription/purchase will continue to be charged as usual.
func (c *ComponentsController) ArchiveComponent(
    ctx context.Context,
    productFamilyId int,
    componentId string) (
    models.ApiResponse[models.Component],
    error) {
    req := c.prepareRequest(ctx, "DELETE", "/product_families/%v/components/%v.json")
    req.AppendTemplateParams(productFamilyId, componentId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.Component
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Component](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListComponentsInput represents the input of the ListComponents endpoint.
type ListComponentsInput struct {
    // The type of filter you would like to apply to your search.
    DateField       *models.BasicDateField       
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
    StartDate       *string                      
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
    EndDate         *string                      
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
    StartDatetime   *string                      
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.  optional
    EndDatetime     *string                      
    // Include archived items
    IncludeArchived *bool                        
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page            *int                         
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage         *int                         
    // Filter to use for List Components operations
    Filter          *models.ListComponentsFilter 
}

// ListComponents takes context, dateField, startDate, endDate, startDatetime, endDatetime, includeArchived, page, perPage, filter as parameters and
// returns an models.ApiResponse with []models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will return a list of components for a site.
func (c *ComponentsController) ListComponents(
    ctx context.Context,
    input ListComponentsInput) (
    models.ApiResponse[[]models.ComponentResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/components.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    if input.DateField != nil {
        req.QueryParam("date_field", *input.DateField)
    }
    if input.StartDate != nil {
        req.QueryParam("start_date", *input.StartDate)
    }
    if input.EndDate != nil {
        req.QueryParam("end_date", *input.EndDate)
    }
    if input.StartDatetime != nil {
        req.QueryParam("start_datetime", *input.StartDatetime)
    }
    if input.EndDatetime != nil {
        req.QueryParam("end_datetime", *input.EndDatetime)
    }
    if input.IncludeArchived != nil {
        req.QueryParam("include_archived", *input.IncludeArchived)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    var result []models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateComponent takes context, componentId, body as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will update a component.
// You may read the component by either the component's id or handle. When using the handle, it must be prefixed with `handle:`.
func (c *ComponentsController) UpdateComponent(
    ctx context.Context,
    componentId string,
    body *models.UpdateComponentRequest) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(ctx, "PUT", "/components/%v.json")
    req.AppendTemplateParams(componentId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListComponentsForProductFamilyInput represents the input of the ListComponentsForProductFamily endpoint.
type ListComponentsForProductFamilyInput struct {
    // The Advanced Billing id of the product family
    ProductFamilyId int                          
    // Include archived items.
    IncludeArchived *bool                        
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page            *int                         
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage         *int                         
    // Filter to use for List Components operations
    Filter          *models.ListComponentsFilter 
    // The type of filter you would like to apply to your search. Use in query `date_field=created_at`.
    DateField       *models.BasicDateField       
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
    EndDate         *string                      
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. optional.
    EndDatetime     *string                      
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
    StartDate       *string                      
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
    StartDatetime   *string                      
}

// ListComponentsForProductFamily takes context, productFamilyId, includeArchived, page, perPage, filter, dateField, endDate, endDatetime, startDate, startDatetime as parameters and
// returns an models.ApiResponse with []models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will return a list of components for a particular product family.
func (c *ComponentsController) ListComponentsForProductFamily(
    ctx context.Context,
    input ListComponentsForProductFamilyInput) (
    models.ApiResponse[[]models.ComponentResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/product_families/%v/components.json")
    req.AppendTemplateParams(input.ProductFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    if input.IncludeArchived != nil {
        req.QueryParam("include_archived", *input.IncludeArchived)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    if input.DateField != nil {
        req.QueryParam("date_field", *input.DateField)
    }
    if input.EndDate != nil {
        req.QueryParam("end_date", *input.EndDate)
    }
    if input.EndDatetime != nil {
        req.QueryParam("end_datetime", *input.EndDatetime)
    }
    if input.StartDate != nil {
        req.QueryParam("start_date", *input.StartDate)
    }
    if input.StartDatetime != nil {
        req.QueryParam("start_datetime", *input.StartDatetime)
    }
    
    var result []models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
