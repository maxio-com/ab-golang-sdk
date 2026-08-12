
# Replay Webhooks Response

## Structure

`ReplayWebhooksResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    replayWebhooksResponse := models.ReplayWebhooksResponse{
        Status:               models.ToPointer("status6"),
    }

}
```

