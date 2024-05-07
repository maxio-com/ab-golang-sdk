package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// DunningStepReached represents a DunningStepReached struct.
type DunningStepReached struct {
    Dunner               DunnerData      `json:"dunner"`
    CurrentStep          DunningStepData `json:"current_step"`
    NextStep             DunningStepData `json:"next_step"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DunningStepReached.
// It customizes the JSON marshaling process for DunningStepReached objects.
func (d DunningStepReached) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DunningStepReached object to a map representation for JSON marshaling.
func (d DunningStepReached) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["dunner"] = d.Dunner.toMap()
    structMap["current_step"] = d.CurrentStep.toMap()
    structMap["next_step"] = d.NextStep.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DunningStepReached.
// It customizes the JSON unmarshaling process for DunningStepReached objects.
func (d *DunningStepReached) UnmarshalJSON(input []byte) error {
    var temp dunningStepReached
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dunner", "current_step", "next_step")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Dunner = *temp.Dunner
    d.CurrentStep = *temp.CurrentStep
    d.NextStep = *temp.NextStep
    return nil
}

// dunningStepReached is a temporary struct used for validating the fields of DunningStepReached.
type dunningStepReached  struct {
    Dunner      *DunnerData      `json:"dunner"`
    CurrentStep *DunningStepData `json:"current_step"`
    NextStep    *DunningStepData `json:"next_step"`
}

func (d *dunningStepReached) validate() error {
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
