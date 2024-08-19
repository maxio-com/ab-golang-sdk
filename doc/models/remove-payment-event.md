
# Remove Payment Event

## Structure

`RemovePaymentEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"remove_payment"` |
| `EventData` | [`models.RemovePaymentEventData`](../../doc/models/remove-payment-event-data.md) | Required | Example schema for an `remove_payment` event |

## Example (as JSON)

```json
{
  "id": 236,
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
  "event_type": "remove_payment",
  "event_data": {
    "transaction_id": 78,
    "memo": "memo0",
    "applied_amount": "applied_amount2",
    "transaction_time": "2016-03-13T12:52:32.123Z",
    "payment_method": {
      "type": "apple_pay"
    },
    "prepayment": false,
    "original_amount": "original_amount0"
  }
}
```

