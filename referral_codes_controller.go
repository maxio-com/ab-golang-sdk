/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
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
// Full documentation on how to use the referrals feature in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/sections/24286965611405-Referrals).
// ## Server Response
// If the referral code is valid the status code will be `200` and the referral code will be returned. If the referral code is invalid, a `404` response will be returned.
func (r *ReferralCodesController) ValidateReferralCode(
    ctx context.Context,
    code string) (
    models.ApiResponse[models.ReferralValidationResponse],
    error) {
    req := r.prepareRequest(ctx, "GET", "/referral_codes/validate.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Invalid referral code.", Unmarshaller: errors.NewSingleStringErrorResponse},
    })
    req.QueryParam("code", code)
    var result models.ReferralValidationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ReferralValidationResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
