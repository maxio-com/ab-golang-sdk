package models

import (
    "encoding/json"
)

// ProductResponse represents a ProductResponse struct.
type ProductResponse struct {
    Product Product `json:"product"`
}

// MarshalJSON implements the json.Marshaler interface for ProductResponse.
// It customizes the JSON marshaling process for ProductResponse objects.
func (p *ProductResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductResponse object to a map representation for JSON marshaling.
func (p *ProductResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["product"] = p.Product.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductResponse.
// It customizes the JSON unmarshaling process for ProductResponse objects.
func (p *ProductResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Product Product `json:"product"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Product = temp.Product
    return nil
}
