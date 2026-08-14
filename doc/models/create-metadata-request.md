
# Create Metadata Request

## Structure

`CreateMetadataRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metadata` | [`[]models.CreateMetadata`](../../doc/models/create-metadata.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createMetadataRequest := models.CreateMetadataRequest{
        Metadata:             []models.CreateMetadata{
            models.CreateMetadata{
                Name:                 models.ToPointer("name6"),
                Value:                models.ToPointer("value8"),
            },
        },
    }

}
```

