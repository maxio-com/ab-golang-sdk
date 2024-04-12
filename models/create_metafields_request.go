package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateMetafieldsRequest represents a CreateMetafieldsRequest struct.
type CreateMetafieldsRequest struct {
    Metafields           CreateMetafieldsRequestMetafields `json:"metafields"`
    AdditionalProperties map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMetafieldsRequest.
// It customizes the JSON marshaling process for CreateMetafieldsRequest objects.
func (c CreateMetafieldsRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMetafieldsRequest object to a map representation for JSON marshaling.
func (c CreateMetafieldsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["metafields"] = c.Metafields.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMetafieldsRequest.
// It customizes the JSON unmarshaling process for CreateMetafieldsRequest objects.
func (c *CreateMetafieldsRequest) UnmarshalJSON(input []byte) error {
    var temp createMetafieldsRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "metafields")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Metafields = *temp.Metafields
    return nil
}

// TODO
type createMetafieldsRequest  struct {
    Metafields *CreateMetafieldsRequestMetafields `json:"metafields"`
}

func (c *createMetafieldsRequest) validate() error {
    var errs []string
    if c.Metafields == nil {
        errs = append(errs, "required field `metafields` is missing for type `Create Metafields Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
