
# Refund Invoice Request

## Structure

`RefundInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Refund` | [`models.RefundInvoiceRequestRefund`](../../doc/models/containers/refund-invoice-request-refund.md) | Required | This is a container for any-of cases. |

## Example (as JSON)

```json
{
  "refund": {
    "amount": "amount8",
    "memo": "memo0",
    "payment_id": 0,
    "external": false,
    "apply_credit": false,
    "void_invoice": false
  }
}
```

