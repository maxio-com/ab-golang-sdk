// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// Site represents a Site struct.
type Site struct {
	Id                                      *int                   `json:"id,omitempty"`
	Name                                    *string                `json:"name,omitempty"`
	Subdomain                               *string                `json:"subdomain,omitempty"`
	Currency                                *string                `json:"currency,omitempty"`
	SellerId                                *int                   `json:"seller_id,omitempty"`
	NonPrimaryCurrencies                    []string               `json:"non_primary_currencies,omitempty"`
	RelationshipInvoicingEnabled            *bool                  `json:"relationship_invoicing_enabled,omitempty"`
	ScheduleSubscriptionCancellationEnabled *bool                  `json:"schedule_subscription_cancellation_enabled,omitempty"`
	CustomerHierarchyEnabled                *bool                  `json:"customer_hierarchy_enabled,omitempty"`
	WhopaysEnabled                          *bool                  `json:"whopays_enabled,omitempty"`
	WhopaysDefaultPayer                     *string                `json:"whopays_default_payer,omitempty"`
	AllocationSettings                      *AllocationSettings    `json:"allocation_settings,omitempty"`
	DefaultPaymentCollectionMethod          *string                `json:"default_payment_collection_method,omitempty"`
	OrganizationAddress                     *OrganizationAddress   `json:"organization_address,omitempty"`
	TaxConfiguration                        *TaxConfiguration      `json:"tax_configuration,omitempty"`
	NetTerms                                *NetTerms              `json:"net_terms,omitempty"`
	Test                                    *bool                  `json:"test,omitempty"`
	AdditionalProperties                    map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Site,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Site) String() string {
	return fmt.Sprintf(
		"Site[Id=%v, Name=%v, Subdomain=%v, Currency=%v, SellerId=%v, NonPrimaryCurrencies=%v, RelationshipInvoicingEnabled=%v, ScheduleSubscriptionCancellationEnabled=%v, CustomerHierarchyEnabled=%v, WhopaysEnabled=%v, WhopaysDefaultPayer=%v, AllocationSettings=%v, DefaultPaymentCollectionMethod=%v, OrganizationAddress=%v, TaxConfiguration=%v, NetTerms=%v, Test=%v, AdditionalProperties=%v]",
		s.Id, s.Name, s.Subdomain, s.Currency, s.SellerId, s.NonPrimaryCurrencies, s.RelationshipInvoicingEnabled, s.ScheduleSubscriptionCancellationEnabled, s.CustomerHierarchyEnabled, s.WhopaysEnabled, s.WhopaysDefaultPayer, s.AllocationSettings, s.DefaultPaymentCollectionMethod, s.OrganizationAddress, s.TaxConfiguration, s.NetTerms, s.Test, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Site.
// It customizes the JSON marshaling process for Site objects.
func (s Site) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"id", "name", "subdomain", "currency", "seller_id", "non_primary_currencies", "relationship_invoicing_enabled", "schedule_subscription_cancellation_enabled", "customer_hierarchy_enabled", "whopays_enabled", "whopays_default_payer", "allocation_settings", "default_payment_collection_method", "organization_address", "tax_configuration", "net_terms", "test"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the Site object to a map representation for JSON marshaling.
func (s Site) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Id != nil {
		structMap["id"] = s.Id
	}
	if s.Name != nil {
		structMap["name"] = s.Name
	}
	if s.Subdomain != nil {
		structMap["subdomain"] = s.Subdomain
	}
	if s.Currency != nil {
		structMap["currency"] = s.Currency
	}
	if s.SellerId != nil {
		structMap["seller_id"] = s.SellerId
	}
	if s.NonPrimaryCurrencies != nil {
		structMap["non_primary_currencies"] = s.NonPrimaryCurrencies
	}
	if s.RelationshipInvoicingEnabled != nil {
		structMap["relationship_invoicing_enabled"] = s.RelationshipInvoicingEnabled
	}
	if s.ScheduleSubscriptionCancellationEnabled != nil {
		structMap["schedule_subscription_cancellation_enabled"] = s.ScheduleSubscriptionCancellationEnabled
	}
	if s.CustomerHierarchyEnabled != nil {
		structMap["customer_hierarchy_enabled"] = s.CustomerHierarchyEnabled
	}
	if s.WhopaysEnabled != nil {
		structMap["whopays_enabled"] = s.WhopaysEnabled
	}
	if s.WhopaysDefaultPayer != nil {
		structMap["whopays_default_payer"] = s.WhopaysDefaultPayer
	}
	if s.AllocationSettings != nil {
		structMap["allocation_settings"] = s.AllocationSettings.toMap()
	}
	if s.DefaultPaymentCollectionMethod != nil {
		structMap["default_payment_collection_method"] = s.DefaultPaymentCollectionMethod
	}
	if s.OrganizationAddress != nil {
		structMap["organization_address"] = s.OrganizationAddress.toMap()
	}
	if s.TaxConfiguration != nil {
		structMap["tax_configuration"] = s.TaxConfiguration.toMap()
	}
	if s.NetTerms != nil {
		structMap["net_terms"] = s.NetTerms.toMap()
	}
	if s.Test != nil {
		structMap["test"] = s.Test
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Site.
// It customizes the JSON unmarshaling process for Site objects.
func (s *Site) UnmarshalJSON(input []byte) error {
	var temp tempSite
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "name", "subdomain", "currency", "seller_id", "non_primary_currencies", "relationship_invoicing_enabled", "schedule_subscription_cancellation_enabled", "customer_hierarchy_enabled", "whopays_enabled", "whopays_default_payer", "allocation_settings", "default_payment_collection_method", "organization_address", "tax_configuration", "net_terms", "test")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Id = temp.Id
	s.Name = temp.Name
	s.Subdomain = temp.Subdomain
	s.Currency = temp.Currency
	s.SellerId = temp.SellerId
	s.NonPrimaryCurrencies = temp.NonPrimaryCurrencies
	s.RelationshipInvoicingEnabled = temp.RelationshipInvoicingEnabled
	s.ScheduleSubscriptionCancellationEnabled = temp.ScheduleSubscriptionCancellationEnabled
	s.CustomerHierarchyEnabled = temp.CustomerHierarchyEnabled
	s.WhopaysEnabled = temp.WhopaysEnabled
	s.WhopaysDefaultPayer = temp.WhopaysDefaultPayer
	s.AllocationSettings = temp.AllocationSettings
	s.DefaultPaymentCollectionMethod = temp.DefaultPaymentCollectionMethod
	s.OrganizationAddress = temp.OrganizationAddress
	s.TaxConfiguration = temp.TaxConfiguration
	s.NetTerms = temp.NetTerms
	s.Test = temp.Test
	return nil
}

// tempSite is a temporary struct used for validating the fields of Site.
type tempSite struct {
	Id                                      *int                 `json:"id,omitempty"`
	Name                                    *string              `json:"name,omitempty"`
	Subdomain                               *string              `json:"subdomain,omitempty"`
	Currency                                *string              `json:"currency,omitempty"`
	SellerId                                *int                 `json:"seller_id,omitempty"`
	NonPrimaryCurrencies                    []string             `json:"non_primary_currencies,omitempty"`
	RelationshipInvoicingEnabled            *bool                `json:"relationship_invoicing_enabled,omitempty"`
	ScheduleSubscriptionCancellationEnabled *bool                `json:"schedule_subscription_cancellation_enabled,omitempty"`
	CustomerHierarchyEnabled                *bool                `json:"customer_hierarchy_enabled,omitempty"`
	WhopaysEnabled                          *bool                `json:"whopays_enabled,omitempty"`
	WhopaysDefaultPayer                     *string              `json:"whopays_default_payer,omitempty"`
	AllocationSettings                      *AllocationSettings  `json:"allocation_settings,omitempty"`
	DefaultPaymentCollectionMethod          *string              `json:"default_payment_collection_method,omitempty"`
	OrganizationAddress                     *OrganizationAddress `json:"organization_address,omitempty"`
	TaxConfiguration                        *TaxConfiguration    `json:"tax_configuration,omitempty"`
	NetTerms                                *NetTerms            `json:"net_terms,omitempty"`
	Test                                    *bool                `json:"test,omitempty"`
}
