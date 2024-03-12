package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateReasonCodeRequest represents a CreateReasonCodeRequest struct.
type CreateReasonCodeRequest struct {
	ReasonCode CreateReasonCode `json:"reason_code"`
}

// MarshalJSON implements the json.Marshaler interface for CreateReasonCodeRequest.
// It customizes the JSON marshaling process for CreateReasonCodeRequest objects.
func (c *CreateReasonCodeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateReasonCodeRequest object to a map representation for JSON marshaling.
func (c *CreateReasonCodeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["reason_code"] = c.ReasonCode.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateReasonCodeRequest.
// It customizes the JSON unmarshaling process for CreateReasonCodeRequest objects.
func (c *CreateReasonCodeRequest) UnmarshalJSON(input []byte) error {
	var temp createReasonCodeRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.ReasonCode = *temp.ReasonCode
	return nil
}

// TODO
type createReasonCodeRequest struct {
	ReasonCode *CreateReasonCode `json:"reason_code"`
}

func (c *createReasonCodeRequest) validate() error {
	var errs []string
	if c.ReasonCode == nil {
		errs = append(errs, "required field `reason_code` is missing for type `Create Reason Code Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
