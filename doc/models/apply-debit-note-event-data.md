
# Apply Debit Note Event Data

Example schema for an `apply_debit_note` event

## Structure

`ApplyDebitNoteEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DebitNoteNumber` | `string` | Required | A unique, identifying string that appears on the debit note and in places it is referenced. |
| `DebitNoteUid` | `string` | Required | Unique identifier for the debit note. It is generated automatically by Chargify and has the prefix "db_" followed by alphanumeric characters. |
| `OriginalAmount` | `string` | Required | The full, original amount of the debit note. |
| `AppliedAmount` | `string` | Required | The amount of the debit note applied to invoice. |
| `Memo` | `models.Optional[string]` | Optional | The debit note memo. |
| `TransactionTime` | `models.Optional[time.Time]` | Optional | The time the debit note was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |

## Example (as JSON)

```json
{
  "debit_note_number": "debit_note_number0",
  "debit_note_uid": "debit_note_uid6",
  "original_amount": "original_amount4",
  "applied_amount": "applied_amount8",
  "memo": "memo4",
  "transaction_time": "2016-03-13T12:52:32.123Z"
}
```

