package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateProductFamilyRequest represents a CreateProductFamilyRequest struct.
type CreateProductFamilyRequest struct {
    ProductFamily        CreateProductFamily `json:"product_family"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductFamilyRequest.
// It customizes the JSON marshaling process for CreateProductFamilyRequest objects.
func (c CreateProductFamilyRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateProductFamilyRequest object to a map representation for JSON marshaling.
func (c CreateProductFamilyRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["product_family"] = c.ProductFamily.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductFamilyRequest.
// It customizes the JSON unmarshaling process for CreateProductFamilyRequest objects.
func (c *CreateProductFamilyRequest) UnmarshalJSON(input []byte) error {
    var temp createProductFamilyRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "product_family")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ProductFamily = *temp.ProductFamily
    return nil
}

// TODO
type createProductFamilyRequest  struct {
    ProductFamily *CreateProductFamily `json:"product_family"`
}

func (c *createProductFamilyRequest) validate() error {
    var errs []string
    if c.ProductFamily == nil {
        errs = append(errs, "required field `product_family` is missing for type `Create Product Family Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
