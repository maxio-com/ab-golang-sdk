
# Create Usage

## Structure

`CreateUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Quantity` | `*float64` | Optional | integer by default or decimal number if fractional quantities are enabled for the component |
| `PricePointId` | `*string` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `BillingSchedule` | [`*models.BillingSchedule`](../../doc/models/billing-schedule.md) | Optional | Billing schedule settings for component allocations or usages on multi-frequency subscriptions. Use this to start a component's billing period on a custom date instead of aligning with the product charge schedule. |
| `CustomPrice` | [`*models.ComponentCustomPrice`](../../doc/models/component-custom-price.md) | Optional | Create or update custom pricing unique to the subscription. Used in place of `price_point_id`. |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    createUsage := models.CreateUsage{
        Quantity:             models.ToPointer(float64(204.7)),
        PricePointId:         models.ToPointer("price_point_id4"),
        Memo:                 models.ToPointer("memo8"),
        BillingSchedule:      models.ToPointer(models.BillingSchedule{
            InitialBillingAt:     models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
        }),
        CustomPrice:          models.ToPointer(models.ComponentCustomPrice{
            TaxIncluded:              models.ToPointer(false),
            PricingScheme:            models.ToPointer(models.PricingScheme_STAIRSTEP),
            Interval:                 models.ToPointer(66),
            IntervalUnit:             models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
            ListPricePointId:         models.NewOptional(models.ToPointer(174)),
            Prices:                   []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                    EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                },
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                    EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                },
            },
        }),
    }

}
```

