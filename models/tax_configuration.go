package models

import (
    "encoding/json"
)

// TaxConfiguration represents a TaxConfiguration struct.
type TaxConfiguration struct {
    Kind                 *TaxConfigurationKind  `json:"kind,omitempty"`
    DestinationAddress   *TaxDestinationAddress `json:"destination_address,omitempty"`
    // Returns `true` when Chargify has been properly configured to charge tax using the specified tax system. More details about taxes: https://maxio-chargify.zendesk.com/hc/en-us/articles/5405488905869-Taxes-Introduction
    FullyConfigured      *bool                  `json:"fully_configured,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TaxConfiguration.
// It customizes the JSON marshaling process for TaxConfiguration objects.
func (t TaxConfiguration) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TaxConfiguration object to a map representation for JSON marshaling.
func (t TaxConfiguration) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
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
    var temp taxConfiguration
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "kind", "destination_address", "fully_configured")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.Kind = temp.Kind
    t.DestinationAddress = temp.DestinationAddress
    t.FullyConfigured = temp.FullyConfigured
    return nil
}

// taxConfiguration is a temporary struct used for validating the fields of TaxConfiguration.
type taxConfiguration  struct {
    Kind               *TaxConfigurationKind  `json:"kind,omitempty"`
    DestinationAddress *TaxDestinationAddress `json:"destination_address,omitempty"`
    FullyConfigured    *bool                  `json:"fully_configured,omitempty"`
}
