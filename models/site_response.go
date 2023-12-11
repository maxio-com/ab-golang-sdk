package models

import (
	"encoding/json"
)

// SiteResponse represents a SiteResponse struct.
type SiteResponse struct {
	Site Site `json:"site"`
}

// MarshalJSON implements the json.Marshaler interface for SiteResponse.
// It customizes the JSON marshaling process for SiteResponse objects.
func (s *SiteResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteResponse object to a map representation for JSON marshaling.
func (s *SiteResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["site"] = s.Site
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteResponse.
// It customizes the JSON unmarshaling process for SiteResponse objects.
func (s *SiteResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Site Site `json:"site"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Site = temp.Site
	return nil
}
