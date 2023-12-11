package ab_golang_sdk

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"maxioadvancedbilling/models"
	"testing"
)

// TestCustomersControllerTestCreateCustomer tests the behavior of the CustomersController's
func TestCustomersControllerTestCreateCustomer(t *testing.T) {
	ctx := context.Background()
	var body models.CreateCustomerRequest
	err := json.Unmarshal([]byte(`{"customer":{"first_name":"Martha","last_name":"Washington","email":"martha@example.com","cc_emails":"george@example.com","organization":"ABC, Inc.","reference":"1234567890","address":"123 Main Street","address_2":"Unit 10","city":"Anytown","state":"MA","zip":"02120","country":"US","phone":"555-555-1212","locale":"es-MX"}}`), &body)
	if err != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := customersController.CreateCustomer(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"customer":{"first_name":"Cathryn","last_name":"Parisian","email":"Stella.McLaughlin6@example.net","cc_emails":null,"organization":"Greenholt - Oberbrunner","reference":null,"id":76,"created_at":"2021-03-29T07:47:00-04:00","updated_at":"2021-03-29T07:47:00-04:00","address":"739 Stephon Bypass","address_2":"Apt. 386","city":"Sedrickchester","state":"KY","state_name":"Kentucky","zip":"46979-7719","country":"US","country_name":"United States","phone":"230-934-3685","verified":false,"portal_customer_created_at":null,"portal_invite_last_sent_at":null,"portal_invite_last_accepted_at":null,"tax_exempt":false,"vat_number":null,"parent_id":null,"locale":"en-US"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestCustomersControllerTestListCustomers tests the behavior of the CustomersController's
func TestCustomersControllerTestListCustomers(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 50
	dateField := models.BasicDateFieldEnum("updated_at")
	apiResponse, err := customersController.ListCustomers(ctx, nil, &page, &perPage, &dateField, nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"customer":{"first_name":"Kayla","last_name":"Test","email":"kayla@example.com","cc_emails":"john@example.com, sue@example.com","organization":"","reference":null,"id":14126091,"created_at":"2016-10-04T15:22:27-04:00","updated_at":"2016-10-04T15:22:30-04:00","address":"","address_2":"","city":"","state":"","zip":"","country":"","phone":"","verified":null,"portal_customer_created_at":"2016-10-04T15:22:29-04:00","portal_invite_last_sent_at":"2016-10-04T15:22:30-04:00","portal_invite_last_accepted_at":null,"tax_exempt":false}},{"customer":{"first_name":"Nick ","last_name":"Test","email":"nick@example.com","cc_emails":"john@example.com, sue@example.com","organization":"","reference":null,"id":14254093,"created_at":"2016-10-13T16:52:51-04:00","updated_at":"2016-10-13T16:52:54-04:00","address":"","address_2":"","city":"","state":"","zip":"","country":"","phone":"","verified":null,"portal_customer_created_at":"2016-10-13T16:52:54-04:00","portal_invite_last_sent_at":"2016-10-13T16:52:54-04:00","portal_invite_last_accepted_at":null,"tax_exempt":false,"parent_id":123}},{"customer":{"first_name":"Don","last_name":"Test","email":"don@example.com","cc_emails":"john@example.com, sue@example.com","organization":"","reference":null,"id":14332342,"created_at":"2016-10-19T10:49:13-04:00","updated_at":"2016-10-19T10:49:19-04:00","address":"1737 15th St","address_2":"","city":"Boulder","state":"CO","zip":"80302","country":"US","phone":"","verified":null,"portal_customer_created_at":"2016-10-19T10:49:19-04:00","portal_invite_last_sent_at":"2016-10-19T10:49:19-04:00","portal_invite_last_accepted_at":null,"tax_exempt":false,"parent_id":null}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}
