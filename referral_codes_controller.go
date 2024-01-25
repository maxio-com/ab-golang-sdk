package advancedbilling

import (
	"context"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
)

// ReferralCodesController represents a controller struct.
type ReferralCodesController struct {
	baseController
}

// NewReferralCodesController creates a new instance of ReferralCodesController.
// It takes a baseController as a parameter and returns a pointer to the ReferralCodesController.
func NewReferralCodesController(baseController baseController) *ReferralCodesController {
	referralCodesController := ReferralCodesController{baseController: baseController}
	return &referralCodesController
}

// ValidateReferralCode takes context, code as parameters and
// returns an models.ApiResponse with models.ReferralValidationResponse data and
// an error if there was an issue with the request or response.
// Use this method to determine if the referral code is valid and applicable within your Site. This method is useful for validating referral codes that are entered by a customer.
// ## Referrals Documentation
// Full documentation on how to use the referrals feature in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407802831643).
// ## Server Response
// If the referral code is valid the status code will be `200` and the referral code will be returned. If the referral code is invalid, a `404` response will be returned.
func (r *ReferralCodesController) ValidateReferralCode(
	ctx context.Context,
	code string) (
	models.ApiResponse[models.ReferralValidationResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", "/referral_codes/validate.json")
	req.Authenticate(true)
	req.QueryParam("code", code)
	var result models.ReferralValidationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReferralValidationResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewSingleStringErrorResponse(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}
