/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "log"
    "time"
)

// Segment represents a Segment struct.
type Segment struct {
    Id                        *int                          `json:"id,omitempty"`
    ComponentId               *int                          `json:"component_id,omitempty"`
    PricePointId              *int                          `json:"price_point_id,omitempty"`
    EventBasedBillingMetricId *int                          `json:"event_based_billing_metric_id,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme             *PricingScheme                `json:"pricing_scheme,omitempty"`
    SegmentProperty1Value     *SegmentSegmentProperty1Value `json:"segment_property_1_value,omitempty"`
    SegmentProperty2Value     *SegmentSegmentProperty2Value `json:"segment_property_2_value,omitempty"`
    SegmentProperty3Value     *SegmentSegmentProperty3Value `json:"segment_property_3_value,omitempty"`
    SegmentProperty4Value     *SegmentSegmentProperty4Value `json:"segment_property_4_value,omitempty"`
    CreatedAt                 *time.Time                    `json:"created_at,omitempty"`
    UpdatedAt                 *time.Time                    `json:"updated_at,omitempty"`
    Prices                    []SegmentPrice                `json:"prices,omitempty"`
    AdditionalProperties      map[string]interface{}        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Segment.
// It customizes the JSON marshaling process for Segment objects.
func (s Segment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "component_id", "price_point_id", "event_based_billing_metric_id", "pricing_scheme", "segment_property_1_value", "segment_property_2_value", "segment_property_3_value", "segment_property_4_value", "created_at", "updated_at", "prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Segment object to a map representation for JSON marshaling.
func (s Segment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.ComponentId != nil {
        structMap["component_id"] = s.ComponentId
    }
    if s.PricePointId != nil {
        structMap["price_point_id"] = s.PricePointId
    }
    if s.EventBasedBillingMetricId != nil {
        structMap["event_based_billing_metric_id"] = s.EventBasedBillingMetricId
    }
    if s.PricingScheme != nil {
        structMap["pricing_scheme"] = s.PricingScheme
    }
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
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt.Format(time.RFC3339)
    }
    if s.UpdatedAt != nil {
        structMap["updated_at"] = s.UpdatedAt.Format(time.RFC3339)
    }
    if s.Prices != nil {
        structMap["prices"] = s.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Segment.
// It customizes the JSON unmarshaling process for Segment objects.
func (s *Segment) UnmarshalJSON(input []byte) error {
    var temp tempSegment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "component_id", "price_point_id", "event_based_billing_metric_id", "pricing_scheme", "segment_property_1_value", "segment_property_2_value", "segment_property_3_value", "segment_property_4_value", "created_at", "updated_at", "prices")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = temp.Id
    s.ComponentId = temp.ComponentId
    s.PricePointId = temp.PricePointId
    s.EventBasedBillingMetricId = temp.EventBasedBillingMetricId
    s.PricingScheme = temp.PricingScheme
    s.SegmentProperty1Value = temp.SegmentProperty1Value
    s.SegmentProperty2Value = temp.SegmentProperty2Value
    s.SegmentProperty3Value = temp.SegmentProperty3Value
    s.SegmentProperty4Value = temp.SegmentProperty4Value
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        s.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        s.UpdatedAt = &UpdatedAtVal
    }
    s.Prices = temp.Prices
    return nil
}

// tempSegment is a temporary struct used for validating the fields of Segment.
type tempSegment  struct {
    Id                        *int                          `json:"id,omitempty"`
    ComponentId               *int                          `json:"component_id,omitempty"`
    PricePointId              *int                          `json:"price_point_id,omitempty"`
    EventBasedBillingMetricId *int                          `json:"event_based_billing_metric_id,omitempty"`
    PricingScheme             *PricingScheme                `json:"pricing_scheme,omitempty"`
    SegmentProperty1Value     *SegmentSegmentProperty1Value `json:"segment_property_1_value,omitempty"`
    SegmentProperty2Value     *SegmentSegmentProperty2Value `json:"segment_property_2_value,omitempty"`
    SegmentProperty3Value     *SegmentSegmentProperty3Value `json:"segment_property_3_value,omitempty"`
    SegmentProperty4Value     *SegmentSegmentProperty4Value `json:"segment_property_4_value,omitempty"`
    CreatedAt                 *string                       `json:"created_at,omitempty"`
    UpdatedAt                 *string                       `json:"updated_at,omitempty"`
    Prices                    []SegmentPrice                `json:"prices,omitempty"`
}
