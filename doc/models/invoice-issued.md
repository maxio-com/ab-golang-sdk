
# Invoice Issued

## Structure

`InvoiceIssued`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `string` | Required | - |
| `Number` | `string` | Required | - |
| `Role` | `string` | Required | - |
| `DueDate` | `time.Time` | Required | - |
| `IssueDate` | `time.Time` | Required | - |
| `PaidDate` | `time.Time` | Required | - |
| `DueAmount` | `string` | Required | - |
| `PaidAmount` | `string` | Required | - |
| `TaxAmount` | `string` | Required | - |
| `RefundAmount` | `string` | Required | - |
| `TotalAmount` | `string` | Required | - |
| `StatusAmount` | `string` | Required | - |
| `ProductName` | `string` | Required | - |
| `ConsolidationLevel` | `string` | Required | - |
| `LineItems` | [`[]models.InvoiceLineItemEventData`](../../doc/models/invoice-line-item-event-data.md) | Required | - |

## Example (as JSON)

```json
{
  "uid": "uid4",
  "number": "number8",
  "role": "role2",
  "due_date": "2016-03-13T12:52:32.123Z",
  "issue_date": "2016-03-13T12:52:32.123Z",
  "paid_date": "2016-03-13T12:52:32.123Z",
  "due_amount": "due_amount6",
  "paid_amount": "paid_amount4",
  "tax_amount": "tax_amount2",
  "refund_amount": "refund_amount0",
  "total_amount": "total_amount0",
  "status_amount": "status_amount4",
  "product_name": "product_name0",
  "consolidation_level": "consolidation_level4",
  "line_items": [
    {
      "uid": "uid8",
      "title": "title4",
      "description": "description8",
      "quantity": 102,
      "quantity_delta": 204
    }
  ]
}
```

