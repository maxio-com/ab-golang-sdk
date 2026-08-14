
# Enable Webhooks Response

## Structure

`EnableWebhooksResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WebhooksEnabled` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    enableWebhooksResponse := models.EnableWebhooksResponse{
        WebhooksEnabled:      models.ToPointer(false),
    }

}
```

