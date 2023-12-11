package ab_golang_sdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"maxioadvancedbilling/errors"
	"maxioadvancedbilling/models"
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

// ListProductsForProductFamily takes context, productFamilyId, page, perPage, dateField, startDate, endDate, startDatetime, endDatetime, includeArchived, include, filterPrepaidProductPricePointProductPricePointId, filterUseSiteExchangeRate as parameters and
// returns an models.ApiResponse with []models.ProductResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve a list of Products belonging to a Product Family.
func (p *ProductFamiliesController) ListProductsForProductFamily(
	ctx context.Context,
	productFamilyId int,
	page *int,
	perPage *int,
	dateField *models.BasicDateFieldEnum,
	startDate *string,
	endDate *string,
	startDatetime *string,
	endDatetime *string,
	includeArchived *bool,
	include *models.ListProductsIncludeEnum,
	filterPrepaidProductPricePointProductPricePointId *models.IncludeNotNullEnum,
	filterUseSiteExchangeRate *bool) (
	models.ApiResponse[[]models.ProductResponse],
	error) {
	req := p.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/product_families/%v/products.json", productFamilyId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
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
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ProductResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
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
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}
	var result models.ProductFamilyResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ProductFamilyResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// ListProductFamilies takes context, dateField, startDate, endDate, startDatetime, endDatetime as parameters and
// returns an models.ApiResponse with []models.ProductFamilyResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve a list of Product Families for a site.
func (p *ProductFamiliesController) ListProductFamilies(
	ctx context.Context,
	dateField *models.BasicDateFieldEnum,
	startDate *string,
	endDate *string,
	startDatetime *string,
	endDatetime *string) (
	models.ApiResponse[[]models.ProductFamilyResponse],
	error) {
	req := p.prepareRequest(ctx, "GET", "/product_families.json")
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
	var result []models.ProductFamilyResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ProductFamilyResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

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
	req.Authenticate(true)

	var result models.ProductFamilyResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ProductFamilyResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
