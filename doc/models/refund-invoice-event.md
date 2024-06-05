
# Refund Invoice Event

## Structure

`RefundInvoiceEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"refund_invoice"` |
| `EventData` | [`models.RefundInvoiceEventData`](../../doc/models/refund-invoice-event-data.md) | Required | Example schema for an `refund_invoice` event |

## Example (as JSON)

```json
{
  "id": 54,
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
  "event_type": "refund_invoice",
  "event_data": {
    "apply_credit": false,
    "consolidation_level": "child",
    "credit_note_attributes": {
      "uid": "uid2",
      "site_id": 72,
      "customer_id": 184,
      "subscription_id": 0,
      "number": "number0"
    },
    "memo": "memo0",
    "original_amount": "original_amount0",
    "payment_id": 204,
    "refund_amount": "refund_amount8",
    "refund_id": 248,
    "transaction_time": "2016-03-13T12:52:32.123Z"
  }
}
```

