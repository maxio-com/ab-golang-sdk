package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateAllocationRequest represents a CreateAllocationRequest struct.
type CreateAllocationRequest struct {
	Allocation CreateAllocation `json:"allocation"`
}

// MarshalJSON implements the json.Marshaler interface for CreateAllocationRequest.
// It customizes the JSON marshaling process for CreateAllocationRequest objects.
func (c *CreateAllocationRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateAllocationRequest object to a map representation for JSON marshaling.
func (c *CreateAllocationRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["allocation"] = c.Allocation.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAllocationRequest.
// It customizes the JSON unmarshaling process for CreateAllocationRequest objects.
func (c *CreateAllocationRequest) UnmarshalJSON(input []byte) error {
	var temp createAllocationRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.Allocation = *temp.Allocation
	return nil
}

// TODO
type createAllocationRequest struct {
	Allocation *CreateAllocation `json:"allocation"`
}

func (c *createAllocationRequest) validate() error {
	var errs []string
	if c.Allocation == nil {
		errs = append(errs, "required field `allocation` is missing for type `Create Allocation Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
