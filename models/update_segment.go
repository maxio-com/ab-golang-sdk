package models

import (
	"encoding/json"
)

// UpdateSegment represents a UpdateSegment struct.
type UpdateSegment struct {
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme PricingSchemeEnum            `json:"pricing_scheme"`
	Prices        []CreateOrUpdateSegmentPrice `json:"prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSegment.
// It customizes the JSON marshaling process for UpdateSegment objects.
func (u *UpdateSegment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSegment object to a map representation for JSON marshaling.
func (u *UpdateSegment) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["pricing_scheme"] = u.PricingScheme
	if u.Prices != nil {
		structMap["prices"] = u.Prices
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSegment.
// It customizes the JSON unmarshaling process for UpdateSegment objects.
func (u *UpdateSegment) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PricingScheme PricingSchemeEnum            `json:"pricing_scheme"`
		Prices        []CreateOrUpdateSegmentPrice `json:"prices,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.PricingScheme = temp.PricingScheme
	u.Prices = temp.Prices
	return nil
}
