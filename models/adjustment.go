// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// Adjustment represents a Adjustment struct.
type Adjustment struct {
    // Signed: positive adds units to the grant, negative removes them. Bounded below by the grant's `posted_remaining` (a correction can't drive it negative) and above by the grant's `original` (a correction can't credit past what was originally granted). An out-of-bounds value is rejected rather than clamped. Zero is rejected. Without a `grant_id` (an overdraft correction), the amount must be positive, can't exceed the usage the overdraft entry still records, and can't exceed the overdraft the wallet still owes. Encoded as a decimal string with at most 8 fractional digits. Send fractional values as strings: a JSON number with a decimal point is rejected with `422`, because a double-precision float may already have lost precision before the request arrives.
    Amount                    AdjustmentAmount       `json:"amount"`
    // Required. Lands on the resulting ledger row's `memo`.
    Reason                    string                 `json:"reason"`
    // The grant to correct. Must belong to this wallet. Required unless `source_transaction_id` names an overdraft entry (a `debit` with no `grant_id`, written when an `allow_negative` wallet consumes past its grants). For an overdraft entry, omit it: sending one returns `422`, and the correction credits the wallet's outstanding overdraft instead of any grant.
    GrantId                   Optional[int]          `json:"grant_id"`
    // (Optional) The `debit` row this correction is fixing. Must belong to this wallet, and must name a row of kind `debit`; grant and adjustment rows are append-only history, corrected by posting a further adjustment rather than by pointing at them. For a grant-attributed debit, it's recorded on the resulting ledger row for audit only: it does not change how `amount` is applied, and more than one adjustment may name the same row. For an overdraft entry, it decides the path: omit `grant_id`, and the correction is credit-only, bounded by the usage the entry still records and by the overdraft the wallet still owes.
    SourceTransactionId       Optional[int]          `json:"source_transaction_id"`
    // (Optional) An optimistic-concurrency check for a caller whose `amount` was derived from the grant's own remainder. Send the grant's `posted_remaining`, not its `remaining`: on a metered wallet read through [Read Wallet]($e/Wallets/readWallet), `remaining` also subtracts usage not yet posted, so it can differ from what the adjustment checks. If given and the grant's current `posted_remaining` no longer matches, the request is rejected with `409` and nothing is written. Compared numerically, so `"100"` and `"100.00000000"` both match a `posted_remaining` of `100`. Omit the field entirely to skip the check; sending it as an empty string is treated as a malformed value (`422`), not as an omission.
    ExpectedRemaining         Optional[string]       `json:"expected_remaining"`
    // (Optional) An optimistic-concurrency check for a caller whose `amount` was derived from a usage entry, sent together with `source_transaction_id`: the entry's `source_corrections_total` as the caller last read it (`"0"` if none). If the sum the server computes no longer matches, the request is rejected with `409` and nothing is written. This is what catches a correction posted between that read and this request. Ignored unless `source_transaction_id` is also given.
    ExpectedSourceCorrections Optional[string]       `json:"expected_source_corrections"`
    AdditionalProperties      map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Adjustment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a Adjustment) String() string {
    return fmt.Sprintf(
    	"Adjustment[Amount=%v, Reason=%v, GrantId=%v, SourceTransactionId=%v, ExpectedRemaining=%v, ExpectedSourceCorrections=%v, AdditionalProperties=%v]",
    	a.Amount, a.Reason, a.GrantId, a.SourceTransactionId, a.ExpectedRemaining, a.ExpectedSourceCorrections, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Adjustment.
// It customizes the JSON marshaling process for Adjustment objects.
func (a Adjustment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "amount", "reason", "grant_id", "source_transaction_id", "expected_remaining", "expected_source_corrections"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the Adjustment object to a map representation for JSON marshaling.
func (a Adjustment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["amount"] = a.Amount.toMap()
    structMap["reason"] = a.Reason
    if a.GrantId.IsValueSet() {
        if a.GrantId.Value() != nil {
            structMap["grant_id"] = a.GrantId.Value()
        } else {
            structMap["grant_id"] = nil
        }
    }
    if a.SourceTransactionId.IsValueSet() {
        if a.SourceTransactionId.Value() != nil {
            structMap["source_transaction_id"] = a.SourceTransactionId.Value()
        } else {
            structMap["source_transaction_id"] = nil
        }
    }
    if a.ExpectedRemaining.IsValueSet() {
        if a.ExpectedRemaining.Value() != nil {
            structMap["expected_remaining"] = a.ExpectedRemaining.Value()
        } else {
            structMap["expected_remaining"] = nil
        }
    }
    if a.ExpectedSourceCorrections.IsValueSet() {
        if a.ExpectedSourceCorrections.Value() != nil {
            structMap["expected_source_corrections"] = a.ExpectedSourceCorrections.Value()
        } else {
            structMap["expected_source_corrections"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Adjustment.
// It customizes the JSON unmarshaling process for Adjustment objects.
func (a *Adjustment) UnmarshalJSON(input []byte) error {
    var temp tempAdjustment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amount", "reason", "grant_id", "source_transaction_id", "expected_remaining", "expected_source_corrections")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Amount = *temp.Amount
    a.Reason = *temp.Reason
    a.GrantId = temp.GrantId
    a.SourceTransactionId = temp.SourceTransactionId
    a.ExpectedRemaining = temp.ExpectedRemaining
    a.ExpectedSourceCorrections = temp.ExpectedSourceCorrections
    return nil
}

// tempAdjustment is a temporary struct used for validating the fields of Adjustment.
type tempAdjustment  struct {
    Amount                    *AdjustmentAmount `json:"amount"`
    Reason                    *string           `json:"reason"`
    GrantId                   Optional[int]     `json:"grant_id"`
    SourceTransactionId       Optional[int]     `json:"source_transaction_id"`
    ExpectedRemaining         Optional[string]  `json:"expected_remaining"`
    ExpectedSourceCorrections Optional[string]  `json:"expected_source_corrections"`
}

func (a *tempAdjustment) validate() error {
    var errs []string
    if a.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Adjustment`")
    }
    if a.Reason == nil {
        errs = append(errs, "required field `reason` is missing for type `Adjustment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
