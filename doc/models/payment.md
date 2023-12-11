
# Payment

## Structure

`Payment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InvoiceUid` | `*string` | Optional | The uid of the paid invoice |
| `Status` | [`*models.StatusEnum`](status-enum.md) | Optional | The current status of the invoice. See [Invoice Statuses](https://chargify.zendesk.com/hc/en-us/articles/4407737494171#line-item-breakdowns) for more. |
| `DueAmount` | `*string` | Optional | The remaining due amount on the invoice |
| `PaidAmount` | `*string` | Optional | The total amount paid on this invoice (including any prior payments) |

## Example (as JSON)

```json
{
  "invoice_uid": "invoice_uid8",
  "status": "voided",
  "due_amount": "due_amount0",
  "paid_amount": "paid_amount0"
}
```

