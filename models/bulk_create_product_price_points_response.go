package models

import (
	"encoding/json"
)

// BulkCreateProductPricePointsResponse represents a BulkCreateProductPricePointsResponse struct.
type BulkCreateProductPricePointsResponse struct {
	PricePoints []ProductPricePoint `json:"price_points,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BulkCreateProductPricePointsResponse.
// It customizes the JSON marshaling process for BulkCreateProductPricePointsResponse objects.
func (b *BulkCreateProductPricePointsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BulkCreateProductPricePointsResponse object to a map representation for JSON marshaling.
func (b *BulkCreateProductPricePointsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if b.PricePoints != nil {
		structMap["price_points"] = b.PricePoints
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkCreateProductPricePointsResponse.
// It customizes the JSON unmarshaling process for BulkCreateProductPricePointsResponse objects.
func (b *BulkCreateProductPricePointsResponse) UnmarshalJSON(input []byte) error {
	var temp bulkCreateProductPricePointsResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	b.PricePoints = temp.PricePoints
	return nil
}

// TODO
type bulkCreateProductPricePointsResponse struct {
	PricePoints []ProductPricePoint `json:"price_points,omitempty"`
}
