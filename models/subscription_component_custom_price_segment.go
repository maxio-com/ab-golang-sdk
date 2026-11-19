// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionComponentCustomPriceSegment represents a SubscriptionComponentCustomPriceSegment struct.
// Custom rates for one segment of an event-based component, set on the subscription through `custom_price.segments`.
// Identify the segment by the property values that occur in your events, in the same positions as the segmenting properties defined on the component's meter. Each segment in a request must have a unique combination of property values.
// Property values are matched against your events as text. Send values for boolean segmenting properties as the strings `"true"` and `"false"`: JSON booleans are not supported here, and a segment whose only property value is `false` is rejected with a 422.
type SubscriptionComponentCustomPriceSegment struct {
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty1Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty1Value  `json:"segment_property_1_value,omitempty"`
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty2Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty2Value  `json:"segment_property_2_value,omitempty"`
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty3Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty3Value  `json:"segment_property_3_value,omitempty"`
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty4Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty4Value  `json:"segment_property_4_value,omitempty"`
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty5Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty5Value  `json:"segment_property_5_value,omitempty"`
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty6Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty6Value  `json:"segment_property_6_value,omitempty"`
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty7Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty7Value  `json:"segment_property_7_value,omitempty"`
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty8Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty8Value  `json:"segment_property_8_value,omitempty"`
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty9Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty9Value  `json:"segment_property_9_value,omitempty"`
    // A value that will occur in your events that you want to bill upon. For a boolean segmenting property, send the value as a string, `"true"` or `"false"`.
    SegmentProperty10Value *SubscriptionComponentCustomPriceSegmentSegmentProperty10Value `json:"segment_property_10_value,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme          PricingScheme                                                  `json:"pricing_scheme"`
    // The rates for this segment. Use one price for a flat rate, or several price brackets for a tiered, volume or stairstep scheme.
    Prices                 []CreateOrUpdateSegmentPrice                                   `json:"prices"`
    AdditionalProperties   map[string]interface{}                                         `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionComponentCustomPriceSegment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionComponentCustomPriceSegment) String() string {
    return fmt.Sprintf(
    	"SubscriptionComponentCustomPriceSegment[SegmentProperty1Value=%v, SegmentProperty2Value=%v, SegmentProperty3Value=%v, SegmentProperty4Value=%v, SegmentProperty5Value=%v, SegmentProperty6Value=%v, SegmentProperty7Value=%v, SegmentProperty8Value=%v, SegmentProperty9Value=%v, SegmentProperty10Value=%v, PricingScheme=%v, Prices=%v, AdditionalProperties=%v]",
    	s.SegmentProperty1Value, s.SegmentProperty2Value, s.SegmentProperty3Value, s.SegmentProperty4Value, s.SegmentProperty5Value, s.SegmentProperty6Value, s.SegmentProperty7Value, s.SegmentProperty8Value, s.SegmentProperty9Value, s.SegmentProperty10Value, s.PricingScheme, s.Prices, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentCustomPriceSegment.
// It customizes the JSON marshaling process for SubscriptionComponentCustomPriceSegment objects.
func (s SubscriptionComponentCustomPriceSegment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "segment_property_1_value", "segment_property_2_value", "segment_property_3_value", "segment_property_4_value", "segment_property_5_value", "segment_property_6_value", "segment_property_7_value", "segment_property_8_value", "segment_property_9_value", "segment_property_10_value", "pricing_scheme", "prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentCustomPriceSegment object to a map representation for JSON marshaling.
func (s SubscriptionComponentCustomPriceSegment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.SegmentProperty1Value != nil {
        structMap["segment_property_1_value"] = s.SegmentProperty1Value.toMap()
    }
    if s.SegmentProperty2Value != nil {
        structMap["segment_property_2_value"] = s.SegmentProperty2Value.toMap()
    }
    if s.SegmentProperty3Value != nil {
        structMap["segment_property_3_value"] = s.SegmentProperty3Value.toMap()
    }
    if s.SegmentProperty4Value != nil {
        structMap["segment_property_4_value"] = s.SegmentProperty4Value.toMap()
    }
    if s.SegmentProperty5Value != nil {
        structMap["segment_property_5_value"] = s.SegmentProperty5Value.toMap()
    }
    if s.SegmentProperty6Value != nil {
        structMap["segment_property_6_value"] = s.SegmentProperty6Value.toMap()
    }
    if s.SegmentProperty7Value != nil {
        structMap["segment_property_7_value"] = s.SegmentProperty7Value.toMap()
    }
    if s.SegmentProperty8Value != nil {
        structMap["segment_property_8_value"] = s.SegmentProperty8Value.toMap()
    }
    if s.SegmentProperty9Value != nil {
        structMap["segment_property_9_value"] = s.SegmentProperty9Value.toMap()
    }
    if s.SegmentProperty10Value != nil {
        structMap["segment_property_10_value"] = s.SegmentProperty10Value.toMap()
    }
    structMap["pricing_scheme"] = s.PricingScheme
    structMap["prices"] = s.Prices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentCustomPriceSegment.
// It customizes the JSON unmarshaling process for SubscriptionComponentCustomPriceSegment objects.
func (s *SubscriptionComponentCustomPriceSegment) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionComponentCustomPriceSegment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "segment_property_1_value", "segment_property_2_value", "segment_property_3_value", "segment_property_4_value", "segment_property_5_value", "segment_property_6_value", "segment_property_7_value", "segment_property_8_value", "segment_property_9_value", "segment_property_10_value", "pricing_scheme", "prices")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.SegmentProperty1Value = temp.SegmentProperty1Value
    s.SegmentProperty2Value = temp.SegmentProperty2Value
    s.SegmentProperty3Value = temp.SegmentProperty3Value
    s.SegmentProperty4Value = temp.SegmentProperty4Value
    s.SegmentProperty5Value = temp.SegmentProperty5Value
    s.SegmentProperty6Value = temp.SegmentProperty6Value
    s.SegmentProperty7Value = temp.SegmentProperty7Value
    s.SegmentProperty8Value = temp.SegmentProperty8Value
    s.SegmentProperty9Value = temp.SegmentProperty9Value
    s.SegmentProperty10Value = temp.SegmentProperty10Value
    s.PricingScheme = *temp.PricingScheme
    s.Prices = *temp.Prices
    return nil
}

// tempSubscriptionComponentCustomPriceSegment is a temporary struct used for validating the fields of SubscriptionComponentCustomPriceSegment.
type tempSubscriptionComponentCustomPriceSegment  struct {
    SegmentProperty1Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty1Value  `json:"segment_property_1_value,omitempty"`
    SegmentProperty2Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty2Value  `json:"segment_property_2_value,omitempty"`
    SegmentProperty3Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty3Value  `json:"segment_property_3_value,omitempty"`
    SegmentProperty4Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty4Value  `json:"segment_property_4_value,omitempty"`
    SegmentProperty5Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty5Value  `json:"segment_property_5_value,omitempty"`
    SegmentProperty6Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty6Value  `json:"segment_property_6_value,omitempty"`
    SegmentProperty7Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty7Value  `json:"segment_property_7_value,omitempty"`
    SegmentProperty8Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty8Value  `json:"segment_property_8_value,omitempty"`
    SegmentProperty9Value  *SubscriptionComponentCustomPriceSegmentSegmentProperty9Value  `json:"segment_property_9_value,omitempty"`
    SegmentProperty10Value *SubscriptionComponentCustomPriceSegmentSegmentProperty10Value `json:"segment_property_10_value,omitempty"`
    PricingScheme          *PricingScheme                                                 `json:"pricing_scheme"`
    Prices                 *[]CreateOrUpdateSegmentPrice                                  `json:"prices"`
}

func (s *tempSubscriptionComponentCustomPriceSegment) validate() error {
    var errs []string
    if s.PricingScheme == nil {
        errs = append(errs, "required field `pricing_scheme` is missing for type `Subscription Component Custom Price Segment`")
    }
    if s.Prices == nil {
        errs = append(errs, "required field `prices` is missing for type `Subscription Component Custom Price Segment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
