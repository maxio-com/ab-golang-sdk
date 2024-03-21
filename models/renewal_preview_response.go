package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// RenewalPreviewResponse represents a RenewalPreviewResponse struct.
type RenewalPreviewResponse struct {
	RenewalPreview RenewalPreview `json:"renewal_preview"`
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewResponse.
// It customizes the JSON marshaling process for RenewalPreviewResponse objects.
func (r *RenewalPreviewResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewResponse object to a map representation for JSON marshaling.
func (r *RenewalPreviewResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["renewal_preview"] = r.RenewalPreview.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewResponse.
// It customizes the JSON unmarshaling process for RenewalPreviewResponse objects.
func (r *RenewalPreviewResponse) UnmarshalJSON(input []byte) error {
	var temp renewalPreviewResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	r.RenewalPreview = *temp.RenewalPreview
	return nil
}

// TODO
type renewalPreviewResponse struct {
	RenewalPreview *RenewalPreview `json:"renewal_preview"`
}

func (r *renewalPreviewResponse) validate() error {
	var errs []string
	if r.RenewalPreview == nil {
		errs = append(errs, "required field `renewal_preview` is missing for type `Renewal Preview Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
