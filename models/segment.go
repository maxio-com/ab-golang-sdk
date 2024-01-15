package models

import (
    "encoding/json"
)

// Segment represents a Segment struct.
type Segment struct {
    Id                        *int           `json:"id,omitempty"`
    ComponentId               *int           `json:"component_id,omitempty"`
    PricePointId              *int           `json:"price_point_id,omitempty"`
    EventBasedBillingMetricId *int           `json:"event_based_billing_metric_id,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme             *PricingScheme `json:"pricing_scheme,omitempty"`
    SegmentProperty1Value     *interface{}   `json:"segment_property_1_value,omitempty"`
    SegmentProperty2Value     *interface{}   `json:"segment_property_2_value,omitempty"`
    SegmentProperty3Value     *interface{}   `json:"segment_property_3_value,omitempty"`
    SegmentProperty4Value     *interface{}   `json:"segment_property_4_value,omitempty"`
    CreatedAt                 *string        `json:"created_at,omitempty"`
    UpdatedAt                 *string        `json:"updated_at,omitempty"`
    Prices                    []SegmentPrice `json:"prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Segment.
// It customizes the JSON marshaling process for Segment objects.
func (s *Segment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Segment object to a map representation for JSON marshaling.
func (s *Segment) toMap() map[string]any {
    structMap := make(map[string]any)
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
        structMap["segment_property_1_value"] = s.SegmentProperty1Value
    }
    if s.SegmentProperty2Value != nil {
        structMap["segment_property_2_value"] = s.SegmentProperty2Value
    }
    if s.SegmentProperty3Value != nil {
        structMap["segment_property_3_value"] = s.SegmentProperty3Value
    }
    if s.SegmentProperty4Value != nil {
        structMap["segment_property_4_value"] = s.SegmentProperty4Value
    }
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt
    }
    if s.UpdatedAt != nil {
        structMap["updated_at"] = s.UpdatedAt
    }
    if s.Prices != nil {
        structMap["prices"] = s.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Segment.
// It customizes the JSON unmarshaling process for Segment objects.
func (s *Segment) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                        *int           `json:"id,omitempty"`
        ComponentId               *int           `json:"component_id,omitempty"`
        PricePointId              *int           `json:"price_point_id,omitempty"`
        EventBasedBillingMetricId *int           `json:"event_based_billing_metric_id,omitempty"`
        PricingScheme             *PricingScheme `json:"pricing_scheme,omitempty"`
        SegmentProperty1Value     *interface{}   `json:"segment_property_1_value,omitempty"`
        SegmentProperty2Value     *interface{}   `json:"segment_property_2_value,omitempty"`
        SegmentProperty3Value     *interface{}   `json:"segment_property_3_value,omitempty"`
        SegmentProperty4Value     *interface{}   `json:"segment_property_4_value,omitempty"`
        CreatedAt                 *string        `json:"created_at,omitempty"`
        UpdatedAt                 *string        `json:"updated_at,omitempty"`
        Prices                    []SegmentPrice `json:"prices,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Id = temp.Id
    s.ComponentId = temp.ComponentId
    s.PricePointId = temp.PricePointId
    s.EventBasedBillingMetricId = temp.EventBasedBillingMetricId
    s.PricingScheme = temp.PricingScheme
    s.SegmentProperty1Value = temp.SegmentProperty1Value
    s.SegmentProperty2Value = temp.SegmentProperty2Value
    s.SegmentProperty3Value = temp.SegmentProperty3Value
    s.SegmentProperty4Value = temp.SegmentProperty4Value
    s.CreatedAt = temp.CreatedAt
    s.UpdatedAt = temp.UpdatedAt
    s.Prices = temp.Prices
    return nil
}
