
# Invoice Status

The current status of the invoice. See [Invoice Statuses](https://docs.maxio.com/billing/invoicing/overview/understand-how-invoices-are-generated#invoice-statuses) for more.

## Enumeration

`InvoiceStatus`

## Fields

| Name |
|  --- |
| `DRAFT` |
| `OPEN` |
| `PAID` |
| `PENDING` |
| `VOIDED` |
| `CANCELED` |
| `PROCESSING` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceStatus := models.InvoiceStatus_PAID

}
```

