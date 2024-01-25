package models

import (
	"encoding/json"
)

// RenewalPreviewComponent represents a RenewalPreviewComponent struct.
type RenewalPreviewComponent struct {
	// Either the component's Chargify id or its handle prefixed with `handle:`
	ComponentId *interface{} `json:"component_id,omitempty"`
	// The quantity for which you wish to preview billing. This is useful if you want to preview a predicted, higher usage value than is currently present on the subscription.
	// This quantity represents:
	// - Whether or not an on/off component is enabled - use 0 for disabled or 1 for enabled
	// - The desired allocated_quantity for a quantity-based component
	// - The desired unit_balance for a metered component
	// - The desired metric quantity for an events-based component
	Quantity *int `json:"quantity,omitempty"`
	// Either the component price point's Chargify id or its handle prefixed with `handle:`
	PricePointId *interface{} `json:"price_point_id,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewComponent.
// It customizes the JSON marshaling process for RenewalPreviewComponent objects.
func (r *RenewalPreviewComponent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewComponent object to a map representation for JSON marshaling.
func (r *RenewalPreviewComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.ComponentId != nil {
		structMap["component_id"] = r.ComponentId
	}
	if r.Quantity != nil {
		structMap["quantity"] = r.Quantity
	}
	if r.PricePointId != nil {
		structMap["price_point_id"] = r.PricePointId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewComponent.
// It customizes the JSON unmarshaling process for RenewalPreviewComponent objects.
func (r *RenewalPreviewComponent) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ComponentId  *interface{} `json:"component_id,omitempty"`
		Quantity     *int         `json:"quantity,omitempty"`
		PricePointId *interface{} `json:"price_point_id,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.ComponentId = temp.ComponentId
	r.Quantity = temp.Quantity
	r.PricePointId = temp.PricePointId
	return nil
}
