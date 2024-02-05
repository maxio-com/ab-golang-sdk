package models

import (
    "encoding/json"
)

// ReasonCodeResponse represents a ReasonCodeResponse struct.
type ReasonCodeResponse struct {
    ReasonCode ReasonCode `json:"reason_code"`
}

// MarshalJSON implements the json.Marshaler interface for ReasonCodeResponse.
// It customizes the JSON marshaling process for ReasonCodeResponse objects.
func (r *ReasonCodeResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReasonCodeResponse object to a map representation for JSON marshaling.
func (r *ReasonCodeResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["reason_code"] = r.ReasonCode.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReasonCodeResponse.
// It customizes the JSON unmarshaling process for ReasonCodeResponse objects.
func (r *ReasonCodeResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ReasonCode ReasonCode `json:"reason_code"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.ReasonCode = temp.ReasonCode
    return nil
}
