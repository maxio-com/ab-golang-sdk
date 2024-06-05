
# Paid Invoice

## Structure

`PaidInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InvoiceId` | `*string` | Optional | The uid of the paid invoice |
| `Status` | [`*models.InvoiceStatus`](../../doc/models/invoice-status.md) | Optional | The current status of the invoice. See [Invoice Statuses](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405078794253-Introduction-to-Invoices#invoice-statuses) for more. |
| `DueAmount` | `*string` | Optional | The remaining due amount on the invoice |
| `PaidAmount` | `*string` | Optional | The total amount paid on this invoice (including any prior payments) |

## Example (as JSON)

```json
{
  "invoice_id": "invoice_id6",
  "status": "draft",
  "due_amount": "due_amount8",
  "paid_amount": "paid_amount8"
}
```

