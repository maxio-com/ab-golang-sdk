
# Billing Manifest Item

## Structure

`BillingManifestItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionType` | `*string` | Optional | - |
| `Kind` | `*string` | Optional | - |
| `AmountInCents` | `*int64` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `DiscountAmountInCents` | `*int64` | Optional | - |
| `TaxableAmountInCents` | `*int64` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `ComponentHandle` | `*string` | Optional | - |
| `ComponentName` | `*string` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `ProductHandle` | `*string` | Optional | - |
| `ProductName` | `*string` | Optional | - |
| `PeriodRangeStart` | `*string` | Optional | - |
| `PeriodRangeEnd` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "transaction_type": "transaction_type2",
  "kind": "kind8",
  "amount_in_cents": 216,
  "memo": "memo4",
  "discount_amount_in_cents": 236
}
```

