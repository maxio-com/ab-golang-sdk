package models

import (
	"encoding/json"
)

// CreateSegment represents a CreateSegment struct.
type CreateSegment struct {
	// A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric.
	SegmentProperty1Value *interface{} `json:"segment_property_1_value,omitempty"`
	// A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric.
	SegmentProperty2Value *interface{} `json:"segment_property_2_value,omitempty"`
	// A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric.
	SegmentProperty3Value *interface{} `json:"segment_property_3_value,omitempty"`
	// A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric.
	SegmentProperty4Value *interface{} `json:"segment_property_4_value,omitempty"`
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme PricingSchemeEnum            `json:"pricing_scheme"`
	Prices        []CreateOrUpdateSegmentPrice `json:"prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSegment.
// It customizes the JSON marshaling process for CreateSegment objects.
func (c *CreateSegment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSegment object to a map representation for JSON marshaling.
func (c *CreateSegment) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.SegmentProperty1Value != nil {
		structMap["segment_property_1_value"] = c.SegmentProperty1Value
	}
	if c.SegmentProperty2Value != nil {
		structMap["segment_property_2_value"] = c.SegmentProperty2Value
	}
	if c.SegmentProperty3Value != nil {
		structMap["segment_property_3_value"] = c.SegmentProperty3Value
	}
	if c.SegmentProperty4Value != nil {
		structMap["segment_property_4_value"] = c.SegmentProperty4Value
	}
	structMap["pricing_scheme"] = c.PricingScheme
	if c.Prices != nil {
		structMap["prices"] = c.Prices
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSegment.
// It customizes the JSON unmarshaling process for CreateSegment objects.
func (c *CreateSegment) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SegmentProperty1Value *interface{}                 `json:"segment_property_1_value,omitempty"`
		SegmentProperty2Value *interface{}                 `json:"segment_property_2_value,omitempty"`
		SegmentProperty3Value *interface{}                 `json:"segment_property_3_value,omitempty"`
		SegmentProperty4Value *interface{}                 `json:"segment_property_4_value,omitempty"`
		PricingScheme         PricingSchemeEnum            `json:"pricing_scheme"`
		Prices                []CreateOrUpdateSegmentPrice `json:"prices,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.SegmentProperty1Value = temp.SegmentProperty1Value
	c.SegmentProperty2Value = temp.SegmentProperty2Value
	c.SegmentProperty3Value = temp.SegmentProperty3Value
	c.SegmentProperty4Value = temp.SegmentProperty4Value
	c.PricingScheme = temp.PricingScheme
	c.Prices = temp.Prices
	return nil
}
