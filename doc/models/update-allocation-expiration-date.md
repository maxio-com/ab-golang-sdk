
# Update Allocation Expiration Date

## Structure

`UpdateAllocationExpirationDate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Allocation` | [`*models.AllocationExpirationDate`](../../doc/models/allocation-expiration-date.md) | Optional | - |

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
    updateAllocationExpirationDate := models.UpdateAllocationExpirationDate{
        Allocation:           models.ToPointer(models.AllocationExpirationDate{
            ExpiresAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        }),
    }

}
```

