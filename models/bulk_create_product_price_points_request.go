package models

import (
    "encoding/json"
)

// BulkCreateProductPricePointsRequest represents a BulkCreateProductPricePointsRequest struct.
type BulkCreateProductPricePointsRequest struct {
    PricePoints []CreateProductPricePoint `json:"price_points"`
}

// MarshalJSON implements the json.Marshaler interface for BulkCreateProductPricePointsRequest.
// It customizes the JSON marshaling process for BulkCreateProductPricePointsRequest objects.
func (b *BulkCreateProductPricePointsRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BulkCreateProductPricePointsRequest object to a map representation for JSON marshaling.
func (b *BulkCreateProductPricePointsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["price_points"] = b.PricePoints
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkCreateProductPricePointsRequest.
// It customizes the JSON unmarshaling process for BulkCreateProductPricePointsRequest objects.
func (b *BulkCreateProductPricePointsRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PricePoints []CreateProductPricePoint `json:"price_points"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.PricePoints = temp.PricePoints
    return nil
}
