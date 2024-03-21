package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// MRRResponse represents a MRRResponse struct.
type MRRResponse struct {
	Mrr MRR `json:"mrr"`
}

// MarshalJSON implements the json.Marshaler interface for MRRResponse.
// It customizes the JSON marshaling process for MRRResponse objects.
func (m *MRRResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MRRResponse object to a map representation for JSON marshaling.
func (m *MRRResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["mrr"] = m.Mrr.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MRRResponse.
// It customizes the JSON unmarshaling process for MRRResponse objects.
func (m *MRRResponse) UnmarshalJSON(input []byte) error {
	var temp mrrResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	m.Mrr = *temp.Mrr
	return nil
}

// TODO
type mrrResponse struct {
	Mrr *MRR `json:"mrr"`
}

func (m *mrrResponse) validate() error {
	var errs []string
	if m.Mrr == nil {
		errs = append(errs, "required field `mrr` is missing for type `MRR response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
