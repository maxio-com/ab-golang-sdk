
# Create or Update Endpoint Request

Used to Create or Update Endpoint.

## Structure

`CreateOrUpdateEndpointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Endpoint` | [`models.CreateOrUpdateEndpoint`](../../doc/models/create-or-update-endpoint.md) | Required | Used to Create or Update Endpoint. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createOrUpdateEndpointRequest := models.CreateOrUpdateEndpointRequest{
        Endpoint:             models.CreateOrUpdateEndpoint{
            Url:                  "url2",
            WebhookSubscriptions: []models.WebhookSubscription{
                models.WebhookSubscription_STATEMENTCLOSED,
            },
        },
    }

}
```

