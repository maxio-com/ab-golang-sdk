package advancedbilling

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}
	var result models.ReasonCodeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReasonCodeResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListReasonCodesInput represents the input of the ListReasonCodes endpoint.
type ListReasonCodesInput struct {
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
}

// ListReasonCodes takes context, page, perPage as parameters and
// returns an models.ApiResponse with []models.ReasonCodeResponse data and
// an error if there was an issue with the request or response.
// This method gives a merchant the option to retrieve a list of all of the current churn codes for a given site.
func (r *ReasonCodesController) ListReasonCodes(
	ctx context.Context,
	input ListReasonCodesInput) (
	models.ApiResponse[[]models.ReasonCodeResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", "/reason_codes.json")
	req.Authenticate(NewAuth("BasicAuth"))
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	var result []models.ReasonCodeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ReasonCodeResponse](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})

	var result models.ReasonCodeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReasonCodeResponse](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ReasonCodeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReasonCodeResponse](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})

	var result models.ReasonCodesJsonResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ReasonCodesJsonResponse](decoder)
	return models.NewApiResponse(result, resp), err
}
