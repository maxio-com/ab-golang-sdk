// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// WalletProductPricePointSummary represents a WalletProductPricePointSummary struct.
// A summary of the Wallet Product's default price point, for rendering a catalog listing without a separate price points call.
type WalletProductPricePointSummary struct {
    // Number of archived price points on this Wallet Product.
    ArchivedCount        *int                   `json:"archived_count,omitempty"`
    // The default price point's pricing scheme. `null` if the Wallet Product has no default price point.
    PricingScheme        Optional[string]       `json:"pricing_scheme"`
    // Number of price brackets on the default price point. A `per_unit` scheme always stores exactly one synthesized bracket, so this is `1` for `per_unit`, not `0`.
    TierCount            *int                   `json:"tier_count,omitempty"`
    // The unit price of the default price point's first bracket, as a decimal string. `null` if there are no brackets.
    EntryPrice           Optional[string]       `json:"entry_price"`
    // The unit price of the default price point's last bracket, as a decimal string. `null` if there are no brackets.
    FinalPrice           Optional[string]       `json:"final_price"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProductPricePointSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProductPricePointSummary) String() string {
    return fmt.Sprintf(
    	"WalletProductPricePointSummary[ArchivedCount=%v, PricingScheme=%v, TierCount=%v, EntryPrice=%v, FinalPrice=%v, AdditionalProperties=%v]",
    	w.ArchivedCount, w.PricingScheme, w.TierCount, w.EntryPrice, w.FinalPrice, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProductPricePointSummary.
// It customizes the JSON marshaling process for WalletProductPricePointSummary objects.
func (w WalletProductPricePointSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "archived_count", "pricing_scheme", "tier_count", "entry_price", "final_price"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProductPricePointSummary object to a map representation for JSON marshaling.
func (w WalletProductPricePointSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.ArchivedCount != nil {
        structMap["archived_count"] = w.ArchivedCount
    }
    if w.PricingScheme.IsValueSet() {
        if w.PricingScheme.Value() != nil {
            structMap["pricing_scheme"] = w.PricingScheme.Value()
        } else {
            structMap["pricing_scheme"] = nil
        }
    }
    if w.TierCount != nil {
        structMap["tier_count"] = w.TierCount
    }
    if w.EntryPrice.IsValueSet() {
        if w.EntryPrice.Value() != nil {
            structMap["entry_price"] = w.EntryPrice.Value()
        } else {
            structMap["entry_price"] = nil
        }
    }
    if w.FinalPrice.IsValueSet() {
        if w.FinalPrice.Value() != nil {
            structMap["final_price"] = w.FinalPrice.Value()
        } else {
            structMap["final_price"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProductPricePointSummary.
// It customizes the JSON unmarshaling process for WalletProductPricePointSummary objects.
func (w *WalletProductPricePointSummary) UnmarshalJSON(input []byte) error {
    var temp tempWalletProductPricePointSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "archived_count", "pricing_scheme", "tier_count", "entry_price", "final_price")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.ArchivedCount = temp.ArchivedCount
    w.PricingScheme = temp.PricingScheme
    w.TierCount = temp.TierCount
    w.EntryPrice = temp.EntryPrice
    w.FinalPrice = temp.FinalPrice
    return nil
}

// tempWalletProductPricePointSummary is a temporary struct used for validating the fields of WalletProductPricePointSummary.
type tempWalletProductPricePointSummary  struct {
    ArchivedCount *int             `json:"archived_count,omitempty"`
    PricingScheme Optional[string] `json:"pricing_scheme"`
    TierCount     *int             `json:"tier_count,omitempty"`
    EntryPrice    Optional[string] `json:"entry_price"`
    FinalPrice    Optional[string] `json:"final_price"`
}
