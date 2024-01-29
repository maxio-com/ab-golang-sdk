package models

import (
	"encoding/json"
)

// ShippingAddress represents a ShippingAddress struct.
type ShippingAddress struct {
	Street  Optional[string] `json:"street"`
	Line2   Optional[string] `json:"line2"`
	City    Optional[string] `json:"city"`
	State   Optional[string] `json:"state"`
	Zip     Optional[string] `json:"zip"`
	Country Optional[string] `json:"country"`
}

// MarshalJSON implements the json.Marshaler interface for ShippingAddress.
// It customizes the JSON marshaling process for ShippingAddress objects.
func (s *ShippingAddress) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the ShippingAddress object to a map representation for JSON marshaling.
func (s *ShippingAddress) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.Street.IsValueSet() {
		structMap["street"] = s.Street.Value()
	}
	if s.Line2.IsValueSet() {
		structMap["line2"] = s.Line2.Value()
	}
	if s.City.IsValueSet() {
		structMap["city"] = s.City.Value()
	}
	if s.State.IsValueSet() {
		structMap["state"] = s.State.Value()
	}
	if s.Zip.IsValueSet() {
		structMap["zip"] = s.Zip.Value()
	}
	if s.Country.IsValueSet() {
		structMap["country"] = s.Country.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ShippingAddress.
// It customizes the JSON unmarshaling process for ShippingAddress objects.
func (s *ShippingAddress) UnmarshalJSON(input []byte) error {
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

	s.Street = temp.Street
	s.Line2 = temp.Line2
	s.City = temp.City
	s.State = temp.State
	s.Zip = temp.Zip
	s.Country = temp.Country
	return nil
}
