package models

import (
    "encoding/json"
)

// SubscriptionPreviewResponse represents a SubscriptionPreviewResponse struct.
type SubscriptionPreviewResponse struct {
    SubscriptionPreview SubscriptionPreview `json:"subscription_preview"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionPreviewResponse.
// It customizes the JSON marshaling process for SubscriptionPreviewResponse objects.
func (s *SubscriptionPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionPreviewResponse object to a map representation for JSON marshaling.
func (s *SubscriptionPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["subscription_preview"] = s.SubscriptionPreview.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionPreviewResponse.
// It customizes the JSON unmarshaling process for SubscriptionPreviewResponse objects.
func (s *SubscriptionPreviewResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        SubscriptionPreview SubscriptionPreview `json:"subscription_preview"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.SubscriptionPreview = temp.SubscriptionPreview
    return nil
}
