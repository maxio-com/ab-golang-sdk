
# Create Prepayment Response

## Structure

`CreatePrepaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayment` | [`models.CreatedPrepayment`](created-prepayment.md) | Required | - |

## Example (as JSON)

```json
{
  "prepayment": {
    "id": 38,
    "subscription_id": 148,
    "amount_in_cents": 124,
    "memo": "memo2",
    "created_at": "created_at6"
  }
}
```

