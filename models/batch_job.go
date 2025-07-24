// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// BatchJob represents a BatchJob struct.
type BatchJob struct {
    Id                   *int                   `json:"id,omitempty"`
    FinishedAt           Optional[time.Time]    `json:"finished_at"`
    RowCount             Optional[int]          `json:"row_count"`
    CreatedAt            Optional[time.Time]    `json:"created_at"`
    Completed            *string                `json:"completed,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for BatchJob,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BatchJob) String() string {
    return fmt.Sprintf(
    	"BatchJob[Id=%v, FinishedAt=%v, RowCount=%v, CreatedAt=%v, Completed=%v, AdditionalProperties=%v]",
    	b.Id, b.FinishedAt, b.RowCount, b.CreatedAt, b.Completed, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BatchJob.
// It customizes the JSON marshaling process for BatchJob objects.
func (b BatchJob) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "id", "finished_at", "row_count", "created_at", "completed"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BatchJob object to a map representation for JSON marshaling.
func (b BatchJob) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Id != nil {
        structMap["id"] = b.Id
    }
    if b.FinishedAt.IsValueSet() {
        var FinishedAtVal *string = nil
        if b.FinishedAt.Value() != nil {
            val := b.FinishedAt.Value().Format(time.RFC3339)
            FinishedAtVal = &val
        }
        if b.FinishedAt.Value() != nil {
            structMap["finished_at"] = FinishedAtVal
        } else {
            structMap["finished_at"] = nil
        }
    }
    if b.RowCount.IsValueSet() {
        if b.RowCount.Value() != nil {
            structMap["row_count"] = b.RowCount.Value()
        } else {
            structMap["row_count"] = nil
        }
    }
    if b.CreatedAt.IsValueSet() {
        var CreatedAtVal *string = nil
        if b.CreatedAt.Value() != nil {
            val := b.CreatedAt.Value().Format(time.RFC3339)
            CreatedAtVal = &val
        }
        if b.CreatedAt.Value() != nil {
            structMap["created_at"] = CreatedAtVal
        } else {
            structMap["created_at"] = nil
        }
    }
    if b.Completed != nil {
        structMap["completed"] = b.Completed
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BatchJob.
// It customizes the JSON unmarshaling process for BatchJob objects.
func (b *BatchJob) UnmarshalJSON(input []byte) error {
    var temp tempBatchJob
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "finished_at", "row_count", "created_at", "completed")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
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

// tempBatchJob is a temporary struct used for validating the fields of BatchJob.
type tempBatchJob  struct {
    Id         *int             `json:"id,omitempty"`
    FinishedAt Optional[string] `json:"finished_at"`
    RowCount   Optional[int]    `json:"row_count"`
    CreatedAt  Optional[string] `json:"created_at"`
    Completed  *string          `json:"completed,omitempty"`
}
