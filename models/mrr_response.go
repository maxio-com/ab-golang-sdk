/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// MRRResponse represents a MRRResponse struct.
type MRRResponse struct {
    Mrr                  MRR                    `json:"mrr"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MRRResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MRRResponse) String() string {
    return fmt.Sprintf(
    	"MRRResponse[Mrr=%v, AdditionalProperties=%v]",
    	m.Mrr, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MRRResponse.
// It customizes the JSON marshaling process for MRRResponse objects.
func (m MRRResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "mrr"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MRRResponse object to a map representation for JSON marshaling.
func (m MRRResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["mrr"] = m.Mrr.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MRRResponse.
// It customizes the JSON unmarshaling process for MRRResponse objects.
func (m *MRRResponse) UnmarshalJSON(input []byte) error {
    var temp tempMRRResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mrr")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Mrr = *temp.Mrr
    return nil
}

// tempMRRResponse is a temporary struct used for validating the fields of MRRResponse.
type tempMRRResponse  struct {
    Mrr *MRR `json:"mrr"`
}

func (m *tempMRRResponse) validate() error {
    var errs []string
    if m.Mrr == nil {
        errs = append(errs, "required field `mrr` is missing for type `MRR response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
