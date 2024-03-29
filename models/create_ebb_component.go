package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateEBBComponent represents a CreateEBBComponent struct.
type CreateEBBComponent struct {
	EventBasedComponent EBBComponent `json:"event_based_component"`
}

// MarshalJSON implements the json.Marshaler interface for CreateEBBComponent.
// It customizes the JSON marshaling process for CreateEBBComponent objects.
func (c *CreateEBBComponent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateEBBComponent object to a map representation for JSON marshaling.
func (c *CreateEBBComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["event_based_component"] = c.EventBasedComponent.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateEBBComponent.
// It customizes the JSON unmarshaling process for CreateEBBComponent objects.
func (c *CreateEBBComponent) UnmarshalJSON(input []byte) error {
	var temp createEBBComponent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.EventBasedComponent = *temp.EventBasedComponent
	return nil
}

// TODO
type createEBBComponent struct {
	EventBasedComponent *EBBComponent `json:"event_based_component"`
}

func (c *createEBBComponent) validate() error {
	var errs []string
	if c.EventBasedComponent == nil {
		errs = append(errs, "required field `event_based_component` is missing for type `Create EBB Component`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
