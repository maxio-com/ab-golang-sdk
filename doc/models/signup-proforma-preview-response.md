
# Signup Proforma Preview Response

## Structure

`SignupProformaPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProformaInvoicePreview` | [`models.SignupProformaPreview`](../../doc/models/signup-proforma-preview.md) | Required | - |

## Example (as JSON)

```json
{
  "proforma_invoice_preview": {
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
}
```

