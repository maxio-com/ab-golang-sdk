/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package errors

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// ComponentAllocationError is a custom error.
type ComponentAllocationError struct {
    https.ApiError
    Errors         []models.ComponentAllocationErrorItem `json:"errors,omitempty"`
}

// NewComponentAllocationError is a constructor for ComponentAllocationError.
// It creates and returns a pointer to a new ComponentAllocationError instance with the given statusCode and body.
func NewComponentAllocationError(apiError https.ApiError) error {
    return &ComponentAllocationError{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ComponentAllocationError.
func (c ComponentAllocationError) Error() string {
    return fmt.Sprintf("ComponentAllocationError occured: %v", c.Message)
}

// ComponentPricePointError is a custom error.
type ComponentPricePointError struct {
    https.ApiError
    Errors         []models.ComponentPricePointErrorItem `json:"errors,omitempty"`
}

// NewComponentPricePointError is a constructor for ComponentPricePointError.
// It creates and returns a pointer to a new ComponentPricePointError instance with the given statusCode and body.
func NewComponentPricePointError(apiError https.ApiError) error {
    return &ComponentPricePointError{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ComponentPricePointError.
func (c ComponentPricePointError) Error() string {
    return fmt.Sprintf("ComponentPricePointError occured: %v", c.Message)
}

// CustomerErrorResponse is a custom error.
type CustomerErrorResponse struct {
    https.ApiError
    Errors         *models.CustomerErrorResponseErrors `json:"errors,omitempty"`
}

// NewCustomerErrorResponse is a constructor for CustomerErrorResponse.
// It creates and returns a pointer to a new CustomerErrorResponse instance with the given statusCode and body.
func NewCustomerErrorResponse(apiError https.ApiError) error {
    return &CustomerErrorResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for CustomerErrorResponse.
func (c CustomerErrorResponse) Error() string {
    return fmt.Sprintf("CustomerErrorResponse occured: %v", c.Message)
}

// ErrorArrayMapResponse is a custom error.
type ErrorArrayMapResponse struct {
    https.ApiError
    Errors         map[string]interface{} `json:"errors,omitempty"`
}

// NewErrorArrayMapResponse is a constructor for ErrorArrayMapResponse.
// It creates and returns a pointer to a new ErrorArrayMapResponse instance with the given statusCode and body.
func NewErrorArrayMapResponse(apiError https.ApiError) error {
    return &ErrorArrayMapResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ErrorArrayMapResponse.
func (e ErrorArrayMapResponse) Error() string {
    return fmt.Sprintf("ErrorArrayMapResponse occured: %v", e.Message)
}

// ErrorListResponse is a custom error.
// Error which contains list of messages.
type ErrorListResponse struct {
    https.ApiError
    Errors         []string `json:"errors"`
}

// NewErrorListResponse is a constructor for ErrorListResponse.
// It creates and returns a pointer to a new ErrorListResponse instance with the given statusCode and body.
func NewErrorListResponse(apiError https.ApiError) error {
    return &ErrorListResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ErrorListResponse.
func (e ErrorListResponse) Error() string {
    return fmt.Sprintf("ErrorListResponse occured: %v", e.Message)
}

// ErrorStringMapResponse is a custom error.
type ErrorStringMapResponse struct {
    https.ApiError
    Errors         map[string]string `json:"errors,omitempty"`
}

// NewErrorStringMapResponse is a constructor for ErrorStringMapResponse.
// It creates and returns a pointer to a new ErrorStringMapResponse instance with the given statusCode and body.
func NewErrorStringMapResponse(apiError https.ApiError) error {
    return &ErrorStringMapResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ErrorStringMapResponse.
func (e ErrorStringMapResponse) Error() string {
    return fmt.Sprintf("ErrorStringMapResponse occured: %v", e.Message)
}

// EventBasedBillingListSegmentsErrors is a custom error.
type EventBasedBillingListSegmentsErrors struct {
    https.ApiError
    Errors         *models.Errors `json:"errors,omitempty"`
}

// NewEventBasedBillingListSegmentsErrors is a constructor for EventBasedBillingListSegmentsErrors.
// It creates and returns a pointer to a new EventBasedBillingListSegmentsErrors instance with the given statusCode and body.
func NewEventBasedBillingListSegmentsErrors(apiError https.ApiError) error {
    return &EventBasedBillingListSegmentsErrors{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for EventBasedBillingListSegmentsErrors.
func (e EventBasedBillingListSegmentsErrors) Error() string {
    return fmt.Sprintf("EventBasedBillingListSegmentsErrors occured: %v", e.Message)
}

// EventBasedBillingSegment is a custom error.
type EventBasedBillingSegment struct {
    https.ApiError
    Errors         models.EventBasedBillingSegmentError `json:"errors"`
}

// NewEventBasedBillingSegment is a constructor for EventBasedBillingSegment.
// It creates and returns a pointer to a new EventBasedBillingSegment instance with the given statusCode and body.
func NewEventBasedBillingSegment(apiError https.ApiError) error {
    return &EventBasedBillingSegment{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for EventBasedBillingSegment.
func (e EventBasedBillingSegment) Error() string {
    return fmt.Sprintf("EventBasedBillingSegment occured: %v", e.Message)
}

// EventBasedBillingSegmentErrors is a custom error.
type EventBasedBillingSegmentErrors struct {
    https.ApiError
    Errors         map[string]interface{} `json:"errors,omitempty"`
}

// NewEventBasedBillingSegmentErrors is a constructor for EventBasedBillingSegmentErrors.
// It creates and returns a pointer to a new EventBasedBillingSegmentErrors instance with the given statusCode and body.
func NewEventBasedBillingSegmentErrors(apiError https.ApiError) error {
    return &EventBasedBillingSegmentErrors{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for EventBasedBillingSegmentErrors.
func (e EventBasedBillingSegmentErrors) Error() string {
    return fmt.Sprintf("EventBasedBillingSegmentErrors occured: %v", e.Message)
}

// ProductPricePointErrorResponse is a custom error.
type ProductPricePointErrorResponse struct {
    https.ApiError
    Errors         models.ProductPricePointErrors `json:"errors"`
}

// NewProductPricePointErrorResponse is a constructor for ProductPricePointErrorResponse.
// It creates and returns a pointer to a new ProductPricePointErrorResponse instance with the given statusCode and body.
func NewProductPricePointErrorResponse(apiError https.ApiError) error {
    return &ProductPricePointErrorResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ProductPricePointErrorResponse.
func (p ProductPricePointErrorResponse) Error() string {
    return fmt.Sprintf("ProductPricePointErrorResponse occured: %v", p.Message)
}

// ProformaBadRequestErrorResponse is a custom error.
type ProformaBadRequestErrorResponse struct {
    https.ApiError
    Errors         *models.ProformaError `json:"errors,omitempty"`
}

// NewProformaBadRequestErrorResponse is a constructor for ProformaBadRequestErrorResponse.
// It creates and returns a pointer to a new ProformaBadRequestErrorResponse instance with the given statusCode and body.
func NewProformaBadRequestErrorResponse(apiError https.ApiError) error {
    return &ProformaBadRequestErrorResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for ProformaBadRequestErrorResponse.
func (p ProformaBadRequestErrorResponse) Error() string {
    return fmt.Sprintf("ProformaBadRequestErrorResponse occured: %v", p.Message)
}

// RefundPrepaymentBaseErrorsResponse is a custom error.
// Errors returned on creating a refund prepayment when bad request
type RefundPrepaymentBaseErrorsResponse struct {
    https.ApiError
    Errors         *models.RefundPrepaymentBaseRefundError `json:"errors,omitempty"`
}

// NewRefundPrepaymentBaseErrorsResponse is a constructor for RefundPrepaymentBaseErrorsResponse.
// It creates and returns a pointer to a new RefundPrepaymentBaseErrorsResponse instance with the given statusCode and body.
func NewRefundPrepaymentBaseErrorsResponse(apiError https.ApiError) error {
    return &RefundPrepaymentBaseErrorsResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for RefundPrepaymentBaseErrorsResponse.
func (r RefundPrepaymentBaseErrorsResponse) Error() string {
    return fmt.Sprintf("RefundPrepaymentBaseErrorsResponse occured: %v", r.Message)
}

// SingleErrorResponse is a custom error.
type SingleErrorResponse struct {
    https.ApiError
    MError         string `json:"error"`
}

// NewSingleErrorResponse is a constructor for SingleErrorResponse.
// It creates and returns a pointer to a new SingleErrorResponse instance with the given statusCode and body.
func NewSingleErrorResponse(apiError https.ApiError) error {
    return &SingleErrorResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SingleErrorResponse.
func (s SingleErrorResponse) Error() string {
    return fmt.Sprintf("SingleErrorResponse occured: %v", s.Message)
}

// SingleStringErrorResponse is a custom error.
type SingleStringErrorResponse struct {
    https.ApiError
    Errors         *string `json:"errors,omitempty"`
}

// NewSingleStringErrorResponse is a constructor for SingleStringErrorResponse.
// It creates and returns a pointer to a new SingleStringErrorResponse instance with the given statusCode and body.
func NewSingleStringErrorResponse(apiError https.ApiError) error {
    return &SingleStringErrorResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SingleStringErrorResponse.
func (s SingleStringErrorResponse) Error() string {
    return fmt.Sprintf("SingleStringErrorResponse occured: %v", s.Message)
}

// SubscriptionAddCouponError is a custom error.
type SubscriptionAddCouponError struct {
    https.ApiError
    Codes          []string `json:"codes,omitempty"`
    CouponCode     []string `json:"coupon_code,omitempty"`
    CouponCodes    []string `json:"coupon_codes,omitempty"`
    Subscription   []string `json:"subscription,omitempty"`
}

// NewSubscriptionAddCouponError is a constructor for SubscriptionAddCouponError.
// It creates and returns a pointer to a new SubscriptionAddCouponError instance with the given statusCode and body.
func NewSubscriptionAddCouponError(apiError https.ApiError) error {
    return &SubscriptionAddCouponError{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionAddCouponError.
func (s SubscriptionAddCouponError) Error() string {
    return fmt.Sprintf("SubscriptionAddCouponError occured: %v", s.Message)
}

// SubscriptionComponentAllocationError is a custom error.
type SubscriptionComponentAllocationError struct {
    https.ApiError
    Errors         []models.SubscriptionComponentAllocationErrorItem `json:"errors,omitempty"`
}

// NewSubscriptionComponentAllocationError is a constructor for SubscriptionComponentAllocationError.
// It creates and returns a pointer to a new SubscriptionComponentAllocationError instance with the given statusCode and body.
func NewSubscriptionComponentAllocationError(apiError https.ApiError) error {
    return &SubscriptionComponentAllocationError{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionComponentAllocationError.
func (s SubscriptionComponentAllocationError) Error() string {
    return fmt.Sprintf("SubscriptionComponentAllocationError occured: %v", s.Message)
}

// SubscriptionGroupCreateErrorResponse is a custom error.
type SubscriptionGroupCreateErrorResponse struct {
    https.ApiError
    Errors         models.SubscriptionGroupCreateErrorResponseErrors `json:"errors"`
}

// NewSubscriptionGroupCreateErrorResponse is a constructor for SubscriptionGroupCreateErrorResponse.
// It creates and returns a pointer to a new SubscriptionGroupCreateErrorResponse instance with the given statusCode and body.
func NewSubscriptionGroupCreateErrorResponse(apiError https.ApiError) error {
    return &SubscriptionGroupCreateErrorResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionGroupCreateErrorResponse.
func (s SubscriptionGroupCreateErrorResponse) Error() string {
    return fmt.Sprintf("SubscriptionGroupCreateErrorResponse occured: %v", s.Message)
}

// SubscriptionGroupSignupErrorResponse is a custom error.
type SubscriptionGroupSignupErrorResponse struct {
    https.ApiError
    Errors         models.SubscriptionGroupSignupError `json:"errors"`
}

// NewSubscriptionGroupSignupErrorResponse is a constructor for SubscriptionGroupSignupErrorResponse.
// It creates and returns a pointer to a new SubscriptionGroupSignupErrorResponse instance with the given statusCode and body.
func NewSubscriptionGroupSignupErrorResponse(apiError https.ApiError) error {
    return &SubscriptionGroupSignupErrorResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionGroupSignupErrorResponse.
func (s SubscriptionGroupSignupErrorResponse) Error() string {
    return fmt.Sprintf("SubscriptionGroupSignupErrorResponse occured: %v", s.Message)
}

// SubscriptionGroupUpdateErrorResponse is a custom error.
type SubscriptionGroupUpdateErrorResponse struct {
    https.ApiError
    Errors         *models.SubscriptionGroupUpdateError `json:"errors,omitempty"`
}

// NewSubscriptionGroupUpdateErrorResponse is a constructor for SubscriptionGroupUpdateErrorResponse.
// It creates and returns a pointer to a new SubscriptionGroupUpdateErrorResponse instance with the given statusCode and body.
func NewSubscriptionGroupUpdateErrorResponse(apiError https.ApiError) error {
    return &SubscriptionGroupUpdateErrorResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionGroupUpdateErrorResponse.
func (s SubscriptionGroupUpdateErrorResponse) Error() string {
    return fmt.Sprintf("SubscriptionGroupUpdateErrorResponse occured: %v", s.Message)
}

// SubscriptionRemoveCouponErrors is a custom error.
type SubscriptionRemoveCouponErrors struct {
    https.ApiError
    Subscription   []string `json:"subscription"`
}

// NewSubscriptionRemoveCouponErrors is a constructor for SubscriptionRemoveCouponErrors.
// It creates and returns a pointer to a new SubscriptionRemoveCouponErrors instance with the given statusCode and body.
func NewSubscriptionRemoveCouponErrors(apiError https.ApiError) error {
    return &SubscriptionRemoveCouponErrors{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionRemoveCouponErrors.
func (s SubscriptionRemoveCouponErrors) Error() string {
    return fmt.Sprintf("SubscriptionRemoveCouponErrors occured: %v", s.Message)
}

// SubscriptionResponseError is a custom error.
type SubscriptionResponseError struct {
    https.ApiError
    Subscription   *models.Subscription `json:"subscription,omitempty"`
}

// NewSubscriptionResponseError is a constructor for SubscriptionResponseError.
// It creates and returns a pointer to a new SubscriptionResponseError instance with the given statusCode and body.
func NewSubscriptionResponseError(apiError https.ApiError) error {
    return &SubscriptionResponseError{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionResponseError.
func (s SubscriptionResponseError) Error() string {
    return fmt.Sprintf("SubscriptionResponseError occured: %v", s.Message)
}

// SubscriptionsMrrErrorResponse is a custom error.
type SubscriptionsMrrErrorResponse struct {
    https.ApiError
    Errors         models.AttributeError `json:"errors"`
}

// NewSubscriptionsMrrErrorResponse is a constructor for SubscriptionsMrrErrorResponse.
// It creates and returns a pointer to a new SubscriptionsMrrErrorResponse instance with the given statusCode and body.
func NewSubscriptionsMrrErrorResponse(apiError https.ApiError) error {
    return &SubscriptionsMrrErrorResponse{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for SubscriptionsMrrErrorResponse.
func (s SubscriptionsMrrErrorResponse) Error() string {
    return fmt.Sprintf("SubscriptionsMrrErrorResponse occured: %v", s.Message)
}

// TooManyManagementLinkRequestsError is a custom error.
type TooManyManagementLinkRequestsError struct {
    https.ApiError
    Errors         models.TooManyManagementLinkRequests `json:"errors"`
}

// NewTooManyManagementLinkRequestsError is a constructor for TooManyManagementLinkRequestsError.
// It creates and returns a pointer to a new TooManyManagementLinkRequestsError instance with the given statusCode and body.
func NewTooManyManagementLinkRequestsError(apiError https.ApiError) error {
    return &TooManyManagementLinkRequestsError{
		ApiError: apiError,
    }
}

// Error implements the Error method for the error interface.
// It returns a formatted error message for TooManyManagementLinkRequestsError.
func (t TooManyManagementLinkRequestsError) Error() string {
    return fmt.Sprintf("TooManyManagementLinkRequestsError occured: %v", t.Message)
}
