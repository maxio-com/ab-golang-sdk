
# Backport Invoice Event

## Structure

`BackportInvoiceEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"backport_invoice"` |
| `EventData` | [`models.Invoice`](../../doc/models/invoice.md) | Required | Example schema for an `backport_invoice` event |

## Example (as JSON)

```json
{
  "id": 78,
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
  "event_type": "backport_invoice",
  "event_data": {
    "issue_date": "2024-01-01",
    "due_date": "2024-01-01",
    "paid_date": "2024-01-01",
    "id": 206,
    "uid": "uid6",
    "site_id": 132,
    "customer_id": 244,
    "subscription_id": 60
  }
}
```

