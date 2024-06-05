
# Failed Payment Event Data

Example schema for an `failed_payment` event

## Structure

`FailedPaymentEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmountInCents` | `int` | Required | The monetary value of the payment, expressed in cents. |
| `AppliedAmount` | `int` | Required | The monetary value of the payment, expressed in dollars. |
| `Memo` | `models.Optional[string]` | Optional | The memo passed when the payment was created. |
| `PaymentMethod` | [`models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Required | - |
| `TransactionId` | `int` | Required | The transaction ID of the failed payment. |

## Example (as JSON)

```json
{
  "amount_in_cents": 128,
  "applied_amount": 154,
  "memo": "memo2",
  "payment_method": "credit_card",
  "transaction_id": 170
}
```

