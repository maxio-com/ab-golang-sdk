
# Create Invoice Payment

## Structure

`CreateInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `*interface{}` | Optional | A string of the dollar amount to be refunded (eg. "10.50" => $10.50) |
| `Memo` | `*string` | Optional | A description to be attached to the payment. |
| `Method` | [`*models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Optional | The type of payment method used.<br>**Default**: `"other"` |
| `Details` | `*string` | Optional | Additional information related to the payment method (eg. Check #) |

## Example (as JSON)

```json
{
  "method": "other",
  "amount": {
    "key1": "val1",
    "key2": "val2"
  },
  "memo": "memo0",
  "details": "details6"
}
```

