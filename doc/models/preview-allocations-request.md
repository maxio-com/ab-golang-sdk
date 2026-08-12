
# Preview Allocations Request

## Structure

`PreviewAllocationsRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Allocations` | [`[]models.CreateAllocation`](../../doc/models/create-allocation.md) | Required | - |
| `EffectiveProrationDate` | `*time.Time` | Optional | To calculate proration amounts for a future time. Only within a current subscription period. Only ISO8601 format is supported. |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |

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
    previewAllocationsRequest := models.PreviewAllocationsRequest{
        Allocations:            []models.CreateAllocation{
            models.CreateAllocation{
                Quantity:                 float64(26.48),
                DecimalQuantity:          models.ToPointer("decimal_quantity8"),
                PreviousQuantity:         models.ToPointer(float64(55.5)),
                DecimalPreviousQuantity:  models.ToPointer("decimal_previous_quantity2"),
                ComponentId:              models.ToPointer(242),
                Memo:                     models.ToPointer("memo6"),
            },
        },
        EffectiveProrationDate: models.ToPointer(parseTime(models.DEFAULT_DATE, "2023-12-01", func(err error) { log.Fatalln(err) })),
        UpgradeCharge:          models.NewOptional(models.ToPointer(models.CreditType_NONE)),
        DowngradeCredit:        models.NewOptional(models.ToPointer(models.CreditType_NONE)),
    }

}
```

