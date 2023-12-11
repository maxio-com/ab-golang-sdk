
# Subscription Group Signup Request

## Structure

`SubscriptionGroupSignupRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroup` | [`models.SubscriptionGroupSignup`](subscription-group-signup.md) | Required | - |

## Example (as JSON)

```json
{
  "subscription_group": {
    "payment_collection_method": "automatic",
    "subscriptions": [
      {
        "metafields": {
          "custom_field_name_1": "custom_field_value_1",
          "custom_field_name_2": "custom_field_value_2"
        },
        "product_handle": "product_handle8",
        "product_id": 144,
        "product_price_point_id": 68,
        "product_price_point_handle": "product_price_point_handle4",
        "offer_id": 40
      }
    ],
    "payment_profile_id": 128,
    "payer_id": 150,
    "payer_reference": "payer_reference6",
    "payer_attributes": {
      "first_name": "first_name2",
      "last_name": "last_name0",
      "email": "email4",
      "cc_emails": "cc_emails2",
      "organization": "organization6"
    }
  }
}
```

