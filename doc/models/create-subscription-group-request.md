
# Create Subscription Group Request

## Structure

`CreateSubscriptionGroupRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroup` | [`models.CreateSubscriptionGroup`](../../doc/models/create-subscription-group.md) | Required | - |

## Example (as JSON)

```json
{
  "subscription_group": {
    "subscription_id": {
      "key1": "val1",
      "key2": "val2"
    },
    "member_ids": [
      164,
      165
    ]
  }
}
```

