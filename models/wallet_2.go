// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Wallet2 represents a Wallet2 struct.
type Wallet2 struct {
    // A merchant-facing label for the wallet. Surrounding spaces are stripped; a blank or `null` value clears it.
    Name                 Optional[string]       `json:"name"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Wallet2,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w Wallet2) String() string {
    return fmt.Sprintf(
    	"Wallet2[Name=%v, AdditionalProperties=%v]",
    	w.Name, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Wallet2.
// It customizes the JSON marshaling process for Wallet2 objects.
func (w Wallet2) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the Wallet2 object to a map representation for JSON marshaling.
func (w Wallet2) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Name.IsValueSet() {
        if w.Name.Value() != nil {
            structMap["name"] = w.Name.Value()
        } else {
            structMap["name"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Wallet2.
// It customizes the JSON unmarshaling process for Wallet2 objects.
func (w *Wallet2) UnmarshalJSON(input []byte) error {
    var temp tempWallet2
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Name = temp.Name
    return nil
}

// tempWallet2 is a temporary struct used for validating the fields of Wallet2.
type tempWallet2  struct {
    Name Optional[string] `json:"name"`
}
