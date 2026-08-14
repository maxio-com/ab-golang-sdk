
# Endpoint Response

## Structure

`EndpointResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Endpoint` | [`*models.Endpoint`](../../doc/models/endpoint.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    endpointResponse := models.EndpointResponse{
        Endpoint:             models.ToPointer(models.Endpoint{
            Id:                   models.ToPointer(202),
            Url:                  models.ToPointer("url2"),
            SiteId:               models.ToPointer(128),
            Status:               models.ToPointer("status0"),
            WebhookSubscriptions: []string{
                "webhook_subscriptions4",
            },
        }),
    }

}
```

