// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// WalletProductActiveWallet represents a WalletProductActiveWallet struct.
// One active wallet drawing on this Wallet Product.
type WalletProductActiveWallet struct {
    // The Advanced Billing id of the wallet.
    Id                   *int                   `json:"id,omitempty"`
    SubscriptionId       *int                   `json:"subscription_id,omitempty"`
    // The owning customer's organization name, falling back to their full name, then to `"Customer #<id>"`. `"Customer"` (with no id) when the wallet's subscription has no customer at all.
    CompanyName          *string                `json:"company_name,omitempty"`
    // See the same field on [Wallet Product Usage Template]($m/Wallet%20Product%20Usage%20Template). The same caveats apply.
    PurchaseMechanism    Optional[string]       `json:"purchase_mechanism"`
    PricePointName       Optional[string]       `json:"price_point_name"`
    Currency             *string                `json:"currency,omitempty"`
    // `false` when the wallet's Feature is archived.
    Active               *bool                  `json:"active,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProductActiveWallet,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProductActiveWallet) String() string {
    return fmt.Sprintf(
    	"WalletProductActiveWallet[Id=%v, SubscriptionId=%v, CompanyName=%v, PurchaseMechanism=%v, PricePointName=%v, Currency=%v, Active=%v, AdditionalProperties=%v]",
    	w.Id, w.SubscriptionId, w.CompanyName, w.PurchaseMechanism, w.PricePointName, w.Currency, w.Active, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProductActiveWallet.
// It customizes the JSON marshaling process for WalletProductActiveWallet objects.
func (w WalletProductActiveWallet) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "id", "subscription_id", "company_name", "purchase_mechanism", "price_point_name", "currency", "active"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProductActiveWallet object to a map representation for JSON marshaling.
func (w WalletProductActiveWallet) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.SubscriptionId != nil {
        structMap["subscription_id"] = w.SubscriptionId
    }
    if w.CompanyName != nil {
        structMap["company_name"] = w.CompanyName
    }
    if w.PurchaseMechanism.IsValueSet() {
        if w.PurchaseMechanism.Value() != nil {
            structMap["purchase_mechanism"] = w.PurchaseMechanism.Value()
        } else {
            structMap["purchase_mechanism"] = nil
        }
    }
    if w.PricePointName.IsValueSet() {
        if w.PricePointName.Value() != nil {
            structMap["price_point_name"] = w.PricePointName.Value()
        } else {
            structMap["price_point_name"] = nil
        }
    }
    if w.Currency != nil {
        structMap["currency"] = w.Currency
    }
    if w.Active != nil {
        structMap["active"] = w.Active
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProductActiveWallet.
// It customizes the JSON unmarshaling process for WalletProductActiveWallet objects.
func (w *WalletProductActiveWallet) UnmarshalJSON(input []byte) error {
    var temp tempWalletProductActiveWallet
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "subscription_id", "company_name", "purchase_mechanism", "price_point_name", "currency", "active")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Id = temp.Id
    w.SubscriptionId = temp.SubscriptionId
    w.CompanyName = temp.CompanyName
    w.PurchaseMechanism = temp.PurchaseMechanism
    w.PricePointName = temp.PricePointName
    w.Currency = temp.Currency
    w.Active = temp.Active
    return nil
}

// tempWalletProductActiveWallet is a temporary struct used for validating the fields of WalletProductActiveWallet.
type tempWalletProductActiveWallet  struct {
    Id                *int             `json:"id,omitempty"`
    SubscriptionId    *int             `json:"subscription_id,omitempty"`
    CompanyName       *string          `json:"company_name,omitempty"`
    PurchaseMechanism Optional[string] `json:"purchase_mechanism"`
    PricePointName    Optional[string] `json:"price_point_name"`
    Currency          *string          `json:"currency,omitempty"`
    Active            *bool            `json:"active,omitempty"`
}
