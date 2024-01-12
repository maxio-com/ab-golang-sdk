package models

import (
    "encoding/json"
)

// BatchJobResponse represents a BatchJobResponse struct.
type BatchJobResponse struct {
    Batchjob BatchJob `json:"batchjob"`
}

// MarshalJSON implements the json.Marshaler interface for BatchJobResponse.
// It customizes the JSON marshaling process for BatchJobResponse objects.
func (b *BatchJobResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BatchJobResponse object to a map representation for JSON marshaling.
func (b *BatchJobResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["batchjob"] = b.Batchjob
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BatchJobResponse.
// It customizes the JSON unmarshaling process for BatchJobResponse objects.
func (b *BatchJobResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Batchjob BatchJob `json:"batchjob"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.Batchjob = temp.Batchjob
    return nil
}
