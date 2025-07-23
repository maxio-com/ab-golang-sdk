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

// ProformaInvoicesController represents a controller struct.
type ProformaInvoicesController struct {
	baseController
}

// NewProformaInvoicesController creates a new instance of ProformaInvoicesController.
// It takes a baseController as a parameter and returns a pointer to the ProformaInvoicesController.
func NewProformaInvoicesController(baseController baseController) *ProformaInvoicesController {
	proformaInvoicesController := ProformaInvoicesController{baseController: baseController}
	return &proformaInvoicesController
}

// CreateConsolidatedProformaInvoice takes context, uid as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This endpoint will trigger the creation of a consolidated proforma invoice asynchronously. It will return a 201 with no message, or a 422 with any errors. To find and view the new consolidated proforma invoice, you may poll the subscription group listing for proforma invoices; only one consolidated proforma invoice may be created per group at a time.
// If the information becomes outdated, simply void the old consolidated proforma invoice and generate a new one.
// ## Restrictions
// Proforma invoices are only available on Relationship Invoicing sites. To create a proforma invoice, the subscription must not be prepaid, and must be in a live state.
func (p *ProformaInvoicesController) CreateConsolidatedProformaInvoice(
	ctx context.Context,
	uid string) (
	*http.Response,
	error) {
	req := p.prepareRequest(
		ctx,
		"POST",
		"/subscription_groups/%v/proforma_invoices.json",
	)
	req.AppendTemplateParams(uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// ListSubscriptionGroupProformaInvoicesInput represents the input of the ListSubscriptionGroupProformaInvoices endpoint.
type ListSubscriptionGroupProformaInvoicesInput struct {
	// The uid of the subscription group
	Uid string
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
}

// ListSubscriptionGroupProformaInvoices takes context, uid, lineItems, discounts, taxes, credits, payments, customFields as parameters and
// returns an models.ApiResponse with models.ListProformaInvoicesResponse data and
// an error if there was an issue with the request or response.
// Only proforma invoices with a `consolidation_level` of parent are returned.
// By default, proforma invoices returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, `custom_fields`. To include breakdowns, pass the specific field as a key in the query with a value set to true.
func (p *ProformaInvoicesController) ListSubscriptionGroupProformaInvoices(
	ctx context.Context,
	input ListSubscriptionGroupProformaInvoicesInput) (
	models.ApiResponse[models.ListProformaInvoicesResponse],
	error) {
	req := p.prepareRequest(
		ctx,
		"GET",
		"/subscription_groups/%v/proforma_invoices.json",
	)
	req.AppendTemplateParams(input.Uid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})
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

	var result models.ListProformaInvoicesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListProformaInvoicesResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ReadProformaInvoice takes context, proformaInvoiceUid as parameters and
// returns an models.ApiResponse with models.ProformaInvoice data and
// an error if there was an issue with the request or response.
// Use this endpoint to read the details of an existing proforma invoice.
// ## Restrictions
// Proforma invoices are only available on Relationship Invoicing sites.
func (p *ProformaInvoicesController) ReadProformaInvoice(
	ctx context.Context,
	proformaInvoiceUid string) (
	models.ApiResponse[models.ProformaInvoice],
	error) {
	req := p.prepareRequest(ctx, "GET", "/proforma_invoices/%v.json")
	req.AppendTemplateParams(proformaInvoiceUid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})

	var result models.ProformaInvoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ProformaInvoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateProformaInvoice takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.ProformaInvoice data and
// an error if there was an issue with the request or response.
// This endpoint will create a proforma invoice and return it as a response. If the information becomes outdated, simply void the old proforma invoice and generate a new one.
// If you would like to preview the next billing amounts without generating a full proforma invoice, please use the renewal preview endpoint.
// ## Restrictions
// Proforma invoices are only available on Relationship Invoicing sites. To create a proforma invoice, the subscription must not be in a group, must not be prepaid, and must be in a live state.
func (p *ProformaInvoicesController) CreateProformaInvoice(
	ctx context.Context,
	subscriptionId int) (
	models.ApiResponse[models.ProformaInvoice],
	error) {
	req := p.prepareRequest(ctx, "POST", "/subscriptions/%v/proforma_invoices.json")
	req.AppendTemplateParams(subscriptionId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})

	var result models.ProformaInvoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ProformaInvoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListProformaInvoicesInput represents the input of the ListProformaInvoices endpoint.
type ListProformaInvoicesInput struct {
	// The Chargify id of the subscription
	SubscriptionId int
	// The beginning date range for the invoice's Due Date, in the YYYY-MM-DD format.
	StartDate *string
	// The ending date range for the invoice's Due Date, in the YYYY-MM-DD format.
	EndDate *string
	// The current status of the invoice.  Allowed Values: draft, open, paid, pending, voided
	Status *models.ProformaInvoiceStatus
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
}

// ListProformaInvoices takes context, subscriptionId, startDate, endDate, status, page, perPage, direction, lineItems, discounts, taxes, credits, payments, customFields as parameters and
// returns an models.ApiResponse with models.ListProformaInvoicesResponse data and
// an error if there was an issue with the request or response.
// By default, proforma invoices returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, or `custom_fields`. To include breakdowns, pass the specific field as a key in the query with a value set to `true`.
func (p *ProformaInvoicesController) ListProformaInvoices(
	ctx context.Context,
	input ListProformaInvoicesInput) (
	models.ApiResponse[models.ListProformaInvoicesResponse],
	error) {
	req := p.prepareRequest(ctx, "GET", "/subscriptions/%v/proforma_invoices.json")
	req.AppendTemplateParams(input.SubscriptionId)
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

	var result models.ListProformaInvoicesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListProformaInvoicesResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// VoidProformaInvoice takes context, proformaInvoiceUid, body as parameters and
// returns an models.ApiResponse with models.ProformaInvoice data and
// an error if there was an issue with the request or response.
// This endpoint will void a proforma invoice that has the status "draft".
// ## Restrictions
// Proforma invoices are only available on Relationship Invoicing sites.
// Only proforma invoices that have the appropriate status may be reopened. If the invoice identified by {uid} does not have the appropriate status, the response will have HTTP status code 422 and an error message.
// A reason for the void operation is required to be included in the request body. If one is not provided, the response will have HTTP status code 422 and an error message.
func (p *ProformaInvoicesController) VoidProformaInvoice(
	ctx context.Context,
	proformaInvoiceUid string,
	body *models.VoidInvoiceRequest) (
	models.ApiResponse[models.ProformaInvoice],
	error) {
	req := p.prepareRequest(ctx, "POST", "/proforma_invoices/%v/void.json")
	req.AppendTemplateParams(proformaInvoiceUid)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ProformaInvoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ProformaInvoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// PreviewProformaInvoice takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.ProformaInvoice data and
// an error if there was an issue with the request or response.
// Return a preview of the data that will be included on a given subscription's proforma invoice if one were to be generated. It will have similar line items and totals as a renewal preview, but the response will be presented in the format of a proforma invoice. Consequently it will include additional information such as the name and addresses that will appear on the proforma invoice.
// The preview endpoint is subject to all the same conditions as the proforma invoice endpoint. For example, previews are only available on the Relationship Invoicing architecture, and previews cannot be made for end-of-life subscriptions.
// If all the data returned in the preview is as expected, you may then create a static proforma invoice and send it to your customer. The data within a preview will not be saved and will not be accessible after the call is made.
// Alternatively, if you have some proforma invoices already, you may make a preview call to determine whether any billing information for the subscription's upcoming renewal has changed.
func (p *ProformaInvoicesController) PreviewProformaInvoice(
	ctx context.Context,
	subscriptionId int) (
	models.ApiResponse[models.ProformaInvoice],
	error) {
	req := p.prepareRequest(
		ctx,
		"POST",
		"/subscriptions/%v/proforma_invoices/preview.json",
	)
	req.AppendTemplateParams(subscriptionId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})

	var result models.ProformaInvoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ProformaInvoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSignupProformaInvoice takes context, body as parameters and
// returns an models.ApiResponse with models.ProformaInvoice data and
// an error if there was an issue with the request or response.
// This endpoint is only available for Relationship Invoicing sites. It cannot be used to create consolidated proforma invoices or preview prepaid subscriptions.
// Create a proforma invoice to preview costs before a subscription's signup. Like other proforma invoices, it can be emailed to the customer, voided, and publicly viewed on the chargifypay domain.
// Pass a payload that resembles a subscription create or signup preview request. For example, you can specify components, coupons/a referral, offers, custom pricing, and an existing customer or payment profile to populate a shipping or billing address.
// A product and customer first name, last name, and email are the minimum requirements. We recommend associating the proforma invoice with a customer_id to easily find their proforma invoices, since the subscription_id will always be blank.
func (p *ProformaInvoicesController) CreateSignupProformaInvoice(
	ctx context.Context,
	body *models.CreateSubscriptionRequest) (
	models.ApiResponse[models.ProformaInvoice],
	error) {
	req := p.prepareRequest(ctx, "POST", "/subscriptions/proforma_invoices.json")

	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewProformaBadRequestErrorResponse},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorArrayMapResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}
	var result models.ProformaInvoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ProformaInvoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// PreviewSignupProformaInvoice takes context, include, body as parameters and
// returns an models.ApiResponse with models.SignupProformaPreviewResponse data and
// an error if there was an issue with the request or response.
// This endpoint is only available for Relationship Invoicing sites. It cannot be used to create consolidated proforma invoice previews or preview prepaid subscriptions.
// Create a signup preview in the format of a proforma invoice to preview costs before a subscription's signup. You have the option of optionally previewing the first renewal's costs as well. The proforma invoice preview will not be persisted.
// Pass a payload that resembles a subscription create or signup preview request. For example, you can specify components, coupons/a referral, offers, custom pricing, and an existing customer or payment profile to populate a shipping or billing address.
// A product and customer first name, last name, and email are the minimum requirements.
func (p *ProformaInvoicesController) PreviewSignupProformaInvoice(
	ctx context.Context,
	include *models.CreateSignupProformaPreviewInclude,
	body *models.CreateSubscriptionRequest) (
	models.ApiResponse[models.SignupProformaPreviewResponse],
	error) {
	req := p.prepareRequest(
		ctx,
		"POST",
		"/subscriptions/proforma_invoices/preview.json",
	)

	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewProformaBadRequestErrorResponse},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorArrayMapResponse},
	})
	req.Header("Content-Type", "application/json")
	if include != nil {
		req.QueryParam("include", *include)
	}
	if body != nil {
		req.Json(body)
	}
	var result models.SignupProformaPreviewResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SignupProformaPreviewResponse](decoder)
	return models.NewApiResponse(result, resp), err
}
