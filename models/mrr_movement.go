/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// MRRMovement represents a MRRMovement struct.
type MRRMovement struct {
    Amount               *int           `json:"amount,omitempty"`
    Category             *string        `json:"category,omitempty"`
    SubscriberDelta      *int           `json:"subscriber_delta,omitempty"`
    LeadDelta            *int           `json:"lead_delta,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MRRMovement.
// It customizes the JSON marshaling process for MRRMovement objects.
func (m MRRMovement) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MRRMovement object to a map representation for JSON marshaling.
func (m MRRMovement) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp tempMRRMovement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "amount", "category", "subscriber_delta", "lead_delta")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Amount = temp.Amount
    m.Category = temp.Category
    m.SubscriberDelta = temp.SubscriberDelta
    m.LeadDelta = temp.LeadDelta
    return nil
}

// tempMRRMovement is a temporary struct used for validating the fields of MRRMovement.
type tempMRRMovement  struct {
    Amount          *int    `json:"amount,omitempty"`
    Category        *string `json:"category,omitempty"`
    SubscriberDelta *int    `json:"subscriber_delta,omitempty"`
    LeadDelta       *int    `json:"lead_delta,omitempty"`
}
