
# Create or Update Endpoint

Used to Create or Update Endpoint.

## Structure

`CreateOrUpdateEndpoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Url` | `string` | Required | - |
| `WebhookSubscriptions` | [`[]models.WebhookSubscription`](../../doc/models/webhook-subscription.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createOrUpdateEndpoint := models.CreateOrUpdateEndpoint{
        Url:                  "url4",
        WebhookSubscriptions: []models.WebhookSubscription{
            models.WebhookSubscription_TRIALENDNOTICE,
            models.WebhookSubscription_SUBSCRIPTIONSTATECHANGE,
            models.WebhookSubscription_SUBSCRIPTIONPRODUCTCHANGESCHEDULED,
        },
    }

}
```

