package advancedbilling

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
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
	// The Chargify id of the product family to which the product belongs
	ProductFamilyId int
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
	// The type of filter you would like to apply to your search.
	// Use in query: `date_field=created_at`.
	DateField *models.BasicDateField
	// The start date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
	StartDate *string
	// The end date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
	EndDate *string
	// The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
	StartDatetime *string
	// The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.
	EndDatetime *string
	// Include archived products
	IncludeArchived *bool
	// Allows including additional data in the response. Use in query `include=prepaid_product_price_point`.
	Include *models.ListProductsInclude
	// Allows fetching products only if a prepaid product price point is present or not. To use this filter you also have to include the following param in the request `include=prepaid_product_price_point`. Use in query `filter[prepaid_product_price_point][product_price_point_id]=not_null`.
	FilterPrepaidProductPricePointProductPricePointId *models.IncludeNotNull
	// Allows fetching products with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`.
	FilterUseSiteExchangeRate *bool
}

// ListProductsForProductFamily takes context, productFamilyId, page, perPage, dateField, startDate, endDate, startDatetime, endDatetime, includeArchived, include, filterPrepaidProductPricePointProductPricePointId, filterUseSiteExchangeRate as parameters and
// returns an models.ApiResponse with []models.ProductResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve a list of Products belonging to a Product Family.
func (p *ProductFamiliesController) ListProductsForProductFamily(
	ctx context.Context,
	input ListProductsForProductFamilyInput) (
	models.ApiResponse[[]models.ProductResponse],
	error) {
	req := p.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/product_families/%v/products.json", input.ProductFamilyId),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {Message: "Not Found"},
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
	if input.Include != nil {
		req.QueryParam("include", *input.Include)
	}
	if input.FilterPrepaidProductPricePointProductPricePointId != nil {
		req.QueryParam("filter[prepaid_product_price_point][product_price_point_id]", *input.FilterPrepaidProductPricePointProductPricePointId)
	}
	if input.FilterUseSiteExchangeRate != nil {
		req.QueryParam("filter[use_site_exchange_rate]", *input.FilterUseSiteExchangeRate)
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
// This method will create a Product Family within your Chargify site. Create a Product Family to act as a container for your products, components and coupons.
// Full documentation on how Product Families operate within the Chargify UI can be located [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405369633421).
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
	DateField *models.BasicDateField
	// The start date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
	StartDate *string
	// The end date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
	EndDate *string
	// The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
	StartDatetime *string
	// The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.
	EndDatetime *string
}

// ListProductFamilies takes context, dateField, startDate, endDate, startDatetime, endDatetime as parameters and
// returns an models.ApiResponse with []models.ProductFamilyResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve a list of Product Families for a site.
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
// This method allows to retrieve a Product Family via the `product_family_id`. The response will contain a Product Family object.
// The product family can be specified either with the id number, or with the `handle:my-family` format.
func (p *ProductFamiliesController) ReadProductFamily(
	ctx context.Context,
	id int) (
	models.ApiResponse[models.ProductFamilyResponse],
	error) {
	req := p.prepareRequest(ctx, "GET", fmt.Sprintf("/product_families/%v.json", id))
	req.Authenticate(NewAuth("BasicAuth"))

	var result models.ProductFamilyResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ProductFamilyResponse](decoder)
	return models.NewApiResponse(result, resp), err
}
