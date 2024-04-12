
# Subscription Group Item

## Structure

`SubscriptionGroupItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Reference` | `models.Optional[string]` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `ProductHandle` | `models.Optional[string]` | Optional | - |
| `ProductPricePointId` | `*int` | Optional | - |
| `ProductPricePointHandle` | `*string` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `CouponCode` | `models.Optional[string]` | Optional | - |
| `TotalRevenueInCents` | `*int64` | Optional | - |
| `BalanceInCents` | `*int64` | Optional | - |

## Example (as JSON)

```json
{
  "id": 16,
  "reference": "reference8",
  "product_id": 214,
  "product_handle": "product_handle4",
  "product_price_point_id": 138
}
```

