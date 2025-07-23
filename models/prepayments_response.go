// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// PrepaymentsResponse represents a PrepaymentsResponse struct.
type PrepaymentsResponse struct {
	Prepayments          []Prepayment           `json:"prepayments,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PrepaymentsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrepaymentsResponse) String() string {
	return fmt.Sprintf(
		"PrepaymentsResponse[Prepayments=%v, AdditionalProperties=%v]",
		p.Prepayments, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PrepaymentsResponse.
// It customizes the JSON marshaling process for PrepaymentsResponse objects.
func (p PrepaymentsResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"prepayments"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PrepaymentsResponse object to a map representation for JSON marshaling.
func (p PrepaymentsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	if p.Prepayments != nil {
		structMap["prepayments"] = p.Prepayments
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaymentsResponse.
// It customizes the JSON unmarshaling process for PrepaymentsResponse objects.
func (p *PrepaymentsResponse) UnmarshalJSON(input []byte) error {
	var temp tempPrepaymentsResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prepayments")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.Prepayments = temp.Prepayments
	return nil
}

// tempPrepaymentsResponse is a temporary struct used for validating the fields of PrepaymentsResponse.
type tempPrepaymentsResponse struct {
	Prepayments []Prepayment `json:"prepayments,omitempty"`
}
