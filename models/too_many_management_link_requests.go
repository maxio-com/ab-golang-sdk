package models

import (
	"encoding/json"
	"log"
	"time"
)

// TooManyManagementLinkRequests represents a TooManyManagementLinkRequests struct.
type TooManyManagementLinkRequests struct {
	Error              string    `json:"error"`
	NewLinkAvailableAt time.Time `json:"new_link_available_at"`
}

// MarshalJSON implements the json.Marshaler interface for TooManyManagementLinkRequests.
// It customizes the JSON marshaling process for TooManyManagementLinkRequests objects.
func (t *TooManyManagementLinkRequests) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TooManyManagementLinkRequests object to a map representation for JSON marshaling.
func (t *TooManyManagementLinkRequests) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["error"] = t.Error
	structMap["new_link_available_at"] = t.NewLinkAvailableAt.Format(time.RFC3339)
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TooManyManagementLinkRequests.
// It customizes the JSON unmarshaling process for TooManyManagementLinkRequests objects.
func (t *TooManyManagementLinkRequests) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Error              string `json:"error"`
		NewLinkAvailableAt string `json:"new_link_available_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.Error = temp.Error
	NewLinkAvailableAtVal, err := time.Parse(time.RFC3339, temp.NewLinkAvailableAt)
	if err != nil {
		log.Fatalf("Cannot Parse new_link_available_at as % s format.", time.RFC3339)
	}
	t.NewLinkAvailableAt = NewLinkAvailableAtVal
	return nil
}
