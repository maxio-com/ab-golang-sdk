
# Sale Rep

## Structure

`SaleRep`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FullName` | `*string` | Optional | - |
| `SubscriptionsCount` | `*int` | Optional | - |
| `TestMode` | `*bool` | Optional | - |
| `Subscriptions` | [`[]models.SaleRepSubscription`](sale-rep-subscription.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": 18,
  "full_name": "full_name0",
  "subscriptions_count": 162,
  "test_mode": false,
  "subscriptions": [
    {
      "id": 202,
      "site_name": "site_name8",
      "subscription_url": "subscription_url2",
      "customer_name": "customer_name8",
      "created_at": "created_at4"
    },
    {
      "id": 202,
      "site_name": "site_name8",
      "subscription_url": "subscription_url2",
      "customer_name": "customer_name8",
      "created_at": "created_at4"
    }
  ]
}
```

