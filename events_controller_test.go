package ab_golang_sdk

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"maxioadvancedbilling/models"
	"testing"
)

// TestEventsControllerTestListEvents tests the behavior of the EventsController's
func TestEventsControllerTestListEvents(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	direction := models.DirectionEnum("desc")
	filter := []models.EventTypeEnum{"custom_field_value_change", "payment_success"}
	dateField := models.ListEventsDateFieldEnum("created_at")
	apiResponse, err := eventsController.ListEvents(ctx, &page, &perPage, nil, nil, &direction, filter, &dateField, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"event":{"id":343087780,"key":"subscription_state_change","message":"State changed on Test subscription to Monthly Product from active to past_due","subscription_id":14950962,"customer_id":12345678,"created_at":"2016-10-27T16:42:22-04:00","event_specific_data":{"previous_subscription_state":"active","new_subscription_state":"past_due"}}},{"event":{"id":343087742,"key":"billing_date_change","message":"Billing date changed on Test's subscription to Monthly Product from 11/27/2016 to 10/27/2016","subscription_id":14950962,"customer_id":12345678,"created_at":"2016-10-27T16:42:19-04:00","event_specific_data":null}},{"event":{"id":343085267,"key":"statement_closed","message":"Statement 79401838 closed (but not settled) for Test's subscription to ANNUAL product","subscription_id":14950975,"customer_id":87654321,"created_at":"2016-10-27T16:40:40-04:00","event_specific_data":null}},{"event":{"id":4481,"key":"custom_field_value_change","message":"Custom field (Extra support included) changed for Subscription 117 from 'Yes' to 'No'.","subscription_id":117,"customer_id":22334455,"created_at":"2022-03-24T07:55:06-04:00","event_specific_data":{"event_type":"updated","metafield_name":"Extra support included","metafield_id":2,"old_value":"Yes","new_value":"No","resource_type":"Subscription","resource_id":117}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestEventsControllerTestListEvents1 tests the behavior of the EventsController's
func TestEventsControllerTestListEvents1(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	direction := models.DirectionEnum("desc")
	filter := []models.EventTypeEnum{"custom_field_value_change", "payment_success"}
	dateField := models.ListEventsDateFieldEnum("created_at")
	apiResponse, err := eventsController.ListEvents(ctx, &page, &perPage, nil, nil, &direction, filter, &dateField, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/xml"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestEventsControllerTestListEvents2 tests the behavior of the EventsController's
func TestEventsControllerTestListEvents2(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	direction := models.DirectionEnum("desc")
	filter := []models.EventTypeEnum{"custom_field_value_change", "payment_success"}
	dateField := models.ListEventsDateFieldEnum("created_at")
	apiResponse, err := eventsController.ListEvents(ctx, &page, &perPage, nil, nil, &direction, filter, &dateField, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "multipart/form-data"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestEventsControllerTestReadEventsCount tests the behavior of the EventsController's
func TestEventsControllerTestReadEventsCount(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	direction := models.DirectionEnum("desc")
	filter := []models.EventTypeEnum{"custom_field_value_change", "payment_success"}
	apiResponse, err := eventsController.ReadEventsCount(ctx, &page, &perPage, nil, nil, &direction, filter)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"count":144}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}
