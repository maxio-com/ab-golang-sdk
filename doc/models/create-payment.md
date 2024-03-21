
# Create Payment

## Structure

`CreatePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `string` | Required | - |
| `Memo` | `string` | Required | - |
| `PaymentDetails` | `string` | Required | - |
| `PaymentMethod` | [`models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Required | The type of payment method used. Defaults to other. |

## Example (as JSON)

```json
{
  "amount": "amount6",
  "memo": "memo8",
  "payment_details": "payment_details4",
  "payment_method": "cash"
}
```

