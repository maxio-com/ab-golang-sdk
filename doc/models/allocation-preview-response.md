
# Allocation Preview Response

## Structure

`AllocationPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllocationPreview` | [`models.AllocationPreview`](../../doc/models/allocation-preview.md) | Required | - |

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
    allocationPreviewResponse := models.AllocationPreviewResponse{
        AllocationPreview:    models.AllocationPreview{
            StartDate:              models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            EndDate:                models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            SubtotalInCents:        models.ToPointer(int64(240)),
            TotalTaxInCents:        models.ToPointer(int64(108)),
            TotalDiscountInCents:   models.ToPointer(int64(142)),
        },
    }

}
```

