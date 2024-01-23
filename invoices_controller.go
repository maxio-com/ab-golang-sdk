package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
)

// InvoicesController represents a controller struct.
type InvoicesController struct {
    baseController
}

// NewInvoicesController creates a new instance of InvoicesController.
// It takes a baseController as a parameter and returns a pointer to the InvoicesController.
func NewInvoicesController(baseController baseController) *InvoicesController {
    invoicesController := InvoicesController{baseController: baseController}
    return &invoicesController
}

// RefundInvoice takes context, uid, body as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// Refund an invoice, segment, or consolidated invoice.
// ## Partial Refund for Consolidated Invoice
// A refund less than the total of a consolidated invoice will be split across its segments.
// A $50.00 refund on a $100.00 consolidated invoice with one $60.00 and one $40.00 segment, the refunded amount will be applied as 50% of each ($30.00 and $20.00 respectively).
func (i *InvoicesController) RefundInvoice(
    ctx context.Context,
    uid string,
    body *models.RefundInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/invoices/%v/refunds.json", uid),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.Invoice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Invoice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ListInvoices takes context, startDate, endDate, status, subscriptionId, subscriptionGroupUid, page, perPage, direction, lineItems, discounts, taxes, credits, payments, customFields, refunds, dateField, startDatetime, endDatetime, customerIds, number, productIds, sort as parameters and
// returns an models.ApiResponse with models.ListInvoicesResponse data and
// an error if there was an issue with the request or response.
// By default, invoices returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, `custom_fields`, or `refunds`. To include breakdowns, pass the specific field as a key in the query with a value set to `true`.
func (i *InvoicesController) ListInvoices(
    ctx context.Context,
    startDate *string,
    endDate *string,
    status *models.InvoiceStatus,
    subscriptionId *int,
    subscriptionGroupUid *string,
    page *int,
    perPage *int,
    direction *models.Direction,
    lineItems *bool,
    discounts *bool,
    taxes *bool,
    credits *bool,
    payments *bool,
    customFields *bool,
    refunds *bool,
    dateField *models.InvoiceDateField,
    startDatetime *string,
    endDatetime *string,
    customerIds []int,
    number []string,
    productIds []int,
    sort *models.InvoiceSortField) (
    models.ApiResponse[models.ListInvoicesResponse],
    error) {
    req := i.prepareRequest(ctx, "GET", "/invoices.json")
    req.Authenticate(true)
    if startDate != nil {
        req.QueryParam("start_date", *startDate)
    }
    if endDate != nil {
        req.QueryParam("end_date", *endDate)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    if subscriptionId != nil {
        req.QueryParam("subscription_id", *subscriptionId)
    }
    if subscriptionGroupUid != nil {
        req.QueryParam("subscription_group_uid", *subscriptionGroupUid)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if direction != nil {
        req.QueryParam("direction", *direction)
    }
    if lineItems != nil {
        req.QueryParam("line_items", *lineItems)
    }
    if discounts != nil {
        req.QueryParam("discounts", *discounts)
    }
    if taxes != nil {
        req.QueryParam("taxes", *taxes)
    }
    if credits != nil {
        req.QueryParam("credits", *credits)
    }
    if payments != nil {
        req.QueryParam("payments", *payments)
    }
    if customFields != nil {
        req.QueryParam("custom_fields", *customFields)
    }
    if refunds != nil {
        req.QueryParam("refunds", *refunds)
    }
    if dateField != nil {
        req.QueryParam("date_field", *dateField)
    }
    if startDatetime != nil {
        req.QueryParam("start_datetime", *startDatetime)
    }
    if endDatetime != nil {
        req.QueryParam("end_datetime", *endDatetime)
    }
    if customerIds != nil {
        req.QueryParam("customer_ids", customerIds)
    }
    if number != nil {
        req.QueryParam("number", number)
    }
    if productIds != nil {
        req.QueryParam("product_ids", productIds)
    }
    if sort != nil {
        req.QueryParam("sort", *sort)
    }
    var result models.ListInvoicesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListInvoicesResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ReadInvoice takes context, uid as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// Use this endpoint to retrieve the details for an invoice.
func (i *InvoicesController) ReadInvoice(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.Invoice],
    error) {
    req := i.prepareRequest(ctx, "GET", fmt.Sprintf("/invoices/%v.json", uid))
    req.Authenticate(true)
    
    var result models.Invoice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Invoice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ListInvoiceEvents takes context, sinceDate, sinceId, page, perPage, invoiceUid, withChangeInvoiceStatus, eventTypes as parameters and
// returns an models.ApiResponse with models.ListInvoiceEventsResponse data and
// an error if there was an issue with the request or response.
// This endpoint returns a list of invoice events. Each event contains event "data" (such as an applied payment) as well as a snapshot of the `invoice` at the time of event completion.
// Exposed event types are:
// + issue_invoice
// + apply_credit_note
// + apply_payment
// + refund_invoice
// + void_invoice
// + void_remainder
// + backport_invoice
// + change_invoice_status
// + change_invoice_collection_method
// + remove_payment
// + failed_payment
// + apply_debit_note
// + create_debit_note
// + change_chargeback_status
// Invoice events are returned in ascending order.
// If both a `since_date` and `since_id` are provided in request parameters, the `since_date` will be used.
// Note - invoice events that occurred prior to 09/05/2018 __will not__ contain an `invoice` snapshot.
func (i *InvoicesController) ListInvoiceEvents(
    ctx context.Context,
    sinceDate *string,
    sinceId *int,
    page *int,
    perPage *int,
    invoiceUid *string,
    withChangeInvoiceStatus *string,
    eventTypes []models.InvoiceEventType) (
    models.ApiResponse[models.ListInvoiceEventsResponse],
    error) {
    req := i.prepareRequest(ctx, "GET", "/invoices/events.json")
    req.Authenticate(true)
    if sinceDate != nil {
        req.QueryParam("since_date", *sinceDate)
    }
    if sinceId != nil {
        req.QueryParam("since_id", *sinceId)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if invoiceUid != nil {
        req.QueryParam("invoice_uid", *invoiceUid)
    }
    if withChangeInvoiceStatus != nil {
        req.QueryParam("with_change_invoice_status", *withChangeInvoiceStatus)
    }
    if eventTypes != nil {
        req.QueryParam("event_types", eventTypes)
    }
    var result models.ListInvoiceEventsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListInvoiceEventsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// RecordPaymentForInvoice takes context, uid, body as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// This API call should be used when you want to record a payment of a given type against a specific invoice. If you would like to apply a payment across multiple invoices, you can use the Bulk Payment endpoint.
// ## Create a Payment from the existing payment profile
// In order to apply a payment to an invoice using an existing payment profile, specify `type` as `payment`, the amount less than the invoice total, and the customer's `payment_profile_id`. The ID of a payment profile might be retrieved via the Payment Profiles API endpoint.
// ```
// {
// "type": "payment",
// "payment": {
// "amount": 10.00,
// "payment_profile_id": 123
// }
// }
// ```
// ## Create a Payment from the Subscription's Prepayment Account
// In order apply a prepayment to an invoice, specify the `type` as `prepayment`, and also the `amount`.
// ```
// {
// "type": "prepayment",
// "payment": {
// "amount": 10.00
// }
// }
// ```
// Note that the `amount` must be less than or equal to the Subscription's Prepayment account balance.
// ## Create a Payment from the Subscription's Service Credit Account
// In order to apply a service credit to an invoice, specify the `type` as `service_credit`, and also the `amount`:
// ```
// {
// "type": "service_credit",
// "payment": {
// "amount": 10.00
// }
// }
// ```
// Note that Chargify will attempt to fully pay the invoice's `due_amount` from the Subscription's Service Credit account. At this time, partial payments from a Service Credit Account are only allowed for consolidated invoices (subscription groups). Therefore, for normal invoices the Service Credit account balance must be greater than or equal to the invoice's `due_amount`.
func (i *InvoicesController) RecordPaymentForInvoice(
    ctx context.Context,
    uid string,
    body *models.CreateInvoicePaymentRequest) (
    models.ApiResponse[models.Invoice],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/invoices/%v/payments.json", uid),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.Invoice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Invoice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// RecordExternalPaymentForInvoices takes context, body as parameters and
// returns an models.ApiResponse with models.MultiInvoicePaymentResponse data and
// an error if there was an issue with the request or response.
// This API call should be used when you want to record an external payment against multiple invoices.
// In order apply a payment to multiple invoices, at minimum, specify the `amount` and `applications` (i.e., `invoice_uid` and `amount`) details.
// ```
// {
// "payment": {
// "memo": "to pay the bills",
// "details": "check number 8675309",
// "method": "check",
// "amount": "250.00",
// "applications": [
// {
// "invoice_uid": "inv_8gk5bwkct3gqt",
// "amount": "100.00"
// },
// {
// "invoice_uid": "inv_7bc6bwkct3lyt",
// "amount": "150.00"
// }
// ]
// }
// }
// ```
// Note that the invoice payment amounts must be greater than 0. Total amount must be greater or equal to invoices payment amount sum.
func (i *InvoicesController) RecordExternalPaymentForInvoices(
    ctx context.Context,
    body *models.CreateMultiInvoicePaymentRequest) (
    models.ApiResponse[models.MultiInvoicePaymentResponse],
    error) {
    req := i.prepareRequest(ctx, "POST", "/invoices/payments.json")
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    var result models.MultiInvoicePaymentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.MultiInvoicePaymentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity")
    }
    return models.NewApiResponse(result, resp), err
}

// ListCreditNotes takes context, subscriptionId, page, perPage, lineItems, discounts, taxes, refunds, applications as parameters and
// returns an models.ApiResponse with models.ListCreditNotesResponse data and
// an error if there was an issue with the request or response.
// Credit Notes are like inverse invoices. They reduce the amount a customer owes.
// By default, the credit notes returned by this endpoint will exclude the arrays of `line_items`, `discounts`, `taxes`, `applications`, or `refunds`. To include these arrays, pass the specific field as a key in the query with a value set to `true`.
func (i *InvoicesController) ListCreditNotes(
    ctx context.Context,
    subscriptionId *int,
    page *int,
    perPage *int,
    lineItems *bool,
    discounts *bool,
    taxes *bool,
    refunds *bool,
    applications *bool) (
    models.ApiResponse[models.ListCreditNotesResponse],
    error) {
    req := i.prepareRequest(ctx, "GET", "/credit_notes.json")
    req.Authenticate(true)
    if subscriptionId != nil {
        req.QueryParam("subscription_id", *subscriptionId)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if lineItems != nil {
        req.QueryParam("line_items", *lineItems)
    }
    if discounts != nil {
        req.QueryParam("discounts", *discounts)
    }
    if taxes != nil {
        req.QueryParam("taxes", *taxes)
    }
    if refunds != nil {
        req.QueryParam("refunds", *refunds)
    }
    if applications != nil {
        req.QueryParam("applications", *applications)
    }
    var result models.ListCreditNotesResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListCreditNotesResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ReadCreditNote takes context, uid as parameters and
// returns an models.ApiResponse with models.CreditNote data and
// an error if there was an issue with the request or response.
// Use this endpoint to retrieve the details for a credit note.
func (i *InvoicesController) ReadCreditNote(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.CreditNote],
    error) {
    req := i.prepareRequest(ctx, "GET", fmt.Sprintf("/credit_notes/%v.json", uid))
    req.Authenticate(true)
    
    var result models.CreditNote
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CreditNote](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// RecordPaymentForSubscription takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.PaymentResponse data and
// an error if there was an issue with the request or response.
// Record an external payment made against a subscription that will pay partially or in full one or more invoices.
// Payment will be applied starting with the oldest open invoice and then next oldest, and so on until the amount of the payment is fully consumed.
// Excess payment will result in the creation of a prepayment on the Invoice Account.
// Only ungrouped or primary subscriptions may be paid using the "bulk" payment request.
func (i *InvoicesController) RecordPaymentForSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.RecordPaymentRequest) (
    models.ApiResponse[models.PaymentResponse],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/payments.json", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.PaymentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PaymentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ReopenInvoice takes context, uid as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// This endpoint allows you to reopen any invoice with the "canceled" status. Invoices enter "canceled" status if they were open at the time the subscription was canceled (whether through dunning or an intentional cancellation).
// Invoices with "canceled" status are no longer considered to be due. Once reopened, they are considered due for payment. Payment may then be captured in one of the following ways:
// - Reactivating the subscription, which will capture all open invoices (See note below about automatic reopening of invoices.)
// - Recording a payment directly against the invoice
// A note about reactivations: any canceled invoices from the most recent active period are automatically opened as a part of the reactivation process. Reactivating via this endpoint prior to reactivation is only necessary when you wish to capture older invoices from previous periods during the reactivation.
// ### Reopening Consolidated Invoices
// When reopening a consolidated invoice, all of its canceled segments will also be reopened.
func (i *InvoicesController) ReopenInvoice(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.Invoice],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/invoices/%v/reopen.json", uid),
    )
    req.Authenticate(true)
    
    var result models.Invoice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Invoice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// VoidInvoice takes context, uid, body as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// This endpoint allows you to void any invoice with the "open" or "canceled" status.  It will also allow voiding of an invoice with the "pending" status if it is not a consolidated invoice.
func (i *InvoicesController) VoidInvoice(
    ctx context.Context,
    uid string,
    body *models.VoidInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error) {
    req := i.prepareRequest(ctx, "POST", fmt.Sprintf("/invoices/%v/void.json", uid))
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.Invoice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Invoice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ListInvoiceSegments takes context, invoiceUid, page, perPage, direction as parameters and
// returns an models.ApiResponse with models.ConsolidatedInvoice data and
// an error if there was an issue with the request or response.
// Invoice segments returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, or `custom_fields`.
func (i *InvoicesController) ListInvoiceSegments(
    ctx context.Context,
    invoiceUid string,
    page *int,
    perPage *int,
    direction *models.Direction) (
    models.ApiResponse[models.ConsolidatedInvoice],
    error) {
    req := i.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/invoices/%v/segments.json", invoiceUid),
    )
    req.Authenticate(true)
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if direction != nil {
        req.QueryParam("direction", *direction)
    }
    
    var result models.ConsolidatedInvoice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ConsolidatedInvoice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// CreateInvoice takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.InvoiceResponse data and
// an error if there was an issue with the request or response.
// This endpoint will allow you to create an ad hoc invoice.
// ### Basic Behavior
// You can create a basic invoice by sending an array of line items to this endpoint. Each line item, at a minimum, must include a title, a quantity and a unit price. Example:
// ```json
// {
// "invoice": {
// "line_items": [
// {
// "title": "A Product",
// "quantity": 12,
// "unit_price": "150.00"
// }
// ]
// }
// }
// ```
// ### Catalog items
// Instead of creating custom products like in above example, You can pass existing items like products, components.
// ```json
// {
// "invoice": {
// "line_items": [
// {
// "product_id": "handle:gold-product",
// "quantity": 2,
// }
// ]
// }
// }
// ```
// The price for each line item will be calculated as well as a total due amount for the invoice. Multiple line items can be sent.
// ### Line items types
// When defining line item, You can choose one of 3 types for one line item:
// #### Custom item
// Like in basic behavior example above, You can pass `title` and `unit_price` for custom item.
// #### Product id
// Product handle (with handle: prefix) or id from the scope of current subscription's site can be provided with `product_id`. By default `unit_price` is taken from product's default price point, but can be overwritten by passing `unit_price` or `product_price_point_id`. If `product_id` is used, following fields cannot be used: `title`, `component_id`.
// #### Component id
// Component handle (with handle: prefix) or id from the scope of current subscription's site can be provided with `component_id`. If `component_id` is used, following fields cannot be used: `title`, `product_id`. By default `unit_price` is taken from product's default price point, but can be overwritten by passing `unit_price` or `price_point_id`. At this moment price points are supportted only for quantity based, on/off and metered components. For prepaid and event based billing components `unit_price` is required.
// ### Coupons
// When creating ad hoc invoice, new discounts can be applied in following way:
// ```json
// {
// "invoice": {
// "line_items": [
// {
// "product_id": "handle:gold-product",
// "quantity": 1
// }
// ],
// "coupons": [
// {
// "code": "COUPONCODE",
// "percentage": 50.0
// }
// ]
// }
// }
// ```
// If You want to use existing coupon for discount creation, only `code` and optional `product_family_id` is needed
// ```json
// ...
// "coupons": [
// {
// "code": "FREESETUP",
// "product_family_id": 1
// }
// ]
// ...
// ```
// ### Coupon options
// #### Code
// Coupon `code` will be displayed on invoice discount section.
// Coupon code can only contain uppercase letters, numbers, and allowed special characters.
// Lowercase letters will be converted to uppercase. It can be used to select an existing coupon from the catalog, or as an ad hoc coupon when passed with `percentage` or `amount`.
// #### Percentage
// Coupon `percentage` can take values from 0 to 100 and up to 4 decimal places. It cannot be used with `amount`. Only for ad hoc coupons, will be ignored if `code` is used to select an existing coupon from the catalog.
// #### Amount
// Coupon `amount` takes number value. It cannot be used with `percentage`. Used only when not matching existing coupon by `code`.
// #### Description
// Optional `description` will be displayed with coupon `code`. Used only when not matching existing coupon by `code`.
// #### Product Family id
// Optional `product_family_id` handle (with handle: prefix) or id is used to match existing coupon within site, when codes are not unique.
// #### Compounding Strategy
// Optional `compounding_strategy` for percentage coupons, can take values `compound` or `full-price`.
// For amount coupons, discounts will be always calculated against the original item price, before other discounts are applied.
// `compound` strategy:
// Percentage-based discounts will be calculated against the remaining price, after prior discounts have been calculated. It is set by default.
// `full-price` strategy:
// Percentage-based discounts will always be calculated against the original item price, before other discounts are applied.
// ### Line Item Options
// #### Period Date Range
// A custom period date range can be defined for each line item with the `period_range_start` and `period_range_end` parameters. Dates must be sent in the `YYYY-MM-DD` format.
// `period_range_end` must be greater or equal `period_range_start`.
// #### Taxes
// The `taxable` parameter can be sent as `true` if taxes should be calculated for a specific line item. For this to work, the site should be configured to use and calculate taxes. Further, if the site uses Avalara for tax calculations, a `tax_code` parameter should also be sent. For existing catalog items: products/components taxes cannot be overwritten.
// #### Price Point
// Price point handle (with handle: prefix) or id from the scope of current subscription's site can be provided with `price_point_id` for components with `component_id` or `product_price_point_id` for products with `product_id` parameter. If price point is passed `unit_price` cannot be used. It can be used only with catalog items products and components.
// #### Description
// Optional `description` parameter, it will overwrite default generated description for line item.
// ### Invoice Options
// #### Issue Date
// By default, invoices will be created with a issue date set to today. `issue_date` parameter can be send to alter that. Only dates in the past can be send. `issue_date` should be send in `YYYY-MM-DD` format.
// #### Net Terms
// By default, invoices will be created with a due date matching the date of invoice creation. If a different due date is desired, the `net_terms` parameter can be sent indicating the number of days in advance the due date should be.
// #### Addresses
// The seller, shipping and billing addresses can be sent to override the site's defaults. Each address requires to send a `first_name` at a minimum in order to work. Please see below for the details on which parameters can be sent for each address object.
// #### Memo and Payment Instructions
// A custom memo can be sent with the `memo` parameter to override the site's default. Likewise, custom payment instructions can be sent with the `payment_instrucions` parameter.
// #### Status
// By default, invoices will be created with open status. Possible alternative is `draft`.
func (i *InvoicesController) CreateInvoice(
    ctx context.Context,
    subscriptionId int,
    body *models.CreateInvoiceRequest) (
    models.ApiResponse[models.InvoiceResponse],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/invoices.json", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.InvoiceResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.InvoiceResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorArrayMapResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// SendInvoice takes context, uid, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This endpoint allows for invoices to be programmatically delivered via email. This endpoint supports the delivery of both ad-hoc and automatically generated invoices. Additionally, this endpoint supports email delivery to direct recipients, carbon-copy (cc) recipients, and blind carbon-copy (bcc) recipients.
// Please note that if no recipient email addresses are specified in the request, then the subscription's default email configuration will be used. For example, if `recipient_emails` is left blank, then the invoice will be delivered to the subscription's customer email address.
// On success, a 204 no-content response will be returned. Please note that this does not indicate that email(s) have been delivered, but instead indicates that emails have been successfully queued for delivery. If _any_ invalid or malformed email address is found in the request body, the entire request will be rejected and a 422 response will be returned.
func (i *InvoicesController) SendInvoice(
    ctx context.Context,
    uid string,
    body *models.SendInvoiceRequest) (
    *http.Response,
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/invoices/%v/deliveries.json", uid),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    if context.Response.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return context.Response, err
}

// PreviewCustomerInformationChanges takes context, uid as parameters and
// returns an models.ApiResponse with models.CustomerChangesPreviewResponse data and
// an error if there was an issue with the request or response.
// Customer information may change after an invoice is issued which may lead to a mismatch between customer information that are present on an open invoice and actual customer information. This endpoint allows to preview these differences, if any.
// The endpoint doesn't accept a request body. Customer information differences are calculated on the application side.
func (i *InvoicesController) PreviewCustomerInformationChanges(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.CustomerChangesPreviewResponse],
    error) {
    req := i.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/invoices/%v/customer_information/preview.json", uid),
    )
    req.Authenticate(true)
    
    var result models.CustomerChangesPreviewResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CustomerChangesPreviewResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewErrorListResponse(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// UpdateCustomerInformation takes context, uid as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// This endpoint updates customer information on an open invoice and returns the updated invoice. If you would like to preview changes that will be applied, use the `/invoices/{uid}/customer_information/preview.json` endpoint before.
// The endpoint doesn't accept a request body. Customer information differences are calculated on the application side.
func (i *InvoicesController) UpdateCustomerInformation(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.Invoice],
    error) {
    req := i.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/invoices/%v/customer_information.json", uid),
    )
    req.Authenticate(true)
    
    var result models.Invoice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Invoice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewErrorListResponse(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// IssueInvoice takes context, uid, body as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// This endpoint allows you to issue an invoice that is in "pending" status. For example, you can issue an invoice that was created when allocating new quantity on a component and using "accrue charges" option.
// You cannot issue a pending child invoice that was created for a member subscription in a group.
// For Remittance subscriptions, the invoice will go into "open" status and payment won't be attempted. The value for `on_failed_payment` would be rejected if sent. Any prepayments or service credits that exist on subscription will be automatically applied. Additionally, if setting is on, an email will be sent for issued invoice.
// For Automatic subscriptions, prepayments and service credits will apply to the invoice and before payment is attempted. On successful payment, the invoice will go into "paid" status and email will be sent to the customer (if setting applies). When payment fails, the next event depends on the `on_failed_payment` value:
// - `leave_open_invoice` - prepayments and credits applied to invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history. This is the default option.
// - `rollback_to_pending` - prepayments and credits not applied; invoice remains in "pending" status; no email sent to the customer; payment failure recorded in the invoice history.
// - `initiate_dunning` - prepayments and credits applied to the invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history; subscription will  most likely go into "past_due" or "canceled" state (depending upon net terms and dunning settings).
func (i *InvoicesController) IssueInvoice(
    ctx context.Context,
    uid string,
    body *models.IssueInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error) {
    req := i.prepareRequest(ctx, "POST", fmt.Sprintf("/invoices/%v/issue.json", uid))
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.Invoice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Invoice](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}
