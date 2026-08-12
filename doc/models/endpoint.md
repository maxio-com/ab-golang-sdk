
# Endpoint

## Structure

`Endpoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Url` | `*string` | Optional | - |
| `SiteId` | `*int` | Optional | - |
| `Status` | `*string` | Optional | - |
| `WebhookSubscriptions` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    endpoint := models.Endpoint{
        Id:                   models.ToPointer(202),
        Url:                  models.ToPointer("url2"),
        SiteId:               models.ToPointer(128),
        Status:               models.ToPointer("status0"),
        WebhookSubscriptions: []string{
            "webhook_subscriptions4",
        },
    }

}
```

