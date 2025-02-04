/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionProductChange represents a SubscriptionProductChange struct.
type SubscriptionProductChange struct {
    PreviousProductId    int                    `json:"previous_product_id"`
    NewProductId         int                    `json:"new_product_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionProductChange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionProductChange) String() string {
    return fmt.Sprintf(
    	"SubscriptionProductChange[PreviousProductId=%v, NewProductId=%v, AdditionalProperties=%v]",
    	s.PreviousProductId, s.NewProductId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionProductChange.
// It customizes the JSON marshaling process for SubscriptionProductChange objects.
func (s SubscriptionProductChange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "previous_product_id", "new_product_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionProductChange object to a map representation for JSON marshaling.
func (s SubscriptionProductChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["previous_product_id"] = s.PreviousProductId
    structMap["new_product_id"] = s.NewProductId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionProductChange.
// It customizes the JSON unmarshaling process for SubscriptionProductChange objects.
func (s *SubscriptionProductChange) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionProductChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "previous_product_id", "new_product_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.PreviousProductId = *temp.PreviousProductId
    s.NewProductId = *temp.NewProductId
    return nil
}

// tempSubscriptionProductChange is a temporary struct used for validating the fields of SubscriptionProductChange.
type tempSubscriptionProductChange  struct {
    PreviousProductId *int `json:"previous_product_id"`
    NewProductId      *int `json:"new_product_id"`
}

func (s *tempSubscriptionProductChange) validate() error {
    var errs []string
    if s.PreviousProductId == nil {
        errs = append(errs, "required field `previous_product_id` is missing for type `Subscription Product Change`")
    }
    if s.NewProductId == nil {
        errs = append(errs, "required field `new_product_id` is missing for type `Subscription Product Change`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
