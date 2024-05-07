
# Billing Manifest

## Structure

`BillingManifest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LineItems` | [`[]models.BillingManifestItem`](../../doc/models/billing-manifest-item.md) | Optional | - |
| `TotalInCents` | `*int64` | Optional | - |
| `TotalDiscountInCents` | `*int64` | Optional | - |
| `TotalTaxInCents` | `*int64` | Optional | - |
| `SubtotalInCents` | `*int64` | Optional | - |
| `StartDate` | `models.Optional[time.Time]` | Optional | - |
| `EndDate` | `models.Optional[time.Time]` | Optional | - |
| `PeriodType` | `models.Optional[string]` | Optional | - |
| `ExistingBalanceInCents` | `*int64` | Optional | - |

## Example (as JSON)

```json
{
  "line_items": [
    {
      "transaction_type": "credit",
      "kind": "component",
      "amount_in_cents": 24,
      "memo": "memo2",
      "discount_amount_in_cents": 172
    },
    {
      "transaction_type": "credit",
      "kind": "component",
      "amount_in_cents": 24,
      "memo": "memo2",
      "discount_amount_in_cents": 172
    },
    {
      "transaction_type": "credit",
      "kind": "component",
      "amount_in_cents": 24,
      "memo": "memo2",
      "discount_amount_in_cents": 172
    }
  ],
  "total_in_cents": 192,
  "total_discount_in_cents": 178,
  "total_tax_in_cents": 172,
  "subtotal_in_cents": 48
}
```

