
# Invoice Refund

## Structure

`InvoiceRefund`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionId` | `*int` | Optional | - |
| `PaymentId` | `*int` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `OriginalAmount` | `*string` | Optional | - |
| `AppliedAmount` | `*string` | Optional | - |
| `GatewayTransactionId` | `Optional[string]` | Optional | The transaction ID for the refund as returned from the payment gateway |

## Example (as JSON)

```json
{
  "transaction_id": 172,
  "payment_id": 42,
  "memo": "memo6",
  "original_amount": "original_amount6",
  "applied_amount": "applied_amount6"
}
```

