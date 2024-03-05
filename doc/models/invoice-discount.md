
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
| `SourceType` | [`*models.InvoiceDiscountSourceType`](../../doc/models/invoice-discount-source-type.md) | Optional | - |
| `SourceId` | `*int` | Optional | - |
| `DiscountType` | [`*models.InvoiceDiscountType`](../../doc/models/invoice-discount-type.md) | Optional | - |
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
  "source_type": "Coupon"
}
```

