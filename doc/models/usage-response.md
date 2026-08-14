
# Usage Response

## Structure

`UsageResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Usage` | [`models.Usage`](../../doc/models/usage.md) | Required | - |

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
    usageResponse := models.UsageResponse{
        Usage:                models.Usage{
            Id:                   models.ToPointer(int64(150)),
            Memo:                 models.NewOptional(models.ToPointer("memo2")),
            CreatedAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            PricePointId:         models.ToPointer(28),
            Quantity:             models.ToPointer(models.UsageQuantityContainer.FromNumber(28)),
        },
    }

}
```

