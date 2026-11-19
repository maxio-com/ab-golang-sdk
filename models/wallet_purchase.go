// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// WalletPurchase represents a WalletPurchase struct.
// A completed unit purchase.
type WalletPurchase struct {
    // The id of the grant this purchase created. When the wallet was overdrawn by at least the purchased quantity, the units repay the debt and the grant is still created, with `remaining` of `0` and `original` equal to the purchased quantity.
    Id                   int                    `json:"id"`
    // The price point the purchase was priced at.
    PricePointId         int                    `json:"price_point_id"`
    // The uid of the invoice created to collect payment for this purchase.
    InvoiceUid           string                 `json:"invoice_uid"`
    // The number of units purchased, as a decimal string.
    Quantity             string                 `json:"quantity"`
    // When the purchased units expire, following the funding entitlement's grant expiration policy. `null` when the wallet has no funding entitlement or its entitlements disagree on an expiration policy.
    ExpiresAt            *time.Time             `json:"expires_at"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletPurchase,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletPurchase) String() string {
    return fmt.Sprintf(
    	"WalletPurchase[Id=%v, PricePointId=%v, InvoiceUid=%v, Quantity=%v, ExpiresAt=%v, AdditionalProperties=%v]",
    	w.Id, w.PricePointId, w.InvoiceUid, w.Quantity, w.ExpiresAt, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletPurchase.
// It customizes the JSON marshaling process for WalletPurchase objects.
func (w WalletPurchase) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "id", "price_point_id", "invoice_uid", "quantity", "expires_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletPurchase object to a map representation for JSON marshaling.
func (w WalletPurchase) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["id"] = w.Id
    structMap["price_point_id"] = w.PricePointId
    structMap["invoice_uid"] = w.InvoiceUid
    structMap["quantity"] = w.Quantity
    if w.ExpiresAt != nil {
        structMap["expires_at"] = w.ExpiresAt.Format(time.RFC3339)
    } else {
        structMap["expires_at"] = nil
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletPurchase.
// It customizes the JSON unmarshaling process for WalletPurchase objects.
func (w *WalletPurchase) UnmarshalJSON(input []byte) error {
    var temp tempWalletPurchase
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "price_point_id", "invoice_uid", "quantity", "expires_at")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Id = *temp.Id
    w.PricePointId = *temp.PricePointId
    w.InvoiceUid = *temp.InvoiceUid
    w.Quantity = *temp.Quantity
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        w.ExpiresAt = &ExpiresAtVal
    }
    return nil
}

// tempWalletPurchase is a temporary struct used for validating the fields of WalletPurchase.
type tempWalletPurchase  struct {
    Id           *int    `json:"id"`
    PricePointId *int    `json:"price_point_id"`
    InvoiceUid   *string `json:"invoice_uid"`
    Quantity     *string `json:"quantity"`
    ExpiresAt    *string `json:"expires_at"`
}

func (w *tempWalletPurchase) validate() error {
    var errs []string
    if w.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Wallet Purchase`")
    }
    if w.PricePointId == nil {
        errs = append(errs, "required field `price_point_id` is missing for type `Wallet Purchase`")
    }
    if w.InvoiceUid == nil {
        errs = append(errs, "required field `invoice_uid` is missing for type `Wallet Purchase`")
    }
    if w.Quantity == nil {
        errs = append(errs, "required field `quantity` is missing for type `Wallet Purchase`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
