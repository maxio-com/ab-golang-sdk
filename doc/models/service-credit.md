
# Service Credit

## Structure

`ServiceCredit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `AmountInCents` | `*int64` | Optional | The amount in cents of the entry |
| `EndingBalanceInCents` | `*int64` | Optional | The new balance for the credit account |
| `EntryType` | [`*models.ServiceCreditType`](../../doc/models/service-credit-type.md) | Optional | The type of entry |
| `Memo` | `*string` | Optional | The memo attached to the entry |

## Example (as JSON)

```json
{
  "id": 216,
  "amount_in_cents": 210,
  "ending_balance_in_cents": 86,
  "entry_type": "Credit",
  "memo": "memo2"
}
```

