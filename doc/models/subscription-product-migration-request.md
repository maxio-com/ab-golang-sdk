
# Subscription Product Migration Request

## Structure

`SubscriptionProductMigrationRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Migration` | [`models.SubscriptionProductMigration`](../../doc/models/subscription-product-migration.md) | Required | - |

## Example (as JSON)

```json
{
  "migration": {
    "include_trial": false,
    "include_initial_charge": false,
    "include_coupons": true,
    "preserve_period": false,
    "product_id": 158,
    "product_price_point_id": 82
  }
}
```

