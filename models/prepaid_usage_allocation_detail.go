package models

import (
    "encoding/json"
)

// PrepaidUsageAllocationDetail represents a PrepaidUsageAllocationDetail struct.
type PrepaidUsageAllocationDetail struct {
    AllocationId         *int           `json:"allocation_id,omitempty"`
    ChargeId             *int           `json:"charge_id,omitempty"`
    UsageQuantity        *int           `json:"usage_quantity,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidUsageAllocationDetail.
// It customizes the JSON marshaling process for PrepaidUsageAllocationDetail objects.
func (p PrepaidUsageAllocationDetail) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidUsageAllocationDetail object to a map representation for JSON marshaling.
func (p PrepaidUsageAllocationDetail) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.AllocationId != nil {
        structMap["allocation_id"] = p.AllocationId
    }
    if p.ChargeId != nil {
        structMap["charge_id"] = p.ChargeId
    }
    if p.UsageQuantity != nil {
        structMap["usage_quantity"] = p.UsageQuantity
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidUsageAllocationDetail.
// It customizes the JSON unmarshaling process for PrepaidUsageAllocationDetail objects.
func (p *PrepaidUsageAllocationDetail) UnmarshalJSON(input []byte) error {
    var temp prepaidUsageAllocationDetail
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allocation_id", "charge_id", "usage_quantity")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.AllocationId = temp.AllocationId
    p.ChargeId = temp.ChargeId
    p.UsageQuantity = temp.UsageQuantity
    return nil
}

// prepaidUsageAllocationDetail is a temporary struct used for validating the fields of PrepaidUsageAllocationDetail.
type prepaidUsageAllocationDetail  struct {
    AllocationId  *int `json:"allocation_id,omitempty"`
    ChargeId      *int `json:"charge_id,omitempty"`
    UsageQuantity *int `json:"usage_quantity,omitempty"`
}
