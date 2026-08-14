
# Revoked Invitation

## Structure

`RevokedInvitation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LastSentAt` | `*string` | Optional | - |
| `LastAcceptedAt` | `*string` | Optional | - |
| `UninvitedCount` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    revokedInvitation := models.RevokedInvitation{
        LastSentAt:           models.ToPointer("last_sent_at8"),
        LastAcceptedAt:       models.ToPointer("last_accepted_at8"),
        UninvitedCount:       models.ToPointer(130),
    }

}
```

