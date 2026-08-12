
# Attribute Error

## Structure

`AttributeError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Attribute` | `[]string` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    attributeError := models.AttributeError{
        Attribute:            []string{
            "attribute8",
            "attribute7",
        },
    }

}
```

