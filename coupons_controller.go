package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
    "time"
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
// ## Coupons Documentation
// Coupons can be administered in the Chargify application or created via API. Please view our section on [creating coupons](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404742830733) for more information.
// Additionally, for documentation on how to apply a coupon to a subscription within the Chargify UI, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404761012877).
// ## Create Coupon
// This request will create a coupon, based on the provided information.
// When creating a coupon, you must specify a product family using the `product_family_id`. If no `product_family_id` is passed, the first product family available is used. You will also need to formulate your URL to cite the Product Family ID in your request.
// You can restrict a coupon to only apply to specific products / components by optionally passing in hashes of `restricted_products` and/or `restricted_components` in the format:
// `{ "<product/component_id>": boolean_value }`
func (c *CouponsController) CreateCoupon(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateOrUpdateCoupon) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/product_families/%v/coupons.json", productFamilyId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ListCouponsForProductFamily takes context, productFamilyId, page, perPage, filterDateField, filterEndDate, filterEndDatetime, filterStartDate, filterStartDatetime, filterIds, filterCodes, currencyPrices, filterUseSiteExchangeRate as parameters and
// returns an models.ApiResponse with []models.CouponResponse data and
// an error if there was an issue with the request or response.
// List coupons for a specific Product Family in a Site.
// If the coupon is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.
func (c *CouponsController) ListCouponsForProductFamily(
    ctx context.Context,
    productFamilyId int,
    page *int,
    perPage *int,
    filterDateField *models.BasicDateField,
    filterEndDate *time.Time,
    filterEndDatetime *time.Time,
    filterStartDate *time.Time,
    filterStartDatetime *time.Time,
    filterIds []int,
    filterCodes []string,
    currencyPrices *bool,
    filterUseSiteExchangeRate *bool) (
    models.ApiResponse[[]models.CouponResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/product_families/%v/coupons.json", productFamilyId),
    )
    req.Authenticate(true)
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
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
    if filterStartDate != nil {
        req.QueryParam("filter[start_date]", filterStartDate.Format(models.DEFAULT_DATE))
    }
    if filterStartDatetime != nil {
        req.QueryParam("filter[start_datetime]", filterStartDatetime.Format(time.RFC3339))
    }
    if filterIds != nil {
        req.QueryParam("filter[ids]", filterIds)
    }
    if filterCodes != nil {
        req.QueryParam("filter[codes]", filterCodes)
    }
    if currencyPrices != nil {
        req.QueryParam("currency_prices", *currencyPrices)
    }
    if filterUseSiteExchangeRate != nil {
        req.QueryParam("filter[use_site_exchange_rate]", *filterUseSiteExchangeRate)
    }
    
    var result []models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.CouponResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ReadCouponByCode takes context, productFamilyId, code as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// You can search for a coupon via the API with the find method. By passing a code parameter, the find will attempt to locate a coupon that matches that code. If no coupon is found, a 404 is returned.
// If you have more than one product family and if the coupon you are trying to find does not belong to the default product family in your site, then you will need to specify (either in the url or as a query string param) the product family id.
func (c *CouponsController) ReadCouponByCode(
    ctx context.Context,
    productFamilyId *int,
    code *string) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/coupons/find.json")
    req.Authenticate(true)
    if productFamilyId != nil {
        req.QueryParam("product_family_id", *productFamilyId)
    }
    if code != nil {
        req.QueryParam("code", *code)
    }
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ReadCoupon takes context, productFamilyId, couponId as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// You can retrieve the Coupon via the API with the Show method. You must identify the Coupon in this call by the ID parameter that Chargify assigns.
// If instead you would like to find a Coupon using a Coupon code, see the Coupon Find method.
// When fetching a coupon, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response.
// If the coupon is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.
func (c *CouponsController) ReadCoupon(
    ctx context.Context,
    productFamilyId int,
    couponId int) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/product_families/%v/coupons/%v.json", productFamilyId, couponId),
    )
    req.Authenticate(true)
    
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// UpdateCoupon takes context, productFamilyId, couponId, body as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// ## Update Coupon
// You can update a Coupon via the API with a PUT request to the resource endpoint.
// You can restrict a coupon to only apply to specific products / components by optionally passing in hashes of `restricted_products` and/or `restricted_components` in the format:
// `{ "<product/component_id>": boolean_value }`
func (c *CouponsController) UpdateCoupon(
    ctx context.Context,
    productFamilyId int,
    couponId int,
    body *models.CreateOrUpdateCoupon) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/product_families/%v/coupons/%v.json", productFamilyId, couponId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ArchiveCoupon takes context, productFamilyId, couponId as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// You can archive a Coupon via the API with the archive method.
