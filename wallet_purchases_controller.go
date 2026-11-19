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

// WalletPurchasesController represents a controller struct.
type WalletPurchasesController struct {
    baseController
}

// NewWalletPurchasesController creates a new instance of WalletPurchasesController.
// It takes a baseController as a parameter and returns a pointer to the WalletPurchasesController.
func NewWalletPurchasesController(baseController baseController) *WalletPurchasesController {
    walletPurchasesController := WalletPurchasesController{baseController: baseController}
    return &walletPurchasesController
}

// CreateWalletPurchase takes context, walletId, idempotencyKey, body as parameters and
// returns an models.ApiResponse with models.WalletPurchaseResultResponse data and
// an error if there was an issue with the request or response.
// Purchases prepaid units for the wallet and charges the subscription's payment method for them, using the wallet's `price_point_id` unless a different price point is named. A successful purchase writes a grant of origin `purchase`, with its expiration following the funding entitlement's grant expiration policy.
// How "charged" works depends on the subscription's collection method. For card/automatic collection, the payment is attempted immediately; a decline writes nothing to the wallet, and the call fails with a `422` naming the gateway's decline reason. For invoice-billed (remittance) subscriptions, there's no immediate charge to decline: an invoice is drafted, the wallet is credited right away, and the invoice is issued afterward. If issuing the invoice fails at that point, the invoice is left in draft state and reported for follow-up, but the call still returns `201` and the units already credited are **not** rolled back. See [Preview Wallet Purchase]($e/Wallet%20Purchases/previewWalletPurchase) to quote a price before charging anything.
func (w *WalletPurchasesController) CreateWalletPurchase(
    ctx context.Context,
    walletId int,
    idempotencyKey *string,
    body *models.CreateWalletPurchaseRequest) (
    models.ApiResponse[models.WalletPurchaseResultResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/wallets/%v/purchases.json")
    req.AppendTemplateParams(walletId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "The `Idempotency-Key` header is empty, longer than 255 characters, or contains characters other than printable ASCII. Nothing is written.", Unmarshaller: errors.NewWalletErrorResponse},
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
        "409": {Message: "Returned in two cases, both with nothing new charged or written. `idempotency_request_in_progress`: a purchase with the same `Idempotency-Key` is still running; retry with the same key once it finishes. `idempotency_key_outcome_unknown`: an earlier purchase with this key ended with a server error or an unexpected failure before its result was recorded, so it may or may not have charged the customer. Check the subscription's invoices before buying again under a new key.", Unmarshaller: errors.NewWalletErrorResponse},
        "422": {Message: "Validation or payment failure, or an `Idempotency-Key` reused with a different request (`idempotency_key_mismatch`). Validation failures include a missing `purchase` root, a missing, malformed, non-positive, or out-of-range `quantity`, a `price_point_id` that doesn't belong to the wallet's billing product, an archived, custom, or otherwise unusable price point, no price point to purchase at (neither sent nor set on the wallet), a memo over 255 characters, an archived wallet, a price point with no price in the subscription's currency, a declined payment (message includes the gateway's own detail), an invoice that couldn't be created at all, a payment that settled but whose units could not be credited (`purchase_paid_but_uncredited`, whose message ends with the paid invoice's uid so it can be reconciled), a declined payment whose invoice could not be voided and is still open (`purchase_declined_invoice_open`), a purchase that failed before payment and left its invoice open (`purchase_invoice_left_open`), and a wallet whose billing product isn't a Wallet Product (only possible for a wallet created before Wallet Products became required).\n\n`purchase_paid_but_uncredited`, `purchase_declined_invoice_open` and `purchase_invoice_left_open` bodies also carry `invoice_uid` and `invoice_number`. With an `Idempotency-Key`, these three responses are stored and replayed like a success.", Unmarshaller: errors.NewWalletErrorResponse},
        "500": {Message: "`purchase_payment_outcome_unknown`: the payment was attempted but the gateway did not confirm whether the card was charged. Nothing is credited by this call; if the invoice turns out to be paid, the units are credited automatically (see [Purchasing Units](page:introduction/basic-concepts/wallets#purchasing-units)). The body includes `invoice_uid` and `invoice_number` when an invoice was created; check that invoice before buying again. Retrying with the same `Idempotency-Key` returns `409` (`idempotency_key_outcome_unknown`) rather than charging again.", Unmarshaller: errors.NewWalletErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if idempotencyKey != nil {
        req.Header("Idempotency-Key", *idempotencyKey)
    }
    if body != nil {
        req.Json(body)
    }
    
    var result models.WalletPurchaseResultResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletPurchaseResultResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// PreviewWalletPurchase takes context, walletId, body as parameters and
// returns an models.ApiResponse with models.WalletPurchasePreviewResponse data and
// an error if there was an issue with the request or response.
// Quotes the price for a purchase at the given quantity and price point, without charging the payment method or writing to the wallet. Takes the same parameters as [Create Wallet Purchase]($e/Wallet%20Purchases/createWalletPurchase).
func (w *WalletPurchasesController) PreviewWalletPurchase(
    ctx context.Context,
    walletId int,
    body *models.CreateWalletPurchaseRequest) (
    models.ApiResponse[models.WalletPurchasePreviewResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/wallets/%v/purchases/preview.json")
    req.AppendTemplateParams(walletId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
        "422": {Message: "Validation failure. Same cases as Create Wallet Purchase, minus the payment-outcome cases, since nothing is charged.", Unmarshaller: errors.NewWalletErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WalletPurchasePreviewResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletPurchasePreviewResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
