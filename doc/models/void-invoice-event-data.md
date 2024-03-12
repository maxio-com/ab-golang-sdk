
# Void Invoice Event Data

Example schema for an `void_invoice` event

## Structure

`VoidInvoiceEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreditNoteAttributes` | [`*models.VoidInvoiceEventDataCreditNoteAttributes`](../../doc/models/containers/void-invoice-event-data-credit-note-attributes.md) | Required | This is a container for one-of cases. |
| `Memo` | `*string` | Required | The memo provided during invoice voiding. |
| `AppliedAmount` | `*string` | Required | The amount of the void. |
| `TransactionTime` | `*time.Time` | Required | The time the refund was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |
| `IsAdvanceInvoice` | `bool` | Required | If true, the invoice is an advance invoice. |
| `Reason` | `string` | Required | The reason for the void. |

## Example (as JSON)

```json
{
  "credit_note_attributes": {
    "uid": "uid2",
    "site_id": 218,
    "customer_id": 74,
    "subscription_id": 146,
    "number": "number0"
  },
  "memo": "memo6",
  "applied_amount": "applied_amount6",
  "transaction_time": "2016-03-13T12:52:32.123Z",
  "is_advance_invoice": false,
  "reason": "reason8"
}
```

