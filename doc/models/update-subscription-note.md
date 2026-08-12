
# Update Subscription Note

Updatable fields for Subscription Note

## Structure

`UpdateSubscriptionNote`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Body` | `string` | Required | - |
| `Sticky` | `bool` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateSubscriptionNote := models.UpdateSubscriptionNote{
        Body:                 "body2",
        Sticky:               false,
    }

}
```

