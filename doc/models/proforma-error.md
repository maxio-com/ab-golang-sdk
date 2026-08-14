
# Proforma Error

## Structure

`ProformaError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`*models.BaseStringError`](../../doc/models/base-string-error.md) | Optional | The error is base if it is not directly associated with a single attribute. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    proformaError := models.ProformaError{
        Subscription:         models.ToPointer(models.BaseStringError{
            Base:                 []string{
                "base3",
                "base4",
            },
        }),
    }

}
```

