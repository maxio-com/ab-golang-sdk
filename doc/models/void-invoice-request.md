
# Void Invoice Request

## Structure

`VoidInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Void` | [`models.VoidInvoice`](../../doc/models/void-invoice.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    voidInvoiceRequest := models.VoidInvoiceRequest{
        Void:                 models.VoidInvoice{
            Reason:               "reason6",
        },
    }

}
```

