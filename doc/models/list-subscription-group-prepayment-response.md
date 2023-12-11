
# List Subscription Group Prepayment Response

## Structure

`ListSubscriptionGroupPrepaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayments` | [`[]models.ListSubscriptionGroupPrepayment`](list-subscription-group-prepayment.md) | Required | - |

## Example (as JSON)

```json
{
  "prepayments": [
    {
      "prepayment": {
        "id": 38,
        "subscription_group_uid": "subscription_group_uid2",
        "amount_in_cents": 124,
        "remaining_amount_in_cents": 182,
        "details": "details8"
      }
    }
  ]
}
```

