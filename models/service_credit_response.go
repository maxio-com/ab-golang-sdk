package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ServiceCreditResponse represents a ServiceCreditResponse struct.
type ServiceCreditResponse struct {
    ServiceCredit        ServiceCredit  `json:"service_credit"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ServiceCreditResponse.
// It customizes the JSON marshaling process for ServiceCreditResponse objects.
func (s ServiceCreditResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the ServiceCreditResponse object to a map representation for JSON marshaling.
func (s ServiceCreditResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["service_credit"] = s.ServiceCredit.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceCreditResponse.
// It customizes the JSON unmarshaling process for ServiceCreditResponse objects.
func (s *ServiceCreditResponse) UnmarshalJSON(input []byte) error {
    var temp serviceCreditResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "service_credit")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ServiceCredit = *temp.ServiceCredit
    return nil
}

// TODO
type serviceCreditResponse  struct {
    ServiceCredit *ServiceCredit `json:"service_credit"`
}

func (s *serviceCreditResponse) validate() error {
    var errs []string
    if s.ServiceCredit == nil {
        errs = append(errs, "required field `service_credit` is missing for type `Service Credit Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
