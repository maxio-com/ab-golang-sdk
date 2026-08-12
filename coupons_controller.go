// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
)

// CouponsController represents a controller struct.
type CouponsController struct {
    baseController
}

// NewCouponsController creates a new instance of CouponsController.
// It takes a baseController as a parameter and returns a pointer to the CouponsController.
func NewCouponsController(baseController baseController) *CouponsController {
    couponsController := CouponsController{baseController: baseController}
    return &couponsController
}

// CreateCoupon takes context, productFamilyId, body as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// Creates a coupon under the specified product family.
// You can create either a flat amount coupon, by specifying `amount_in_cents`, or percentage coupon by specifying `percentage`.
// See [Apply Coupons to Subscriptions](https://maxio.zendesk.com/hc/en-us/articles/24261259337101-Coupons-and-Subscriptions) for information on applying a coupon to a subscription in the Advanced Billing UI.
func (c *CouponsController) CreateCoupon(
    ctx context.Context,
    productFamilyId int,
    body *models.CouponRequest) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "POST", "/product_families/%v/coupons.json")
    req.AppendTemplateParams(productFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListCouponsForProductFamilyInput represents the input of the ListCouponsForProductFamily endpoint.
type ListCouponsForProductFamilyInput struct {
    // The Advanced Billing id of the product family to which the coupon belongs
    ProductFamilyId int                       
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page            *int                      
    // This parameter indicates how many records to fetch in each request. Default value is 30. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage         *int                      
    // Filter to use for List Coupons operations
    Filter          *models.ListCouponsFilter 
    // (Optional) If you have defined multiple currencies at the site level, you can pass `?currency_prices=true` to include an array of currency price data in the response. Use in query `currency_prices=true`.
    CurrencyPrices  *bool                     
}

// ListCouponsForProductFamily takes context, productFamilyId, page, perPage, filter, currencyPrices as parameters and
// returns an models.ApiResponse with []models.CouponResponse data and
// an error if there was an issue with the request or response.
// Lists coupons for a specific product family in a site.
func (c *CouponsController) ListCouponsForProductFamily(
    ctx context.Context,
    input ListCouponsForProductFamilyInput) (
    models.ApiResponse[[]models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/product_families/%v/coupons.json")
    req.AppendTemplateParams(input.ProductFamilyId)
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    if input.CurrencyPrices != nil {
        req.QueryParam("currency_prices", *input.CurrencyPrices)
    }
    
    var result []models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.CouponResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// FindCoupon takes context, productFamilyId, code, currencyPrices as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// Searches for a coupon by code.
// If you have more than one product family and if the coupon you are trying to find does not belong to the default product family in your site, you need to specify (either in the URL or as a query string param) the `product_family_id`.
func (c *CouponsController) FindCoupon(
    ctx context.Context,
    productFamilyId *int,
    code *string,
    currencyPrices *bool) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/coupons/find.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    if productFamilyId != nil {
        req.QueryParam("product_family_id", *productFamilyId)
    }
    if code != nil {
        req.QueryParam("code", *code)
    }
    if currencyPrices != nil {
        req.QueryParam("currency_prices", *currencyPrices)
    }
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadCoupon takes context, productFamilyId, couponId, currencyPrices as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// Returns a coupon by its system-assigned ID. You must identify the Coupon in this call by the ID parameter assigned to it.
// If instead you would like to find a Coupon using a Coupon code, use the [Find Coupon]($e/Coupons/findCoupon) endpoint.
// If the coupon is set to `use_site_exchange_rate: true`, it returns pricing based on the current exchange rate. If the flag is set to false, it returns all of the defined prices for each currency.
func (c *CouponsController) ReadCoupon(
    ctx context.Context,
    productFamilyId int,
    couponId int,
    currencyPrices *bool) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/product_families/%v/coupons/%v.json")
    req.AppendTemplateParams(productFamilyId, couponId)
    req.Authenticate(NewAuth("BasicAuth"))
    if currencyPrices != nil {
        req.QueryParam("currency_prices", *currencyPrices)
    }
    
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateCoupon takes context, productFamilyId, couponId, body as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// Updates a coupon. 
// You can restrict a coupon to only apply to specific products / components by optionally passing in hashes of `restricted_products` and/or `restricted_components` in the format:
// `{ "<product/component_id>": boolean_value }`
func (c *CouponsController) UpdateCoupon(
    ctx context.Context,
    productFamilyId int,
    couponId int,
    body *models.CouponRequest) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "PUT", "/product_families/%v/coupons/%v.json")
    req.AppendTemplateParams(productFamilyId, couponId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ArchiveCoupon takes context, productFamilyId, couponId as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// Archives a coupon, making it unavailable for future use while remaining active on existing subscriptions.
// Archiving makes that Coupon unavailable for future use, but allows it to remain attached and functional on existing Subscriptions that are using it.
// The `archived_at` date and time will be assigned.
func (c *CouponsController) ArchiveCoupon(
    ctx context.Context,
    productFamilyId int,
    couponId int) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "DELETE", "/product_families/%v/coupons/%v.json")
    req.AppendTemplateParams(productFamilyId, couponId)
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListCouponsInput represents the input of the ListCoupons endpoint.
type ListCouponsInput struct {
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page           *int                      
    // This parameter indicates how many records to fetch in each request. Default value is 30. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage        *int                      
    // Filter to use for List Coupons operations
    Filter         *models.ListCouponsFilter 
    // (Optional) If you have defined multiple currencies at the site level, you can pass `?currency_prices=true` to include an array of currency price data in the response. Use in query `currency_prices=true`.
    CurrencyPrices *bool                     
}

// ListCoupons takes context, page, perPage, filter, currencyPrices as parameters and
// returns an models.ApiResponse with []models.CouponResponse data and
// an error if there was an issue with the request or response.
// Lists coupons for a site.
func (c *CouponsController) ListCoupons(
    ctx context.Context,
    input ListCouponsInput) (
    models.ApiResponse[[]models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/coupons.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    if input.CurrencyPrices != nil {
        req.QueryParam("currency_prices", *input.CurrencyPrices)
    }
    var result []models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.CouponResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadCouponUsage takes context, productFamilyId, couponId as parameters and
// returns an models.ApiResponse with []models.CouponUsage data and
// an error if there was an issue with the request or response.
// Lists coupon usage details, one entry per product.
func (c *CouponsController) ReadCouponUsage(
    ctx context.Context,
    productFamilyId int,
    couponId int) (
    models.ApiResponse[[]models.CouponUsage],
    error) {
    req := c.prepareRequest(ctx, "GET", "/product_families/%v/coupons/%v/usage.json")
    req.AppendTemplateParams(productFamilyId, couponId)
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result []models.CouponUsage
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.CouponUsage](decoder)
    return models.NewApiResponse(result, resp), err
}

// ValidateCoupon takes context, code, productFamilyId as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// Verifies whether a specific coupon code is valid. This method is useful for validating coupon codes that are entered by a customer.
// If you have more than one product family and if the coupon you are validating does not belong to the first product family in your site, you need to specify the product family, either in the URL or as a query string param. This can be done by supplying the id or the handle in the `handle:my-family` format.
// Supplying the `product_family_handle` in the URL:
// ```
// https://<subdomain>.chargify.com/product_families/handle:<product_family_handle>/coupons/validate.<format>?code=<coupon_code>
// ```
// Supplying the `product_family_id` as a query parameter:
// ```
// https://<subdomain>.chargify.com/coupons/validate.<format>?code=<coupon_code>&product_family_id=<id>
// ```
func (c *CouponsController) ValidateCoupon(
    ctx context.Context,
    code string,
    productFamilyId *int) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/coupons/validate.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found: '{$response.body}'", Unmarshaller: errors.NewSingleStringErrorResponse},
    })
    req.QueryParam("code", code)
    if productFamilyId != nil {
        req.QueryParam("product_family_id", *productFamilyId)
    }
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrUpdateCouponCurrencyPrices takes context, couponId, body as parameters and
// returns an models.ApiResponse with models.CouponCurrencyResponse data and
// an error if there was an issue with the request or response.
// Creates and/or updates currency prices for an existing coupon. Multiple prices can be created or updated in a single request but each of the currencies must be defined on the site level already and the coupon must be an amount-based coupon, not percentage.
// Currency pricing for coupons must mirror the setup of the primary coupon pricing - if the primary coupon is percentage based, you will not be able to define pricing in non-primary currencies.
func (c *CouponsController) CreateOrUpdateCouponCurrencyPrices(
    ctx context.Context,
    couponId int,
    body *models.CouponCurrencyRequest) (
    models.ApiResponse[models.CouponCurrencyResponse],
    error) {
    req := c.prepareRequest(ctx, "PUT", "/coupons/%v/currency_prices.json")
    req.AppendTemplateParams(couponId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorStringMapResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.CouponCurrencyResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponCurrencyResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateCouponSubcodes takes context, couponId, body as parameters and
// returns an models.ApiResponse with models.CouponSubcodesResponse data and
// an error if there was an issue with the request or response.
// Creates subcodes for an existing coupon.
// Coupon Subcodes allow you to create a set of unique codes that allow you to expand the use of one coupon.
// For example:
// Master Coupon Code:
// + SPRING2020
// Coupon Subcodes:
// + SPRING90210
// + DP80302
// + SPRINGBALTIMORE
// When creating a coupon subcode, you must specify a coupon to attach it to using the coupon_id. Valid coupon subcodes are all capital letters, contain only letters and numbers, and do not have any spaces. Lowercase letters are capitalized before the subcode is created.
// Note: If you are using any of the allowed special characters ("%", "@", "+", "-", "_", and "."), you must encode them for use in the URL.
// % to %25
// @ to %40
// + to %2B
// - to %2D
// _ to %5F
// . to %2E
// So, if the coupon subcode is `20%OFF`, the URL to delete this coupon subcode would be: `https://<subdomain>.chargify.com/coupons/567/codes/20%25OFF.<format>`.
// For more information on coupon codes and applying coupons to subscriptions, see [Coupon Codes](https://maxio.zendesk.com/hc/en-us/articles/24261208729229-Coupon-Codes) and [Coupons and Subscriptions](https://maxio.zendesk.com/hc/en-us/articles/24261259337101-Coupons-and-Subscriptions).
func (c *CouponsController) CreateCouponSubcodes(
    ctx context.Context,
    couponId int,
    body *models.CouponSubcodes) (
    models.ApiResponse[models.CouponSubcodesResponse],
    error) {
    req := c.prepareRequest(ctx, "POST", "/coupons/%v/codes.json")
    req.AppendTemplateParams(couponId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.CouponSubcodesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponSubcodesResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListCouponSubcodesInput represents the input of the ListCouponSubcodes endpoint.
type ListCouponSubcodesInput struct {
    // The Advanced Billing id of the coupon
    CouponId int  
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page     *int 
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage  *int 
}

// ListCouponSubcodes takes context, couponId, page, perPage as parameters and
// returns an models.ApiResponse with models.CouponSubcodes data and
// an error if there was an issue with the request or response.
// Lists the subcodes attached to a coupon.
func (c *CouponsController) ListCouponSubcodes(
    ctx context.Context,
    input ListCouponSubcodesInput) (
    models.ApiResponse[models.CouponSubcodes],
    error) {
    req := c.prepareRequest(ctx, "GET", "/coupons/%v/codes.json")
    req.AppendTemplateParams(input.CouponId)
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    
    var result models.CouponSubcodes
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponSubcodes](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateCouponSubcodes takes context, couponId, body as parameters and
// returns an models.ApiResponse with models.CouponSubcodesResponse data and
// an error if there was an issue with the request or response.
// Updates the subcodes for a coupon, replacing all existing subcodes with the new list.
// Send an array of new coupon subcodes.
// **Note**: All current subcodes for that Coupon will be deleted first, and replaced with the list of subcodes sent to this endpoint.
// The response will contain:
// + The created subcodes,
// + Subcodes that were not created because they already exist,
// + Any subcodes not created because they are invalid.
func (c *CouponsController) UpdateCouponSubcodes(
    ctx context.Context,
    couponId int,
    body *models.CouponSubcodes) (
    models.ApiResponse[models.CouponSubcodesResponse],
    error) {
    req := c.prepareRequest(ctx, "PUT", "/coupons/%v/codes.json")
    req.AppendTemplateParams(couponId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.CouponSubcodesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponSubcodesResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteCouponSubcode takes context, couponId, subcode as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Deletes a specific subcode from a coupon.
// ## Example
// Given a coupon with an ID of 567, and a coupon subcode of 20OFF, the URL to `DELETE` this coupon subcode would be:
// ```
// http://subdomain.chargify.com/coupons/567/codes/20OFF.<format>
// ```
// Note: If you are using any of the allowed special characters (“%”, “@”, “+”, “-”, “_”, and “.”), you must encode them for use in the URL.
// | Special character | Encoding |
// |-------------------|----------|
// | %                 | %25      |
// | @                 | %40      |
// | +                 | %2B      |
// | –                 | %2D      |
// | _                 | %5F      |
// | .                 | %2E      |
// ## Percent Encoding Example
// Or if the coupon subcode is 20%OFF, the URL to delete this coupon subcode would be: @https://<subdomain>.chargify.com/coupons/567/codes/20%25OFF.<format>.
func (c *CouponsController) DeleteCouponSubcode(
    ctx context.Context,
    couponId int,
    subcode string) (
    *http.Response,
    error) {
    req := c.prepareRequest(ctx, "DELETE", "/coupons/%v/codes/%v.json")
    req.AppendTemplateParams(couponId, subcode)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
    })
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
