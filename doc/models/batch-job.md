
# Batch Job

## Structure

`BatchJob`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FinishedAt` | `models.Optional[time.Time]` | Optional | - |
| `RowCount` | `models.Optional[int]` | Optional | - |
| `CreatedAt` | `models.Optional[time.Time]` | Optional | - |
| `Completed` | `*string` | Optional | - |

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
    batchJob := models.BatchJob{
        Id:                   models.ToPointer(60),
        FinishedAt:           models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
        RowCount:             models.NewOptional(models.ToPointer(68)),
        CreatedAt:            models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
        Completed:            models.ToPointer("completed6"),
    }

}
```

