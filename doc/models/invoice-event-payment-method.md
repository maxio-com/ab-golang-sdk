
# Invoice Event Payment Method

## Enumeration

`InvoiceEventPaymentMethod`

## Fields

| Name |
|  --- |
| `APPLEPAY` |
| `BANKACCOUNT` |
| `CREDITCARD` |
| `EXTERNAL` |
| `PAYPALACCOUNT` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceEventPaymentMethod := models.InvoiceEventPaymentMethod_CREDITCARD

}
```

