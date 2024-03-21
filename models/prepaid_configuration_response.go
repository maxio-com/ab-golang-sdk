package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// PrepaidConfigurationResponse represents a PrepaidConfigurationResponse struct.
type PrepaidConfigurationResponse struct {
	PrepaidConfiguration PrepaidConfiguration `json:"prepaid_configuration"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidConfigurationResponse.
// It customizes the JSON marshaling process for PrepaidConfigurationResponse objects.
func (p *PrepaidConfigurationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PrepaidConfigurationResponse object to a map representation for JSON marshaling.
func (p *PrepaidConfigurationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["prepaid_configuration"] = p.PrepaidConfiguration.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidConfigurationResponse.
// It customizes the JSON unmarshaling process for PrepaidConfigurationResponse objects.
func (p *PrepaidConfigurationResponse) UnmarshalJSON(input []byte) error {
	var temp prepaidConfigurationResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	p.PrepaidConfiguration = *temp.PrepaidConfiguration
	return nil
}

// TODO
type prepaidConfigurationResponse struct {
	PrepaidConfiguration *PrepaidConfiguration `json:"prepaid_configuration"`
}

func (p *prepaidConfigurationResponse) validate() error {
	var errs []string
	if p.PrepaidConfiguration == nil {
		errs = append(errs, "required field `prepaid_configuration` is missing for type `Prepaid Configuration Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
