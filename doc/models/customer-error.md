
# Customer Error

## Structure

`CustomerError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Customer` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    customerError := models.CustomerError{
        Customer:             models.ToPointer("customer8"),
    }

}
```

