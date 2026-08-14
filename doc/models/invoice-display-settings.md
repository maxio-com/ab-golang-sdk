
# Invoice Display Settings

## Structure

`InvoiceDisplaySettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `HideZeroSubtotalLines` | `*bool` | Optional | - |
| `IncludeDiscountsOnLines` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceDisplaySettings := models.InvoiceDisplaySettings{
        HideZeroSubtotalLines:   models.ToPointer(false),
        IncludeDiscountsOnLines: models.ToPointer(false),
    }

}
```

