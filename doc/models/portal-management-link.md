
# Portal Management Link

## Structure

`PortalManagementLink`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Url` | `*string` | Optional | - |
| `FetchCount` | `*int` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `NewLinkAvailableAt` | `*time.Time` | Optional | - |
| `ExpiresAt` | `*time.Time` | Optional | - |
| `LastInviteSentAt` | `models.Optional[time.Time]` | Optional | - |

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
    portalManagementLink := models.PortalManagementLink{
        Url:                  models.ToPointer("url8"),
        FetchCount:           models.ToPointer(88),
        CreatedAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        NewLinkAvailableAt:   models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        ExpiresAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

