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
    "time"
)

// SubscriptionsController represents a controller struct.
type SubscriptionsController struct {
    baseController
}

// NewSubscriptionsController creates a new instance of SubscriptionsController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionsController.
func NewSubscriptionsController(baseController baseController) *SubscriptionsController {
    subscriptionsController := SubscriptionsController{baseController: baseController}
    return &subscriptionsController
}

// CreateSubscription takes context, body as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Creates a Subscription for a customer and product
// Specify the product with `product_id` or `product_handle`. To set a specific product pricepPoint, use `product_price_point_handle` or `product_price_point_id`.
// Identify an existing customer with `customer_id` or `customer_reference`. Optionally, include an existing payment profile using `payment_profile_id`. To create a new customer, pass customer_attributes. 
// Select an option from the **Request Examples** drop-down on the right side of the portal to see examples of common scenarios for creating subscriptions. 
// Payment information may be required to create a subscription, depending on the options for the Product being subscribed. See [product options](https://docs.maxio.com/hc/en-us/articles/24261076617869-Edit-Products) for more information. See the [Payments Profile]($e/Payment%20Profiles/createPaymentProfile) endpoint for details on payment parameters.
// Do not use real card information for testing. See the Sites articles that cover [testing your site setup](https://docs.maxio.com/hc/en-us/articles/24250712113165-Testing-Overview#testing-overview-0-0) for more details on testing in your sandbox.
// Note that collecting and sending raw card details in production requires [PCI compliance](https://docs.maxio.com/hc/en-us/articles/24183956938381-PCI-Compliance#pci-compliance-0-0) on your end. If your business is not PCI compliant, use [Chargify.js](https://docs.maxio.com/hc/en-us/articles/38163190843789-Chargify-js-Overview#chargify-js-overview-0-0) to collect credit card or bank account information.
// See the [Subscription Signups](page:introduction/basic-concepts/subscription-signup) article for more information on working with subscriptions in Advanced Billing.
func (s *SubscriptionsController) CreateSubscription(
    ctx context.Context,
    body *models.CreateSubscriptionRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscriptions.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSubscriptionsInput represents the input of the ListSubscriptions endpoint.
type ListSubscriptionsInput struct {
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page                *int                             
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage             *int                             
    // The current state of the subscription
    State               *models.SubscriptionStateFilter  
    // The product id of the subscription. (Note that the product handle cannot be used.)
    Product             *int                             
    // The ID of the product price point. If supplied, product is required
    ProductPricePointId *int                             
    // The numeric id of the coupon currently applied to the subscription. (This can be found in the URL when editing a coupon. Note that the coupon code cannot be used.)
    Coupon              *int                             
    // The coupon code currently applied to the subscription
    CouponCode          *string                          
    // The type of filter you'd like to apply to your search.  Allowed Values: , current_period_ends_at, current_period_starts_at, created_at, activated_at, canceled_at, expires_at, trial_started_at, trial_ended_at, updated_at
    DateField           *models.SubscriptionDateField    
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. Use in query `start_date=2022-07-01`.
    StartDate           *time.Time                       
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. Use in query `end_date=2022-08-01`.
    EndDate             *time.Time                       
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. Use in query `start_datetime=2022-07-01 09:00:05`.
    StartDatetime       *time.Time                       
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. Use in query `end_datetime=2022-08-01 10:00:05`.
    EndDatetime         *time.Time                       
    // The value of the metadata field specified in the parameter. Use in query `metadata[my-field]=value&metadata[other-field]=another_value`.
    Metadata            map[string]string                
    // Controls the order in which results are returned.
    // Use in query `direction=asc`.
    Direction           *models.SortingDirection         
    // The attribute by which to sort
    Sort                *models.SubscriptionSort         
    // Allows including additional data in the response. Use in query: `include[]=self_service_page_token`.
    Include             []models.SubscriptionListInclude 
}

// ListSubscriptions takes context, page, perPage, state, product, productPricePointId, coupon, couponCode, dateField, startDate, endDate, startDatetime, endDatetime, metadata, direction, sort, include as parameters and
// returns an models.ApiResponse with []models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// This method will return an array of subscriptions from a Site. Pay close attention to query string filters and pagination in order to control responses from the server.
// ## Search for a subscription
// Use the query strings below to search for a subscription using the criteria available. The return value will be an array.
// ## Self-Service Page token
// Self-Service Page token for the subscriptions is not returned by default. If this information is desired, the include[]=self_service_page_token parameter must be provided with the request.
func (s *SubscriptionsController) ListSubscriptions(
    ctx context.Context,
    input ListSubscriptionsInput) (
    models.ApiResponse[[]models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/subscriptions.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParamWithArraySerializationOption("page", *input.Page, https.UnIndexed)
    }
    if input.PerPage != nil {
        req.QueryParamWithArraySerializationOption("per_page", *input.PerPage, https.UnIndexed)
    }
    if input.State != nil {
        req.QueryParamWithArraySerializationOption("state", *input.State, https.UnIndexed)
    }
    if input.Product != nil {
        req.QueryParamWithArraySerializationOption("product", *input.Product, https.UnIndexed)
    }
    if input.ProductPricePointId != nil {
        req.QueryParamWithArraySerializationOption("product_price_point_id", *input.ProductPricePointId, https.UnIndexed)
    }
    if input.Coupon != nil {
        req.QueryParamWithArraySerializationOption("coupon", *input.Coupon, https.UnIndexed)
    }
    if input.CouponCode != nil {
        req.QueryParamWithArraySerializationOption("coupon_code", *input.CouponCode, https.UnIndexed)
    }
    if input.DateField != nil {
        req.QueryParamWithArraySerializationOption("date_field", *input.DateField, https.UnIndexed)
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
    if input.Metadata != nil {
        req.QueryParamWithArraySerializationOption("metadata", input.Metadata, https.UnIndexed)
    }
    if input.Direction != nil {
        req.QueryParamWithArraySerializationOption("direction", *input.Direction, https.UnIndexed)
    }
    if input.Sort != nil {
        req.QueryParamWithArraySerializationOption("sort", *input.Sort, https.UnIndexed)
    }
    if input.Include != nil {
        req.QueryParamWithArraySerializationOption("include", input.Include, https.UnIndexed)
    }
    var result []models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSubscription takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Updates one or more attributes of a subscription.
// ## Update Subscription Payment Method
// Change the card that your subscriber uses for their subscription. You can also use this method to change the expiration date of the card **if your gateway allows**.
// Do not use real card information for testing. See the Sites articles that cover [testing your site setup](https://docs.maxio.com/hc/en-us/articles/24250712113165-Testing-Overview#testing-overview-0-0) for more details on testing in your sandbox.
// Note that collecting and sending raw card details in production requires [PCI compliance](https://docs.maxio.com/hc/en-us/articles/24183956938381-PCI-Compliance#pci-compliance-0-0) on your end. If your business is not PCI compliant, use [Chargify.js](https://docs.maxio.com/hc/en-us/articles/38163190843789-Chargify-js-Overview#chargify-js-overview-0-0) to collect credit card or bank account information.
// > Note: Partial card updates for **Authorize.Net** are not allowed via this endpoint. The existing Payment Profile must be directly updated instead.
// ## Update Product
// You also use this method to change the subscription to a different product by setting a new value for product_handle. A product change can be done in two different ways, **product change** or **delayed product change**.
// ### Product Change
// You can change a subscription's product. The new payment amount is calculated and charged at the normal start of the next period. If you require complex product changes or prorated upgrades and downgrades instead, please see the documentation on [Migrating Subscription Products](https://docs.maxio.com/hc/en-us/articles/24252069837581-Product-Changes-and-Migrations#product-changes-and-migrations-0-0).
// To perform a product change, set either the `product_handle` or `product_id` attribute to that of a different product from the same site as the subscription. You can also change the price point by passing in either `product_price_point_id` or `product_price_point_handle` - otherwise the new product's default price point is used.
// ### Delayed Product Change
// This method also changes the product and/or price point, and the new payment amount is calculated and charged at the normal start of the next period.
// This method schedules the product change to happen automatically at the subscription’s next renewal date. To perform a delayed product change, set the `product_handle` attribute as you would in a regular product change, but also set the `product_change_delayed` attribute to `true`. No proration applies in this case.
// You can also perform a delayed change to the price point by passing in either `product_price_point_id` or `product_price_point_handle`
// > **Note:** To cancel a delayed product change, set `next_product_id` to an empty string.
// ## Billing Date Changes
// You can update dates for a subscrption. 
// ### Regular Billing Date Changes
// Send the `next_billing_at` to set the next billing date for the subscription. After that date passes and the subscription is processed, the following billing date will be set according to the subscription's product period.
// > Note: If you pass an invalid date, the correct date is automatically set to he correct date. For example, if February 30 is passed, the next billing would be set to March 2nd in a non-leap year.
// The server response will not return data under the key/value pair of `next_billing_at`. View the key/value pair of `current_period_ends_at` to verify that the `next_billing_at` date has been changed successfully.
// ### Calendar Billing  and Snap Day Changes
// For a subscription using Calendar Billing, setting the next billing date is a bit different. Send the `snap_day` attribute to change the calendar billing date for **a subscription using a product eligible for calendar billing**.
// > Note: If you change the product associated with a subscription that contains a `snap_day` and immediately `READ/GET` the subscription data, it will still contain original `snap_day`. The `snap_day`will will reset to 'null on the next billing cycle. This is because  a product change is instantanous and only affects the product associated with a subscription.
func (s *SubscriptionsController) UpdateSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.UpdateSubscriptionRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/subscriptions/%v.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadSubscription takes context, subscriptionId, include as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to find subscription details.
// ## Self-Service Page token
// Self-Service Page token for the subscription is not returned by default. If this information is desired, the include[]=self_service_page_token parameter must be provided with the request.
func (s *SubscriptionsController) ReadSubscription(
    ctx context.Context,
    subscriptionId int,
    include []models.SubscriptionInclude) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/subscriptions/%v.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    if include != nil {
        req.QueryParamWithArraySerializationOption("include", include, https.UnIndexed)
    }
    
    var result models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// OverrideSubscription takes context, subscriptionId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This API endpoint allows you to set certain subscription fields that are usually managed for you automatically. Some of the fields can be set via the normal Subscriptions Update API, but others can only be set using this endpoint.
// This endpoint is provided for cases where you need to “align” Advanced Billing data with data that happened in your system, perhaps before you started using Advanced Billing. For example, you may choose to import your historical subscription data, and would like the activation and cancellation dates in Advanced Billing to match your existing historical dates. Advanced Billing does not backfill historical events (i.e. from the Events API), but some static data can be changed via this API.
// Why are some fields only settable from this endpoint, and not the normal subscription create and update endpoints? Because we want users of this endpoint to be aware that these fields are usually managed by Advanced Billing, and using this API means **you are stepping out on your own.**
// Changing these fields will not affect any other attributes. For example, adding an expiration date will not affect the next assessment date on the subscription.
// If you regularly need to override the current_period_starts_at for new subscriptions, this can also be accomplished by setting both `previous_billing_at` and `next_billing_at` at subscription creation. See the documentation on [Importing Subscriptions](./b3A6MTQxMDgzODg-create-subscription#subscriptions-import) for more information.
// ## Limitations
// When passing `current_period_starts_at` some validations are made:
// 1. The subscription needs to be unbilled (no statements or invoices).
// 2. The value passed must be a valid date/time. We recommend using the iso 8601 format.
// 3. The value passed must be before the current date/time.
// If unpermitted parameters are sent, a 400 HTTP response is sent along with a string giving the reason for the problem.
func (s *SubscriptionsController) OverrideSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.OverrideSubscriptionRequest) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "PUT", "/subscriptions/%v/override.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSingleErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// FindSubscription takes context, reference as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to find a subscription by its reference.
func (s *SubscriptionsController) FindSubscription(
    ctx context.Context,
    reference *string) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/subscriptions/lookup.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
    })
    if reference != nil {
        req.QueryParam("reference", *reference)
    }
    var result models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// PurgeSubscription takes context, subscriptionId, ack, cascade as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// For sites in test mode, you may purge individual subscriptions.
