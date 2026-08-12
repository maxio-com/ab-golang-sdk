
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

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupSubscriptionError := models.SubscriptionGroupSubscriptionError{
        Product:                       []string{
            "product7",
            "product8",
        },
        ProductPricePointId:           []string{
            "product_price_point_id3",
            "product_price_point_id4",
        },
        PaymentProfile:                []string{
            "payment_profile8",
            "payment_profile9",
        },
        PaymentProfileChargifyToken:   []string{
            "payment_profile.chargify_token2",
            "payment_profile.chargify_token3",
        },
        Base:                          []string{
            "base1",
            "base2",
            "base3",
        },
    }

}
```

