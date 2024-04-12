
# Create Multi Invoice Payment

## Structure

`CreateMultiInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Memo` | `*string` | Optional | A description to be attached to the payment. |
| `Details` | `*string` | Optional | Additional information related to the payment method (eg. Check #). |
| `Method` | [`*models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Optional | The type of payment method used. Defaults to other. |
| `Amount` | [`models.CreateMultiInvoicePaymentAmount`](../../doc/models/containers/create-multi-invoice-payment-amount.md) | Required | This is a container for one-of cases. |
| `ReceivedOn` | `*string` | Optional | Date reflecting when the payment was received from a customer. Must be in the past. |
| `Applications` | [`[]models.CreateInvoicePaymentApplication`](../../doc/models/create-invoice-payment-application.md) | Required | - |

## Example (as JSON)

```json
{
  "amount": "String7",
  "applications": [
    {
      "invoice_uid": "invoice_uid8",
      "amount": "amount0"
    }
  ],
  "memo": "memo8",
  "details": "details4",
  "method": "credit_card",
  "received_on": "received_on6"
}
```

