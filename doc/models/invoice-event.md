
# Invoice Event

## Structure

`InvoiceEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `EventType` | [`*models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Optional | Invoice Event Type |
| `EventData` | [`*models.InvoiceEvent1`](../../doc/models/invoice-event-1.md) | Optional | The event data is the data that, when combined with the command, results in the output invoice found in the invoice field. |
| `Timestamp` | `*string` | Optional | - |
| `Invoice` | [`*models.Invoice`](../../doc/models/invoice.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": 78,
  "event_type": "void_invoice",
  "event_data": {
    "uid": "uid6",
    "credit_note_number": "credit_note_number0",
    "credit_note_uid": "credit_note_uid0",
    "original_amount": "original_amount0",
    "applied_amount": "applied_amount2"
  },
  "timestamp": "timestamp8",
  "invoice": {
    "id": 166,
    "uid": "uid6",
    "site_id": 92,
    "customer_id": 204,
    "subscription_id": 20
  }
}
```

