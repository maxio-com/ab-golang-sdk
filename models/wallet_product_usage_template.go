// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// WalletProductUsageTemplate represents a WalletProductUsageTemplate struct.
// One wallet template billing through this Wallet Product, joined to the Feature and plan/component that assigns it, when there is one.
type WalletProductUsageTemplate struct {
    // The name of the product or component the Feature is assigned to. `null` when the wallet template has no assigning Feature.
    PlanName             Optional[string]       `json:"plan_name"`
    // The label set in `wallet_behavior.name` on the Feature, or the standalone wallet's own `name`.
    WalletTemplateName   Optional[string]       `json:"wallet_template_name"`
    UnitName             *string                `json:"unit_name,omitempty"`
    // A best-effort label for how this wallet is expected to be replenished. Currently reports `"Overage"` when the wallet template allows a negative balance, and `null` otherwise; there's no persisted field yet distinguishing, for example, an intended "Auto top-up" from "Purchase more" setup. Treat this as informational, not a stable enum.
    PurchaseMechanism    Optional[string]       `json:"purchase_mechanism"`
    PricePointName       Optional[string]       `json:"price_point_name"`
    Currency             *string                `json:"currency,omitempty"`
    // `false` when the assigning Feature is archived.
    Active               *bool                  `json:"active,omitempty"`
    FeatureTemplateId    Optional[int]          `json:"feature_template_id"`
    FeatureTemplateName  Optional[string]       `json:"feature_template_name"`
    // `true` when this wallet template is priced at a price point other than the Wallet Product's default.
    PricePointOverride   *bool                  `json:"price_point_override,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProductUsageTemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProductUsageTemplate) String() string {
    return fmt.Sprintf(
    	"WalletProductUsageTemplate[PlanName=%v, WalletTemplateName=%v, UnitName=%v, PurchaseMechanism=%v, PricePointName=%v, Currency=%v, Active=%v, FeatureTemplateId=%v, FeatureTemplateName=%v, PricePointOverride=%v, AdditionalProperties=%v]",
    	w.PlanName, w.WalletTemplateName, w.UnitName, w.PurchaseMechanism, w.PricePointName, w.Currency, w.Active, w.FeatureTemplateId, w.FeatureTemplateName, w.PricePointOverride, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProductUsageTemplate.
// It customizes the JSON marshaling process for WalletProductUsageTemplate objects.
func (w WalletProductUsageTemplate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "plan_name", "wallet_template_name", "unit_name", "purchase_mechanism", "price_point_name", "currency", "active", "feature_template_id", "feature_template_name", "price_point_override"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProductUsageTemplate object to a map representation for JSON marshaling.
func (w WalletProductUsageTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.PlanName.IsValueSet() {
        if w.PlanName.Value() != nil {
            structMap["plan_name"] = w.PlanName.Value()
        } else {
            structMap["plan_name"] = nil
        }
    }
    if w.WalletTemplateName.IsValueSet() {
        if w.WalletTemplateName.Value() != nil {
            structMap["wallet_template_name"] = w.WalletTemplateName.Value()
        } else {
            structMap["wallet_template_name"] = nil
        }
    }
    if w.UnitName != nil {
        structMap["unit_name"] = w.UnitName
    }
    if w.PurchaseMechanism.IsValueSet() {
        if w.PurchaseMechanism.Value() != nil {
            structMap["purchase_mechanism"] = w.PurchaseMechanism.Value()
        } else {
            structMap["purchase_mechanism"] = nil
        }
    }
    if w.PricePointName.IsValueSet() {
        if w.PricePointName.Value() != nil {
            structMap["price_point_name"] = w.PricePointName.Value()
        } else {
            structMap["price_point_name"] = nil
        }
    }
    if w.Currency != nil {
        structMap["currency"] = w.Currency
    }
    if w.Active != nil {
        structMap["active"] = w.Active
    }
    if w.FeatureTemplateId.IsValueSet() {
        if w.FeatureTemplateId.Value() != nil {
            structMap["feature_template_id"] = w.FeatureTemplateId.Value()
        } else {
            structMap["feature_template_id"] = nil
        }
    }
    if w.FeatureTemplateName.IsValueSet() {
        if w.FeatureTemplateName.Value() != nil {
            structMap["feature_template_name"] = w.FeatureTemplateName.Value()
        } else {
            structMap["feature_template_name"] = nil
        }
    }
    if w.PricePointOverride != nil {
        structMap["price_point_override"] = w.PricePointOverride
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProductUsageTemplate.
// It customizes the JSON unmarshaling process for WalletProductUsageTemplate objects.
func (w *WalletProductUsageTemplate) UnmarshalJSON(input []byte) error {
    var temp tempWalletProductUsageTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "plan_name", "wallet_template_name", "unit_name", "purchase_mechanism", "price_point_name", "currency", "active", "feature_template_id", "feature_template_name", "price_point_override")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.PlanName = temp.PlanName
    w.WalletTemplateName = temp.WalletTemplateName
    w.UnitName = temp.UnitName
    w.PurchaseMechanism = temp.PurchaseMechanism
    w.PricePointName = temp.PricePointName
    w.Currency = temp.Currency
    w.Active = temp.Active
    w.FeatureTemplateId = temp.FeatureTemplateId
    w.FeatureTemplateName = temp.FeatureTemplateName
    w.PricePointOverride = temp.PricePointOverride
    return nil
}

// tempWalletProductUsageTemplate is a temporary struct used for validating the fields of WalletProductUsageTemplate.
type tempWalletProductUsageTemplate  struct {
    PlanName            Optional[string] `json:"plan_name"`
    WalletTemplateName  Optional[string] `json:"wallet_template_name"`
    UnitName            *string          `json:"unit_name,omitempty"`
    PurchaseMechanism   Optional[string] `json:"purchase_mechanism"`
    PricePointName      Optional[string] `json:"price_point_name"`
    Currency            *string          `json:"currency,omitempty"`
    Active              *bool            `json:"active,omitempty"`
    FeatureTemplateId   Optional[int]    `json:"feature_template_id"`
    FeatureTemplateName Optional[string] `json:"feature_template_name"`
    PricePointOverride  *bool            `json:"price_point_override,omitempty"`
}
