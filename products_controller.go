// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "context"
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
// Use this method to create a product within your Advanced Billing site.
// + [Products Documentation](https://maxio.zendesk.com/hc/en-us/articles/24261090117645-Products-Overview)
// + [Changing a Subscription's Product](https://maxio.zendesk.com/hc/en-us/articles/24252069837581-Product-Changes-and-Migrations)
func (p *ProductsController) CreateProduct(
    ctx context.Context,
    productFamilyId string,
    body *models.CreateOrUpdateProductRequest) (
    models.ApiResponse[models.ProductResponse],
    error) {
    req := p.prepareRequest(ctx, "POST", "/product_families/%v/products.json")
    req.AppendTemplateParams(productFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
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
// This endpoint allows you to read the current details of a product that you've created in Advanced Billing.
func (p *ProductsController) ReadProduct(
    ctx context.Context,
    productId int) (
    models.ApiResponse[models.ProductResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/products/%v.json")
    req.AppendTemplateParams(productId)
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
    req := p.prepareRequest(ctx, "PUT", "/products/%v.json")
    req.AppendTemplateParams(productId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
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
    req := p.prepareRequest(ctx, "DELETE", "/products/%v.json")
    req.AppendTemplateParams(productId)
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
    req := p.prepareRequest(ctx, "GET", "/products/handle/%v.json")
    req.AppendTemplateParams(apiHandle)
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.ProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListProductsInput represents the input of the ListProducts endpoint.
type ListProductsInput struct {
    // The type of filter you would like to apply to your search.
    // Use in query: `date_field=created_at`.
    DateField       *models.BasicDateField      
    // Filter to use for List Products operations
    Filter          *models.ListProductsFilter  
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
    EndDate         *time.Time                  
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of end_date.
    EndDatetime     *time.Time                  
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
    StartDate       *time.Time                  
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of start_date.
    StartDatetime   *time.Time                  
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page            *int                        
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage         *int                        
    // Include archived products. Use in query: `include_archived=true`.
    IncludeArchived *bool                       
    // Allows including additional data in the response. Use in query `include=prepaid_product_price_point`.
    Include         *models.ListProductsInclude 
}

// ListProducts takes context, dateField, filter, endDate, endDatetime, startDate, startDatetime, page, perPage, includeArchived, include as parameters and
// returns an models.ApiResponse with []models.ProductResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve a list of Products belonging to a Site.
func (p *ProductsController) ListProducts(
    ctx context.Context,
    input ListProductsInput) (
    models.ApiResponse[[]models.ProductResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/products.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    if input.DateField != nil {
        req.QueryParam("date_field", *input.DateField)
    }
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    if input.EndDate != nil {
        req.QueryParam("end_date", input.EndDate.Format(models.DEFAULT_DATE))
    }
    if input.EndDatetime != nil {
        req.QueryParam("end_datetime", input.EndDatetime.Format(time.RFC3339))
    }
    if input.StartDate != nil {
        req.QueryParam("start_date", input.StartDate.Format(models.DEFAULT_DATE))
    }
    if input.StartDatetime != nil {
        req.QueryParam("start_datetime", input.StartDatetime.Format(time.RFC3339))
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.IncludeArchived != nil {
        req.QueryParam("include_archived", *input.IncludeArchived)
    }
    if input.Include != nil {
        req.QueryParam("include", *input.Include)
    }
    var result []models.ProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
