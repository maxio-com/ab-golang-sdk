// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// MeteredUsage represents a MeteredUsage struct.
type MeteredUsage struct {
    PreviousUnitBalance  string                 `json:"previous_unit_balance"`
    NewUnitBalance       int                    `json:"new_unit_balance"`
    UsageQuantity        int                    `json:"usage_quantity"`
    ComponentId          int                    `json:"component_id"`
    ComponentHandle      string                 `json:"component_handle"`
    Memo                 string                 `json:"memo"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MeteredUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MeteredUsage) String() string {
    return fmt.Sprintf(
    	"MeteredUsage[PreviousUnitBalance=%v, NewUnitBalance=%v, UsageQuantity=%v, ComponentId=%v, ComponentHandle=%v, Memo=%v, AdditionalProperties=%v]",
    	m.PreviousUnitBalance, m.NewUnitBalance, m.UsageQuantity, m.ComponentId, m.ComponentHandle, m.Memo, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MeteredUsage.
// It customizes the JSON marshaling process for MeteredUsage objects.
func (m MeteredUsage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "previous_unit_balance", "new_unit_balance", "usage_quantity", "component_id", "component_handle", "memo"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MeteredUsage object to a map representation for JSON marshaling.
func (m MeteredUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["previous_unit_balance"] = m.PreviousUnitBalance
    structMap["new_unit_balance"] = m.NewUnitBalance
    structMap["usage_quantity"] = m.UsageQuantity
    structMap["component_id"] = m.ComponentId
    structMap["component_handle"] = m.ComponentHandle
    structMap["memo"] = m.Memo
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MeteredUsage.
// It customizes the JSON unmarshaling process for MeteredUsage objects.
func (m *MeteredUsage) UnmarshalJSON(input []byte) error {
    var temp tempMeteredUsage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "previous_unit_balance", "new_unit_balance", "usage_quantity", "component_id", "component_handle", "memo")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.PreviousUnitBalance = *temp.PreviousUnitBalance
    m.NewUnitBalance = *temp.NewUnitBalance
    m.UsageQuantity = *temp.UsageQuantity
    m.ComponentId = *temp.ComponentId
    m.ComponentHandle = *temp.ComponentHandle
    m.Memo = *temp.Memo
    return nil
}

// tempMeteredUsage is a temporary struct used for validating the fields of MeteredUsage.
type tempMeteredUsage  struct {
    PreviousUnitBalance *string `json:"previous_unit_balance"`
    NewUnitBalance      *int    `json:"new_unit_balance"`
    UsageQuantity       *int    `json:"usage_quantity"`
    ComponentId         *int    `json:"component_id"`
    ComponentHandle     *string `json:"component_handle"`
    Memo                *string `json:"memo"`
}

func (m *tempMeteredUsage) validate() error {
    var errs []string
    if m.PreviousUnitBalance == nil {
        errs = append(errs, "required field `previous_unit_balance` is missing for type `Metered Usage`")
    }
    if m.NewUnitBalance == nil {
        errs = append(errs, "required field `new_unit_balance` is missing for type `Metered Usage`")
    }
    if m.UsageQuantity == nil {
        errs = append(errs, "required field `usage_quantity` is missing for type `Metered Usage`")
    }
    if m.ComponentId == nil {
        errs = append(errs, "required field `component_id` is missing for type `Metered Usage`")
    }
    if m.ComponentHandle == nil {
        errs = append(errs, "required field `component_handle` is missing for type `Metered Usage`")
    }
    if m.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Metered Usage`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
