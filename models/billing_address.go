package models

import (
    "encoding/json"
)

// BillingAddress represents a BillingAddress struct.
type BillingAddress struct {
    Street               Optional[string] `json:"street"`
    Line2                Optional[string] `json:"line2"`
    City                 Optional[string] `json:"city"`
    State                Optional[string] `json:"state"`
    Zip                  Optional[string] `json:"zip"`
    Country              Optional[string] `json:"country"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BillingAddress.
// It customizes the JSON marshaling process for BillingAddress objects.
func (b BillingAddress) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BillingAddress object to a map representation for JSON marshaling.
func (b BillingAddress) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Street.IsValueSet() {
        if b.Street.Value() != nil {
            structMap["street"] = b.Street.Value()
        } else {
            structMap["street"] = nil
        }
    }
    if b.Line2.IsValueSet() {
        if b.Line2.Value() != nil {
            structMap["line2"] = b.Line2.Value()
        } else {
            structMap["line2"] = nil
        }
    }
    if b.City.IsValueSet() {
        if b.City.Value() != nil {
            structMap["city"] = b.City.Value()
        } else {
            structMap["city"] = nil
        }
    }
    if b.State.IsValueSet() {
        if b.State.Value() != nil {
            structMap["state"] = b.State.Value()
        } else {
            structMap["state"] = nil
        }
    }
    if b.Zip.IsValueSet() {
        if b.Zip.Value() != nil {
            structMap["zip"] = b.Zip.Value()
        } else {
            structMap["zip"] = nil
        }
    }
    if b.Country.IsValueSet() {
        if b.Country.Value() != nil {
            structMap["country"] = b.Country.Value()
        } else {
            structMap["country"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BillingAddress.
// It customizes the JSON unmarshaling process for BillingAddress objects.
func (b *BillingAddress) UnmarshalJSON(input []byte) error {
    var temp billingAddress
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "street", "line2", "city", "state", "zip", "country")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.Street = temp.Street
    b.Line2 = temp.Line2
    b.City = temp.City
    b.State = temp.State
    b.Zip = temp.Zip
    b.Country = temp.Country
    return nil
}

// TODO
type billingAddress  struct {
    Street  Optional[string] `json:"street"`
    Line2   Optional[string] `json:"line2"`
    City    Optional[string] `json:"city"`
    State   Optional[string] `json:"state"`
    Zip     Optional[string] `json:"zip"`
    Country Optional[string] `json:"country"`
}
