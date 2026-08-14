
# Base String Error

The error is base if it is not directly associated with a single attribute.

## Structure

`BaseStringError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Base` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    baseStringError := models.BaseStringError{
        Base:                 []string{
            "base5",
            "base6",
        },
    }

}
```

