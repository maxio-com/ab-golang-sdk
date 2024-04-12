package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// BatchJobResponse represents a BatchJobResponse struct.
type BatchJobResponse struct {
    Batchjob             BatchJob       `json:"batchjob"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BatchJobResponse.
// It customizes the JSON marshaling process for BatchJobResponse objects.
func (b BatchJobResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BatchJobResponse object to a map representation for JSON marshaling.
func (b BatchJobResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    structMap["batchjob"] = b.Batchjob.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BatchJobResponse.
// It customizes the JSON unmarshaling process for BatchJobResponse objects.
func (b *BatchJobResponse) UnmarshalJSON(input []byte) error {
    var temp batchJobResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "batchjob")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.Batchjob = *temp.Batchjob
    return nil
}

// TODO
type batchJobResponse  struct {
    Batchjob *BatchJob `json:"batchjob"`
}

func (b *batchJobResponse) validate() error {
    var errs []string
    if b.Batchjob == nil {
        errs = append(errs, "required field `batchjob` is missing for type `Batch Job Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
