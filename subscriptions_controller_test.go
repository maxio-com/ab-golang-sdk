package ab_golang_sdk

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"maxioadvancedbilling/models"
	"testing"
	"time"
)

// TestSubscriptionsControllerTestCreateSubscription tests the behavior of the SubscriptionsController's
func TestSubscriptionsControllerTestCreateSubscription(t *testing.T) {
	ctx := context.Background()
	var body models.CreateSubscriptionRequest
	err := json.Unmarshal([]byte(`{"subscription":{"product_handle":"basic","customer_attributes":{"first_name":"Joe","last_name":"Blow","email":"joe@example.com","zip":"02120","state":"MA","reference":"XYZ","phone":"(617) 111 - 0000","organization":"Acme","country":"US","city":"Boston","address_2":null,"address":"123 Mass Ave."},"credit_card_attributes":{"last_name":"Smith","first_name":"Joe","full_number":"4111111111111111","expiration_year":"2021","expiration_month":"1","card_type":"visa","billing_zip":"02120","billing_state":"MA","billing_country":"US","billing_city":"Boston","billing_address_2":null,"billing_address":"123 Mass Ave."}}}`), &body)
	if err != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := subscriptionsController.CreateSubscription(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 201)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"subscription":{"id":15236915,"state":"active","balance_in_cents":0,"total_revenue_in_cents":14000,"product_price_in_cents":1000,"product_version_number":7,"current_period_ends_at":"2016-11-15T14:48:10-05:00","next_assessment_at":"2016-11-15T14:48:10-05:00","trial_started_at":null,"trial_ended_at":null,"activated_at":"2016-11-14T14:48:12-05:00","expires_at":null,"created_at":"2016-11-14T14:48:10-05:00","updated_at":"2016-11-14T15:24:41-05:00","cancellation_message":null,"cancellation_method":"merchant_api","cancel_at_end_of_period":null,"canceled_at":null,"current_period_started_at":"2016-11-14T14:48:10-05:00","previous_state":"active","signup_payment_id":162269766,"signup_revenue":"260.00","delayed_cancel_at":null,"coupon_code":"5SNN6HFK3GBH","payment_collection_method":"automatic","snap_day":null,"reason_code":null,"receives_invoice_emails":false,"customer":{"first_name":"Curtis","last_name":"Test","email":"curtis@example.com","cc_emails":"jeff@example.com","organization":"","reference":null,"id":14714298,"created_at":"2016-11-14T14:48:10-05:00","updated_at":"2016-11-14T14:48:13-05:00","address":"123 Anywhere Street","address_2":"","city":"Boulder","state":"CO","zip":"80302","country":"US","phone":"","verified":false,"portal_customer_created_at":"2016-11-14T14:48:13-05:00","portal_invite_last_sent_at":"2016-11-14T14:48:13-05:00","portal_invite_last_accepted_at":null,"tax_exempt":false,"vat_number":"012345678"},"product":{"id":3792003,"name":"$10 Basic Plan","handle":"basic","description":"lorem ipsum","accounting_code":"basic","price_in_cents":1000,"interval":1,"interval_unit":"day","initial_charge_in_cents":null,"expiration_interval":null,"expiration_interval_unit":"never","trial_price_in_cents":null,"trial_interval":null,"trial_interval_unit":"month","initial_charge_after_trial":false,"return_params":"","request_credit_card":false,"require_credit_card":false,"created_at":"2016-03-24T13:38:39-04:00","updated_at":"2016-11-03T13:03:05-04:00","archived_at":null,"update_return_url":"","update_return_params":"","product_family":{"id":527890,"name":"Acme Projects","handle":"billing-plans","accounting_code":null,"description":""},"public_signup_pages":[{"id":281054,"url":"https://general-goods.chargify.com/subscribe/kqvmfrbgd89q/basic"},{"id":281240,"url":"https://general-goods.chargify.com/subscribe/dkffht5dxfd8/basic"},{"id":282694,"url":"https://general-goods.chargify.com/subscribe/jwffwgdd95s8/basic"}],"taxable":false,"version_number":7,"product_price_point_name":"Default"},"credit_card":{"id":10191713,"payment_type":"credit_card","first_name":"Curtis","last_name":"Test","masked_card_number":"XXXX-XXXX-XXXX-1","card_type":"bogus","expiration_month":1,"expiration_year":2026,"billing_address":"123 Anywhere Street","billing_address_2":"","billing_city":"Boulder","billing_state":null,"billing_country":"","billing_zip":"80302","current_vault":"bogus","vault_token":"1","customer_vault_token":null,"customer_id":14714298},"payment_type":"credit_card","referral_code":"w7kjc9","next_product_id":null,"coupon_use_count":1,"coupon_uses_allowed":1,"next_product_handle":null,"stored_credential_transaction_id":125566112256688,"dunning_communication_delay_enabled":true,"dunning_communication_delay_time_zone":"Eastern Time (US & Canada)"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestSubscriptionsControllerTestListSubscriptions tests the behavior of the SubscriptionsController's
func TestSubscriptionsControllerTestListSubscriptions(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	startDate, err := time.Parse(models.DEFAULT_DATE, "2022-07-01")
	if err != nil {
		t.Errorf("Cannot parse the time.")
	}
	endDate, err := time.Parse(models.DEFAULT_DATE, "2022-08-01")
	if err != nil {
		t.Errorf("Cannot parse the time.")
	}
	startDatetime, err := time.Parse(time.RFC3339, "2022-07-01 09:00:05")
	if err != nil {
		t.Errorf("Cannot parse the time.")
	}
	endDatetime, err := time.Parse(time.RFC3339, "2022-08-01 10:00:05")
	if err != nil {
		t.Errorf("Cannot parse the time.")
	}
	sort := models.SubscriptionSortEnum("signup_date")
	apiResponse, err := subscriptionsController.ListSubscriptions(ctx, &page, &perPage, nil, nil, nil, nil, nil, &startDate, &endDate, &startDatetime, &endDatetime, nil, nil, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSubscriptionsControllerTestReadSubscriptionByReference tests the behavior of the SubscriptionsController's
func TestSubscriptionsControllerTestReadSubscriptionByReference(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := subscriptionsController.ReadSubscriptionByReference(ctx, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSubscriptionsControllerTestPreviewSubscription tests the behavior of the SubscriptionsController's
func TestSubscriptionsControllerTestPreviewSubscription(t *testing.T) {
	ctx := context.Background()
	var body models.CreateSubscriptionRequest
	err := json.Unmarshal([]byte(`{"subscription":{"product_handle":"gold-product"}}`), &body)
	if err != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := subscriptionsController.PreviewSubscription(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"subscription_preview":{"current_billing_manifest":{"line_items":[{"transaction_type":"charge","kind":"baseline","amount_in_cents":5000,"memo":"Gold Product (08/21/2018 - 09/21/2018)","discount_amount_in_cents":0,"taxable_amount_in_cents":0,"product_id":1,"product_handle":"gold-product","product_name":"Gold Product","period_range_start":"13 Oct 2023","period_range_end":"13 Nov 2023"},{"transaction_type":"charge","kind":"component","amount_in_cents":28000,"memo":"Component name: 14 Unit names","discount_amount_in_cents":0,"taxable_amount_in_cents":0,"component_id":462149,"component_handle":"handle","component_name":"Component name"},{"transaction_type":"charge","kind":"component","amount_in_cents":2000,"memo":"Fractional Metered Components: 20.0 Fractional Metereds","discount_amount_in_cents":0,"taxable_amount_in_cents":0,"component_id":426665,"component_handle":"handle","component_name":"Fractional Metered Components"},{"transaction_type":"charge","kind":"component","amount_in_cents":0,"memo":"On/Off Component","discount_amount_in_cents":0,"taxable_amount_in_cents":0,"component_id":426670,"component_handle":"handle","component_name":"On/Off Component"},{"transaction_type":"adjustment","kind":"coupon","amount_in_cents":0,"memo":"Coupon: 1DOLLAR - You only get $1.00 off","discount_amount_in_cents":0,"taxable_amount_in_cents":0}],"total_in_cents":35000,"total_discount_in_cents":0,"total_tax_in_cents":0,"subtotal_in_cents":35000,"start_date":"2018-08-21T21:25:21Z","end_date":"2018-09-21T21:25:21Z","period_type":"recurring","existing_balance_in_cents":0},"next_billing_manifest":{"line_items":[{"transaction_type":"charge","kind":"baseline","amount_in_cents":5000,"memo":"Gold Product (09/21/2018 - 10/21/2018)","discount_amount_in_cents":0,"taxable_amount_in_cents":0,"product_id":1,"product_handle":"gold-product","product_name":"Gold Product"},{"transaction_type":"charge","kind":"component","amount_in_cents":28000,"memo":"Component name: 14 Unit names","discount_amount_in_cents":0,"taxable_amount_in_cents":0,"component_id":462149,"component_handle":"handle","component_name":"Component name"},{"transaction_type":"charge","kind":"component","amount_in_cents":0,"memo":"On/Off Component","discount_amount_in_cents":0,"taxable_amount_in_cents":0,"component_id":426670,"component_handle":"handle","component_name":"On/Off Component"}],"total_in_cents":33000,"total_discount_in_cents":0,"total_tax_in_cents":0,"subtotal_in_cents":33000,"start_date":"2018-09-21T21:25:21Z","end_date":"2018-10-21T21:25:21Z","period_type":"recurring","existing_balance_in_cents":0}}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}
