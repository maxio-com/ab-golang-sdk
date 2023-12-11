package ab_golang_sdk

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"maxioadvancedbilling/models"
	"testing"
)

// TestPaymentProfilesControllerTestCreatePaymentProfile tests the behavior of the PaymentProfilesController's
func TestPaymentProfilesControllerTestCreatePaymentProfile(t *testing.T) {
	ctx := context.Background()
	var body models.CreatePaymentProfileRequest
	err := json.Unmarshal([]byte(`{"payment_profile":{"customer_id":123,"bank_name":"Best Bank","bank_routing_number":"021000089","bank_account_number":"111111111111","bank_account_type":"checking","bank_account_holder_type":"business","payment_type":"bank_account"}}`), &body)
	if err != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := paymentProfilesController.CreatePaymentProfile(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"payment_profile":{"first_name":"Jessica","last_name":"Test","card_type":"visa","expiration_month":10,"expiration_year":2018,"customer_id":19195410,"current_vault":"bogus","vault_token":"1","billing_address":"123 Main St.","billing_city":"Boston","billing_state":"MA","billing_zip":"02120","billing_country":"US","customer_vault_token":null,"billing_address_2":null,"payment_type":"credit_card","site_gateway_setting_id":1,"gateway_handle":"handle"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestPaymentProfilesControllerTestListPaymentProfiles tests the behavior of the PaymentProfilesController's
func TestPaymentProfilesControllerTestListPaymentProfiles(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	apiResponse, err := paymentProfilesController.ListPaymentProfiles(ctx, &page, &perPage, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"payment_profile":{"id":10089892,"first_name":"Chester","last_name":"Tester","customer_id":14543792,"current_vault":"bogus","vault_token":"0011223344","billing_address":"456 Juniper Court","billing_city":"Boulder","billing_state":"CO","billing_zip":"80302","billing_country":"US","customer_vault_token":null,"billing_address_2":"","bank_name":"Bank of Kansas City","masked_bank_routing_number":"XXXX6789","masked_bank_account_number":"XXXX3344","bank_account_type":"checking","bank_account_holder_type":"personal","payment_type":"bank_account","site_gateway_setting_id":1,"gateway_handle":"handle"}},{"payment_profile":{"id":10188522,"first_name":"Frankie","last_name":"Tester","customer_id":14543712,"current_vault":"bogus","vault_token":"123456789","billing_address":"123 Montana Way","billing_city":"Los Angeles","billing_state":"CA","billing_zip":"90210","billing_country":"US","customer_vault_token":null,"billing_address_2":"","bank_name":"Bank of Kansas City","masked_bank_routing_number":"XXXX6789","masked_bank_account_number":"XXXX6789","bank_account_type":"checking","bank_account_holder_type":"personal","payment_type":"bank_account","site_gateway_setting_id":1,"gateway_handle":"handle"}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}
