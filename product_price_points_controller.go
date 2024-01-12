package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "time"
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
    productId interface{},
    body *models.CreateProductPricePointRequest) (
    models.ApiResponse[models.ProductPricePointResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/products/%v/price_points.json", productId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewProductPricePointErrorResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ListProductPricePoints takes context, productId, page, perPage, currencyPrices, filterType as parameters and
// returns an models.ApiResponse with models.ListProductPricePointsResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to retrieve a list of product price points.
func (p *ProductPricePointsController) ListProductPricePoints(
    ctx context.Context,
    productId interface{},
    page *int,
    perPage *int,
    currencyPrices *bool,
    filterType []models.PricePointType) (
    models.ApiResponse[models.ListProductPricePointsResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/products/%v/price_points.json", productId),
    )
    req.Authenticate(true)
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if currencyPrices != nil {
        req.QueryParam("currency_prices", *currencyPrices)
    }
    if filterType != nil {
        req.QueryParam("filter[type]", filterType)
    }
    
    var result models.ListProductPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListProductPricePointsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// UpdateProductPricePoint takes context, productId, pricePointId, body as parameters and
// returns an models.ApiResponse with models.ProductPricePointResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to update a product price point.
// Note: Custom product price points are not able to be updated.
func (p *ProductPricePointsController) UpdateProductPricePoint(
    ctx context.Context,
    productId interface{},
    pricePointId interface{},
    body *models.UpdateProductPricePointRequest) (
    models.ApiResponse[models.ProductPricePointResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/products/%v/price_points/%v.json", productId, pricePointId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ReadProductPricePoint takes context, productId, pricePointId, currencyPrices as parameters and
// returns an models.ApiResponse with models.ProductPricePointResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to retrieve details for a specific product price point.
func (p *ProductPricePointsController) ReadProductPricePoint(
    ctx context.Context,
    productId interface{},
    pricePointId interface{},
    currencyPrices *bool) (
    models.ApiResponse[models.ProductPricePointResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/products/%v/price_points/%v.json", productId, pricePointId),
    )
    req.Authenticate(true)
    if currencyPrices != nil {
        req.QueryParam("currency_prices", *currencyPrices)
    }
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ArchiveProductPricePoint takes context, productId, pricePointId as parameters and
// returns an models.ApiResponse with models.ProductPricePointResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to archive a product price point.
func (p *ProductPricePointsController) ArchiveProductPricePoint(
    ctx context.Context,
    productId interface{},
    pricePointId interface{}) (
    models.ApiResponse[models.ProductPricePointResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/products/%v/price_points/%v.json", productId, pricePointId),
    )
    req.Authenticate(true)
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
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
    req.Authenticate(true)
    
    var result models.ProductPricePointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
    
    var result models.ProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// CreateProductPricePoints takes context, productId, body as parameters and
// returns an models.ApiResponse with models.BulkCreateProductPricePointsResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to create multiple product price points in one request.
func (p *ProductPricePointsController) CreateProductPricePoints(
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
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.BulkCreateProductPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.BulkCreateProductPricePointsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewApiError(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// CreateProductCurrencyPrices takes context, productPricePointId, body as parameters and
// returns an models.ApiResponse with models.ProductPricePointCurrencyPrice data and
// an error if there was an issue with the request or response.
// This endpoint allows you to create currency prices for a given currency that has been defined on the site level in your settings.
// When creating currency prices, they need to mirror the structure of your primary pricing. If the product price point defines a trial and/or setup fee, each currency must also define a trial and/or setup fee.
// Note: Currency Prices are not able to be created for custom product price points.
func (p *ProductPricePointsController) CreateProductCurrencyPrices(
    ctx context.Context,
    productPricePointId int,
    body *models.CreateProductCurrencyPricesRequest) (
    models.ApiResponse[models.ProductPricePointCurrencyPrice],
    error) {
    req := p.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_price_points/%v/currency_prices.json", productPricePointId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ProductPricePointCurrencyPrice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointCurrencyPrice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorMapResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// UpdateProductCurrencyPrices takes context, productPricePointId, body as parameters and
// returns an models.ApiResponse with models.ProductPricePointCurrencyPrice data and
// an error if there was an issue with the request or response.
// This endpoint allows you to update the `price`s of currency prices for a given currency that exists on the product price point.
// When updating the pricing, it needs to mirror the structure of your primary pricing. If the product price point defines a trial and/or setup fee, each currency must also define a trial and/or setup fee.
// Note: Currency Prices are not able to be updated for custom product price points.
func (p *ProductPricePointsController) UpdateProductCurrencyPrices(
    ctx context.Context,
    productPricePointId int,
    body *models.UpdateCurrencyPricesRequest) (
    models.ApiResponse[models.ProductPricePointCurrencyPrice],
    error) {
    req := p.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/product_price_points/%v/currency_prices.json", productPricePointId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ProductPricePointCurrencyPrice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductPricePointCurrencyPrice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorMapResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ListAllProductPricePoints takes context, direction, filterArchivedAt, filterDateField, filterEndDate, filterEndDatetime, filterIds, filterStartDate, filterStartDatetime, filterType, include, page, perPage as parameters and
// returns an models.ApiResponse with models.ListProductPricePointsResponse data and
// an error if there was an issue with the request or response.
// This method allows retrieval of a list of Products Price Points belonging to a Site.
func (p *ProductPricePointsController) ListAllProductPricePoints(
    ctx context.Context,
    direction *models.SortingDirection,
    filterArchivedAt *models.IncludeNotNull,
    filterDateField *models.BasicDateField,
    filterEndDate *time.Time,
    filterEndDatetime *time.Time,
    filterIds []int,
    filterStartDate *time.Time,
    filterStartDatetime *time.Time,
    filterType []models.PricePointType,
    include *models.ListProductsPricePointsInclude,
    page *int,
    perPage *int) (
    models.ApiResponse[models.ListProductPricePointsResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/products_price_points.json")
    req.Authenticate(true)
    if direction != nil {
        req.QueryParam("direction", *direction)
    }
    if filterArchivedAt != nil {
        req.QueryParam("filter[archived_at]", *filterArchivedAt)
    }
    if filterDateField != nil {
        req.QueryParam("filter[date_field]", *filterDateField)
    }
    if filterEndDate != nil {
        req.QueryParam("filter[end_date]", filterEndDate.Format(models.DEFAULT_DATE))
    }
    if filterEndDatetime != nil {
        req.QueryParam("filter[end_datetime]", filterEndDatetime.Format(time.RFC3339))
    }
    if filterIds != nil {
        req.QueryParam("filter[ids]", filterIds)
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
    if include != nil {
        req.QueryParam("include", *include)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    var result models.ListProductPricePointsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListProductPricePointsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}
