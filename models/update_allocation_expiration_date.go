package models

import (
    "encoding/json"
)

// UpdateAllocationExpirationDate represents a UpdateAllocationExpirationDate struct.
type UpdateAllocationExpirationDate struct {
    Allocation           *AllocationExpirationDate `json:"allocation,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateAllocationExpirationDate.
// It customizes the JSON marshaling process for UpdateAllocationExpirationDate objects.
func (u UpdateAllocationExpirationDate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateAllocationExpirationDate object to a map representation for JSON marshaling.
func (u UpdateAllocationExpirationDate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Allocation != nil {
        structMap["allocation"] = u.Allocation.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateAllocationExpirationDate.
// It customizes the JSON unmarshaling process for UpdateAllocationExpirationDate objects.
func (u *UpdateAllocationExpirationDate) UnmarshalJSON(input []byte) error {
    var temp updateAllocationExpirationDate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allocation")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Allocation = temp.Allocation
    return nil
}

// updateAllocationExpirationDate is a temporary struct used for validating the fields of UpdateAllocationExpirationDate.
type updateAllocationExpirationDate  struct {
    Allocation *AllocationExpirationDate `json:"allocation,omitempty"`
}
