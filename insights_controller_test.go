package ab_golang_sdk

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"maxioadvancedbilling/models"
	"testing"
)

// TestInsightsControllerTestReadSiteStats tests the behavior of the InsightsController's
func TestInsightsControllerTestReadSiteStats(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := insightsController.ReadSiteStats(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"seller_name":"Acme, Inc.","site_name":"Production","site_id":12345,"site_currency":"USD","stats":{"total_subscriptions":120,"subscriptions_today":4,"total_revenue":"$45,978.81","revenue_today":"$1,405.12","revenue_this_month":"$10,000.00","revenue_this_year":"$27,935.24"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestInsightsControllerTestReadMrr tests the behavior of the InsightsController's
func TestInsightsControllerTestReadMrr(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := insightsController.ReadMrr(ctx, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"mrr":{"amount_in_cents":9915593,"amount_formatted":"$99,155.93","currency":"USD","currency_symbol":"$","at_time":"2021-02-03T14:23:17-05:00","breakouts":{"plan_amount_in_cents":9913593,"plan_amount_formatted":"$99,135.93","usage_amount_in_cents":2000,"usage_amount_formatted":"$20.00"}}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestInsightsControllerTestReadMrrMovements tests the behavior of the InsightsController's
func TestInsightsControllerTestReadMrrMovements(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 10
	apiResponse, err := insightsController.ReadMrrMovements(ctx, nil, &page, &perPage, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"mrr":{"page":0,"per_page":10,"total_pages":80,"total_entries":791,"currency":"USD","currency_symbol":"$","movements":[{"timestamp":"2014-12-03T13:59:46-05:00","amount_in_cents":2173,"amount_formatted":"$21.73","description":"Awesome Company signed up for Super Product ($21.73/mo)","category":"new_business","breakouts":{"plan_amount_in_cents":2173,"plan_amount_formatted":"$21.73","usage_amount_in_cents":0,"usage_amount_formatted":"$0.00"},"line_items":[{"product_id":306386,"component_id":0,"price_point_id":3856987,"name":"Cached Queries","mrr":2173,"mrr_movements":[{"amount":2173,"category":"new_business","subscriber_delta":0,"lead_delta":0}],"quantity":1,"prev_quantity":0,"recurring":true}],"subscription_id":12355,"subscriber_name":"Amy Smith"}]}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestInsightsControllerTestListMrrPerSubscription tests the behavior of the InsightsController's
func TestInsightsControllerTestListMrrPerSubscription(t *testing.T) {
	ctx := context.Background()
	filterSubscriptionIds := []int{1, 2, 3}
	atTime := "at_time=2022-01-10T10:00:00-05:00"
	page := 1
	perPage := 20
	direction := models.DirectionEnum("desc")
	apiResponse, err := insightsController.ListMrrPerSubscription(ctx, filterSubscriptionIds, &atTime, &page, &perPage, &direction)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
