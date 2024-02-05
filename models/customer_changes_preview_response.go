package models

import (
    "encoding/json"
)

// CustomerChangesPreviewResponse represents a CustomerChangesPreviewResponse struct.
type CustomerChangesPreviewResponse struct {
    Changes CustomerChange `json:"changes"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerChangesPreviewResponse.
// It customizes the JSON marshaling process for CustomerChangesPreviewResponse objects.
func (c *CustomerChangesPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerChangesPreviewResponse object to a map representation for JSON marshaling.
func (c *CustomerChangesPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["changes"] = c.Changes.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChangesPreviewResponse.
// It customizes the JSON unmarshaling process for CustomerChangesPreviewResponse objects.
func (c *CustomerChangesPreviewResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Changes CustomerChange `json:"changes"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Changes = temp.Changes
    return nil
}
