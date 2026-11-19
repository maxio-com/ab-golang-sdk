
# Wallet Transaction

One append-only row on the wallet's ledger. Always appears bare inside a `wallet_transactions` array, never wrapped in its own key.

## Structure

`WalletTransaction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Advanced Billing id of the ledger row. |
| `WalletId` | `*int` | Optional | - |
| `Kind` | `*string` | Optional | The kind of ledger row. Known values: `grant`, `debit`, `adjustment`, `initial_balance`, and `expiration`. More kinds are expected over time, so treat this as an open string and don't fail on one you don't recognize. Only `grant` and `debit` are accepted on [Create Wallet Transaction](../../doc/controllers/wallet-transactions.md#create-wallet-transaction). `adjustment` rows are written by [Create Wallet Adjustment](../../doc/controllers/wallet-adjustments.md#create-wallet-adjustment), `initial_balance` rows only by a seeded [Create Wallet](../../doc/controllers/wallets.md#create-wallet), and `expiration` rows only by the automatic expiry sweep that writes off an unconsumed grant once its `expires_at` passes. |
| `Amount` | `*string` | Optional | A decimal string, **signed** on this resource. Non-negative for grants, negative for debits and expirations, either sign for adjustments. A grant row written after an overdraft repayment carries only what was left after the repayment (possibly `"0.0"`), not the grant's `original`. |
| `GrantId` | `models.Optional[int]` | Optional | The grant this row is attributed to. `null` for an overdraft row (a `kind: debit` row draining the below-zero portion of a balance, attributed to no grant), for an overdraft-repayment row (the `kind: grant` row that pays an existing negative balance back to zero, written before any remainder becomes a new grant), and for an overdraft correction (a `kind: adjustment` row posted against an overdraft entry, which credits no grant). |
| `SourceTransactionId` | `models.Optional[int]` | Optional | For an `adjustment` row posted against a specific usage entry, the id of the `debit` row it corrects. `null` for every other kind, and for adjustments made against a grant's balance directly. |
| `SourceCorrectionsTotal` | `models.Optional[string]` | Optional | For a `debit` row only, the signed sum (as a decimal string) of every adjustment posted against it so far (`"0.0"` when none). `null` for every other kind. The row's own `amount` is never rewritten, so a caller deriving the entry's corrected quantity nets this out of `amount`, and sends this value back as `expected_source_corrections` on a further correction. |
| `BalanceAfter` | `*string` | Optional | The wallet's posted balance immediately after this row was applied, as a decimal string. Compare it with the wallet's `posted_balance`, not its `balance`. |
| `MeteredFromAt` | `models.Optional[time.Time]` | Optional | For a row that closed a metered usage window (an hourly close, a depletion, or an expiry write-off), the start of that window. `null` on every other row. |
| `MeteredThroughAt` | `models.Optional[time.Time]` | Optional | The end of the usage window this row closed. `null` on every other row. |
| `ConversionRateApplied` | `models.Optional[string]` | Optional | The meter-to-wallet-unit conversion rate applied to compute this row's `amount`, as a decimal string. `null` on every other row. |
| `InvoiceUid` | `models.Optional[string]` | Optional | The uid of the invoice behind a purchase's rows: the grant-creating row and, when the purchase repaid an overdraft, the repayment row. `null` on every other row. |
| `Memo` | `models.Optional[string]` | Optional | Free-text note. For an `adjustment` row, carries the required `reason`. |
| `CreatedAt` | `*time.Time` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletTransaction := models.WalletTransaction{
        Id:                     models.ToPointer(178),
        WalletId:               models.ToPointer(26),
        Kind:                   models.ToPointer("grant"),
        Amount:                 models.ToPointer("-30.0"),
        GrantId:                models.NewOptional(models.ToPointer(198)),
    }

}
```

