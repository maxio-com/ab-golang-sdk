
# List Subscription Groups Response

## Structure

`ListSubscriptionGroupsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroups` | [`[]models.ListSubscriptionGroupsItem`](../../doc/models/list-subscription-groups-item.md) | Optional | - |
| `Meta` | [`*models.ListSubscriptionGroupsMeta`](../../doc/models/list-subscription-groups-meta.md) | Optional | - |

## Example (as JSON)

```json
{
  "subscription_groups": [
    {
      "uid": "uid2",
      "scheme": 166,
      "customer_id": 186,
      "payment_profile_id": 162,
      "subscription_ids": [
        40
      ]
    },
    {
      "uid": "uid2",
      "scheme": 166,
      "customer_id": 186,
      "payment_profile_id": 162,
      "subscription_ids": [
        40
      ]
    },
    {
      "uid": "uid2",
      "scheme": 166,
      "customer_id": 186,
      "payment_profile_id": 162,
      "subscription_ids": [
        40
      ]
    }
  ],
  "meta": {
    "current_page": 126,
    "total_count": 150
  }
}
```

