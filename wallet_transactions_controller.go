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

// WalletTransactionsController represents a controller struct.
type WalletTransactionsController struct {
    baseController
}

// NewWalletTransactionsController creates a new instance of WalletTransactionsController.
// It takes a baseController as a parameter and returns a pointer to the WalletTransactionsController.
func NewWalletTransactionsController(baseController baseController) *WalletTransactionsController {
    walletTransactionsController := WalletTransactionsController{baseController: baseController}
    return &walletTransactionsController
}

// ListWalletTransactions takes context, walletId, page, perPage as parameters and
// returns an models.ApiResponse with models.WalletTransactionsListResponse data and
// an error if there was an issue with the request or response.
// Lists the wallet's append-only ledger, newest first, paginated.
func (w *WalletTransactionsController) ListWalletTransactions(
    ctx context.Context,
    walletId int,
    page *int,
    perPage *int) (
    models.ApiResponse[models.WalletTransactionsListResponse],
    error) {
    req := w.prepareRequest(ctx, "GET", "/wallets/%v/transactions.json")
    req.AppendTemplateParams(walletId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
    })
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    
    var result models.WalletTransactionsListResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletTransactionsListResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateWalletTransaction takes context, walletId, idempotencyKey, body as parameters and
// returns an models.ApiResponse with models.WalletTransactionResultResponse data and
// an error if there was an issue with the request or response.
// Grants units to the wallet, or debits units from it, depending on `kind`. `amount` is always sent positive, and `kind` determines the direction written to the ledger.
// **`kind: grant`**: a one-time manual top-up. On a non-negative balance, this creates a grant and writes one ledger row. On a negative balance, it repays the overdraft first (a `kind: grant` row with `grant_id: null`, for the repaid amount), then creates the grant: its `original` is always the full `amount` sent, and its `remaining` is what's left after the overdraft was repaid, which is `0` if the overdraft absorbed the whole top-up. The grant row is still written, and its `amount` is that `remaining`, not `original`, so the ledger rows sum to the balance change. A grant may carry `expires_at`; a debit may not.
// **`kind: debit`**: direct consumption. Drains grants earliest-expiration-first, one `kind: debit` row per grant touched, plus any overdue expiration rows this same call triggers first (see below):
// - Under the wallet's default `hard_stop_overflow: discard`, the drain is capped at the available balance. For a metered wallet, that cap is the live metered balance, not the stored one, so a debit can drain less than its stored balance would suggest, or even drain nothing at all (`201` with `drained_amount: "0.0"`) if live usage has already consumed what the ledger hadn't caught up to yet. A debit against a wallet whose *stored* balance is already zero or negative (`discard` only ever reaches zero) returns `422 "Wallet is depleted"`. Before that failure is returned, this call still expires any grants that were already overdue, so a `422` response can still leave new `kind: expiration` rows written.
// - Under `hard_stop_overflow: allow_negative`, the full amount is deducted. The excess past the wallet's grants is booked as one overdraft row with `grant_id: null`, `rejected_amount` is always `"0.0"`, and the balance goes negative.
func (w *WalletTransactionsController) CreateWalletTransaction(
    ctx context.Context,
    walletId int,
    idempotencyKey *string,
    body *models.CreateWalletTransactionRequest) (
    models.ApiResponse[models.WalletTransactionResultResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/wallets/%v/transactions.json")
    req.AppendTemplateParams(walletId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "The `Idempotency-Key` header is empty, longer than 255 characters, or contains characters other than printable ASCII. Nothing is written.", Unmarshaller: errors.NewWalletErrorResponse},
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
        "409": {Message: "A request with the same `Idempotency-Key` was still being processed when this retry's wait for it timed out (`idempotency_request_in_progress`). Nothing is written; retry with the same key once it finishes.", Unmarshaller: errors.NewWalletErrorResponse},
        "422": {Message: "Validation failure, or an `Idempotency-Key` reused with a different request (`idempotency_key_mismatch`). Validation failures include an unrecognized `kind`, a missing `wallet_transaction` root, a malformed, non-positive, or out-of-range `amount`, a memo over 255 characters, `expires_at` sent with `kind: debit`, a malformed or past-dated `expires_at`, an archived wallet, and, for `kind: debit` only, a depleted wallet.", Unmarshaller: errors.NewWalletErrorResponse},
        "500": {Message: "For a metered wallet, usage could be read but settling it against the ledger failed.", Unmarshaller: errors.NewWalletErrorResponse},
        "503": {Message: "For a metered wallet, live usage could not be read in time to compute the debit. Not returned for `kind: grant`.", Unmarshaller: errors.NewWalletErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if idempotencyKey != nil {
        req.Header("Idempotency-Key", *idempotencyKey)
    }
    if body != nil {
        req.Json(body)
    }
    
    var result models.WalletTransactionResultResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletTransactionResultResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
