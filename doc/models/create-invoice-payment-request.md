
# Create Invoice Payment Request

## Structure

`CreateInvoicePaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Payment` | [`models.CreateInvoicePayment`](create-invoice-payment.md) | Required | - |
| `Type` | [`*models.InvoicePaymentTypeEnum`](invoice-payment-type-enum.md) | Optional | The type of payment to be applied to an Invoice.<br>**Default**: `"external"` |

## Example (as JSON)

```json
{
  "payment": {
    "method": "other",
    "amount": {
      "key1": "val1",
      "key2": "val2"
    },
    "memo": "memo0",
    "details": "details6"
  },
  "type": "external"
}
```

