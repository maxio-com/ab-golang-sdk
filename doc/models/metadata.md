
# Metadata

## Structure

`Metadata`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `models.Optional[int]` | Optional | - |
| `Value` | `models.Optional[string]` | Optional | - |
| `ResourceId` | `models.Optional[int]` | Optional | - |
| `Name` | `*string` | Optional | - |
| `DeletedAt` | `models.Optional[time.Time]` | Optional | - |
| `MetafieldId` | `models.Optional[int]` | Optional | - |

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
    metadata := models.Metadata{
        Id:                   models.NewOptional(models.ToPointer(50)),
        Value:                models.NewOptional(models.ToPointer("value8")),
        ResourceId:           models.NewOptional(models.ToPointer(134)),
        Name:                 models.ToPointer("name6"),
        DeletedAt:            models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
    }

}
```

