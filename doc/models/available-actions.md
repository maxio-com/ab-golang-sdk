
# Available Actions

## Structure

`AvailableActions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SendEmail` | [`*models.SendEmail`](../../doc/models/send-email.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    availableActions := models.AvailableActions{
        SendEmail:            models.ToPointer(models.SendEmail{
            CanExecute:           false,
            Url:                  "url0",
        }),
    }

}
```

