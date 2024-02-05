package models

import (
    "encoding/json"
)

// BillingAddress represents a BillingAddress struct.
type BillingAddress struct {
    Street  Optional[string] `json:"street"`
    Line2   Optional[string] `json:"line2"`
    City    Optional[string] `json:"city"`
    State   Optional[string] `json:"state"`
    Zip     Optional[string] `json:"zip"`
    Country Optional[string] `json:"country"`
}

// MarshalJSON implements the json.Marshaler interface for BillingAddress.
// It customizes the JSON marshaling process for BillingAddress objects.
func (b *BillingAddress) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BillingAddress object to a map representation for JSON marshaling.
func (b *BillingAddress) toMap() map[string]any {
    structMap := make(map[string]any)
    if b.Street.IsValueSet() {
        structMap["street"] = b.Street.Value()
    }
    if b.Line2.IsValueSet() {
        structMap["line2"] = b.Line2.Value()
    }
    if b.City.IsValueSet() {
        structMap["city"] = b.City.Value()
    }
    if b.State.IsValueSet() {
        structMap["state"] = b.State.Value()
    }
    if b.Zip.IsValueSet() {
        structMap["zip"] = b.Zip.Value()
    }
    if b.Country.IsValueSet() {
        structMap["country"] = b.Country.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BillingAddress.
// It customizes the JSON unmarshaling process for BillingAddress objects.
func (b *BillingAddress) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Street  Optional[string] `json:"street"`
        Line2   Optional[string] `json:"line2"`
        City    Optional[string] `json:"city"`
        State   Optional[string] `json:"state"`
        Zip     Optional[string] `json:"zip"`
        Country Optional[string] `json:"country"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.Street = temp.Street
    b.Line2 = temp.Line2
    b.City = temp.City
    b.State = temp.State
    b.Zip = temp.Zip
    b.Country = temp.Country
    return nil
}
