
# Create Usage Request

## Structure

`CreateUsageRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Usage` | [`models.CreateUsage`](../../doc/models/create-usage.md) | Required | - |

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
    createUsageRequest := models.CreateUsageRequest{
        Usage:                models.CreateUsage{
            Quantity:             models.ToPointer(float64(162.34)),
            PricePointId:         models.ToPointer("price_point_id0"),
            Memo:                 models.ToPointer("memo2"),
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
        },
    }

}
```

