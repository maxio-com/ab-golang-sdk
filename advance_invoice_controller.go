package advancedbilling

import (
    "context"
    "fmt"
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
// Generate an invoice in advance for a subscription's next renewal date. [Please see our docs](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404811062541-Issue-Invoice-In-Advance  ) for more information on advance invoices, including eligibility on generating one; for the most part, they function like any other invoice, except they are issued early and have special behavior upon being voided.
// A subscription may only have one advance invoice per billing period. Attempting to issue an advance invoice when one already exists will return an error.
// That said, regeneration of the invoice may be forced with the params `force: true`, which will void an advance invoice if one exists and generate a new one. If no advance invoice exists, a new one will be generated.
// We recommend using either the create or preview endpoints for proforma invoices to preview this advance invoice before using this endpoint to generate it.
func (a *AdvanceInvoiceController) IssueAdvanceInvoice(
    ctx context.Context,
    subscriptionId int,
    body *models.IssueAdvanceInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error) {
    req := a.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/advance_invoice/issue.json", subscriptionId),
    )
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
// Once an advance invoice has been generated for a subscription's upcoming renewal, it can be viewed through this endpoint. There can only be one advance invoice per subscription per billing cycle.
func (a *AdvanceInvoiceController) ReadAdvanceInvoice(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.Invoice],
    error) {
    req := a.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/subscriptions/%v/advance_invoice.json", subscriptionId),
    )
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
// Void a subscription's existing advance invoice. Once voided, it can later be regenerated if desired.
// A `reason` is required in order to void, and the invoice must have an open status. Voiding will cause any prepayments and credits that were applied to the invoice to be returned to the subscription. For a full overview of the impact of voiding, please [see our help docs]($m/Invoice).
func (a *AdvanceInvoiceController) VoidAdvanceInvoice(
    ctx context.Context,
    subscriptionId int,
    body *models.VoidInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error) {
    req := a.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/advance_invoice/void.json", subscriptionId),
    )
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
