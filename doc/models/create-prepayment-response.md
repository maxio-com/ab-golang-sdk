
# Create Prepayment Response

## Structure

`CreatePrepaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayment` | [`models.CreatedPrepayment`](../../doc/models/created-prepayment.md) | Required | - |

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
    createPrepaymentResponse := models.CreatePrepaymentResponse{
        Prepayment:           models.CreatedPrepayment{
            Id:                     models.ToPointer(int64(38)),
            SubscriptionId:         models.ToPointer(148),
            AmountInCents:          models.ToPointer(int64(124)),
            Memo:                   models.ToPointer("memo2"),
            CreatedAt:              models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        },
    }

}
```

