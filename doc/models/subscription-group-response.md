
# Subscription Group Response

## Structure

`SubscriptionGroupResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroup` | [`models.SubscriptionGroup`](../../doc/models/subscription-group.md) | Required | - |

## Example (as JSON)

```json
{
  "subscription_group": {
    "payment_collection_method": "automatic",
    "customer_id": 220,
    "payment_profile": {
      "id": 44,
      "first_name": "first_name4",
      "last_name": "last_name2",
      "masked_card_number": "masked_card_number2"
    },
    "subscription_ids": [
      74,
      75
    ],
    "created_at": "2016-03-13T12:52:32.123Z"
  }
}
```

