
# Subscription Group Prepayment Response

## Structure

`SubscriptionGroupPrepaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `AmountInCents` | `*int64` | Optional | The amount in cents of the entry. |
| `EndingBalanceInCents` | `*int64` | Optional | The ending balance in cents of the account. |
| `EntryType` | [`*models.ServiceCreditType`](../../doc/models/service-credit-type.md) | Optional | The type of entry |
| `Memo` | `models.Optional[string]` | Optional | A memo attached to the entry. |

## Example (as JSON)

```json
{
  "id": 110,
  "amount_in_cents": 196,
  "ending_balance_in_cents": 236,
  "entry_type": "Credit",
  "memo": "memo2"
}
```

