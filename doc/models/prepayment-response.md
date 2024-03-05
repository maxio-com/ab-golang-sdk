
# Prepayment Response

## Structure

`PrepaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayment` | [`models.Prepayment`](../../doc/models/prepayment.md) | Required | - |

## Example (as JSON)

```json
{
  "prepayment": {
    "id": 38,
    "subscription_id": 148,
    "amount_in_cents": 124,
    "remaining_amount_in_cents": 182,
    "refunded_amount_in_cents": 132,
    "details": "details8",
    "external": false,
    "memo": "memo2",
    "payment_type": "credit_card",
    "created_at": "2016-03-13T12:52:32.123Z"
  }
}
```

