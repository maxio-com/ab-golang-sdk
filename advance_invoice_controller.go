// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// AdvanceInvoiceController represents a controller struct.
type AdvanceInvoiceController struct {
    baseController
}

// NewAdvanceInvoiceController creates a new instance of AdvanceInvoiceController.
// It takes a baseController as a parameter and returns a pointer to the AdvanceInvoiceController.
func NewAdvanceInvoiceController(baseController baseController) *AdvanceInvoiceController {
    advanceInvoiceController := AdvanceInvoiceController{baseController: baseController}
    return &advanceInvoiceController
}

// IssueAdvanceInvoice takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// Issues an invoice in advance for a subscription's next renewal date. For the most part, advance invoices function like any other invoice, except they are issued early and have special behavior upon being voided. For more information on advance invoices, including eligibility for generating one, see [Issue Invoice In Advance](https://maxio.zendesk.com/hc/en-us/articles/24252026404749-Issue-Invoice-In-Advance).
// A subscription can only have one advance invoice per billing period. Attempting to issue an advance invoice when one already exists returns an error.
// Regeneration of the invoice can be forced with the params `force: true`, which voids an advance invoice if one exists and generates a new one. If no advance invoice exists, a new one is generated.
// Consider using either the create or preview endpoints for proforma invoices to preview this advance invoice before using this endpoint to generate it.
func (a *AdvanceInvoiceController) IssueAdvanceInvoice(
    ctx context.Context,
    subscriptionId int,
    body *models.IssueAdvanceInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error) {
    req := a.prepareRequest(
      ctx,
      "POST",
      "/subscriptions/%v/advance_invoice/issue.json",
    )
    req.AppendTemplateParams(subscriptionId)
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

// ReadAdvanceInvoice takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// Returns the advance invoice generated for a subscription's upcoming renewal. There can only be one advance invoice per subscription per billing cycle.
func (a *AdvanceInvoiceController) ReadAdvanceInvoice(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.Invoice],
    error) {
    req := a.prepareRequest(ctx, "GET", "/subscriptions/%v/advance_invoice.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
    })
    
    var result models.Invoice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Invoice](decoder)
    return models.NewApiResponse(result, resp), err
}

// VoidAdvanceInvoice takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.Invoice data and
// an error if there was an issue with the request or response.
// Voids a subscription's existing advance invoice. Once voided, it can later be regenerated if desired.
// A `reason` is required to void, and the invoice must have an open status. Voiding causes any prepayments and credits that were applied to the invoice to be returned to the subscription.
// For a full overview of the impact of voiding, see [Invoice]($m/Invoice).
func (a *AdvanceInvoiceController) VoidAdvanceInvoice(
    ctx context.Context,
    subscriptionId int,
    body *models.VoidInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error) {
    req := a.prepareRequest(
      ctx,
      "POST",
      "/subscriptions/%v/advance_invoice/void.json",
    )
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
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
