
# Remove Payment Event Data

Example schema for an `remove_payment` event

## Structure

`RemovePaymentEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionId` | `int` | Required | Transaction ID of the original payment that was removed |
| `Memo` | `string` | Required | Memo of the original payment |
| `OriginalAmount` | `*string` | Optional | Full amount of the original payment |
| `AppliedAmount` | `string` | Required | Applied amount of the original payment |
| `TransactionTime` | `time.Time` | Required | Transaction time of the original payment, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |
| `PaymentMethod` | [`models.InvoiceEventPayment2`](../../doc/models/containers/invoice-event-payment-2.md) | Required | A nested data structure detailing the method of payment |
| `Prepayment` | `bool` | Required | The flag that shows whether the original payment was a prepayment or not |

## Example (as JSON)

```json
{
  "transaction_id": 180,
  "memo": "memo0",
  "original_amount": "original_amount0",
  "applied_amount": "applied_amount2",
  "transaction_time": "2016-03-13T12:52:32.123Z",
  "payment_method": {
    "type": "apple_pay"
  },
  "prepayment": false
}
```

