# Wallet Adjustments

```go
walletAdjustmentsController := client.WalletAdjustmentsController()
```

## Class Name

`WalletAdjustmentsController`


# Create Wallet Adjustment

Posts a signed correction against a specific grant (or, for an overdraft entry, against the overdraft): positive adds units to the grant, negative removes them. `reason` is required and is written to the ledger row's `memo`. Writes exactly one ledger row, of kind `adjustment`.

Use `source_transaction_id` to name a specific `debit` row when you're correcting one recorded consumption rather than the grant's balance as a whole. When it names a grant-attributed debit, it's recorded for audit only: naming it doesn't change how `amount` is applied, and more than one adjustment may name the same row. When it names an overdraft entry, it switches the correction to the overdraft path described below.

An overdraft entry (a `debit` row with no `grant_id`, written when an `allow_negative` wallet consumes past its grants) can be corrected without a `grant_id`. That correction is credit-only, credits no grant, and is bounded by the usage the entry still records and by the overdraft the wallet still owes. Once a later grant has repaid the overdraft, correct that grant's balance instead.

An adjustment corrects the ledger, so it works on the grant's recorded remainder, `posted_remaining`, not on `remaining`, which on a metered wallet read through [Read Wallet](../../doc/controllers/wallets.md#read-wallet) also subtracts usage not yet posted. Send `posted_remaining` back as `expected_remaining`; sending `remaining` on a metered wallet with pending usage returns `409`.

`expected_remaining` and `expected_source_corrections` are optional optimistic-concurrency checks; see their field descriptions. Use them when the correction's `amount` was derived from a value read earlier that might have moved since.

```go
CreateWalletAdjustment(
    ctx context.Context,
    walletId int,
    idempotencyKey *string,
    body *models.CreateWalletAdjustmentRequest) (
    models.ApiResponse[models.WalletGrantResultResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `walletId` | `int` | Template, Required | The Advanced Billing id of the wallet. |
| `idempotencyKey` | `*string` | Header, Optional | A key you generate, such as a UUID, that makes retrying this request safe. 1 to 255 printable ASCII characters; surrounding double quotes are ignored. Keys are scoped to your site and kept for at least 24 hours.<br><br>The first request with a key runs normally. A retry with the same key, method, path and query string, and a byte-identical body returns the stored status and body of that first request without running it again. The same key with a different request returns `422` (`idempotency_key_mismatch`). A retry that arrives while the first request is still running waits for it to finish, and returns `409` (`idempotency_request_in_progress`) only if that wait times out; on Create Wallet Purchase the `409` is immediate. Only successful responses are stored, so a failed request can be retried with the same key. When this header is sent, `uniqueness_token` is not checked.<br><br>See [Idempotent retries](http://localhost:8080/go) for the details, including how wallet purchases differ.<br><br>**Constraints**: *Minimum Length*: `1`, *Maximum Length*: `255` |
| `body` | [`*models.CreateWalletAdjustmentRequest`](../../doc/models/create-wallet-adjustment-request.md) | Body, Optional | A signed correction to a specific grant's balance, or to an overdraft entry's usage, intended for support and ops use. |

## Response Type

**201**: The ledger row written, plus the wallet's new balance.

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletGrantResultResponse](../../doc/models/wallet-grant-result-response.md).

## Example Usage

```go
ctx := context.Background()

walletId := 216

body := models.CreateWalletAdjustmentRequest{
    Adjustment:           models.Adjustment{
        Amount:                    models.AdjustmentAmountContainer.FromString("10.0"),
        Reason:                    "Support credit for outage on 2026-07-30",
        GrantId:                   models.NewOptional(models.ToPointer(9)),
    },
}

apiResponse, err := walletAdjustmentsController.CreateWalletAdjustment(ctx, walletId, nil, &body)
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
      "id": 44,
      "wallet_id": 5,
      "kind": "adjustment",
      "amount": "10.0",
      "grant_id": 9,
      "source_transaction_id": null,
      "source_corrections_total": null,
      "balance_after": "60.0",
      "metered_from_at": null,
      "metered_through_at": null,
      "conversion_rate_applied": null,
      "invoice_uid": null,
      "memo": "Support credit for outage on 2026-07-30",
      "created_at": "2026-08-02T09:10:00Z"
    }
  ],
  "balance": "60.0"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | The `Idempotency-Key` header is empty, longer than 255 characters, or contains characters other than printable ASCII. Nothing is written. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 403 | Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site. | `ApiError` |
| 404 | Not Found | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 409 | Returned in two cases. A request with the same `Idempotency-Key` was still being processed when this retry's wait for it timed out (`idempotency_request_in_progress`); retry with the same key once it finishes. Or the state the `amount` was derived from has moved: either the grant's `posted_remaining` no longer matches `expected_remaining`, or the corrections already posted against `source_transaction_id` no longer sum to `expected_source_corrections`. Nothing is written. Reload and re-derive the amount. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 422 | Validation failure, or an `Idempotency-Key` reused with a different request (`idempotency_key_mismatch`). Validation failures include a missing `reason`, a `reason` over 255 characters, a missing or not-found `grant_id`, a zero, malformed, or out-of-range `amount`, a negative adjustment exceeding the grant's `posted_remaining`, a positive adjustment exceeding the grant's `original`, a malformed `expected_remaining` or `expected_source_corrections` (including an empty string, which is treated as malformed rather than omitted), a `source_transaction_id` not on this wallet or not a `debit` row, an overdraft correction that debits instead of credits ("Overdraft usage can only be corrected downward"), that exceeds the usage the entry still records ("Adjustment exceeds the usage this entry still records"), or that exceeds the overdraft the wallet still owes ("Adjustment exceeds the wallet's outstanding overdraft", followed by how much can still be corrected), a `grant_id` sent with an overdraft entry, and an archived wallet. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |

