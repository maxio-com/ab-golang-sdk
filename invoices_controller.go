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
	req := i.prepareRequest(ctx, "POST", "/invoices/%v/refunds.json")
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Invoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Invoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListInvoicesInput represents the input of the ListInvoices endpoint.
type ListInvoicesInput struct {
	// The start date (format YYYY-MM-DD) with which to filter the date_field. Returns invoices with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
	StartDate *string
	// The end date (format YYYY-MM-DD) with which to filter the date_field. Returns invoices with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
	EndDate *string
	// The current status of the invoice.  Allowed Values: draft, open, paid, pending, voided
	Status *models.InvoiceStatus
	// The subscription's ID.
	SubscriptionId *int
	// The UID of the subscription group you want to fetch consolidated invoices for. This will return a paginated list of consolidated invoices for the specified group.
	SubscriptionGroupUid *string
	// The consolidation level of the invoice. Allowed Values: none, parent, child or comma-separated lists of thereof, e.g. none,parent.
	ConsolidationLevel *string
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
	// The sort direction of the returned invoices.
	Direction *models.Direction
	// Include line items data
	LineItems *bool
	// Include discounts data
	Discounts *bool
	// Include taxes data
	Taxes *bool
	// Include credits data
	Credits *bool
	// Include payments data
	Payments *bool
	// Include custom fields data
	CustomFields *bool
	// Include refunds data
	Refunds *bool
	// The type of filter you would like to apply to your search. Use in query `date_field=issue_date`.
	DateField *models.InvoiceDateField
	// The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns invoices with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. Allowed to be used only along with date_field set to created_at or updated_at.
	StartDatetime *string
	// The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns invoices with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. Allowed to be used only along with date_field set to created_at or updated_at.
	EndDatetime *string
	// Allows fetching invoices with matching customer id based on provided values. Use in query `customer_ids=1,2,3`.
	CustomerIds []int
	// Allows fetching invoices with matching invoice number based on provided values. Use in query `number=1234,1235`.
	Number []string
	// Allows fetching invoices with matching line items product ids based on provided values. Use in query `product_ids=23,34`.
	ProductIds []int
	// Allows specification of the order of the returned list. Use in query `sort=total_amount`.
	Sort *models.InvoiceSortField
}

