package models

import (
	"encoding/json"
)

// SiteSummary represents a SiteSummary struct.
type SiteSummary struct {
	SellerName   *string         `json:"seller_name,omitempty"`
	SiteName     *string         `json:"site_name,omitempty"`
	SiteId       *int            `json:"site_id,omitempty"`
	SiteCurrency *string         `json:"site_currency,omitempty"`
	Stats        *SiteStatistics `json:"stats,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSummary.
// It customizes the JSON marshaling process for SiteSummary objects.
func (s *SiteSummary) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSummary object to a map representation for JSON marshaling.
func (s *SiteSummary) toMap() map[string]any {
	structMap := make(map[string]any)
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
		structMap["stats"] = s.Stats
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSummary.
// It customizes the JSON unmarshaling process for SiteSummary objects.
func (s *SiteSummary) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SellerName   *string         `json:"seller_name,omitempty"`
		SiteName     *string         `json:"site_name,omitempty"`
		SiteId       *int            `json:"site_id,omitempty"`
		SiteCurrency *string         `json:"site_currency,omitempty"`
		Stats        *SiteStatistics `json:"stats,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.SellerName = temp.SellerName
	s.SiteName = temp.SiteName
	s.SiteId = temp.SiteId
	s.SiteCurrency = temp.SiteCurrency
	s.Stats = temp.Stats
	return nil
}
