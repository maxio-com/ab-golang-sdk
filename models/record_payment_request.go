package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// RecordPaymentRequest represents a RecordPaymentRequest struct.
type RecordPaymentRequest struct {
    Payment              CreatePayment  `json:"payment"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RecordPaymentRequest.
// It customizes the JSON marshaling process for RecordPaymentRequest objects.
func (r RecordPaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RecordPaymentRequest object to a map representation for JSON marshaling.
func (r RecordPaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["payment"] = r.Payment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RecordPaymentRequest.
// It customizes the JSON unmarshaling process for RecordPaymentRequest objects.
func (r *RecordPaymentRequest) UnmarshalJSON(input []byte) error {
    var temp recordPaymentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "payment")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Payment = *temp.Payment
    return nil
}

// recordPaymentRequest is a temporary struct used for validating the fields of RecordPaymentRequest.
type recordPaymentRequest  struct {
    Payment *CreatePayment `json:"payment"`
}

func (r *recordPaymentRequest) validate() error {
    var errs []string
    if r.Payment == nil {
        errs = append(errs, "required field `payment` is missing for type `Record Payment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
