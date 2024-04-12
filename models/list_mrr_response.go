package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListMRRResponse represents a ListMRRResponse struct.
type ListMRRResponse struct {
    Mrr                  ListMRRResponseResult `json:"mrr"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListMRRResponse.
// It customizes the JSON marshaling process for ListMRRResponse objects.
func (l ListMRRResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListMRRResponse object to a map representation for JSON marshaling.
func (l ListMRRResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["mrr"] = l.Mrr.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListMRRResponse.
// It customizes the JSON unmarshaling process for ListMRRResponse objects.
func (l *ListMRRResponse) UnmarshalJSON(input []byte) error {
    var temp listMRRResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mrr")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Mrr = *temp.Mrr
    return nil
}

// TODO
type listMRRResponse  struct {
    Mrr *ListMRRResponseResult `json:"mrr"`
}

func (l *listMRRResponse) validate() error {
    var errs []string
    if l.Mrr == nil {
        errs = append(errs, "required field `mrr` is missing for type `List MRR Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
