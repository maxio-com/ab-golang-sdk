package advancedbilling

import (
    "context"
    "fmt"
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

// CreateComponent takes context, productFamilyId, componentKind, body as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will create a component definition under the specified product family. These component definitions determine what components are named, how they are measured, and how much they cost.
// Components can then be added and “allocated” for each subscription to a product in the product family. These component line-items affect how much a subscription will be charged, depending on the current allocations (i.e. 4 IP Addresses, or SSL “enabled”)
// This documentation covers both component definitions and component line-items. Please understand the difference.
// Please note that you may not edit components via API. To do so, please log into the application.
// ### Component Documentation
// For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).
// For information on how to record component usage against a subscription, please see the following resources:
// + [Proration and Component Allocations](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677#applying-proration-and-recording-components)
// + [Recording component usage against a subscription](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404606587917#recording-component-usage)
func (c *ComponentsController) CreateComponent(
    ctx context.Context,
    productFamilyId int,
    componentKind models.ComponentKindPath,
    body *interface{}) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_families/%v/%v.json", productFamilyId, componentKind),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ReadComponentByHandle takes context, handle as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will return information regarding a component having the handle you provide. You can identify your components with a handle so you don't have to save or reference the IDs we generate.
func (c *ComponentsController) ReadComponentByHandle(
    ctx context.Context,
    handle string) (
    models.ApiResponse[models.ComponentResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/components/lookup.json")
    req.Authenticate(true)
    req.QueryParam("handle", handle)
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ReadComponentById takes context, productFamilyId, componentId as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// This request will return information regarding a component from a specific product family.
// You may read the component by either the component's id or handle. When using the handle, it must be prefixed with `handle:`.
func (c *ComponentsController) ReadComponentById(
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
    req.Authenticate(true)
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
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
    req.Authenticate(true)
    
    var result models.Component
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Component](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
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
    req.Authenticate(true)
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
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ComponentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// UpdateDefaultPricePointForComponent takes context, componentId, pricePointId as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// Sets a new default price point for the component. This new default will apply to all new subscriptions going forward - existing subscriptions will remain on their current price point.
// See [Price Points Documentation](https://chargify.zendesk.com/hc/en-us/articles/4407755865883#price-points) for more information on price points and moving subscriptions between price points.
// Note: Custom price points are not able to be set as the default for a component.
func (c *ComponentsController) UpdateDefaultPricePointForComponent(
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
    req.Authenticate(true)
    
    var result models.ComponentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
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
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ComponentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
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
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// CreateComponentPricePoints takes context, componentId, body as parameters and
// returns an models.ApiResponse with models.ComponentPricePointsResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to create multiple component price points in one request.
func (c *ComponentsController) CreateComponentPricePoints(
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
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// CreateCurrencyPrices takes context, pricePointId, body as parameters and
// returns an models.ApiResponse with []models.CurrencyPrice data and
// an error if there was an issue with the request or response.
// This endpoint allows you to create currency prices for a given currency that has been defined on the site level in your settings.
// When creating currency prices, they need to mirror the structure of your primary pricing. For each price level defined on the component price point, there should be a matching price level created in the given currency.
// Note: Currency Prices are not able to be created for custom price points.
func (c *ComponentsController) CreateCurrencyPrices(
    ctx context.Context,
    pricePointId int,
    body *models.CreateCurrencyPricesRequest) (
    models.ApiResponse[[]models.CurrencyPrice],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/price_points/%v/currency_prices.json", pricePointId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result []models.CurrencyPrice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.CurrencyPrice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// UpdateCurrencyPrices takes context, pricePointId, body as parameters and
// returns an models.ApiResponse with []models.CurrencyPrice data and
// an error if there was an issue with the request or response.
// This endpoint allows you to update currency prices for a given currency that has been defined on the site level in your settings.
// Note: Currency Prices are not able to be updated for custom price points.
func (c *ComponentsController) UpdateCurrencyPrices(
    ctx context.Context,
    pricePointId int,
    body *models.UpdateCurrencyPricesRequest) (
    models.ApiResponse[[]models.CurrencyPrice],
    error) {
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/price_points/%v/currency_prices.json", pricePointId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result []models.CurrencyPrice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.CurrencyPrice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
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
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListComponentsPricePointsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}
