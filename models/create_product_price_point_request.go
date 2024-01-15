package models

import (
    "encoding/json"
)

// CreateProductPricePointRequest represents a CreateProductPricePointRequest struct.
type CreateProductPricePointRequest struct {
    PricePoint CreateProductPricePoint `json:"price_point"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductPricePointRequest.
// It customizes the JSON marshaling process for CreateProductPricePointRequest objects.
func (c *CreateProductPricePointRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateProductPricePointRequest object to a map representation for JSON marshaling.
func (c *CreateProductPricePointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["price_point"] = c.PricePoint
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductPricePointRequest.
// It customizes the JSON unmarshaling process for CreateProductPricePointRequest objects.
func (c *CreateProductPricePointRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PricePoint CreateProductPricePoint `json:"price_point"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.PricePoint = temp.PricePoint
    return nil
}
