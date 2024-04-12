package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateOffer represents a CreateOffer struct.
type CreateOffer struct {
    Name                 string                 `json:"name"`
    Handle               string                 `json:"handle"`
    Description          *string                `json:"description,omitempty"`
    ProductId            int                    `json:"product_id"`
    ProductPricePointId  *int                   `json:"product_price_point_id,omitempty"`
    Components           []CreateOfferComponent `json:"components,omitempty"`
    Coupons              []string               `json:"coupons,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOffer.
// It customizes the JSON marshaling process for CreateOffer objects.
func (c CreateOffer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOffer object to a map representation for JSON marshaling.
func (c CreateOffer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["name"] = c.Name
    structMap["handle"] = c.Handle
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    structMap["product_id"] = c.ProductId
    if c.ProductPricePointId != nil {
        structMap["product_price_point_id"] = c.ProductPricePointId
    }
    if c.Components != nil {
        structMap["components"] = c.Components
    }
    if c.Coupons != nil {
        structMap["coupons"] = c.Coupons
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOffer.
// It customizes the JSON unmarshaling process for CreateOffer objects.
func (c *CreateOffer) UnmarshalJSON(input []byte) error {
    var temp createOffer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "handle", "description", "product_id", "product_price_point_id", "components", "coupons")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Name = *temp.Name
    c.Handle = *temp.Handle
    c.Description = temp.Description
    c.ProductId = *temp.ProductId
    c.ProductPricePointId = temp.ProductPricePointId
    c.Components = temp.Components
    c.Coupons = temp.Coupons
    return nil
}

// TODO
type createOffer  struct {
    Name                *string                `json:"name"`
    Handle              *string                `json:"handle"`
    Description         *string                `json:"description,omitempty"`
    ProductId           *int                   `json:"product_id"`
    ProductPricePointId *int                   `json:"product_price_point_id,omitempty"`
    Components          []CreateOfferComponent `json:"components,omitempty"`
    Coupons             []string               `json:"coupons,omitempty"`
}

func (c *createOffer) validate() error {
    var errs []string
    if c.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Create Offer`")
    }
    if c.Handle == nil {
        errs = append(errs, "required field `handle` is missing for type `Create Offer`")
    }
    if c.ProductId == nil {
        errs = append(errs, "required field `product_id` is missing for type `Create Offer`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
