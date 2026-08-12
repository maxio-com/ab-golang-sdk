
# Replay Webhooks Request

## Structure

`ReplayWebhooksRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ids` | `[]int64` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    replayWebhooksRequest := models.ReplayWebhooksRequest{
        Ids:                  []int64{
            int64(209),
            int64(210),
            int64(211),
        },
    }

}
```

