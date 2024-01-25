
# Create Multi Invoice Payment Request

## Structure

`CreateMultiInvoicePaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Payment` | [`models.CreateMultiInvoicePayment`](../../doc/models/create-multi-invoice-payment.md) | Required | - |

## Example (as JSON)

```json
{
  "payment": {
    "method": "other",
    "amount": {
      "key1": "val1",
      "key2": "val2"
    },
    "applications": [
      {
        "invoice_uid": "invoice_uid8",
        "amount": "amount0"
      }
    ],
    "memo": "memo0",
    "details": "details6",
    "received_on": "received_on8"
  }
}
```

