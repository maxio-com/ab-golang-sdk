package advancedbilling

import (
    "context"
    "fmt"
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
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.SubscriptionGroupPrepaymentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionGroupPrepaymentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ListPrepaymentsForSubscriptionGroup takes context, uid, filterDateField, filterEndDate, filterStartDate, page, perPage as parameters and
// returns an models.ApiResponse with models.ListSubscriptionGroupPrepaymentResponse data and
// an error if there was an issue with the request or response.
// This request will list a subscription group's prepayments.
func (s *SubscriptionGroupInvoiceAccountController) ListPrepaymentsForSubscriptionGroup(
    ctx context.Context,
    uid string,
    filterDateField *models.ListSubscriptionGroupPrepaymentDateField,
    filterEndDate *time.Time,
    filterStartDate *time.Time,
    page *int,
    perPage *int) (
    models.ApiResponse[models.ListSubscriptionGroupPrepaymentResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/subscription_groups/%v/prepayments.json", uid),
    )
    req.Authenticate(true)
    if filterDateField != nil {
        req.QueryParam("filter[date_field]", *filterDateField)
    }
    if filterEndDate != nil {
        req.QueryParam("filter[end_date]", filterEndDate.Format(models.DEFAULT_DATE))
    }
    if filterStartDate != nil {
        req.QueryParam("filter[start_date]", filterStartDate.Format(models.DEFAULT_DATE))
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    
    var result models.ListSubscriptionGroupPrepaymentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListSubscriptionGroupPrepaymentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
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
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ServiceCreditResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ServiceCreditResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
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
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}
