
# Subscription Group

## Structure

`SubscriptionGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomerId` | `*int` | Optional | - |
| `PaymentProfile` | [`*models.SubscriptionGroupPaymentProfile`](../../doc/models/subscription-group-payment-profile.md) | Optional | - |
| `PaymentCollectionMethod` | `*string` | Optional | - |
| `SubscriptionIds` | `[]int` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |

## Example (as JSON)

```json
{
  "customer_id": 36,
  "payment_profile": {
    "id": 44,
    "first_name": "first_name4",
    "last_name": "last_name2",
    "masked_card_number": "masked_card_number2"
  },
  "payment_collection_method": "payment_collection_method8",
  "subscription_ids": [
    146,
    147,
    148
  ],
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

