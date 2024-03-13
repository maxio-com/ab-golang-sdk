package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateSegment represents a CreateSegment struct.
type CreateSegment struct {
	// A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric.
	SegmentProperty1Value *CreateSegmentSegmentProperty1Value `json:"segment_property_1_value,omitempty"`
	// A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric.
	SegmentProperty2Value *CreateSegmentSegmentProperty2Value `json:"segment_property_2_value,omitempty"`
	// A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric.
	SegmentProperty3Value *CreateSegmentSegmentProperty3Value `json:"segment_property_3_value,omitempty"`
	// A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric.
	SegmentProperty4Value *CreateSegmentSegmentProperty4Value `json:"segment_property_4_value,omitempty"`
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme PricingScheme                `json:"pricing_scheme"`
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
		structMap["segment_property_1_value"] = c.SegmentProperty1Value.toMap()
	}
	if c.SegmentProperty2Value != nil {
		structMap["segment_property_2_value"] = c.SegmentProperty2Value.toMap()
	}
	if c.SegmentProperty3Value != nil {
		structMap["segment_property_3_value"] = c.SegmentProperty3Value.toMap()
	}
	if c.SegmentProperty4Value != nil {
		structMap["segment_property_4_value"] = c.SegmentProperty4Value.toMap()
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
	var temp createSegment
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.SegmentProperty1Value = temp.SegmentProperty1Value
	c.SegmentProperty2Value = temp.SegmentProperty2Value
	c.SegmentProperty3Value = temp.SegmentProperty3Value
	c.SegmentProperty4Value = temp.SegmentProperty4Value
	c.PricingScheme = *temp.PricingScheme
	c.Prices = temp.Prices
	return nil
}

// TODO
type createSegment struct {
	SegmentProperty1Value *CreateSegmentSegmentProperty1Value `json:"segment_property_1_value,omitempty"`
	SegmentProperty2Value *CreateSegmentSegmentProperty2Value `json:"segment_property_2_value,omitempty"`
	SegmentProperty3Value *CreateSegmentSegmentProperty3Value `json:"segment_property_3_value,omitempty"`
	SegmentProperty4Value *CreateSegmentSegmentProperty4Value `json:"segment_property_4_value,omitempty"`
	PricingScheme         *PricingScheme                      `json:"pricing_scheme"`
	Prices                []CreateOrUpdateSegmentPrice        `json:"prices,omitempty"`
}

func (c *createSegment) validate() error {
	var errs []string
	if c.PricingScheme == nil {
		errs = append(errs, "required field `pricing_scheme` is missing for type `Create Segment`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
