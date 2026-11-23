# Wallet Transactions

```go
walletTransactionsController := client.WalletTransactionsController()
```

## Class Name

`WalletTransactionsController`

## Methods

* [List Wallet Transactions](../../doc/controllers/wallet-transactions.md#list-wallet-transactions)
* [Create Wallet Transaction](../../doc/controllers/wallet-transactions.md#create-wallet-transaction)


# List Wallet Transactions

Lists the wallet's append-only ledger, newest first, paginated.

```go
ListWalletTransactions(
    ctx context.Context,
    walletId int,
    page *int,
    perPage *int) (
    models.ApiResponse[models.WalletTransactionsListResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `walletId` | `int` | Template, Required | The Advanced Billing id of the wallet. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletTransactionsListResponse](../../doc/models/wallet-transactions-list-response.md).

## Example Usage

```go
ctx := context.Background()

walletId := 216

page := 1

perPage := 50

apiResponse, err := walletTransactionsController.ListWalletTransactions(ctx, walletId, &page, &perPage)
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
  "wallet_transactions": [
    {
      "id": 41,
      "wallet_id": 5,
      "kind": "debit",
      "amount": "-30.0",
      "grant_id": 9,
      "source_transaction_id": null,
      "source_corrections_total": "0.0",
      "balance_after": "0.0",
      "metered_from_at": null,
      "metered_through_at": null,
      "conversion_rate_applied": null,
      "invoice_uid": null,
      "memo": null,
      "created_at": "2026-08-01T09:30:00Z"
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
| 404 | Not Found | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |


# Create Wallet Transaction

Grants units to the wallet, or debits units from it, depending on `kind`. `amount` is always sent positive, and `kind` determines the direction written to the ledger.

**`kind: grant`**: a one-time manual top-up. On a non-negative balance, this creates a grant and writes one ledger row. On a negative balance, it repays the overdraft first (a `kind: grant` row with `grant_id: null`, for the repaid amount), then creates the grant: its `original` is always the full `amount` sent, and its `remaining` is what's left after the overdraft was repaid, which is `0` if the overdraft absorbed the whole top-up. The grant row is still written, and its `amount` is that `remaining`, not `original`, so the ledger rows sum to the balance change. A grant may carry `expires_at`; a debit may not.

**`kind: debit`**: direct consumption. Drains grants earliest-expiration-first, one `kind: debit` row per grant touched, plus any overdue expiration rows this same call triggers first (see below):

- Under the wallet's default `hard_stop_overflow: discard`, the drain is capped at the available balance. For a metered wallet, that cap is the live metered balance, not the stored one, so a debit can drain less than its stored balance would suggest, or even drain nothing at all (`201` with `drained_amount: "0.0"`) if live usage has already consumed what the ledger hadn't caught up to yet. A debit against a wallet whose *stored* balance is already zero or negative (`discard` only ever reaches zero) returns `422 "Wallet is depleted"`. Before that failure is returned, this call still expires any grants that were already overdue, so a `422` response can still leave new `kind: expiration` rows written.
- Under `hard_stop_overflow: allow_negative`, the full amount is deducted. The excess past the wallet's grants is booked as one overdraft row with `grant_id: null`, `rejected_amount` is always `"0.0"`, and the balance goes negative.

```go
CreateWalletTransaction(
    ctx context.Context,
    walletId int,
    idempotencyKey *string,
    body *models.CreateWalletTransactionRequest) (
    models.ApiResponse[models.WalletTransactionResultResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `walletId` | `int` | Template, Required | The Advanced Billing id of the wallet. |
| `idempotencyKey` | `*string` | Header, Optional | A key you generate, such as a UUID, that makes retrying this request safe. 1 to 255 printable ASCII characters; surrounding double quotes are ignored. Keys are scoped to your site and kept for at least 24 hours.<br><br>The first request with a key runs normally. A retry with the same key, method, path and query string, and a byte-identical body returns the stored status and body of that first request without running it again. The same key with a different request returns `422` (`idempotency_key_mismatch`). A retry that arrives while the first request is still running waits for it to finish, and returns `409` (`idempotency_request_in_progress`) only if that wait times out; on Create Wallet Purchase the `409` is immediate. Only successful responses are stored, so a failed request can be retried with the same key. When this header is sent, `uniqueness_token` is not checked.<br><br>See [Idempotent retries](http://localhost:8080/go) for the details, including how wallet purchases differ.<br><br>**Constraints**: *Minimum Length*: `1`, *Maximum Length*: `255` |
| `body` | [`*models.CreateWalletTransactionRequest`](../../doc/models/create-wallet-transaction-request.md) | Body, Optional | - |

## Response Type

**201**: Transaction applied. For `kind: grant`, the ledger row(s) written plus the new balance. For `kind: debit`, one row per grant touched plus the requested/drained/rejected accounting and the new balance.

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletTransactionResultResponse](../../doc/models/wallet-transaction-result-response.md).

## Example Usage

```go
ctx := context.Background()

walletId := 216

body := models.CreateWalletTransactionRequest{
    WalletTransaction:    models.WalletTransaction1{
        Kind:                 models.Kind1_GRANT,
        Amount:               models.WalletTransaction1AmountContainer.FromString("100.0"),
        Memo:                 models.NewOptional(models.ToPointer("Manual top-up")),
    },
}

apiResponse, err := walletTransactionsController.CreateWalletTransaction(ctx, walletId, nil, &body)
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
  "wallet_transactions": [
    {
      "id": 42,
      "wallet_id": 5,
      "kind": "grant",
      "amount": "100.0",
      "grant_id": 13,
      "source_transaction_id": null,
      "source_corrections_total": null,
      "balance_after": "250.0",
      "metered_from_at": null,
      "metered_through_at": null,
      "conversion_rate_applied": null,
      "invoice_uid": null,
      "memo": "Manual top-up",
      "created_at": "2026-08-02T09:00:00Z"
    }
  ],
  "balance": "250.0"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | The `Idempotency-Key` header is empty, longer than 255 characters, or contains characters other than printable ASCII. Nothing is written. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 403 | Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site. | `ApiError` |
| 404 | Not Found | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 409 | A request with the same `Idempotency-Key` was still being processed when this retry's wait for it timed out (`idempotency_request_in_progress`). Nothing is written; retry with the same key once it finishes. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 422 | Validation failure, or an `Idempotency-Key` reused with a different request (`idempotency_key_mismatch`). Validation failures include an unrecognized `kind`, a missing `wallet_transaction` root, a malformed, non-positive, or out-of-range `amount`, a memo over 255 characters, `expires_at` sent with `kind: debit`, a malformed or past-dated `expires_at`, an archived wallet, and, for `kind: debit` only, a depleted wallet. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 500 | For a metered wallet, usage could be read but settling it against the ledger failed. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 503 | For a metered wallet, live usage could not be read in time to compute the debit. Not returned for `kind: grant`. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |

