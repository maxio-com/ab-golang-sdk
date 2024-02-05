package models

import (
    "encoding/json"
)

// MRRMovement represents a MRRMovement struct.
type MRRMovement struct {
    Amount          *int    `json:"amount,omitempty"`
    Category        *string `json:"category,omitempty"`
    SubscriberDelta *int    `json:"subscriber_delta,omitempty"`
    LeadDelta       *int    `json:"lead_delta,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for MRRMovement.
// It customizes the JSON marshaling process for MRRMovement objects.
func (m *MRRMovement) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MRRMovement object to a map representation for JSON marshaling.
func (m *MRRMovement) toMap() map[string]any {
    structMap := make(map[string]any)
    if m.Amount != nil {
        structMap["amount"] = m.Amount
    }
    if m.Category != nil {
        structMap["category"] = m.Category
    }
    if m.SubscriberDelta != nil {
        structMap["subscriber_delta"] = m.SubscriberDelta
    }
    if m.LeadDelta != nil {
        structMap["lead_delta"] = m.LeadDelta
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MRRMovement.
// It customizes the JSON unmarshaling process for MRRMovement objects.
func (m *MRRMovement) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Amount          *int    `json:"amount,omitempty"`
        Category        *string `json:"category,omitempty"`
        SubscriberDelta *int    `json:"subscriber_delta,omitempty"`
        LeadDelta       *int    `json:"lead_delta,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    m.Amount = temp.Amount
    m.Category = temp.Category
    m.SubscriberDelta = temp.SubscriberDelta
    m.LeadDelta = temp.LeadDelta
    return nil
}
