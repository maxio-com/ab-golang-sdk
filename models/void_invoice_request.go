// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// VoidInvoiceRequest represents a VoidInvoiceRequest struct.
type VoidInvoiceRequest struct {
	Void                 VoidInvoice            `json:"void"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VoidInvoiceRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VoidInvoiceRequest) String() string {
	return fmt.Sprintf(
		"VoidInvoiceRequest[Void=%v, AdditionalProperties=%v]",
		v.Void, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VoidInvoiceRequest.
// It customizes the JSON marshaling process for VoidInvoiceRequest objects.
func (v VoidInvoiceRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(v.AdditionalProperties,
		"void"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(v.toMap())
}

// toMap converts the VoidInvoiceRequest object to a map representation for JSON marshaling.
func (v VoidInvoiceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, v.AdditionalProperties)
	structMap["void"] = v.Void.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidInvoiceRequest.
// It customizes the JSON unmarshaling process for VoidInvoiceRequest objects.
func (v *VoidInvoiceRequest) UnmarshalJSON(input []byte) error {
	var temp tempVoidInvoiceRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "void")
	if err != nil {
		return err
	}
	v.AdditionalProperties = additionalProperties

	v.Void = *temp.Void
	return nil
}

// tempVoidInvoiceRequest is a temporary struct used for validating the fields of VoidInvoiceRequest.
type tempVoidInvoiceRequest struct {
	Void *VoidInvoice `json:"void"`
}

func (v *tempVoidInvoiceRequest) validate() error {
	var errs []string
	if v.Void == nil {
		errs = append(errs, "required field `void` is missing for type `Void Invoice Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
