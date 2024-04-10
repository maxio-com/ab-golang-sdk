
# Void Remainder Event Data

Example schema for an `void_remainder` event

## Structure

`VoidRemainderEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreditNoteAttributes` | [`models.CreditNote`](../../doc/models/credit-note.md) | Required | - |
| `Memo` | `string` | Required | The memo provided during invoice remainder voiding. |
| `AppliedAmount` | `string` | Required | The amount of the void. |
| `TransactionTime` | `time.Time` | Required | The time the refund was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |

## Example (as JSON)

```json
{
  "credit_note_attributes": {
    "uid": "uid2",
    "site_id": 72,
    "customer_id": 184,
    "subscription_id": 0,
    "number": "number0"
  },
  "memo": "memo6",
  "applied_amount": "applied_amount4",
  "transaction_time": "2016-03-13T12:52:32.123Z"
}
```

