
# Billing Manifest Item

## Structure

`BillingManifestItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionType` | [`*models.LineItemTransactionType`](../../doc/models/line-item-transaction-type.md) | Optional | A handle for the line item transaction type |
| `Kind` | [`*models.BillingManifestLineItemKind`](../../doc/models/billing-manifest-line-item-kind.md) | Optional | A handle for the billing manifest line item kind |
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
  "transaction_type": "info_transaction",
  "kind": "baseline",
  "amount_in_cents": 216,
  "memo": "memo4",
  "discount_amount_in_cents": 236
}
```

