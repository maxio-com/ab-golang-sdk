package models

import (
    "encoding/json"
    "log"
    "time"
)

// BatchJob represents a BatchJob struct.
type BatchJob struct {
    Id         *int                `json:"id,omitempty"`
    FinishedAt Optional[time.Time] `json:"finished_at"`
    RowCount   Optional[int]       `json:"row_count"`
    CreatedAt  Optional[time.Time] `json:"created_at"`
    Completed  *string             `json:"completed,omitempty"`
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
        var FinishedAtVal *string = nil
        if b.FinishedAt.Value() != nil {
            val := b.FinishedAt.Value().Format(time.RFC3339)
            FinishedAtVal = &val
        }
        structMap["finished_at"] = FinishedAtVal
    }
    if b.RowCount.IsValueSet() {
        structMap["row_count"] = b.RowCount.Value()
    }
    if b.CreatedAt.IsValueSet() {
        var CreatedAtVal *string = nil
        if b.CreatedAt.Value() != nil {
            val := b.CreatedAt.Value().Format(time.RFC3339)
            CreatedAtVal = &val
        }
        structMap["created_at"] = CreatedAtVal
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
    b.FinishedAt.ShouldSetValue(temp.FinishedAt.IsValueSet())
    if temp.FinishedAt.Value() != nil {
        FinishedAtVal, err := time.Parse(time.RFC3339, (*temp.FinishedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse finished_at as % s format.", time.RFC3339)
        }
        b.FinishedAt.SetValue(&FinishedAtVal)
    }
    b.RowCount = temp.RowCount
    b.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
    if temp.CreatedAt.Value() != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        b.CreatedAt.SetValue(&CreatedAtVal)
    }
    b.Completed = temp.Completed
    return nil
}
