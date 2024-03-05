package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "time"
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
// For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).
func (c *ComponentsController) CreateMeteredComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateMeteredComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_families/%v/metered_components.json", productFamilyId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
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
// For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).
func (c *ComponentsController) CreateQuantityBasedComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateQuantityBasedComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_families/%v/quantity_based_components.json", productFamilyId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
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
// For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).
func (c *ComponentsController) CreateOnOffComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateOnOffComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_families/%v/on_off_components.json", productFamilyId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
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
// For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).
func (c *ComponentsController) CreatePrepaidUsageComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreatePrepaidComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_families/%v/prepaid_usage_components.json", productFamilyId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
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
// For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).
func (c *ComponentsController) CreateEventBasedComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateEBBComponent) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_families/%v/event_based_components.json", productFamilyId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
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
// This request will return information regarding a component from a specific product family.
// You may read the component by either the component's id or handle. When using the handle, it must be prefixed with `handle:`.
func (c *ComponentsController) ReadComponent(
    ctx context.Context,
    productFamilyId int,
    componentId string) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/product_families/%v/components/%v.json", productFamilyId, componentId),
    )
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
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/product_families/%v/components/%v.json", productFamilyId, componentId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
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
    req := c.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/product_families/%v/components/%v.json", productFamilyId, componentId),
    )
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

