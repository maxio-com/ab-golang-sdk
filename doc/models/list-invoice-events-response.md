
# List Invoice Events Response

## Structure

`ListInvoiceEventsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.InvoiceEvent`](../../doc/models/containers/invoice-event.md) | Optional | - |
| `Page` | `*int` | Optional | - |
| `PerPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "event_type": "apply_credit_note",
      "id": 214,
      "timestamp": "2016-03-13T12:52:32.123Z",
      "invoice": {
        "id": 166,
        "uid": "uid6",
        "site_id": 92,
        "customer_id": 204,
        "subscription_id": 20
      },
      "event_data": {
        "uid": "uid6",
        "credit_note_number": "credit_note_number0",
        "credit_note_uid": "credit_note_uid0",
        "original_amount": "original_amount0",
        "applied_amount": "applied_amount2",
        "transaction_time": "2016-03-13T12:52:32.123Z",
        "memo": "memo0",
        "role": "role0",
        "consolidated_invoice": false,
        "applied_credit_notes": [
          {
            "uid": "uid4",
            "number": "number8"
          },
          {
            "uid": "uid4",
            "number": "number8"
          },
          {
            "uid": "uid4",
            "number": "number8"
          }
        ]
      }
    }
  ],
  "page": 184,
  "per_page": 96,
  "total_pages": 194
}
```

