package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// PrepaymentResponse represents a PrepaymentResponse struct.
type PrepaymentResponse struct {
	Prepayment Prepayment `json:"prepayment"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaymentResponse.
// It customizes the JSON marshaling process for PrepaymentResponse objects.
func (p *PrepaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PrepaymentResponse object to a map representation for JSON marshaling.
func (p *PrepaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["prepayment"] = p.Prepayment.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaymentResponse.
// It customizes the JSON unmarshaling process for PrepaymentResponse objects.
func (p *PrepaymentResponse) UnmarshalJSON(input []byte) error {
	var temp prepaymentResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	p.Prepayment = *temp.Prepayment
	return nil
}

// TODO
type prepaymentResponse struct {
	Prepayment *Prepayment `json:"prepayment"`
}

func (p *prepaymentResponse) validate() error {
	var errs []string
	if p.Prepayment == nil {
		errs = append(errs, "required field `prepayment` is missing for type `Prepayment Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
