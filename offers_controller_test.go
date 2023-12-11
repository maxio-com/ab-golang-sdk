package ab_golang_sdk

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"maxioadvancedbilling/models"
	"testing"
)

// TestOffersControllerTestCreateOffer tests the behavior of the OffersController's
func TestOffersControllerTestCreateOffer(t *testing.T) {
	ctx := context.Background()
	var body models.CreateOfferRequest
	err := json.Unmarshal([]byte(`{"offer":{"name":"Solo","handle":"han_shot_first","description":"A Star Wars Story","product_id":31,"product_price_point_id":102,"components":[{"component_id":24,"starting_quantity":1}],"coupons":["DEF456"]}}`), &body)
	if err != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := offersController.CreateOffer(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"offer":{"id":3,"site_id":2,"product_family_id":4,"product_family_name":"Chargify","product_id":31,"product_name":"30-Day Square Trial","product_price_in_cents":2000,"product_revisable_number":0,"name":"Solo","handle":"han_shot_first","description":"A Star Wars Story","created_at":"2018-06-08T14:51:52-04:00","updated_at":"2018-06-08T14:51:52-04:00","archived_at":null,"product_price_point_name":"Default","offer_items":[{"component_id":24,"component_name":"Invoices","component_unit_price":"3.0","price_point_id":104,"price_point_name":"Original","starting_quantity":"1.0","editable":false}],"offer_discounts":[{"coupon_id":3,"coupon_code":"DEF456","coupon_name":"IB Loyalty"}]}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}

// TestOffersControllerTestListOffers tests the behavior of the OffersController's
func TestOffersControllerTestListOffers(t *testing.T) {
	ctx := context.Background()
	page := 1
	perPage := 20
	includeArchived := true
	apiResponse, err := offersController.ListOffers(ctx, &page, &perPage, &includeArchived)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"offers":[{"id":239,"site_id":48110,"product_family_id":1025627,"product_family_name":"Gold","product_id":110,"product_name":"Pro","product_price_in_cents":1000,"product_revisable_number":0,"product_price_point_id":138,"product_price_point_name":"Default","name":"Third Offer","handle":"third","description":"","created_at":"2018-08-03T09:56:11-05:00","updated_at":"2018-08-03T09:56:11-05:00","archived_at":null,"offer_items":[{"component_id":426665,"component_name":"Database Size (GB)","component_unit_price":"1.0","price_point_id":149438,"price_point_name":"Auto-created","starting_quantity":"0.0","editable":false,"currency_prices":[]}],"offer_discounts":[{"coupon_id":234,"coupon_code":"GR8_CUSTOMER","coupon_name":"Multi-service Discount"}],"offer_signup_pages":[{"id":356482,"nickname":"ggoods","enabled":true,"return_url":"","return_params":"","url":"https://general-goods.chargifypay.com/subscribe/hjpvhnw63tzy"}]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Data, false, false)
}
