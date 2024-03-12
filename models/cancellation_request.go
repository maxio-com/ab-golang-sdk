package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CancellationRequest represents a CancellationRequest struct.
type CancellationRequest struct {
	Subscription CancellationOptions `json:"subscription"`
}

// MarshalJSON implements the json.Marshaler interface for CancellationRequest.
// It customizes the JSON marshaling process for CancellationRequest objects.
func (c *CancellationRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CancellationRequest object to a map representation for JSON marshaling.
func (c *CancellationRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription"] = c.Subscription.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CancellationRequest.
// It customizes the JSON unmarshaling process for CancellationRequest objects.
func (c *CancellationRequest) UnmarshalJSON(input []byte) error {
	var temp cancellationRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.Subscription = *temp.Subscription
	return nil
}

// TODO
type cancellationRequest struct {
	Subscription *CancellationOptions `json:"subscription"`
}

func (c *cancellationRequest) validate() error {
	var errs []string
	if c.Subscription == nil {
		errs = append(errs, "required field `subscription` is missing for type `Cancellation Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
