
# Component Price Point Currency Overage Response

## Structure

`ComponentPricePointCurrencyOverageResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.CurrencyOveragePrices`](../../doc/models/currency-overage-prices.md) | Required | Extends a component price point with currency overage prices. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentPricePointCurrencyOverageResponse := models.ComponentPricePointCurrencyOverageResponse{
        PricePoint:           models.CurrencyOveragePrices{
            Id:                       models.ToPointer(248),
            Type:                     models.ToPointer(models.PricePointType_ENUMDEFAULT),
            Default:                  models.ToPointer(false),
            Name:                     models.ToPointer("name0"),
            PricingScheme:            models.ToPointer(models.PricingScheme_PERUNIT),
        },
    }

}
```