// ListComponents takes context, dateField, startDate, endDate, startDatetime, endDatetime, includeArchived, page, perPage, filterIds, filterUseSiteExchangeRate as parameters and
// returns an models.ApiResponse with []models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will return a list of components for a site.
func (c *ComponentsController) ListComponents(
    ctx context.Context,
    dateField *models.BasicDateField,
    startDate *string,
    endDate *string,
    startDatetime *string,
    endDatetime *string,
    includeArchived *bool,
    page *int,
    perPage *int,
    filterIds []string,
    filterUseSiteExchangeRate *bool) (
    models.ApiResponse[[]models.ComponentResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/components.json")
    req.Authenticate(NewAuth("BasicAuth"))
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
    if includeArchived != nil {
        req.QueryParam("include_archived", *includeArchived)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if filterIds != nil {
        req.QueryParam("filter[ids]", filterIds)
    }
    if filterUseSiteExchangeRate != nil {
        req.QueryParam("filter[use_site_exchange_rate]", *filterUseSiteExchangeRate)
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
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/components/%v.json", componentId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// PromoteComponentPricePointToDefault takes context, componentId, pricePointId as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// Sets a new default price point for the component. This new default will apply to all new subscriptions going forward - existing subscriptions will remain on their current price point.
// See [Price Points Documentation](https://chargify.zendesk.com/hc/en-us/articles/4407755865883#price-points) for more information on price points and moving subscriptions between price points.
// Note: Custom price points are not able to be set as the default for a component.
func (c *ComponentsController) PromoteComponentPricePointToDefault(
    ctx context.Context,
    componentId int,
    pricePointId int) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/components/%v/price_points/%v/default.json", componentId, pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListComponentsForProductFamily takes context, productFamilyId, includeArchived, filterIds, page, perPage, dateField, endDate, endDatetime, startDate, startDatetime, filterUseSiteExchangeRate as parameters and
// returns an models.ApiResponse with []models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will return a list of components for a particular product family.
func (c *ComponentsController) ListComponentsForProductFamily(
    ctx context.Context,
    productFamilyId int,
    includeArchived *bool,
    filterIds []int,
    page *int,
    perPage *int,
    dateField *models.BasicDateField,
    endDate *string,
    endDatetime *string,
    startDate *string,
    startDatetime *string,
    filterUseSiteExchangeRate *bool) (
    models.ApiResponse[[]models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/product_families/%v/components.json", productFamilyId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if includeArchived != nil {
        req.QueryParam("include_archived", *includeArchived)
    }
    if filterIds != nil {
        req.QueryParam("filter[ids]", filterIds)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if dateField != nil {
        req.QueryParam("date_field", *dateField)
    }
    if endDate != nil {
        req.QueryParam("end_date", *endDate)
    }
    if endDatetime != nil {
        req.QueryParam("end_datetime", *endDatetime)
    }
    if startDate != nil {
        req.QueryParam("start_date", *startDate)
    }
    if startDatetime != nil {
        req.QueryParam("start_datetime", *startDatetime)
    }
    if filterUseSiteExchangeRate != nil {
        req.QueryParam("filter[use_site_exchange_rate]", *filterUseSiteExchangeRate)
    }
    
    var result []models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ComponentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateComponentPricePoint takes context, componentId, body as parameters and
// returns an models.ApiResponse with models.ComponentPricePointResponse data and
// an error if there was an issue with the request or response.
// This endpoint can be used to create a new price point for an existing component.
func (c *ComponentsController) CreateComponentPricePoint(
    ctx context.Context,
    componentId int,
    body *models.CreateComponentPricePointRequest) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/components/%v/price_points.json", componentId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListComponentPricePoints takes context, componentId, currencyPrices, page, perPage, filterType as parameters and
// returns an models.ApiResponse with models.ComponentPricePointsResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to read current price points that are associated with a component.
// You may specify the component by using either the numeric id or the `handle:gold` syntax.
// When fetching a component's price points, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response.
// If the price point is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.
func (c *ComponentsController) ListComponentPricePoints(
    ctx context.Context,
    componentId int,
    currencyPrices *bool,
    page *int,
    perPage *int,
    filterType []models.PricePointType) (
    models.ApiResponse[models.ComponentPricePointsResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/components/%v/price_points.json", componentId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if currencyPrices != nil {
        req.QueryParam("currency_prices", *currencyPrices)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if filterType != nil {
        req.QueryParam("filter[type]", filterType)
    }
    
    var result models.ComponentPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// BulkCreateComponentPricePoints takes context, componentId, body as parameters and
// returns an models.ApiResponse with models.ComponentPricePointsResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to create multiple component price points in one request.
func (c *ComponentsController) BulkCreateComponentPricePoints(
    ctx context.Context,
    componentId string,
    body *models.CreateComponentPricePointsRequest) (
    models.ApiResponse[models.ComponentPricePointsResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/components/%v/price_points/bulk.json", componentId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateComponentPricePoint takes context, componentId, pricePointId, body as parameters and
// returns an models.ApiResponse with models.ComponentPricePointResponse data and
// an error if there was an issue with the request or response.
// When updating a price point, it's prices can be updated as well by creating new prices or editing / removing existing ones.
// Passing in a price bracket without an `id` will attempt to create a new price.
// Including an `id` will update the corresponding price, and including the `_destroy` flag set to true along with the `id` will remove that price.
// Note: Custom price points cannot be updated directly. They must be edited through the Subscription.
func (c *ComponentsController) UpdateComponentPricePoint(
    ctx context.Context,
    componentId int,
    pricePointId int,
    body *models.UpdateComponentPricePointRequest) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/components/%v/price_points/%v.json", componentId, pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorArrayMapResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ArchiveComponentPricePoint takes context, componentId, pricePointId as parameters and
// returns an models.ApiResponse with models.ComponentPricePointResponse data and
// an error if there was an issue with the request or response.
// A price point can be archived at any time. Subscriptions using a price point that has been archived will continue using it until they're moved to another price point.
func (c *ComponentsController) ArchiveComponentPricePoint(
    ctx context.Context,
    componentId int,
    pricePointId int) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/components/%v/price_points/%v.json", componentId, pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UnarchiveComponentPricePoint takes context, componentId, pricePointId as parameters and
// returns an models.ApiResponse with models.ComponentPricePointResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to unarchive a component price point.
func (c *ComponentsController) UnarchiveComponentPricePoint(
    ctx context.Context,
    componentId int,
    pricePointId int) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/components/%v/price_points/%v/unarchive.json", componentId, pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateCurrencyPrices takes context, pricePointId, body as parameters and
// returns an models.ApiResponse with models.ComponentCurrencyPricesResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to create currency prices for a given currency that has been defined on the site level in your settings.
// When creating currency prices, they need to mirror the structure of your primary pricing. For each price level defined on the component price point, there should be a matching price level created in the given currency.
// Note: Currency Prices are not able to be created for custom price points.
func (c *ComponentsController) CreateCurrencyPrices(
    ctx context.Context,
    pricePointId int,
    body *models.CreateCurrencyPricesRequest) (
    models.ApiResponse[models.ComponentCurrencyPricesResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/price_points/%v/currency_prices.json", pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorArrayMapResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentCurrencyPricesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentCurrencyPricesResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateCurrencyPrices takes context, pricePointId, body as parameters and
// returns an models.ApiResponse with models.ComponentCurrencyPricesResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to update currency prices for a given currency that has been defined on the site level in your settings.
// Note: Currency Prices are not able to be updated for custom price points.
func (c *ComponentsController) UpdateCurrencyPrices(
    ctx context.Context,
    pricePointId int,
    body *models.UpdateCurrencyPricesRequest) (
    models.ApiResponse[models.ComponentCurrencyPricesResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/price_points/%v/currency_prices.json", pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorArrayMapResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentCurrencyPricesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentCurrencyPricesResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListAllComponentPricePoints takes context, filterDateField, filterEndDate, filterEndDatetime, include, page, perPage, filterStartDate, filterStartDatetime, filterType, direction, filterIds, filterArchivedAt as parameters and
// returns an models.ApiResponse with models.ListComponentsPricePointsResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve a list of Components Price Points belonging to a Site.
func (c *ComponentsController) ListAllComponentPricePoints(
    ctx context.Context,
    filterDateField *models.BasicDateField,
    filterEndDate *time.Time,
    filterEndDatetime *time.Time,
    include *models.ListComponentsPricePointsInclude,
    page *int,
    perPage *int,
    filterStartDate *time.Time,
    filterStartDatetime *time.Time,
    filterType []models.PricePointType,
    direction *models.SortingDirection,
    filterIds []int,
    filterArchivedAt *models.IncludeNotNull) (
    models.ApiResponse[models.ListComponentsPricePointsResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/components_price_points.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    if filterDateField != nil {
        req.QueryParam("filter[date_field]", *filterDateField)
    }
    if filterEndDate != nil {
        req.QueryParam("filter[end_date]", filterEndDate.Format(models.DEFAULT_DATE))
    }
    if filterEndDatetime != nil {
        req.QueryParam("filter[end_datetime]", filterEndDatetime.Format(time.RFC3339))
    }
    if include != nil {
        req.QueryParam("include", *include)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if filterStartDate != nil {
        req.QueryParam("filter[start_date]", filterStartDate.Format(models.DEFAULT_DATE))
    }
    if filterStartDatetime != nil {
        req.QueryParam("filter[start_datetime]", filterStartDatetime.Format(time.RFC3339))
    }
    if filterType != nil {
        req.QueryParam("filter[type]", filterType)
    }
    if direction != nil {
        req.QueryParam("direction", *direction)
    }
    if filterIds != nil {
        req.QueryParam("filter[ids]", filterIds)
    }
    if filterArchivedAt != nil {
        req.QueryParam("filter[archived_at]", *filterArchivedAt)
    }
    var result models.ListComponentsPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListComponentsPricePointsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
