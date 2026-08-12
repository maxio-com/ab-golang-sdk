
# Send Invoice Request

## Structure

`SendInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RecipientEmails` | `[]string` | Optional | **Constraints**: *Maximum Items*: `5` |
| `CcRecipientEmails` | `[]string` | Optional | **Constraints**: *Maximum Items*: `5` |
| `BccRecipientEmails` | `[]string` | Optional | **Constraints**: *Maximum Items*: `5` |
| `AttachmentUrls` | `[]string` | Optional | Array of URLs to files to attach to the invoice email. Max 10 files, 10MB each.<br><br>**Constraints**: *Maximum Items*: `10` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    sendInvoiceRequest := models.SendInvoiceRequest{
        RecipientEmails:      []string{
            "recipient_emails7",
        },
        CcRecipientEmails:    []string{
            "cc_recipient_emails2",
        },
        BccRecipientEmails:   []string{
            "bcc_recipient_emails0",
            "bcc_recipient_emails1",
            "bcc_recipient_emails2",
        },
        AttachmentUrls:       []string{
            "attachment_urls4",
        },
    }

}
```

