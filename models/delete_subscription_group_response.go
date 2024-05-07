package models

import (
    "encoding/json"
)

// DeleteSubscriptionGroupResponse represents a DeleteSubscriptionGroupResponse struct.
type DeleteSubscriptionGroupResponse struct {
    Uid                  *string        `json:"uid,omitempty"`
    Deleted              *bool          `json:"deleted,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeleteSubscriptionGroupResponse.
// It customizes the JSON marshaling process for DeleteSubscriptionGroupResponse objects.
func (d DeleteSubscriptionGroupResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeleteSubscriptionGroupResponse object to a map representation for JSON marshaling.
func (d DeleteSubscriptionGroupResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
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
    var temp deleteSubscriptionGroupResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "deleted")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Uid = temp.Uid
    d.Deleted = temp.Deleted
    return nil
}

// deleteSubscriptionGroupResponse is a temporary struct used for validating the fields of DeleteSubscriptionGroupResponse.
type deleteSubscriptionGroupResponse  struct {
    Uid     *string `json:"uid,omitempty"`
    Deleted *bool   `json:"deleted,omitempty"`
}
