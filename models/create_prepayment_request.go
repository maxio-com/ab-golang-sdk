package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreatePrepaymentRequest represents a CreatePrepaymentRequest struct.
type CreatePrepaymentRequest struct {
    Prepayment           CreatePrepayment `json:"prepayment"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaymentRequest.
// It customizes the JSON marshaling process for CreatePrepaymentRequest objects.
func (c CreatePrepaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaymentRequest object to a map representation for JSON marshaling.
func (c CreatePrepaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["prepayment"] = c.Prepayment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaymentRequest.
// It customizes the JSON unmarshaling process for CreatePrepaymentRequest objects.
func (c *CreatePrepaymentRequest) UnmarshalJSON(input []byte) error {
    var temp createPrepaymentRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "prepayment")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Prepayment = *temp.Prepayment
    return nil
}

// createPrepaymentRequest is a temporary struct used for validating the fields of CreatePrepaymentRequest.
type createPrepaymentRequest  struct {
    Prepayment *CreatePrepayment `json:"prepayment"`
}

func (c *createPrepaymentRequest) validate() error {
    var errs []string
    if c.Prepayment == nil {
        errs = append(errs, "required field `prepayment` is missing for type `Create Prepayment Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
