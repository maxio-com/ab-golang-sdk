
# Customer Payer Change

## Structure

`CustomerPayerChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Before` | [`models.InvoicePayerChange`](../../doc/models/invoice-payer-change.md) | Required | - |
| `After` | [`models.InvoicePayerChange`](../../doc/models/invoice-payer-change.md) | Required | - |

## Example (as JSON)

```json
{
  "before": {
    "first_name": "first_name0",
    "last_name": "last_name8",
    "organization": "organization4",
    "email": "email6"
  },
  "after": {
    "first_name": "first_name2",
    "last_name": "last_name0",
    "organization": "organization4",
    "email": "email4"
  }
}
```

