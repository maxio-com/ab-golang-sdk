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

// WalletsController represents a controller struct.
type WalletsController struct {
    baseController
}

// NewWalletsController creates a new instance of WalletsController.
// It takes a baseController as a parameter and returns a pointer to the WalletsController.
func NewWalletsController(baseController baseController) *WalletsController {
    walletsController := WalletsController{baseController: baseController}
    return &walletsController
}

// ListWallets takes context, page, perPage, subscriptionId, status as parameters and
// returns an models.ApiResponse with models.WalletsListResponse data and
// an error if there was an issue with the request or response.
// Lists the wallets on your site, in creation order, paginated. Without `subscription_id`, returns every wallet on the site. The API key is the tenancy boundary; `subscription_id` only narrows the result set further. Archived wallets are included unless you filter them out with `status=active`.
func (w *WalletsController) ListWallets(
    ctx context.Context,
    page *int,
    perPage *int,
    subscriptionId *int,
    status *models.Status2) (
    models.ApiResponse[models.WalletsListResponse],
    error) {
    req := w.prepareRequest(ctx, "GET", "/wallets.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewWalletErrorResponse},
    })
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if subscriptionId != nil {
        req.QueryParam("subscription_id", *subscriptionId)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    var result models.WalletsListResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletsListResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateWallet takes context, idempotencyKey, body as parameters and
// returns an models.ApiResponse with models.WalletResponse data and
// an error if there was an issue with the request or response.
// Creates a wallet owned by a subscription and denominated by a [Wallet Product]($e/Wallet%20Products/createWalletProduct): only a Wallet Product can denominate a wallet, not an ordinary component. The Wallet Product must be active; it never belongs to a product family, so there's no family check against the subscription. It's **immutable for the life of the wallet**; there's no way to repoint an existing wallet at a different one.
// This endpoint always creates a `manual` wallet. To provision a wallet automatically from a usage-limit entitlement instead, configure `wallet_behavior` on the Feature behind that entitlement (see [the Wallets concept guide](page:introduction/basic-concepts/wallets#provisioning-a-wallet-from-a-feature)) rather than sending `entitlement_id` here.
func (w *WalletsController) CreateWallet(
    ctx context.Context,
    idempotencyKey *string,
    body *models.CreateWalletRequest) (
    models.ApiResponse[models.WalletResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/wallets.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "The `Idempotency-Key` header is empty, longer than 255 characters, or contains characters other than printable ASCII. Nothing is written.", Unmarshaller: errors.NewWalletErrorResponse},
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Unknown or cross-site subscription or meter, or a `wallet_product_id` that doesn't resolve to a Wallet Product on this site.", Unmarshaller: errors.NewWalletErrorResponse},
        "409": {Message: "A request with the same `Idempotency-Key` was still being processed when this retry's wait for it timed out (`idempotency_request_in_progress`). Nothing is written; retry with the same key once it finishes.", Unmarshaller: errors.NewWalletErrorResponse},
        "422": {Message: "Validation failure, or an `Idempotency-Key` reused with a different request (`idempotency_key_mismatch`). Validation failures include a missing or archived `wallet_product_id`; a `price_point_id` that's missing, archived, a subscription-specific custom price point, or doesn't belong to the Wallet Product; an invalid `depletion_policy` or `hard_stop_overflow` value; a meter that isn't a `sum` or `count` meter; an invalid `conversion_rate`; a subscription with no reference, when pairing a meter whose stream bills by subscription reference; and the `initial_balance` / `initial_balance_expires_at` error cases described on those fields. Also returned for `entitlement_id`: this endpoint never accepts it, and rejects it with the message shown below rather than silently ignoring it. Also returned when `subscription_id` names a prepaid subscription: a wallet would double-meter a balance the subscription already tracks on its own, so the two are mutually exclusive.", Unmarshaller: errors.NewWalletErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if idempotencyKey != nil {
        req.Header("Idempotency-Key", *idempotencyKey)
    }
    if body != nil {
        req.Json(body)
    }
    var result models.WalletResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadWallet takes context, id as parameters and
// returns an models.ApiResponse with models.WalletResponse data and
// an error if there was an issue with the request or response.
// Returns the wallet's balance and grant breakdown. For an active, metered wallet on a site with metered consumption enabled, this is the only operation that computes a live balance from meter usage; see the wallet resource's `metered_*` fields. It is also the only operation where `balance` can differ from `posted_balance`, and a grant's `remaining` from its `posted_remaining`: the pending usage is drawn from the grants in drain order, so the grants' `remaining` add up to `balance`. Readable regardless of status. Archiving a wallet doesn't affect reads.
func (w *WalletsController) ReadWallet(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.WalletResponse],
    error) {
    req := w.prepareRequest(ctx, "GET", "/wallets/%v.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
    })
    
    var result models.WalletResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ArchiveWallet takes context, id as parameters and
// returns an models.ApiResponse with models.WalletResponse data and
// an error if there was an issue with the request or response.
// Archives the wallet: blocks further grants, debits, adjustments, and purchases (each returns `422`), stops any metered drawdown, and stops recurring funding (its schedule is deactivated with reason `wallet_archived`, not removed), while reads and the ledger stay available. Nothing already on the wallet is destroyed or clawed back, but nothing on it can be spent either, by any path, once archived. Idempotent: archiving an already-archived wallet returns the same `200`. There is no restore endpoint.
func (w *WalletsController) ArchiveWallet(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.WalletResponse],
    error) {
    req := w.prepareRequest(ctx, "DELETE", "/wallets/%v.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
    })
    
    var result models.WalletResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateWallet takes context, id, body as parameters and
// returns an models.ApiResponse with models.WalletResponse data and
// an error if there was an issue with the request or response.
// Updates the wallet's editable attributes. Only `name` is editable; any other attribute in the body is ignored. Send only the attributes to change.
func (w *WalletsController) UpdateWallet(
    ctx context.Context,
    id int,
    body *models.UpdateWalletRequest) (
    models.ApiResponse[models.WalletResponse],
    error) {
    req := w.prepareRequest(ctx, "PATCH", "/wallets/%v.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
        "422": {Message: "Validation failure. Includes a body with no editable attribute (`parameter_missing`), a `name` over 255 characters (`validation_failed`), and an archived wallet (`wallet_archived`).", Unmarshaller: errors.NewWalletErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WalletResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