// Archiving makes that Coupon unavailable for future use, but allows it to remain attached and functional on existing Subscriptions that are using it.
// The `archived_at` date and time will be assigned.
func (c *CouponsController) ArchiveCoupon(
    ctx context.Context,
    productFamilyId int,
    couponId int) (
    models.ApiResponse[models.CouponResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/product_families/%v/coupons/%v.json", productFamilyId, couponId),
    )
    req.Authenticate(true)
    
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ListCoupons takes context, page, perPage, dateField, startDate, endDate, startDatetime, endDatetime, filterIds, filterCodes, currencyPrices, filterEndDate, filterEndDatetime, filterStartDate, filterStartDatetime, filterDateField, filterUseSiteExchangeRate as parameters and
// returns an models.ApiResponse with []models.CouponResponse data and
// an error if there was an issue with the request or response.
// You can retrieve a list of coupons.
// If the coupon is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.
func (c *CouponsController) ListCoupons(
    ctx context.Context,
    page *int,
    perPage *int,
    dateField *models.BasicDateField,
    startDate *time.Time,
    endDate *time.Time,
    startDatetime *time.Time,
    endDatetime *time.Time,
    filterIds []int,
    filterCodes []string,
    currencyPrices *bool,
    filterEndDate *time.Time,
    filterEndDatetime *time.Time,
    filterStartDate *time.Time,
    filterStartDatetime *time.Time,
    filterDateField *models.BasicDateField,
    filterUseSiteExchangeRate *bool) (
    models.ApiResponse[[]models.CouponResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/coupons.json")
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
        req.QueryParam("start_date", startDate.Format(models.DEFAULT_DATE))
    }
    if endDate != nil {
        req.QueryParam("end_date", endDate.Format(models.DEFAULT_DATE))
    }
    if startDatetime != nil {
        req.QueryParam("start_datetime", startDatetime.Format(time.RFC3339))
    }
    if endDatetime != nil {
        req.QueryParam("end_datetime", endDatetime.Format(time.RFC3339))
    }
    if filterIds != nil {
        req.QueryParam("filter[ids]", filterIds)
    }
    if filterCodes != nil {
        req.QueryParam("filter[codes]", filterCodes)
    }
    if currencyPrices != nil {
        req.QueryParam("currency_prices", *currencyPrices)
    }
    if filterEndDate != nil {
        req.QueryParam("filter[end_date]", filterEndDate.Format(models.DEFAULT_DATE))
    }
    if filterEndDatetime != nil {
        req.QueryParam("filter[end_datetime]", filterEndDatetime.Format(time.RFC3339))
    }
    if filterStartDate != nil {
        req.QueryParam("filter[start_date]", filterStartDate.Format(models.DEFAULT_DATE))
    }
    if filterStartDatetime != nil {
        req.QueryParam("filter[start_datetime]", filterStartDatetime.Format(time.RFC3339))
    }
    if filterDateField != nil {
        req.QueryParam("filter[date_field]", *filterDateField)
    }
    if filterUseSiteExchangeRate != nil {
        req.QueryParam("filter[use_site_exchange_rate]", *filterUseSiteExchangeRate)
    }
    var result []models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.CouponResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ReadCouponUsage takes context, productFamilyId, couponId as parameters and
// returns an models.ApiResponse with []models.CouponUsage data and
// an error if there was an issue with the request or response.
// This request will provide details about the coupon usage as an array of data hashes, one per product.
func (c *CouponsController) ReadCouponUsage(
    ctx context.Context,
    productFamilyId int,
    couponId int) (
    models.ApiResponse[[]models.CouponUsage],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/product_families/%v/coupons/%v/usage.json", productFamilyId, couponId),
    )
    req.Authenticate(true)
    
    var result []models.CouponUsage
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.CouponUsage](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ValidateCoupon takes context, code, productFamilyId as parameters and
// returns an models.ApiResponse with models.CouponResponse data and
// an error if there was an issue with the request or response.
// You can verify if a specific coupon code is valid using the `validate` method. This method is useful for validating coupon codes that are entered by a customer. If the coupon is found and is valid, the coupon will be returned with a 200 status code.
// If the coupon is invalid, the status code will be 404 and the response will say why it is invalid. If the coupon is valid, the status code will be 200 and the coupon will be returned. The following reasons for invalidity are supported:
// + Coupon not found
// + Coupon is invalid
// + Coupon expired
// If you have more than one product family and if the coupon you are validating does not belong to the first product family in your site, then you will need to specify the product family, either in the url or as a query string param. This can be done by supplying the id or the handle in the `handle:my-family` format.
// Eg.
// ```
// https://<subdomain>.chargify.com/product_families/handle:<product_family_handle>/coupons/validate.<format>?code=<coupon_code>
// ```
// Or:
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
    req.Authenticate(true)
    req.QueryParam("code", code)
    if productFamilyId != nil {
        req.QueryParam("product_family_id", *productFamilyId)
    }
    var result models.CouponResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewSingleStringErrorResponse(404, "Not Found")
    }
    return models.NewApiResponse(result, resp), err
}

