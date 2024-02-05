
# Proforma Invoice Tax

## Structure

`ProformaInvoiceTax`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Title` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `SourceType` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Percentage` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `TaxableAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `TaxAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `LineItemBreakouts` | [`[]models.ProformaInvoiceTaxBreakout`](../../doc/models/proforma-invoice-tax-breakout.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "uid": "uid8",
  "title": "title4",
  "source_type": "source_type8",
  "percentage": "percentage6",
  "taxable_amount": "taxable_amount2"
}
```

