
# Create Prepayment Response

## Structure

`CreatePrepaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayment` | [`models.CreatedPrepayment`](../../doc/models/created-prepayment.md) | Required | - |

## Example (as JSON)

```json
{
  "prepayment": {
    "id": 38,
    "subscription_id": 148,
    "amount_in_cents": 124,
    "memo": "memo2",
    "created_at": "2016-03-13T12:52:32.123Z"
  }
}
```

