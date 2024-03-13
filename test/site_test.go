package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/apimatic/go-core-runtime/https"
	advancedbilling "github.com/maxio-com/ab-golang-sdk"
	"github.com/maxio-com/ab-golang-sdk/models"
	"github.com/stretchr/testify/suite"
)

type SiteSuite struct {
	APISuite
}

func TestSiteSuite(t *testing.T) {
	suite.Run(t, new(SiteSuite))
}

func (s *SiteSuite) TestReadSite() {
	cases := []struct {
		name        string
		client      advancedbilling.ClientInterface
		assert      func(*testing.T, models.ApiResponse[models.SiteResponse], error)
		expectedErr bool
	}{
		{
			name:        "success",
			client:      s.client,
			expectedErr: false,
			assert: func(t *testing.T, resp models.ApiResponse[models.SiteResponse], err error) {
				s.Equal(http.StatusOK, resp.Response.StatusCode, "status code")
				s.NoError(err)

				respSite := resp.Data.Site
				s.Equal(4718, *respSite.Id, "ID")
				s.Equal("GO SDK env", *respSite.Name, "Name")
				s.Equal("go-sdk", *respSite.Subdomain, "Subdomain")
				s.Equal("USD", *respSite.Currency, "Currency")
				s.Equal(722159, *respSite.SellerId, "SellerID")
				s.EqualValues([]string{}, respSite.NonPrimaryCurrencies, "NonPrimaryCurrencies")
				s.True(*respSite.RelationshipInvoicingEnabled, "RelationshipInvoiceEnabled")
				s.False(*respSite.CustomerHierarchyEnabled, "CustomerHierarchyEnabled,")
				s.False(*respSite.WhopaysEnabled, "WhopaysEnabled")
				s.Equal("self-ungrouped", *respSite.WhopaysDefaultPayer, "WhopaysDefaultPayer")
				s.Equal(string(models.CollectionMethod_AUTOMATIC), *respSite.DefaultPaymentCollectionMethod, "DefaultPaymentCollectionMethod")

				allocationSettings := respSite.AllocationSettings
				s.Equal(models.CreditType_PRORATED, *allocationSettings.UpgradeCharge.Value(), "UpgradeCharge")
				s.Equal(models.CreditType_NONE, *allocationSettings.DowngradeCredit.Value(), "DowngradeCredit")
				s.Equal("true", *allocationSettings.AccrueCharge, "AccrueCharge")

				organizationAddress := respSite.OrganizationAddress
				s.Equal("Asdf Street", *organizationAddress.Street.Value(), "Street")
				s.Equal("123/444", *organizationAddress.Line2.Value(), "Line2")
				s.Equal("San Antonio", *organizationAddress.City.Value(), "City")
				s.Equal("TX", *organizationAddress.State.Value(), "State")
				s.Equal("78015", *organizationAddress.Zip.Value(), "Zip")
				s.Equal("US", *organizationAddress.Country.Value(), "Country")
				s.Equal("Developer Experience", *organizationAddress.Name.Value(), "AddressName")
				s.Equal("555 111 222", *organizationAddress.Phone.Value(), "Phone")

				taxConfiguration := respSite.TaxConfiguration
				s.Equal(models.TaxConfigurationKind_CUSTOM, *taxConfiguration.Kind, "Kind")
				s.Equal(models.TaxDestinationAddress_SHIPPINGTHENBILLING, *taxConfiguration.DestinationAddress, "TaxDestinationAddress")
				s.False(*taxConfiguration.FullyConfigured, "FullyConfigured")

				netTerms := respSite.NetTerms
				s.Equal(0, *netTerms.DefaultNetTerms, "DefaultNetTerms")
				s.Equal(0, *netTerms.AutomaticNetTerms, "AutomaticNetTerms")
				s.Equal(0, *netTerms.RemittanceNetTerms, "RemittanceNetTerms")
				s.False(*netTerms.NetTermsOnRemittanceSignupsEnabled, "NetTermsOnRemittanceSignupsEnabled")
				s.False(*netTerms.CustomNetTermsEnabled, "CustomNetTermsEnabled")

				s.True(*respSite.Test, "Test")
			},
		},
		{
			name:   "unauthorized",
			client: s.unauthorizedClient,
			assert: func(t *testing.T, resp models.ApiResponse[models.SiteResponse], err error) {
				s.Equal(err.Error(), "ApiError occured: HTTP Response Not OK. Status code: 401. Response: 'HTTP Basic: Access denied.\n'.")
				actualErr, ok := err.(https.ApiError)
				s.Equal(http.StatusUnauthorized, actualErr.StatusCode)
				s.True(ok)
			},
			expectedErr: true,
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			resp, err := c.client.SitesController().ReadSite(context.Background())
			if !c.expectedErr {
				s.NoErrorf(err, "could not read site")
			}
			c.assert(t, resp, err)
		})
	}

}
