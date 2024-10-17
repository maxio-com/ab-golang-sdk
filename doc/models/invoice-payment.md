
# Invoice Payment

## Structure

`InvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionTime` | `*time.Time` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `OriginalAmount` | `*string` | Optional | - |
| `AppliedAmount` | `*string` | Optional | - |
| `PaymentMethod` | [`*models.InvoicePaymentMethod`](../../doc/models/invoice-payment-method.md) | Optional | - |
| `TransactionId` | `*int` | Optional | - |
| `Prepayment` | `*bool` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |
| `GatewayUsed` | `*string` | Optional | - |
| `GatewayTransactionId` | `models.Optional[string]` | Optional | The transaction ID for the payment as returned from the payment gateway |
| `ReceivedOn` | `models.Optional[time.Time]` | Optional | Date reflecting when the payment was received from a customer. Must be in the past. Applicable only to<br>`external` payments. |
| `Uid` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "transaction_time": "2016-03-13T12:52:32.123Z",
  "memo": "memo6",
  "original_amount": "original_amount6",
  "applied_amount": "applied_amount6",
  "payment_method": {
    "details": "details0",
    "kind": "kind8",
    "memo": "memo4",
    "type": "type0",
    "card_brand": "card_brand6"
  }
}
```

