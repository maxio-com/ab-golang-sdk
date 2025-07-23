// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateOrUpdateSegmentPrice represents a CreateOrUpdateSegmentPrice struct.
type CreateOrUpdateSegmentPrice struct {
    StartingQuantity     *int                                `json:"starting_quantity,omitempty"`
    EndingQuantity       *int                                `json:"ending_quantity,omitempty"`
    // The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
    UnitPrice            CreateOrUpdateSegmentPriceUnitPrice `json:"unit_price"`
    AdditionalProperties map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for CreateOrUpdateSegmentPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateOrUpdateSegmentPrice) String() string {
    return fmt.Sprintf(
    	"CreateOrUpdateSegmentPrice[StartingQuantity=%v, EndingQuantity=%v, UnitPrice=%v, AdditionalProperties=%v]",
    	c.StartingQuantity, c.EndingQuantity, c.UnitPrice, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateSegmentPrice.
// It customizes the JSON marshaling process for CreateOrUpdateSegmentPrice objects.
func (c CreateOrUpdateSegmentPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "starting_quantity", "ending_quantity", "unit_price"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateSegmentPrice object to a map representation for JSON marshaling.
func (c CreateOrUpdateSegmentPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp tempCreateOrUpdateSegmentPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "starting_quantity", "ending_quantity", "unit_price")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.StartingQuantity = temp.StartingQuantity
    c.EndingQuantity = temp.EndingQuantity
    c.UnitPrice = *temp.UnitPrice
    return nil
}

// tempCreateOrUpdateSegmentPrice is a temporary struct used for validating the fields of CreateOrUpdateSegmentPrice.
type tempCreateOrUpdateSegmentPrice  struct {
    StartingQuantity *int                                 `json:"starting_quantity,omitempty"`
    EndingQuantity   *int                                 `json:"ending_quantity,omitempty"`
    UnitPrice        *CreateOrUpdateSegmentPriceUnitPrice `json:"unit_price"`
}

func (c *tempCreateOrUpdateSegmentPrice) validate() error {
    var errs []string
    if c.UnitPrice == nil {
        errs = append(errs, "required field `unit_price` is missing for type `Create or Update Segment Price`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
