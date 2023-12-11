package models

import (
	"encoding/json"
)

// TaxConfiguration represents a TaxConfiguration struct.
type TaxConfiguration struct {
	Kind               *TaxConfigurationKindEnum  `json:"kind,omitempty"`
	DestinationAddress *TaxDestinationAddressEnum `json:"destination_address,omitempty"`
	// Returns `true` when Chargify has been properly configured to charge tax using the specified tax system. More details about taxes: https://maxio-chargify.zendesk.com/hc/en-us/articles/5405488905869-Taxes-Introduction
	FullyConfigured *bool `json:"fully_configured,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for TaxConfiguration.
// It customizes the JSON marshaling process for TaxConfiguration objects.
func (t *TaxConfiguration) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TaxConfiguration object to a map representation for JSON marshaling.
func (t *TaxConfiguration) toMap() map[string]any {
	structMap := make(map[string]any)
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
	temp := &struct {
		Kind               *TaxConfigurationKindEnum  `json:"kind,omitempty"`
		DestinationAddress *TaxDestinationAddressEnum `json:"destination_address,omitempty"`
		FullyConfigured    *bool                      `json:"fully_configured,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.Kind = temp.Kind
	t.DestinationAddress = temp.DestinationAddress
	t.FullyConfigured = temp.FullyConfigured
	return nil
}
