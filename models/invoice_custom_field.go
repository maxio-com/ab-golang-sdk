/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// InvoiceCustomField represents a InvoiceCustomField struct.
type InvoiceCustomField struct {
    OwnerId              *int                   `json:"owner_id,omitempty"`
    OwnerType            *CustomFieldOwner      `json:"owner_type,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Value                *string                `json:"value,omitempty"`
    MetadatumId          *int                   `json:"metadatum_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceCustomField,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceCustomField) String() string {
    return fmt.Sprintf(
    	"InvoiceCustomField[OwnerId=%v, OwnerType=%v, Name=%v, Value=%v, MetadatumId=%v, AdditionalProperties=%v]",
    	i.OwnerId, i.OwnerType, i.Name, i.Value, i.MetadatumId, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceCustomField.
// It customizes the JSON marshaling process for InvoiceCustomField objects.
func (i InvoiceCustomField) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "owner_id", "owner_type", "name", "value", "metadatum_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceCustomField object to a map representation for JSON marshaling.
func (i InvoiceCustomField) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.OwnerId != nil {
        structMap["owner_id"] = i.OwnerId
    }
    if i.OwnerType != nil {
        structMap["owner_type"] = i.OwnerType
    }
    if i.Name != nil {
        structMap["name"] = i.Name
    }
    if i.Value != nil {
        structMap["value"] = i.Value
    }
    if i.MetadatumId != nil {
        structMap["metadatum_id"] = i.MetadatumId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceCustomField.
// It customizes the JSON unmarshaling process for InvoiceCustomField objects.
func (i *InvoiceCustomField) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceCustomField
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "owner_id", "owner_type", "name", "value", "metadatum_id")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.OwnerId = temp.OwnerId
    i.OwnerType = temp.OwnerType
    i.Name = temp.Name
    i.Value = temp.Value
    i.MetadatumId = temp.MetadatumId
    return nil
}

// tempInvoiceCustomField is a temporary struct used for validating the fields of InvoiceCustomField.
type tempInvoiceCustomField  struct {
    OwnerId     *int              `json:"owner_id,omitempty"`
    OwnerType   *CustomFieldOwner `json:"owner_type,omitempty"`
    Name        *string           `json:"name,omitempty"`
    Value       *string           `json:"value,omitempty"`
    MetadatumId *int              `json:"metadatum_id,omitempty"`
}