// ListInvoices takes context, startDate, endDate, status, subscriptionId, subscriptionGroupUid, consolidationLevel, page, perPage, direction, lineItems, discounts, taxes, credits, payments, customFields, refunds, dateField, startDatetime, endDatetime, customerIds, number, productIds, sort as parameters and
// returns an models.ApiResponse with models.ListInvoicesResponse data and
// an error if there was an issue with the request or response.
// By default, invoices returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, `custom_fields`, or `refunds`. To include breakdowns, pass the specific field as a key in the query with a value set to `true`.
func (i *InvoicesController) ListInvoices(
	ctx context.Context,
	input ListInvoicesInput) (
	models.ApiResponse[models.ListInvoicesResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/invoices.json")

	req.Authenticate(NewAuth("BasicAuth"))
	if input.StartDate != nil {
		req.QueryParam("start_date", *input.StartDate)
	}
	if input.EndDate != nil {
		req.QueryParam("end_date", *input.EndDate)
	}
	if input.Status != nil {
		req.QueryParam("status", *input.Status)
	}
	if input.SubscriptionId != nil {
		req.QueryParam("subscription_id", *input.SubscriptionId)
	}
	if input.SubscriptionGroupUid != nil {
		req.QueryParam("subscription_group_uid", *input.SubscriptionGroupUid)
	}
	if input.ConsolidationLevel != nil {
		req.QueryParam("consolidation_level", *input.ConsolidationLevel)
	}
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.Direction != nil {
		req.QueryParam("direction", *input.Direction)
	}
	if input.LineItems != nil {
		req.QueryParam("line_items", *input.LineItems)
	}
	if input.Discounts != nil {
		req.QueryParam("discounts", *input.Discounts)
	}
	if input.Taxes != nil {
		req.QueryParam("taxes", *input.Taxes)
	}
	if input.Credits != nil {
		req.QueryParam("credits", *input.Credits)
	}
	if input.Payments != nil {
		req.QueryParam("payments", *input.Payments)
	}
	if input.CustomFields != nil {
		req.QueryParam("custom_fields", *input.CustomFields)
	}
	if input.Refunds != nil {
		req.QueryParam("refunds", *input.Refunds)
	}
	if input.DateField != nil {
		req.QueryParam("date_field", *input.DateField)
	}
	if input.StartDatetime != nil {
		req.QueryParam("start_datetime", *input.StartDatetime)
	}
	if input.EndDatetime != nil {
		req.QueryParam("end_datetime", *input.EndDatetime)
	}
	if input.CustomerIds != nil {
		req.QueryParam("customer_ids", input.CustomerIds)
	}
	if input.Number != nil {
		req.QueryParam("number", input.Number)
	}
	if input.ProductIds != nil {
		req.QueryParam("product_ids", input.ProductIds)
	}
	if input.Sort != nil {
		req.QueryParam("sort", *input.Sort)
	}
	var result models.ListInvoicesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListInvoicesResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ReadInvoice takes context, uid as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// Use this endpoint to retrieve the details for an invoice.
// ## PDF Invoice retrieval
// Individual PDF Invoices can be retrieved by using the "Accept" header application/pdf or appending .pdf as the format portion of the URL:
// ```curl -u <api_key>:x -H
// Accept:application/pdf -H
// https://acme.chargify.com/invoices/inv_8gd8tdhtd3hgr.pdf > output_file.pdf
// URL: `https://<subdomain>.chargify.com/invoices/<uid>.<format>`
// Method: GET
// Required parameters: `uid`
// Response: A single Invoice.
// ```
func (i *InvoicesController) ReadInvoice(
	ctx context.Context,
	uid string) (
	models.ApiResponse[models.Invoice],
	error) {
	req := i.prepareRequest(ctx, "GET", "/invoices/%v.json")
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))

	var result models.Invoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Invoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListInvoiceEventsInput represents the input of the ListInvoiceEvents endpoint.
type ListInvoiceEventsInput struct {
	// The timestamp in a format `YYYY-MM-DD T HH:MM:SS Z`, or `YYYY-MM-DD`(in this case, it returns data from the beginning of the day). of the event from which you want to start the search. All the events before the `since_date` timestamp are not returned in the response.
	SinceDate *string
	// The ID of the event from which you want to start the search(ID is not included. e.g. if ID is set to 2, then all events with ID 3 and more will be shown) This parameter is not used if since_date is defined.
	SinceId *int64
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 100. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	PerPage *int
	// Providing an invoice_uid allows for scoping of the invoice events to a single invoice or credit note.
	InvoiceUid *string
	// Use this parameter if you want to fetch also invoice events with change_invoice_status type.
	WithChangeInvoiceStatus *string
	// Filter results by event_type. Supply a comma separated list of event types (listed above). Use in query: `event_types=void_invoice,void_remainder`.
	EventTypes []models.InvoiceEventType
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
	input ListInvoiceEventsInput) (
	models.ApiResponse[models.ListInvoiceEventsResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/invoices/events.json")

	req.Authenticate(NewAuth("BasicAuth"))
	if input.SinceDate != nil {
		req.QueryParam("since_date", *input.SinceDate)
	}
	if input.SinceId != nil {
		req.QueryParam("since_id", *input.SinceId)
	}
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.InvoiceUid != nil {
		req.QueryParam("invoice_uid", *input.InvoiceUid)
	}
	if input.WithChangeInvoiceStatus != nil {
		req.QueryParam("with_change_invoice_status", *input.WithChangeInvoiceStatus)
	}
	if input.EventTypes != nil {
		req.QueryParam("event_types", input.EventTypes)
	}
	var result models.ListInvoiceEventsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListInvoiceEventsResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// RecordPaymentForInvoice takes context, uid, body as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// Applies a payment of a given type against a specific invoice. If you would like to apply a payment across multiple invoices, you can use the Bulk Payment endpoint.
func (i *InvoicesController) RecordPaymentForInvoice(
	ctx context.Context,
	uid string,
	body *models.CreateInvoicePaymentRequest) (
	models.ApiResponse[models.Invoice],
	error) {
	req := i.prepareRequest(ctx, "POST", "/invoices/%v/payments.json")
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Invoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Invoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// RecordPaymentForMultipleInvoices takes context, body as parameters and
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
func (i *InvoicesController) RecordPaymentForMultipleInvoices(
	ctx context.Context,
	body *models.CreateMultiInvoicePaymentRequest) (
	models.ApiResponse[models.MultiInvoicePaymentResponse],
	error) {
	req := i.prepareRequest(ctx, "POST", "/invoices/payments.json")

	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}
	var result models.MultiInvoicePaymentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MultiInvoicePaymentResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListCreditNotesInput represents the input of the ListCreditNotes endpoint.
type ListCreditNotesInput struct {
	// The subscription's Advanced Billing id
	SubscriptionId *int
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
	// Include line items data
	LineItems *bool
	// Include discounts data
	Discounts *bool
	// Include taxes data
	Taxes *bool
	// Include refunds data
	Refunds *bool
	// Include applications data
	Applications *bool
}

// ListCreditNotes takes context, subscriptionId, page, perPage, lineItems, discounts, taxes, refunds, applications as parameters and
// returns an models.ApiResponse with models.ListCreditNotesResponse data and
// an error if there was an issue with the request or response.
// Credit Notes are like inverse invoices. They reduce the amount a customer owes.
// By default, the credit notes returned by this endpoint will exclude the arrays of `line_items`, `discounts`, `taxes`, `applications`, or `refunds`. To include these arrays, pass the specific field as a key in the query with a value set to `true`.
func (i *InvoicesController) ListCreditNotes(
	ctx context.Context,
	input ListCreditNotesInput) (
	models.ApiResponse[models.ListCreditNotesResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/credit_notes.json")

	req.Authenticate(NewAuth("BasicAuth"))
	if input.SubscriptionId != nil {
		req.QueryParam("subscription_id", *input.SubscriptionId)
	}
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.LineItems != nil {
		req.QueryParam("line_items", *input.LineItems)
	}
	if input.Discounts != nil {
		req.QueryParam("discounts", *input.Discounts)
	}
	if input.Taxes != nil {
		req.QueryParam("taxes", *input.Taxes)
	}
	if input.Refunds != nil {
		req.QueryParam("refunds", *input.Refunds)
	}
	if input.Applications != nil {
		req.QueryParam("applications", *input.Applications)
	}
	var result models.ListCreditNotesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListCreditNotesResponse](decoder)
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
	req := i.prepareRequest(ctx, "GET", "/credit_notes/%v.json")
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))

	var result models.CreditNote
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.CreditNote](decoder)
	return models.NewApiResponse(result, resp), err
}

// RecordPaymentForSubscription takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.RecordPaymentResponse data and
// an error if there was an issue with the request or response.
// Record an external payment made against a subscription that will pay partially or in full one or more invoices.
// Payment will be applied starting with the oldest open invoice and then next oldest, and so on until the amount of the payment is fully consumed.
// Excess payment will result in the creation of a prepayment on the Invoice Account.
// Only ungrouped or primary subscriptions may be paid using the "bulk" payment request.
func (i *InvoicesController) RecordPaymentForSubscription(
	ctx context.Context,
	subscriptionId int,
	body *models.RecordPaymentRequest) (
	models.ApiResponse[models.RecordPaymentResponse],
	error) {
	req := i.prepareRequest(ctx, "POST", "/subscriptions/%v/payments.json")
	req.AppendTemplateParams(subscriptionId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.RecordPaymentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RecordPaymentResponse](decoder)
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
	req := i.prepareRequest(ctx, "POST", "/invoices/%v/reopen.json")
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})

	var result models.Invoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Invoice](decoder)
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
	req := i.prepareRequest(ctx, "POST", "/invoices/%v/void.json")
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Invoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Invoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListConsolidatedInvoiceSegmentsInput represents the input of the ListConsolidatedInvoiceSegments endpoint.
type ListConsolidatedInvoiceSegmentsInput struct {
	// The unique identifier of the consolidated invoice
	InvoiceUid string
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
	// Sort direction of the returned segments.
	Direction *models.Direction
}

