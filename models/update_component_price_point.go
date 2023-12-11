package models

import (
	"encoding/json"
)

// UpdateComponentPricePoint represents a UpdateComponentPricePoint struct.
type UpdateComponentPricePoint struct {
	Name   *string       `json:"name,omitempty"`
	Prices []UpdatePrice `json:"prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentPricePoint.
// It customizes the JSON marshaling process for UpdateComponentPricePoint objects.
func (u *UpdateComponentPricePoint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentPricePoint object to a map representation for JSON marshaling.
func (u *UpdateComponentPricePoint) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.Name != nil {
		structMap["name"] = u.Name
	}
	if u.Prices != nil {
		structMap["prices"] = u.Prices
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentPricePoint.
// It customizes the JSON unmarshaling process for UpdateComponentPricePoint objects.
func (u *UpdateComponentPricePoint) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name   *string       `json:"name,omitempty"`
		Prices []UpdatePrice `json:"prices,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Name = temp.Name
	u.Prices = temp.Prices
	return nil
}
