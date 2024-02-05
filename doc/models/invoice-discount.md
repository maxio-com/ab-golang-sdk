
# Invoice Discount

## Structure

`InvoiceDiscount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Title` | `*string` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `Code` | `*string` | Optional | - |
| `SourceType` | `*string` | Optional | - |
| `SourceId` | `*int` | Optional | - |
| `DiscountType` | `*string` | Optional | - |
| `Percentage` | `*string` | Optional | - |
| `EligibleAmount` | `*string` | Optional | - |
| `DiscountAmount` | `*string` | Optional | - |
| `TransactionId` | `*int` | Optional | - |
| `LineItemBreakouts` | [`[]models.InvoiceDiscountBreakout`](../../doc/models/invoice-discount-breakout.md) | Optional | - |

## Example (as JSON)

```json
{
  "uid": "uid0",
  "title": "title4",
  "description": "description0",
  "code": "code8",
  "source_type": "source_type0"
}
```

