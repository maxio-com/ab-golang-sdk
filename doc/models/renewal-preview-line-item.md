
# Renewal Preview Line Item

## Structure

`RenewalPreviewLineItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionType` | [`*models.LineItemTransactionType`](../../doc/models/line-item-transaction-type.md) | Optional | A handle for the line item transaction type |
| `Kind` | [`*models.LineItemKind`](../../doc/models/line-item-kind.md) | Optional | A handle for the line item kind |
| `AmountInCents` | `*int64` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `DiscountAmountInCents` | `*int64` | Optional | - |
| `TaxableAmountInCents` | `*int64` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `ProductName` | `*string` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `ComponentHandle` | `*string` | Optional | - |
| `ComponentName` | `*string` | Optional | - |
| `ProductHandle` | `*string` | Optional | - |
| `PeriodRangeStart` | `*string` | Optional | - |
| `PeriodRangeEnd` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "transaction_type": "charge",
  "kind": "prepaid_usage_component",
  "amount_in_cents": 154,
  "memo": "memo0",
  "discount_amount_in_cents": 214
}
```

