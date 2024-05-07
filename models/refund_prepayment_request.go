package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// RefundPrepaymentRequest represents a RefundPrepaymentRequest struct.
type RefundPrepaymentRequest struct {
    Refund               RefundPrepayment `json:"refund"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RefundPrepaymentRequest.
// It customizes the JSON marshaling process for RefundPrepaymentRequest objects.
func (r RefundPrepaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RefundPrepaymentRequest object to a map representation for JSON marshaling.
func (r RefundPrepaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["refund"] = r.Refund.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundPrepaymentRequest.
// It customizes the JSON unmarshaling process for RefundPrepaymentRequest objects.
func (r *RefundPrepaymentRequest) UnmarshalJSON(input []byte) error {
    var temp refundPrepaymentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "refund")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Refund = *temp.Refund
    return nil
}

// refundPrepaymentRequest is a temporary struct used for validating the fields of RefundPrepaymentRequest.
type refundPrepaymentRequest  struct {
    Refund *RefundPrepayment `json:"refund"`
}

func (r *refundPrepaymentRequest) validate() error {
    var errs []string
    if r.Refund == nil {
        errs = append(errs, "required field `refund` is missing for type `Refund Prepayment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