// UpdateCouponCurrencyPrices takes context, couponId, body as parameters and
// returns an models.ApiResponse with models.CouponCurrencyResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to create and/or update currency prices for an existing coupon. Multiple prices can be created or updated in a single request but each of the currencies must be defined on the site level already and the coupon must be an amount-based coupon, not percentage.
// Currency pricing for coupons must mirror the setup of the primary coupon pricing - if the primary coupon is percentage based, you will not be able to define pricing in non-primary currencies.
func (c *CouponsController) UpdateCouponCurrencyPrices(
    ctx context.Context,
    couponId int,
    body *models.CouponCurrencyRequest) (
    models.ApiResponse[models.CouponCurrencyResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/coupons/%v/currency_prices.json", couponId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.CouponCurrencyResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponCurrencyResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// CreateCouponSubcodes takes context, couponId, body as parameters and
// returns an models.ApiResponse with models.CouponSubcodesResponse data and
// an error if there was an issue with the request or response.
// ## Coupon Subcodes Intro
// Coupon Subcodes allow you to create a set of unique codes that allow you to expand the use of one coupon.
// For example:
// Master Coupon Code:
// + SPRING2020
// Coupon Subcodes:
// + SPRING90210
// + DP80302
// + SPRINGBALTIMORE
// Coupon subcodes can be administered in the Admin Interface or via the API.
// When creating a coupon subcode, you must specify a coupon to attach it to using the coupon_id. Valid coupon subcodes are all capital letters, contain only letters and numbers, and do not have any spaces. Lowercase letters will be capitalized before the subcode is created.
// ## Coupon Subcodes Documentation
// Full documentation on how to create coupon subcodes in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407755909531#coupon-codes).
// Additionally, for documentation on how to apply a coupon to a Subscription within the Chargify UI, please see our documentation [here](https://chargify.zendesk.com/hc/en-us/articles/4407884887835#coupon).
// ## Create Coupon Subcode
// This request allows you to create specific subcodes underneath an existing coupon code.
// *Note*: If you are using any of the allowed special characters ("%", "@", "+", "-", "_", and "."), you must encode them for use in the URL.
// % to %25
// @ to %40
// + to %2B
// - to %2D
// _ to %5F
// . to %2E
// So, if the coupon subcode is `20%OFF`, the URL to delete this coupon subcode would be: `https://<subdomain>.chargify.com/coupons/567/codes/20%25OFF.<format>`
func (c *CouponsController) CreateCouponSubcodes(
    ctx context.Context,
    couponId int,
    body *models.CouponSubcodes) (
    models.ApiResponse[models.CouponSubcodesResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/coupons/%v/codes.json", couponId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.CouponSubcodesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponSubcodesResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ListCouponSubcodes takes context, couponId, page, perPage as parameters and
// returns an models.ApiResponse with models.CouponSubcodes data and
// an error if there was an issue with the request or response.
// This request allows you to request the subcodes that are attached to a coupon.
func (c *CouponsController) ListCouponSubcodes(
    ctx context.Context,
    couponId int,
    page *int,
    perPage *int) (
    models.ApiResponse[models.CouponSubcodes],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/coupons/%v/codes.json", couponId),
    )
    req.Authenticate(true)
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    
    var result models.CouponSubcodes
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponSubcodes](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// UpdateCouponSubcodes takes context, couponId, body as parameters and
// returns an models.ApiResponse with models.CouponSubcodesResponse data and
// an error if there was an issue with the request or response.
// You can update the subcodes for the given Coupon via the API with a PUT request to the resource endpoint.
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
    req := c.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/coupons/%v/codes.json", couponId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.CouponSubcodesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CouponSubcodesResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// DeleteCouponSubcode takes context, couponId, subcode as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
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
// Or if the coupon subcode is 20%OFF, the URL to delete this coupon subcode would be: @https://<subdomain>.chargify.com/coupons/567/codes/20%25OFF.<format>
func (c *CouponsController) DeleteCouponSubcode(
    ctx context.Context,
    couponId int,
    subcode string) (
    *http.Response,
    error) {
    req := c.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/coupons/%v/codes/%v.json", couponId, subcode),
    )
    req.Authenticate(true)
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    if context.Response.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    return context.Response, err
}
