package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// UsageResponse represents a UsageResponse struct.
type UsageResponse struct {
	Usage Usage `json:"usage"`
}

// MarshalJSON implements the json.Marshaler interface for UsageResponse.
// It customizes the JSON marshaling process for UsageResponse objects.
func (u *UsageResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UsageResponse object to a map representation for JSON marshaling.
func (u *UsageResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["usage"] = u.Usage.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UsageResponse.
// It customizes the JSON unmarshaling process for UsageResponse objects.
func (u *UsageResponse) UnmarshalJSON(input []byte) error {
	var temp usageResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	u.Usage = *temp.Usage
	return nil
}

// TODO
type usageResponse struct {
	Usage *Usage `json:"usage"`
}

func (u *usageResponse) validate() error {
	var errs []string
	if u.Usage == nil {
		errs = append(errs, "required field `usage` is missing for type `Usage Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
