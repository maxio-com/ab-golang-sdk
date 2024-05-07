package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CustomerResponse represents a CustomerResponse struct.
type CustomerResponse struct {
    Customer             Customer       `json:"customer"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerResponse.
// It customizes the JSON marshaling process for CustomerResponse objects.
func (c CustomerResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerResponse object to a map representation for JSON marshaling.
func (c CustomerResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["customer"] = c.Customer.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerResponse.
// It customizes the JSON unmarshaling process for CustomerResponse objects.
func (c *CustomerResponse) UnmarshalJSON(input []byte) error {
    var temp customerResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "customer")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Customer = *temp.Customer
    return nil
}

// customerResponse is a temporary struct used for validating the fields of CustomerResponse.
type customerResponse  struct {
    Customer *Customer `json:"customer"`
}

func (c *customerResponse) validate() error {
    var errs []string
    if c.Customer == nil {
        errs = append(errs, "required field `customer` is missing for type `Customer Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
