
# Payment Method External

## Structure

`PaymentMethodExternal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | `*string` | Required | - |
| `Kind` | `string` | Required | - |
| `Memo` | `*string` | Required | - |
| `Type` | [`models.InvoiceEventPaymentMethod`](../../doc/models/invoice-event-payment-method.md) | Required | - |

## Example (as JSON)

```json
{
  "details": "details4",
  "kind": "kind2",
  "memo": "memo8",
  "type": "external"
}
```

