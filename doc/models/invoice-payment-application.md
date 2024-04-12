
# Invoice Payment Application

## Structure

`InvoicePaymentApplication`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InvoiceUid` | `*string` | Optional | Unique identifier for the paid invoice. It has the prefix "inv_" followed by alphanumeric characters. |
| `ApplicationUid` | `*string` | Optional | Unique identifier for the payment. It has the prefix "pmt_" followed by alphanumeric characters. |
| `AppliedAmount` | `*string` | Optional | Dollar amount of the paid invoice. |

## Example (as JSON)

```json
{
  "invoice_uid": "invoice_uid2",
  "application_uid": "application_uid4",
  "applied_amount": "applied_amount6"
}
```

