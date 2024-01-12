
# Signup Proforma Preview

## Structure

`SignupProformaPreview`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrentProformaInvoice` | [`*models.ProformaInvoice`](../../doc/models/proforma-invoice.md) | Optional | - |
| `NextProformaInvoice` | [`*models.ProformaInvoice`](../../doc/models/proforma-invoice.md) | Optional | - |

## Example (as JSON)

```json
{
  "current_proforma_invoice": {
    "uid": "uid6",
    "site_id": 72,
    "customer_id": 184,
    "subscription_id": 0,
    "number": 132
  },
  "next_proforma_invoice": {
    "uid": "uid8",
    "site_id": 212,
    "customer_id": 68,
    "subscription_id": 140,
    "number": 16
  }
}
```