// ListConsolidatedInvoiceSegments takes context, invoiceUid, page, perPage, direction as parameters and
// returns an models.ApiResponse with models.ConsolidatedInvoice data and
// an error if there was an issue with the request or response.
// Invoice segments returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, or `custom_fields`.
func (i *InvoicesController) ListConsolidatedInvoiceSegments(
	ctx context.Context,
	input ListConsolidatedInvoiceSegmentsInput) (
	models.ApiResponse[models.ConsolidatedInvoice],
	error) {
	req := i.prepareRequest(ctx, "GET", "/invoices/%v/segments.json")
	req.AppendTemplateParams(input.InvoiceUid)
	req.Authenticate(NewAuth("BasicAuth"))
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.Direction != nil {
		req.QueryParam("direction", *input.Direction)
	}

	var result models.ConsolidatedInvoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ConsolidatedInvoice](decoder)
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
	req := i.prepareRequest(ctx, "POST", "/subscriptions/%v/invoices.json")
	req.AppendTemplateParams(subscriptionId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorArrayMapResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.InvoiceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.InvoiceResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// SendInvoice takes context, uid, body as parameters and
// returns an *Response and
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
	req := i.prepareRequest(ctx, "POST", "/invoices/%v/deliveries.json")
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
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
		"/invoices/%v/customer_information/preview.json",
	)
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'", Unmarshaller: errors.NewErrorListResponse},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})

	var result models.CustomerChangesPreviewResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.CustomerChangesPreviewResponse](decoder)
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
	req := i.prepareRequest(ctx, "PUT", "/invoices/%v/customer_information.json")
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'", Unmarshaller: errors.NewErrorListResponse},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})

	var result models.Invoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Invoice](decoder)
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
	req := i.prepareRequest(ctx, "POST", "/invoices/%v/issue.json")
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Invoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Invoice](decoder)
	return models.NewApiResponse(result, resp), err
}
