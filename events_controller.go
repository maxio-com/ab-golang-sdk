/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// EventsController represents a controller struct.
type EventsController struct {
    baseController
}

// NewEventsController creates a new instance of EventsController.
// It takes a baseController as a parameter and returns a pointer to the EventsController.
func NewEventsController(baseController baseController) *EventsController {
    eventsController := EventsController{baseController: baseController}
    return &eventsController
}

// ListEventsInput represents the input of the ListEvents endpoint.
type ListEventsInput struct {
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page          *int                        
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage       *int                        
    // Returns events with an id greater than or equal to the one specified
    SinceId       *int64                      
    // Returns events with an id less than or equal to the one specified
    MaxId         *int64                      
    // The sort direction of the returned events.
    Direction     *models.Direction           
    // You can pass multiple event keys after comma.
    // Use in query `filter=signup_success,payment_success`.
    Filter        []models.EventType          
    // The type of filter you would like to apply to your search.
    DateField     *models.ListEventsDateField 
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
    StartDate     *string                     
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
    EndDate       *string                     
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
    StartDatetime *string                     
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.
    EndDatetime   *string                     
}

// ListEvents takes context, page, perPage, sinceId, maxId, direction, filter, dateField, startDate, endDate, startDatetime, endDatetime as parameters and
// returns an models.ApiResponse with []models.EventResponse data and
// an error if there was an issue with the request or response.
// ## Events Intro
// Advanced Billing Events include various activity that happens around a Site. This information is **especially** useful to track down issues that arise when subscriptions are not created due to errors.
// Within the Advanced Billing UI, "Events" are referred to as "Site Activity".  Full documentation on how to record view Events / Site Activty in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24250671733517-Site-Activity).
// ## List Events for a Site
// This method will retrieve a list of events for a site. Use query string filters to narrow down results. You may use the `key` filter as part of your query string to narrow down results.
// ### Legacy Filters
// The following keys are no longer supported.
// + `payment_failure_recreated`
// + `payment_success_recreated`
// + `renewal_failure_recreated`
// + `renewal_success_recreated`
// + `zferral_revenue_post_failure` - (Specific to the deprecated Zferral integration)
// + `zferral_revenue_post_success` - (Specific to the deprecated Zferral integration)
// ## Event Specific Data
// Event Specific Data
// Each event type has its own `event_specific_data` specified.
// Here’s an example event for the `subscription_product_change` event:
// ```
// {
// "event": {
// "id": 351,
// "key": "subscription_product_change",
// "message": "Product changed on Marky Mark's subscription from 'Basic' to 'Pro'",
// "subscription_id": 205,
// "event_specific_data": {
// "new_product_id": 3,
// "previous_product_id": 2
// },
// "created_at": "2012-01-30T10:43:31-05:00"
// }
// }
// ```
// Here’s an example event for the `subscription_state_change` event:
// ```
// {
// "event": {
// "id": 353,
// "key": "subscription_state_change",
// "message": "State changed on Marky Mark's subscription to Pro from trialing to active",
// "subscription_id": 205,
// "event_specific_data": {
// "new_subscription_state": "active",
// "previous_subscription_state": "trialing"
// },
// "created_at": "2012-01-30T10:43:33-05:00"
// }
// }
// ```
func (e *EventsController) ListEvents(
    ctx context.Context,
    input ListEventsInput) (
    models.ApiResponse[[]models.EventResponse],
    error) {
    req := e.prepareRequest(ctx, "GET", "/events.json")
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.SinceId != nil {
        req.QueryParam("since_id", *input.SinceId)
    }
    if input.MaxId != nil {
        req.QueryParam("max_id", *input.MaxId)
    }
    if input.Direction != nil {
        req.QueryParam("direction", *input.Direction)
    }
    if input.Filter != nil {
        req.QueryParam("filter", input.Filter)
    }
    if input.DateField != nil {
        req.QueryParam("date_field", *input.DateField)
    }
    if input.StartDate != nil {
        req.QueryParam("start_date", *input.StartDate)
    }
    if input.EndDate != nil {
        req.QueryParam("end_date", *input.EndDate)
    }
    if input.StartDatetime != nil {
        req.QueryParam("start_datetime", *input.StartDatetime)
    }
    if input.EndDatetime != nil {
        req.QueryParam("end_datetime", *input.EndDatetime)
    }
    var result []models.EventResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.EventResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSubscriptionEventsInput represents the input of the ListSubscriptionEvents endpoint.
type ListSubscriptionEventsInput struct {
    // The Chargify id of the subscription
    SubscriptionId int                
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page           *int               
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage        *int               
    // Returns events with an id greater than or equal to the one specified
    SinceId        *int64             
    // Returns events with an id less than or equal to the one specified
    MaxId          *int64             
    // The sort direction of the returned events.
    Direction      *models.Direction  
    // You can pass multiple event keys after comma.
    // Use in query `filter=signup_success,payment_success`.
    Filter         []models.EventType 
}

// ListSubscriptionEvents takes context, subscriptionId, page, perPage, sinceId, maxId, direction, filter as parameters and
// returns an models.ApiResponse with []models.EventResponse data and
// an error if there was an issue with the request or response.
// The following request will return a list of events for a subscription.
// Each event type has its own `event_specific_data` specified.
func (e *EventsController) ListSubscriptionEvents(
    ctx context.Context,
    input ListSubscriptionEventsInput) (
    models.ApiResponse[[]models.EventResponse],
    error) {
    req := e.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/subscriptions/%v/events.json", input.SubscriptionId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.SinceId != nil {
        req.QueryParam("since_id", *input.SinceId)
    }
    if input.MaxId != nil {
        req.QueryParam("max_id", *input.MaxId)
    }
    if input.Direction != nil {
        req.QueryParam("direction", *input.Direction)
    }
    if input.Filter != nil {
        req.QueryParam("filter", input.Filter)
    }
    
    var result []models.EventResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.EventResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadEventsCountInput represents the input of the ReadEventsCount endpoint.
type ReadEventsCountInput struct {
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page      *int               
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage   *int               
    // Returns events with an id greater than or equal to the one specified
    SinceId   *int64             
    // Returns events with an id less than or equal to the one specified
    MaxId     *int64             
    // The sort direction of the returned events.
    Direction *models.Direction  
    // You can pass multiple event keys after comma.
    // Use in query `filter=signup_success,payment_success`.
    Filter    []models.EventType 
}

// ReadEventsCount takes context, page, perPage, sinceId, maxId, direction, filter as parameters and
// returns an models.ApiResponse with models.CountResponse data and
// an error if there was an issue with the request or response.
// Get a count of all the events for a given site by using this method.
func (e *EventsController) ReadEventsCount(
    ctx context.Context,
    input ReadEventsCountInput) (
    models.ApiResponse[models.CountResponse],
    error) {
    req := e.prepareRequest(ctx, "GET", "/events/count.json")
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.SinceId != nil {
        req.QueryParam("since_id", *input.SinceId)
    }
    if input.MaxId != nil {
        req.QueryParam("max_id", *input.MaxId)
    }
    if input.Direction != nil {
        req.QueryParam("direction", *input.Direction)
    }
    if input.Filter != nil {
        req.QueryParam("filter", input.Filter)
    }
    var result models.CountResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CountResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
