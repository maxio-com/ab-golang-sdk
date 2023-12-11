
# List Sale Rep Item

## Structure

`ListSaleRepItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FullName` | `*string` | Optional | - |
| `SubscriptionsCount` | `*int` | Optional | - |
| `MrrData` | [`map[string]models.SaleRepItemMrr`](sale-rep-item-mrr.md) | Optional | - |
| `TestMode` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "mrr_data": {
    "november_2019": {
      "mrr": "$0.00",
      "usage": "$0.00",
      "recurring": "$0.00"
    },
    "december_2019": {
      "mrr": "$0.00",
      "usage": "$0.00",
      "recurring": "$0.00"
    },
    "january_2020": {
      "mrr": "$400.00",
      "usage": "$0.00",
      "recurring": "$400.00"
    }
  },
  "id": 26,
  "full_name": "full_name8",
  "subscriptions_count": 154,
  "test_mode": false
}
```

