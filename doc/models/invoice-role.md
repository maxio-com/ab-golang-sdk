
# Invoice Role

## Enumeration

`InvoiceRole`

## Fields

| Name |
|  --- |
| `UNSET` |
| `SIGNUP` |
| `RENEWAL` |
| `USAGE` |
| `REACTIVATION` |
| `PRORATION` |
| `MIGRATION` |
| `ADHOC` |
| `BACKPORT` |
| `BACKPORTBALANCERECONCILIATION` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceRole := models.InvoiceRole_RENEWAL

}
```

