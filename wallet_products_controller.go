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

// WalletProductsController represents a controller struct.
type WalletProductsController struct {
    baseController
}

// NewWalletProductsController creates a new instance of WalletProductsController.
// It takes a baseController as a parameter and returns a pointer to the WalletProductsController.
func NewWalletProductsController(baseController baseController) *WalletProductsController {
    walletProductsController := WalletProductsController{baseController: baseController}
    return &walletProductsController
}

// ListWalletProducts takes context, includeArchived, status, q, sortBy, sortDirection, page, perPage as parameters and
// returns an models.ApiResponse with models.WalletProductsListResponse data and
// an error if there was an issue with the request or response.
// Lists the Wallet Products on your site.
func (w *WalletProductsController) ListWalletProducts(
    ctx context.Context,
    includeArchived *bool,
    status *models.Status1,
    q *string,
    sortBy *models.SortBy1,
    sortDirection *models.SortDirection,
    page *int,
    perPage *int) (
    models.ApiResponse[models.WalletProductsListResponse],
    error) {
    req := w.prepareRequest(ctx, "GET", "/wallet_products.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
    })
    if includeArchived != nil {
        req.QueryParam("include_archived", *includeArchived)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    if q != nil {
        req.QueryParam("q", *q)
    }
    if sortBy != nil {
        req.QueryParam("sort_by", *sortBy)
    }
    if sortDirection != nil {
        req.QueryParam("sort_direction", *sortDirection)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    var result models.WalletProductsListResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletProductsListResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateWalletProduct takes context, body as parameters and
// returns an models.ApiResponse with models.WalletProductResponse data and
// an error if there was an issue with the request or response.
// Creates a Wallet Product: a billing product with no product family, meant to denominate wallets and to be purchased through [Create Wallet Purchase]($e/Wallet%20Purchases/createWalletPurchase). It's a Component under the hood, so it returns the full Component field set, and can optionally be priced with one or more price points in the same call.
func (w *WalletProductsController) CreateWalletProduct(
    ctx context.Context,
    body *models.CreateWalletProductRequest) (
    models.ApiResponse[models.WalletProductResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/wallet_products.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "422": {Message: "Validation failure. Includes a `product_family_id` (rejected rather than ignored), and any failure validating a price point in `price_point`/`price_points`, which rolls back the whole call, including the Wallet Product itself.", Unmarshaller: errors.NewWalletErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.WalletProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadWalletProduct takes context, id as parameters and
// returns an models.ApiResponse with models.WalletProductResponse data and
// an error if there was an issue with the request or response.
// Returns the Wallet Product.
func (w *WalletProductsController) ReadWalletProduct(
    ctx context.Context,
    id string) (
    models.ApiResponse[models.WalletProductResponse],
    error) {
    req := w.prepareRequest(ctx, "GET", "/wallet_products/%v.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
    })
    
    var result models.WalletProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateWalletProduct takes context, id, body as parameters and
// returns an models.ApiResponse with models.WalletProductResponse data and
// an error if there was an issue with the request or response.
// Updates the Wallet Product's catalog-facing fields. Omitted fields are left unchanged. Unlike [Create Wallet Product]($e/Wallet%20Products/createWalletProduct), this does not accept `unit_name`, `handle`, `price_point`, or `price_points`; see the request schema for which fields are actually mutable.
func (w *WalletProductsController) UpdateWalletProduct(
    ctx context.Context,
    id string,
    body *models.UpdateWalletProductRequest) (
    models.ApiResponse[models.WalletProductResponse],
    error) {
    req := w.prepareRequest(ctx, "PUT", "/wallet_products/%v.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
        "422": {Message: "Validation failure.", Unmarshaller: errors.NewWalletErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WalletProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadWalletProductUsage takes context, id as parameters and
// returns an models.ApiResponse with models.WalletProductUsageResponse data and
// an error if there was an issue with the request or response.
// Returns everything drawing on this Wallet Product: the wallet templates billing through it, split into ones using its default price point and ones overriding it, plus a sample of up to 50 of the active wallets actually created from them. A wallet template can appear more than once in `wallet_templates`/`price_point_overrides`, once per Feature and owning plan/component sharing it, so these arrays aren't one row per template.
func (w *WalletProductsController) ReadWalletProductUsage(
    ctx context.Context,
    id string) (
    models.ApiResponse[models.WalletProductUsageResponse],
    error) {
    req := w.prepareRequest(ctx, "GET", "/wallet_products/%v/usage.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
    })
    
    var result models.WalletProductUsageResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletProductUsageResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ArchiveWalletProduct takes context, id as parameters and
// returns an models.ApiResponse with models.WalletProductResponse data and
// an error if there was an issue with the request or response.
// Archives the Wallet Product. Refuses when an active Feature's `wallet_behavior` still points at it, or when it's used in a scheduled renewal. This protects Feature-provisioned wallets, but not a wallet created manually against this product: a manual wallet can be left denominated by an archived Wallet Product, and [Create Wallet Purchase]($e/Wallet%20Purchases/createWalletPurchase) doesn't check whether the billing product is archived before pricing and charging a purchase against it.
func (w *WalletProductsController) ArchiveWalletProduct(
    ctx context.Context,
    id string) (
    models.ApiResponse[models.WalletProductResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/wallet_products/%v/archive.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
        "422": {Message: "The Wallet Product is already archived, is in use by one or more Features, or is in use by a scheduled renewal.", Unmarshaller: errors.NewWalletErrorResponse},
    })
    
    var result models.WalletProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UnarchiveWalletProduct takes context, id as parameters and
// returns an models.ApiResponse with models.WalletProductResponse data and
// an error if there was an issue with the request or response.
// Restores an archived Wallet Product.
func (w *WalletProductsController) UnarchiveWalletProduct(
    ctx context.Context,
    id string) (
    models.ApiResponse[models.WalletProductResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/wallet_products/%v/unarchive.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site."},
        "404": {Message: "Not Found", Unmarshaller: errors.NewWalletErrorResponse},
        "422": {Message: "The Wallet Product is not archived.", Unmarshaller: errors.NewWalletErrorResponse},
    })
    
    var result models.WalletProductResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WalletProductResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
