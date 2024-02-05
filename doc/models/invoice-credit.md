
# Invoice Credit

## Structure

`InvoiceCredit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `CreditNoteNumber` | `*string` | Optional | - |
| `CreditNoteUid` | `*string` | Optional | - |
| `TransactionTime` | `*time.Time` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `OriginalAmount` | `*string` | Optional | - |
| `AppliedAmount` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "uid": "uid6",
  "credit_note_number": "credit_note_number0",
  "credit_note_uid": "credit_note_uid0",
  "transaction_time": "2016-03-13T12:52:32.123Z",
  "memo": "memo0"
}
```

