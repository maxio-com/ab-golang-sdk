/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package errors

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/apiError"
	"github.com/maxio-com/ab-golang-sdk/models"
)

var NewApiError = apiError.NewApiError

// ErrorListResponse is a custom error.
// Error which contains list of messages.
type ErrorListResponse struct {
	apiError.ApiError
	Errors []string `json:"errors"`
}

// NewErrorListResponse is a constructor for ErrorListResponse.
// It creates and returns a pointer to a new ErrorListResponse instance with the given statusCode and body.
func NewErrorListResponse(
	statusCode int,
	body string) *ErrorListResponse {
	return &ErrorListResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ErrorListResponse.
func (e *ErrorListResponse) Error() string {
	return fmt.Sprintf("ErrorListResponse occured %v", e.Body)
}

// CustomerErrorResponse is a custom error.
type CustomerErrorResponse struct {
	apiError.ApiError
	Errors *interface{} `json:"errors,omitempty"`
}

// NewCustomerErrorResponse is a constructor for CustomerErrorResponse.
// It creates and returns a pointer to a new CustomerErrorResponse instance with the given statusCode and body.
func NewCustomerErrorResponse(
	statusCode int,
	body string) *CustomerErrorResponse {
	return &CustomerErrorResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for CustomerErrorResponse.
func (c *CustomerErrorResponse) Error() string {
	return fmt.Sprintf("CustomerErrorResponse occured %v", c.Body)
}

// TooManyManagementLinkRequestsError is a custom error.
type TooManyManagementLinkRequestsError struct {
	apiError.ApiError
	Errors models.TooManyManagementLinkRequests `json:"errors"`
}

// NewTooManyManagementLinkRequestsError is a constructor for TooManyManagementLinkRequestsError.
// It creates and returns a pointer to a new TooManyManagementLinkRequestsError instance with the given statusCode and body.
func NewTooManyManagementLinkRequestsError(
	statusCode int,
	body string) *TooManyManagementLinkRequestsError {
	return &TooManyManagementLinkRequestsError{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for TooManyManagementLinkRequestsError.
func (t *TooManyManagementLinkRequestsError) Error() string {
	return fmt.Sprintf("TooManyManagementLinkRequestsError occured %v", t.Body)
}

// SingleStringErrorResponse is a custom error.
type SingleStringErrorResponse struct {
	apiError.ApiError
	Errors *string `json:"errors,omitempty"`
}

// NewSingleStringErrorResponse is a constructor for SingleStringErrorResponse.
// It creates and returns a pointer to a new SingleStringErrorResponse instance with the given statusCode and body.
func NewSingleStringErrorResponse(
	statusCode int,
	body string) *SingleStringErrorResponse {
	return &SingleStringErrorResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SingleStringErrorResponse.
func (s *SingleStringErrorResponse) Error() string {
	return fmt.Sprintf("SingleStringErrorResponse occured %v", s.Body)
}

// SingleErrorResponse is a custom error.
type SingleErrorResponse struct {
	apiError.ApiError
	MError string `json:"error"`
}

// NewSingleErrorResponse is a constructor for SingleErrorResponse.
// It creates and returns a pointer to a new SingleErrorResponse instance with the given statusCode and body.
func NewSingleErrorResponse(
	statusCode int,
	body string) *SingleErrorResponse {
	return &SingleErrorResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SingleErrorResponse.
func (s *SingleErrorResponse) Error() string {
	return fmt.Sprintf("SingleErrorResponse occured %v", s.Body)
}

// ErrorArrayMapResponse is a custom error.
type ErrorArrayMapResponse struct {
	apiError.ApiError
	Errors map[string]interface{} `json:"errors,omitempty"`
}

// NewErrorArrayMapResponse is a constructor for ErrorArrayMapResponse.
// It creates and returns a pointer to a new ErrorArrayMapResponse instance with the given statusCode and body.
func NewErrorArrayMapResponse(
	statusCode int,
	body string) *ErrorArrayMapResponse {
	return &ErrorArrayMapResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ErrorArrayMapResponse.
func (e *ErrorArrayMapResponse) Error() string {
	return fmt.Sprintf("ErrorArrayMapResponse occured %v", e.Body)
}

// ProductPricePointErrorResponse is a custom error.
type ProductPricePointErrorResponse struct {
	apiError.ApiError
	Errors models.ProductPricePointErrors `json:"errors"`
}

// NewProductPricePointErrorResponse is a constructor for ProductPricePointErrorResponse.
// It creates and returns a pointer to a new ProductPricePointErrorResponse instance with the given statusCode and body.
func NewProductPricePointErrorResponse(
	statusCode int,
	body string) *ProductPricePointErrorResponse {
	return &ProductPricePointErrorResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ProductPricePointErrorResponse.
func (p *ProductPricePointErrorResponse) Error() string {
	return fmt.Sprintf("ProductPricePointErrorResponse occured %v", p.Body)
}

// ErrorStringMapResponse is a custom error.
type ErrorStringMapResponse struct {
	apiError.ApiError
	Errors map[string]string `json:"errors,omitempty"`
}

// NewErrorStringMapResponse is a constructor for ErrorStringMapResponse.
// It creates and returns a pointer to a new ErrorStringMapResponse instance with the given statusCode and body.
func NewErrorStringMapResponse(
	statusCode int,
	body string) *ErrorStringMapResponse {
	return &ErrorStringMapResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ErrorStringMapResponse.
func (e *ErrorStringMapResponse) Error() string {
	return fmt.Sprintf("ErrorStringMapResponse occured %v", e.Body)
}

// ComponentPricePointError is a custom error.
type ComponentPricePointError struct {
	apiError.ApiError
	Errors []models.ComponentPricePointErrorItem `json:"errors,omitempty"`
}

// NewComponentPricePointError is a constructor for ComponentPricePointError.
// It creates and returns a pointer to a new ComponentPricePointError instance with the given statusCode and body.
func NewComponentPricePointError(
	statusCode int,
	body string) *ComponentPricePointError {
	return &ComponentPricePointError{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ComponentPricePointError.
func (c *ComponentPricePointError) Error() string {
	return fmt.Sprintf("ComponentPricePointError occured %v", c.Body)
}

// ComponentAllocationError is a custom error.
type ComponentAllocationError struct {
	apiError.ApiError
	Errors []models.ComponentAllocationErrorItem `json:"errors,omitempty"`
}

// NewComponentAllocationError is a constructor for ComponentAllocationError.
// It creates and returns a pointer to a new ComponentAllocationError instance with the given statusCode and body.
func NewComponentAllocationError(
	statusCode int,
	body string) *ComponentAllocationError {
	return &ComponentAllocationError{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ComponentAllocationError.
func (c *ComponentAllocationError) Error() string {
	return fmt.Sprintf("ComponentAllocationError occured %v", c.Body)
}

// SubscriptionComponentAllocationError is a custom error.
type SubscriptionComponentAllocationError struct {
	apiError.ApiError
	Errors []models.SubscriptionComponentAllocationErrorItem `json:"errors,omitempty"`
}

// NewSubscriptionComponentAllocationError is a constructor for SubscriptionComponentAllocationError.
// It creates and returns a pointer to a new SubscriptionComponentAllocationError instance with the given statusCode and body.
func NewSubscriptionComponentAllocationError(
	statusCode int,
	body string) *SubscriptionComponentAllocationError {
	return &SubscriptionComponentAllocationError{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionComponentAllocationError.
func (s *SubscriptionComponentAllocationError) Error() string {
	return fmt.Sprintf("SubscriptionComponentAllocationError occured %v", s.Body)
}

// SubscriptionGroupSignupErrorResponse is a custom error.
type SubscriptionGroupSignupErrorResponse struct {
	apiError.ApiError
	Errors models.SubscriptionGroupSignupError `json:"errors"`
}

// NewSubscriptionGroupSignupErrorResponse is a constructor for SubscriptionGroupSignupErrorResponse.
// It creates and returns a pointer to a new SubscriptionGroupSignupErrorResponse instance with the given statusCode and body.
func NewSubscriptionGroupSignupErrorResponse(
	statusCode int,
	body string) *SubscriptionGroupSignupErrorResponse {
	return &SubscriptionGroupSignupErrorResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionGroupSignupErrorResponse.
func (s *SubscriptionGroupSignupErrorResponse) Error() string {
	return fmt.Sprintf("SubscriptionGroupSignupErrorResponse occured %v", s.Body)
}

// SubscriptionGroupUpdateErrorResponse is a custom error.
type SubscriptionGroupUpdateErrorResponse struct {
	apiError.ApiError
	Errors *models.SubscriptionGroupUpdateError `json:"errors,omitempty"`
}

// NewSubscriptionGroupUpdateErrorResponse is a constructor for SubscriptionGroupUpdateErrorResponse.
// It creates and returns a pointer to a new SubscriptionGroupUpdateErrorResponse instance with the given statusCode and body.
func NewSubscriptionGroupUpdateErrorResponse(
	statusCode int,
	body string) *SubscriptionGroupUpdateErrorResponse {
	return &SubscriptionGroupUpdateErrorResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionGroupUpdateErrorResponse.
func (s *SubscriptionGroupUpdateErrorResponse) Error() string {
	return fmt.Sprintf("SubscriptionGroupUpdateErrorResponse occured %v", s.Body)
}

// SubscriptionAddCouponError is a custom error.
type SubscriptionAddCouponError struct {
	apiError.ApiError
	Codes        []string `json:"codes,omitempty"`
	CouponCode   []string `json:"coupon_code,omitempty"`
	CouponCodes  []string `json:"coupon_codes,omitempty"`
	Subscription []string `json:"subscription,omitempty"`
}

// NewSubscriptionAddCouponError is a constructor for SubscriptionAddCouponError.
// It creates and returns a pointer to a new SubscriptionAddCouponError instance with the given statusCode and body.
func NewSubscriptionAddCouponError(
	statusCode int,
	body string) *SubscriptionAddCouponError {
	return &SubscriptionAddCouponError{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionAddCouponError.
func (s *SubscriptionAddCouponError) Error() string {
	return fmt.Sprintf("SubscriptionAddCouponError occured %v", s.Body)
}

// SubscriptionRemoveCouponErrors is a custom error.
type SubscriptionRemoveCouponErrors struct {
	apiError.ApiError
	Subscription []string `json:"subscription"`
}

// NewSubscriptionRemoveCouponErrors is a constructor for SubscriptionRemoveCouponErrors.
// It creates and returns a pointer to a new SubscriptionRemoveCouponErrors instance with the given statusCode and body.
func NewSubscriptionRemoveCouponErrors(
	statusCode int,
	body string) *SubscriptionRemoveCouponErrors {
	return &SubscriptionRemoveCouponErrors{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionRemoveCouponErrors.
func (s *SubscriptionRemoveCouponErrors) Error() string {
	return fmt.Sprintf("SubscriptionRemoveCouponErrors occured %v", s.Body)
}

// EventBasedBillingSegmentErrors is a custom error.
type EventBasedBillingSegmentErrors struct {
	apiError.ApiError
	Errors map[string]interface{} `json:"errors,omitempty"`
}

// NewEventBasedBillingSegmentErrors is a constructor for EventBasedBillingSegmentErrors.
// It creates and returns a pointer to a new EventBasedBillingSegmentErrors instance with the given statusCode and body.
func NewEventBasedBillingSegmentErrors(
	statusCode int,
	body string) *EventBasedBillingSegmentErrors {
	return &EventBasedBillingSegmentErrors{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for EventBasedBillingSegmentErrors.
func (e *EventBasedBillingSegmentErrors) Error() string {
	return fmt.Sprintf("EventBasedBillingSegmentErrors occured %v", e.Body)
}

// EventBasedBillingListSegmentsErrors is a custom error.
type EventBasedBillingListSegmentsErrors struct {
	apiError.ApiError
	Errors *models.Errors `json:"errors,omitempty"`
}

// NewEventBasedBillingListSegmentsErrors is a constructor for EventBasedBillingListSegmentsErrors.
// It creates and returns a pointer to a new EventBasedBillingListSegmentsErrors instance with the given statusCode and body.
func NewEventBasedBillingListSegmentsErrors(
	statusCode int,
	body string) *EventBasedBillingListSegmentsErrors {
	return &EventBasedBillingListSegmentsErrors{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for EventBasedBillingListSegmentsErrors.
func (e *EventBasedBillingListSegmentsErrors) Error() string {
	return fmt.Sprintf("EventBasedBillingListSegmentsErrors occured %v", e.Body)
}

// EventBasedBillingSegment is a custom error.
type EventBasedBillingSegment struct {
	apiError.ApiError
	Errors models.EventBasedBillingSegmentError `json:"errors"`
}

// NewEventBasedBillingSegment is a constructor for EventBasedBillingSegment.
// It creates and returns a pointer to a new EventBasedBillingSegment instance with the given statusCode and body.
func NewEventBasedBillingSegment(
	statusCode int,
	body string) *EventBasedBillingSegment {
	return &EventBasedBillingSegment{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for EventBasedBillingSegment.
func (e *EventBasedBillingSegment) Error() string {
	return fmt.Sprintf("EventBasedBillingSegment occured %v", e.Body)
}

// SubscriptionsMrrErrorResponse is a custom error.
type SubscriptionsMrrErrorResponse struct {
	apiError.ApiError
	Errors models.AttributeError `json:"errors"`
}

// NewSubscriptionsMrrErrorResponse is a constructor for SubscriptionsMrrErrorResponse.
// It creates and returns a pointer to a new SubscriptionsMrrErrorResponse instance with the given statusCode and body.
func NewSubscriptionsMrrErrorResponse(
	statusCode int,
	body string) *SubscriptionsMrrErrorResponse {
	return &SubscriptionsMrrErrorResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionsMrrErrorResponse.
func (s *SubscriptionsMrrErrorResponse) Error() string {
	return fmt.Sprintf("SubscriptionsMrrErrorResponse occured %v", s.Body)
}

// RefundPrepaymentBaseErrorsResponse is a custom error.
// Errors returned on creating a refund prepayment when bad request
type RefundPrepaymentBaseErrorsResponse struct {
	apiError.ApiError
	Errors *models.RefundPrepaymentBaseRefundError `json:"errors,omitempty"`
}

// NewRefundPrepaymentBaseErrorsResponse is a constructor for RefundPrepaymentBaseErrorsResponse.
// It creates and returns a pointer to a new RefundPrepaymentBaseErrorsResponse instance with the given statusCode and body.
func NewRefundPrepaymentBaseErrorsResponse(
	statusCode int,
	body string) *RefundPrepaymentBaseErrorsResponse {
	return &RefundPrepaymentBaseErrorsResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for RefundPrepaymentBaseErrorsResponse.
func (r *RefundPrepaymentBaseErrorsResponse) Error() string {
	return fmt.Sprintf("RefundPrepaymentBaseErrorsResponse occured %v", r.Body)
}

// RefundPrepaymentAggregatedErrorsResponse is a custom error.
// Errors returned on creating a refund prepayment, grouped by field, as arrays of strings.
type RefundPrepaymentAggregatedErrorsResponse struct {
	apiError.ApiError
	Errors *models.RefundPrepaymentAggregatedError `json:"errors,omitempty"`
}

// NewRefundPrepaymentAggregatedErrorsResponse is a constructor for RefundPrepaymentAggregatedErrorsResponse.
// It creates and returns a pointer to a new RefundPrepaymentAggregatedErrorsResponse instance with the given statusCode and body.
func NewRefundPrepaymentAggregatedErrorsResponse(
	statusCode int,
	body string) *RefundPrepaymentAggregatedErrorsResponse {
	return &RefundPrepaymentAggregatedErrorsResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for RefundPrepaymentAggregatedErrorsResponse.
func (r *RefundPrepaymentAggregatedErrorsResponse) Error() string {
	return fmt.Sprintf("RefundPrepaymentAggregatedErrorsResponse occured %v", r.Body)
}

// ProformaBadRequestErrorResponse is a custom error.
type ProformaBadRequestErrorResponse struct {
	apiError.ApiError
	Errors *models.ProformaError `json:"errors,omitempty"`
}

// NewProformaBadRequestErrorResponse is a constructor for ProformaBadRequestErrorResponse.
// It creates and returns a pointer to a new ProformaBadRequestErrorResponse instance with the given statusCode and body.
func NewProformaBadRequestErrorResponse(
	statusCode int,
	body string) *ProformaBadRequestErrorResponse {
	return &ProformaBadRequestErrorResponse{
		ApiError: apiError.ApiError{
			StatusCode: statusCode,
			Body:       body,
		},
	}
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ProformaBadRequestErrorResponse.
func (p *ProformaBadRequestErrorResponse) Error() string {
	return fmt.Sprintf("ProformaBadRequestErrorResponse occured %v", p.Body)
}
