
# Change Chargeback Status Event

## Structure

`ChangeChargebackStatusEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"change_chargeback_status"` |
| `EventData` | [`models.ChangeChargebackStatusEventData`](../../doc/models/change-chargeback-status-event-data.md) | Required | Example schema for an `change_chargeback_status` event |

## Example (as JSON)

```json
{
  "id": 214,
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
  "event_type": "change_chargeback_status",
  "event_data": {
    "chargeback_status": "won"
  }
}
```

