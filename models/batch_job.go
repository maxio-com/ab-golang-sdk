package models

import (
    "encoding/json"
)

// BatchJob represents a BatchJob struct.
type BatchJob struct {
    Id         *int             `json:"id,omitempty"`
    FinishedAt Optional[string] `json:"finished_at"`
    RowCount   Optional[int]    `json:"row_count"`
    CreatedAt  Optional[string] `json:"created_at"`
    Completed  *string          `json:"completed,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BatchJob.
// It customizes the JSON marshaling process for BatchJob objects.
func (b *BatchJob) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BatchJob object to a map representation for JSON marshaling.
func (b *BatchJob) toMap() map[string]any {
    structMap := make(map[string]any)
    if b.Id != nil {
        structMap["id"] = b.Id
    }
    if b.FinishedAt.IsValueSet() {
        structMap["finished_at"] = b.FinishedAt.Value()
    }
    if b.RowCount.IsValueSet() {
        structMap["row_count"] = b.RowCount.Value()
    }
    if b.CreatedAt.IsValueSet() {
        structMap["created_at"] = b.CreatedAt.Value()
    }
    if b.Completed != nil {
        structMap["completed"] = b.Completed
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BatchJob.
// It customizes the JSON unmarshaling process for BatchJob objects.
func (b *BatchJob) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id         *int             `json:"id,omitempty"`
        FinishedAt Optional[string] `json:"finished_at"`
        RowCount   Optional[int]    `json:"row_count"`
        CreatedAt  Optional[string] `json:"created_at"`
        Completed  *string          `json:"completed,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.Id = temp.Id
    b.FinishedAt = temp.FinishedAt
    b.RowCount = temp.RowCount
    b.CreatedAt = temp.CreatedAt
    b.Completed = temp.Completed
    return nil
}
