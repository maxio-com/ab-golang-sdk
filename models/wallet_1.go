// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// Wallet1 represents a Wallet1 struct.
type Wallet1 struct {
    // The subscription that will own the wallet.
    SubscriptionId          int                     `json:"subscription_id"`
    // The [Wallet Product]($e/Wallet%20Products/createWalletProduct) that will denominate the wallet: its numeric id, or its handle prefixed with `handle:` (for example `handle:token-pack`). A bare handle with no prefix is looked up as an id and returns `404`. Immutable once the wallet is created.
    WalletProductId         Wallet1WalletProductId  `json:"wallet_product_id"`
    // (Optional) A merchant-facing label for the wallet. Surrounding spaces are stripped; a blank value is stored as `null`.
    Name                    Optional[string]        `json:"name"`
    // Defaults to `hard_stop`, currently the only supported value.
    DepletionPolicy         *string                 `json:"depletion_policy,omitempty"`
    // Defaults to `discard`.
    HardStopOverflow        *WalletHardStopOverflow `json:"hard_stop_overflow,omitempty"`
    // Pairs the wallet with a merchant-owned metering meter, so its balance factors in live usage (see the wallet resource's `meter` and `metered_through_at` fields). Only a `sum` or `count` meter can be paired. Requires metered wallet consumption to be enabled for the site, checked whenever this or `conversion_rate` is sent: `403` (plain text, not JSON) if it isn't. A meter id that doesn't resolve to a merchant-owned meter on this site, including a system-maintained shadow meter, is `404`. Immutable once the wallet is created.
    MeterId                 *int                    `json:"meter_id,omitempty"`
    // How many wallet units one meter unit is worth, as a decimal string with at most 8 fractional digits, greater than `0` and below 10^16. Send fractional values as strings: a JSON number with a decimal point is rejected with `422`, because a double-precision float may already have lost precision before the request arrives. Only meaningful with `meter_id`. Defaults to `"1.0"`. Immutable once the wallet is created.
    ConversionRate          *Wallet1ConversionRate  `json:"conversion_rate,omitempty"`
    // The price point used to price [Create Wallet Purchase]($e/Wallet%20Purchases/createWalletPurchase) calls against this wallet. Must belong to the Wallet Product, and can't be archived or a subscription-specific custom price point. Defaults to the Wallet Product's default price point. Immutable once the wallet is created.
    PricePointId            *int                    `json:"price_point_id,omitempty"`
    // (Optional) A free opening balance, not a purchase and not reflected on any invoice. When non-zero, wallet creation also writes an initial-balance grant and a matching `initial_balance` ledger row. Omitted or exactly `0` creates a zero-balance wallet, not an error. Encoded as a decimal string with at most 8 fractional digits, and must be non-negative and below 10^16. Send fractional values as strings: a JSON number with a decimal point is rejected with `422`, because a double-precision float may already have lost precision before the request arrives.
    InitialBalance          *Wallet1InitialBalance  `json:"initial_balance,omitempty"`
    // An expiration for the initial-balance grant. Only accepted together with a non-zero `initial_balance`, and must be in the future. Accepts either a bare calendar date (`YYYY-MM-DD`), resolved to the end of that day in the site's time zone, or a full timestamp. A calendar day that doesn't exist, such as `2027-02-30`, is rejected with `422` rather than rolled into the next month.
    InitialBalanceExpiresAt Optional[string]        `json:"initial_balance_expires_at"`
    AdditionalProperties    map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for Wallet1,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w Wallet1) String() string {
    return fmt.Sprintf(
    	"Wallet1[SubscriptionId=%v, WalletProductId=%v, Name=%v, DepletionPolicy=%v, HardStopOverflow=%v, MeterId=%v, ConversionRate=%v, PricePointId=%v, InitialBalance=%v, InitialBalanceExpiresAt=%v, AdditionalProperties=%v]",
    	w.SubscriptionId, w.WalletProductId, w.Name, w.DepletionPolicy, w.HardStopOverflow, w.MeterId, w.ConversionRate, w.PricePointId, w.InitialBalance, w.InitialBalanceExpiresAt, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Wallet1.
// It customizes the JSON marshaling process for Wallet1 objects.
func (w Wallet1) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "subscription_id", "wallet_product_id", "name", "depletion_policy", "hard_stop_overflow", "meter_id", "conversion_rate", "price_point_id", "initial_balance", "initial_balance_expires_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the Wallet1 object to a map representation for JSON marshaling.
func (w Wallet1) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["subscription_id"] = w.SubscriptionId
    structMap["wallet_product_id"] = w.WalletProductId.toMap()
    if w.Name.IsValueSet() {
        if w.Name.Value() != nil {
            structMap["name"] = w.Name.Value()
        } else {
            structMap["name"] = nil
        }
    }
    if w.DepletionPolicy != nil {
        structMap["depletion_policy"] = w.DepletionPolicy
    }
    if w.HardStopOverflow != nil {
        structMap["hard_stop_overflow"] = w.HardStopOverflow
    }
    if w.MeterId != nil {
        structMap["meter_id"] = w.MeterId
    }
    if w.ConversionRate != nil {
        structMap["conversion_rate"] = w.ConversionRate.toMap()
    }
    if w.PricePointId != nil {
        structMap["price_point_id"] = w.PricePointId
    }
    if w.InitialBalance != nil {
        structMap["initial_balance"] = w.InitialBalance.toMap()
    }
    if w.InitialBalanceExpiresAt.IsValueSet() {
        if w.InitialBalanceExpiresAt.Value() != nil {
            structMap["initial_balance_expires_at"] = w.InitialBalanceExpiresAt.Value()
        } else {
            structMap["initial_balance_expires_at"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Wallet1.
// It customizes the JSON unmarshaling process for Wallet1 objects.
func (w *Wallet1) UnmarshalJSON(input []byte) error {
    var temp tempWallet1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_id", "wallet_product_id", "name", "depletion_policy", "hard_stop_overflow", "meter_id", "conversion_rate", "price_point_id", "initial_balance", "initial_balance_expires_at")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.SubscriptionId = *temp.SubscriptionId
    w.WalletProductId = *temp.WalletProductId
    w.Name = temp.Name
    w.DepletionPolicy = temp.DepletionPolicy
    w.HardStopOverflow = temp.HardStopOverflow
    w.MeterId = temp.MeterId
    w.ConversionRate = temp.ConversionRate
    w.PricePointId = temp.PricePointId
    w.InitialBalance = temp.InitialBalance
    w.InitialBalanceExpiresAt = temp.InitialBalanceExpiresAt
    return nil
}

// tempWallet1 is a temporary struct used for validating the fields of Wallet1.
type tempWallet1  struct {
    SubscriptionId          *int                    `json:"subscription_id"`
    WalletProductId         *Wallet1WalletProductId `json:"wallet_product_id"`
    Name                    Optional[string]        `json:"name"`
    DepletionPolicy         *string                 `json:"depletion_policy,omitempty"`
    HardStopOverflow        *WalletHardStopOverflow `json:"hard_stop_overflow,omitempty"`
    MeterId                 *int                    `json:"meter_id,omitempty"`
    ConversionRate          *Wallet1ConversionRate  `json:"conversion_rate,omitempty"`
    PricePointId            *int                    `json:"price_point_id,omitempty"`
    InitialBalance          *Wallet1InitialBalance  `json:"initial_balance,omitempty"`
    InitialBalanceExpiresAt Optional[string]        `json:"initial_balance_expires_at"`
}

func (w *tempWallet1) validate() error {
    var errs []string
    if w.SubscriptionId == nil {
        errs = append(errs, "required field `subscription_id` is missing for type `Wallet1`")
    }
    if w.WalletProductId == nil {
        errs = append(errs, "required field `wallet_product_id` is missing for type `Wallet1`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
