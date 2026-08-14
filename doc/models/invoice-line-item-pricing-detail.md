
# Invoice Line Item Pricing Detail

## Structure

`InvoiceLineItemPricingDetail`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Label` | `*string` | Optional | - |
| `Amount` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceLineItemPricingDetail := models.InvoiceLineItemPricingDetail{
        Label:                models.ToPointer("label4"),
        Amount:               models.ToPointer("amount6"),
    }

}
```

