
# Subscription Group Balances

## Structure

`SubscriptionGroupBalances`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayments` | [`*models.AccountBalance`](account-balance.md) | Optional | - |
| `ServiceCredits` | [`*models.AccountBalance`](account-balance.md) | Optional | - |
| `OpenInvoices` | [`*models.AccountBalance`](account-balance.md) | Optional | - |
| `PendingDiscounts` | [`*models.AccountBalance`](account-balance.md) | Optional | - |

## Example (as JSON)

```json
{
  "prepayments": {
    "balance_in_cents": 192
  },
  "service_credits": {
    "balance_in_cents": 84
  },
  "open_invoices": {
    "balance_in_cents": 40
  },
  "pending_discounts": {
    "balance_in_cents": 88
  }
}
```

