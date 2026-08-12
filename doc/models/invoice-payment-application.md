
# Invoice Payment Application

## Structure

`InvoicePaymentApplication`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InvoiceUid` | `*string` | Optional | Unique identifier for the paid invoice. It has the prefix "inv_" followed by alphanumeric characters. |
| `ApplicationUid` | `*string` | Optional | Unique identifier for the payment. It has the prefix "pmt_" followed by alphanumeric characters. |
| `AppliedAmount` | `*string` | Optional | Dollar amount of the paid invoice. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoicePaymentApplication := models.InvoicePaymentApplication{
        InvoiceUid:           models.ToPointer("invoice_uid8"),
        ApplicationUid:       models.ToPointer("application_uid8"),
        AppliedAmount:        models.ToPointer("applied_amount0"),
    }

}
```

