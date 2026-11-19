// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Meter represents a Meter struct.
// The metering meter this wallet is paired with, or `null` for a wallet with no meter. Set on create and immutable afterward.
type Meter struct {
    Id                   *int                   `json:"id,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Handle               *string                `json:"handle,omitempty"`
    // How many wallet units one meter unit is worth, as a decimal string. Defaults to `"1.0"`.
    ConversionRate       *string                `json:"conversion_rate,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Meter,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Meter) String() string {
    return fmt.Sprintf(
    	"Meter[Id=%v, Name=%v, Handle=%v, ConversionRate=%v, AdditionalProperties=%v]",
    	m.Id, m.Name, m.Handle, m.ConversionRate, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Meter.
// It customizes the JSON marshaling process for Meter objects.
func (m Meter) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "id", "name", "handle", "conversion_rate"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Meter object to a map representation for JSON marshaling.
func (m Meter) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Id != nil {
        structMap["id"] = m.Id
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.Handle != nil {
        structMap["handle"] = m.Handle
    }
    if m.ConversionRate != nil {
        structMap["conversion_rate"] = m.ConversionRate
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Meter.
// It customizes the JSON unmarshaling process for Meter objects.
func (m *Meter) UnmarshalJSON(input []byte) error {
    var temp tempMeter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "name", "handle", "conversion_rate")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Id = temp.Id
    m.Name = temp.Name
    m.Handle = temp.Handle
    m.ConversionRate = temp.ConversionRate
    return nil
}

// tempMeter is a temporary struct used for validating the fields of Meter.
type tempMeter  struct {
    Id             *int    `json:"id,omitempty"`
    Name           *string `json:"name,omitempty"`
    Handle         *string `json:"handle,omitempty"`
    ConversionRate *string `json:"conversion_rate,omitempty"`
}
