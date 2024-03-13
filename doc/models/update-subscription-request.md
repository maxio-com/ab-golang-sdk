
# Update Subscription Request

## Structure

`UpdateSubscriptionRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`models.UpdateSubscription`](../../doc/models/update-subscription.md) | Required | - |

## Example (as JSON)

```json
{
  "subscription": {
    "dunning_communication_delay_time_zone": "\"Eastern Time (US & Canada)\"",
    "credit_card_attributes": {
      "full_number": "full_number2",
      "expiration_month": "expiration_month6",
      "expiration_year": "expiration_year2"
    },
    "product_handle": "product_handle6",
    "product_id": 206,
    "product_change_delayed": false,
    "next_product_id": "next_product_id6"
  }
}
```

