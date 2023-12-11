package models

import (
	"encoding/json"
)

// PreviewAllocationsRequest represents a PreviewAllocationsRequest struct.
type PreviewAllocationsRequest struct {
	Allocations []CreateAllocation `json:"allocations"`
	// To calculate proration amounts for a future time. Only within a current subscription period. Only ISO8601 format is supported.
	EffectiveProrationDate *string `json:"effective_proration_date,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PreviewAllocationsRequest.
// It customizes the JSON marshaling process for PreviewAllocationsRequest objects.
func (p *PreviewAllocationsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PreviewAllocationsRequest object to a map representation for JSON marshaling.
func (p *PreviewAllocationsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["allocations"] = p.Allocations
	if p.EffectiveProrationDate != nil {
		structMap["effective_proration_date"] = p.EffectiveProrationDate
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PreviewAllocationsRequest.
// It customizes the JSON unmarshaling process for PreviewAllocationsRequest objects.
func (p *PreviewAllocationsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Allocations            []CreateAllocation `json:"allocations"`
		EffectiveProrationDate *string            `json:"effective_proration_date,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Allocations = temp.Allocations
	p.EffectiveProrationDate = temp.EffectiveProrationDate
	return nil
}
