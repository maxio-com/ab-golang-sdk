package ab_golang_sdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"maxioadvancedbilling/models"
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

// ListEvents takes context, page, perPage, sinceId, maxId, direction, filter, dateField, startDate, endDate, startDatetime, endDatetime as parameters and
// returns an models.ApiResponse with []models.EventResponse data and
// an error if there was an issue with the request or response.
// ## Events Intro
// Chargify Events include various activity that happens around a Site. This information is **especially** useful to track down issues that arise when subscriptions are not created due to errors.
// Within the Chargify UI, "Events" are referred to as "Site Activity".  Full documentation on how to record view Events / Site Activty in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407864698139).
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
	page *int,
	perPage *int,
	sinceId *int,
	maxId *int,
	direction *models.DirectionEnum,
	filter []models.EventTypeEnum,
	dateField *models.ListEventsDateFieldEnum,
	startDate *string,
	endDate *string,
	startDatetime *string,
	endDatetime *string) (
	models.ApiResponse[[]models.EventResponse],
	error) {
	req := e.prepareRequest(ctx, "GET", "/events.json")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if sinceId != nil {
		req.QueryParam("since_id", *sinceId)
	}
	if maxId != nil {
		req.QueryParam("max_id", *maxId)
	}
	if direction != nil {
		req.QueryParam("direction", *direction)
	}
	if filter != nil {
		req.QueryParam("filter", filter)
	}
	if dateField != nil {
		req.QueryParam("date_field", *dateField)
	}
	if startDate != nil {
		req.QueryParam("start_date", *startDate)
	}
	if endDate != nil {
		req.QueryParam("end_date", *endDate)
	}
	if startDatetime != nil {
		req.QueryParam("start_datetime", *startDatetime)
	}
	if endDatetime != nil {
		req.QueryParam("end_datetime", *endDatetime)
	}
	var result []models.EventResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.EventResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// ListSubscriptionEvents takes context, subscriptionId, page, perPage, sinceId, maxId, direction, filter as parameters and
// returns an models.ApiResponse with []models.EventResponse data and
// an error if there was an issue with the request or response.
// The following request will return a list of events for a subscription.
// Each event type has its own `event_specific_data` specified.
func (e *EventsController) ListSubscriptionEvents(
	ctx context.Context,
	subscriptionId int,
	page *int,
	perPage *int,
	sinceId *int,
	maxId *int,
	direction *models.DirectionEnum,
	filter []models.EventTypeEnum) (
	models.ApiResponse[[]models.EventResponse],
	error) {
	req := e.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/events.json", subscriptionId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if sinceId != nil {
		req.QueryParam("since_id", *sinceId)
	}
	if maxId != nil {
		req.QueryParam("max_id", *maxId)
	}
	if direction != nil {
		req.QueryParam("direction", *direction)
	}
	if filter != nil {
		req.QueryParam("filter", filter)
	}

	var result []models.EventResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.EventResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// ReadEventsCount takes context, page, perPage, sinceId, maxId, direction, filter as parameters and
// returns an models.ApiResponse with models.CountResponse data and
// an error if there was an issue with the request or response.
// Get a count of all the events for a given site by using this method.
func (e *EventsController) ReadEventsCount(
	ctx context.Context,
	page *int,
	perPage *int,
	sinceId *int,
	maxId *int,
	direction *models.DirectionEnum,
	filter []models.EventTypeEnum) (
	models.ApiResponse[models.CountResponse],
	error) {
	req := e.prepareRequest(ctx, "GET", "/events/count.json")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if sinceId != nil {
		req.QueryParam("since_id", *sinceId)
	}
	if maxId != nil {
		req.QueryParam("max_id", *maxId)
	}
	if direction != nil {
		req.QueryParam("direction", *direction)
	}
	if filter != nil {
		req.QueryParam("filter", filter)
	}
	var result models.CountResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.CountResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
