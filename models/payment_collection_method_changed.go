package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PaymentCollectionMethodChanged represents a PaymentCollectionMethodChanged struct.
type PaymentCollectionMethodChanged struct {
    PreviousValue        string         `json:"previous_value"`
    CurrentValue         string         `json:"current_value"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentCollectionMethodChanged.
// It customizes the JSON marshaling process for PaymentCollectionMethodChanged objects.
func (p PaymentCollectionMethodChanged) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentCollectionMethodChanged object to a map representation for JSON marshaling.
func (p PaymentCollectionMethodChanged) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["previous_value"] = p.PreviousValue
    structMap["current_value"] = p.CurrentValue
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentCollectionMethodChanged.
// It customizes the JSON unmarshaling process for PaymentCollectionMethodChanged objects.
func (p *PaymentCollectionMethodChanged) UnmarshalJSON(input []byte) error {
    var temp paymentCollectionMethodChanged
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "previous_value", "current_value")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.PreviousValue = *temp.PreviousValue
    p.CurrentValue = *temp.CurrentValue
    return nil
}

// TODO
type paymentCollectionMethodChanged  struct {
    PreviousValue *string `json:"previous_value"`
    CurrentValue  *string `json:"current_value"`
}

func (p *paymentCollectionMethodChanged) validate() error {
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
    return errors.New(strings.Join(errs, "\n"))
}
