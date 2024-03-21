package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreatePrepaymentRequest represents a CreatePrepaymentRequest struct.
type CreatePrepaymentRequest struct {
	Prepayment CreatePrepayment `json:"prepayment"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaymentRequest.
// It customizes the JSON marshaling process for CreatePrepaymentRequest objects.
func (c *CreatePrepaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaymentRequest object to a map representation for JSON marshaling.
func (c *CreatePrepaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
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

	c.Prepayment = *temp.Prepayment
	return nil
}

// TODO
type createPrepaymentRequest struct {
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
