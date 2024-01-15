package models

import (
    "encoding/json"
)

// MRRResponse represents a MRRResponse struct.
type MRRResponse struct {
    Mrr MRR `json:"mrr"`
}

// MarshalJSON implements the json.Marshaler interface for MRRResponse.
// It customizes the JSON marshaling process for MRRResponse objects.
func (m *MRRResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MRRResponse object to a map representation for JSON marshaling.
func (m *MRRResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["mrr"] = m.Mrr
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MRRResponse.
// It customizes the JSON unmarshaling process for MRRResponse objects.
func (m *MRRResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Mrr MRR `json:"mrr"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    m.Mrr = temp.Mrr
    return nil
}
