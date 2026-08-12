
# Update Metadata Request

## Structure

`UpdateMetadataRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metadata` | [`*models.UpdateMetadata`](../../doc/models/update-metadata.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateMetadataRequest := models.UpdateMetadataRequest{
        Metadata:             models.ToPointer(models.UpdateMetadata{
            CurrentName:          models.ToPointer("current_name0"),
            Name:                 models.ToPointer("name6"),
            Value:                models.ToPointer("value8"),
        }),
    }

}
```

