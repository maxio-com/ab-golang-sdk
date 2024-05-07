package models

import (
    "encoding/json"
)

// ShippingAddress represents a ShippingAddress struct.
type ShippingAddress struct {
    Street               Optional[string] `json:"street"`
    Line2                Optional[string] `json:"line2"`
    City                 Optional[string] `json:"city"`
    State                Optional[string] `json:"state"`
    Zip                  Optional[string] `json:"zip"`
    Country              Optional[string] `json:"country"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ShippingAddress.
// It customizes the JSON marshaling process for ShippingAddress objects.
func (s ShippingAddress) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the ShippingAddress object to a map representation for JSON marshaling.
func (s ShippingAddress) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Street.IsValueSet() {
        if s.Street.Value() != nil {
            structMap["street"] = s.Street.Value()
        } else {
            structMap["street"] = nil
        }
    }
    if s.Line2.IsValueSet() {
        if s.Line2.Value() != nil {
            structMap["line2"] = s.Line2.Value()
        } else {
            structMap["line2"] = nil
        }
    }
    if s.City.IsValueSet() {
        if s.City.Value() != nil {
            structMap["city"] = s.City.Value()
        } else {
            structMap["city"] = nil
        }
    }
    if s.State.IsValueSet() {
        if s.State.Value() != nil {
            structMap["state"] = s.State.Value()
        } else {
            structMap["state"] = nil
        }
    }
    if s.Zip.IsValueSet() {
        if s.Zip.Value() != nil {
            structMap["zip"] = s.Zip.Value()
        } else {
            structMap["zip"] = nil
        }
    }
    if s.Country.IsValueSet() {
        if s.Country.Value() != nil {
            structMap["country"] = s.Country.Value()
        } else {
            structMap["country"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ShippingAddress.
// It customizes the JSON unmarshaling process for ShippingAddress objects.
func (s *ShippingAddress) UnmarshalJSON(input []byte) error {
    var temp shippingAddress
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "street", "line2", "city", "state", "zip", "country")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Street = temp.Street
    s.Line2 = temp.Line2
    s.City = temp.City
    s.State = temp.State
    s.Zip = temp.Zip
    s.Country = temp.Country
    return nil
}

// shippingAddress is a temporary struct used for validating the fields of ShippingAddress.
type shippingAddress  struct {
    Street  Optional[string] `json:"street"`
    Line2   Optional[string] `json:"line2"`
    City    Optional[string] `json:"city"`
    State   Optional[string] `json:"state"`
    Zip     Optional[string] `json:"zip"`
    Country Optional[string] `json:"country"`
}
