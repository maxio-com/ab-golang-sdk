package models

import (
    "encoding/json"
)

// CustomerCustomFieldsChange represents a CustomerCustomFieldsChange struct.
type CustomerCustomFieldsChange struct {
    Before []ProformaCustomField `json:"before,omitempty"`
    After  []ProformaCustomField `json:"after,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerCustomFieldsChange.
// It customizes the JSON marshaling process for CustomerCustomFieldsChange objects.
func (c *CustomerCustomFieldsChange) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerCustomFieldsChange object to a map representation for JSON marshaling.
func (c *CustomerCustomFieldsChange) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Before != nil {
        structMap["before"] = c.Before
    }
    if c.After != nil {
        structMap["after"] = c.After
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerCustomFieldsChange.
// It customizes the JSON unmarshaling process for CustomerCustomFieldsChange objects.
func (c *CustomerCustomFieldsChange) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Before []ProformaCustomField `json:"before,omitempty"`
        After  []ProformaCustomField `json:"after,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Before = temp.Before
    c.After = temp.After
    return nil
}
