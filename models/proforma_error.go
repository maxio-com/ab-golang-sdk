package models

import (
    "encoding/json"
)

// ProformaError represents a ProformaError struct.
type ProformaError struct {
    // The error is base if it is not directly associated with a single attribute.
    Subscription         *BaseStringError `json:"subscription,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaError.
// It customizes the JSON marshaling process for ProformaError objects.
func (p ProformaError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaError object to a map representation for JSON marshaling.
func (p ProformaError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Subscription != nil {
        structMap["subscription"] = p.Subscription.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaError.
// It customizes the JSON unmarshaling process for ProformaError objects.
func (p *ProformaError) UnmarshalJSON(input []byte) error {
    var temp proformaError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscription")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Subscription = temp.Subscription
    return nil
}

// TODO
type proformaError  struct {
    Subscription *BaseStringError `json:"subscription,omitempty"`
}
