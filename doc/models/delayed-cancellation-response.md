
# Delayed Cancellation Response

## Structure

`DelayedCancellationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Message` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    delayedCancellationResponse := models.DelayedCancellationResponse{
        Message:              models.ToPointer("message4"),
    }

}
```

