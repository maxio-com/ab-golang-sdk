
# List Invoice Events Response

## Structure

`ListInvoiceEventsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.InvoiceEvent`](../../doc/models/invoice-event.md) | Optional | - |
| `Page` | `*int` | Optional | - |
| `PerPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": 68,
      "event_type": "void_invoice",
      "event_data": {
        "uid": "uid2",
        "credit_note_number": "credit_note_number4",
        "credit_note_uid": "credit_note_uid4",
        "original_amount": "original_amount6",
        "applied_amount": "applied_amount6",
        "transaction_time": "2016-03-13T12:52:32.123Z",
        "memo": "memo6",
        "role": "role4",
        "consolidated_invoice": false,
        "applied_credit_notes": [
          {
            "uid": "uid4",
            "number": "number8"
          },
          {
            "uid": "uid4",
            "number": "number8"
          }
        ]
      },
      "timestamp": "2016-03-13T12:52:32.123Z",
      "invoice": {
        "id": 166,
        "uid": "uid6",
        "site_id": 92,
        "customer_id": 204,
        "subscription_id": 20
      }
    }
  ],
  "page": 184,
  "per_page": 96,
  "total_pages": 194
}
```

