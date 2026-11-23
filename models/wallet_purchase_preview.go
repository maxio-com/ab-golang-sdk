// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletPurchasePreview represents a WalletPurchasePreview struct.
// A price quote for a purchase, without charging anything.
type WalletPurchasePreview struct {
    PricePointId         int                    `json:"price_point_id"`
    // The three-letter currency code the `amount` is denominated in, matching the subscription's currency.
    Currency             string                 `json:"currency"`
    // The number of units the quote is for, as a decimal string.
    Quantity             string                 `json:"quantity"`
    // The price for `quantity` units at this price point, as a decimal string.
    Amount               string                 `json:"amount"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletPurchasePreview,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletPurchasePreview) String() string {
    return fmt.Sprintf(
    	"WalletPurchasePreview[PricePointId=%v, Currency=%v, Quantity=%v, Amount=%v, AdditionalProperties=%v]",
    	w.PricePointId, w.Currency, w.Quantity, w.Amount, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletPurchasePreview.
// It customizes the JSON marshaling process for WalletPurchasePreview objects.
func (w WalletPurchasePreview) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "price_point_id", "currency", "quantity", "amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletPurchasePreview object to a map representation for JSON marshaling.
func (w WalletPurchasePreview) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["price_point_id"] = w.PricePointId
    structMap["currency"] = w.Currency
    structMap["quantity"] = w.Quantity
    structMap["amount"] = w.Amount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletPurchasePreview.
// It customizes the JSON unmarshaling process for WalletPurchasePreview objects.
func (w *WalletPurchasePreview) UnmarshalJSON(input []byte) error {
    var temp tempWalletPurchasePreview
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_point_id", "currency", "quantity", "amount")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.PricePointId = *temp.PricePointId
    w.Currency = *temp.Currency
    w.Quantity = *temp.Quantity
    w.Amount = *temp.Amount
    return nil
}

// tempWalletPurchasePreview is a temporary struct used for validating the fields of WalletPurchasePreview.
type tempWalletPurchasePreview  struct {
    PricePointId *int    `json:"price_point_id"`
    Currency     *string `json:"currency"`
    Quantity     *string `json:"quantity"`
    Amount       *string `json:"amount"`
}

func (w *tempWalletPurchasePreview) validate() error {
    var errs []string
    if w.PricePointId == nil {
        errs = append(errs, "required field `price_point_id` is missing for type `Wallet Purchase Preview`")
    }
    if w.Currency == nil {
        errs = append(errs, "required field `currency` is missing for type `Wallet Purchase Preview`")
    }
    if w.Quantity == nil {
        errs = append(errs, "required field `quantity` is missing for type `Wallet Purchase Preview`")
    }
    if w.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Wallet Purchase Preview`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
