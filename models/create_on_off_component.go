package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateOnOffComponent represents a CreateOnOffComponent struct.
type CreateOnOffComponent struct {
	OnOffComponent OnOffComponent `json:"on_off_component"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOnOffComponent.
// It customizes the JSON marshaling process for CreateOnOffComponent objects.
func (c *CreateOnOffComponent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateOnOffComponent object to a map representation for JSON marshaling.
func (c *CreateOnOffComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["on_off_component"] = c.OnOffComponent.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOnOffComponent.
// It customizes the JSON unmarshaling process for CreateOnOffComponent objects.
func (c *CreateOnOffComponent) UnmarshalJSON(input []byte) error {
	var temp createOnOffComponent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.OnOffComponent = *temp.OnOffComponent
	return nil
}

// TODO
type createOnOffComponent struct {
	OnOffComponent *OnOffComponent `json:"on_off_component"`
}

func (c *createOnOffComponent) validate() error {
	var errs []string
	if c.OnOffComponent == nil {
		errs = append(errs, "required field `on_off_component` is missing for type `Create On/Off Component`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
