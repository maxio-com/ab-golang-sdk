package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateComponentPricePointRequest represents a CreateComponentPricePointRequest struct.
type CreateComponentPricePointRequest struct {
    PricePoint           CreateComponentPricePointRequestPricePoint `json:"price_point"`
    AdditionalProperties map[string]any                             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateComponentPricePointRequest.
// It customizes the JSON marshaling process for CreateComponentPricePointRequest objects.
func (c CreateComponentPricePointRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateComponentPricePointRequest object to a map representation for JSON marshaling.
func (c CreateComponentPricePointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["price_point"] = c.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateComponentPricePointRequest.
// It customizes the JSON unmarshaling process for CreateComponentPricePointRequest objects.
func (c *CreateComponentPricePointRequest) UnmarshalJSON(input []byte) error {
    var temp createComponentPricePointRequest
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

// createComponentPricePointRequest is a temporary struct used for validating the fields of CreateComponentPricePointRequest.
type createComponentPricePointRequest  struct {
    PricePoint *CreateComponentPricePointRequestPricePoint `json:"price_point"`
}

func (c *createComponentPricePointRequest) validate() error {
    var errs []string
    if c.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Create Component Price Point Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
