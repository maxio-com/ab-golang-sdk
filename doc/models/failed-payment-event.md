
# Failed Payment Event

## Structure

`FailedPaymentEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"failed_payment"` |
| `EventData` | [`models.FailedPaymentEventData`](../../doc/models/failed-payment-event-data.md) | Required | Example schema for an `failed_payment` event |

## Example (as JSON)

```json
{
  "id": 120,
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
  "event_type": "failed_payment",
  "event_data": {
    "amount_in_cents": 220,
    "applied_amount": 194,
    "memo": "memo0",
    "payment_method": "cash",
    "transaction_id": 78
  }
}
```

