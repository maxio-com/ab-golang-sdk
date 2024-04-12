package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreatePrepaymentResponse represents a CreatePrepaymentResponse struct.
type CreatePrepaymentResponse struct {
    Prepayment           CreatedPrepayment `json:"prepayment"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaymentResponse.
// It customizes the JSON marshaling process for CreatePrepaymentResponse objects.
func (c CreatePrepaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaymentResponse object to a map representation for JSON marshaling.
func (c CreatePrepaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["prepayment"] = c.Prepayment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaymentResponse.
// It customizes the JSON unmarshaling process for CreatePrepaymentResponse objects.
func (c *CreatePrepaymentResponse) UnmarshalJSON(input []byte) error {
    var temp createPrepaymentResponse
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

// TODO
type createPrepaymentResponse  struct {
    Prepayment *CreatedPrepayment `json:"prepayment"`
}

func (c *createPrepaymentResponse) validate() error {
    var errs []string
    if c.Prepayment == nil {
        errs = append(errs, "required field `prepayment` is missing for type `Create Prepayment Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
