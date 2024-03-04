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

// ProductsController represents a controller struct.
type ProductsController struct {
    baseController
}

// NewProductsController creates a new instance of ProductsController.
// It takes a baseController as a parameter and returns a pointer to the ProductsController.
func NewProductsController(baseController baseController) *ProductsController {
    productsController := ProductsController{baseController: baseController}
    return &productsController
}

// CreateProduct takes context, productFamilyId, body as parameters and
// returns an models.ApiResponse with models.ProductResponse data and
// an error if there was an issue with the request or response.
// Use this method to create a product within your Chargify site.
// + [Products Documentation](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405561405709)
// + [Changing a Subscription's Product](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404225334669-Product-Changes-Migrations)
func (p *ProductsController) CreateProduct(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateOrUpdateProductRequest) (
    models.ApiResponse[models.ProductResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_families/%v/products.json", productFamilyId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadProduct takes context, productId as parameters and
// returns an models.ApiResponse with models.ProductResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to read the current details of a product that you've created in Chargify.
func (p *ProductsController) ReadProduct(
    ctx context.Context,
    productId int) (
    models.ApiResponse[models.ProductResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", fmt.Sprintf("/products/%v.json", productId))
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.ProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateProduct takes context, productId, body as parameters and
// returns an models.ApiResponse with models.ProductResponse data and
// an error if there was an issue with the request or response.
// Use this method to change aspects of an existing product.
// ### Input Attributes Update Notes
// + `update_return_params` The parameters we will append to your `update_return_url`. See Return URLs and Parameters
// ### Product Price Point
// Updating a product using this endpoint will create a new price point and set it as the default price point for this product. If you should like to update an existing product price point, that must be done separately.
func (p *ProductsController) UpdateProduct(
    ctx context.Context,
    productId int,
    body *models.CreateOrUpdateProductRequest) (
    models.ApiResponse[models.ProductResponse],
    error) {
    req := p.prepareRequest(ctx, "PUT", fmt.Sprintf("/products/%v.json", productId))
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ArchiveProduct takes context, productId as parameters and
// returns an models.ApiResponse with models.ProductResponse data and
// an error if there was an issue with the request or response.
// Sending a DELETE request to this endpoint will archive the product. All current subscribers will be unffected; their subscription/purchase will continue to be charged monthly.
// This will restrict the option to chose the product for purchase via the Billing Portal, as well as disable Public Signup Pages for the product.
func (p *ProductsController) ArchiveProduct(
    ctx context.Context,
    productId int) (
    models.ApiResponse[models.ProductResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/products/%v.json", productId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.ProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadProductByHandle takes context, apiHandle as parameters and
// returns an models.ApiResponse with models.ProductResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve a Product object by its `api_handle`.
func (p *ProductsController) ReadProductByHandle(
    ctx context.Context,
    apiHandle string) (
    models.ApiResponse[models.ProductResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/products/handle/%v.json", apiHandle),
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

// ListProducts takes context, dateField, endDate, endDatetime, startDate, startDatetime, page, perPage, includeArchived, include, filterPrepaidProductPricePointProductPricePointId, filterUseSiteExchangeRate as parameters and
// returns an models.ApiResponse with []models.ProductResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve a list of Products belonging to a Site.
func (p *ProductsController) ListProducts(
    ctx context.Context,
    dateField *models.BasicDateField,
    endDate *time.Time,
    endDatetime *time.Time,
    startDate *time.Time,
    startDatetime *time.Time,
    page *int,
    perPage *int,
    includeArchived *bool,
    include *models.ListProductsInclude,
    filterPrepaidProductPricePointProductPricePointId *models.IncludeNotNull,
    filterUseSiteExchangeRate *bool) (
    models.ApiResponse[[]models.ProductResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/products.json")
    req.Authenticate(NewAuth("BasicAuth"))
    if dateField != nil {
        req.QueryParam("date_field", *dateField)
    }
    if endDate != nil {
        req.QueryParam("end_date", endDate.Format(models.DEFAULT_DATE))
    }
    if endDatetime != nil {
        req.QueryParam("end_datetime", endDatetime.Format(time.RFC3339))
    }
    if startDate != nil {
        req.QueryParam("start_date", startDate.Format(models.DEFAULT_DATE))
    }
    if startDatetime != nil {
        req.QueryParam("start_datetime", startDatetime.Format(time.RFC3339))
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if includeArchived != nil {
        req.QueryParam("include_archived", *includeArchived)
    }
    if include != nil {
        req.QueryParam("include", *include)
    }
    if filterPrepaidProductPricePointProductPricePointId != nil {
        req.QueryParam("filter[prepaid_product_price_point][product_price_point_id]", *filterPrepaidProductPricePointProductPricePointId)
    }
    if filterUseSiteExchangeRate != nil {
        req.QueryParam("filter[use_site_exchange_rate]", *filterUseSiteExchangeRate)
    }
    var result []models.ProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
