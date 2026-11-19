// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// WalletGrant represents a WalletGrant struct.
// A funding tranche embedded on the wallet. `original` is the amount the grant was created with; `posted_remaining` is what debits and adjustments written to the ledger have left; `remaining` is what can still be spent. Debits drain grants earliest-expiration-first (grants with no expiration are drained last).
type WalletGrant struct {
    // The Advanced Billing id of the grant.
    Id                    *int                   `json:"id,omitempty"`
    // The memo the grant was created with, or `Grant
    Name                  *string                `json:"name,omitempty"`
    // The amount the grant was created with, encoded as a decimal string. Never changes after creation.
    Original              *string                `json:"original,omitempty"`
    // What the grant can still spend, encoded as a decimal string. On [Read Wallet]($e/Wallets/readWallet) for a metered wallet whose `balance` is the live figure, this is `posted_remaining` less the pending usage drawn from this grant in drain order: usage past a grant spills to the next one, and usage past every grant is an overdraft that belongs to no grant. It never goes below zero, and the grants' `remaining` add up to the wallet's `balance` whenever `balance` is zero or more. Everywhere else it equals `posted_remaining`.
    Remaining             *string                `json:"remaining,omitempty"`
    // The grant's remainder from written ledger entries alone, encoded as a decimal string. Always present. This is the figure an adjustment checks `expected_remaining` and its bounds against, so send it back, not `remaining`, when correcting a grant.
    PostedRemaining       *string                `json:"posted_remaining,omitempty"`
    // When this grant's unconsumed remainder is automatically written off. `null` for an evergreen grant.
    ExpiresAt             Optional[time.Time]    `json:"expires_at"`
    // The calendar day `expires_at` falls on, already resolved in the site's time zone.
    ExpiresOn             Optional[time.Time]    `json:"expires_on"`
    // When the expiry sweep actually wrote this grant's remainder off. Usually falls shortly after `expires_at`; use `expires_on` for the date the grant expired, not this field. `null` if the grant was never written off (including a grant simply spent to zero before its `expires_at`).
    ExpiredAt             Optional[time.Time]    `json:"expired_at"`
    // The calendar day `expired_at` falls on, resolved in the site's time zone.
    ExpiredOn             Optional[time.Time]    `json:"expired_on"`
    // How the grant was created. `recurring` comes from the wallet's automatic funding schedule (see the wallet's `recurring_funding` array). `manual` comes from a one-time grant posted to [Create Wallet Transaction]($e/Wallet%20Transactions/createWalletTransaction). `initial_balance` comes from a non-zero `initial_balance` at wallet creation. `allowance_increase` comes from raising a funding entitlement's allowance while it's already funding a wallet. `purchase` comes from [Create Wallet Purchase]($e/Wallet%20Purchases/createWalletPurchase).
    Origin                *Origin                `json:"origin,omitempty"`
    // For a `recurring` grant, the start of the funding cadence period this grant covers. `null` for every other origin. After a gap in funding (see `skipped_periods`), this is the period currently in progress, not the oldest period that was missed.
    GrantedForPeriodStart Optional[time.Time]    `json:"granted_for_period_start"`
    // The uid of the invoice that paid for this grant, for a grant with `origin` of `purchase`. `null` for every other origin.
    InvoiceUid            Optional[string]       `json:"invoice_uid"`
    CreatedAt             *time.Time             `json:"created_at,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletGrant,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletGrant) String() string {
    return fmt.Sprintf(
    	"WalletGrant[Id=%v, Name=%v, Original=%v, Remaining=%v, PostedRemaining=%v, ExpiresAt=%v, ExpiresOn=%v, ExpiredAt=%v, ExpiredOn=%v, Origin=%v, GrantedForPeriodStart=%v, InvoiceUid=%v, CreatedAt=%v, AdditionalProperties=%v]",
    	w.Id, w.Name, w.Original, w.Remaining, w.PostedRemaining, w.ExpiresAt, w.ExpiresOn, w.ExpiredAt, w.ExpiredOn, w.Origin, w.GrantedForPeriodStart, w.InvoiceUid, w.CreatedAt, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletGrant.
// It customizes the JSON marshaling process for WalletGrant objects.
func (w WalletGrant) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "id", "name", "original", "remaining", "posted_remaining", "expires_at", "expires_on", "expired_at", "expired_on", "origin", "granted_for_period_start", "invoice_uid", "created_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletGrant object to a map representation for JSON marshaling.
func (w WalletGrant) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.Name != nil {
        structMap["name"] = w.Name
    }
    if w.Original != nil {
        structMap["original"] = w.Original
    }
    if w.Remaining != nil {
        structMap["remaining"] = w.Remaining
    }
    if w.PostedRemaining != nil {
        structMap["posted_remaining"] = w.PostedRemaining
    }
    if w.ExpiresAt.IsValueSet() {
        var ExpiresAtVal *string = nil
        if w.ExpiresAt.Value() != nil {
            val := w.ExpiresAt.Value().Format(time.RFC3339)
            ExpiresAtVal = &val
        }
        if w.ExpiresAt.Value() != nil {
            structMap["expires_at"] = ExpiresAtVal
        } else {
            structMap["expires_at"] = nil
        }
    }
    if w.ExpiresOn.IsValueSet() {
        var ExpiresOnVal *string = nil
        if w.ExpiresOn.Value() != nil {
            val := w.ExpiresOn.Value().Format(DEFAULT_DATE)
            ExpiresOnVal = &val
        }
        if w.ExpiresOn.Value() != nil {
            structMap["expires_on"] = ExpiresOnVal
        } else {
            structMap["expires_on"] = nil
        }
    }
    if w.ExpiredAt.IsValueSet() {
        var ExpiredAtVal *string = nil
        if w.ExpiredAt.Value() != nil {
            val := w.ExpiredAt.Value().Format(time.RFC3339)
            ExpiredAtVal = &val
        }
        if w.ExpiredAt.Value() != nil {
            structMap["expired_at"] = ExpiredAtVal
        } else {
            structMap["expired_at"] = nil
        }
    }
    if w.ExpiredOn.IsValueSet() {
        var ExpiredOnVal *string = nil
        if w.ExpiredOn.Value() != nil {
            val := w.ExpiredOn.Value().Format(DEFAULT_DATE)
            ExpiredOnVal = &val
        }
        if w.ExpiredOn.Value() != nil {
            structMap["expired_on"] = ExpiredOnVal
        } else {
            structMap["expired_on"] = nil
        }
    }
    if w.Origin != nil {
        structMap["origin"] = w.Origin
    }
    if w.GrantedForPeriodStart.IsValueSet() {
        var GrantedForPeriodStartVal *string = nil
        if w.GrantedForPeriodStart.Value() != nil {
            val := w.GrantedForPeriodStart.Value().Format(time.RFC3339)
            GrantedForPeriodStartVal = &val
        }
        if w.GrantedForPeriodStart.Value() != nil {
            structMap["granted_for_period_start"] = GrantedForPeriodStartVal
        } else {
            structMap["granted_for_period_start"] = nil
        }
    }
    if w.InvoiceUid.IsValueSet() {
        if w.InvoiceUid.Value() != nil {
            structMap["invoice_uid"] = w.InvoiceUid.Value()
        } else {
            structMap["invoice_uid"] = nil
        }
    }
    if w.CreatedAt != nil {
        structMap["created_at"] = w.CreatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletGrant.
// It customizes the JSON unmarshaling process for WalletGrant objects.
func (w *WalletGrant) UnmarshalJSON(input []byte) error {
    var temp tempWalletGrant
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "name", "original", "remaining", "posted_remaining", "expires_at", "expires_on", "expired_at", "expired_on", "origin", "granted_for_period_start", "invoice_uid", "created_at")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Id = temp.Id
    w.Name = temp.Name
    w.Original = temp.Original
    w.Remaining = temp.Remaining
    w.PostedRemaining = temp.PostedRemaining
    w.ExpiresAt.ShouldSetValue(temp.ExpiresAt.IsValueSet())
    if temp.ExpiresAt.Value() != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, (*temp.ExpiresAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        w.ExpiresAt.SetValue(&ExpiresAtVal)
    }
    w.ExpiresOn.ShouldSetValue(temp.ExpiresOn.IsValueSet())
    if temp.ExpiresOn.Value() != nil {
        ExpiresOnVal, err := time.Parse(DEFAULT_DATE, (*temp.ExpiresOn.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse expires_on as % s format.", DEFAULT_DATE)
        }
        w.ExpiresOn.SetValue(&ExpiresOnVal)
    }
    w.ExpiredAt.ShouldSetValue(temp.ExpiredAt.IsValueSet())
    if temp.ExpiredAt.Value() != nil {
        ExpiredAtVal, err := time.Parse(time.RFC3339, (*temp.ExpiredAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse expired_at as % s format.", time.RFC3339)
        }
        w.ExpiredAt.SetValue(&ExpiredAtVal)
    }
    w.ExpiredOn.ShouldSetValue(temp.ExpiredOn.IsValueSet())
    if temp.ExpiredOn.Value() != nil {
        ExpiredOnVal, err := time.Parse(DEFAULT_DATE, (*temp.ExpiredOn.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse expired_on as % s format.", DEFAULT_DATE)
        }
        w.ExpiredOn.SetValue(&ExpiredOnVal)
    }
    w.Origin = temp.Origin
    w.GrantedForPeriodStart.ShouldSetValue(temp.GrantedForPeriodStart.IsValueSet())
    if temp.GrantedForPeriodStart.Value() != nil {
        GrantedForPeriodStartVal, err := time.Parse(time.RFC3339, (*temp.GrantedForPeriodStart.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse granted_for_period_start as % s format.", time.RFC3339)
        }
        w.GrantedForPeriodStart.SetValue(&GrantedForPeriodStartVal)
    }
    w.InvoiceUid = temp.InvoiceUid
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        w.CreatedAt = &CreatedAtVal
    }
    return nil
}

// tempWalletGrant is a temporary struct used for validating the fields of WalletGrant.
type tempWalletGrant  struct {
    Id                    *int             `json:"id,omitempty"`
    Name                  *string          `json:"name,omitempty"`
    Original              *string          `json:"original,omitempty"`
    Remaining             *string          `json:"remaining,omitempty"`
    PostedRemaining       *string          `json:"posted_remaining,omitempty"`
    ExpiresAt             Optional[string] `json:"expires_at"`
    ExpiresOn             Optional[string] `json:"expires_on"`
    ExpiredAt             Optional[string] `json:"expired_at"`
    ExpiredOn             Optional[string] `json:"expired_on"`
    Origin                *Origin          `json:"origin,omitempty"`
    GrantedForPeriodStart Optional[string] `json:"granted_for_period_start"`
    InvoiceUid            Optional[string] `json:"invoice_uid"`
    CreatedAt             *string          `json:"created_at,omitempty"`
}
