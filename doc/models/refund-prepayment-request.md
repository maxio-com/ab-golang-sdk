
# Refund Prepayment Request

## Structure

`RefundPrepaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Refund` | [`models.RefundPrepayment`](../../doc/models/refund-prepayment.md) | Required | - |

## Example (as JSON)

```json
{
  "refund": {
    "amount_in_cents": 132,
    "amount": {
      "key1": "val1",
      "key2": "val2"
    },
    "memo": "memo2",
    "external": false
  }
}
```

