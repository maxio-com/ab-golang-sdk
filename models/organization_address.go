package models

import (
	"encoding/json"
)

// OrganizationAddress represents a OrganizationAddress struct.
type OrganizationAddress struct {
	Street  Optional[string] `json:"street"`
	Line2   Optional[string] `json:"line2"`
	City    Optional[string] `json:"city"`
	State   Optional[string] `json:"state"`
	Zip     Optional[string] `json:"zip"`
	Country Optional[string] `json:"country"`
	Name    Optional[string] `json:"name"`
	Phone   Optional[string] `json:"phone"`
}

// MarshalJSON implements the json.Marshaler interface for OrganizationAddress.
// It customizes the JSON marshaling process for OrganizationAddress objects.
func (o *OrganizationAddress) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrganizationAddress object to a map representation for JSON marshaling.
func (o *OrganizationAddress) toMap() map[string]any {
	structMap := make(map[string]any)
	if o.Street.IsValueSet() {
		if o.Street.Value() != nil {
			structMap["street"] = o.Street.Value()
		} else {
			structMap["street"] = nil
		}
	}
	if o.Line2.IsValueSet() {
		if o.Line2.Value() != nil {
			structMap["line2"] = o.Line2.Value()
		} else {
			structMap["line2"] = nil
		}
	}
	if o.City.IsValueSet() {
		if o.City.Value() != nil {
			structMap["city"] = o.City.Value()
		} else {
			structMap["city"] = nil
		}
	}
	if o.State.IsValueSet() {
		if o.State.Value() != nil {
			structMap["state"] = o.State.Value()
		} else {
			structMap["state"] = nil
		}
	}
	if o.Zip.IsValueSet() {
		if o.Zip.Value() != nil {
			structMap["zip"] = o.Zip.Value()
		} else {
			structMap["zip"] = nil
		}
	}
	if o.Country.IsValueSet() {
		if o.Country.Value() != nil {
			structMap["country"] = o.Country.Value()
		} else {
			structMap["country"] = nil
		}
	}
	if o.Name.IsValueSet() {
		if o.Name.Value() != nil {
			structMap["name"] = o.Name.Value()
		} else {
			structMap["name"] = nil
		}
	}
	if o.Phone.IsValueSet() {
		if o.Phone.Value() != nil {
			structMap["phone"] = o.Phone.Value()
		} else {
			structMap["phone"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrganizationAddress.
// It customizes the JSON unmarshaling process for OrganizationAddress objects.
func (o *OrganizationAddress) UnmarshalJSON(input []byte) error {
	var temp organizationAddress
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	o.Street = temp.Street
	o.Line2 = temp.Line2
	o.City = temp.City
	o.State = temp.State
	o.Zip = temp.Zip
	o.Country = temp.Country
	o.Name = temp.Name
	o.Phone = temp.Phone
	return nil
}

// TODO
type organizationAddress struct {
	Street  Optional[string] `json:"street"`
	Line2   Optional[string] `json:"line2"`
	City    Optional[string] `json:"city"`
	State   Optional[string] `json:"state"`
	Zip     Optional[string] `json:"zip"`
	Country Optional[string] `json:"country"`
	Name    Optional[string] `json:"name"`
	Phone   Optional[string] `json:"phone"`
}
