package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
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
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.AccountBalances
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountBalances](decoder)
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
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'."},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.CreatePrepaymentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CreatePrepaymentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListPrepaymentsInput represents the input of the ListPrepayments endpoint.
type ListPrepaymentsInput struct {
    // The Chargify id of the subscription
    SubscriptionId int                           
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page           *int                          
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage        *int                          
    // Filter to use for List Prepayments operations
    Filter         *models.ListPrepaymentsFilter 
}

// ListPrepayments takes context, subscriptionId, page, perPage, filter as parameters and
// returns an models.ApiResponse with models.PrepaymentsResponse data and
// an error if there was an issue with the request or response.
// This request will list a subscription's prepayments.
func (s *SubscriptionInvoiceAccountController) ListPrepayments(
    ctx context.Context,
    input ListPrepaymentsInput) (
    models.ApiResponse[models.PrepaymentsResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/subscriptions/%v/prepayments.json", input.SubscriptionId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
    })
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    
    var result models.PrepaymentsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PrepaymentsResponse](decoder)
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
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'."},
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
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'."},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
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
    prepaymentId int64,
    body *models.RefundPrepaymentRequest) (
    models.ApiResponse[models.PrepaymentResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/prepayments/%v/refunds.json", subscriptionId, prepaymentId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "400": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewRefundPrepaymentBaseErrorsResponse},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'."},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.PrepaymentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PrepaymentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
