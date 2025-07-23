// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// OriginInvoice represents a OriginInvoice struct.
type OriginInvoice struct {
    // The UID of the invoice serving as an origin invoice.
    Uid                  *string                `json:"uid,omitempty"`
    // The number of the invoice serving as an origin invoice.
    Number               *string                `json:"number,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OriginInvoice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OriginInvoice) String() string {
    return fmt.Sprintf(
    	"OriginInvoice[Uid=%v, Number=%v, AdditionalProperties=%v]",
    	o.Uid, o.Number, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OriginInvoice.
// It customizes the JSON marshaling process for OriginInvoice objects.
func (o OriginInvoice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "uid", "number"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OriginInvoice object to a map representation for JSON marshaling.
func (o OriginInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Uid != nil {
        structMap["uid"] = o.Uid
    }
    if o.Number != nil {
        structMap["number"] = o.Number
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OriginInvoice.
// It customizes the JSON unmarshaling process for OriginInvoice objects.
func (o *OriginInvoice) UnmarshalJSON(input []byte) error {
    var temp tempOriginInvoice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "number")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Uid = temp.Uid
    o.Number = temp.Number
    return nil
}

// tempOriginInvoice is a temporary struct used for validating the fields of OriginInvoice.
type tempOriginInvoice  struct {
    Uid    *string `json:"uid,omitempty"`
    Number *string `json:"number,omitempty"`
}
