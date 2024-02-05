package models

import (
    "encoding/json"
)

// UpdateReasonCodeRequest represents a UpdateReasonCodeRequest struct.
type UpdateReasonCodeRequest struct {
    ReasonCode UpdateReasonCode `json:"reason_code"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateReasonCodeRequest.
// It customizes the JSON marshaling process for UpdateReasonCodeRequest objects.
func (u *UpdateReasonCodeRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateReasonCodeRequest object to a map representation for JSON marshaling.
func (u *UpdateReasonCodeRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["reason_code"] = u.ReasonCode.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateReasonCodeRequest.
// It customizes the JSON unmarshaling process for UpdateReasonCodeRequest objects.
func (u *UpdateReasonCodeRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ReasonCode UpdateReasonCode `json:"reason_code"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.ReasonCode = temp.ReasonCode
    return nil
}
