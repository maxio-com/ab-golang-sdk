package models

import (
	"encoding/json"
)

// BulkUpdateSegmentsItem represents a BulkUpdateSegmentsItem struct.
type BulkUpdateSegmentsItem struct {
	// The ID of the segment you want to update.
	Id int `json:"id"`
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme PricingSchemeEnum            `json:"pricing_scheme"`
	Prices        []CreateOrUpdateSegmentPrice `json:"prices"`
}

// MarshalJSON implements the json.Marshaler interface for BulkUpdateSegmentsItem.
// It customizes the JSON marshaling process for BulkUpdateSegmentsItem objects.
func (b *BulkUpdateSegmentsItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BulkUpdateSegmentsItem object to a map representation for JSON marshaling.
func (b *BulkUpdateSegmentsItem) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["id"] = b.Id
	structMap["pricing_scheme"] = b.PricingScheme
	structMap["prices"] = b.Prices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkUpdateSegmentsItem.
// It customizes the JSON unmarshaling process for BulkUpdateSegmentsItem objects.
func (b *BulkUpdateSegmentsItem) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id            int                          `json:"id"`
		PricingScheme PricingSchemeEnum            `json:"pricing_scheme"`
		Prices        []CreateOrUpdateSegmentPrice `json:"prices"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Id = temp.Id
	b.PricingScheme = temp.PricingScheme
	b.Prices = temp.Prices
	return nil
}
