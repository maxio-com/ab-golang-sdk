package ab_golang_sdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"maxioadvancedbilling/errors"
	"maxioadvancedbilling/models"
	"net/http"
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
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/components/%v.json", subscriptionId, componentId),
	)
	req.Authenticate(true)

	var result models.SubscriptionComponentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SubscriptionComponentResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}

// ListSubscriptionComponents takes context, subscriptionId, dateField, direction, endDate, endDatetime, pricePointIds, productFamilyIds, sort, startDate, startDatetime, include, filterUseSiteExchangeRate, filterCurrencies as parameters and
// returns an models.ApiResponse with []models.SubscriptionComponentResponse data and
// an error if there was an issue with the request or response.
// This request will list a subscription's applied components.
// ## Archived Components
// When requesting to list components for a given subscription, if the subscription contains **archived** components they will be listed in the server response.
func (s *SubscriptionComponentsController) ListSubscriptionComponents(
	ctx context.Context,
	subscriptionId int,
	dateField *models.SubscriptionListDateFieldEnum,
	direction *models.SortingDirectionEnum,
	endDate *string,
	endDatetime *string,
	pricePointIds *models.IncludeNotNullEnum,
	productFamilyIds []int,
	sort *models.ListSubscriptionComponentsSortEnum,
	startDate *string,
	startDatetime *string,
	include *models.ListSubscriptionComponentsIncludeEnum,
	filterUseSiteExchangeRate *bool,
	filterCurrencies []string) (
	models.ApiResponse[[]models.SubscriptionComponentResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/components.json", subscriptionId),
	)
	req.Authenticate(true)
	if dateField != nil {
		req.QueryParam("date_field", *dateField)
	}
	if direction != nil {
		req.QueryParam("direction", *direction)
	}
	if endDate != nil {
		req.QueryParam("end_date", *endDate)
	}
	if endDatetime != nil {
		req.QueryParam("end_datetime", *endDatetime)
	}
	if pricePointIds != nil {
		req.QueryParam("price_point_ids", *pricePointIds)
	}
	if productFamilyIds != nil {
		req.QueryParam("product_family_ids", productFamilyIds)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if startDate != nil {
		req.QueryParam("start_date", *startDate)
	}
	if startDatetime != nil {
		req.QueryParam("start_datetime", *startDatetime)
	}
	if include != nil {
		req.QueryParam("include", *include)
	}
	if filterUseSiteExchangeRate != nil {
		req.QueryParam("filter[use_site_exchange_rate]", *filterUseSiteExchangeRate)
	}
	if filterCurrencies != nil {
		req.QueryParam("filter[currencies]", filterCurrencies)
	}

	var result []models.SubscriptionComponentResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.SubscriptionComponentResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionComponentsPricePoints takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.BulkComponentSPricePointAssignment data and
// an error if there was an issue with the request or response.
// Updates the price points on one or more of a subscription's components.
// The `price_point` key can take either a:
// 1. Price point id (integer)
// 2. Price point handle (string)
// 3. `"_default"` string, which will reset the price point to the component's current default price point.
func (s *SubscriptionComponentsController) UpdateSubscriptionComponentsPricePoints(
	ctx context.Context,
	subscriptionId int,
	body *models.BulkComponentSPricePointAssignment) (
	models.ApiResponse[models.BulkComponentSPricePointAssignment],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/price_points.json", subscriptionId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result models.BulkComponentSPricePointAssignment
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BulkComponentSPricePointAssignment](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 422 {
		err = errors.NewComponentPricePointError(422, "Unprocessable Entity (WebDAV)")
	}
	return models.NewApiResponse(result, resp), err
}

// ResetSubscriptionComponentsPricePoints takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Resets all of a subscription's components to use the current default.
// **Note**: this will update the price point for all of the subscription's components, even ones that have not been allocated yet.
func (s *SubscriptionComponentsController) ResetSubscriptionComponentsPricePoints(
	ctx context.Context,
	subscriptionId int) (
	models.ApiResponse[models.SubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/price_points/reset.json", subscriptionId),
	)
	req.Authenticate(true)

	var result models.SubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// AllocateComponent takes context, subscriptionId, componentId, body as parameters and
// returns an models.ApiResponse with models.AllocationResponse data and
// an error if there was an issue with the request or response.
// This endpoint creates a new allocation, setting the current allocated quantity for the Component and recording a memo.
// **Notice**: Allocations can only be updated for Quantity, On/Off, and Prepaid Components.
// ## Allocations Documentation
// Full documentation on how to record Allocations in the Chargify UI can be located [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404527849997). It is focused on how allocations operate within the Chargify UI.It goes into greater detail on how the user interface will react when recording allocations.
// This documentation also goes into greater detail on how proration is taken into consideration when applying component allocations.
// ## Proration Schemes
// Changing the allocated quantity of a component mid-period can result in either a Charge or Credit being applied to the subscription. When creating an allocation via the API, you can pass the `upgrade_charge`, `downgrade_credit`, and `accrue_charge` to be applied.
// **Notice:** These proration and accural fields will be ignored for Prepaid Components since this component type always generate charges immediately without proration.
// For background information on prorated components and upgrade/downgrade schemes, see [Setting Component Allocations.](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404527849997#proration-upgrades-vs-downgrades).
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
// 2. [Component-level default value](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404527849997-Component-Allocations#component-allocations-0-0)
// 3. Allocation API call top level (outside of the `allocations` array)
// 4. [Site-level default value](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404527849997#proration-schemes)
// ### Order of Resolution for accrue charge
// 1. Allocation API call top level (outside of the `allocations` array)
// 2. [Site-level default value](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404527849997#proration-schemes)
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
		fmt.Sprintf("/subscriptions/%v/components/%v/allocations.json", subscriptionId, componentId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result models.AllocationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AllocationResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// ListAllocations takes context, subscriptionId, componentId, page as parameters and
// returns an models.ApiResponse with []models.AllocationResponse data and
// an error if there was an issue with the request or response.
// This endpoint returns the 50 most recent Allocations, ordered by most recent first.
// ## On/Off Components
// When a subscription's on/off component has been toggled to on (`1`) or off (`0`), usage will be logged in this response.
// ## Querying data via Chargify gem
// You can also query the current quantity via the [official Chargify Gem.](http://github.com/chargify/chargify_api_ares)
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
		fmt.Sprintf("/subscriptions/%v/components/%v/allocations.json", subscriptionId, componentId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.AllocationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.AllocationResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 401 {
		err = errors.NewApiError(401, "Unauthorized")
	}
	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	if resp.StatusCode == 422 {
		err = errors.NewApiError(422, "Unprocessable Entity (WebDAV)")
	}
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
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/allocations.json", subscriptionId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result []models.AllocationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.AllocationResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 401 {
		err = errors.NewApiError(401, "Unauthorized")
	}
	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	if resp.StatusCode == 422 {
		err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
	}
	return models.NewApiResponse(result, resp), err
}

// PreviewAllocations takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.AllocationPreviewResponse data and
// an error if there was an issue with the request or response.
// Chargify offers the ability to preview a potential subscription's **quantity-based** or **on/off** component allocation in the middle of the current billing period.  This is useful if you want users to be able to see the effect of a component operation before actually doing it.
// ## Fine-grained Component Control: Use with multiple `upgrade_charge`s or `downgrade_credits`
// When the allocation uses multiple different types of `upgrade_charge`s or `downgrade_credit`s, the Allocation is viewed as an Allocation which uses "Fine-Grained Component Control". As a result, the response will not include `direction` and `proration` within the `allocation_preview` at the `line_items` and `allocations` level respectfully.
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
		fmt.Sprintf("/subscriptions/%v/allocations/preview.json", subscriptionId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result models.AllocationPreviewResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AllocationPreviewResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 422 {
		err = errors.NewComponentAllocationError(422, "Unprocessable Entity (WebDAV)")
	}
	return models.NewApiResponse(result, resp), err
}

// UpdatePrepaidUsageAllocation takes context, subscriptionId, componentId, allocationId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// When the expiration interval options are selected on a prepaid usage component price point, all allocations will be created with an expiration date. This expiration date can be changed after the fact to allow for extending or shortening the allocation's active window.
// In order to change a prepaid usage allocation's expiration date, a PUT call must be made to the allocation's endpoint with a new expiration date.
// ## Limitations
// A few limitations exist when changing an allocation's expiration date:
// - An expiration date can only be changed for an allocation that belongs to a price point with expiration interval options explicitly set.
// - An expiration date can be changed towards the future with no limitations.
// - An expiration date can be changed towards the past (essentially expiring it) up to the subscription's current period beginning date.
func (s *SubscriptionComponentsController) UpdatePrepaidUsageAllocation(
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
		fmt.Sprintf("/subscriptions/%v/components/%v/allocations/%v.json", subscriptionId, componentId, allocationId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	if context.Response.StatusCode == 422 {
		err = errors.NewSubscriptionComponentAllocationError(422, "Unprocessable Entity (WebDAV)")
	}
	return context.Response, err
}

// DeletePrepaidUsageAllocation takes context, subscriptionId, componentId, allocationId, body as parameters and
// returns an models.ApiResponse with  data and
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
		fmt.Sprintf("/subscriptions/%v/components/%v/allocations/%v.json", subscriptionId, componentId, allocationId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	if context.Response.StatusCode == 422 {
		err = errors.NewSubscriptionComponentAllocationError(422, "Unprocessable Entity (WebDAV)")
	}
	return context.Response, err
}

// CreateUsage takes context, subscriptionId, componentId, body as parameters and
// returns an models.ApiResponse with models.UsageResponse data and
// an error if there was an issue with the request or response.
// ## Documentation
// Full documentation on how to create Components in the Chargify UI can be located [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677#creating-components). Additionally, for information on how to record component usage against a subscription, please see the following resources:
// + [Recording Metered Component Usage](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404527849997#reporting-metered-component-usage)
// + [Reporting Prepaid Component Status](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404527849997#reporting-prepaid-component-status)
// You may choose to report metered or prepaid usage to Chargify as often as you wish. You may report usage as it happens. You may also report usage periodically, such as each night or once per billing period. If usage events occur in your system very frequently (on the order of thousands of times an hour), it is best to accumulate usage into batches on your side, and then report those batches less frequently, such as daily. This will ensure you remain below any API throttling limits. If your use case requires higher rates of usage reporting, we recommend utilizing Events Based Components.
// ## Create Usage for Subscription
// This endpoint allows you to record an instance of metered or prepaid usage for a subscription. The `quantity` from usage for each component is accumulated to the `unit_balance` on the [Component Line Item](./b3A6MTQxMDgzNzQ-read-subscription-component) for the subscription.
// ## Price Point ID usage
// If you are using price points, for metered and prepaid usage components, Chargify gives you the option to specify a price point in your request.
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
	componentId int,
	body *models.CreateUsageRequest) (
	models.ApiResponse[models.UsageResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/components/%v/usages.json", subscriptionId, componentId),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(*body)
	}

	var result models.UsageResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UsageResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 422 {
		err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
	}
	return models.NewApiResponse(result, resp), err
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
	subscriptionId int,
	componentId int,
	sinceId *int,
	maxId *int,
	sinceDate *string,
	untilDate *string,
	page *int,
	perPage *int) (
	models.ApiResponse[[]models.UsageResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/components/%v/usages.json", subscriptionId, componentId),
	)
	req.Authenticate(true)
	if sinceId != nil {
		req.QueryParam("since_id", *sinceId)
	}
	if maxId != nil {
		req.QueryParam("max_id", *maxId)
	}
	if sinceDate != nil {
		req.QueryParam("since_date", *sinceDate)
	}
	if untilDate != nil {
		req.QueryParam("until_date", *untilDate)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}

	var result []models.UsageResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.UsageResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// ActivateEventBasedComponent takes context, subscriptionId, componentId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// In order to bill your subscribers on your Events data under the Events-Based Billing feature, the components must be activated for the subscriber.
// Learn more about the role of activation in the [Events-Based Billing docs](https://chargify.zendesk.com/hc/en-us/articles/4407720810907#activating-components-for-subscribers).
// Use this endpoint to activate an event-based component for a single subscription. Activating an event-based component causes Chargify to bill for events when the subscription is renewed.
// *Note: it is possible to stream events for a subscription at any time, regardless of component activation status. The activation status only determines if the subscription should be billed for event-based component usage at renewal.*
func (s *SubscriptionComponentsController) ActivateEventBasedComponent(
	ctx context.Context,
	subscriptionId int,
	componentId int) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/event_based_billing/subscriptions/%v/components/%v/activate.json", subscriptionId, componentId),
	)
	req.Authenticate(true)

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// DeactivateEventBasedComponent takes context, subscriptionId, componentId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Use this endpoint to deactivate an event-based component for a single subscription. Deactivating the event-based component causes Chargify to ignore related events at subscription renewal.
func (s *SubscriptionComponentsController) DeactivateEventBasedComponent(
	ctx context.Context,
	subscriptionId int,
	componentId int) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/event_based_billing/subscriptions/%v/components/%v/deactivate.json", subscriptionId, componentId),
	)
	req.Authenticate(true)

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// RecordEvent takes context, subdomain, apiHandle, storeUid, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// ## Documentation
// Events-Based Billing is an evolved form of metered billing that is based on data-rich events streamed in real-time from your system to Chargify.
// These events can then be transformed, enriched, or analyzed to form the computed totals of usage charges billed to your customers.
// This API allows you to stream events into the Chargify data ingestion engine.
// Learn more about the feature in general in the [Events-Based Billing help docs](https://chargify.zendesk.com/hc/en-us/articles/4407720613403).
// ## Record Event
// Use this endpoint to record a single event.
// *Note: this endpoint differs from the standard Chargify endpoints in that the URL subdomain will be `events` and your site subdomain will be included in the URL path. For example:*
// ```
// https://events.chargify.com/my-site-subdomain/events/my-stream-api-handle
// ```
func (s *SubscriptionComponentsController) RecordEvent(
	ctx context.Context,
	subdomain string,
	apiHandle string,
	storeUid *string,
	body *models.EBBEvent) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/%v/events/%v.json", subdomain, apiHandle),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if storeUid != nil {
		req.QueryParam("store_uid", *storeUid)
	}
	if body != nil {
		req.Json(*body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// RecordEvents takes context, subdomain, apiHandle, storeUid, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Use this endpoint to record a collection of events.
// *Note: this endpoint differs from the standard Chargify endpoints in that the subdomain will be `events` and your site subdomain will be included in the URL path.*
// A maximum of 1000 events can be published in a single request. A 422 will be returned if this limit is exceeded.
func (s *SubscriptionComponentsController) RecordEvents(
	ctx context.Context,
	subdomain string,
	apiHandle string,
	storeUid *string,
	body []models.EBBEvent) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/%v/events/%v/bulk.json", subdomain, apiHandle),
	)
	req.Authenticate(true)
	req.Header("Content-Type", "application/json")
	if storeUid != nil {
		req.QueryParam("store_uid", *storeUid)
	}
	if body != nil {
		req.Json(body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// ListSubscriptionComponentsForSite takes context, page, perPage, sort, direction, dateField, startDate, startDatetime, endDate, endDatetime, subscriptionIds, pricePointIds, productFamilyIds, include, filterUseSiteExchangeRate, filterCurrencies, filterSubscriptionStates, filterSubscriptionDateField, filterSubscriptionStartDate, filterSubscriptionStartDatetime, filterSubscriptionEndDate, filterSubscriptionEndDatetime as parameters and
// returns an models.ApiResponse with models.ListSubscriptionComponentsResponse data and
// an error if there was an issue with the request or response.
// This request will list components applied to each subscription.
func (s *SubscriptionComponentsController) ListSubscriptionComponentsForSite(
	ctx context.Context,
	page *int,
	perPage *int,
	sort *models.ListSubscriptionComponentsSortEnum,
	direction *models.SortingDirectionEnum,
	dateField *models.SubscriptionListDateFieldEnum,
	startDate *string,
	startDatetime *string,
	endDate *string,
	endDatetime *string,
	subscriptionIds []int,
	pricePointIds *models.IncludeNotNullEnum,
	productFamilyIds []int,
	include *models.ListSubscriptionComponentsIncludeEnum,
	filterUseSiteExchangeRate *bool,
	filterCurrencies []string,
	filterSubscriptionStates []models.SubscriptionStateFilterEnum,
	filterSubscriptionDateField *models.SubscriptionListDateFieldEnum,
	filterSubscriptionStartDate *string,
	filterSubscriptionStartDatetime *string,
	filterSubscriptionEndDate *string,
	filterSubscriptionEndDatetime *string) (
	models.ApiResponse[models.ListSubscriptionComponentsResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/subscriptions_components.json")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if direction != nil {
		req.QueryParam("direction", *direction)
	}
	if dateField != nil {
		req.QueryParam("date_field", *dateField)
	}
	if startDate != nil {
		req.QueryParam("start_date", *startDate)
	}
	if startDatetime != nil {
		req.QueryParam("start_datetime", *startDatetime)
	}
	if endDate != nil {
		req.QueryParam("end_date", *endDate)
	}
	if endDatetime != nil {
		req.QueryParam("end_datetime", *endDatetime)
	}
	if subscriptionIds != nil {
		req.QueryParam("subscription_ids", subscriptionIds)
	}
	if pricePointIds != nil {
		req.QueryParam("price_point_ids", *pricePointIds)
	}
	if productFamilyIds != nil {
		req.QueryParam("product_family_ids", productFamilyIds)
	}
	if include != nil {
		req.QueryParam("include", *include)
	}
	if filterUseSiteExchangeRate != nil {
		req.QueryParam("filter[use_site_exchange_rate]", *filterUseSiteExchangeRate)
	}
	if filterCurrencies != nil {
		req.QueryParam("filter[currencies]", filterCurrencies)
	}
	if filterSubscriptionStates != nil {
		req.QueryParam("filter[subscription][states]", filterSubscriptionStates)
	}
	if filterSubscriptionDateField != nil {
		req.QueryParam("filter[subscription][date_field]", *filterSubscriptionDateField)
	}
	if filterSubscriptionStartDate != nil {
		req.QueryParam("filter[subscription][start_date]", *filterSubscriptionStartDate)
	}
	if filterSubscriptionStartDatetime != nil {
		req.QueryParam("filter[subscription][start_datetime]", *filterSubscriptionStartDatetime)
	}
	if filterSubscriptionEndDate != nil {
		req.QueryParam("filter[subscription][end_date]", *filterSubscriptionEndDate)
	}
	if filterSubscriptionEndDatetime != nil {
		req.QueryParam("filter[subscription][end_datetime]", *filterSubscriptionEndDatetime)
	}
	var result models.ListSubscriptionComponentsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListSubscriptionComponentsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
