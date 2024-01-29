package advancedbilling

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
)

// ReasonCodesController represents a controller struct.
type ReasonCodesController struct {
	baseController
}

// NewReasonCodesController creates a new instance of ReasonCodesController.
// It takes a baseController as a parameter and returns a pointer to the ReasonCodesController.
func NewReasonCodesController(baseController baseController) *ReasonCodesController {
	reasonCodesController := ReasonCodesController{baseController: baseController}
	return &reasonCodesController
}

// CreateReasonCode takes context, body as parameters and
// returns an models.ApiResponse with models.ReasonCodeResponse data and
// an error if there was an issue with the request or response.
// # Reason Codes Intro
// ReasonCodes are a way to gain a high level view of why your customers are cancelling the subcription to your product or service.
// Add a set of churn reason codes to be displayed in-app and/or the Chargify Billing Portal. As your subscribers decide to cancel their subscription, learn why they decided to cancel.
// ## Reason Code Documentation
// Full documentation on how Reason Codes operate within Chargify can be located under the following links.
// [Churn Reason Codes](https://chargify.zendesk.com/hc/en-us/articles/4407896775579#churn-reason-codes)
// ## Create Reason Code
// This method gives a merchant the option to create a reason codes for a given Site.
func (r *ReasonCodesController) CreateReasonCode(
	ctx context.Context,
	body *models.CreateReasonCodeRequest) (
	models.ApiResponse[models.ReasonCodeResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/reason_codes.json")
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}
	var result models.ReasonCodeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReasonCodeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 422 {
		err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
	}
	return models.NewApiResponse(result, resp), err
}

// ListReasonCodes takes context, page, perPage as parameters and
// returns an models.ApiResponse with []models.ReasonCodeResponse data and
// an error if there was an issue with the request or response.
// This method gives a merchant the option to retrieve a list of all of the current churn codes for a given site.
func (r *ReasonCodesController) ListReasonCodes(
	ctx context.Context,
	page *int,
	perPage *int) (
	models.ApiResponse[[]models.ReasonCodeResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", "/reason_codes.json")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	var result []models.ReasonCodeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ReasonCodeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// ReadReasonCode takes context, reasonCodeId as parameters and
// returns an models.ApiResponse with models.ReasonCodeResponse data and
// an error if there was an issue with the request or response.
// This method gives a merchant the option to retrieve a list of a particular code for a given Site by providing the unique numerical ID of the code.
func (r *ReasonCodesController) ReadReasonCode(
	ctx context.Context,
	reasonCodeId int) (
	models.ApiResponse[models.ReasonCodeResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/reason_codes/%v.json", reasonCodeId),
	)
	req.Authenticate(true)

	var result models.ReasonCodeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReasonCodeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}

// UpdateReasonCode takes context, reasonCodeId, body as parameters and
// returns an models.ApiResponse with models.ReasonCodeResponse data and
// an error if there was an issue with the request or response.
// This method gives a merchant the option to update an existing reason code for a given site.
func (r *ReasonCodesController) UpdateReasonCode(
	ctx context.Context,
	reasonCodeId int,
	body *models.UpdateReasonCodeRequest) (
	models.ApiResponse[models.ReasonCodeResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/reason_codes/%v.json", reasonCodeId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result models.ReasonCodeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReasonCodeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}

// DeleteReasonCode takes context, reasonCodeId as parameters and
// returns an models.ApiResponse with models.ReasonCodesJsonResponse data and
// an error if there was an issue with the request or response.
// This method gives a merchant the option to delete one reason code from the Churn Reason Codes. This code will be immediately removed. This action is not reversable.
func (r *ReasonCodesController) DeleteReasonCode(
	ctx context.Context,
	reasonCodeId int) (
	models.ApiResponse[models.ReasonCodesJsonResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/reason_codes/%v.json", reasonCodeId),
	)
	req.Authenticate(true)

	var result models.ReasonCodesJsonResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReasonCodesJsonResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}
