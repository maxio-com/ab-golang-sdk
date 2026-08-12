
# Invoice Payment Method Type

The type of payment method used. Defaults to other.

## Enumeration

`InvoicePaymentMethodType`

## Fields

| Name |
|  --- |
| `CREDITCARD` |
| `CHECK` |
| `CASH` |
| `MONEYORDER` |
| `ACH` |
| `OTHER` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoicePaymentMethodType := models.InvoicePaymentMethodType_CASH

}
```

