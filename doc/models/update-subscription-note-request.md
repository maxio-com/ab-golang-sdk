
# Update Subscription Note Request

Updatable fields for Subscription Note

## Structure

`UpdateSubscriptionNoteRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Note` | [`models.UpdateSubscriptionNote`](../../doc/models/update-subscription-note.md) | Required | Updatable fields for Subscription Note |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateSubscriptionNoteRequest := models.UpdateSubscriptionNoteRequest{
        Note:                 models.UpdateSubscriptionNote{
            Body:                 "body0",
            Sticky:               false,
        },
    }

}
```

