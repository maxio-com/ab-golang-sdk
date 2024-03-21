package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// UpdateComponentRequest represents a UpdateComponentRequest struct.
type UpdateComponentRequest struct {
	Component UpdateComponent `json:"component"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentRequest.
// It customizes the JSON marshaling process for UpdateComponentRequest objects.
func (u *UpdateComponentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentRequest object to a map representation for JSON marshaling.
func (u *UpdateComponentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["component"] = u.Component.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentRequest.
// It customizes the JSON unmarshaling process for UpdateComponentRequest objects.
func (u *UpdateComponentRequest) UnmarshalJSON(input []byte) error {
	var temp updateComponentRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	u.Component = *temp.Component
	return nil
}

// TODO
type updateComponentRequest struct {
	Component *UpdateComponent `json:"component"`
}

func (u *updateComponentRequest) validate() error {
	var errs []string
	if u.Component == nil {
		errs = append(errs, "required field `component` is missing for type `Update Component Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
