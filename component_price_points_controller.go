/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// ComponentPricePointsController represents a controller struct.
type ComponentPricePointsController struct {
    baseController
}

// NewComponentPricePointsController creates a new instance of ComponentPricePointsController.
// It takes a baseController as a parameter and returns a pointer to the ComponentPricePointsController.
func NewComponentPricePointsController(baseController baseController) *ComponentPricePointsController {
    componentPricePointsController := ComponentPricePointsController{baseController: baseController}
    return &componentPricePointsController
}

// PromoteComponentPricePointToDefault takes context, componentId, pricePointId as parameters and
// returns an models.ApiResponse with models.ComponentResponse data and
// an error if there was an issue with the request or response.
// Sets a new default price point for the component. This new default will apply to all new subscriptions going forward - existing subscriptions will remain on their current price point.
// See [Price Points Documentation](https://maxio.zendesk.com/hc/en-us/articles/24261191737101-Price-Points-Components) for more information on price points and moving subscriptions between price points.
// Note: Custom price points are not able to be set as the default for a component.
func (c *ComponentPricePointsController) PromoteComponentPricePointToDefault(
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

// CreateComponentPricePoint takes context, componentId, body as parameters and
// returns an models.ApiResponse with models.ComponentPricePointResponse data and
// an error if there was an issue with the request or response.
// This endpoint can be used to create a new price point for an existing component.
func (c *ComponentPricePointsController) CreateComponentPricePoint(
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
        req.Json(body)
    }
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListComponentPricePointsInput represents the input of the ListComponentPricePoints endpoint.
type ListComponentPricePointsInput struct {
    // The Advanced Billing id of the component
    ComponentId    int                     
    // Include an array of currency price data
    CurrencyPrices *bool                   
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page           *int                    
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage        *int                    
    // Use in query: `filter[type]=catalog,default`.
    FilterType     []models.PricePointType 
}

// ListComponentPricePoints takes context, componentId, currencyPrices, page, perPage, filterType as parameters and
// returns an models.ApiResponse with models.ComponentPricePointsResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to read current price points that are associated with a component.
// You may specify the component by using either the numeric id or the `handle:gold` syntax.
// When fetching a component's price points, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response.
// If the price point is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.
func (c *ComponentPricePointsController) ListComponentPricePoints(
    ctx context.Context,
    input ListComponentPricePointsInput) (
    models.ApiResponse[models.ComponentPricePointsResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/components/%v/price_points.json", input.ComponentId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if input.CurrencyPrices != nil {
        req.QueryParam("currency_prices", *input.CurrencyPrices)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.FilterType != nil {
        req.QueryParam("filter[type]", input.FilterType)
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
func (c *ComponentPricePointsController) BulkCreateComponentPricePoints(
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
        req.Json(body)
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
func (c *ComponentPricePointsController) UpdateComponentPricePoint(
    ctx context.Context,
    componentId models.UpdateComponentPricePointComponentId,
    pricePointId models.UpdateComponentPricePointPricePointId,
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
        req.Json(body)
    }
    
    var result models.ComponentPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadComponentPricePoint takes context, componentId, pricePointId as parameters and
// returns an models.ApiResponse with models.ComponentPricePointResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to retrieve details for a specific component price point. You can achieve this by using either the component price point ID or handle.
func (c *ComponentPricePointsController) ReadComponentPricePoint(
    ctx context.Context,
    componentId models.ReadComponentPricePointComponentId,
    pricePointId models.ReadComponentPricePointPricePointId) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/components/%v/price_points/%v.json", componentId, pricePointId),
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

// ArchiveComponentPricePoint takes context, componentId, pricePointId as parameters and
// returns an models.ApiResponse with models.ComponentPricePointResponse data and
// an error if there was an issue with the request or response.
// A price point can be archived at any time. Subscriptions using a price point that has been archived will continue using it until they're moved to another price point.
func (c *ComponentPricePointsController) ArchiveComponentPricePoint(
    ctx context.Context,
    componentId models.ArchiveComponentPricePointComponentId,
    pricePointId models.ArchiveComponentPricePointPricePointId) (
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
func (c *ComponentPricePointsController) UnarchiveComponentPricePoint(
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
func (c *ComponentPricePointsController) CreateCurrencyPrices(
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
        req.Json(body)
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
func (c *ComponentPricePointsController) UpdateCurrencyPrices(
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
        req.Json(body)
    }
    
    var result models.ComponentCurrencyPricesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ComponentCurrencyPricesResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListAllComponentPricePointsInput represents the input of the ListAllComponentPricePoints endpoint.
type ListAllComponentPricePointsInput struct {
    // Allows including additional data in the response. Use in query: `include=currency_prices`.
    Include   *models.ListComponentsPricePointsInclude 
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page      *int                                     
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage   *int                                     
    // Controls the order in which results are returned.
    // Use in query `direction=asc`.
    Direction *models.SortingDirection                 
    // Filter to use for List PricePoints operations
    Filter    *models.ListPricePointsFilter            
}

// ListAllComponentPricePoints takes context, include, page, perPage, direction, filter as parameters and
// returns an models.ApiResponse with models.ListComponentsPricePointsResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve a list of Components Price Points belonging to a Site.
func (c *ComponentPricePointsController) ListAllComponentPricePoints(
    ctx context.Context,
    input ListAllComponentPricePointsInput) (
    models.ApiResponse[models.ListComponentsPricePointsResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/components_price_points.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    if input.Include != nil {
        req.QueryParam("include", *input.Include)
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
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    var result models.ListComponentsPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListComponentsPricePointsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
