package models

import (
    "encoding/json"
)

// PrepaymentAggregatedError represents a PrepaymentAggregatedError struct.
type PrepaymentAggregatedError struct {
    AmountInCents []string `json:"amount_in_cents,omitempty"`
    Base          []string `json:"base,omitempty"`
    External      []string `json:"external,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaymentAggregatedError.
// It customizes the JSON marshaling process for PrepaymentAggregatedError objects.
func (p *PrepaymentAggregatedError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaymentAggregatedError object to a map representation for JSON marshaling.
func (p *PrepaymentAggregatedError) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.AmountInCents != nil {
        structMap["amount_in_cents"] = p.AmountInCents
    }
    if p.Base != nil {
        structMap["base"] = p.Base
    }
    if p.External != nil {
        structMap["external"] = p.External
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaymentAggregatedError.
// It customizes the JSON unmarshaling process for PrepaymentAggregatedError objects.
func (p *PrepaymentAggregatedError) UnmarshalJSON(input []byte) error {
    temp := &struct {
        AmountInCents []string `json:"amount_in_cents,omitempty"`
        Base          []string `json:"base,omitempty"`
        External      []string `json:"external,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.AmountInCents = temp.AmountInCents
    p.Base = temp.Base
    p.External = temp.External
    return nil
}
