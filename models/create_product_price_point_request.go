package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateProductPricePointRequest represents a CreateProductPricePointRequest struct.
type CreateProductPricePointRequest struct {
    PricePoint           CreateProductPricePoint `json:"price_point"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductPricePointRequest.
// It customizes the JSON marshaling process for CreateProductPricePointRequest objects.
func (c CreateProductPricePointRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateProductPricePointRequest object to a map representation for JSON marshaling.
func (c CreateProductPricePointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["price_point"] = c.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductPricePointRequest.
// It customizes the JSON unmarshaling process for CreateProductPricePointRequest objects.
func (c *CreateProductPricePointRequest) UnmarshalJSON(input []byte) error {
    var temp createProductPricePointRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "price_point")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.PricePoint = *temp.PricePoint
    return nil
}

// createProductPricePointRequest is a temporary struct used for validating the fields of CreateProductPricePointRequest.
type createProductPricePointRequest  struct {
    PricePoint *CreateProductPricePoint `json:"price_point"`
}

func (c *createProductPricePointRequest) validate() error {
    var errs []string
    if c.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Create Product Price Point Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
