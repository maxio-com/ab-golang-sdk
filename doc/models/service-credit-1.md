
# Service Credit 1

## Structure

`ServiceCredit1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `AmountInCents` | `*int64` | Optional | The amount in cents of the entry |
| `EndingBalanceInCents` | `*int64` | Optional | The new balance for the credit account |
| `EntryType` | [`*models.ServiceCreditType`](../../doc/models/service-credit-type.md) | Optional | The type of entry |
| `Memo` | `*string` | Optional | The memo attached to the entry |
| `InvoiceUid` | `models.Optional[string]` | Optional | The invoice uid associated with the entry. Only present for debit entries |
| `RemainingBalanceInCents` | `*int64` | Optional | The remaining balance for the entry |
| `CreatedAt` | `*time.Time` | Optional | The date and time the entry was created |

## Example (as JSON)

```json
{
  "id": 174,
  "amount_in_cents": 4,
  "ending_balance_in_cents": 44,
  "entry_type": "Credit",
  "memo": "memo8"
}
```

