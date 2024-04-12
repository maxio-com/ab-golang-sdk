package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// ProductPricePointsController represents a controller struct.
type ProductPricePointsController struct {
    baseController
}

// NewProductPricePointsController creates a new instance of ProductPricePointsController.
// It takes a baseController as a parameter and returns a pointer to the ProductPricePointsController.
func NewProductPricePointsController(baseController baseController) *ProductPricePointsController {
    productPricePointsController := ProductPricePointsController{baseController: baseController}
    return &productPricePointsController
}

// CreateProductPricePoint takes context, productId, body as parameters and
// returns an models.ApiResponse with models.ProductPricePointResponse data and
// an error if there was an issue with the request or response.
// [Product Price Point Documentation](https://chargify.zendesk.com/hc/en-us/articles/4407755824155)
func (p *ProductPricePointsController) CreateProductPricePoint(
    ctx context.Context,
    productId models.CreateProductPricePointProductId,
    body *models.CreateProductPricePointRequest) (
    models.ApiResponse[models.ProductPricePointResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/products/%v/price_points.json", productId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewProductPricePointErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListProductPricePointsInput represents the input of the ListProductPricePoints endpoint.
type ListProductPricePointsInput struct {
    // The id or handle of the product. When using the handle, it must be prefixed with `handle:`
    ProductId      models.ListProductPricePointsInputProductId 
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page           *int                                        
    // This parameter indicates how many records to fetch in each request. Default value is 10. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    PerPage        *int                                        
    // When fetching a product's price points, if you have defined multiple currencies at the site level, you can optionally pass the ?currency_prices=true query param to include an array of currency price data in the response. If the product price point is set to use_site_exchange_rate: true, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.
    CurrencyPrices *bool                                       
    // Use in query: `filter[type]=catalog,default`.
    FilterType     []models.PricePointType                     
}

// ListProductPricePoints takes context, productId, page, perPage, currencyPrices, filterType as parameters and
// returns an models.ApiResponse with models.ListProductPricePointsResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to retrieve a list of product price points.
func (p *ProductPricePointsController) ListProductPricePoints(
    ctx context.Context,
    input ListProductPricePointsInput) (
    models.ApiResponse[models.ListProductPricePointsResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/products/%v/price_points.json", input.ProductId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.CurrencyPrices != nil {
        req.QueryParam("currency_prices", *input.CurrencyPrices)
    }
    if input.FilterType != nil {
        req.QueryParam("filter[type]", input.FilterType)
    }
    
    var result models.ListProductPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListProductPricePointsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateProductPricePoint takes context, productId, pricePointId, body as parameters and
// returns an models.ApiResponse with models.ProductPricePointResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to update a product price point.
// Note: Custom product price points are not able to be updated.
func (p *ProductPricePointsController) UpdateProductPricePoint(
    ctx context.Context,
    productId models.UpdateProductPricePointProductId,
    pricePointId models.UpdateProductPricePointPricePointId,
    body *models.UpdateProductPricePointRequest) (
    models.ApiResponse[models.ProductPricePointResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/products/%v/price_points/%v.json", productId, pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadProductPricePoint takes context, productId, pricePointId, currencyPrices as parameters and
// returns an models.ApiResponse with models.ProductPricePointResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to retrieve details for a specific product price point.
func (p *ProductPricePointsController) ReadProductPricePoint(
    ctx context.Context,
    productId models.ReadProductPricePointProductId,
    pricePointId models.ReadProductPricePointPricePointId,
    currencyPrices *bool) (
    models.ApiResponse[models.ProductPricePointResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/products/%v/price_points/%v.json", productId, pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if currencyPrices != nil {
        req.QueryParam("currency_prices", *currencyPrices)
    }
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ArchiveProductPricePoint takes context, productId, pricePointId as parameters and
// returns an models.ApiResponse with models.ProductPricePointResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to archive a product price point.
func (p *ProductPricePointsController) ArchiveProductPricePoint(
    ctx context.Context,
    productId models.ArchiveProductPricePointProductId,
    pricePointId models.ArchiveProductPricePointPricePointId) (
    models.ApiResponse[models.ProductPricePointResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/products/%v/price_points/%v.json", productId, pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UnarchiveProductPricePoint takes context, productId, pricePointId as parameters and
// returns an models.ApiResponse with models.ProductPricePointResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to unarchive an archived product price point.
func (p *ProductPricePointsController) UnarchiveProductPricePoint(
    ctx context.Context,
    productId int,
    pricePointId int) (
    models.ApiResponse[models.ProductPricePointResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "PATCH",
      fmt.Sprintf("/products/%v/price_points/%v/unarchive.json", productId, pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// PromoteProductPricePointToDefault takes context, productId, pricePointId as parameters and
// returns an models.ApiResponse with models.ProductResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to make a product price point the default for the product.
// Note: Custom product price points are not able to be set as the default for a product.
func (p *ProductPricePointsController) PromoteProductPricePointToDefault(
    ctx context.Context,
    productId int,
    pricePointId int) (
    models.ApiResponse[models.ProductResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "PATCH",
      fmt.Sprintf("/products/%v/price_points/%v/default.json", productId, pricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.ProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// BulkCreateProductPricePoints takes context, productId, body as parameters and
// returns an models.ApiResponse with models.BulkCreateProductPricePointsResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to create multiple product price points in one request.
func (p *ProductPricePointsController) BulkCreateProductPricePoints(
    ctx context.Context,
    productId int,
    body *models.BulkCreateProductPricePointsRequest) (
    models.ApiResponse[models.BulkCreateProductPricePointsResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/products/%v/price_points/bulk.json", productId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'."},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.BulkCreateProductPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.BulkCreateProductPricePointsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateProductCurrencyPrices takes context, productPricePointId, body as parameters and
// returns an models.ApiResponse with models.CurrencyPricesResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to create currency prices for a given currency that has been defined on the site level in your settings.
// When creating currency prices, they need to mirror the structure of your primary pricing. If the product price point defines a trial and/or setup fee, each currency must also define a trial and/or setup fee.
// Note: Currency Prices are not able to be created for custom product price points.
func (p *ProductPricePointsController) CreateProductCurrencyPrices(
    ctx context.Context,
    productPricePointId int,
    body *models.CreateProductCurrencyPricesRequest) (
    models.ApiResponse[models.CurrencyPricesResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_price_points/%v/currency_prices.json", productPricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorArrayMapResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.CurrencyPricesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CurrencyPricesResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateProductCurrencyPrices takes context, productPricePointId, body as parameters and
// returns an models.ApiResponse with models.CurrencyPricesResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to update the `price`s of currency prices for a given currency that exists on the product price point.
// When updating the pricing, it needs to mirror the structure of your primary pricing. If the product price point defines a trial and/or setup fee, each currency must also define a trial and/or setup fee.
// Note: Currency Prices are not able to be updated for custom product price points.
func (p *ProductPricePointsController) UpdateProductCurrencyPrices(
    ctx context.Context,
    productPricePointId int,
    body *models.UpdateCurrencyPricesRequest) (
    models.ApiResponse[models.CurrencyPricesResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/product_price_points/%v/currency_prices.json", productPricePointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorArrayMapResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.CurrencyPricesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CurrencyPricesResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListAllProductPricePointsInput represents the input of the ListAllProductPricePoints endpoint.
type ListAllProductPricePointsInput struct {
    // Controls the order in which results are returned.
    // Use in query `direction=asc`.
    Direction *models.SortingDirection               
    // Filter to use for List PricePoints operations
    Filter    *models.ListPricePointsFilter          
    // Allows including additional data in the response. Use in query: `include=currency_prices`.
    Include   *models.ListProductsPricePointsInclude 
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page      *int                                   
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage   *int                                   
}

// ListAllProductPricePoints takes context, direction, filter, include, page, perPage as parameters and
// returns an models.ApiResponse with models.ListProductPricePointsResponse data and
// an error if there was an issue with the request or response.
// This method allows retrieval of a list of Products Price Points belonging to a Site.
func (p *ProductPricePointsController) ListAllProductPricePoints(
    ctx context.Context,
    input ListAllProductPricePointsInput) (
    models.ApiResponse[models.ListProductPricePointsResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/products_price_points.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    if input.Direction != nil {
        req.QueryParam("direction", *input.Direction)
    }
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    if input.Include != nil {
        req.QueryParam("include", *input.Include)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    var result models.ListProductPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListProductPricePointsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
