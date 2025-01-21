/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// PrepaidUsageAllocationDetail represents a PrepaidUsageAllocationDetail struct.
type PrepaidUsageAllocationDetail struct {
    AllocationId         *int                   `json:"allocation_id,omitempty"`
    ChargeId             *int                   `json:"charge_id,omitempty"`
    UsageQuantity        *int                   `json:"usage_quantity,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PrepaidUsageAllocationDetail,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrepaidUsageAllocationDetail) String() string {
    return fmt.Sprintf(
    	"PrepaidUsageAllocationDetail[AllocationId=%v, ChargeId=%v, UsageQuantity=%v, AdditionalProperties=%v]",
    	p.AllocationId, p.ChargeId, p.UsageQuantity, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PrepaidUsageAllocationDetail.
// It customizes the JSON marshaling process for PrepaidUsageAllocationDetail objects.
func (p PrepaidUsageAllocationDetail) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "allocation_id", "charge_id", "usage_quantity"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidUsageAllocationDetail object to a map representation for JSON marshaling.
func (p PrepaidUsageAllocationDetail) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
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
    var temp tempPrepaidUsageAllocationDetail
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allocation_id", "charge_id", "usage_quantity")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.AllocationId = temp.AllocationId
    p.ChargeId = temp.ChargeId
    p.UsageQuantity = temp.UsageQuantity
    return nil
}

// tempPrepaidUsageAllocationDetail is a temporary struct used for validating the fields of PrepaidUsageAllocationDetail.
type tempPrepaidUsageAllocationDetail  struct {
    AllocationId  *int `json:"allocation_id,omitempty"`
    ChargeId      *int `json:"charge_id,omitempty"`
    UsageQuantity *int `json:"usage_quantity,omitempty"`
}
