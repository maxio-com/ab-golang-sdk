
# Subscription Custom Price

(Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription. A subscription can have up to 30 custom price points. Exceeding this limit will result in an API error.

## Structure

`SubscriptionCustomPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | (Optional) |
| `Handle` | `*string` | Optional | (Optional) |
| `PriceInCents` | [`models.SubscriptionCustomPricePriceInCents`](../../doc/models/containers/subscription-custom-price-price-in-cents.md) | Required | This is a container for one-of cases. |
| `Interval` | [`models.SubscriptionCustomPriceInterval`](../../doc/models/containers/subscription-custom-price-interval.md) | Required | This is a container for one-of cases. |
| `IntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Required | Required if using `custom_price` attribute. |
| `TrialPriceInCents` | [`*models.SubscriptionCustomPriceTrialPriceInCents`](../../doc/models/containers/subscription-custom-price-trial-price-in-cents.md) | Optional | This is a container for one-of cases. |
| `TrialInterval` | [`*models.SubscriptionCustomPriceTrialInterval`](../../doc/models/containers/subscription-custom-price-trial-interval.md) | Optional | This is a container for one-of cases. |
| `TrialIntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | (Optional) |
| `TrialType` | [`models.Optional[models.TrialType]`](../../doc/models/trial-type.md) | Optional | Indicates how a trial is handled when the trial period ends and there is no credit card on file. For `no_obligation`, the subscription transitions to a Trial Ended state. Maxio will not send any emails or statements. For `payment_expected`, the subscription transitions to a Past Due state. Maxio will send normal dunning emails and statements according to your other settings. |
| `InitialChargeInCents` | [`*models.SubscriptionCustomPriceInitialChargeInCents`](../../doc/models/containers/subscription-custom-price-initial-charge-in-cents.md) | Optional | This is a container for one-of cases. |
| `InitialChargeAfterTrial` | `*bool` | Optional | (Optional) |
| `ExpirationInterval` | [`*models.SubscriptionCustomPriceExpirationInterval`](../../doc/models/containers/subscription-custom-price-expiration-interval.md) | Optional | This is a container for one-of cases. |
| `ExpirationIntervalUnit` | [`models.Optional[models.ExpirationIntervalUnit]`](../../doc/models/expiration-interval-unit.md) | Optional | (Optional) |
| `TaxIncluded` | `*bool` | Optional | (Optional) |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionCustomPrice := models.SubscriptionCustomPrice{
        Name:                    models.ToPointer("name2"),
        Handle:                  models.ToPointer("handle8"),
        PriceInCents:            models.SubscriptionCustomPricePriceInCentsContainer.FromString("String5"),
        Interval:                models.SubscriptionCustomPriceIntervalContainer.FromString("String1"),
        IntervalUnit:            models.ToPointer(models.IntervalUnit_DAY),
        TrialPriceInCents:       models.ToPointer(models.SubscriptionCustomPriceTrialPriceInCentsContainer.FromString("String5")),
        TrialInterval:           models.ToPointer(models.SubscriptionCustomPriceTrialIntervalContainer.FromString("String9")),
        TrialIntervalUnit:       models.ToPointer(models.IntervalUnit_DAY),
    }

}
```

