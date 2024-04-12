package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateComponentPricePointsRequest represents a CreateComponentPricePointsRequest struct.
type CreateComponentPricePointsRequest struct {
    PricePoints          []CreateComponentPricePointsRequestPricePoints `json:"price_points"`
    AdditionalProperties map[string]any                                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateComponentPricePointsRequest.
// It customizes the JSON marshaling process for CreateComponentPricePointsRequest objects.
func (c CreateComponentPricePointsRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateComponentPricePointsRequest object to a map representation for JSON marshaling.
func (c CreateComponentPricePointsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["price_points"] = c.PricePoints
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateComponentPricePointsRequest.
// It customizes the JSON unmarshaling process for CreateComponentPricePointsRequest objects.
func (c *CreateComponentPricePointsRequest) UnmarshalJSON(input []byte) error {
    var temp createComponentPricePointsRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "price_points")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.PricePoints = *temp.PricePoints
    return nil
}

// TODO
type createComponentPricePointsRequest  struct {
    PricePoints *[]CreateComponentPricePointsRequestPricePoints `json:"price_points"`
}

func (c *createComponentPricePointsRequest) validate() error {
    var errs []string
    if c.PricePoints == nil {
        errs = append(errs, "required field `price_points` is missing for type `Create Component Price Points Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
