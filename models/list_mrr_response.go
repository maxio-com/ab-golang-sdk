// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ListMRRResponse represents a ListMRRResponse struct.
type ListMRRResponse struct {
    Mrr                  ListMRRResponseResult  `json:"mrr"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListMRRResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListMRRResponse) String() string {
    return fmt.Sprintf(
    	"ListMRRResponse[Mrr=%v, AdditionalProperties=%v]",
    	l.Mrr, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListMRRResponse.
// It customizes the JSON marshaling process for ListMRRResponse objects.
func (l ListMRRResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "mrr"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListMRRResponse object to a map representation for JSON marshaling.
func (l ListMRRResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["mrr"] = l.Mrr.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListMRRResponse.
// It customizes the JSON unmarshaling process for ListMRRResponse objects.
func (l *ListMRRResponse) UnmarshalJSON(input []byte) error {
    var temp tempListMRRResponse
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
    l.AdditionalProperties = additionalProperties
    
    l.Mrr = *temp.Mrr
    return nil
}

// tempListMRRResponse is a temporary struct used for validating the fields of ListMRRResponse.
type tempListMRRResponse  struct {
    Mrr *ListMRRResponseResult `json:"mrr"`
}

func (l *tempListMRRResponse) validate() error {
    var errs []string
    if l.Mrr == nil {
        errs = append(errs, "required field `mrr` is missing for type `List MRR Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
