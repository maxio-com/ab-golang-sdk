// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletsListResponse represents a WalletsListResponse struct.
type WalletsListResponse struct {
    Wallets              []Wallet               `json:"wallets"`
    Meta                 ListPublicKeysMeta     `json:"meta"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletsListResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletsListResponse) String() string {
    return fmt.Sprintf(
    	"WalletsListResponse[Wallets=%v, Meta=%v, AdditionalProperties=%v]",
    	w.Wallets, w.Meta, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletsListResponse.
// It customizes the JSON marshaling process for WalletsListResponse objects.
func (w WalletsListResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "wallets", "meta"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletsListResponse object to a map representation for JSON marshaling.
func (w WalletsListResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["wallets"] = w.Wallets
    structMap["meta"] = w.Meta.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletsListResponse.
// It customizes the JSON unmarshaling process for WalletsListResponse objects.
func (w *WalletsListResponse) UnmarshalJSON(input []byte) error {
    var temp tempWalletsListResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "wallets", "meta")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Wallets = *temp.Wallets
    w.Meta = *temp.Meta
    return nil
}

// tempWalletsListResponse is a temporary struct used for validating the fields of WalletsListResponse.
type tempWalletsListResponse  struct {
    Wallets *[]Wallet           `json:"wallets"`
    Meta    *ListPublicKeysMeta `json:"meta"`
}

func (w *tempWalletsListResponse) validate() error {
    var errs []string
    if w.Wallets == nil {
        errs = append(errs, "required field `wallets` is missing for type `Wallets List Response`")
    }
    if w.Meta == nil {
        errs = append(errs, "required field `meta` is missing for type `Wallets List Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
