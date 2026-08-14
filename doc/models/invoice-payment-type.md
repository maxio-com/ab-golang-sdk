
# Invoice Payment Type

The type of payment to be applied to an Invoice. Defaults to external.

## Enumeration

`InvoicePaymentType`

## Fields

| Name |
|  --- |
| `EXTERNAL` |
| `PREPAYMENT` |
| `SERVICECREDIT` |
| `PAYMENT` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoicePaymentType := models.InvoicePaymentType_EXTERNAL

}
```

