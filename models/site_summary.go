/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// SiteSummary represents a SiteSummary struct.
type SiteSummary struct {
    SellerName           *string                `json:"seller_name,omitempty"`
    SiteName             *string                `json:"site_name,omitempty"`
    SiteId               *int                   `json:"site_id,omitempty"`
    SiteCurrency         *string                `json:"site_currency,omitempty"`
    Stats                *SiteStatistics        `json:"stats,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSummary) String() string {
    return fmt.Sprintf(
    	"SiteSummary[SellerName=%v, SiteName=%v, SiteId=%v, SiteCurrency=%v, Stats=%v, AdditionalProperties=%v]",
    	s.SellerName, s.SiteName, s.SiteId, s.SiteCurrency, s.Stats, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSummary.
// It customizes the JSON marshaling process for SiteSummary objects.
func (s SiteSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "seller_name", "site_name", "site_id", "site_currency", "stats"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSummary object to a map representation for JSON marshaling.
func (s SiteSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.SellerName != nil {
        structMap["seller_name"] = s.SellerName
    }
    if s.SiteName != nil {
        structMap["site_name"] = s.SiteName
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.SiteCurrency != nil {
        structMap["site_currency"] = s.SiteCurrency
    }
    if s.Stats != nil {
        structMap["stats"] = s.Stats.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSummary.
// It customizes the JSON unmarshaling process for SiteSummary objects.
func (s *SiteSummary) UnmarshalJSON(input []byte) error {
    var temp tempSiteSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "seller_name", "site_name", "site_id", "site_currency", "stats")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.SellerName = temp.SellerName
    s.SiteName = temp.SiteName
    s.SiteId = temp.SiteId
    s.SiteCurrency = temp.SiteCurrency
    s.Stats = temp.Stats
    return nil
}

// tempSiteSummary is a temporary struct used for validating the fields of SiteSummary.
type tempSiteSummary  struct {
    SellerName   *string         `json:"seller_name,omitempty"`
    SiteName     *string         `json:"site_name,omitempty"`
    SiteId       *int            `json:"site_id,omitempty"`
    SiteCurrency *string         `json:"site_currency,omitempty"`
    Stats        *SiteStatistics `json:"stats,omitempty"`
}
