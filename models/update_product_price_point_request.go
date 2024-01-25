package models

import (
	"encoding/json"
)

// UpdateProductPricePointRequest represents a UpdateProductPricePointRequest struct.
type UpdateProductPricePointRequest struct {
	PricePoint UpdateProductPricePoint `json:"price_point"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateProductPricePointRequest.
// It customizes the JSON marshaling process for UpdateProductPricePointRequest objects.
func (u *UpdateProductPricePointRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateProductPricePointRequest object to a map representation for JSON marshaling.
func (u *UpdateProductPricePointRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["price_point"] = u.PricePoint
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateProductPricePointRequest.
// It customizes the JSON unmarshaling process for UpdateProductPricePointRequest objects.
func (u *UpdateProductPricePointRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PricePoint UpdateProductPricePoint `json:"price_point"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.PricePoint = temp.PricePoint
	return nil
}
