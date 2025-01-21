/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CountResponse represents a CountResponse struct.
type CountResponse struct {
    Count                *int                   `json:"count,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CountResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CountResponse) String() string {
    return fmt.Sprintf(
    	"CountResponse[Count=%v, AdditionalProperties=%v]",
    	c.Count, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CountResponse.
// It customizes the JSON marshaling process for CountResponse objects.
func (c CountResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CountResponse object to a map representation for JSON marshaling.
func (c CountResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Count != nil {
        structMap["count"] = c.Count
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CountResponse.
// It customizes the JSON unmarshaling process for CountResponse objects.
func (c *CountResponse) UnmarshalJSON(input []byte) error {
    var temp tempCountResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "count")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Count = temp.Count
    return nil
}

// tempCountResponse is a temporary struct used for validating the fields of CountResponse.
type tempCountResponse  struct {
    Count *int `json:"count,omitempty"`
}
