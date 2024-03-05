
# Create Invoice Payment

## Structure

`CreateInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `*interface{}` | Optional | A string of the dollar amount to be refunded (eg. "10.50" => $10.50) |
| `Memo` | `*string` | Optional | A description to be attached to the payment. |
| `Method` | [`*models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Optional | The type of payment method used. Defaults to other. |
| `Details` | `*string` | Optional | Additional information related to the payment method (eg. Check #) |
| `PaymentProfileId` | `*int` | Optional | The ID of the payment profile to be used for the payment. |

## Example (as JSON)

```json
{
  "amount": {
    "key1": "val1",
    "key2": "val2"
  },
  "memo": "memo0",
  "method": "cash",
  "details": "details6",
  "payment_profile_id": 122
}
```

