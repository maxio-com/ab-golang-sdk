
# Create Subscription Request

## Structure

`CreateSubscriptionRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`models.CreateSubscription`](../../doc/models/create-subscription.md) | Required | - |

## Example (as JSON)

```json
{
  "subscription": {
    "payment_collection_method": "automatic",
    "metafields": {
      "custom_field_name_1": "custom_field_value_1",
      "custom_field_name_2": "custom_field_value_2"
    },
    "dunning_communication_delay_enabled": false,
    "dunning_communication_delay_time_zone": "\"Eastern Time (US & Canada)\"",
    "skip_billing_manifest_taxes": false,
    "product_handle": "product_handle6",
    "product_id": 206,
    "product_price_point_handle": "product_price_point_handle2",
    "product_price_point_id": 130,
    "custom_price": {
      "name": "name4",
      "handle": "handle0",
      "price_in_cents": "String3",
      "interval": "String3",
      "interval_unit": "day",
      "trial_price_in_cents": "String3",
      "trial_interval": "String5",
      "trial_interval_unit": "day"
    }
  }
}
```

