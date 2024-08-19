/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateOfferRequest represents a CreateOfferRequest struct.
type CreateOfferRequest struct {
    Offer                CreateOffer    `json:"offer"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOfferRequest.
// It customizes the JSON marshaling process for CreateOfferRequest objects.
func (c CreateOfferRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOfferRequest object to a map representation for JSON marshaling.
func (c CreateOfferRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["offer"] = c.Offer.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOfferRequest.
// It customizes the JSON unmarshaling process for CreateOfferRequest objects.
func (c *CreateOfferRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateOfferRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "offer")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Offer = *temp.Offer
    return nil
}

// tempCreateOfferRequest is a temporary struct used for validating the fields of CreateOfferRequest.
type tempCreateOfferRequest  struct {
    Offer *CreateOffer `json:"offer"`
}

func (c *tempCreateOfferRequest) validate() error {
    var errs []string
    if c.Offer == nil {
        errs = append(errs, "required field `offer` is missing for type `Create Offer Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
