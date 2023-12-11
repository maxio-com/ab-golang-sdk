package ab_golang_sdk

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"maxioadvancedbilling/models"
	"testing"
)

// TestReasonCodesControllerTestCreateReasonCode tests the behavior of the ReasonCodesController's
func TestReasonCodesControllerTestCreateReasonCode(t *testing.T) {
	ctx := context.Background()
	var body models.CreateReasonCodeRequest
	err := json.Unmarshal([]byte(`{"reason_code":{"code":"NOTHANKYOU","description":"No thank you!","position":5}}`), &body)
	if err != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := reasonCodesController.CreateReasonCode(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestReasonCodesControllerTestListReasonCodes tests the behavior of the ReasonCodesController's
func TestReasonCodesControllerTestListReasonCodes(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	apiResponse, err := reasonCodesController.ListReasonCodes(ctx, &page, &perPage)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"reason_code":{"id":2,"site_id":2,"code":"LARGE","description":"This is too complicated","position":1,"created_at":"2017-02-16T16:49:07-05:00","updated_at":"2017-02-17T16:29:51-05:00"}},{"reason_code":{"id":1,"site_id":2,"code":"CH1","description":"This doesnt meet my needs","position":2,"created_at":"2017-02-16T16:48:45-05:00","updated_at":"2017-02-17T16:29:59-05:00"}},{"reason_code":{"id":5,"site_id":2,"code":"HAN99","description":"Hard to setup","position":3,"created_at":"2017-02-17T16:29:42-05:00","updated_at":"2017-02-17T16:29:59-05:00"}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}
