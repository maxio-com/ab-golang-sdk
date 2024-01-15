package models

import (
    "encoding/json"
)

// ListMRRResponse represents a ListMRRResponse struct.
type ListMRRResponse struct {
    Mrr ListMRRResponseResult `json:"mrr"`
}

// MarshalJSON implements the json.Marshaler interface for ListMRRResponse.
// It customizes the JSON marshaling process for ListMRRResponse objects.
func (l *ListMRRResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListMRRResponse object to a map representation for JSON marshaling.
func (l *ListMRRResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["mrr"] = l.Mrr
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListMRRResponse.
// It customizes the JSON unmarshaling process for ListMRRResponse objects.
func (l *ListMRRResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Mrr ListMRRResponseResult `json:"mrr"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.Mrr = temp.Mrr
    return nil
}
