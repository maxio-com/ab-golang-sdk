
# Renewal Preview Response

## Structure

`RenewalPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RenewalPreview` | [`models.RenewalPreview`](../../doc/models/renewal-preview.md) | Required | - |

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
    renewalPreviewResponse := models.RenewalPreviewResponse{
        RenewalPreview:       models.RenewalPreview{
            NextAssessmentAt:       models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            SubtotalInCents:        models.ToPointer(int64(132)),
            TotalTaxInCents:        models.ToPointer(int64(0)),
            TotalDiscountInCents:   models.ToPointer(int64(250)),
            TotalInCents:           models.ToPointer(int64(20)),
        },
    }

}
```

