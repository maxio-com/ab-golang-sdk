package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// TooManyManagementLinkRequests represents a TooManyManagementLinkRequests struct.
type TooManyManagementLinkRequests struct {
    Error                string         `json:"error"`
    NewLinkAvailableAt   time.Time      `json:"new_link_available_at"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TooManyManagementLinkRequests.
// It customizes the JSON marshaling process for TooManyManagementLinkRequests objects.
func (t TooManyManagementLinkRequests) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TooManyManagementLinkRequests object to a map representation for JSON marshaling.
func (t TooManyManagementLinkRequests) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    structMap["error"] = t.Error
    structMap["new_link_available_at"] = t.NewLinkAvailableAt.Format(time.RFC3339)
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TooManyManagementLinkRequests.
// It customizes the JSON unmarshaling process for TooManyManagementLinkRequests objects.
func (t *TooManyManagementLinkRequests) UnmarshalJSON(input []byte) error {
    var temp tooManyManagementLinkRequests
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "error", "new_link_available_at")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.Error = *temp.Error
    NewLinkAvailableAtVal, err := time.Parse(time.RFC3339, *temp.NewLinkAvailableAt)
    if err != nil {
        log.Fatalf("Cannot Parse new_link_available_at as % s format.", time.RFC3339)
    }
    t.NewLinkAvailableAt = NewLinkAvailableAtVal
    return nil
}

// tooManyManagementLinkRequests is a temporary struct used for validating the fields of TooManyManagementLinkRequests.
type tooManyManagementLinkRequests  struct {
    Error              *string `json:"error"`
    NewLinkAvailableAt *string `json:"new_link_available_at"`
}

func (t *tooManyManagementLinkRequests) validate() error {
    var errs []string
    if t.Error == nil {
        errs = append(errs, "required field `error` is missing for type `Too Many Management Link Requests`")
    }
    if t.NewLinkAvailableAt == nil {
        errs = append(errs, "required field `new_link_available_at` is missing for type `Too Many Management Link Requests`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
