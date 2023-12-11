package ab_golang_sdk

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"maxioadvancedbilling/models"
	"testing"
)

// TestSubscriptionGroupsControllerTestSignupWithSubscriptionGroup tests the behavior of the SubscriptionGroupsController's
func TestSubscriptionGroupsControllerTestSignupWithSubscriptionGroup(t *testing.T) {
	ctx := context.Background()
	var body models.SubscriptionGroupSignupRequest
	err := json.Unmarshal([]byte(`{"subscription_group":{"payment_profile_id":123,"payer_id":123,"subscriptions":[{"product_id":11,"primary":true},{"product_id":12},{"product_id":13}]}}`), &body)
	if err != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := subscriptionGroupsController.SignupWithSubscriptionGroup(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 201)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSubscriptionGroupsControllerTestCreateSubscriptionGroup tests the behavior of the SubscriptionGroupsController's
func TestSubscriptionGroupsControllerTestCreateSubscriptionGroup(t *testing.T) {
	ctx := context.Background()
	var body models.CreateSubscriptionGroupRequest
	err := json.Unmarshal([]byte(`{"subscription_group":{"subscription_id":1,"member_ids":[2,3,4]}}`), &body)
	if err != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := subscriptionGroupsController.CreateSubscriptionGroup(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"subscription_group":{"customer_id":1,"payment_profile":{"id":1,"first_name":"t","last_name":"t","masked_card_number":"XXXX-XXXX-XXXX-1"},"payment_collection_method":"automatic","subscription_ids":[1,2],"created_at":"2021-01-21T05:47:38-05:00"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestSubscriptionGroupsControllerTestListSubscriptionGroups tests the behavior of the SubscriptionGroupsController's
func TestSubscriptionGroupsControllerTestListSubscriptionGroups(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	apiResponse, err := subscriptionGroupsController.ListSubscriptionGroups(ctx, &page, &perPage, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"subscription_groups":[{"uid":"grp_952mvqcnk53wq","scheme":1,"customer_id":88498000,"payment_profile_id":93063018,"subscription_ids":[42768907,82370782],"primary_subscription_id":69844395,"next_assessment_at":"Sun, 09 Aug 2020 15:59:06 EDT -04:00","state":"active","cancel_at_end_of_period":false,"account_balances":{"prepayments":{"balance_in_cents":0},"service_credits":{"balance_in_cents":0},"pending_discounts":{"balance_in_cents":0}}}],"meta":{"current_page":1,"total_count":1}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}
