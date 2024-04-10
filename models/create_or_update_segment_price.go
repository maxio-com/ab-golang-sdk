package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateOrUpdateSegmentPrice represents a CreateOrUpdateSegmentPrice struct.
type CreateOrUpdateSegmentPrice struct {
    StartingQuantity     *int                                `json:"starting_quantity,omitempty"`
    EndingQuantity       *int                                `json:"ending_quantity,omitempty"`
    // The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
    UnitPrice            CreateOrUpdateSegmentPriceUnitPrice `json:"unit_price"`
    AdditionalProperties map[string]any                      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateSegmentPrice.
// It customizes the JSON marshaling process for CreateOrUpdateSegmentPrice objects.
func (c CreateOrUpdateSegmentPrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateSegmentPrice object to a map representation for JSON marshaling.
func (c CreateOrUpdateSegmentPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.StartingQuantity != nil {
        structMap["starting_quantity"] = c.StartingQuantity
    }
    if c.EndingQuantity != nil {
        structMap["ending_quantity"] = c.EndingQuantity
    }
    structMap["unit_price"] = c.UnitPrice.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateSegmentPrice.
// It customizes the JSON unmarshaling process for CreateOrUpdateSegmentPrice objects.
func (c *CreateOrUpdateSegmentPrice) UnmarshalJSON(input []byte) error {
    var temp createOrUpdateSegmentPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "starting_quantity", "ending_quantity", "unit_price")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.StartingQuantity = temp.StartingQuantity
    c.EndingQuantity = temp.EndingQuantity
    c.UnitPrice = *temp.UnitPrice
    return nil
}

// TODO
type createOrUpdateSegmentPrice  struct {
    StartingQuantity *int                                 `json:"starting_quantity,omitempty"`
    EndingQuantity   *int                                 `json:"ending_quantity,omitempty"`
    UnitPrice        *CreateOrUpdateSegmentPriceUnitPrice `json:"unit_price"`
}

func (c *createOrUpdateSegmentPrice) validate() error {
    var errs []string
    if c.UnitPrice == nil {
        errs = append(errs, "required field `unit_price` is missing for type `Create or Update Segment Price`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
