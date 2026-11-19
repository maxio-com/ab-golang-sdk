
# Adjustment

## Structure

`Adjustment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | [`models.AdjustmentAmount`](../../doc/models/containers/adjustment-amount.md) | Required | This is a container for one-of cases. |
| `Reason` | `string` | Required | Required. Lands on the resulting ledger row's `memo`.<br><br>**Constraints**: *Maximum Length*: `255` |
| `GrantId` | `models.Optional[int]` | Optional | The grant to correct. Must belong to this wallet. Required unless `source_transaction_id` names an overdraft entry (a `debit` with no `grant_id`, written when an `allow_negative` wallet consumes past its grants). For an overdraft entry, omit it: sending one returns `422`, and the correction credits the wallet's outstanding overdraft instead of any grant. |
| `SourceTransactionId` | `models.Optional[int]` | Optional | (Optional) The `debit` row this correction is fixing. Must belong to this wallet, and must name a row of kind `debit`; grant and adjustment rows are append-only history, corrected by posting a further adjustment rather than by pointing at them. For a grant-attributed debit, it's recorded on the resulting ledger row for audit only: it does not change how `amount` is applied, and more than one adjustment may name the same row. For an overdraft entry, it decides the path: omit `grant_id`, and the correction is credit-only, bounded by the usage the entry still records and by the overdraft the wallet still owes. |
| `ExpectedRemaining` | `models.Optional[string]` | Optional | (Optional) An optimistic-concurrency check for a caller whose `amount` was derived from the grant's own remainder. Send the grant's `posted_remaining`, not its `remaining`: on a metered wallet read through [Read Wallet](../../doc/controllers/wallets.md#read-wallet), `remaining` also subtracts usage not yet posted, so it can differ from what the adjustment checks. If given and the grant's current `posted_remaining` no longer matches, the request is rejected with `409` and nothing is written. Compared numerically, so `"100"` and `"100.00000000"` both match a `posted_remaining` of `100`. Omit the field entirely to skip the check; sending it as an empty string is treated as a malformed value (`422`), not as an omission. |
| `ExpectedSourceCorrections` | `models.Optional[string]` | Optional | (Optional) An optimistic-concurrency check for a caller whose `amount` was derived from a usage entry, sent together with `source_transaction_id`: the entry's `source_corrections_total` as the caller last read it (`"0"` if none). If the sum the server computes no longer matches, the request is rejected with `409` and nothing is written. This is what catches a correction posted between that read and this request. Ignored unless `source_transaction_id` is also given. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    adjustment := models.Adjustment{
        Amount:                    models.AdjustmentAmountContainer.FromString("String7"),
        Reason:                    "reason0",
        GrantId:                   models.NewOptional(models.ToPointer(104)),
        SourceTransactionId:       models.NewOptional(models.ToPointer(126)),
        ExpectedRemaining:         models.NewOptional(models.ToPointer("expected_remaining6")),
        ExpectedSourceCorrections: models.NewOptional(models.ToPointer("expected_source_corrections8")),
    }

}
```

