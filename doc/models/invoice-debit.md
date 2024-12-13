
# Invoice Debit

## Structure

`InvoiceDebit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `DebitNoteNumber` | `*string` | Optional | - |
| `DebitNoteUid` | `*string` | Optional | - |
| `Role` | [`*models.DebitNoteRole`](../../doc/models/debit-note-role.md) | Optional | The role of the debit note. |
| `TransactionTime` | `*time.Time` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `OriginalAmount` | `*string` | Optional | - |
| `AppliedAmount` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "uid": "uid2",
  "debit_note_number": "debit_note_number2",
  "debit_note_uid": "debit_note_uid8",
  "role": "chargeback",
  "transaction_time": "2016-03-13T12:52:32.123Z"
}
```

