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
		if i.Street.Value() != nil {
			structMap["street"] = i.Street.Value()
		} else {
			structMap["street"] = nil
		}
	}
	if i.Line2.IsValueSet() {
		if i.Line2.Value() != nil {
			structMap["line2"] = i.Line2.Value()
		} else {
			structMap["line2"] = nil
		}
	}
	if i.City.IsValueSet() {
		if i.City.Value() != nil {
			structMap["city"] = i.City.Value()
		} else {
			structMap["city"] = nil
		}
	}
	if i.State.IsValueSet() {
		if i.State.Value() != nil {
			structMap["state"] = i.State.Value()
		} else {
			structMap["state"] = nil
		}
	}
	if i.Zip.IsValueSet() {
		if i.Zip.Value() != nil {
			structMap["zip"] = i.Zip.Value()
		} else {
			structMap["zip"] = nil
		}
	}
	if i.Country.IsValueSet() {
		if i.Country.Value() != nil {
			structMap["country"] = i.Country.Value()
		} else {
			structMap["country"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceAddress.
// It customizes the JSON unmarshaling process for InvoiceAddress objects.
func (i *InvoiceAddress) UnmarshalJSON(input []byte) error {
	var temp invoiceAddress
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

// TODO
type invoiceAddress struct {
	Street  Optional[string] `json:"street"`
	Line2   Optional[string] `json:"line2"`
	City    Optional[string] `json:"city"`
	State   Optional[string] `json:"state"`
	Zip     Optional[string] `json:"zip"`
	Country Optional[string] `json:"country"`
}
