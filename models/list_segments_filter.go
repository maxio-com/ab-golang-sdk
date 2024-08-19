/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ListSegmentsFilter represents a ListSegmentsFilter struct.
type ListSegmentsFilter struct {
    // The value passed here would be used to filter segments. Pass a value related to `segment_property_1` on attached Metric. If empty string is passed, this filter would be rejected. Use in query `filter[segment_property_1_value]=EU`.
    SegmentProperty1Value *string        `json:"segment_property_1_value,omitempty"`
    // The value passed here would be used to filter segments. Pass a value related to `segment_property_2` on attached Metric. If empty string is passed, this filter would be rejected.
    SegmentProperty2Value *string        `json:"segment_property_2_value,omitempty"`
    // The value passed here would be used to filter segments. Pass a value related to `segment_property_3` on attached Metric. If empty string is passed, this filter would be rejected.
    SegmentProperty3Value *string        `json:"segment_property_3_value,omitempty"`
    // The value passed here would be used to filter segments. Pass a value related to `segment_property_4` on attached Metric. If empty string is passed, this filter would be rejected.
    SegmentProperty4Value *string        `json:"segment_property_4_value,omitempty"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListSegmentsFilter.
// It customizes the JSON marshaling process for ListSegmentsFilter objects.
func (l ListSegmentsFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListSegmentsFilter object to a map representation for JSON marshaling.
func (l ListSegmentsFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.SegmentProperty1Value != nil {
        structMap["segment_property_1_value"] = l.SegmentProperty1Value
    }
    if l.SegmentProperty2Value != nil {
        structMap["segment_property_2_value"] = l.SegmentProperty2Value
    }
    if l.SegmentProperty3Value != nil {
        structMap["segment_property_3_value"] = l.SegmentProperty3Value
    }
    if l.SegmentProperty4Value != nil {
        structMap["segment_property_4_value"] = l.SegmentProperty4Value
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSegmentsFilter.
// It customizes the JSON unmarshaling process for ListSegmentsFilter objects.
func (l *ListSegmentsFilter) UnmarshalJSON(input []byte) error {
    var temp tempListSegmentsFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "segment_property_1_value", "segment_property_2_value", "segment_property_3_value", "segment_property_4_value")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.SegmentProperty1Value = temp.SegmentProperty1Value
    l.SegmentProperty2Value = temp.SegmentProperty2Value
    l.SegmentProperty3Value = temp.SegmentProperty3Value
    l.SegmentProperty4Value = temp.SegmentProperty4Value
    return nil
}

// tempListSegmentsFilter is a temporary struct used for validating the fields of ListSegmentsFilter.
type tempListSegmentsFilter  struct {
    SegmentProperty1Value *string `json:"segment_property_1_value,omitempty"`
    SegmentProperty2Value *string `json:"segment_property_2_value,omitempty"`
    SegmentProperty3Value *string `json:"segment_property_3_value,omitempty"`
    SegmentProperty4Value *string `json:"segment_property_4_value,omitempty"`
}
