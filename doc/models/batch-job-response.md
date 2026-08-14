
# Batch Job Response

## Structure

`BatchJobResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Batchjob` | [`models.BatchJob`](../../doc/models/batch-job.md) | Required | - |

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
    batchJobResponse := models.BatchJobResponse{
        Batchjob:             models.BatchJob{
            Id:                   models.ToPointer(54),
            FinishedAt:           models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
            RowCount:             models.NewOptional(models.ToPointer(62)),
            CreatedAt:            models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
            Completed:            models.ToPointer("completed4"),
        },
    }

}
```

