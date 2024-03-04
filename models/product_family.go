package models

import (
    "encoding/json"
    "log"
    "time"
)

// ProductFamily represents a ProductFamily struct.
type ProductFamily struct {
    Id             *int             `json:"id,omitempty"`
    Name           *string          `json:"name,omitempty"`
    Handle         *string          `json:"handle,omitempty"`
    AccountingCode Optional[string] `json:"accounting_code"`
    Description    Optional[string] `json:"description"`
    CreatedAt      *time.Time       `json:"created_at,omitempty"`
    UpdatedAt      *time.Time       `json:"updated_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProductFamily.
// It customizes the JSON marshaling process for ProductFamily objects.
func (p *ProductFamily) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductFamily object to a map representation for JSON marshaling.
func (p *ProductFamily) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.Name != nil {
        structMap["name"] = p.Name
    }
    if p.Handle != nil {
        structMap["handle"] = p.Handle
    }
    if p.AccountingCode.IsValueSet() {
        structMap["accounting_code"] = p.AccountingCode.Value()
    }
    if p.Description.IsValueSet() {
        structMap["description"] = p.Description.Value()
    }
    if p.CreatedAt != nil {
        structMap["created_at"] = p.CreatedAt.Format(time.RFC3339)
    }
    if p.UpdatedAt != nil {
        structMap["updated_at"] = p.UpdatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductFamily.
// It customizes the JSON unmarshaling process for ProductFamily objects.
func (p *ProductFamily) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id             *int             `json:"id,omitempty"`
        Name           *string          `json:"name,omitempty"`
        Handle         *string          `json:"handle,omitempty"`
        AccountingCode Optional[string] `json:"accounting_code"`
        Description    Optional[string] `json:"description"`
        CreatedAt      *string          `json:"created_at,omitempty"`
        UpdatedAt      *string          `json:"updated_at,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Id = temp.Id
    p.Name = temp.Name
    p.Handle = temp.Handle
    p.AccountingCode = temp.AccountingCode
    p.Description = temp.Description
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        p.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        p.UpdatedAt = &UpdatedAtVal
    }
    return nil
}
