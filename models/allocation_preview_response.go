package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// AllocationPreviewResponse represents a AllocationPreviewResponse struct.
type AllocationPreviewResponse struct {
	AllocationPreview AllocationPreview `json:"allocation_preview"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewResponse.
// It customizes the JSON marshaling process for AllocationPreviewResponse objects.
func (a *AllocationPreviewResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreviewResponse object to a map representation for JSON marshaling.
func (a *AllocationPreviewResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["allocation_preview"] = a.AllocationPreview.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewResponse.
// It customizes the JSON unmarshaling process for AllocationPreviewResponse objects.
func (a *AllocationPreviewResponse) UnmarshalJSON(input []byte) error {
	var temp allocationPreviewResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	a.AllocationPreview = *temp.AllocationPreview
	return nil
}

// TODO
type allocationPreviewResponse struct {
	AllocationPreview *AllocationPreview `json:"allocation_preview"`
}

func (a *allocationPreviewResponse) validate() error {
	var errs []string
	if a.AllocationPreview == nil {
		errs = append(errs, "required field `allocation_preview` is missing for type `Allocation Preview Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
