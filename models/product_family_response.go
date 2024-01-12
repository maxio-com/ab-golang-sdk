package models

import (
    "encoding/json"
)

// ProductFamilyResponse represents a ProductFamilyResponse struct.
type ProductFamilyResponse struct {
    ProductFamily *ProductFamily `json:"product_family,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProductFamilyResponse.
// It customizes the JSON marshaling process for ProductFamilyResponse objects.
func (p *ProductFamilyResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductFamilyResponse object to a map representation for JSON marshaling.
func (p *ProductFamilyResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.ProductFamily != nil {
        structMap["product_family"] = p.ProductFamily
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductFamilyResponse.
// It customizes the JSON unmarshaling process for ProductFamilyResponse objects.
func (p *ProductFamilyResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ProductFamily *ProductFamily `json:"product_family,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.ProductFamily = temp.ProductFamily
    return nil
}
