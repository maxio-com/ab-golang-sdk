
# Enable Webhooks Request

## Structure

`EnableWebhooksRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WebhooksEnabled` | `bool` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    enableWebhooksRequest := models.EnableWebhooksRequest{
        WebhooksEnabled:      false,
    }

}
```

