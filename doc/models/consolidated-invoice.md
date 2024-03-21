
# Consolidated Invoice

## Structure

`ConsolidatedInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Invoices` | [`[]models.Invoice`](../../doc/models/invoice.md) | Optional | - |

## Example (as JSON)

```json
{
  "invoices": [
    {
      "id": 196,
      "uid": "uid6",
      "site_id": 122,
      "customer_id": 234,
      "subscription_id": 50
    },
    {
      "id": 196,
      "uid": "uid6",
      "site_id": 122,
      "customer_id": 234,
      "subscription_id": 50
    }
  ]
}
```

