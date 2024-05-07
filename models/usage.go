package models

import (
    "encoding/json"
    "log"
    "time"
)

// Usage represents a Usage struct.
type Usage struct {
    Id                   *int64           `json:"id,omitempty"`
    Memo                 Optional[string] `json:"memo"`
    CreatedAt            *time.Time       `json:"created_at,omitempty"`
    PricePointId         *int             `json:"price_point_id,omitempty"`
    Quantity             *UsageQuantity   `json:"quantity,omitempty"`
    OverageQuantity      *int             `json:"overage_quantity,omitempty"`
    ComponentId          *int             `json:"component_id,omitempty"`
    ComponentHandle      *string          `json:"component_handle,omitempty"`
    SubscriptionId       *int             `json:"subscription_id,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Usage.
// It customizes the JSON marshaling process for Usage objects.
func (u Usage) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the Usage object to a map representation for JSON marshaling.
func (u Usage) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.Memo.IsValueSet() {
        if u.Memo.Value() != nil {
            structMap["memo"] = u.Memo.Value()
        } else {
            structMap["memo"] = nil
        }
    }
    if u.CreatedAt != nil {
        structMap["created_at"] = u.CreatedAt.Format(time.RFC3339)
    }
    if u.PricePointId != nil {
        structMap["price_point_id"] = u.PricePointId
    }
    if u.Quantity != nil {
        structMap["quantity"] = u.Quantity.toMap()
    }
    if u.OverageQuantity != nil {
        structMap["overage_quantity"] = u.OverageQuantity
    }
    if u.ComponentId != nil {
        structMap["component_id"] = u.ComponentId
    }
    if u.ComponentHandle != nil {
        structMap["component_handle"] = u.ComponentHandle
    }
    if u.SubscriptionId != nil {
        structMap["subscription_id"] = u.SubscriptionId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Usage.
// It customizes the JSON unmarshaling process for Usage objects.
func (u *Usage) UnmarshalJSON(input []byte) error {
    var temp usage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "memo", "created_at", "price_point_id", "quantity", "overage_quantity", "component_id", "component_handle", "subscription_id")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Id = temp.Id
    u.Memo = temp.Memo
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        u.CreatedAt = &CreatedAtVal
    }
    u.PricePointId = temp.PricePointId
    u.Quantity = temp.Quantity
    u.OverageQuantity = temp.OverageQuantity
    u.ComponentId = temp.ComponentId
    u.ComponentHandle = temp.ComponentHandle
    u.SubscriptionId = temp.SubscriptionId
    return nil
}

// usage is a temporary struct used for validating the fields of Usage.
type usage  struct {
    Id              *int64           `json:"id,omitempty"`
    Memo            Optional[string] `json:"memo"`
    CreatedAt       *string          `json:"created_at,omitempty"`
    PricePointId    *int             `json:"price_point_id,omitempty"`
    Quantity        *UsageQuantity   `json:"quantity,omitempty"`
    OverageQuantity *int             `json:"overage_quantity,omitempty"`
    ComponentId     *int             `json:"component_id,omitempty"`
    ComponentHandle *string          `json:"component_handle,omitempty"`
    SubscriptionId  *int             `json:"subscription_id,omitempty"`
}
