
# Proforma Invoice Tax

## Structure

`ProformaInvoiceTax`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Title` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `SourceType` | [`*models.ProformaInvoiceTaxSourceType`](../../doc/models/proforma-invoice-tax-source-type.md) | Optional | - |
| `Percentage` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `TaxableAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `TaxAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `LineItemBreakouts` | [`[]models.InvoiceTaxBreakout`](../../doc/models/invoice-tax-breakout.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "uid": "uid8",
  "title": "title4",
  "source_type": "Tax",
  "percentage": "percentage6",
  "taxable_amount": "taxable_amount2"
}
```

