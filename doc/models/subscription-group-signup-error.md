
# Subscription Group Signup Error

## Structure

`SubscriptionGroupSignupError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscriptions` | [`map[string]models.SubscriptionGroupSubscriptionError`](../../doc/models/subscription-group-subscription-error.md) | Optional | Subscription errors keyed by the subscription's `reference` from the request, or by its zero-based position in the `subscriptions` array when it has no reference. |
| `PayerReference` | `*string` | Optional | - |
| `Payer` | [`*models.PayerError`](../../doc/models/payer-error.md) | Optional | - |
| `SubscriptionGroup` | `[]string` | Optional | - |
| `PaymentProfileId` | `*string` | Optional | - |
| `PayerId` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupSignupError := models.SubscriptionGroupSignupError{
        Subscriptions:        map[string]models.SubscriptionGroupSubscriptionError{
            "key0": models.SubscriptionGroupSubscriptionError{
                Customer:                      []string{
                    "customer1",
                    "customer2",
                },
                Product:                       []string{
                    "product9",
                },
                ProductPricePointId:           []string{
                    "product_price_point_id7",
                },
                PaymentProfile:                []string{
                    "payment_profile2",
                },
                PaymentProfileChargifyToken:   []string{
                    "payment_profile.chargify_token6",
                },
            },
            "key1": models.SubscriptionGroupSubscriptionError{
                Customer:                      []string{
                    "customer1",
                    "customer2",
                },
                Product:                       []string{
                    "product9",
                },
                ProductPricePointId:           []string{
                    "product_price_point_id7",
                },
                PaymentProfile:                []string{
                    "payment_profile2",
                },
                PaymentProfileChargifyToken:   []string{
                    "payment_profile.chargify_token6",
                },
            },
            "key2": models.SubscriptionGroupSubscriptionError{
                Customer:                      []string{
                    "customer1",
                    "customer2",
                },
                Product:                       []string{
                    "product9",
                },
                ProductPricePointId:           []string{
                    "product_price_point_id7",
                },
                PaymentProfile:                []string{
                    "payment_profile2",
                },
                PaymentProfileChargifyToken:   []string{
                    "payment_profile.chargify_token6",
                },
            },
        },
        PayerReference:       models.ToPointer("payer_reference8"),
        Payer:                models.ToPointer(models.PayerError{
            LastName:             []string{
                "last_name5",
                "last_name6",
            },
            FirstName:            []string{
                "first_name8",
            },
            Email:                []string{
                "email0",
                "email9",
            },
        }),
        SubscriptionGroup:    []string{
            "subscription_group3",
            "subscription_group4",
            "subscription_group5",
        },
        PaymentProfileId:     models.ToPointer("payment_profile_id4"),
    }

}
```

