package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SiteResponse represents a SiteResponse struct.
type SiteResponse struct {
    Site                 Site           `json:"site"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteResponse.
// It customizes the JSON marshaling process for SiteResponse objects.
func (s SiteResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteResponse object to a map representation for JSON marshaling.
func (s SiteResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["site"] = s.Site.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteResponse.
// It customizes the JSON unmarshaling process for SiteResponse objects.
func (s *SiteResponse) UnmarshalJSON(input []byte) error {
    var temp siteResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "site")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Site = *temp.Site
    return nil
}

// TODO
type siteResponse  struct {
    Site *Site `json:"site"`
}

func (s *siteResponse) validate() error {
    var errs []string
    if s.Site == nil {
        errs = append(errs, "required field `site` is missing for type `Site Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
