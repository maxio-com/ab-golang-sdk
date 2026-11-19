# Wallets

The Wallet object is a balance, owned by a subscription and denominated in the unit of a Wallet Product, that you can create, fund, drain, and read through the API.

Wallets suit prepaid allowances, credits, and consumption pools that don't map cleanly onto the regular usage-billing cycle. Every balance change is recorded on the wallet's append-only ledger. For more information, see [Understand Wallets](https://docs.maxio.com/billing/catalog/wallets/understand-wallets).

```go
walletsController := client.WalletsController()
```

## Class Name

`WalletsController`

## Methods

* [List Wallets](../../doc/controllers/wallets.md#list-wallets)
* [Create Wallet](../../doc/controllers/wallets.md#create-wallet)
* [Read Wallet](../../doc/controllers/wallets.md#read-wallet)
* [Archive Wallet](../../doc/controllers/wallets.md#archive-wallet)
* [Update Wallet](../../doc/controllers/wallets.md#update-wallet)


# List Wallets

Lists the wallets on your site, in creation order, paginated. Without `subscription_id`, returns every wallet on the site. The API key is the tenancy boundary; `subscription_id` only narrows the result set further. Archived wallets are included unless you filter them out with `status=active`.

```go
ListWallets(
    ctx context.Context,
    page *int,
    perPage *int,
    subscriptionId *int,
    status *models.Status2) (
    models.ApiResponse[models.WalletsListResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `subscriptionId` | `*int` | Query, Optional | Restricts the result to wallets owned by this subscription. |
| `status` | [`*models.Status2`](../../doc/models/status-2.md) | Query, Optional | Restricts the result by wallet status. An unrecognized value returns `422` rather than being silently ignored. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletsListResponse](../../doc/models/wallets-list-response.md).

## Example Usage

```go
ctx := context.Background()

page := 1

perPage := 50

apiResponse, err := walletsController.ListWallets(ctx, &page, &perPage, nil, nil)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.WalletErrorResponse:
            log.Fatalln("WalletErrorResponseException: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "wallets": [
    {
      "id": 5,
      "owner_type": "subscription",
      "owner_id": 23,
      "name": "Onboarding credits",
      "wallet_product_id": 76,
      "wallet_product_name": "API Calls",
      "unit": "calls",
      "price_point_id": 140,
      "balance": "150.0",
      "posted_balance": "150.0",
      "meter": null,
      "metered_through_at": null,
      "metered_recheck_after_seconds": null,
      "metered_stale": null,
      "status": "active",
      "depletion_policy": "hard_stop",
      "hard_stop_overflow": "discard",
      "origin": "manual",
      "feature_key": null,
      "grants": [
        {
          "id": 9,
          "name": "Grant",
          "original": "100.0",
          "remaining": "50.0",
          "posted_remaining": "50.0",
          "expires_at": null,
          "expires_on": null,
          "expired_at": null,
          "expired_on": null,
          "origin": "manual",
          "granted_for_period_start": null,
          "invoice_uid": null,
          "created_at": "2026-07-15T10:00:00Z"
        }
      ],
      "recurring_funding": [],
      "created_at": "2026-07-15T10:00:00Z",
      "updated_at": "2026-08-01T09:30:00Z"
    }
  ],
  "meta": {
    "total_count": 1,
    "current_page": 1,
    "total_pages": 1,
    "per_page": 20
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site. | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |


# Create Wallet

Creates a wallet owned by a subscription and denominated by a [Wallet Product](../../doc/controllers/wallet-products.md#create-wallet-product): only a Wallet Product can denominate a wallet, not an ordinary component. The Wallet Product must be active; it never belongs to a product family, so there's no family check against the subscription. It's **immutable for the life of the wallet**; there's no way to repoint an existing wallet at a different one.

This endpoint always creates a `manual` wallet. To provision a wallet automatically from a usage-limit entitlement instead, configure `wallet_behavior` on the Feature behind that entitlement (see [the Wallets concept guide](http://localhost:8080/go)) rather than sending `entitlement_id` here.

```go
CreateWallet(
    ctx context.Context,
    idempotencyKey *string,
    body *models.CreateWalletRequest) (
    models.ApiResponse[models.WalletResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `idempotencyKey` | `*string` | Header, Optional | A key you generate, such as a UUID, that makes retrying this request safe. 1 to 255 printable ASCII characters; surrounding double quotes are ignored. Keys are scoped to your site and kept for at least 24 hours.<br><br>The first request with a key runs normally. A retry with the same key, method, path and query string, and a byte-identical body returns the stored status and body of that first request without running it again. The same key with a different request returns `422` (`idempotency_key_mismatch`). A retry that arrives while the first request is still running waits for it to finish, and returns `409` (`idempotency_request_in_progress`) only if that wait times out; on Create Wallet Purchase the `409` is immediate. Only successful responses are stored, so a failed request can be retried with the same key. When this header is sent, `uniqueness_token` is not checked.<br><br>See [Idempotent retries](http://localhost:8080/go) for the details, including how wallet purchases differ.<br><br>**Constraints**: *Minimum Length*: `1`, *Maximum Length*: `255` |
| `body` | [`*models.CreateWalletRequest`](../../doc/models/create-wallet-request.md) | Body, Optional | - |

## Response Type

**201**: Wallet created. The balance is `"0.0"` with no grants, unless a non-zero `initial_balance` was given. In that case, the balance and its seeded grant are already reflected.

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletResponse](../../doc/models/wallet-response.md).

## Example Usage

```go
ctx := context.Background()

body := models.CreateWalletRequest{
    Wallet:               models.Wallet1{
        SubscriptionId:          23,
        WalletProductId:         models.Wallet1WalletProductIdContainer.FromNumber(76),
    },
}

apiResponse, err := walletsController.CreateWallet(ctx, nil, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.WalletErrorResponse:
            log.Fatalln("WalletErrorResponseException: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "wallet": {
    "id": 6,
    "owner_type": "subscription",
    "owner_id": 23,
    "name": null,
    "wallet_product_id": 76,
    "wallet_product_name": "API Calls",
    "unit": "calls",
    "price_point_id": 140,
    "balance": "0.0",
    "posted_balance": "0.0",
    "meter": null,
    "metered_through_at": null,
    "metered_recheck_after_seconds": null,
    "metered_stale": null,
    "status": "active",
    "depletion_policy": "hard_stop",
    "hard_stop_overflow": "discard",
    "origin": "manual",
    "feature_key": null,
    "grants": [],
    "recurring_funding": [],
    "created_at": "2026-08-01T09:30:00Z",
    "updated_at": "2026-08-01T09:30:00Z"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | The `Idempotency-Key` header is empty, longer than 255 characters, or contains characters other than printable ASCII. Nothing is written. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 403 | Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site. | `ApiError` |
| 404 | Unknown or cross-site subscription or meter, or a `wallet_product_id` that doesn't resolve to a Wallet Product on this site. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 409 | A request with the same `Idempotency-Key` was still being processed when this retry's wait for it timed out (`idempotency_request_in_progress`). Nothing is written; retry with the same key once it finishes. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 422 | Validation failure, or an `Idempotency-Key` reused with a different request (`idempotency_key_mismatch`). Validation failures include a missing or archived `wallet_product_id`; a `price_point_id` that's missing, archived, a subscription-specific custom price point, or doesn't belong to the Wallet Product; an invalid `depletion_policy` or `hard_stop_overflow` value; a meter that isn't a `sum` or `count` meter; an invalid `conversion_rate`; a subscription with no reference, when pairing a meter whose stream bills by subscription reference; and the `initial_balance` / `initial_balance_expires_at` error cases described on those fields. Also returned for `entitlement_id`: this endpoint never accepts it, and rejects it with the message shown below rather than silently ignoring it. Also returned when `subscription_id` names a prepaid subscription: a wallet would double-meter a balance the subscription already tracks on its own, so the two are mutually exclusive. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |


# Read Wallet

Returns the wallet's balance and grant breakdown. For an active, metered wallet on a site with metered consumption enabled, this is the only operation that computes a live balance from meter usage; see the wallet resource's `metered_*` fields. It is also the only operation where `balance` can differ from `posted_balance`, and a grant's `remaining` from its `posted_remaining`: the pending usage is drawn from the grants in drain order, so the grants' `remaining` add up to `balance`. Readable regardless of status. Archiving a wallet doesn't affect reads.

```go
ReadWallet(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.WalletResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `int` | Template, Required | The Advanced Billing id of the wallet. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletResponse](../../doc/models/wallet-response.md).

## Example Usage

```go
ctx := context.Background()

id := 112

apiResponse, err := walletsController.ReadWallet(ctx, id)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.WalletErrorResponse:
            log.Fatalln("WalletErrorResponseException: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site. | `ApiError` |
| 404 | Not Found | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |


# Archive Wallet

Archives the wallet: blocks further grants, debits, adjustments, and purchases (each returns `422`), stops any metered drawdown, and stops recurring funding (its schedule is deactivated with reason `wallet_archived`, not removed), while reads and the ledger stay available. Nothing already on the wallet is destroyed or clawed back, but nothing on it can be spent either, by any path, once archived. Idempotent: archiving an already-archived wallet returns the same `200`. There is no restore endpoint.

```go
ArchiveWallet(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.WalletResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `int` | Template, Required | The Advanced Billing id of the wallet. |

## Response Type

**200**: The wallet, with `"status": "archived"`.

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletResponse](../../doc/models/wallet-response.md).

## Example Usage

```go
ctx := context.Background()

id := 112

apiResponse, err := walletsController.ArchiveWallet(ctx, id)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.WalletErrorResponse:
            log.Fatalln("WalletErrorResponseException: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site. | `ApiError` |
| 404 | Not Found | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |


# Update Wallet

Updates the wallet's editable attributes. Only `name` is editable; any other attribute in the body is ignored. Send only the attributes to change.

```go
UpdateWallet(
    ctx context.Context,
    id int,
    body *models.UpdateWalletRequest) (
    models.ApiResponse[models.WalletResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `int` | Template, Required | The Advanced Billing id of the wallet. |
| `body` | [`*models.UpdateWalletRequest`](../../doc/models/update-wallet-request.md) | Body, Optional | - |

## Response Type

**200**: The updated wallet. Reports the stored balance, equal to `posted_balance`, with the `metered_*` fields `null`, as on list, create, and archive.

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletResponse](../../doc/models/wallet-response.md).

## Example Usage

```go
ctx := context.Background()

id := 112

body := models.UpdateWalletRequest{
    Wallet:               models.Wallet2{
        Name:                 models.NewOptional(models.ToPointer("Onboarding credits")),
    },
}

apiResponse, err := walletsController.UpdateWallet(ctx, id, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.WalletErrorResponse:
            log.Fatalln("WalletErrorResponseException: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site. | `ApiError` |
| 404 | Not Found | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 422 | Validation failure. Includes a body with no editable attribute (`parameter_missing`), a `name` over 255 characters (`validation_failed`), and an archived wallet (`wallet_archived`). | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |

