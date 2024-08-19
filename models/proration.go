/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// Proration represents a Proration struct.
type Proration struct {
    // The alternative to sending preserve_period as a direct attribute to migration
    PreservePeriod       *bool          `json:"preserve_period,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Proration.
// It customizes the JSON marshaling process for Proration objects.
func (p Proration) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the Proration object to a map representation for JSON marshaling.
func (p Proration) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.PreservePeriod != nil {
        structMap["preserve_period"] = p.PreservePeriod
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Proration.
// It customizes the JSON unmarshaling process for Proration objects.
func (p *Proration) UnmarshalJSON(input []byte) error {
    var temp tempProration
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "preserve_period")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.PreservePeriod = temp.PreservePeriod
    return nil
}

// tempProration is a temporary struct used for validating the fields of Proration.
type tempProration  struct {
    PreservePeriod *bool `json:"preserve_period,omitempty"`
}
