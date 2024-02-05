package advancedbilling

import (
    "context"
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
    req.Authenticate(true)
    var result models.SiteSummary
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteSummary](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
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
    req.Authenticate(true)
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
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.MRRResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
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
    subscriptionId *int,
    page *int,
    perPage *int,
    direction *models.SortingDirection) (
    models.ApiResponse[models.ListMRRResponse],
    error) {
    req := i.prepareRequest(ctx, "GET", "/mrr_movements.json")
    req.Authenticate(true)
    if subscriptionId != nil {
        req.QueryParam("subscription_id", *subscriptionId)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if direction != nil {
        req.QueryParam("direction", *direction)
    }
    var result models.ListMRRResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListMRRResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ListMrrPerSubscription takes context, filterSubscriptionIds, atTime, page, perPage, direction as parameters and
// returns an models.ApiResponse with models.SubscriptionMRRResponse data and
// an error if there was an issue with the request or response.
// Deprecated: listMrrPerSubscription is deprecated
// This endpoint returns your site's current MRR, including plan and usage breakouts split per subscription.
func (i *InsightsController) ListMrrPerSubscription(
    ctx context.Context,
    filterSubscriptionIds []int,
    atTime *string,
    page *int,
    perPage *int,
    direction *models.Direction) (
    models.ApiResponse[models.SubscriptionMRRResponse],
    error) {
    req := i.prepareRequest(ctx, "GET", "/subscriptions_mrr.json")
    req.Authenticate(true)
    if filterSubscriptionIds != nil {
        req.QueryParam("filter[subscription_ids]", filterSubscriptionIds)
    }
    if atTime != nil {
        req.QueryParam("at_time", *atTime)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if direction != nil {
        req.QueryParam("direction", *direction)
    }
    var result models.SubscriptionMRRResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionMRRResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 400 {
        err = errors.NewSubscriptionsMrrErrorResponse(400, "Bad Request")
    }
    return models.NewApiResponse(result, resp), err
}
