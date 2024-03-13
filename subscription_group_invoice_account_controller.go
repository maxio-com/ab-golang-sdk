package advancedbilling

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
	"time"
)

// SubscriptionGroupInvoiceAccountController represents a controller struct.
type SubscriptionGroupInvoiceAccountController struct {
	baseController
}

// NewSubscriptionGroupInvoiceAccountController creates a new instance of SubscriptionGroupInvoiceAccountController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionGroupInvoiceAccountController.
func NewSubscriptionGroupInvoiceAccountController(baseController baseController) *SubscriptionGroupInvoiceAccountController {
	subscriptionGroupInvoiceAccountController := SubscriptionGroupInvoiceAccountController{baseController: baseController}
	return &subscriptionGroupInvoiceAccountController
}

// CreateSubscriptionGroupPrepayment takes context, uid, body as parameters and
// returns an models.ApiResponse with models.SubscriptionGroupPrepaymentResponse data and
// an error if there was an issue with the request or response.
// A prepayment can be added for a subscription group identified by the group's `uid`. This endpoint requires a `amount`, `details`, `method`, and `memo`. On success, the prepayment will be added to the group's prepayment balance.
func (s *SubscriptionGroupInvoiceAccountController) CreateSubscriptionGroupPrepayment(
	ctx context.Context,
	uid string,
	body *models.SubscriptionGroupPrepaymentRequest) (
	models.ApiResponse[models.SubscriptionGroupPrepaymentResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscription_groups/%v/prepayments.json", uid),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.SubscriptionGroupPrepaymentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SubscriptionGroupPrepaymentResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListPrepaymentsForSubscriptionGroupInput represents the input of the ListPrepaymentsForSubscriptionGroup endpoint.
type ListPrepaymentsForSubscriptionGroupInput struct {
	// The uid of the subscription group
	Uid string
	// The type of filter you would like to apply to your search.
	// Use in query: `filter[date_field]=created_at`.
	FilterDateField *models.ListSubscriptionGroupPrepaymentDateField
	// The end date (format YYYY-MM-DD) with which to filter the date_field.
	// Returns prepayments with a timestamp up to and including 11:59:59PM in your site's time zone on the date specified.
	// Use in query: `filter[end_date]=2011-12-15`.
	FilterEndDate *time.Time
	// The start date (format YYYY-MM-DD) with which to filter the date_field.
	// Returns prepayments with a timestamp at or after midnight (12:00:00 AM) in your site's time zone on the date specified.
	// Use in query: `filter[start_date]=2011-12-15`.
	FilterStartDate *time.Time
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
}

// ListPrepaymentsForSubscriptionGroup takes context, uid, filterDateField, filterEndDate, filterStartDate, page, perPage as parameters and
// returns an models.ApiResponse with models.ListSubscriptionGroupPrepaymentResponse data and
// an error if there was an issue with the request or response.
// This request will list a subscription group's prepayments.
func (s *SubscriptionGroupInvoiceAccountController) ListPrepaymentsForSubscriptionGroup(
	ctx context.Context,
	input ListPrepaymentsForSubscriptionGroupInput) (
	models.ApiResponse[models.ListSubscriptionGroupPrepaymentResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscription_groups/%v/prepayments.json", input.Uid),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})
	if input.FilterDateField != nil {
		req.QueryParam("filter[date_field]", *input.FilterDateField)
	}
	if input.FilterEndDate != nil {
		req.QueryParam("filter[end_date]", input.FilterEndDate.Format(models.DEFAULT_DATE))
	}
	if input.FilterStartDate != nil {
		req.QueryParam("filter[start_date]", input.FilterStartDate.Format(models.DEFAULT_DATE))
	}
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}

	var result models.ListSubscriptionGroupPrepaymentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListSubscriptionGroupPrepaymentResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// IssueSubscriptionGroupServiceCredit takes context, uid, body as parameters and
// returns an models.ApiResponse with models.ServiceCreditResponse data and
// an error if there was an issue with the request or response.
// Credit can be issued for a subscription group identified by the group's `uid`. Credit will be added to the group in the amount specified in the request body. The credit will be applied to group member invoices as they are generated.
func (s *SubscriptionGroupInvoiceAccountController) IssueSubscriptionGroupServiceCredit(
	ctx context.Context,
	uid string,
	body *models.IssueServiceCreditRequest) (
	models.ApiResponse[models.ServiceCreditResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscription_groups/%v/service_credits.json", uid),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ServiceCreditResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ServiceCreditResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeductSubscriptionGroupServiceCredit takes context, uid, body as parameters and
// returns an models.ApiResponse with models.ServiceCredit data and
// an error if there was an issue with the request or response.
// Credit can be deducted for a subscription group identified by the group's `uid`. Credit will be deducted from the group in the amount specified in the request body.
func (s *SubscriptionGroupInvoiceAccountController) DeductSubscriptionGroupServiceCredit(
	ctx context.Context,
	uid string,
	body *models.DeductServiceCreditRequest) (
	models.ApiResponse[models.ServiceCredit],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscription_groups/%v/service_credit_deductions.json", uid),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ServiceCredit
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ServiceCredit](decoder)
	return models.NewApiResponse(result, resp), err
}
