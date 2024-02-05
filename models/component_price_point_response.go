package models

import (
    "encoding/json"
)

// ComponentPricePointResponse represents a ComponentPricePointResponse struct.
type ComponentPricePointResponse struct {
    PricePoint ComponentPricePoint `json:"price_point"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointResponse.
// It customizes the JSON marshaling process for ComponentPricePointResponse objects.
func (c *ComponentPricePointResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointResponse object to a map representation for JSON marshaling.
func (c *ComponentPricePointResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["price_point"] = c.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointResponse.
// It customizes the JSON unmarshaling process for ComponentPricePointResponse objects.
func (c *ComponentPricePointResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PricePoint ComponentPricePoint `json:"price_point"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.PricePoint = temp.PricePoint
    return nil
}
