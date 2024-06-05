
# Void Invoice Event

## Structure

`VoidInvoiceEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"void_invoice"` |
| `EventData` | [`models.VoidInvoiceEventData`](../../doc/models/void-invoice-event-data.md) | Required | Example schema for an `void_invoice` event |

## Example (as JSON)

```json
{
  "id": 16,
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
  "event_type": "void_invoice",
  "event_data": {
    "credit_note_attributes": {
      "uid": "uid2",
      "site_id": 72,
      "customer_id": 184,
      "subscription_id": 0,
      "number": "number0"
    },
    "memo": "memo0",
    "applied_amount": "applied_amount2",
    "transaction_time": "2016-03-13T12:52:32.123Z",
    "is_advance_invoice": false,
    "reason": "reason2"
  }
}
```

