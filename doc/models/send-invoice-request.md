
# Send Invoice Request

## Structure

`SendInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RecipientEmails` | `[]string` | Optional | **Constraints**: *Maximum Items*: `5` |
| `CcRecipientEmails` | `[]string` | Optional | **Constraints**: *Maximum Items*: `5` |
| `BccRecipientEmails` | `[]string` | Optional | **Constraints**: *Maximum Items*: `5` |

## Example (as JSON)

```json
{
  "recipient_emails": [
    "recipient_emails3",
    "recipient_emails4"
  ],
  "cc_recipient_emails": [
    "cc_recipient_emails6",
    "cc_recipient_emails5"
  ],
  "bcc_recipient_emails": [
    "bcc_recipient_emails6"
  ]
}
```

