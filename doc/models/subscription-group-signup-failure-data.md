
# Subscription Group Signup Failure Data

## Structure

`SubscriptionGroupSignupFailureData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PayerId` | `*int` | Optional | - |
| `PayerReference` | `*string` | Optional | - |
| `PaymentProfileId` | `*int` | Optional | - |
| `PaymentCollectionMethod` | `*string` | Optional | - |
| `PayerAttributes` | [`*models.PayerAttributes`](../../doc/models/payer-attributes.md) | Optional | - |
| `CreditCardAttributes` | [`*models.SubscriptionGroupCreditCard`](../../doc/models/subscription-group-credit-card.md) | Optional | - |
| `BankAccountAttributes` | [`*models.SubscriptionGroupBankAccount`](../../doc/models/subscription-group-bank-account.md) | Optional | - |
| `Subscriptions` | [`[]models.SubscriptionGroupSignupItem`](../../doc/models/subscription-group-signup-item.md) | Optional | - |

## Example (as JSON)

```json
{
  "payer_id": 16,
  "payer_reference": "payer_reference8",
  "payment_profile_id": 6,
  "payment_collection_method": "payment_collection_method0",
  "payer_attributes": {
    "first_name": "first_name2",
    "last_name": "last_name0",
    "email": "email4",
    "cc_emails": "cc_emails2",
    "organization": "organization6"
  }
}
```

