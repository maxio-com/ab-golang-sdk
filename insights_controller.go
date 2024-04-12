package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "time"
)

// InsightsController represents a controller struct.
type InsightsController struct {
    baseController
}

// NewInsightsController creates a new instance of InsightsController.
// It takes a baseController as a parameter and returns a pointer to the InsightsController.
func NewInsightsController(baseController baseController) *InsightsController {
    insightsController := InsightsController{baseController: baseController}
    return &insightsController
}

// ReadSiteStats takes context as parameters and
// returns an models.ApiResponse with models.SiteSummary data and
// an error if there was an issue with the request or response.
// The Stats API is a very basic view of some Site-level stats. This API call only answers with JSON responses. An XML version is not provided.
// ## Stats Documentation
// There currently is not a complimentary matching set of documentation that compliments this endpoint. However, each Site's dashboard will reflect the summary of information provided in the Stats reposnse.
// ```
// https://subdomain.chargify.com/dashboard
// ```
func (i *InsightsController) ReadSiteStats(ctx context.Context) (
    models.ApiResponse[models.SiteSummary],
    error) {
    req := i.prepareRequest(ctx, "GET", "/stats.json")
    req.Authenticate(NewAuth("BasicAuth"))
    var result models.SiteSummary
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteSummary](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadMrr takes context, atTime, subscriptionId as parameters and
// returns an models.ApiResponse with models.MRRResponse data and
// an error if there was an issue with the request or response.
// Deprecated: readMrr is deprecated
// This endpoint returns your site's current MRR, including plan and usage breakouts.
func (i *InsightsController) ReadMrr(
    ctx context.Context,
    atTime *time.Time,
    subscriptionId *int) (
    models.ApiResponse[models.MRRResponse],
    error) {
    req := i.prepareRequest(ctx, "GET", "/mrr.json")
    req.Authenticate(NewAuth("BasicAuth"))
    if atTime != nil {
        req.QueryParam("at_time", atTime.Format(time.RFC3339))
    }
    if subscriptionId != nil {
        req.QueryParam("subscription_id", *subscriptionId)
    }
    var result models.MRRResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.MRRResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListMrrMovementsInput represents the input of the ListMrrMovements endpoint.
type ListMrrMovementsInput struct {
    // optionally filter results by subscription
    SubscriptionId *int                     
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page           *int                     
    // This parameter indicates how many records to fetch in each request. Default value is 10. The maximum allowed values is 50; any per_page value over 50 will be changed to 50.
    // Use in query `per_page=20`.
    PerPage        *int                     
    // Controls the order in which results are returned.
    // Use in query `direction=asc`.
    Direction      *models.SortingDirection 
}

// ListMrrMovements takes context, subscriptionId, page, perPage, direction as parameters and
// returns an models.ApiResponse with models.ListMRRResponse data and
// an error if there was an issue with the request or response.
// Deprecated: listMrrMovements is deprecated
// This endpoint returns your site's MRR movements.
// ## Understanding MRR movements
// This endpoint will aid in accessing your site's [MRR Report](https://chargify.zendesk.com/hc/en-us/articles/4407838249627) data.
// Whenever a subscription event occurs that causes your site's MRR to change (such as a signup or upgrade), we record an MRR movement. These records are accessible via the MRR Movements endpoint.
// Each MRR Movement belongs to a subscription and contains a timestamp, category, and an amount. `line_items` represent the subscription's product configuration at the time of the movement.
// ### Plan & Usage Breakouts
// In the MRR Report UI, we support a setting to [include or exclude](https://chargify.zendesk.com/hc/en-us/articles/4407838249627#displaying-component-based-metered-usage-in-mrr) usage revenue. In the MRR APIs, responses include `plan` and `usage` breakouts.
// Plan includes revenue from:
// * Products
// * Quantity-Based Components
// * On/Off Components
// Usage includes revenue from:
// * Metered Components
// * Prepaid Usage Components
func (i *InsightsController) ListMrrMovements(
    ctx context.Context,
    input ListMrrMovementsInput) (
    models.ApiResponse[models.ListMRRResponse],
    error) {
    req := i.prepareRequest(ctx, "GET", "/mrr_movements.json")
    req.Authenticate(NewAuth("BasicAuth"))
    if input.SubscriptionId != nil {
        req.QueryParam("subscription_id", *input.SubscriptionId)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.Direction != nil {
        req.QueryParam("direction", *input.Direction)
    }
    var result models.ListMRRResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListMRRResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListMrrPerSubscriptionInput represents the input of the ListMrrPerSubscription endpoint.
type ListMrrPerSubscriptionInput struct {
    // Filter to use for List MRR per subscription operation
    Filter    *models.ListMrrFilter 
    // Submit a timestamp in ISO8601 format to request MRR for a historic time. Use in query: `at_time=2022-01-10T10:00:00-05:00`.
    AtTime    *string               
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page      *int                  
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage   *int                  
    // Controls the order in which results are returned. Records are ordered by subscription_id in ascending order by default. Use in query `direction=desc`.
    Direction *models.Direction     
}

// ListMrrPerSubscription takes context, filter, atTime, page, perPage, direction as parameters and
// returns an models.ApiResponse with models.SubscriptionMRRResponse data and
// an error if there was an issue with the request or response.
// Deprecated: listMrrPerSubscription is deprecated
// This endpoint returns your site's current MRR, including plan and usage breakouts split per subscription.
func (i *InsightsController) ListMrrPerSubscription(
    ctx context.Context,
    input ListMrrPerSubscriptionInput) (
    models.ApiResponse[models.SubscriptionMRRResponse],
    error) {
    req := i.prepareRequest(ctx, "GET", "/subscriptions_mrr.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSubscriptionsMrrErrorResponse},
    })
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    if input.AtTime != nil {
        req.QueryParam("at_time", *input.AtTime)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.Direction != nil {
        req.QueryParam("direction", *input.Direction)
    }
    var result models.SubscriptionMRRResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionMRRResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
