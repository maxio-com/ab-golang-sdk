package models

import (
    "encoding/json"
)

// DeleteSubscriptionGroupResponse represents a DeleteSubscriptionGroupResponse struct.
type DeleteSubscriptionGroupResponse struct {
    Uid     *string `json:"uid,omitempty"`
    Deleted *bool   `json:"deleted,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for DeleteSubscriptionGroupResponse.
// It customizes the JSON marshaling process for DeleteSubscriptionGroupResponse objects.
func (d *DeleteSubscriptionGroupResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeleteSubscriptionGroupResponse object to a map representation for JSON marshaling.
func (d *DeleteSubscriptionGroupResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if d.Uid != nil {
        structMap["uid"] = d.Uid
    }
    if d.Deleted != nil {
        structMap["deleted"] = d.Deleted
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeleteSubscriptionGroupResponse.
// It customizes the JSON unmarshaling process for DeleteSubscriptionGroupResponse objects.
func (d *DeleteSubscriptionGroupResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Uid     *string `json:"uid,omitempty"`
        Deleted *bool   `json:"deleted,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    d.Uid = temp.Uid
    d.Deleted = temp.Deleted
    return nil
}
