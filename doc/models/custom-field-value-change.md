
# Custom Field Value Change

## Structure

`CustomFieldValueChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EventType` | `string` | Required | - |
| `MetafieldName` | `string` | Required | - |
| `MetafieldId` | `int` | Required | - |
| `OldValue` | `*string` | Required | - |
| `NewValue` | `*string` | Required | - |
| `ResourceType` | `string` | Required | - |
| `ResourceId` | `int` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    customFieldValueChange := models.CustomFieldValueChange{
        EventType:            "event_type8",
        MetafieldName:        "metafield_name2",
        MetafieldId:          138,
        OldValue:             models.ToPointer("old_value6"),
        NewValue:             models.ToPointer("new_value2"),
        ResourceType:         "resource_type6",
        ResourceId:           14,
    }

}
```

