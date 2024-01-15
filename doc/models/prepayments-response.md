
# Prepayments Response

## Structure

`PrepaymentsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayments` | [`[]models.Prepayment`](../../doc/models/prepayment.md) | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "prepayments": [
    {
      "id": 76,
      "subscription_id": 186,
      "amount_in_cents": 94,
      "remaining_amount_in_cents": 220,
      "refunded_amount_in_cents": 170,
      "details": "details6",
      "external": false,
      "memo": "memo0",
      "payment_type": "cash",
      "created_at": "created_at4"
    },
    {
      "id": 76,
      "subscription_id": 186,
      "amount_in_cents": 94,
      "remaining_amount_in_cents": 220,
      "refunded_amount_in_cents": 170,
      "details": "details6",
      "external": false,
      "memo": "memo0",
      "payment_type": "cash",
      "created_at": "created_at4"
    }
  ]
}
```

