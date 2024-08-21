/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ListOffersResponse represents a ListOffersResponse struct.
type ListOffersResponse struct {
    Offers               []Offer        `json:"offers,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListOffersResponse.
// It customizes the JSON marshaling process for ListOffersResponse objects.
func (l ListOffersResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListOffersResponse object to a map representation for JSON marshaling.
func (l ListOffersResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Offers != nil {
        structMap["offers"] = l.Offers
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListOffersResponse.
// It customizes the JSON unmarshaling process for ListOffersResponse objects.
func (l *ListOffersResponse) UnmarshalJSON(input []byte) error {
    var temp tempListOffersResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "offers")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Offers = temp.Offers
    return nil
}

// tempListOffersResponse is a temporary struct used for validating the fields of ListOffersResponse.
type tempListOffersResponse  struct {
    Offers []Offer `json:"offers,omitempty"`
}
