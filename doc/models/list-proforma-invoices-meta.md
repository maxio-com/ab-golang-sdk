
# List Proforma Invoices Meta

## Structure

`ListProformaInvoicesMeta`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TotalCount` | `*int` | Optional | - |
| `CurrentPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |
| `StatusCode` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listProformaInvoicesMeta := models.ListProformaInvoicesMeta{
        TotalCount:           models.ToPointer(50),
        CurrentPage:          models.ToPointer(26),
        TotalPages:           models.ToPointer(38),
        StatusCode:           models.ToPointer(68),
    }

}
```

