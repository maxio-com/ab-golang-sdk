
# Subscription Group Signup Request

## Structure

`SubscriptionGroupSignupRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroup` | [`models.SubscriptionGroupSignup`](../../doc/models/subscription-group-signup.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupSignupRequest := models.SubscriptionGroupSignupRequest{
        SubscriptionGroup:    models.SubscriptionGroupSignup{
            PaymentProfileId:        models.ToPointer(128),
            PayerId:                 models.ToPointer(150),
            PayerReference:          models.ToPointer("payer_reference6"),
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
        },
    }

}
```

