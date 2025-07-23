// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ServiceCreditResponse represents a ServiceCreditResponse struct.
type ServiceCreditResponse struct {
    ServiceCredit        ServiceCredit          `json:"service_credit"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServiceCreditResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServiceCreditResponse) String() string {
    return fmt.Sprintf(
    	"ServiceCreditResponse[ServiceCredit=%v, AdditionalProperties=%v]",
    	s.ServiceCredit, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServiceCreditResponse.
// It customizes the JSON marshaling process for ServiceCreditResponse objects.
func (s ServiceCreditResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "service_credit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServiceCreditResponse object to a map representation for JSON marshaling.
func (s ServiceCreditResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["service_credit"] = s.ServiceCredit.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceCreditResponse.
// It customizes the JSON unmarshaling process for ServiceCreditResponse objects.
func (s *ServiceCreditResponse) UnmarshalJSON(input []byte) error {
    var temp tempServiceCreditResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "service_credit")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ServiceCredit = *temp.ServiceCredit
    return nil
}

// tempServiceCreditResponse is a temporary struct used for validating the fields of ServiceCreditResponse.
type tempServiceCreditResponse  struct {
    ServiceCredit *ServiceCredit `json:"service_credit"`
}

func (s *tempServiceCreditResponse) validate() error {
    var errs []string
    if s.ServiceCredit == nil {
        errs = append(errs, "required field `service_credit` is missing for type `Service Credit Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
