package models

import (
    "encoding/json"
)

// CreateOrUpdateProductRequest represents a CreateOrUpdateProductRequest struct.
type CreateOrUpdateProductRequest struct {
    Product CreateOrUpdateProduct `json:"product"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateProductRequest.
// It customizes the JSON marshaling process for CreateOrUpdateProductRequest objects.
func (c *CreateOrUpdateProductRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateProductRequest object to a map representation for JSON marshaling.
func (c *CreateOrUpdateProductRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["product"] = c.Product
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateProductRequest.
// It customizes the JSON unmarshaling process for CreateOrUpdateProductRequest objects.
func (c *CreateOrUpdateProductRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Product CreateOrUpdateProduct `json:"product"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Product = temp.Product
    return nil
}
