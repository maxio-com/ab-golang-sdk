
# Update Metadata

## Structure

`UpdateMetadata`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrentName` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Value` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateMetadata := models.UpdateMetadata{
        CurrentName:          models.ToPointer("current_name8"),
        Name:                 models.ToPointer("name4"),
        Value:                models.ToPointer("value6"),
    }

}
```

