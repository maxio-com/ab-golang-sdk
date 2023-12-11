package ab_golang_sdk

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"maxioadvancedbilling/models"
	"testing"
)

// TestSitesControllerTestReadSite tests the behavior of the SitesController's
func TestSitesControllerTestReadSite(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := sitesController.ReadSite(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"site":{"id":0,"name":"string","subdomain":"string","currency":"string","seller_id":0,"non_primary_currencies":["string"],"relationship_invoicing_enabled":true,"customer_hierarchy_enabled":true,"whopays_enabled":true,"whopays_default_payer":"string","default_payment_collection_method":"string","organization_address":{"street":null,"line2":null,"city":null,"state":null,"zip":null,"country":null,"name":"string","phone":"string"},"tax_configuration":{"kind":"custom","fully_configured":true,"destination_address":"shipping_then_billing"},"net_terms":{"default_net_terms":0,"automatic_net_terms":0,"remittance_net_terms":0,"net_terms_on_remittance_signups_enabled":false,"custom_net_terms_enabled":false},"test":true}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestSitesControllerTestClearSite tests the behavior of the SitesController's
func TestSitesControllerTestClearSite(t *testing.T) {
	ctx := context.Background()
	cleanupScope := models.CleanupScopeEnum("all")
	resp, err := sitesController.ClearSite(ctx, &cleanupScope)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesControllerTestListChargifyJsPublicKeys tests the behavior of the SitesController's
func TestSitesControllerTestListChargifyJsPublicKeys(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	apiResponse, err := sitesController.ListChargifyJsPublicKeys(ctx, &page, &perPage)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"chargify_js_keys":[{"public_key":"chjs_ftrxt7c4fv6f74wchjs_5zyn7gnwv","requires_security_token":false,"created_at":"2021-01-01T05:00:00-04:00"}],"meta":{"total_count":1,"current_page":1,"total_pages":1,"per_page":10}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}
