/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionPreviewResponse represents a SubscriptionPreviewResponse struct.
type SubscriptionPreviewResponse struct {
    SubscriptionPreview  SubscriptionPreview    `json:"subscription_preview"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionPreviewResponse.
// It customizes the JSON marshaling process for SubscriptionPreviewResponse objects.
func (s SubscriptionPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "subscription_preview"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionPreviewResponse object to a map representation for JSON marshaling.
func (s SubscriptionPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["subscription_preview"] = s.SubscriptionPreview.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionPreviewResponse.
// It customizes the JSON unmarshaling process for SubscriptionPreviewResponse objects.
func (s *SubscriptionPreviewResponse) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionPreviewResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_preview")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.SubscriptionPreview = *temp.SubscriptionPreview
    return nil
}

// tempSubscriptionPreviewResponse is a temporary struct used for validating the fields of SubscriptionPreviewResponse.
type tempSubscriptionPreviewResponse  struct {
    SubscriptionPreview *SubscriptionPreview `json:"subscription_preview"`
}

func (s *tempSubscriptionPreviewResponse) validate() error {
    var errs []string
    if s.SubscriptionPreview == nil {
        errs = append(errs, "required field `subscription_preview` is missing for type `Subscription Preview Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
