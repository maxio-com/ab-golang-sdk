
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

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupSignupFailureData := models.SubscriptionGroupSignupFailureData{
        PayerId:                 models.ToPointer(236),
        PayerReference:          models.ToPointer("payer_reference2"),
        PaymentProfileId:        models.ToPointer(42),
        PaymentCollectionMethod: models.ToPointer("payment_collection_method4"),
        PayerAttributes:         models.ToPointer(models.PayerAttributes{
            FirstName:            models.ToPointer("first_name2"),
            LastName:             models.ToPointer("last_name0"),
            Email:                models.ToPointer("email4"),
            CcEmails:             models.ToPointer("cc_emails2"),
            Organization:         models.ToPointer("organization6"),
        }),
    }

}
```

