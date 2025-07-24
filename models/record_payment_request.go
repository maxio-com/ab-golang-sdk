// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RecordPaymentRequest represents a RecordPaymentRequest struct.
type RecordPaymentRequest struct {
    Payment              CreatePayment          `json:"payment"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RecordPaymentRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RecordPaymentRequest) String() string {
    return fmt.Sprintf(
    	"RecordPaymentRequest[Payment=%v, AdditionalProperties=%v]",
    	r.Payment, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RecordPaymentRequest.
// It customizes the JSON marshaling process for RecordPaymentRequest objects.
func (r RecordPaymentRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "payment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RecordPaymentRequest object to a map representation for JSON marshaling.
func (r RecordPaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["payment"] = r.Payment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RecordPaymentRequest.
// It customizes the JSON unmarshaling process for RecordPaymentRequest objects.
func (r *RecordPaymentRequest) UnmarshalJSON(input []byte) error {
    var temp tempRecordPaymentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payment")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Payment = *temp.Payment
    return nil
}

// tempRecordPaymentRequest is a temporary struct used for validating the fields of RecordPaymentRequest.
type tempRecordPaymentRequest  struct {
    Payment *CreatePayment `json:"payment"`
}

func (r *tempRecordPaymentRequest) validate() error {
    var errs []string
    if r.Payment == nil {
        errs = append(errs, "required field `payment` is missing for type `Record Payment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
