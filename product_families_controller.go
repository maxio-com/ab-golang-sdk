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

// ProductFamiliesController represents a controller struct.
type ProductFamiliesController struct {
    baseController
}

// NewProductFamiliesController creates a new instance of ProductFamiliesController.
// It takes a baseController as a parameter and returns a pointer to the ProductFamiliesController.
func NewProductFamiliesController(baseController baseController) *ProductFamiliesController {
    productFamiliesController := ProductFamiliesController{baseController: baseController}
    return &productFamiliesController
}

// ListProductsForProductFamilyInput represents the input of the ListProductsForProductFamily endpoint.
type ListProductsForProductFamilyInput struct {
    // Either the product family's id or its handle prefixed with `handle:`
    ProductFamilyId string                      
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page            *int                        
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage         *int                        
    // The type of filter you would like to apply to your search.
    // Use in query: `date_field=created_at`.
    DateField       *models.BasicDateField      
    // Filter to use for List Products operations
    Filter          *models.ListProductsFilter  
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
    StartDate       *time.Time                  
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
    EndDate         *time.Time                  
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
    StartDatetime   *time.Time                  
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.
    EndDatetime     *time.Time                  
    // Include archived products
    IncludeArchived *bool                       
    // Allows including additional data in the response. Use in query `include=prepaid_product_price_point`.
    Include         *models.ListProductsInclude 
}

// ListProductsForProductFamily takes context, productFamilyId, page, perPage, dateField, filter, startDate, endDate, startDatetime, endDatetime, includeArchived, include as parameters and
// returns an models.ApiResponse with []models.ProductResponse data and
// an error if there was an issue with the request or response.
// Retrieves a list of Products belonging to a Product Family.
func (p *ProductFamiliesController) ListProductsForProductFamily(
    ctx context.Context,
    input ListProductsForProductFamilyInput) (
    models.ApiResponse[[]models.ProductResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/product_families/%v/products.json")
    req.AppendTemplateParams(input.ProductFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
    })
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.DateField != nil {
        req.QueryParam("date_field", *input.DateField)
    }
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
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

// CreateProductFamily takes context, body as parameters and
// returns an models.ApiResponse with models.ProductFamilyResponse data and
// an error if there was an issue with the request or response.
// Creates a Product Family within your Advanced Billing site. Create a Product Family to act as a container for your products, components and coupons.
// Full documentation on how Product Families operate within the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24261098936205-Product-Families).
func (p *ProductFamiliesController) CreateProductFamily(
    ctx context.Context,
    body *models.CreateProductFamilyRequest) (
    models.ApiResponse[models.ProductFamilyResponse],
    error) {
    req := p.prepareRequest(ctx, "POST", "/product_families.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.ProductFamilyResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductFamilyResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListProductFamiliesInput represents the input of the ListProductFamilies endpoint.
type ListProductFamiliesInput struct {
    // The type of filter you would like to apply to your search.
    // Use in query: `date_field=created_at`.
    DateField     *models.BasicDateField 
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
    StartDate     *time.Time             
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
    EndDate       *time.Time             
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
    StartDatetime *time.Time             
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.
    EndDatetime   *time.Time             
}

// ListProductFamilies takes context, dateField, startDate, endDate, startDatetime, endDatetime as parameters and
// returns an models.ApiResponse with []models.ProductFamilyResponse data and
// an error if there was an issue with the request or response.
// Retrieve a list of Product Families for a site.
func (p *ProductFamiliesController) ListProductFamilies(
    ctx context.Context,
    input ListProductFamiliesInput) (
    models.ApiResponse[[]models.ProductFamilyResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/product_families.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
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
    var result []models.ProductFamilyResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ProductFamilyResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadProductFamily takes context, id as parameters and
// returns an models.ApiResponse with models.ProductFamilyResponse data and
// an error if there was an issue with the request or response.
// Retrieves a Product Family via the `product_family_id`. The response will contain a Product Family object.
// The product family can be specified either with the id number, or with the `handle:my-family` format.
func (p *ProductFamiliesController) ReadProductFamily(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.ProductFamilyResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/product_families/%v.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.ProductFamilyResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ProductFamilyResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
