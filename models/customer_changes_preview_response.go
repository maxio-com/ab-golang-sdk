package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CustomerChangesPreviewResponse represents a CustomerChangesPreviewResponse struct.
type CustomerChangesPreviewResponse struct {
    Changes              CustomerChange `json:"changes"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerChangesPreviewResponse.
// It customizes the JSON marshaling process for CustomerChangesPreviewResponse objects.
func (c CustomerChangesPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerChangesPreviewResponse object to a map representation for JSON marshaling.
func (c CustomerChangesPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["changes"] = c.Changes.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChangesPreviewResponse.
// It customizes the JSON unmarshaling process for CustomerChangesPreviewResponse objects.
func (c *CustomerChangesPreviewResponse) UnmarshalJSON(input []byte) error {
    var temp customerChangesPreviewResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "changes")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Changes = *temp.Changes
    return nil
}

// TODO
type customerChangesPreviewResponse  struct {
    Changes *CustomerChange `json:"changes"`
}

func (c *customerChangesPreviewResponse) validate() error {
    var errs []string
    if c.Changes == nil {
        errs = append(errs, "required field `changes` is missing for type `Customer Changes Preview Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
