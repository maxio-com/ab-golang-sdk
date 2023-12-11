package ab_golang_sdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"maxioadvancedbilling/errors"
	"maxioadvancedbilling/models"
	"net/http"
	"time"
)

// SubscriptionInvoiceAccountController represents a controller struct.
type SubscriptionInvoiceAccountController struct {
	baseController
}

// NewSubscriptionInvoiceAccountController creates a new instance of SubscriptionInvoiceAccountController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionInvoiceAccountController.
func NewSubscriptionInvoiceAccountController(baseController baseController) *SubscriptionInvoiceAccountController {
	subscriptionInvoiceAccountController := SubscriptionInvoiceAccountController{baseController: baseController}
	return &subscriptionInvoiceAccountController
}

// ReadAccountBalances takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.AccountBalances data and
// an error if there was an issue with the request or response.
// Returns the `balance_in_cents` of the Subscription's Pending Discount, Service Credit, and Prepayment accounts, as well as the sum of the Subscription's open, payable invoices.
func (s *SubscriptionInvoiceAccountController) ReadAccountBalances(
	ctx context.Context,
	subscriptionId int) (
	models.ApiResponse[models.AccountBalances],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/account_balances.json", subscriptionId),
	)
	req.Authenticate(true)

	var result models.AccountBalances
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountBalances](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreatePrepayment takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.CreatePrepaymentResponse data and
// an error if there was an issue with the request or response.
// ## Create Prepayment
// In order to specify a prepayment made against a subscription, specify the `amount, memo, details, method`.
// When the `method` specified is `"credit_card_on_file"`, the prepayment amount will be collected using the default credit card payment profile and applied to the prepayment account balance.  This is especially useful for manual replenishment of prepaid subscriptions.
// Please note that you **can't** pass `amount_in_cents`.
func (s *SubscriptionInvoiceAccountController) CreatePrepayment(
	ctx context.Context,
	subscriptionId int,
	body *models.CreatePrepaymentRequest) (
	models.ApiResponse[models.CreatePrepaymentResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/prepayments.json", subscriptionId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result models.CreatePrepaymentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.CreatePrepaymentResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// ListPrepayments takes context, subscriptionId, page, perPage, filterDateField, filterStartDate, filterEndDate as parameters and
// returns an models.ApiResponse with models.PrepaymentsResponse data and
// an error if there was an issue with the request or response.
// This request will list a subscription's prepayments.
func (s *SubscriptionInvoiceAccountController) ListPrepayments(
	ctx context.Context,
	subscriptionId int,
	page *int,
	perPage *int,
	filterDateField *models.BasicDateFieldEnum,
	filterStartDate *time.Time,
	filterEndDate *time.Time) (
	models.ApiResponse[models.PrepaymentsResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/prepayments.json", subscriptionId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if filterDateField != nil {
		req.QueryParam("filter[date_field]", *filterDateField)
	}
	if filterStartDate != nil {
		req.QueryParam("filter[start_date]", filterStartDate.Format(models.DEFAULT_DATE))
	}
	if filterEndDate != nil {
		req.QueryParam("filter[end_date]", filterEndDate.Format(models.DEFAULT_DATE))
	}

	var result models.PrepaymentsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.PrepaymentsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 401 {
		err = errors.NewApiError(401, "Unauthorized")
	}
	if resp.StatusCode == 403 {
		err = errors.NewApiError(403, "Forbidden")
	}
	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}

// IssueServiceCredit takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.ServiceCredit data and
// an error if there was an issue with the request or response.
// Credit will be added to the subscription in the amount specified in the request body. The credit is subsequently applied to the next generated invoice.
func (s *SubscriptionInvoiceAccountController) IssueServiceCredit(
	ctx context.Context,
	subscriptionId int,
	body *models.IssueServiceCreditRequest) (
	models.ApiResponse[models.ServiceCredit],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/service_credits.json", subscriptionId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result models.ServiceCredit
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ServiceCredit](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeductServiceCredit takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Credit will be removed from the subscription in the amount specified in the request body. The credit amount being deducted must be equal to or less than the current credit balance.
func (s *SubscriptionInvoiceAccountController) DeductServiceCredit(
	ctx context.Context,
	subscriptionId int,
	body *models.DeductServiceCreditRequest) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/service_credit_deductions.json", subscriptionId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	if context.Response.StatusCode == 422 {
		err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
	}
	return context.Response, err
}

// RefundPrepayment takes context, subscriptionId, prepaymentId, body as parameters and
// returns an models.ApiResponse with models.PrepaymentResponse data and
// an error if there was an issue with the request or response.
// This endpoint will refund, completely or partially, a particular prepayment applied to a subscription. The `prepayment_id` will be the account transaction ID of the original payment. The prepayment must have some amount remaining in order to be refunded.
// The amount may be passed either as a decimal, with `amount`, or an integer in cents, with `amount_in_cents`.
func (s *SubscriptionInvoiceAccountController) RefundPrepayment(
	ctx context.Context,
	subscriptionId int,
	prepaymentId string,
	body *models.RefundPrepaymentRequest) (
	models.ApiResponse[models.PrepaymentResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/prepayments/%v/refunds.json", subscriptionId, prepaymentId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result models.PrepaymentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.PrepaymentResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 400 {
		err = errors.NewRefundPrepaymentBaseErrorsResponse(400, "Bad Request")
	}
	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	if resp.StatusCode == 422 {
		err = errors.NewRefundPrepaymentAggregatedErrorsResponse(422, "Unprocessable Entity")
	}
	return models.NewApiResponse(result, resp), err
}
