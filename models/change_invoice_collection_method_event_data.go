package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ChangeInvoiceCollectionMethodEventData represents a ChangeInvoiceCollectionMethodEventData struct.
// Example schema for an `change_invoice_collection_method` event
type ChangeInvoiceCollectionMethodEventData struct {
    // The previous collection method of the invoice.
    FromCollectionMethod string         `json:"from_collection_method"`
    // The new collection method of the invoice.
    ToCollectionMethod   string         `json:"to_collection_method"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ChangeInvoiceCollectionMethodEventData.
// It customizes the JSON marshaling process for ChangeInvoiceCollectionMethodEventData objects.
func (c ChangeInvoiceCollectionMethodEventData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ChangeInvoiceCollectionMethodEventData object to a map representation for JSON marshaling.
func (c ChangeInvoiceCollectionMethodEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["from_collection_method"] = c.FromCollectionMethod
    structMap["to_collection_method"] = c.ToCollectionMethod
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChangeInvoiceCollectionMethodEventData.
// It customizes the JSON unmarshaling process for ChangeInvoiceCollectionMethodEventData objects.
func (c *ChangeInvoiceCollectionMethodEventData) UnmarshalJSON(input []byte) error {
    var temp changeInvoiceCollectionMethodEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "from_collection_method", "to_collection_method")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.FromCollectionMethod = *temp.FromCollectionMethod
    c.ToCollectionMethod = *temp.ToCollectionMethod
    return nil
}

// TODO
type changeInvoiceCollectionMethodEventData  struct {
    FromCollectionMethod *string `json:"from_collection_method"`
    ToCollectionMethod   *string `json:"to_collection_method"`
}

func (c *changeInvoiceCollectionMethodEventData) validate() error {
    var errs []string
    if c.FromCollectionMethod == nil {
        errs = append(errs, "required field `from_collection_method` is missing for type `Change Invoice Collection Method Event Data`")
    }
    if c.ToCollectionMethod == nil {
        errs = append(errs, "required field `to_collection_method` is missing for type `Change Invoice Collection Method Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
