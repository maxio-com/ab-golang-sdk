package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PrepaidUsage represents a PrepaidUsage struct.
type PrepaidUsage struct {
    PreviousUnitBalance        string                         `json:"previous_unit_balance"`
    PreviousOverageUnitBalance string                         `json:"previous_overage_unit_balance"`
    NewUnitBalance             int                            `json:"new_unit_balance"`
    NewOverageUnitBalance      int                            `json:"new_overage_unit_balance"`
    UsageQuantity              int                            `json:"usage_quantity"`
    OverageUsageQuantity       int                            `json:"overage_usage_quantity"`
    ComponentId                int                            `json:"component_id"`
    ComponentHandle            string                         `json:"component_handle"`
    Memo                       string                         `json:"memo"`
    AllocationDetails          []PrepaidUsageAllocationDetail `json:"allocation_details"`
    AdditionalProperties       map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidUsage.
// It customizes the JSON marshaling process for PrepaidUsage objects.
func (p PrepaidUsage) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidUsage object to a map representation for JSON marshaling.
func (p PrepaidUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["previous_unit_balance"] = p.PreviousUnitBalance
    structMap["previous_overage_unit_balance"] = p.PreviousOverageUnitBalance
    structMap["new_unit_balance"] = p.NewUnitBalance
    structMap["new_overage_unit_balance"] = p.NewOverageUnitBalance
    structMap["usage_quantity"] = p.UsageQuantity
    structMap["overage_usage_quantity"] = p.OverageUsageQuantity
    structMap["component_id"] = p.ComponentId
    structMap["component_handle"] = p.ComponentHandle
    structMap["memo"] = p.Memo
    structMap["allocation_details"] = p.AllocationDetails
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidUsage.
// It customizes the JSON unmarshaling process for PrepaidUsage objects.
func (p *PrepaidUsage) UnmarshalJSON(input []byte) error {
    var temp prepaidUsage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "previous_unit_balance", "previous_overage_unit_balance", "new_unit_balance", "new_overage_unit_balance", "usage_quantity", "overage_usage_quantity", "component_id", "component_handle", "memo", "allocation_details")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.PreviousUnitBalance = *temp.PreviousUnitBalance
    p.PreviousOverageUnitBalance = *temp.PreviousOverageUnitBalance
    p.NewUnitBalance = *temp.NewUnitBalance
    p.NewOverageUnitBalance = *temp.NewOverageUnitBalance
    p.UsageQuantity = *temp.UsageQuantity
    p.OverageUsageQuantity = *temp.OverageUsageQuantity
    p.ComponentId = *temp.ComponentId
    p.ComponentHandle = *temp.ComponentHandle
    p.Memo = *temp.Memo
    p.AllocationDetails = *temp.AllocationDetails
    return nil
}

// TODO
type prepaidUsage  struct {
    PreviousUnitBalance        *string                         `json:"previous_unit_balance"`
    PreviousOverageUnitBalance *string                         `json:"previous_overage_unit_balance"`
    NewUnitBalance             *int                            `json:"new_unit_balance"`
    NewOverageUnitBalance      *int                            `json:"new_overage_unit_balance"`
    UsageQuantity              *int                            `json:"usage_quantity"`
    OverageUsageQuantity       *int                            `json:"overage_usage_quantity"`
    ComponentId                *int                            `json:"component_id"`
    ComponentHandle            *string                         `json:"component_handle"`
    Memo                       *string                         `json:"memo"`
    AllocationDetails          *[]PrepaidUsageAllocationDetail `json:"allocation_details"`
}

func (p *prepaidUsage) validate() error {
    var errs []string
    if p.PreviousUnitBalance == nil {
        errs = append(errs, "required field `previous_unit_balance` is missing for type `Prepaid Usage`")
    }
    if p.PreviousOverageUnitBalance == nil {
        errs = append(errs, "required field `previous_overage_unit_balance` is missing for type `Prepaid Usage`")
    }
    if p.NewUnitBalance == nil {
        errs = append(errs, "required field `new_unit_balance` is missing for type `Prepaid Usage`")
    }
    if p.NewOverageUnitBalance == nil {
        errs = append(errs, "required field `new_overage_unit_balance` is missing for type `Prepaid Usage`")
    }
    if p.UsageQuantity == nil {
        errs = append(errs, "required field `usage_quantity` is missing for type `Prepaid Usage`")
    }
    if p.OverageUsageQuantity == nil {
        errs = append(errs, "required field `overage_usage_quantity` is missing for type `Prepaid Usage`")
    }
    if p.ComponentId == nil {
        errs = append(errs, "required field `component_id` is missing for type `Prepaid Usage`")
    }
    if p.ComponentHandle == nil {
        errs = append(errs, "required field `component_handle` is missing for type `Prepaid Usage`")
    }
    if p.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Prepaid Usage`")
    }
    if p.AllocationDetails == nil {
        errs = append(errs, "required field `allocation_details` is missing for type `Prepaid Usage`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
