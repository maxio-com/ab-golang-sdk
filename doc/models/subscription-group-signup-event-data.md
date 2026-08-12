
# Subscription Group Signup Event Data

## Structure

`SubscriptionGroupSignupEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroup` | [`models.SubscriptionGroupSignupFailureData`](../../doc/models/subscription-group-signup-failure-data.md) | Required | - |
| `Customer` | [`*models.Customer`](../../doc/models/customer.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupSignupEventData := models.SubscriptionGroupSignupEventData{
        SubscriptionGroup:    models.SubscriptionGroupSignupFailureData{
            PayerId:                 models.ToPointer(150),
            PayerReference:          models.ToPointer("payer_reference6"),
            PaymentProfileId:        models.ToPointer(128),
            PaymentCollectionMethod: models.ToPointer("payment_collection_method8"),
            PayerAttributes:         models.ToPointer(models.PayerAttributes{
                FirstName:            models.ToPointer("first_name2"),
                LastName:             models.ToPointer("last_name0"),
                Email:                models.ToPointer("email4"),
                CcEmails:             models.ToPointer("cc_emails2"),
                Organization:         models.ToPointer("organization6"),
            }),
        },
        Customer:             models.ToPointer(models.Customer{
            FirstName:                   models.ToPointer("first_name0"),
            LastName:                    models.ToPointer("last_name8"),
            Email:                       models.ToPointer("email6"),
            CcEmails:                    models.NewOptional(models.ToPointer("cc_emails0")),
            Organization:                models.NewOptional(models.ToPointer("organization6")),
        }),
    }

}
```

