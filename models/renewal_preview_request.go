package models

import (
	"encoding/json"
)

// RenewalPreviewRequest represents a RenewalPreviewRequest struct.
type RenewalPreviewRequest struct {
	// An optional array of component definitions to preview. Providing any component definitions here will override the actual components on the subscription (and their quantities), and the billing preview will contain only these components (in addition to any product base fees).
	Components []RenewalPreviewComponent `json:"components,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewRequest.
// It customizes the JSON marshaling process for RenewalPreviewRequest objects.
func (r *RenewalPreviewRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewRequest object to a map representation for JSON marshaling.
func (r *RenewalPreviewRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.Components != nil {
		structMap["components"] = r.Components
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewRequest.
// It customizes the JSON unmarshaling process for RenewalPreviewRequest objects.
func (r *RenewalPreviewRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Components []RenewalPreviewComponent `json:"components,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Components = temp.Components
	return nil
}
