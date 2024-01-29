package models

import (
	"encoding/json"
)

// InvoiceAddress represents a InvoiceAddress struct.
type InvoiceAddress struct {
	Street  Optional[string] `json:"street"`
	Line2   Optional[string] `json:"line2"`
	City    Optional[string] `json:"city"`
	State   Optional[string] `json:"state"`
	Zip     Optional[string] `json:"zip"`
	Country Optional[string] `json:"country"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceAddress.
// It customizes the JSON marshaling process for InvoiceAddress objects.
func (i *InvoiceAddress) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceAddress object to a map representation for JSON marshaling.
func (i *InvoiceAddress) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.Street.IsValueSet() {
		structMap["street"] = i.Street.Value()
	}
	if i.Line2.IsValueSet() {
		structMap["line2"] = i.Line2.Value()
	}
	if i.City.IsValueSet() {
		structMap["city"] = i.City.Value()
	}
	if i.State.IsValueSet() {
		structMap["state"] = i.State.Value()
	}
	if i.Zip.IsValueSet() {
		structMap["zip"] = i.Zip.Value()
	}
	if i.Country.IsValueSet() {
		structMap["country"] = i.Country.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceAddress.
// It customizes the JSON unmarshaling process for InvoiceAddress objects.
func (i *InvoiceAddress) UnmarshalJSON(input []byte) error {
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

	i.Street = temp.Street
	i.Line2 = temp.Line2
	i.City = temp.City
	i.State = temp.State
	i.Zip = temp.Zip
	i.Country = temp.Country
	return nil
}
