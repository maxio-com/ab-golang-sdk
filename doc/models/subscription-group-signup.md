
# Subscription Group Signup

## Structure

`SubscriptionGroupSignup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfileId` | `*int` | Optional | - |
| `PayerId` | `*int` | Optional | - |
| `PayerReference` | `*string` | Optional | - |
| `PaymentCollectionMethod` | [`*models.PaymentCollectionMethodEnum`](payment-collection-method-enum.md) | Optional | The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.<br>**Default**: `"automatic"` |
| `PayerAttributes` | [`*models.PayerAttributes`](payer-attributes.md) | Optional | - |
| `CreditCardAttributes` | [`*models.SubscriptionGroupCreditCard`](subscription-group-credit-card.md) | Optional | - |
| `BankAccountAttributes` | [`*models.SubscriptionGroupBankAccount`](subscription-group-bank-account.md) | Optional | - |
| `Subscriptions` | [`[]models.SubscriptionGroupSignupItem`](subscription-group-signup-item.md) | Required | - |

## Example (as JSON)

```json
{
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
  "payment_profile_id": 42,
  "payer_id": 64,
  "payer_reference": "payer_reference8",
  "payer_attributes": {
    "first_name": "first_name2",
    "last_name": "last_name0",
    "email": "email4",
    "cc_emails": "cc_emails2",
    "organization": "organization6"
  }
}
```

