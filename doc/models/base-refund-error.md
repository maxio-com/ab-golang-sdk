
# Base Refund Error

## Structure

`BaseRefundError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Base` | `[]interface{}` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    baseRefundError := models.BaseRefundError{
        Base:                 []interface{}{
            interface{}("[key1, val1][key2, val2]"),
            interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

