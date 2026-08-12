
# Create Invoice Payment Application

## Structure

`CreateInvoicePaymentApplication`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InvoiceUid` | `string` | Required | Unique identifier for the invoice. It has the prefix "inv_" followed by alphanumeric characters. |
| `Amount` | `string` | Required | Dollar amount of the invoice payment (eg. "10.50" => $10.50). |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createInvoicePaymentApplication := models.CreateInvoicePaymentApplication{
        InvoiceUid:           "invoice_uid6",
        Amount:               "amount8",
    }

}
```

