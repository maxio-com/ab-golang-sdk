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
		structMap["street"] = o.Street.Value()
	}
	if o.Line2.IsValueSet() {
		structMap["line2"] = o.Line2.Value()
	}
	if o.City.IsValueSet() {
		structMap["city"] = o.City.Value()
	}
	if o.State.IsValueSet() {
		structMap["state"] = o.State.Value()
	}
	if o.Zip.IsValueSet() {
		structMap["zip"] = o.Zip.Value()
	}
	if o.Country.IsValueSet() {
		structMap["country"] = o.Country.Value()
	}
	if o.Name.IsValueSet() {
		structMap["name"] = o.Name.Value()
	}
	if o.Phone.IsValueSet() {
		structMap["phone"] = o.Phone.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrganizationAddress.
// It customizes the JSON unmarshaling process for OrganizationAddress objects.
func (o *OrganizationAddress) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Street  Optional[string] `json:"street"`
		Line2   Optional[string] `json:"line2"`
		City    Optional[string] `json:"city"`
		State   Optional[string] `json:"state"`
		Zip     Optional[string] `json:"zip"`
		Country Optional[string] `json:"country"`
		Name    Optional[string] `json:"name"`
		Phone   Optional[string] `json:"phone"`
	}{}
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
