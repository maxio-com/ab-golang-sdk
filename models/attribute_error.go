package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// AttributeError represents a AttributeError struct.
type AttributeError struct {
    Attribute            []string       `json:"attribute"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AttributeError.
// It customizes the JSON marshaling process for AttributeError objects.
func (a AttributeError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AttributeError object to a map representation for JSON marshaling.
func (a AttributeError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["attribute"] = a.Attribute
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AttributeError.
// It customizes the JSON unmarshaling process for AttributeError objects.
func (a *AttributeError) UnmarshalJSON(input []byte) error {
    var temp attributeError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "attribute")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Attribute = *temp.Attribute
    return nil
}

// attributeError is a temporary struct used for validating the fields of AttributeError.
type attributeError  struct {
    Attribute *[]string `json:"attribute"`
}

func (a *attributeError) validate() error {
    var errs []string
    if a.Attribute == nil {
        errs = append(errs, "required field `attribute` is missing for type `Attribute Error`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
