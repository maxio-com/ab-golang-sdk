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

// WalletAdjustmentsController represents a controller struct.
type WalletAdjustmentsController struct {
    baseController
}

// NewWalletAdjustmentsController creates a new instance of WalletAdjustmentsController.
// It takes a baseController as a parameter and returns a pointer to the WalletAdjustmentsController.
func NewWalletAdjustmentsController(baseController baseController) *WalletAdjustmentsController {
    walletAdjustmentsController := WalletAdjustmentsController{baseController: baseController}
    return &walletAdjustmentsController
}

// CreateWalletAdjustment takes context, walletId, idempotencyKey, body as parameters and
// returns an models.ApiResponse with models.WalletGrantResultResponse data and
// an error if there was an issue with the request or response.
// Posts a signed correction against a specific grant (or, for an overdraft entry, against the overdraft): positive adds units to the grant, negative removes them. `reason` is required and is written to the ledger row's `memo`. Writes exactly one ledger row, of kind `adjustment`.
// Use `source_transaction_id` to name a specific `debit` row when you're correcting one recorded consumption rather than the grant's balance as a whole. When it names a grant-attributed debit, it's recorded for audit only: naming it doesn't change how `amount` is applied, and more than one adjustment may name the same row. When it names an overdraft entry, it switches the correction to the overdraft path described below.
// An overdraft entry (a `debit` row with no `grant_id`, written when an `allow_negative` wallet consumes past its grants) can be corrected without a `grant_id`. That correction is credit-only, credits no grant, and is bounded by the usage the entry still records and by the overdraft the wallet still owes. Once a later grant has repaid the overdraft, correct that grant's balance instead.
// An adjustment corrects the ledger, so it works on the grant's recorded remainder, `posted_remaining`, not on `remaining`, which on a metered wallet read through [Read Wallet]($e/Wallets/readWallet) also subtracts usage not yet posted. Send `posted_remaining` back as `expected_remaining`; sending `remaining` on a metered wallet with pending usage returns `409`.
// `expected_remaining` and `expected_source_corrections` are optional optimistic-concurrency checks; see their field descriptions. Use them when the correction's `amount` was derived from a value read earlier that might have moved since.
func (w *WalletAdjustmentsController) CreateWalletAdjustment(
    ctx context.Context,
    walletId int,
    idempotencyKey *string,
    body *models.CreateWalletAdjustmentRequest) (
    models.ApiResponse[models.WalletGrantResultResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/wallets/%v/adjustments.json")
    req.AppendTemplateParams(walletId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "The `Idempotency-Key` header is empty, longer than 255 characters, or contains characters other than printable ASCII. Nothing is written.", Unmarshaller: errors.NewWalletErrorResponse},
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
        "409": {Message: "Returned in two cases. A request with the same `Idempotency-Key` was still being processed when this retry's wait for it timed out (`idempotency_request_in_progress`); retry with the same key once it finishes. Or the state the `amount` was derived from has moved: either the grant's `posted_remaining` no longer matches `expected_remaining`, or the corrections already posted against `source_transaction_id` no longer sum to `expected_source_corrections`. Nothing is written. Reload and re-derive the amount.", Unmarshaller: errors.NewWalletErrorResponse},
        "422": {Message: "Validation failure, or an `Idempotency-Key` reused with a different request (`idempotency_key_mismatch`). Validation failures include a missing `reason`, a `reason` over 255 characters, a missing or not-found `grant_id`, a zero, malformed, or out-of-range `amount`, a negative adjustment exceeding the grant's `posted_remaining`, a positive adjustment exceeding the grant's `original`, a malformed `expected_remaining` or `expected_source_corrections` (including an empty string, which is treated as malformed rather than omitted), a `source_transaction_id` not on this wallet or not a `debit` row, an overdraft correction that debits instead of credits (\"Overdraft usage can only be corrected downward\"), that exceeds the usage the entry still records (\"Adjustment exceeds the usage this entry still records\"), or that exceeds the overdraft the wallet still owes (\"Adjustment exceeds the wallet's outstanding overdraft\", followed by how much can still be corrected), a `grant_id` sent with an overdraft entry, and an archived wallet.", Unmarshaller: errors.NewWalletErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if idempotencyKey != nil {
        req.Header("Idempotency-Key", *idempotencyKey)
    }
    if body != nil {
        req.Json(body)
    }
    
    var result models.WalletGrantResultResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletGrantResultResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
