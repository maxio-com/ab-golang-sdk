// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
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
// Validates whether a referral code is valid and applicable within your site. This method is useful for validating referral codes that are entered by a customer.
// For more information, see [Understanding Referrals](https://docs.maxio.com/hc/en-us/articles/24286981223693-Understanding-Referrals) in the product documentation.
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
