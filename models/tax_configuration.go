// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// TaxConfiguration represents a TaxConfiguration struct.
type TaxConfiguration struct {
    Kind                 *TaxConfigurationKind  `json:"kind,omitempty"`
    DestinationAddress   *TaxDestinationAddress `json:"destination_address,omitempty"`
    // Returns `true` when Chargify has been properly configured to charge tax using the specified tax system. More details about taxes: https://maxio.zendesk.com/hc/en-us/articles/24287012608909-Taxes-Overview
    FullyConfigured      *bool                  `json:"fully_configured,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TaxConfiguration,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TaxConfiguration) String() string {
    return fmt.Sprintf(
    	"TaxConfiguration[Kind=%v, DestinationAddress=%v, FullyConfigured=%v, AdditionalProperties=%v]",
    	t.Kind, t.DestinationAddress, t.FullyConfigured, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TaxConfiguration.
// It customizes the JSON marshaling process for TaxConfiguration objects.
func (t TaxConfiguration) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "kind", "destination_address", "fully_configured"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TaxConfiguration object to a map representation for JSON marshaling.
func (t TaxConfiguration) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Kind != nil {
        structMap["kind"] = t.Kind
    }
    if t.DestinationAddress != nil {
        structMap["destination_address"] = t.DestinationAddress
    }
    if t.FullyConfigured != nil {
        structMap["fully_configured"] = t.FullyConfigured
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TaxConfiguration.
// It customizes the JSON unmarshaling process for TaxConfiguration objects.
func (t *TaxConfiguration) UnmarshalJSON(input []byte) error {
    var temp tempTaxConfiguration
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "kind", "destination_address", "fully_configured")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.Kind = temp.Kind
    t.DestinationAddress = temp.DestinationAddress
    t.FullyConfigured = temp.FullyConfigured
    return nil
}

// tempTaxConfiguration is a temporary struct used for validating the fields of TaxConfiguration.
type tempTaxConfiguration  struct {
    Kind               *TaxConfigurationKind  `json:"kind,omitempty"`
    DestinationAddress *TaxDestinationAddress `json:"destination_address,omitempty"`
    FullyConfigured    *bool                  `json:"fully_configured,omitempty"`
}
