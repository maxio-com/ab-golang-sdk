/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// OfferResponse represents a OfferResponse struct.
type OfferResponse struct {
    Offer                *Offer         `json:"offer,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OfferResponse.
// It customizes the JSON marshaling process for OfferResponse objects.
func (o OfferResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OfferResponse object to a map representation for JSON marshaling.
func (o OfferResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "offer")
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
