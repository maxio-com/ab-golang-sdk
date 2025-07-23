// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// RefundPrepaymentBaseRefundError represents a RefundPrepaymentBaseRefundError struct.
type RefundPrepaymentBaseRefundError struct {
	Refund               *BaseRefundError       `json:"refund,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RefundPrepaymentBaseRefundError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RefundPrepaymentBaseRefundError) String() string {
	return fmt.Sprintf(
		"RefundPrepaymentBaseRefundError[Refund=%v, AdditionalProperties=%v]",
		r.Refund, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RefundPrepaymentBaseRefundError.
// It customizes the JSON marshaling process for RefundPrepaymentBaseRefundError objects.
func (r RefundPrepaymentBaseRefundError) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"refund"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RefundPrepaymentBaseRefundError object to a map representation for JSON marshaling.
func (r RefundPrepaymentBaseRefundError) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Refund != nil {
		structMap["refund"] = r.Refund.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundPrepaymentBaseRefundError.
// It customizes the JSON unmarshaling process for RefundPrepaymentBaseRefundError objects.
func (r *RefundPrepaymentBaseRefundError) UnmarshalJSON(input []byte) error {
	var temp tempRefundPrepaymentBaseRefundError
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "refund")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Refund = temp.Refund
	return nil
}

// tempRefundPrepaymentBaseRefundError is a temporary struct used for validating the fields of RefundPrepaymentBaseRefundError.
type tempRefundPrepaymentBaseRefundError struct {
	Refund *BaseRefundError `json:"refund,omitempty"`
}
