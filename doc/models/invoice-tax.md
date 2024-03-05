
# Invoice Tax

## Structure

`InvoiceTax`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Title` | `*string` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `SourceType` | [`*models.ProformaInvoiceTaxSourceType`](../../doc/models/proforma-invoice-tax-source-type.md) | Optional | - |
| `SourceId` | `*int` | Optional | - |
| `Percentage` | `*string` | Optional | - |
| `TaxableAmount` | `*string` | Optional | - |
| `TaxAmount` | `*string` | Optional | - |
| `TransactionId` | `*int` | Optional | - |
| `LineItemBreakouts` | [`[]models.InvoiceTaxBreakout`](../../doc/models/invoice-tax-breakout.md) | Optional | - |
| `TaxComponentBreakouts` | [`[]models.InvoiceTaxComponentBreakout`](../../doc/models/invoice-tax-component-breakout.md) | Optional | - |

## Example (as JSON)

```json
{
  "uid": "uid8",
  "title": "title6",
  "description": "description2",
  "source_type": "Tax",
  "source_id": 164
}
```

