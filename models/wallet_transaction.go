// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// WalletTransaction represents a WalletTransaction struct.
// One append-only row on the wallet's ledger. Always appears bare inside a `wallet_transactions` array, never wrapped in its own key.
type WalletTransaction struct {
    // The Advanced Billing id of the ledger row.
    Id                     *int                   `json:"id,omitempty"`
    WalletId               *int                   `json:"wallet_id,omitempty"`
    // The kind of ledger row. Known values: `grant`, `debit`, `adjustment`, `initial_balance`, and `expiration`. More kinds are expected over time, so treat this as an open string and don't fail on one you don't recognize. Only `grant` and `debit` are accepted on [Create Wallet Transaction]($e/Wallet%20Transactions/createWalletTransaction). `adjustment` rows are written by [Create Wallet Adjustment]($e/Wallet%20Adjustments/createWalletAdjustment), `initial_balance` rows only by a seeded [Create Wallet]($e/Wallets/createWallet), and `expiration` rows only by the automatic expiry sweep that writes off an unconsumed grant once its `expires_at` passes.
    Kind                   *string                `json:"kind,omitempty"`
    // A decimal string, **signed** on this resource. Non-negative for grants, negative for debits and expirations, either sign for adjustments. A grant row written after an overdraft repayment carries only what was left after the repayment (possibly `"0.0"`), not the grant's `original`.
    Amount                 *string                `json:"amount,omitempty"`
    // The grant this row is attributed to. `null` for an overdraft row (a `kind: debit` row draining the below-zero portion of a balance, attributed to no grant), for an overdraft-repayment row (the `kind: grant` row that pays an existing negative balance back to zero, written before any remainder becomes a new grant), and for an overdraft correction (a `kind: adjustment` row posted against an overdraft entry, which credits no grant).
    GrantId                Optional[int]          `json:"grant_id"`
    // For an `adjustment` row posted against a specific usage entry, the id of the `debit` row it corrects. `null` for every other kind, and for adjustments made against a grant's balance directly.
    SourceTransactionId    Optional[int]          `json:"source_transaction_id"`
    // For a `debit` row only, the signed sum (as a decimal string) of every adjustment posted against it so far (`"0.0"` when none). `null` for every other kind. The row's own `amount` is never rewritten, so a caller deriving the entry's corrected quantity nets this out of `amount`, and sends this value back as `expected_source_corrections` on a further correction.
    SourceCorrectionsTotal Optional[string]       `json:"source_corrections_total"`
    // The wallet's posted balance immediately after this row was applied, as a decimal string. Compare it with the wallet's `posted_balance`, not its `balance`.
    BalanceAfter           *string                `json:"balance_after,omitempty"`
    // For a row that closed a metered usage window (an hourly close, a depletion, or an expiry write-off), the start of that window. `null` on every other row.
    MeteredFromAt          Optional[time.Time]    `json:"metered_from_at"`
    // The end of the usage window this row closed. `null` on every other row.
    MeteredThroughAt       Optional[time.Time]    `json:"metered_through_at"`
    // The meter-to-wallet-unit conversion rate applied to compute this row's `amount`, as a decimal string. `null` on every other row.
    ConversionRateApplied  Optional[string]       `json:"conversion_rate_applied"`
    // The uid of the invoice behind a purchase's rows: the grant-creating row and, when the purchase repaid an overdraft, the repayment row. `null` on every other row.
    InvoiceUid             Optional[string]       `json:"invoice_uid"`
    // Free-text note. For an `adjustment` row, carries the required `reason`.
    Memo                   Optional[string]       `json:"memo"`
    CreatedAt              *time.Time             `json:"created_at,omitempty"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletTransaction,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletTransaction) String() string {
    return fmt.Sprintf(
    	"WalletTransaction[Id=%v, WalletId=%v, Kind=%v, Amount=%v, GrantId=%v, SourceTransactionId=%v, SourceCorrectionsTotal=%v, BalanceAfter=%v, MeteredFromAt=%v, MeteredThroughAt=%v, ConversionRateApplied=%v, InvoiceUid=%v, Memo=%v, CreatedAt=%v, AdditionalProperties=%v]",
    	w.Id, w.WalletId, w.Kind, w.Amount, w.GrantId, w.SourceTransactionId, w.SourceCorrectionsTotal, w.BalanceAfter, w.MeteredFromAt, w.MeteredThroughAt, w.ConversionRateApplied, w.InvoiceUid, w.Memo, w.CreatedAt, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletTransaction.
// It customizes the JSON marshaling process for WalletTransaction objects.
func (w WalletTransaction) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "id", "wallet_id", "kind", "amount", "grant_id", "source_transaction_id", "source_corrections_total", "balance_after", "metered_from_at", "metered_through_at", "conversion_rate_applied", "invoice_uid", "memo", "created_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletTransaction object to a map representation for JSON marshaling.
func (w WalletTransaction) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.WalletId != nil {
        structMap["wallet_id"] = w.WalletId
    }
    if w.Kind != nil {
        structMap["kind"] = w.Kind
    }
    if w.Amount != nil {
        structMap["amount"] = w.Amount
    }
    if w.GrantId.IsValueSet() {
        if w.GrantId.Value() != nil {
            structMap["grant_id"] = w.GrantId.Value()
        } else {
            structMap["grant_id"] = nil
        }
    }
    if w.SourceTransactionId.IsValueSet() {
        if w.SourceTransactionId.Value() != nil {
            structMap["source_transaction_id"] = w.SourceTransactionId.Value()
        } else {
            structMap["source_transaction_id"] = nil
        }
    }
    if w.SourceCorrectionsTotal.IsValueSet() {
        if w.SourceCorrectionsTotal.Value() != nil {
            structMap["source_corrections_total"] = w.SourceCorrectionsTotal.Value()
        } else {
            structMap["source_corrections_total"] = nil
        }
    }
    if w.BalanceAfter != nil {
        structMap["balance_after"] = w.BalanceAfter
    }
    if w.MeteredFromAt.IsValueSet() {
        var MeteredFromAtVal *string = nil
        if w.MeteredFromAt.Value() != nil {
            val := w.MeteredFromAt.Value().Format(time.RFC3339)
            MeteredFromAtVal = &val
        }
        if w.MeteredFromAt.Value() != nil {
            structMap["metered_from_at"] = MeteredFromAtVal
        } else {
            structMap["metered_from_at"] = nil
        }
    }
    if w.MeteredThroughAt.IsValueSet() {
        var MeteredThroughAtVal *string = nil
        if w.MeteredThroughAt.Value() != nil {
            val := w.MeteredThroughAt.Value().Format(time.RFC3339)
            MeteredThroughAtVal = &val
        }
        if w.MeteredThroughAt.Value() != nil {
            structMap["metered_through_at"] = MeteredThroughAtVal
        } else {
            structMap["metered_through_at"] = nil
        }
    }
    if w.ConversionRateApplied.IsValueSet() {
        if w.ConversionRateApplied.Value() != nil {
            structMap["conversion_rate_applied"] = w.ConversionRateApplied.Value()
        } else {
            structMap["conversion_rate_applied"] = nil
        }
    }
    if w.InvoiceUid.IsValueSet() {
        if w.InvoiceUid.Value() != nil {
            structMap["invoice_uid"] = w.InvoiceUid.Value()
        } else {
            structMap["invoice_uid"] = nil
        }
    }
    if w.Memo.IsValueSet() {
        if w.Memo.Value() != nil {
            structMap["memo"] = w.Memo.Value()
        } else {
            structMap["memo"] = nil
        }
    }
    if w.CreatedAt != nil {
        structMap["created_at"] = w.CreatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletTransaction.
// It customizes the JSON unmarshaling process for WalletTransaction objects.
func (w *WalletTransaction) UnmarshalJSON(input []byte) error {
    var temp tempWalletTransaction
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "wallet_id", "kind", "amount", "grant_id", "source_transaction_id", "source_corrections_total", "balance_after", "metered_from_at", "metered_through_at", "conversion_rate_applied", "invoice_uid", "memo", "created_at")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Id = temp.Id
    w.WalletId = temp.WalletId
    w.Kind = temp.Kind
    w.Amount = temp.Amount
    w.GrantId = temp.GrantId
    w.SourceTransactionId = temp.SourceTransactionId
    w.SourceCorrectionsTotal = temp.SourceCorrectionsTotal
    w.BalanceAfter = temp.BalanceAfter
    w.MeteredFromAt.ShouldSetValue(temp.MeteredFromAt.IsValueSet())
    if temp.MeteredFromAt.Value() != nil {
        MeteredFromAtVal, err := time.Parse(time.RFC3339, (*temp.MeteredFromAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse metered_from_at as % s format.", time.RFC3339)
        }
        w.MeteredFromAt.SetValue(&MeteredFromAtVal)
    }
    w.MeteredThroughAt.ShouldSetValue(temp.MeteredThroughAt.IsValueSet())
    if temp.MeteredThroughAt.Value() != nil {
        MeteredThroughAtVal, err := time.Parse(time.RFC3339, (*temp.MeteredThroughAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse metered_through_at as % s format.", time.RFC3339)
        }
        w.MeteredThroughAt.SetValue(&MeteredThroughAtVal)
    }
    w.ConversionRateApplied = temp.ConversionRateApplied
    w.InvoiceUid = temp.InvoiceUid
    w.Memo = temp.Memo
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        w.CreatedAt = &CreatedAtVal
    }
    return nil
}

// tempWalletTransaction is a temporary struct used for validating the fields of WalletTransaction.
type tempWalletTransaction  struct {
    Id                     *int             `json:"id,omitempty"`
    WalletId               *int             `json:"wallet_id,omitempty"`
    Kind                   *string          `json:"kind,omitempty"`
    Amount                 *string          `json:"amount,omitempty"`
    GrantId                Optional[int]    `json:"grant_id"`
    SourceTransactionId    Optional[int]    `json:"source_transaction_id"`
    SourceCorrectionsTotal Optional[string] `json:"source_corrections_total"`
    BalanceAfter           *string          `json:"balance_after,omitempty"`
    MeteredFromAt          Optional[string] `json:"metered_from_at"`
    MeteredThroughAt       Optional[string] `json:"metered_through_at"`
    ConversionRateApplied  Optional[string] `json:"conversion_rate_applied"`
    InvoiceUid             Optional[string] `json:"invoice_uid"`
    Memo                   Optional[string] `json:"memo"`
    CreatedAt              *string          `json:"created_at,omitempty"`
}
