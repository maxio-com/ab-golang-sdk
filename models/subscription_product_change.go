// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// SubscriptionProductChange represents a SubscriptionProductChange struct.
// Event data for both `subscription_product_change` and `subscription_product_change_scheduled`. The price point and `effective_at` fields are only populated for scheduled changes.
type SubscriptionProductChange struct {
    PreviousProductId           int                    `json:"previous_product_id"`
    NewProductId                int                    `json:"new_product_id"`
    PreviousProductPricePointId Optional[int]          `json:"previous_product_price_point_id"`
    NewProductPricePointId      Optional[int]          `json:"new_product_price_point_id"`
    // When the scheduled product change takes effect (the subscription's next renewal). Only sent for `subscription_product_change_scheduled`.
    EffectiveAt                 Optional[time.Time]    `json:"effective_at"`
    AdditionalProperties        map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionProductChange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionProductChange) String() string {
    return fmt.Sprintf(
    	"SubscriptionProductChange[PreviousProductId=%v, NewProductId=%v, PreviousProductPricePointId=%v, NewProductPricePointId=%v, EffectiveAt=%v, AdditionalProperties=%v]",
    	s.PreviousProductId, s.NewProductId, s.PreviousProductPricePointId, s.NewProductPricePointId, s.EffectiveAt, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionProductChange.
// It customizes the JSON marshaling process for SubscriptionProductChange objects.
func (s SubscriptionProductChange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "previous_product_id", "new_product_id", "previous_product_price_point_id", "new_product_price_point_id", "effective_at"); err != nil {
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
    if s.PreviousProductPricePointId.IsValueSet() {
        if s.PreviousProductPricePointId.Value() != nil {
            structMap["previous_product_price_point_id"] = s.PreviousProductPricePointId.Value()
        } else {
            structMap["previous_product_price_point_id"] = nil
        }
    }
    if s.NewProductPricePointId.IsValueSet() {
        if s.NewProductPricePointId.Value() != nil {
            structMap["new_product_price_point_id"] = s.NewProductPricePointId.Value()
        } else {
            structMap["new_product_price_point_id"] = nil
        }
    }
    if s.EffectiveAt.IsValueSet() {
        var EffectiveAtVal *string = nil
        if s.EffectiveAt.Value() != nil {
            val := s.EffectiveAt.Value().Format(time.RFC3339)
            EffectiveAtVal = &val
        }
        if s.EffectiveAt.Value() != nil {
            structMap["effective_at"] = EffectiveAtVal
        } else {
            structMap["effective_at"] = nil
        }
    }
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "previous_product_id", "new_product_id", "previous_product_price_point_id", "new_product_price_point_id", "effective_at")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.PreviousProductId = *temp.PreviousProductId
    s.NewProductId = *temp.NewProductId
    s.PreviousProductPricePointId = temp.PreviousProductPricePointId
    s.NewProductPricePointId = temp.NewProductPricePointId
    s.EffectiveAt.ShouldSetValue(temp.EffectiveAt.IsValueSet())
    if temp.EffectiveAt.Value() != nil {
        EffectiveAtVal, err := time.Parse(time.RFC3339, (*temp.EffectiveAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse effective_at as % s format.", time.RFC3339)
        }
        s.EffectiveAt.SetValue(&EffectiveAtVal)
    }
    return nil
}

// tempSubscriptionProductChange is a temporary struct used for validating the fields of SubscriptionProductChange.
type tempSubscriptionProductChange  struct {
    PreviousProductId           *int             `json:"previous_product_id"`
    NewProductId                *int             `json:"new_product_id"`
    PreviousProductPricePointId Optional[int]    `json:"previous_product_price_point_id"`
    NewProductPricePointId      Optional[int]    `json:"new_product_price_point_id"`
    EffectiveAt                 Optional[string] `json:"effective_at"`
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
