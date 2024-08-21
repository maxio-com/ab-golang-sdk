
# Change Invoice Collection Method Event

## Structure

`ChangeInvoiceCollectionMethodEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"change_invoice_collection_method"` |
| `EventData` | [`models.ChangeInvoiceCollectionMethodEventData`](../../doc/models/change-invoice-collection-method-event-data.md) | Required | Example schema for an `change_invoice_collection_method` event |

## Example (as JSON)

```json
{
  "id": 246,
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
  "event_type": "change_invoice_collection_method",
  "event_data": {
    "from_collection_method": "from_collection_method4",
    "to_collection_method": "to_collection_method8"
  }
}
```

