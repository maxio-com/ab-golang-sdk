
# Account Balance

## Structure

`AccountBalance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BalanceInCents` | `*int64` | Optional | The balance in cents. |
| `AutomaticBalanceInCents` | `models.Optional[int64]` | Optional | The automatic balance in cents. |
| `RemittanceBalanceInCents` | `models.Optional[int64]` | Optional | The remittance balance in cents. |

## Example (as JSON)

```json
{
  "balance_in_cents": 16,
  "automatic_balance_in_cents": 226,
  "remittance_balance_in_cents": 62
}
```

