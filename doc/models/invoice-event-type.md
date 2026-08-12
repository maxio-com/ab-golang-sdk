
# Invoice Event Type

Invoice Event Type

## Enumeration

`InvoiceEventType`

## Fields

| Name |
|  --- |
| `ISSUEINVOICE` |
| `APPLYCREDITNOTE` |
| `CREATECREDITNOTE` |
| `APPLYPAYMENT` |
| `APPLYDEBITNOTE` |
| `CREATEDEBITNOTE` |
| `REFUNDINVOICE` |
| `VOIDINVOICE` |
| `VOIDREMAINDER` |
| `BACKPORTINVOICE` |
| `CHANGEINVOICESTATUS` |
| `CHANGEINVOICECOLLECTIONMETHOD` |
| `REMOVEPAYMENT` |
| `FAILEDPAYMENT` |
| `CHANGECHARGEBACKSTATUS` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceEventType := models.InvoiceEventType_CREATECREDITNOTE

}
```

