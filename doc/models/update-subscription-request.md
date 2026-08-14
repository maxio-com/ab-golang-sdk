
# Update Subscription Request

## Structure

`UpdateSubscriptionRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`models.UpdateSubscription`](../../doc/models/update-subscription.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateSubscriptionRequest := models.UpdateSubscriptionRequest{
        Subscription:         models.UpdateSubscription{
            CreditCardAttributes:              models.ToPointer(models.CreditCardAttributes{
                FullNumber:           models.ToPointer("full_number2"),
                ExpirationMonth:      models.ToPointer("expiration_month6"),
                ExpirationYear:       models.ToPointer("expiration_year2"),
            }),
            ProductHandle:                     models.ToPointer("product_handle6"),
            ProductId:                         models.ToPointer(206),
            ProductChangeDelayed:              models.ToPointer(false),
            NextProductId:                     models.ToPointer("next_product_id6"),
            DeferSignup:                       models.ToPointer(false),
            DunningCommunicationDelayTimeZone: models.NewOptional(models.ToPointer("\"Eastern Time (US & Canada)\"")),
        },
    }

}
```

