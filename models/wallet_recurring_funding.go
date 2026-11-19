// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// WalletRecurringFunding represents a WalletRecurringFunding struct.
// Describes automatic funding of the wallet from a usage-limit entitlement. Set up by configuring `wallet_behavior` on the Feature behind that entitlement, not by any parameter on the wallet itself. This field is an array as of its first release. The array currently holds zero or one element, and is expected to hold several once a wallet can be funded by more than one entitlement.
type WalletRecurringFunding struct {
    // The Advanced Billing id of the recurring funding schedule.
    Id                   *int                          `json:"id,omitempty"`
    // The usage-limit entitlement this schedule reads its allowance and cadence from.
    EntitlementId        *int                          `json:"entitlement_id,omitempty"`
    // The entitlement's current allowance, read fresh at request time rather than stored on the wallet. Editing the entitlement changes this immediately, with no separate sync step. `"0"` means the entitlement is configured to grant nothing, a valid, working configuration and not an error. `null` when `status` is `stopped`, since there is no longer an entitlement to read.
    Allowance            Optional[string]              `json:"allowance"`
    // When the next automatic grant is due. Grants are applied by a background process shortly after this moment passes, not exactly on it, so a wallet read taken in that window still shows the pre-grant balance. `null` when `status` is `stopped`, and also `null` once the owning subscription has terminated, even while `status` is still `active`: at its next due moment the schedule is stopped with reason `subscription_terminated`.
    NextGrantAt          Optional[time.Time]           `json:"next_grant_at"`
    // The calendar day `next_grant_at` falls on, resolved in the site's time zone. `null` under the same conditions as `next_grant_at`.
    NextGrantOn          Optional[time.Time]           `json:"next_grant_on"`
    // Cumulative count of due moments that passed without a grant being made (for example, during an outage). A gap in funding is visible here without cross-referencing the ledger.
    SkippedPeriods       *int                          `json:"skipped_periods,omitempty"`
    // `active`: refilling normally. `erroring`: the schedule is still due and retrying, but the last attempt couldn't be written (see `reason`); correcting the underlying entitlement resumes funding on the next run with no action needed on the wallet itself. `stopped`: funding has ended for good (see `reason`); units already granted stay on the wallet and remain spendable, but no further automatic grants will be made.
    Status               *WalletRecurringFundingStatus `json:"status,omitempty"`
    // Set whenever `status` isn't `active`. Accompanying `stopped`: `entitlement_missing` (an ordinary product or component change removed the entitlement), `wallet_archived` (the wallet was archived), `subscription_terminated` (the subscription reached a terminal state; funding resumes if the subscription returns to a granting state, for example on reactivation, without back-granting the periods it missed), `feature_removed` (the subscription no longer has this Feature), `wallet_behavior_disabled` (the Feature no longer has `wallet_behavior` configured), `superseded_by_replacement` (a later entitlement now funds this wallet instead), or `subscription_prepaid` (the subscription switched to prepaid, which cannot carry an active wallet). Accompanying `erroring`: `allowance_not_representable` (the entitlement's current allowance doesn't fit the wallet's amount encoding) or `balance_at_capacity` (granting it would push the balance past the maximum supported value).
    Reason               Optional[Reason]              `json:"reason"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for WalletRecurringFunding,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletRecurringFunding) String() string {
    return fmt.Sprintf(
    	"WalletRecurringFunding[Id=%v, EntitlementId=%v, Allowance=%v, NextGrantAt=%v, NextGrantOn=%v, SkippedPeriods=%v, Status=%v, Reason=%v, AdditionalProperties=%v]",
    	w.Id, w.EntitlementId, w.Allowance, w.NextGrantAt, w.NextGrantOn, w.SkippedPeriods, w.Status, w.Reason, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletRecurringFunding.
// It customizes the JSON marshaling process for WalletRecurringFunding objects.
func (w WalletRecurringFunding) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "id", "entitlement_id", "allowance", "next_grant_at", "next_grant_on", "skipped_periods", "status", "reason"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletRecurringFunding object to a map representation for JSON marshaling.
func (w WalletRecurringFunding) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.EntitlementId != nil {
        structMap["entitlement_id"] = w.EntitlementId
    }
    if w.Allowance.IsValueSet() {
        if w.Allowance.Value() != nil {
            structMap["allowance"] = w.Allowance.Value()
        } else {
            structMap["allowance"] = nil
        }
    }
    if w.NextGrantAt.IsValueSet() {
        var NextGrantAtVal *string = nil
        if w.NextGrantAt.Value() != nil {
            val := w.NextGrantAt.Value().Format(time.RFC3339)
            NextGrantAtVal = &val
        }
        if w.NextGrantAt.Value() != nil {
            structMap["next_grant_at"] = NextGrantAtVal
        } else {
            structMap["next_grant_at"] = nil
        }
    }
    if w.NextGrantOn.IsValueSet() {
        var NextGrantOnVal *string = nil
        if w.NextGrantOn.Value() != nil {
            val := w.NextGrantOn.Value().Format(DEFAULT_DATE)
            NextGrantOnVal = &val
        }
        if w.NextGrantOn.Value() != nil {
            structMap["next_grant_on"] = NextGrantOnVal
        } else {
            structMap["next_grant_on"] = nil
        }
    }
    if w.SkippedPeriods != nil {
        structMap["skipped_periods"] = w.SkippedPeriods
    }
    if w.Status != nil {
        structMap["status"] = w.Status
    }
    if w.Reason.IsValueSet() {
        if w.Reason.Value() != nil {
            structMap["reason"] = w.Reason.Value()
        } else {
            structMap["reason"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletRecurringFunding.
// It customizes the JSON unmarshaling process for WalletRecurringFunding objects.
func (w *WalletRecurringFunding) UnmarshalJSON(input []byte) error {
    var temp tempWalletRecurringFunding
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "entitlement_id", "allowance", "next_grant_at", "next_grant_on", "skipped_periods", "status", "reason")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Id = temp.Id
    w.EntitlementId = temp.EntitlementId
    w.Allowance = temp.Allowance
    w.NextGrantAt.ShouldSetValue(temp.NextGrantAt.IsValueSet())
    if temp.NextGrantAt.Value() != nil {
        NextGrantAtVal, err := time.Parse(time.RFC3339, (*temp.NextGrantAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse next_grant_at as % s format.", time.RFC3339)
        }
        w.NextGrantAt.SetValue(&NextGrantAtVal)
    }
    w.NextGrantOn.ShouldSetValue(temp.NextGrantOn.IsValueSet())
    if temp.NextGrantOn.Value() != nil {
        NextGrantOnVal, err := time.Parse(DEFAULT_DATE, (*temp.NextGrantOn.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse next_grant_on as % s format.", DEFAULT_DATE)
        }
        w.NextGrantOn.SetValue(&NextGrantOnVal)
    }
    w.SkippedPeriods = temp.SkippedPeriods
    w.Status = temp.Status
    w.Reason = temp.Reason
    return nil
}

// tempWalletRecurringFunding is a temporary struct used for validating the fields of WalletRecurringFunding.
type tempWalletRecurringFunding  struct {
    Id             *int                          `json:"id,omitempty"`
    EntitlementId  *int                          `json:"entitlement_id,omitempty"`
    Allowance      Optional[string]              `json:"allowance"`
    NextGrantAt    Optional[string]              `json:"next_grant_at"`
    NextGrantOn    Optional[string]              `json:"next_grant_on"`
    SkippedPeriods *int                          `json:"skipped_periods,omitempty"`
    Status         *WalletRecurringFundingStatus `json:"status,omitempty"`
    Reason         Optional[Reason]              `json:"reason"`
}
