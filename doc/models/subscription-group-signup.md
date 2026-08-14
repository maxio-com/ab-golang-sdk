
# Subscription Group Signup

## Structure

`SubscriptionGroupSignup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfileId` | `*int` | Optional | - |
| `PayerId` | `*int` | Optional | - |
| `PayerReference` | `*string` | Optional | - |
| `PaymentCollectionMethod` | [`*models.CollectionMethod`](../../doc/models/collection-method.md) | Optional | The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`. |
| `PayerAttributes` | [`*models.PayerAttributes`](../../doc/models/payer-attributes.md) | Optional | - |
| `CreditCardAttributes` | [`*models.SubscriptionGroupCreditCard`](../../doc/models/subscription-group-credit-card.md) | Optional | - |
| `BankAccountAttributes` | [`*models.SubscriptionGroupBankAccount`](../../doc/models/subscription-group-bank-account.md) | Optional | - |
| `Subscriptions` | [`[]models.SubscriptionGroupSignupItem`](../../doc/models/subscription-group-signup-item.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupSignup := models.SubscriptionGroupSignup{
        PaymentProfileId:        models.ToPointer(124),
        PayerId:                 models.ToPointer(146),
        PayerReference:          models.ToPointer("payer_reference0"),
        PaymentCollectionMethod: models.ToPointer(models.CollectionMethod_PREPAID),
        PayerAttributes:         models.ToPointer(models.PayerAttributes{
            FirstName:            models.ToPointer("first_name2"),
            LastName:             models.ToPointer("last_name0"),
            Email:                models.ToPointer("email4"),
            CcEmails:             models.ToPointer("cc_emails2"),
            Organization:         models.ToPointer("organization6"),
        }),
        Subscriptions:           []models.SubscriptionGroupSignupItem{
            models.SubscriptionGroupSignupItem{
                ProductHandle:           models.ToPointer("product_handle8"),
                ProductId:               models.ToPointer(144),
                ProductPricePointId:     models.ToPointer(68),
                ProductPricePointHandle: models.ToPointer("product_price_point_handle4"),
                OfferId:                 models.ToPointer(40),
                Metafields:              map[string]string{
                    "custom_field_name_1": "custom_field_value_1",
                    "custom_field_name_2": "custom_field_value_2",
                },
            },
        },
    }

}
```

