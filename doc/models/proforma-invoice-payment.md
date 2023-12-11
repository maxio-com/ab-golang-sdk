
# Proforma Invoice Payment

## Structure

`ProformaInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Memo` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `OriginalAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `AppliedAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Prepayment` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "memo": "memo4",
  "original_amount": "original_amount4",
  "applied_amount": "applied_amount8",
  "prepayment": false
}
```

