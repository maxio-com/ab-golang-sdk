
# Activate Subscription Request

## Structure

`ActivateSubscriptionRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RevertOnFailure` | `models.Optional[bool]` | Optional | You may choose how to handle the activation failure. `true` means do not change the subscription’s state and billing period. `false` means to continue through with the activation and enter an end-of-life state. If this parameter is omitted or `null` is passed it will default to the value set in the site settings (default: `true`). |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    activateSubscriptionRequest := models.ActivateSubscriptionRequest{
        RevertOnFailure:      models.NewOptional(models.ToPointer(false)),
    }

}
```

