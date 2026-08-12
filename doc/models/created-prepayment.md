
# Created Prepayment

## Structure

`CreatedPrepayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int64` | Optional | **Constraints**: `>= 1` |
| `SubscriptionId` | `*int` | Optional | **Constraints**: `>= 1` |
| `AmountInCents` | `*int64` | Optional | **Constraints**: `>= 0.01` |
| `Memo` | `*string` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `StartingBalanceInCents` | `*int64` | Optional | **Constraints**: `>= 0` |
| `EndingBalanceInCents` | `*int64` | Optional | - |

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
    createdPrepayment := models.CreatedPrepayment{
        Id:                     models.ToPointer(int64(186)),
        SubscriptionId:         models.ToPointer(40),
        AmountInCents:          models.ToPointer(int64(240)),
        Memo:                   models.ToPointer("memo6"),
        CreatedAt:              models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

