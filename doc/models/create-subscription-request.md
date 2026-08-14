
# Create Subscription Request

## Structure

`CreateSubscriptionRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`models.CreateSubscription`](../../doc/models/create-subscription.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createSubscriptionRequest := models.CreateSubscriptionRequest{
        Subscription:         models.CreateSubscription{
            ProductHandle:                     models.ToPointer("product_handle6"),
            ProductId:                         models.ToPointer(206),
            ProductPricePointHandle:           models.ToPointer("product_price_point_handle2"),
            ProductPricePointId:               models.ToPointer(130),
            CustomPrice:                       models.ToPointer(models.SubscriptionCustomPrice{
                Name:                    models.ToPointer("name4"),
                Handle:                  models.ToPointer("handle0"),
                PriceInCents:            models.SubscriptionCustomPricePriceInCentsContainer.FromString("String3"),
                Interval:                models.SubscriptionCustomPriceIntervalContainer.FromString("String3"),
                IntervalUnit:            models.ToPointer(models.IntervalUnit_DAY),
                TrialPriceInCents:       models.ToPointer(models.SubscriptionCustomPriceTrialPriceInCentsContainer.FromString("String3")),
                TrialInterval:           models.ToPointer(models.SubscriptionCustomPriceTrialIntervalContainer.FromString("String5")),
                TrialIntervalUnit:       models.ToPointer(models.IntervalUnit_DAY),
            }),
            DeferSignup:                       models.ToPointer(false),
            Metafields:                        map[string]string{
                "custom_field_name_1": "custom_field_value_1",
                "custom_field_name_2": "custom_field_value_2",
            },
            DunningCommunicationDelayEnabled:  models.ToPointer(false),
            DunningCommunicationDelayTimeZone: models.NewOptional(models.ToPointer("\"Eastern Time (US & Canada)\"")),
        },
    }

}
```

