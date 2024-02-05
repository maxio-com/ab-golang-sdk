package models

import (
    "encoding/json"
)

// AllocationPreviewResponse represents a AllocationPreviewResponse struct.
type AllocationPreviewResponse struct {
    AllocationPreview AllocationPreview `json:"allocation_preview"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewResponse.
// It customizes the JSON marshaling process for AllocationPreviewResponse objects.
func (a *AllocationPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreviewResponse object to a map representation for JSON marshaling.
func (a *AllocationPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["allocation_preview"] = a.AllocationPreview.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewResponse.
// It customizes the JSON unmarshaling process for AllocationPreviewResponse objects.
func (a *AllocationPreviewResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        AllocationPreview AllocationPreview `json:"allocation_preview"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    a.AllocationPreview = temp.AllocationPreview
    return nil
}
