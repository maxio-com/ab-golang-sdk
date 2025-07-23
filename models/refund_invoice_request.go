// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// RefundInvoiceRequest represents a RefundInvoiceRequest struct.
type RefundInvoiceRequest struct {
	Refund               RefundInvoiceRequestRefund `json:"refund"`
	AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for RefundInvoiceRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RefundInvoiceRequest) String() string {
	return fmt.Sprintf(
		"RefundInvoiceRequest[Refund=%v, AdditionalProperties=%v]",
		r.Refund, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RefundInvoiceRequest.
// It customizes the JSON marshaling process for RefundInvoiceRequest objects.
func (r RefundInvoiceRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"refund"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RefundInvoiceRequest object to a map representation for JSON marshaling.
func (r RefundInvoiceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["refund"] = r.Refund.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundInvoiceRequest.
// It customizes the JSON unmarshaling process for RefundInvoiceRequest objects.
func (r *RefundInvoiceRequest) UnmarshalJSON(input []byte) error {
	var temp tempRefundInvoiceRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "refund")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Refund = *temp.Refund
	return nil
}

// tempRefundInvoiceRequest is a temporary struct used for validating the fields of RefundInvoiceRequest.
type tempRefundInvoiceRequest struct {
	Refund *RefundInvoiceRequestRefund `json:"refund"`
}

func (r *tempRefundInvoiceRequest) validate() error {
	var errs []string
	if r.Refund == nil {
		errs = append(errs, "required field `refund` is missing for type `Refund Invoice Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
