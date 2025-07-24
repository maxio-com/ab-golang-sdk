// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// OfferResponse represents a OfferResponse struct.
type OfferResponse struct {
    Offer                *Offer                 `json:"offer,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OfferResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OfferResponse) String() string {
    return fmt.Sprintf(
    	"OfferResponse[Offer=%v, AdditionalProperties=%v]",
    	o.Offer, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OfferResponse.
// It customizes the JSON marshaling process for OfferResponse objects.
func (o OfferResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "offer"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OfferResponse object to a map representation for JSON marshaling.
func (o OfferResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Offer != nil {
        structMap["offer"] = o.Offer.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OfferResponse.
// It customizes the JSON unmarshaling process for OfferResponse objects.
func (o *OfferResponse) UnmarshalJSON(input []byte) error {
    var temp tempOfferResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "offer")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Offer = temp.Offer
    return nil
}

// tempOfferResponse is a temporary struct used for validating the fields of OfferResponse.
type tempOfferResponse  struct {
    Offer *Offer `json:"offer,omitempty"`
}
