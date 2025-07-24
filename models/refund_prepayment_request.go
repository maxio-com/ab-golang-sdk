// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RefundPrepaymentRequest represents a RefundPrepaymentRequest struct.
type RefundPrepaymentRequest struct {
    Refund               RefundPrepayment       `json:"refund"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RefundPrepaymentRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RefundPrepaymentRequest) String() string {
    return fmt.Sprintf(
    	"RefundPrepaymentRequest[Refund=%v, AdditionalProperties=%v]",
    	r.Refund, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RefundPrepaymentRequest.
// It customizes the JSON marshaling process for RefundPrepaymentRequest objects.
func (r RefundPrepaymentRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "refund"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RefundPrepaymentRequest object to a map representation for JSON marshaling.
func (r RefundPrepaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["refund"] = r.Refund.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundPrepaymentRequest.
// It customizes the JSON unmarshaling process for RefundPrepaymentRequest objects.
func (r *RefundPrepaymentRequest) UnmarshalJSON(input []byte) error {
    var temp tempRefundPrepaymentRequest
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

// tempRefundPrepaymentRequest is a temporary struct used for validating the fields of RefundPrepaymentRequest.
type tempRefundPrepaymentRequest  struct {
    Refund *RefundPrepayment `json:"refund"`
}

func (r *tempRefundPrepaymentRequest) validate() error {
    var errs []string
    if r.Refund == nil {
        errs = append(errs, "required field `refund` is missing for type `Refund Prepayment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
