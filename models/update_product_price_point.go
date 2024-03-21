package models

import (
	"encoding/json"
)

// UpdateProductPricePoint represents a UpdateProductPricePoint struct.
type UpdateProductPricePoint struct {
	Handle       *string `json:"handle,omitempty"`
	PriceInCents *int64  `json:"price_in_cents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateProductPricePoint.
// It customizes the JSON marshaling process for UpdateProductPricePoint objects.
func (u *UpdateProductPricePoint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateProductPricePoint object to a map representation for JSON marshaling.
func (u *UpdateProductPricePoint) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.Handle != nil {
		structMap["handle"] = u.Handle
	}
	if u.PriceInCents != nil {
		structMap["price_in_cents"] = u.PriceInCents
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateProductPricePoint.
// It customizes the JSON unmarshaling process for UpdateProductPricePoint objects.
func (u *UpdateProductPricePoint) UnmarshalJSON(input []byte) error {
	var temp updateProductPricePoint
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	u.Handle = temp.Handle
	u.PriceInCents = temp.PriceInCents
	return nil
}

// TODO
type updateProductPricePoint struct {
	Handle       *string `json:"handle,omitempty"`
	PriceInCents *int64  `json:"price_in_cents,omitempty"`
}
