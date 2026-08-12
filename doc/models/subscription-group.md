
# Subscription Group

## Structure

`SubscriptionGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `PaymentProfile` | [`*models.SubscriptionGroupPaymentProfile`](../../doc/models/subscription-group-payment-profile.md) | Optional | - |
| `PaymentCollectionMethod` | [`*models.CollectionMethod`](../../doc/models/collection-method.md) | Optional | The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`. |
| `SubscriptionIds` | `[]int` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroup := models.SubscriptionGroup{
        Uid:                     models.ToPointer("uid8"),
        CustomerId:              models.ToPointer(78),
        PaymentProfile:          models.ToPointer(models.SubscriptionGroupPaymentProfile{
            Id:                   models.ToPointer(44),
            FirstName:            models.ToPointer("first_name4"),
            LastName:             models.ToPointer("last_name2"),
            MaskedCardNumber:     models.ToPointer("masked_card_number2"),
        }),
        PaymentCollectionMethod: models.ToPointer(models.CollectionMethod_AUTOMATIC),
        SubscriptionIds:         []int{
            188,
            189,
            190,
        },
    }

}
```

