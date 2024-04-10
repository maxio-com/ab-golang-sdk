
# Subscription Group Signup Error Response Exception

## Structure

`SubscriptionGroupSignupErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.SubscriptionGroupSignupError`](../../doc/models/subscription-group-signup-error.md) | Required | - |

## Example (as JSON)

```json
{
  "errors": {
    "subscriptions": {
      "0": {
        "payment_profile.chargify_token": [
          "Chargify token not found"
        ],
        "product": [
          "Product must be among the Products for this Site"
        ],
        "product_price_point_id": [
          "Product price point must belong to product."
        ],
        "base": [
          "No payment method was on file for the $30.00 balance'"
        ],
        "payment_profile.expiration_month": [
          "Credit card expiration month: cannot be blank."
        ],
        "payment_profile.expiration_year": [
          "Credit card expiration year: cannot be blank."
        ],
        "payment_profile.full_number": [
          "Credit card number: cannot be blank."
        ],
        "payment_profile": [
          "payment_profile2"
        ]
      }
    },
    "payer_reference": "payer_reference4",
    "payer": {
      "last_name": [
        "last_name5",
        "last_name6"
      ],
      "first_name": [
        "first_name8"
      ],
      "email": [
        "email0",
        "email9"
      ]
    },
    "subscription_group": [
      "subscription_group7",
      "subscription_group8",
      "subscription_group9"
    ],
    "payment_profile_id": "payment_profile_id8"
  }
}
```

