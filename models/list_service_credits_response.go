// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ListServiceCreditsResponse represents a ListServiceCreditsResponse struct.
type ListServiceCreditsResponse struct {
    ServiceCredits       []ServiceCredit1       `json:"service_credits,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListServiceCreditsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListServiceCreditsResponse) String() string {
    return fmt.Sprintf(
    	"ListServiceCreditsResponse[ServiceCredits=%v, AdditionalProperties=%v]",
    	l.ServiceCredits, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListServiceCreditsResponse.
// It customizes the JSON marshaling process for ListServiceCreditsResponse objects.
func (l ListServiceCreditsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "service_credits"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListServiceCreditsResponse object to a map representation for JSON marshaling.
func (l ListServiceCreditsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.ServiceCredits != nil {
        structMap["service_credits"] = l.ServiceCredits
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListServiceCreditsResponse.
// It customizes the JSON unmarshaling process for ListServiceCreditsResponse objects.
func (l *ListServiceCreditsResponse) UnmarshalJSON(input []byte) error {
    var temp tempListServiceCreditsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "service_credits")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.ServiceCredits = temp.ServiceCredits
    return nil
}

// tempListServiceCreditsResponse is a temporary struct used for validating the fields of ListServiceCreditsResponse.
type tempListServiceCreditsResponse  struct {
    ServiceCredits []ServiceCredit1 `json:"service_credits,omitempty"`
}
