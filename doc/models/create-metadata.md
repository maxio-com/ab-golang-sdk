
# Create Metadata

## Structure

`CreateMetadata`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Value` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createMetadata := models.CreateMetadata{
        Name:                 models.ToPointer("name0"),
        Value:                models.ToPointer("value2"),
    }

}
```

