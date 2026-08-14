
# Deliver Proforma Invoice Request

## Structure

`DeliverProformaInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RecipientEmails` | `[]string` | Optional | - |
| `CcRecipientEmails` | `[]string` | Optional | - |
| `BccRecipientEmails` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    deliverProformaInvoiceRequest := models.DeliverProformaInvoiceRequest{
        RecipientEmails:      []string{
            "recipient_emails3",
            "recipient_emails4",
        },
        CcRecipientEmails:    []string{
            "cc_recipient_emails2",
            "cc_recipient_emails1",
            "cc_recipient_emails0",
        },
        BccRecipientEmails:   []string{
            "bcc_recipient_emails6",
        },
    }

}
```

