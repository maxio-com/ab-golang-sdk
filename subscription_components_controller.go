// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
	"context"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
	"net/http"
	"time"
)

// SubscriptionComponentsController represents a controller struct.
type SubscriptionComponentsController struct {
	baseController
}

// NewSubscriptionComponentsController creates a new instance of SubscriptionComponentsController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionComponentsController.
func NewSubscriptionComponentsController(baseController baseController) *SubscriptionComponentsController {
	subscriptionComponentsController := SubscriptionComponentsController{baseController: baseController}
	return &subscriptionComponentsController
}

// ReadSubscriptionComponent takes context, subscriptionId, componentId as parameters and
// returns an models.ApiResponse with models.SubscriptionComponentResponse data and
// an error if there was an issue with the request or response.
// This request will list information regarding a specific component owned by a subscription.
func (s *SubscriptionComponentsController) ReadSubscriptionComponent(
	ctx context.Context,
	subscriptionId int,
	componentId int) (
	models.ApiResponse[models.SubscriptionComponentResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/subscriptions/%v/components/%v.json")
	req.AppendTemplateParams(subscriptionId, componentId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})

	var result models.SubscriptionComponentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SubscriptionComponentResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSubscriptionComponentsInput represents the input of the ListSubscriptionComponents endpoint.
type ListSubscriptionComponentsInput struct {
	// The Chargify id of the subscription
	SubscriptionId int
	// The type of filter you'd like to apply to your search. Use in query `date_field=updated_at`.
	DateField *models.SubscriptionListDateField
	// Controls the order in which results are returned.
	// Use in query `direction=asc`.
	Direction *models.SortingDirection
	// Filter to use for List Subscription Components operation
	Filter *models.ListSubscriptionComponentsFilter
	// The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
	EndDate *string
	// The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of end_date.
	EndDatetime *string
	// Allows fetching components allocation only if price point id is present. Use in query `price_point_ids=not_null`.
	PricePointIds *models.IncludeNotNull
	// Allows fetching components allocation with matching product family id based on provided ids. Use in query `product_family_ids=1,2,3`.
	ProductFamilyIds []int
	// The attribute by which to sort. Use in query `sort=updated_at`.
	Sort *models.ListSubscriptionComponentsSort
	// The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
	StartDate *string
	// The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of start_date.
	StartDatetime *string
	// Allows including additional data in the response. Use in query `include=subscription,historic_usages`.
	Include []models.ListSubscriptionComponentsInclude
	// If in_use is set to true, it returns only components that are currently in use. However, if it's set to false or not provided, it returns all components connected with the subscription.
	InUse *bool
}

// ListSubscriptionComponents takes context, subscriptionId, dateField, direction, filter, endDate, endDatetime, pricePointIds, productFamilyIds, sort, startDate, startDatetime, include, inUse as parameters and
// returns an models.ApiResponse with []models.SubscriptionComponentResponse data and
// an error if there was an issue with the request or response.
// This request will list a subscription's applied components.
// ## Archived Components
// When requesting to list components for a given subscription, if the subscription contains **archived** components they will be listed in the server response.
func (s *SubscriptionComponentsController) ListSubscriptionComponents(
	ctx context.Context,
	input ListSubscriptionComponentsInput) (
	models.ApiResponse[[]models.SubscriptionComponentResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/subscriptions/%v/components.json")
	req.AppendTemplateParams(input.SubscriptionId)
	req.Authenticate(NewAuth("BasicAuth"))
	if input.DateField != nil {
		req.QueryParam("date_field", *input.DateField)
	}
	if input.Direction != nil {
		req.QueryParam("direction", *input.Direction)
	}
	if input.Filter != nil {
		req.QueryParam("filter", *input.Filter)
	}
	if input.EndDate != nil {
		req.QueryParam("end_date", *input.EndDate)
	}
	if input.EndDatetime != nil {
		req.QueryParam("end_datetime", *input.EndDatetime)
	}
	if input.PricePointIds != nil {
		req.QueryParam("price_point_ids", *input.PricePointIds)
	}
	if input.ProductFamilyIds != nil {
		req.QueryParam("product_family_ids", input.ProductFamilyIds)
	}
	if input.Sort != nil {
		req.QueryParam("sort", *input.Sort)
	}
	if input.StartDate != nil {
		req.QueryParam("start_date", *input.StartDate)
	}
	if input.StartDatetime != nil {
		req.QueryParam("start_datetime", *input.StartDatetime)
	}
	if input.Include != nil {
		req.QueryParam("include", input.Include)
	}
	if input.InUse != nil {
		req.QueryParam("in_use", *input.InUse)
	}

	var result []models.SubscriptionComponentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.SubscriptionComponentResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// BulkUpdateSubscriptionComponentsPricePoints takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.BulkComponentsPricePointAssignment data and
// an error if there was an issue with the request or response.
// Updates the price points on one or more of a subscription's components.
// The `price_point` key can take either a:
// 1. Price point id (integer)
// 2. Price point handle (string)
// 3. `"_default"` string, which will reset the price point to the component's current default price point.
func (s *SubscriptionComponentsController) BulkUpdateSubscriptionComponentsPricePoints(
	ctx context.Context,
	subscriptionId int,
	body *models.BulkComponentsPricePointAssignment) (
	models.ApiResponse[models.BulkComponentsPricePointAssignment],
	error) {
	req := s.prepareRequest(ctx, "POST", "/subscriptions/%v/price_points.json")
	req.AppendTemplateParams(subscriptionId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewComponentPricePointError},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.BulkComponentsPricePointAssignment
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BulkComponentsPricePointAssignment](decoder)
	return models.NewApiResponse(result, resp), err
}

// BulkResetSubscriptionComponentsPricePoints takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Resets all of a subscription's components to use the current default.
// **Note**: this will update the price point for all of the subscription's components, even ones that have not been allocated yet.
func (s *SubscriptionComponentsController) BulkResetSubscriptionComponentsPricePoints(
	ctx context.Context,
	subscriptionId int) (
	models.ApiResponse[models.SubscriptionResponse],
	error) {
	req := s.prepareRequest(ctx, "POST", "/subscriptions/%v/price_points/reset.json")
	req.AppendTemplateParams(subscriptionId)
	req.Authenticate(NewAuth("BasicAuth"))

	var result models.SubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// AllocateComponent takes context, subscriptionId, componentId, body as parameters and
// returns an models.ApiResponse with models.AllocationResponse data and
// an error if there was an issue with the request or response.
// This endpoint creates a new allocation, setting the current allocated quantity for the Component and recording a memo.
// **Notice**: Allocations can only be updated for Quantity, On/Off, and Prepaid Components.
// ## Allocations Documentation
// Full documentation on how to record Allocations in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24251883961485-Component-Allocations-Overview). It is focused on how allocations operate within the Advanced Billing UI.It goes into greater detail on how the user interface will react when recording allocations.
// This documentation also goes into greater detail on how proration is taken into consideration when applying component allocations.
// ## Proration Schemes
// Changing the allocated quantity of a component mid-period can result in either a Charge or Credit being applied to the subscription. When creating an allocation via the API, you can pass the `upgrade_charge`, `downgrade_credit`, and `accrue_charge` to be applied.
// **Notice:** These proration and accural fields will be ignored for Prepaid Components since this component type always generate charges immediately without proration.
// For background information on prorated components and upgrade/downgrade schemes, see [Setting Component Allocations.](https://maxio.zendesk.com/hc/en-us/articles/24251906165133-Component-Allocations-Proration).
// See the tables below for valid values.
// | upgrade_charge | Definition                                                        |
// |----------------|-------------------------------------------------------------------|
// | `full`         | A charge is added for the full price of the component.            |
// | `prorated`     | A charge is added for the prorated price of the component change. |
// | `none`         | No charge is added.                                               |
// | downgrade_credit | Definition                                        |
// |------------------|---------------------------------------------------|
// | `full`           | A full price credit is added for the amount owed. |
// | `prorated`       | A prorated credit is added for the amount owed.   |
// | `none`           | No charge is added.                               |
// | accrue_charge | Definition                                                                                                 |
// |---------------|------------------------------------------------------------------------------------------------------------|
// | `true`        | Attempt to charge the customer at next renewal.                                                            |
// | `false`       | Attempt to charge the customer right away. If it fails, the charge will be accrued until the next renewal. |
// ### Order of Resolution for upgrade_charge and downgrade_credit
// 1. Per allocation in API call (within a single allocation of the `allocations` array)
// 2. [Component-level default value](https://maxio.zendesk.com/hc/en-us/articles/24251883961485-Component-Allocations-Overview)
// 3. Allocation API call top level (outside of the `allocations` array)
// 4. [Site-level default value](https://maxio.zendesk.com/hc/en-us/articles/24251906165133-Component-Allocations-Proration#proration-schemes)
// ### Order of Resolution for accrue charge
// 1. Allocation API call top level (outside of the `allocations` array)
// 2. [Site-level default value](https://maxio.zendesk.com/hc/en-us/articles/24251906165133-Component-Allocations-Proration#proration-schemes)
// **NOTE: Proration uses the current price of the component as well as the current tax rates. Changes to either may cause the prorated charge/credit to be wrong.**
func (s *SubscriptionComponentsController) AllocateComponent(
	ctx context.Context,
	subscriptionId int,
	componentId int,
	body *models.CreateAllocationRequest) (
	models.ApiResponse[models.AllocationResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/subscriptions/%v/components/%v/allocations.json",
	)
	req.AppendTemplateParams(subscriptionId, componentId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.AllocationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AllocationResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListAllocations takes context, subscriptionId, componentId, page as parameters and
// returns an models.ApiResponse with []models.AllocationResponse data and
// an error if there was an issue with the request or response.
// This endpoint returns the 50 most recent Allocations, ordered by most recent first.
// ## On/Off Components
// When a subscription's on/off component has been toggled to on (`1`) or off (`0`), usage will be logged in this response.
// ## Querying data via Advanced Billing gem
// You can also query the current quantity via the [official Advanced Billing Gem.](http://github.com/chargify/chargify_api_ares)
// ```# First way
// component = Chargify::Subscription::Component.find(1, :params => {:subscription_id => 7})
// puts component.allocated_quantity
// # => 23
// # Second way
// component = Chargify::Subscription.find(7).component(1)
// puts component.allocated_quantity
// # => 23
// ```
func (s *SubscriptionComponentsController) ListAllocations(
	ctx context.Context,
	subscriptionId int,
	componentId int,
	page *int) (
	models.ApiResponse[[]models.AllocationResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		"/subscriptions/%v/components/%v/allocations.json",
	)
	req.AppendTemplateParams(subscriptionId, componentId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.AllocationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.AllocationResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// AllocateComponents takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with []models.AllocationResponse data and
// an error if there was an issue with the request or response.
// Creates multiple allocations, setting the current allocated quantity for each of the components and recording a memo. The charges and/or credits that are created will be rolled up into a single total which is used to determine whether this is an upgrade or a downgrade. Be aware of the Order of Resolutions explained below in determining the proration scheme.
// A `component_id` is required for each allocation.
// This endpoint only responds to JSON. It is not available for XML.
func (s *SubscriptionComponentsController) AllocateComponents(
	ctx context.Context,
	subscriptionId int,
	body *models.AllocateComponents) (
	models.ApiResponse[[]models.AllocationResponse],
	error) {
	req := s.prepareRequest(ctx, "POST", "/subscriptions/%v/allocations.json")
	req.AppendTemplateParams(subscriptionId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result []models.AllocationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.AllocationResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// PreviewAllocations takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.AllocationPreviewResponse data and
// an error if there was an issue with the request or response.
// Advanced Billing offers the ability to preview a potential subscription's **quantity-based** or **on/off** component allocation in the middle of the current billing period.  This is useful if you want users to be able to see the effect of a component operation before actually doing it.
// ## Fine-grained Component Control: Use with multiple `upgrade_charge`s or `downgrade_credits`
// When the allocation uses multiple different types of `upgrade_charge`s or `downgrade_credit`s, the Allocation is viewed as an Allocation which uses "Fine-Grained Component Control". As a result, the response will not include `direction` and `proration` within the `allocation_preview`, but at the `line_items` and `allocations` level respectfully.
// See example below for Fine-Grained Component Control response.
func (s *SubscriptionComponentsController) PreviewAllocations(
	ctx context.Context,
	subscriptionId int,
	body *models.PreviewAllocationsRequest) (
	models.ApiResponse[models.AllocationPreviewResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/subscriptions/%v/allocations/preview.json",
	)
	req.AppendTemplateParams(subscriptionId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewComponentAllocationError},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.AllocationPreviewResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AllocationPreviewResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdatePrepaidUsageAllocationExpirationDate takes context, subscriptionId, componentId, allocationId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// When the expiration interval options are selected on a prepaid usage component price point, all allocations will be created with an expiration date. This expiration date can be changed after the fact to allow for extending or shortening the allocation's active window.
// In order to change a prepaid usage allocation's expiration date, a PUT call must be made to the allocation's endpoint with a new expiration date.
// ## Limitations
// A few limitations exist when changing an allocation's expiration date:
// - An expiration date can only be changed for an allocation that belongs to a price point with expiration interval options explicitly set.
// - An expiration date can be changed towards the future with no limitations.
// - An expiration date can be changed towards the past (essentially expiring it) up to the subscription's current period beginning date.
func (s *SubscriptionComponentsController) UpdatePrepaidUsageAllocationExpirationDate(
	ctx context.Context,
	subscriptionId int,
	componentId int,
	allocationId int,
	body *models.UpdateAllocationExpirationDate) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"PUT",
		"/subscriptions/%v/components/%v/allocations/%v.json",
	)
	req.AppendTemplateParams(subscriptionId, componentId, allocationId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSubscriptionComponentAllocationError},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// DeletePrepaidUsageAllocation takes context, subscriptionId, componentId, allocationId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Prepaid Usage components are unique in that their allocations are always additive. In order to reduce a subscription's allocated quantity for a prepaid usage component each allocation must be destroyed individually via this endpoint.
// ## Credit Scheme
// By default, destroying an allocation will generate a service credit on the subscription. This behavior can be modified with the optional `credit_scheme` parameter on this endpoint. The accepted values are:
// 1. `none`: The allocation will be destroyed and the balances will be updated but no service credit or refund will be created.
// 2. `credit`: The allocation will be destroyed and the balances will be updated and a service credit will be generated. This is also the default behavior if the `credit_scheme` param is not passed.
// 3. `refund`: The allocation will be destroyed and the balances will be updated and a refund will be issued along with a Credit Note.
func (s *SubscriptionComponentsController) DeletePrepaidUsageAllocation(
	ctx context.Context,
	subscriptionId int,
	componentId int,
	allocationId int,
	body *models.CreditSchemeRequest) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		"/subscriptions/%v/components/%v/allocations/%v.json",
	)
	req.AppendTemplateParams(subscriptionId, componentId, allocationId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSubscriptionComponentAllocationError},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// CreateUsage takes context, subscriptionId, componentId, body as parameters and
// returns an models.ApiResponse with models.UsageResponse data and
// an error if there was an issue with the request or response.
// ## Documentation
// Full documentation on how to create Components in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24261149711501-Create-Edit-and-Archive-Components). Additionally, for information on how to record component usage against a subscription, please see the following resources:
// + [Recording Metered Component Usage](https://maxio.zendesk.com/hc/en-us/articles/24251890500109-Reporting-Component-Allocations#reporting-metered-component-usage)
// + [Reporting Prepaid Component Status](https://maxio.zendesk.com/hc/en-us/articles/24251890500109-Reporting-Component-Allocations#reporting-prepaid-component-status)
// You may choose to report metered or prepaid usage to Advanced Billing as often as you wish. You may report usage as it happens. You may also report usage periodically, such as each night or once per billing period. If usage events occur in your system very frequently (on the order of thousands of times an hour), it is best to accumulate usage into batches on your side, and then report those batches less frequently, such as daily. This will ensure you remain below any API throttling limits. If your use case requires higher rates of usage reporting, we recommend utilizing Events Based Components.
// ## Create Usage for Subscription
// This endpoint allows you to record an instance of metered or prepaid usage for a subscription. The `quantity` from usage for each component is accumulated to the `unit_balance` on the [Component Line Item](./b3A6MTQxMDgzNzQ-read-subscription-component) for the subscription.
// ## Price Point ID usage
// If you are using price points, for metered and prepaid usage components, Advanced Billing gives you the option to specify a price point in your request.
// You do not need to specify a price point ID. If a price point is not included, the default price point for the component will be used when the usage is recorded.
// If an invalid `price_point_id` is submitted, the endpoint will return an error.
// ## Deducting Usage
// In the event that you need to reverse a previous usage report or otherwise deduct from the current usage balance, you may provide a negative quantity.
// Example:
// Previously recorded:
// ```json
// {
// "usage": {
// "quantity": 5000,
// "memo": "Recording 5000 units"
// }
// }
// ```
// At this point, `unit_balance` would be `5000`. To reduce the balance to `0`, POST the following payload:
// ```json
// {
// "usage": {
// "quantity": -5000,
// "memo": "Deducting 5000 units"
// }
// }
// ```
// The `unit_balance` has a floor of `0`; negative unit balances are never allowed. For example, if the usage balance is 100 and you deduct 200 units, the unit balance would then be `0`, not `-100`.
// ## FAQ
// Q. Is it possible to record metered usage for more than one component at a time?
// A. No. Usage should be reported as one API call per component on a single subscription. For example, to record that a subscriber has sent both an SMS Message and an Email, send an API call for each.
func (s *SubscriptionComponentsController) CreateUsage(
	ctx context.Context,
	subscriptionId int,
	componentId models.CreateUsageComponentId,
	body *models.CreateUsageRequest) (
	models.ApiResponse[models.UsageResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/subscriptions/%v/components/%v/usages.json",
	)
	req.AppendTemplateParams(subscriptionId, componentId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.UsageResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UsageResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListUsagesInput represents the input of the ListUsages endpoint.
type ListUsagesInput struct {
	// The Chargify id of the subscription
	SubscriptionId int
	// Either the Advanced Billing id for the component or the component's handle prefixed by `handle:`
	ComponentId models.ListUsagesInputComponentId
	// Returns usages with an id greater than or equal to the one specified
	SinceId *int64
	// Returns usages with an id less than or equal to the one specified
	MaxId *int64
	// Returns usages with a created_at date greater than or equal to midnight (12:00 AM) on the date specified.
	SinceDate *time.Time
	// Returns usages with a created_at date less than or equal to midnight (12:00 AM) on the date specified.
	UntilDate *time.Time
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
}

// ListUsages takes context, subscriptionId, componentId, sinceId, maxId, sinceDate, untilDate, page, perPage as parameters and
// returns an models.ApiResponse with []models.UsageResponse data and
// an error if there was an issue with the request or response.
// This request will return a list of the usages associated with a subscription for a particular metered component. This will display the previously recorded components for a subscription.
// This endpoint is not compatible with quantity-based components.
// ## Since Date and Until Date Usage
// Note: The `since_date` and `until_date` attributes each default to midnight on the date specified. For example, in order to list usages for January 20th, you would need to append the following to the URL.
// ```
// ?since_date=2016-01-20&until_date=2016-01-21
// ```
// ## Read Usage by Handle
// Use this endpoint to read the previously recorded components for a subscription.  You can now specify either the component id (integer) or the component handle prefixed by "handle:" to specify the unique identifier for the component you are working with.
func (s *SubscriptionComponentsController) ListUsages(
	ctx context.Context,
	input ListUsagesInput) (
	models.ApiResponse[[]models.UsageResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		"/subscriptions/%v/components/%v/usages.json",
	)
	req.AppendTemplateParams(input.SubscriptionId, input.ComponentId)
	req.Authenticate(NewAuth("BasicAuth"))
	if input.SinceId != nil {
		req.QueryParam("since_id", *input.SinceId)
	}
	if input.MaxId != nil {
		req.QueryParam("max_id", *input.MaxId)
	}
	if input.SinceDate != nil {
		req.QueryParam("since_date", input.SinceDate.Format(models.DEFAULT_DATE))
	}
	if input.UntilDate != nil {
		req.QueryParam("until_date", input.UntilDate.Format(models.DEFAULT_DATE))
	}
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}

	var result []models.UsageResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.UsageResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ActivateEventBasedComponent takes context, subscriptionId, componentId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// In order to bill your subscribers on your Events data under the Events-Based Billing feature, the components must be activated for the subscriber.
// Learn more about the role of activation in the [Events-Based Billing docs](https://maxio.zendesk.com/hc/en-us/articles/24260323329805-Events-Based-Billing-Overview).
// Use this endpoint to activate an event-based component for a single subscription. Activating an event-based component causes Advanced Billing to bill for events when the subscription is renewed.
// *Note: it is possible to stream events for a subscription at any time, regardless of component activation status. The activation status only determines if the subscription should be billed for event-based component usage at renewal.*
func (s *SubscriptionComponentsController) ActivateEventBasedComponent(
	ctx context.Context,
	subscriptionId int,
	componentId int,
	body *models.ActivateEventBasedComponent) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/event_based_billing/subscriptions/%v/components/%v/activate.json",
	)
	req.AppendTemplateParams(subscriptionId, componentId)
	req.Authenticate(NewAuth("BasicAuth"))
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// DeactivateEventBasedComponent takes context, subscriptionId, componentId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Use this endpoint to deactivate an event-based component for a single subscription. Deactivating the event-based component causes Advanced Billing to ignore related events at subscription renewal.
func (s *SubscriptionComponentsController) DeactivateEventBasedComponent(
	ctx context.Context,
	subscriptionId int,
	componentId int) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/event_based_billing/subscriptions/%v/components/%v/deactivate.json",
	)
	req.AppendTemplateParams(subscriptionId, componentId)
	req.Authenticate(NewAuth("BasicAuth"))

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// RecordEvent takes context, apiHandle, storeUid, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// ## Documentation
// Events-Based Billing is an evolved form of metered billing that is based on data-rich events streamed in real-time from your system to Advanced Billing.
// These events can then be transformed, enriched, or analyzed to form the computed totals of usage charges billed to your customers.
// This API allows you to stream events into the Advanced Billing data ingestion engine.
// Learn more about the feature in general in the [Events-Based Billing help docs](https://maxio.zendesk.com/hc/en-us/articles/24260323329805-Events-Based-Billing-Overview).
// ## Record Event
// Use this endpoint to record a single event.
// *Note: this endpoint differs from the standard Chargify API endpoints in that the URL subdomain will be `events` and your site subdomain will be included in the URL path. For example:*
// ```
// https://events.chargify.com/my-site-subdomain/events/my-stream-api-handle
// ```
func (s *SubscriptionComponentsController) RecordEvent(
	ctx context.Context,
	apiHandle string,
	storeUid *string,
	body *models.EBBEvent) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/events/%v.json")
	req.AppendTemplateParams(apiHandle)
	req.BaseUrl("ebb")
	req.Authenticate(NewAuth("BasicAuth"))
	req.Header("Content-Type", "application/json")
	if storeUid != nil {
		req.QueryParam("store_uid", *storeUid)
	}
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// BulkRecordEvents takes context, apiHandle, storeUid, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Use this endpoint to record a collection of events.
// *Note: this endpoint differs from the standard Chargify API endpoints in that the subdomain will be `events` and your site subdomain will be included in the URL path.*
// A maximum of 1000 events can be published in a single request. A 422 will be returned if this limit is exceeded.
func (s *SubscriptionComponentsController) BulkRecordEvents(
	ctx context.Context,
	apiHandle string,
	storeUid *string,
	body []models.EBBEvent) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/events/%v/bulk.json")
	req.AppendTemplateParams(apiHandle)
	req.BaseUrl("ebb")
	req.Authenticate(NewAuth("BasicAuth"))
	req.Header("Content-Type", "application/json")
	if storeUid != nil {
		req.QueryParam("store_uid", *storeUid)
	}
	if body != nil {
		req.Json(body)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// ListSubscriptionComponentsForSiteInput represents the input of the ListSubscriptionComponentsForSite endpoint.
type ListSubscriptionComponentsForSiteInput struct {
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
	// This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
	// Use in query `per_page=200`.
	PerPage *int
	// The attribute by which to sort. Use in query: `sort=updated_at`.
	Sort *models.ListSubscriptionComponentsSort
	// Controls the order in which results are returned.
	// Use in query `direction=asc`.
	Direction *models.SortingDirection
	// Filter to use for List Subscription Components For Site operation
	Filter *models.ListSubscriptionComponentsForSiteFilter
	// The type of filter you'd like to apply to your search. Use in query: `date_field=updated_at`.
	DateField *models.SubscriptionListDateField
	// The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. Use in query `start_date=2011-12-15`.
	StartDate *string
	// The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of start_date. Use in query `start_datetime=2022-07-01 09:00:05`.
	StartDatetime *string
	// The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. Use in query `end_date=2011-12-16`.
	EndDate *string
	// The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of end_date. Use in query `end_datetime=2022-07-01 09:00:05`.
	EndDatetime *string
	// Allows fetching components allocation with matching subscription id based on provided ids. Use in query `subscription_ids=1,2,3`.
	SubscriptionIds []int
	// Allows fetching components allocation only if price point id is present. Use in query `price_point_ids=not_null`.
	PricePointIds *models.IncludeNotNull
	// Allows fetching components allocation with matching product family id based on provided ids. Use in query `product_family_ids=1,2,3`.
	ProductFamilyIds []int
	// Allows including additional data in the response. Use in query `include=subscription,historic_usages`.
	Include *models.ListSubscriptionComponentsInclude
}

// ListSubscriptionComponentsForSite takes context, page, perPage, sort, direction, filter, dateField, startDate, startDatetime, endDate, endDatetime, subscriptionIds, pricePointIds, productFamilyIds, include as parameters and
// returns an models.ApiResponse with models.ListSubscriptionComponentsResponse data and
// an error if there was an issue with the request or response.
// This request will list components applied to each subscription.
func (s *SubscriptionComponentsController) ListSubscriptionComponentsForSite(
	ctx context.Context,
	input ListSubscriptionComponentsForSiteInput) (
	models.ApiResponse[models.ListSubscriptionComponentsResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/subscriptions_components.json")

	req.Authenticate(NewAuth("BasicAuth"))
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.Sort != nil {
		req.QueryParam("sort", *input.Sort)
	}
	if input.Direction != nil {
		req.QueryParam("direction", *input.Direction)
	}
	if input.Filter != nil {
		req.QueryParam("filter", *input.Filter)
	}
	if input.DateField != nil {
		req.QueryParam("date_field", *input.DateField)
	}
	if input.StartDate != nil {
		req.QueryParam("start_date", *input.StartDate)
	}
	if input.StartDatetime != nil {
		req.QueryParam("start_datetime", *input.StartDatetime)
	}
	if input.EndDate != nil {
		req.QueryParam("end_date", *input.EndDate)
	}
	if input.EndDatetime != nil {
		req.QueryParam("end_datetime", *input.EndDatetime)
	}
	if input.SubscriptionIds != nil {
		req.QueryParam("subscription_ids", input.SubscriptionIds)
	}
	if input.PricePointIds != nil {
		req.QueryParam("price_point_ids", *input.PricePointIds)
	}
	if input.ProductFamilyIds != nil {
		req.QueryParam("product_family_ids", input.ProductFamilyIds)
	}
	if input.Include != nil {
		req.QueryParam("include", *input.Include)
	}
	var result models.ListSubscriptionComponentsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListSubscriptionComponentsResponse](decoder)
	return models.NewApiResponse(result, resp), err
}
