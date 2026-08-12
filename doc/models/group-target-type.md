
# Group Target Type

The type of object indicated by the id attribute.

## Enumeration

`GroupTargetType`

## Fields

| Name |
|  --- |
| `CUSTOMER` |
| `SUBSCRIPTION` |
| `SELF` |
| `PARENT` |
| `ELDEST` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    groupTargetType := models.GroupTargetType_SELF

}
```

