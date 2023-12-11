package models

import (
	"encoding/json"
)

// CreateProductFamilyRequest represents a CreateProductFamilyRequest struct.
type CreateProductFamilyRequest struct {
	ProductFamily CreateProductFamily `json:"product_family"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductFamilyRequest.
// It customizes the JSON marshaling process for CreateProductFamilyRequest objects.
func (c *CreateProductFamilyRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateProductFamilyRequest object to a map representation for JSON marshaling.
func (c *CreateProductFamilyRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["product_family"] = c.ProductFamily
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductFamilyRequest.
// It customizes the JSON unmarshaling process for CreateProductFamilyRequest objects.
func (c *CreateProductFamilyRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ProductFamily CreateProductFamily `json:"product_family"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ProductFamily = temp.ProductFamily
	return nil
}
