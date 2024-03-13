package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ComponentResponse represents a ComponentResponse struct.
type ComponentResponse struct {
	Component Component `json:"component"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentResponse.
// It customizes the JSON marshaling process for ComponentResponse objects.
func (c *ComponentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentResponse object to a map representation for JSON marshaling.
func (c *ComponentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["component"] = c.Component.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentResponse.
// It customizes the JSON unmarshaling process for ComponentResponse objects.
func (c *ComponentResponse) UnmarshalJSON(input []byte) error {
	var temp componentResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.Component = *temp.Component
	return nil
}

// TODO
type componentResponse struct {
	Component *Component `json:"component"`
}

func (c *componentResponse) validate() error {
	var errs []string
	if c.Component == nil {
		errs = append(errs, "required field `component` is missing for type `Component Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
