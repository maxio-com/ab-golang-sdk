
# Resent Invitation

## Structure

`ResentInvitation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LastSentAt` | `*string` | Optional | - |
| `LastAcceptedAt` | `*string` | Optional | - |
| `SendInviteLinkText` | `*string` | Optional | - |
| `UninvitedCount` | `*int` | Optional | - |
| `LastInviteSentAt` | `*time.Time` | Optional | - |
| `LastInviteAcceptedAt` | `*time.Time` | Optional | - |

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
    resentInvitation := models.ResentInvitation{
        LastSentAt:           models.ToPointer("last_sent_at0"),
        LastAcceptedAt:       models.ToPointer("last_accepted_at0"),
        SendInviteLinkText:   models.ToPointer("send_invite_link_text8"),
        UninvitedCount:       models.ToPointer(178),
        LastInviteSentAt:     models.ToPointer(parseTime(time.RFC3339, "2024-01-01T00:30:00-04:00", func(err error) { log.Fatalln(err) })),
        LastInviteAcceptedAt: models.ToPointer(parseTime(time.RFC3339, "2024-01-01T00:35:00-04:00", func(err error) { log.Fatalln(err) })),
    }

}
```

