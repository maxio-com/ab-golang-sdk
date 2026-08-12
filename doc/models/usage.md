
# Usage

## Structure

`Usage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int64` | Optional | **Constraints**: `>= 0` |
| `Memo` | `models.Optional[string]` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `Quantity` | [`*models.UsageQuantity`](../../doc/models/containers/usage-quantity.md) | Optional | This is a container for one-of cases. |
| `OverageQuantity` | `*int` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `ComponentHandle` | `*string` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |

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
    usage := models.Usage{
        Id:                   models.ToPointer(int64(150)),
        Memo:                 models.NewOptional(models.ToPointer("memo2")),
        CreatedAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        PricePointId:         models.ToPointer(28),
        Quantity:             models.ToPointer(models.UsageQuantityContainer.FromNumber(28)),
    }

}
```

