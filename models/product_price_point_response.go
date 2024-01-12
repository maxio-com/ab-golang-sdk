package models

import (
    "encoding/json"
)

// ProductPricePointResponse represents a ProductPricePointResponse struct.
type ProductPricePointResponse struct {
    PricePoint ProductPricePoint `json:"price_point"`
}

// MarshalJSON implements the json.Marshaler interface for ProductPricePointResponse.
// It customizes the JSON marshaling process for ProductPricePointResponse objects.
func (p *ProductPricePointResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductPricePointResponse object to a map representation for JSON marshaling.
func (p *ProductPricePointResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["price_point"] = p.PricePoint
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductPricePointResponse.
// It customizes the JSON unmarshaling process for ProductPricePointResponse objects.
func (p *ProductPricePointResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PricePoint ProductPricePoint `json:"price_point"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.PricePoint = temp.PricePoint
    return nil
}
