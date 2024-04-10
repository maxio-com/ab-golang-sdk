
# Prepayment Account Balance Changed

## Structure

`PrepaymentAccountBalanceChanged`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `string` | Required | - |
| `PrepaymentAccountBalanceInCents` | `int64` | Required | - |
| `PrepaymentBalanceChangeInCents` | `int64` | Required | - |
| `CurrencyCode` | `string` | Required | - |

## Example (as JSON)

```json
{
  "reason": "reason4",
  "prepayment_account_balance_in_cents": 182,
  "prepayment_balance_change_in_cents": 206,
  "currency_code": "currency_code4"
}
```

