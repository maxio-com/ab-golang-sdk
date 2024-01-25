
# Create Multi Invoice Payment

## Structure

`CreateMultiInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Memo` | `*string` | Optional | A description to be attached to the payment. |
| `Details` | `*string` | Optional | Additional information related to the payment method (eg. Check #). |
| `Method` | [`*models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Optional | The type of payment method used.<br>**Default**: `"other"` |
| `Amount` | `interface{}` | Required | Dollar amount of the sum of the invoices payment (eg. "10.50" => $10.50). |
| `ReceivedOn` | `*string` | Optional | Date reflecting when the payment was received from a customer. Must be in the past. |
| `Applications` | [`[]models.CreateInvoicePaymentApplication`](../../doc/models/create-invoice-payment-application.md) | Required | - |

## Example (as JSON)

```json
{
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
  "memo": "memo8",
  "details": "details4",
  "received_on": "received_on6"
}
```

