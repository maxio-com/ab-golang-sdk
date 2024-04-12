
# Subscription Group Subscription Error

Object which contains subscription errors.

## Structure

`SubscriptionGroupSubscriptionError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Product` | `[]string` | Optional | - |
| `ProductPricePointId` | `[]string` | Optional | - |
| `PaymentProfile` | `[]string` | Optional | - |
| `PaymentProfileChargifyToken` | `[]string` | Optional | - |
| `Base` | `[]string` | Optional | - |
| `PaymentProfileExpirationMonth` | `[]string` | Optional | - |
| `PaymentProfileExpirationYear` | `[]string` | Optional | - |
| `PaymentProfileFullNumber` | `[]string` | Optional | - |

## Example (as JSON)

```json
{
  "product": [
    "product7",
    "product6"
  ],
  "product_price_point_id": [
    "product_price_point_id9",
    "product_price_point_id0"
  ],
  "payment_profile": [
    "payment_profile4",
    "payment_profile5"
  ],
  "payment_profile.chargify_token": [
    "payment_profile.chargify_token8",
    "payment_profile.chargify_token9"
  ],
  "base": [
    "base7",
    "base8",
    "base9"
  ]
}
```

