// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// DunningStepReached represents a DunningStepReached struct.
type DunningStepReached struct {
	Dunner               DunnerData             `json:"dunner"`
	CurrentStep          DunningStepData        `json:"current_step"`
	NextStep             DunningStepData        `json:"next_step"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DunningStepReached,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DunningStepReached) String() string {
	return fmt.Sprintf(
		"DunningStepReached[Dunner=%v, CurrentStep=%v, NextStep=%v, AdditionalProperties=%v]",
		d.Dunner, d.CurrentStep, d.NextStep, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DunningStepReached.
// It customizes the JSON marshaling process for DunningStepReached objects.
func (d DunningStepReached) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"dunner", "current_step", "next_step"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DunningStepReached object to a map representation for JSON marshaling.
func (d DunningStepReached) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	structMap["dunner"] = d.Dunner.toMap()
	structMap["current_step"] = d.CurrentStep.toMap()
	structMap["next_step"] = d.NextStep.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DunningStepReached.
// It customizes the JSON unmarshaling process for DunningStepReached objects.
func (d *DunningStepReached) UnmarshalJSON(input []byte) error {
	var temp tempDunningStepReached
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dunner", "current_step", "next_step")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.Dunner = *temp.Dunner
	d.CurrentStep = *temp.CurrentStep
	d.NextStep = *temp.NextStep
	return nil
}

// tempDunningStepReached is a temporary struct used for validating the fields of DunningStepReached.
type tempDunningStepReached struct {
	Dunner      *DunnerData      `json:"dunner"`
	CurrentStep *DunningStepData `json:"current_step"`
	NextStep    *DunningStepData `json:"next_step"`
}

func (d *tempDunningStepReached) validate() error {
	var errs []string
	if d.Dunner == nil {
		errs = append(errs, "required field `dunner` is missing for type `Dunning Step Reached`")
	}
	if d.CurrentStep == nil {
		errs = append(errs, "required field `current_step` is missing for type `Dunning Step Reached`")
	}
	if d.NextStep == nil {
		errs = append(errs, "required field `next_step` is missing for type `Dunning Step Reached`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
