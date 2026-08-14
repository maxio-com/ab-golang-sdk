
# Subscription Group Update Error

## Structure

`SubscriptionGroupUpdateError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Members` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupUpdateError := models.SubscriptionGroupUpdateError{
        Members:              []string{
            "members6",
            "members7",
        },
    }

}
```

