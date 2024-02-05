package models

import (
    "encoding/json"
)

// GroupSettings represents a GroupSettings struct.
type GroupSettings struct {
    // Attributes of the target customer who will be the responsible payer of the created subscription. Required.
    Target  GroupTarget   `json:"target"`
    // Optional attributes related to billing date and accrual. Note: Only applicable for new subscriptions.
    Billing *GroupBilling `json:"billing,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for GroupSettings.
// It customizes the JSON marshaling process for GroupSettings objects.
func (g *GroupSettings) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GroupSettings object to a map representation for JSON marshaling.
func (g *GroupSettings) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["target"] = g.Target.toMap()
    if g.Billing != nil {
        structMap["billing"] = g.Billing.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GroupSettings.
// It customizes the JSON unmarshaling process for GroupSettings objects.
func (g *GroupSettings) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Target  GroupTarget   `json:"target"`
        Billing *GroupBilling `json:"billing,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Target = temp.Target
    g.Billing = temp.Billing
    return nil
}
