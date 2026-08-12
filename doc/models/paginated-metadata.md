
# Paginated Metadata

## Structure

`PaginatedMetadata`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TotalCount` | `*int` | Optional | - |
| `CurrentPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |
| `PerPage` | `*int` | Optional | - |
| `Metadata` | [`[]models.Metadata`](../../doc/models/metadata.md) | Optional | - |

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
    paginatedMetadata := models.PaginatedMetadata{
        TotalCount:           models.ToPointer(166),
        CurrentPage:          models.ToPointer(142),
        TotalPages:           models.ToPointer(154),
        PerPage:              models.ToPointer(136),
        Metadata:             []models.Metadata{
            models.Metadata{
                Id:                   models.NewOptional(models.ToPointer(50)),
                Value:                models.NewOptional(models.ToPointer("value8")),
                ResourceId:           models.NewOptional(models.ToPointer(134)),
                Name:                 models.ToPointer("name6"),
                DeletedAt:            models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
            },
        },
    }

}
```

