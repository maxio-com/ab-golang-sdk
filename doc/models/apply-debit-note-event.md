
# Apply Debit Note Event

## Structure

`ApplyDebitNoteEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"apply_debit_note"` |
| `EventData` | [`models.ApplyDebitNoteEventData`](../../doc/models/apply-debit-note-event-data.md) | Required | Example schema for an `apply_debit_note` event |

## Example (as JSON)

```json
{
  "id": 164,
  "timestamp": "2016-03-13T12:52:32.123Z",
  "invoice": {
    "issue_date": "2024-01-01",
    "due_date": "2024-01-01",
    "paid_date": "2024-01-01",
    "id": 166,
    "uid": "uid6",
    "site_id": 92,
    "customer_id": 204,
    "subscription_id": 20
  },
  "event_type": "apply_debit_note",
  "event_data": {
    "debit_note_number": "debit_note_number6",
    "debit_note_uid": "debit_note_uid2",
    "original_amount": "original_amount0",
    "applied_amount": "applied_amount2",
    "memo": "memo0",
    "transaction_time": "2016-03-13T12:52:32.123Z"
  }
}
```

