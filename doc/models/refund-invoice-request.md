
# Refund Invoice Request

## Structure

`RefundInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Refund` | [`models.Refund`](../../doc/models/refund.md) | Required | - |

## Example (as JSON)

```json
{
  "refund": {
    "amount": "amount0",
    "memo": "memo2",
    "payment_id": 44,
    "external": false,
    "apply_credit": false
  }
}
```

