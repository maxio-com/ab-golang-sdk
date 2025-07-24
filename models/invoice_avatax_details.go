// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// InvoiceAvataxDetails represents a InvoiceAvataxDetails struct.
type InvoiceAvataxDetails struct {
    Id                   Optional[int64]        `json:"id"`
    Status               Optional[string]       `json:"status"`
    DocumentCode         Optional[string]       `json:"document_code"`
    CommitDate           Optional[time.Time]    `json:"commit_date"`
    ModifyDate           Optional[time.Time]    `json:"modify_date"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceAvataxDetails,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceAvataxDetails) String() string {
    return fmt.Sprintf(
    	"InvoiceAvataxDetails[Id=%v, Status=%v, DocumentCode=%v, CommitDate=%v, ModifyDate=%v, AdditionalProperties=%v]",
    	i.Id, i.Status, i.DocumentCode, i.CommitDate, i.ModifyDate, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceAvataxDetails.
// It customizes the JSON marshaling process for InvoiceAvataxDetails objects.
func (i InvoiceAvataxDetails) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "id", "status", "document_code", "commit_date", "modify_date"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceAvataxDetails object to a map representation for JSON marshaling.
func (i InvoiceAvataxDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Id.IsValueSet() {
        if i.Id.Value() != nil {
            structMap["id"] = i.Id.Value()
        } else {
            structMap["id"] = nil
        }
    }
    if i.Status.IsValueSet() {
        if i.Status.Value() != nil {
            structMap["status"] = i.Status.Value()
        } else {
            structMap["status"] = nil
        }
    }
    if i.DocumentCode.IsValueSet() {
        if i.DocumentCode.Value() != nil {
            structMap["document_code"] = i.DocumentCode.Value()
        } else {
            structMap["document_code"] = nil
        }
    }
    if i.CommitDate.IsValueSet() {
        var CommitDateVal *string = nil
        if i.CommitDate.Value() != nil {
            val := i.CommitDate.Value().Format(time.RFC3339)
            CommitDateVal = &val
        }
        if i.CommitDate.Value() != nil {
            structMap["commit_date"] = CommitDateVal
        } else {
            structMap["commit_date"] = nil
        }
    }
    if i.ModifyDate.IsValueSet() {
        var ModifyDateVal *string = nil
        if i.ModifyDate.Value() != nil {
            val := i.ModifyDate.Value().Format(time.RFC3339)
            ModifyDateVal = &val
        }
        if i.ModifyDate.Value() != nil {
            structMap["modify_date"] = ModifyDateVal
        } else {
            structMap["modify_date"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceAvataxDetails.
// It customizes the JSON unmarshaling process for InvoiceAvataxDetails objects.
func (i *InvoiceAvataxDetails) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceAvataxDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "status", "document_code", "commit_date", "modify_date")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Id = temp.Id
    i.Status = temp.Status
    i.DocumentCode = temp.DocumentCode
    i.CommitDate.ShouldSetValue(temp.CommitDate.IsValueSet())
    if temp.CommitDate.Value() != nil {
        CommitDateVal, err := time.Parse(time.RFC3339, (*temp.CommitDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse commit_date as % s format.", time.RFC3339)
        }
        i.CommitDate.SetValue(&CommitDateVal)
    }
    i.ModifyDate.ShouldSetValue(temp.ModifyDate.IsValueSet())
    if temp.ModifyDate.Value() != nil {
        ModifyDateVal, err := time.Parse(time.RFC3339, (*temp.ModifyDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse modify_date as % s format.", time.RFC3339)
        }
        i.ModifyDate.SetValue(&ModifyDateVal)
    }
    return nil
}

// tempInvoiceAvataxDetails is a temporary struct used for validating the fields of InvoiceAvataxDetails.
type tempInvoiceAvataxDetails  struct {
    Id           Optional[int64]  `json:"id"`
    Status       Optional[string] `json:"status"`
    DocumentCode Optional[string] `json:"document_code"`
    CommitDate   Optional[string] `json:"commit_date"`
    ModifyDate   Optional[string] `json:"modify_date"`
}
