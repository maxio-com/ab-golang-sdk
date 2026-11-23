// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SegmentPricePointErrorDetails represents a SegmentPricePointErrorDetails struct.
// Price-point validation error details for segment writes.
type SegmentPricePointErrorDetails struct {
    PricePoint           []string               `json:"price_point"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SegmentPricePointErrorDetails,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SegmentPricePointErrorDetails) String() string {
    return fmt.Sprintf(
    	"SegmentPricePointErrorDetails[PricePoint=%v, AdditionalProperties=%v]",
    	s.PricePoint, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SegmentPricePointErrorDetails.
// It customizes the JSON marshaling process for SegmentPricePointErrorDetails objects.
func (s SegmentPricePointErrorDetails) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SegmentPricePointErrorDetails object to a map representation for JSON marshaling.
func (s SegmentPricePointErrorDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["price_point"] = s.PricePoint
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SegmentPricePointErrorDetails.
// It customizes the JSON unmarshaling process for SegmentPricePointErrorDetails objects.
func (s *SegmentPricePointErrorDetails) UnmarshalJSON(input []byte) error {
    var temp tempSegmentPricePointErrorDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_point")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.PricePoint = *temp.PricePoint
    return nil
}

// tempSegmentPricePointErrorDetails is a temporary struct used for validating the fields of SegmentPricePointErrorDetails.
type tempSegmentPricePointErrorDetails  struct {
    PricePoint *[]string `json:"price_point"`
}

func (s *tempSegmentPricePointErrorDetails) validate() error {
    var errs []string
    if s.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Segment Price Point Error Details`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
