
# Apply Payment Event

## Structure

`ApplyPaymentEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"apply_payment"` |
| `EventData` | [`models.ApplyPaymentEventData`](../../doc/models/apply-payment-event-data.md) | Required | Example schema for an `apply_payment` event |

## Example (as JSON)

```json
{
  "id": 234,
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
  "event_type": "apply_payment",
  "event_data": {
    "consolidation_level": "child",
    "memo": "memo0",
    "original_amount": "original_amount0",
    "applied_amount": "applied_amount2",
    "transaction_time": "2016-03-13T12:52:32.123Z",
    "payment_method": {
      "type": "apple_pay"
    },
    "transaction_id": 78,
    "parent_invoice_number": 36,
    "remaining_prepayment_amount": "remaining_prepayment_amount6",
    "prepayment": false,
    "external": false
  }
}
```

