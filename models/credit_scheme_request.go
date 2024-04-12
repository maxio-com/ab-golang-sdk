package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreditSchemeRequest represents a CreditSchemeRequest struct.
type CreditSchemeRequest struct {
    CreditScheme         CreditScheme   `json:"credit_scheme"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreditSchemeRequest.
// It customizes the JSON marshaling process for CreditSchemeRequest objects.
func (c CreditSchemeRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreditSchemeRequest object to a map representation for JSON marshaling.
func (c CreditSchemeRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["credit_scheme"] = c.CreditScheme
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditSchemeRequest.
// It customizes the JSON unmarshaling process for CreditSchemeRequest objects.
func (c *CreditSchemeRequest) UnmarshalJSON(input []byte) error {
    var temp creditSchemeRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "credit_scheme")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.CreditScheme = *temp.CreditScheme
    return nil
}

// TODO
type creditSchemeRequest  struct {
    CreditScheme *CreditScheme `json:"credit_scheme"`
}

func (c *creditSchemeRequest) validate() error {
    var errs []string
    if c.CreditScheme == nil {
        errs = append(errs, "required field `credit_scheme` is missing for type `Credit Scheme Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