// Provide the subscription ID in the url.  To confirm, supply the customer ID in the query string `ack` parameter. You may also delete the customer record and/or payment profiles by passing `cascade` parameters. For example, to delete just the customer record, the query params would be: `?ack={customer_id}&cascade[]=customer`
// If you need to remove subscriptions from a live site, contact support to discuss your use case.
// ### Delete customer and payment profile
// The query params will be: `?ack={customer_id}&cascade[]=customer&cascade[]=payment_profile`
func (s *SubscriptionsController) PurgeSubscription(
    ctx context.Context,
    subscriptionId int,
    ack int,
    cascade []models.SubscriptionPurgeType) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscriptions/%v/purge.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSubscriptionResponseError},
    })
    req.QueryParamWithArraySerializationOption("ack", ack, https.UnIndexed)
    if cascade != nil {
        req.QueryParamWithArraySerializationOption("cascade", cascade, https.UnIndexed)
    }
    
    var result models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdatePrepaidSubscriptionConfiguration takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.PrepaidConfigurationResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to update a subscription's prepaid configuration.
func (s *SubscriptionsController) UpdatePrepaidSubscriptionConfiguration(
    ctx context.Context,
    subscriptionId int,
    body *models.UpsertPrepaidConfigurationRequest) (
    models.ApiResponse[models.PrepaidConfigurationResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      "/subscriptions/%v/prepaid_configurations.json",
    )
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'."},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.PrepaidConfigurationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PrepaidConfigurationResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// PreviewSubscription takes context, body as parameters and
// returns an models.ApiResponse with models.SubscriptionPreviewResponse data and
// an error if there was an issue with the request or response.
// The Chargify API allows you to preview a subscription by POSTing the same JSON or XML as for a subscription creation.
// The "Next Billing" amount and "Next Billing" date are represented in each Subscriber's Summary.
// A subscription will not be created by utilizing this endpoint; it is meant to serve as a prediction.
// For more information, see our documentation [here](https://maxio.zendesk.com/hc/en-us/articles/24252493695757-Subscriber-Interface-Overview).
// ## Taxable Subscriptions
// This endpoint will preview taxes applicable to a purchase. In order for taxes to be previewed, the following conditions must be met:
// + Taxes must be configured on the subscription
// + The preview must be for the purchase of a taxable product or component, or combination of the two.
// + The subscription payload must contain a full billing or shipping address in order to calculate tax
// For more information about creating taxable previews, see our documentation guide on how to create [taxable subscriptions.](https://maxio.zendesk.com/hc/en-us/sections/24287012349325-Taxes)
// You do **not** need to include a card number to generate tax information when you are previewing a subscription. However, when you actually want to create the subscription, you must include the credit card information if you want the billing address to be stored in Advanced Billing. The billing address and the credit card information are stored together within the payment profile object. Also, you may not send a billing address to Advanced Billing without payment profile information, as the address is stored on the card.
// You can pass shipping and billing addresses and still decide not to calculate taxes. To do that, pass `skip_billing_manifest_taxes: true` attribute.
// ## Non-taxable Subscriptions
// If you'd like to calculate subscriptions that do not include tax you may leave off the billing information.
func (s *SubscriptionsController) PreviewSubscription(
    ctx context.Context,
    body *models.CreateSubscriptionRequest) (
    models.ApiResponse[models.SubscriptionPreviewResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscriptions/preview.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.SubscriptionPreviewResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionPreviewResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ApplyCouponsToSubscription takes context, subscriptionId, code, body as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// An existing subscription can accommodate multiple discounts/coupon codes. This is only applicable if each coupon is stackable. For more information on stackable coupons, we recommend reviewing our [coupon documentation.](https://maxio.zendesk.com/hc/en-us/articles/24261259337101-Coupons-and-Subscriptions#stackability-rules)
// ## Query Parameters vs Request Body Parameters
// Passing in a coupon code as a query parameter will add the code to the subscription, completely replacing all existing coupon codes on the subscription.
// For this reason, using this query parameter on this endpoint has been deprecated in favor of using the request body parameters as described below. When passing in request body parameters, the list of coupon codes will simply be added to any existing list of codes on the subscription.
func (s *SubscriptionsController) ApplyCouponsToSubscription(
    ctx context.Context,
    subscriptionId int,
    code *string,
    body *models.AddCouponsRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscriptions/%v/add_coupon.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSubscriptionAddCouponError},
    })
    req.Header("Content-Type", "application/json")
    if code != nil {
        req.QueryParam("code", *code)
    }
    if body != nil {
        req.Json(body)
    }
    
    var result models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// RemoveCouponFromSubscription takes context, subscriptionId, couponCode as parameters and
// returns an models.ApiResponse with string data and
// an error if there was an issue with the request or response.
// Use this endpoint to remove a coupon from an existing subscription.
// For more information on the expected behaviour of removing a coupon from a subscription, See our documentation [here.](https://maxio.zendesk.com/hc/en-us/articles/24261259337101-Coupons-and-Subscriptions#removing-a-coupon)
func (s *SubscriptionsController) RemoveCouponFromSubscription(
    ctx context.Context,
    subscriptionId int,
    couponCode *string) (
    models.ApiResponse[string],
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/subscriptions/%v/remove_coupon.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSubscriptionRemoveCouponErrors},
    })
    if couponCode != nil {
        req.QueryParam("coupon_code", *couponCode)
    }
    
    str, resp, err := req.CallAsText()
    var result string = str

    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    return models.NewApiResponse(result, resp), err
}

