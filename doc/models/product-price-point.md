
# Product Price Point

## Structure

`ProductPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | The product price point name |
| `Handle` | `models.Optional[string]` | Optional | The product price point API handle |
| `PriceInCents` | `*int64` | Optional | The product price point price, in integer cents |
| `Interval` | `*int` | Optional | The numerical interval. e.g., an interval of ‘30’ coupled with an interval_unit of day would mean this product price point would renew every 30 days. |
| `IntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this product price point, either month or day |
| `TrialPriceInCents` | `models.Optional[int64]` | Optional | The product price point trial price, in integer cents |
| `TrialInterval` | `models.Optional[int]` | Optional | The numerical trial interval. e.g., an interval of ‘30’ coupled with a trial_interval_unit of day would mean this product price point trial would last 30 days. |
| `TrialIntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the trial interval unit for this product price point, either month or day |
| `TrialType` | [`models.Optional[models.TrialType]`](../../doc/models/trial-type.md) | Optional | Indicates how a trial is handled when the trial period ends and there is no credit card on file. For `no_obligation`, the subscription transitions to a Trial Ended state. Maxio will not send any emails or statements. For `payment_expected`, the subscription transitions to a Past Due state. Maxio will send normal dunning emails and statements according to your other settings. |
| `IntroductoryOffer` | `models.Optional[bool]` | Optional | reserved for future use |
| `InitialChargeInCents` | `models.Optional[int64]` | Optional | The product price point initial charge, in integer cents |
| `InitialChargeAfterTrial` | `models.Optional[bool]` | Optional | - |
| `ExpirationInterval` | `models.Optional[int]` | Optional | The numerical expiration interval. e.g., an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product price point would expire after 30 days. |
| `ExpirationIntervalUnit` | [`models.Optional[models.ExpirationIntervalUnit]`](../../doc/models/expiration-interval-unit.md) | Optional | A string representing the expiration interval unit for this product price point, either month, day or never |
| `ProductId` | `*int` | Optional | The product id this price point belongs to |
| `ArchivedAt` | `models.Optional[time.Time]` | Optional | Timestamp indicating when this price point was archived |
| `CreatedAt` | `*time.Time` | Optional | Timestamp indicating when this price point was created |
| `UpdatedAt` | `*time.Time` | Optional | Timestamp indicating when this price point was last updated |
| `UseSiteExchangeRate` | `*bool` | Optional | Whether or not to use the site's exchange rate or define your own pricing when your site has multiple currencies defined. |
| `Type` | [`*models.PricePointType`](../../doc/models/price-point-type.md) | Optional | The type of price point |
| `TaxIncluded` | `*bool` | Optional | Whether or not the price point includes tax |
| `SubscriptionId` | `models.Optional[int]` | Optional | The subscription id this price point belongs to |
| `CurrencyPrices` | [`[]models.CurrencyPrice`](../../doc/models/currency-price.md) | Optional | An array of currency pricing data is available when multiple currencies are defined for the site. It varies based on the use_site_exchange_rate setting for the price point. This parameter is present only in the response of read endpoints, after including the appropriate query parameter. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    productPricePoint := models.ProductPricePoint{
        Id:                      models.ToPointer(10),
        Name:                    models.ToPointer("name0"),
        Handle:                  models.NewOptional(models.ToPointer("handle6")),
        PriceInCents:            models.ToPointer(int64(178)),
        Interval:                models.ToPointer(194),
    }

}
```

