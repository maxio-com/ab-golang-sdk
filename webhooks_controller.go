package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// WebhooksController represents a controller struct.
type WebhooksController struct {
    baseController
}

// NewWebhooksController creates a new instance of WebhooksController.
// It takes a baseController as a parameter and returns a pointer to the WebhooksController.
func NewWebhooksController(baseController baseController) *WebhooksController {
    webhooksController := WebhooksController{baseController: baseController}
    return &webhooksController
}

// ListWebhooksInput represents the input of the ListWebhooks endpoint.
type ListWebhooksInput struct {
    // Webhooks with matching status would be returned.
    Status       *models.WebhookStatus 
    // Format YYYY-MM-DD. Returns Webhooks with the created_at date greater than or equal to the one specified.
    SinceDate    *string               
    // Format YYYY-MM-DD. Returns Webhooks with the created_at date less than or equal to the one specified.
    UntilDate    *string               
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page         *int                  
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage      *int                  
    // The order in which the Webhooks are returned.
    Order        *models.WebhookOrder  
    // The Chargify id of a subscription you'd like to filter for
    Subscription *int                  
}

// ListWebhooks takes context, status, sinceDate, untilDate, page, perPage, order, subscription as parameters and
// returns an models.ApiResponse with []models.WebhookResponse data and
// an error if there was an issue with the request or response.
// ## Webhooks Intro
// The Webhooks API allows you to view a list of all webhooks and to selectively resend individual or groups of webhooks. Webhooks will be sent on endpoints specified by you. Endpoints can be added via API or Web UI. There is also an option to enable / disable webhooks via API request.
// We recommend that you review Chargify's webhook documentation located in our help site. The following resources will help guide you on how to use webhooks in Chargify, in addition to these webhook endpoints:
// + [Adding/editing new webhooks](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404448450317#configure-webhook-url)
// + [Webhooks introduction and delivery information](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405568068365#webhooks-introduction-0-0)
// + [Main webhook overview](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405357509645-Webhooks-Reference#webhooks-reference-0-0)
// + [Available webhooks and payloads](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405357509645-Webhooks-Reference#events)
// ## List Webhooks for a Site
// This method allows you to fetch data about webhooks. You can pass query parameters if you want to filter webhooks.
func (w *WebhooksController) ListWebhooks(
    ctx context.Context,
    input ListWebhooksInput) (
    models.ApiResponse[[]models.WebhookResponse],
    error) {
    req := w.prepareRequest(ctx, "GET", "/webhooks.json")
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Status != nil {
        req.QueryParam("status", *input.Status)
    }
    if input.SinceDate != nil {
        req.QueryParam("since_date", *input.SinceDate)
    }
    if input.UntilDate != nil {
        req.QueryParam("until_date", *input.UntilDate)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.Order != nil {
        req.QueryParam("order", *input.Order)
    }
    if input.Subscription != nil {
        req.QueryParam("subscription", *input.Subscription)
    }
    var result []models.WebhookResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.WebhookResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// EnableWebhooks takes context, body as parameters and
// returns an models.ApiResponse with models.EnableWebhooksResponse data and
// an error if there was an issue with the request or response.
// This method allows you to enable webhooks via API for your site
func (w *WebhooksController) EnableWebhooks(
    ctx context.Context,
    body *models.EnableWebhooksRequest) (
    models.ApiResponse[models.EnableWebhooksResponse],
    error) {
    req := w.prepareRequest(ctx, "PUT", "/webhooks/settings.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.EnableWebhooksResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.EnableWebhooksResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReplayWebhooks takes context, body as parameters and
// returns an models.ApiResponse with models.ReplayWebhooksResponse data and
// an error if there was an issue with the request or response.
// Posting to the replay endpoint does not immediately resend the webhooks. They are added to a queue and will be sent as soon as possible, depending on available system resources.
// You may submit an array of up to 1000 webhook IDs to replay in the request.
func (w *WebhooksController) ReplayWebhooks(
    ctx context.Context,
    body *models.ReplayWebhooksRequest) (
    models.ApiResponse[models.ReplayWebhooksResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/webhooks/replay.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.ReplayWebhooksResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ReplayWebhooksResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateEndpoint takes context, body as parameters and
// returns an models.ApiResponse with models.EndpointResponse data and
// an error if there was an issue with the request or response.
// The Chargify API allows you to create an endpoint and assign a list of webhooks subscriptions (events) to it.
// You can check available events here.
// [Event keys](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405357509645-Webhooks-Reference#example-payloads)
func (w *WebhooksController) CreateEndpoint(
    ctx context.Context,
    body *models.CreateOrUpdateEndpointRequest) (
    models.ApiResponse[models.EndpointResponse],
    error) {
    req := w.prepareRequest(ctx, "POST", "/endpoints.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.EndpointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.EndpointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListEndpoints takes context as parameters and
// returns an models.ApiResponse with []models.Endpoint data and
// an error if there was an issue with the request or response.
// This method returns created endpoints for site.
func (w *WebhooksController) ListEndpoints(ctx context.Context) (
    models.ApiResponse[[]models.Endpoint],
    error) {
    req := w.prepareRequest(ctx, "GET", "/endpoints.json")
    req.Authenticate(NewAuth("BasicAuth"))
    var result []models.Endpoint
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Endpoint](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateEndpoint takes context, endpointId, body as parameters and
// returns an models.ApiResponse with models.EndpointResponse data and
// an error if there was an issue with the request or response.
// You can update an Endpoint via the API with a PUT request to the resource endpoint.
// You can change the `url` of your endpoint which consumes webhooks or list of `webhook_subscriptions`.
// Check available [Event keys](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404448450317-Webhooks#configure-webhook-url).
// Always send a complete list of events which you want subscribe/watch.
// Sending an PUT request for existing endpoint with empty list of `webhook_subscriptions` will end with unsubscribe from all events.
// If you want unsubscribe from specific event, just send a list of `webhook_subscriptions` without the specific event key.
func (w *WebhooksController) UpdateEndpoint(
    ctx context.Context,
    endpointId int,
    body *models.CreateOrUpdateEndpointRequest) (
    models.ApiResponse[models.EndpointResponse],
    error) {
    req := w.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/endpoints/%v.json", endpointId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.EndpointResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.EndpointResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
