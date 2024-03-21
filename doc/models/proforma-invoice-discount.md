
# Proforma Invoice Discount

## Structure

`ProformaInvoiceDiscount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Title` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Code` | `*string` | Optional | - |
| `SourceType` | [`*models.ProformaInvoiceDiscountSourceType`](../../doc/models/proforma-invoice-discount-source-type.md) | Optional | - |
| `DiscountType` | [`*models.InvoiceDiscountType`](../../doc/models/invoice-discount-type.md) | Optional | - |
| `EligibleAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `DiscountAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `LineItemBreakouts` | [`[]models.InvoiceDiscountBreakout`](../../doc/models/invoice-discount-breakout.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "uid": "uid2",
  "title": "title8",
  "code": "code0",
  "source_type": "Coupon",
  "discount_type": "percentage"
}
```

