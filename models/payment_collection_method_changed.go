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

// PaymentCollectionMethodChanged represents a PaymentCollectionMethodChanged struct.
type PaymentCollectionMethodChanged struct {
    PreviousValue        string                 `json:"previous_value"`
    CurrentValue         string                 `json:"current_value"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PaymentCollectionMethodChanged,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaymentCollectionMethodChanged) String() string {
    return fmt.Sprintf(
    	"PaymentCollectionMethodChanged[PreviousValue=%v, CurrentValue=%v, AdditionalProperties=%v]",
    	p.PreviousValue, p.CurrentValue, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PaymentCollectionMethodChanged.
// It customizes the JSON marshaling process for PaymentCollectionMethodChanged objects.
func (p PaymentCollectionMethodChanged) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "previous_value", "current_value"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentCollectionMethodChanged object to a map representation for JSON marshaling.
func (p PaymentCollectionMethodChanged) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["previous_value"] = p.PreviousValue
    structMap["current_value"] = p.CurrentValue
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentCollectionMethodChanged.
// It customizes the JSON unmarshaling process for PaymentCollectionMethodChanged objects.
func (p *PaymentCollectionMethodChanged) UnmarshalJSON(input []byte) error {
    var temp tempPaymentCollectionMethodChanged
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "previous_value", "current_value")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.PreviousValue = *temp.PreviousValue
    p.CurrentValue = *temp.CurrentValue
    return nil
}

// tempPaymentCollectionMethodChanged is a temporary struct used for validating the fields of PaymentCollectionMethodChanged.
type tempPaymentCollectionMethodChanged  struct {
    PreviousValue *string `json:"previous_value"`
    CurrentValue  *string `json:"current_value"`
}

func (p *tempPaymentCollectionMethodChanged) validate() error {
    var errs []string
    if p.PreviousValue == nil {
        errs = append(errs, "required field `previous_value` is missing for type `Payment Collection Method Changed`")
    }
    if p.CurrentValue == nil {
        errs = append(errs, "required field `current_value` is missing for type `Payment Collection Method Changed`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
