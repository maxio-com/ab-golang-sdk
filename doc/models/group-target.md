
# Group Target

Attributes of the target customer who will be the responsible payer of the created subscription. Required.

## Structure

`GroupTarget`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`models.GroupTargetType`](../../doc/models/group-target-type.md) | Required | The type of object indicated by the id attribute. |
| `Id` | `*int` | Optional | The id of the target customer or subscription to group the existing subscription with. Ignored and should not be included if type is "self", "parent", or "eldest". |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    groupTarget := models.GroupTarget{
        Type:                 models.GroupTargetType_PARENT,
        Id:                   models.ToPointer(234),
    }

}
```

