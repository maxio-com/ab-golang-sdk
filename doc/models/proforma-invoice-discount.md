
# Proforma Invoice Discount

## Structure

`ProformaInvoiceDiscount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Title` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `SourceType` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `DiscountType` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `EligibleAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `DiscountAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `LineItemBreakouts` | [`[]models.ProformaInvoiceDiscountBreakout`](proforma-invoice-discount-breakout.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "title": "title8",
  "source_type": "source_type2",
  "discount_type": "discount_type0",
  "eligible_amount": "eligible_amount4",
  "discount_amount": "discount_amount6"
}
```