// ActivateSubscription takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Advanced Billing offers the ability to activate awaiting signup and trialing subscriptions. This feature is only available on the Relationship Invoicing architecture. Subscriptions in a group may not be activated immediately.
// For details on how the activation works, and how to activate subscriptions through the application, see [activation](#).
// The `revert_on_failure` parameter controls the behavior upon activation failure.
// - If set to `true` and something goes wrong i.e. payment fails, then Advanced Billing will not change the subscription's state. The subscription’s billing period will also remain the same.
// - If set to `false` and something goes wrong i.e. payment fails, then Advanced Billing will continue through with the activation and enter an end of life state. For trialing subscriptions, that will either be trial ended (if the trial is no obligation), past due (if the trial has an obligation), or canceled (if the site has no dunning strategy, or has a strategy that says to cancel immediately). For awaiting signup subscriptions, that will always be canceled.
// The default activation failure behavior can be configured per activation attempt, or you may set a default value under Config > Settings > Subscription Activation Settings.
// ## Activation Scenarios
// ### Activate Awaiting Signup subscription
// - Given you have a product without trial
// - Given you have a site without dunning strategy
// ```mermaid
// flowchart LR
// AS[Awaiting Signup] --> A{Activate}
// A -->|Success| Active
// A -->|Failure| ROF{revert_on_failure}
// ROF -->|true| AS
// ROF -->|false| Canceled
// ```
// - Given you have a product with trial
// - Given you have a site with dunning strategy
// ```mermaid
// flowchart LR
// AS[Awaiting Signup] --> A{Activate}
// A -->|Success| Trialing
// A -->|Failure| ROF{revert_on_failure}
// ROF -->|true| AS
// ROF -->|false| PD[Past Due]
// ```
// ### Activate Trialing subscription
// You can read more about the behavior of trialing subscriptions [here](https://maxio.zendesk.com/hc/en-us/articles/24252155721869-Trialing-Subscriptions).
// When the `revert_on_failure` parameter is set to `true`, the subscription's state will remain as Trialing, we will void the invoice from activation and return any prepayments and credits applied to the invoice back to the subscription.
func (s *SubscriptionsController) ActivateSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.ActivateSubscriptionRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/subscriptions/%v/activate.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorArrayMapResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
