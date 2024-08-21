
# Create Debit Note Event

## Structure

`CreateDebitNoteEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"create_debit_note"` |
| `EventData` | [`models.DebitNote`](../../doc/models/debit-note.md) | Required | Example schema for an `create_debit_note` event |

## Example (as JSON)

```json
{
  "id": 98,
  "timestamp": "2016-03-13T12:52:32.123Z",
  "invoice": {
    "issue_date": "2024-01-01",
    "due_date": "2024-01-01",
    "paid_date": "2024-01-01",
    "public_url_expires_on": "2024-01-21",
    "id": 166,
    "uid": "uid6",
    "site_id": 92,
    "customer_id": 204,
    "subscription_id": 20
  },
  "event_type": "create_debit_note",
  "event_data": {
    "uid": "uid6",
    "site_id": 132,
    "customer_id": 244,
    "subscription_id": 60,
    "number": 64
  }
}
```

