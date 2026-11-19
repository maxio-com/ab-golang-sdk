# Wallet Purchases

```go
walletPurchasesController := client.WalletPurchasesController()
```

## Class Name

`WalletPurchasesController`

## Methods

* [Create Wallet Purchase](../../doc/controllers/wallet-purchases.md#create-wallet-purchase)
* [Preview Wallet Purchase](../../doc/controllers/wallet-purchases.md#preview-wallet-purchase)


# Create Wallet Purchase

Purchases prepaid units for the wallet and charges the subscription's payment method for them, using the wallet's `price_point_id` unless a different price point is named. A successful purchase writes a grant of origin `purchase`, with its expiration following the funding entitlement's grant expiration policy.

How "charged" works depends on the subscription's collection method. For card/automatic collection, the payment is attempted immediately; a decline writes nothing to the wallet, and the call fails with a `422` naming the gateway's decline reason. For invoice-billed (remittance) subscriptions, there's no immediate charge to decline: an invoice is drafted, the wallet is credited right away, and the invoice is issued afterward. If issuing the invoice fails at that point, the invoice is left in draft state and reported for follow-up, but the call still returns `201` and the units already credited are **not** rolled back. See [Preview Wallet Purchase](../../doc/controllers/wallet-purchases.md#preview-wallet-purchase) to quote a price before charging anything.

```go
CreateWalletPurchase(
    ctx context.Context,
    walletId int,
    idempotencyKey *string,
    body *models.CreateWalletPurchaseRequest) (
    models.ApiResponse[models.WalletPurchaseResultResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `walletId` | `int` | Template, Required | The Advanced Billing id of the wallet. |
| `idempotencyKey` | `*string` | Header, Optional | A key you generate, such as a UUID, that makes retrying this request safe. 1 to 255 printable ASCII characters; surrounding double quotes are ignored. Keys are scoped to your site and kept for at least 24 hours.<br><br>The first request with a key runs normally. A retry with the same key, method, path and query string, and a byte-identical body returns the stored status and body of that first request without running it again. The same key with a different request returns `422` (`idempotency_key_mismatch`). A retry that arrives while the first request is still running waits for it to finish, and returns `409` (`idempotency_request_in_progress`) only if that wait times out; on Create Wallet Purchase the `409` is immediate. Only successful responses are stored, so a failed request can be retried with the same key. When this header is sent, `uniqueness_token` is not checked.<br><br>See [Idempotent retries](http://localhost:8080/go) for the details, including how wallet purchases differ.<br><br>**Constraints**: *Minimum Length*: `1`, *Maximum Length*: `255` |
| `body` | [`*models.CreateWalletPurchaseRequest`](../../doc/models/create-wallet-purchase-request.md) | Body, Optional | Purchases prepaid units for the wallet at a price point, charging the subscription's payment method. |

## Response Type

**201**: Purchase completed. Charges the subscription's payment method and, once payment succeeds, credits the wallet.

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletPurchaseResultResponse](../../doc/models/wallet-purchase-result-response.md).

## Example Usage

```go
ctx := context.Background()

walletId := 216

body := models.CreateWalletPurchaseRequest{
    Purchase:             models.Purchase{
        Quantity:             models.PurchaseQuantityContainer.FromString("500.0"),
    },
}

apiResponse, err := walletPurchasesController.CreateWalletPurchase(ctx, walletId, nil, &body)
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
  "wallet_purchases": [
    {
      "id": 15,
      "price_point_id": 140,
      "invoice_uid": "abc123",
      "quantity": "500.0",
      "expires_at": null
    }
  ],
  "balance": "650.0"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | The `Idempotency-Key` header is empty, longer than 255 characters, or contains characters other than printable ASCII. Nothing is written. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 403 | Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site. | `ApiError` |
| 404 | Not Found | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 409 | Returned in two cases, both with nothing new charged or written. `idempotency_request_in_progress`: a purchase with the same `Idempotency-Key` is still running; retry with the same key once it finishes. `idempotency_key_outcome_unknown`: an earlier purchase with this key ended with a server error or an unexpected failure before its result was recorded, so it may or may not have charged the customer. Check the subscription's invoices before buying again under a new key. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 422 | Validation or payment failure, or an `Idempotency-Key` reused with a different request (`idempotency_key_mismatch`). Validation failures include a missing `purchase` root, a missing, malformed, non-positive, or out-of-range `quantity`, a `price_point_id` that doesn't belong to the wallet's billing product, an archived, custom, or otherwise unusable price point, no price point to purchase at (neither sent nor set on the wallet), a memo over 255 characters, an archived wallet, a price point with no price in the subscription's currency, a declined payment (message includes the gateway's own detail), an invoice that couldn't be created at all, a payment that settled but whose units could not be credited (`purchase_paid_but_uncredited`, whose message ends with the paid invoice's uid so it can be reconciled), a declined payment whose invoice could not be voided and is still open (`purchase_declined_invoice_open`), a purchase that failed before payment and left its invoice open (`purchase_invoice_left_open`), and a wallet whose billing product isn't a Wallet Product (only possible for a wallet created before Wallet Products became required).<br><br>`purchase_paid_but_uncredited`, `purchase_declined_invoice_open` and `purchase_invoice_left_open` bodies also carry `invoice_uid` and `invoice_number`. With an `Idempotency-Key`, these three responses are stored and replayed like a success. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 500 | `purchase_payment_outcome_unknown`: the payment was attempted but the gateway did not confirm whether the card was charged. Nothing is credited by this call; if the invoice turns out to be paid, the units are credited automatically (see [Purchasing Units](http://localhost:8080/go)). The body includes `invoice_uid` and `invoice_number` when an invoice was created; check that invoice before buying again. Retrying with the same `Idempotency-Key` returns `409` (`idempotency_key_outcome_unknown`) rather than charging again. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |


# Preview Wallet Purchase

Quotes the price for a purchase at the given quantity and price point, without charging the payment method or writing to the wallet. Takes the same parameters as [Create Wallet Purchase](../../doc/controllers/wallet-purchases.md#create-wallet-purchase).

```go
PreviewWalletPurchase(
    ctx context.Context,
    walletId int,
    body *models.CreateWalletPurchaseRequest) (
    models.ApiResponse[models.WalletPurchasePreviewResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `walletId` | `int` | Template, Required | The Advanced Billing id of the wallet. |
| `body` | [`*models.CreateWalletPurchaseRequest`](../../doc/models/create-wallet-purchase-request.md) | Body, Optional | Purchases prepaid units for the wallet at a price point, charging the subscription's payment method. |

## Response Type

**200**: A price quote. Charges nothing and writes nothing to the wallet.

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WalletPurchasePreviewResponse](../../doc/models/wallet-purchase-preview-response.md).

## Example Usage

```go
ctx := context.Background()

walletId := 216

body := models.CreateWalletPurchaseRequest{
    Purchase:             models.Purchase{
        Quantity:             models.PurchaseQuantityContainer.FromString("500.0"),
    },
}

apiResponse, err := walletPurchasesController.PreviewWalletPurchase(ctx, walletId, &body)
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
  "wallet_purchase_previews": [
    {
      "price_point_id": 140,
      "currency": "USD",
      "quantity": "500.0",
      "amount": "25.0"
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden. The response body is plain text (`403 Forbidden`), not JSON. Returned when the Wallets feature isn't enabled for the site. | `ApiError` |
| 404 | Not Found | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |
| 422 | Validation failure. Same cases as Create Wallet Purchase, minus the payment-outcome cases, since nothing is charged. | [`WalletErrorResponseException`](../../doc/models/wallet-error-response-exception.md) |

