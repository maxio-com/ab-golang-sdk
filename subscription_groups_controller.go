package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
)

// SubscriptionGroupsController represents a controller struct.
type SubscriptionGroupsController struct {
    baseController
}

// NewSubscriptionGroupsController creates a new instance of SubscriptionGroupsController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionGroupsController.
func NewSubscriptionGroupsController(baseController baseController) *SubscriptionGroupsController {
    subscriptionGroupsController := SubscriptionGroupsController{baseController: baseController}
    return &subscriptionGroupsController
}

// SignupWithSubscriptionGroup takes context, body as parameters and
// returns an models.ApiResponse with models.SubscriptionGroupSignupResponse data and
// an error if there was an issue with the request or response.
// Create multiple subscriptions at once under the same customer and consolidate them into a subscription group.
// You must provide one and only one of the `payer_id`/`payer_reference`/`payer_attributes` for the customer attached to the group.
// You must provide one and only one of the `payment_profile_id`/`credit_card_attributes`/`bank_account_attributes` for the payment profile attached to the group.
// Only one of the `subscriptions` can have `"primary": true` attribute set.
// When passing product to a subscription you can use either `product_id` or `product_handle` or `offer_id`. You can also use `custom_price` instead.
func (s *SubscriptionGroupsController) SignupWithSubscriptionGroup(
    ctx context.Context,
    body *models.SubscriptionGroupSignupRequest) (
    models.ApiResponse[models.SubscriptionGroupSignupResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscription_groups/signup.json")
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    var result models.SubscriptionGroupSignupResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionGroupSignupResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewSubscriptionGroupSignupErrorResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// CreateSubscriptionGroup takes context, body as parameters and
// returns an models.ApiResponse with models.SubscriptionGroupResponse data and
// an error if there was an issue with the request or response.
// Creates a subscription group with given members.
func (s *SubscriptionGroupsController) CreateSubscriptionGroup(
    ctx context.Context,
    body *models.CreateSubscriptionGroupRequest) (
    models.ApiResponse[models.SubscriptionGroupResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscription_groups.json")
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    var result models.SubscriptionGroupResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionGroupResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewSingleStringErrorResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ListSubscriptionGroups takes context, page, perPage, include as parameters and
// returns an models.ApiResponse with models.ListSubscriptionGroupsResponse data and
// an error if there was an issue with the request or response.
// Returns an array of subscription groups for the site. The response is paginated and will return a `meta` key with pagination information.
// #### Account Balance Information
// Account balance information for the subscription groups is not returned by default. If this information is desired, the `include[]=account_balances` parameter must be provided with the request.
func (s *SubscriptionGroupsController) ListSubscriptionGroups(
    ctx context.Context,
    page *int,
    perPage *int,
    include *string) (
    models.ApiResponse[models.ListSubscriptionGroupsResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/subscription_groups.json")
    req.Authenticate(true)
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if include != nil {
        req.QueryParam("include", *include)
    }
    var result models.ListSubscriptionGroupsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListSubscriptionGroupsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// ReadSubscriptionGroup takes context, uid as parameters and
// returns an models.ApiResponse with models.FullSubscriptionGroupResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to find subscription group details.
// #### Current Billing Amount in Cents
// Current billing amount for the subscription group is not returned by default. If this information is desired, the `include[]=current_billing_amount_in_cents` parameter must be provided with the request.
func (s *SubscriptionGroupsController) ReadSubscriptionGroup(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.FullSubscriptionGroupResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/subscription_groups/%v.json", uid),
    )
    req.Authenticate(true)
    
    var result models.FullSubscriptionGroupResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FullSubscriptionGroupResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionGroupMembers takes context, uid, body as parameters and
// returns an models.ApiResponse with models.SubscriptionGroupResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to update subscription group members.
// `"member_ids": []` should contain an array of both subscription IDs to set as group members and subscription IDs already present in the groups. Not including them will result in removing them from subscription group. To clean up members, just leave the array empty.
func (s *SubscriptionGroupsController) UpdateSubscriptionGroupMembers(
    ctx context.Context,
    uid string,
    body *models.UpdateSubscriptionGroupRequest) (
    models.ApiResponse[models.SubscriptionGroupResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/subscription_groups/%v.json", uid),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.SubscriptionGroupResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionGroupResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewSubscriptionGroupUpdateErrorResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// DeleteSubscriptionGroup takes context, uid as parameters and
// returns an models.ApiResponse with models.DeleteSubscriptionGroupResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to delete subscription group.
// Only groups without members can be deleted
func (s *SubscriptionGroupsController) DeleteSubscriptionGroup(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.DeleteSubscriptionGroupResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/subscription_groups/%v.json", uid),
    )
    req.Authenticate(true)
    
    var result models.DeleteSubscriptionGroupResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.DeleteSubscriptionGroupResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    return models.NewApiResponse(result, resp), err
}

// FindSubscriptionGroup takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.FullSubscriptionGroupResponse data and
// an error if there was an issue with the request or response.
// Use this endpoint to find subscription group associated with subscription.
// If the subscription is not in a group endpoint will return 404 code.
func (s *SubscriptionGroupsController) FindSubscriptionGroup(
    ctx context.Context,
    subscriptionId string) (
    models.ApiResponse[models.FullSubscriptionGroupResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/subscription_groups/lookup.json")
    req.Authenticate(true)
    req.QueryParam("subscription_id", subscriptionId)
    var result models.FullSubscriptionGroupResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FullSubscriptionGroupResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    return models.NewApiResponse(result, resp), err
}

// AddSubscriptionToGroup takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionGroupResponse data and
// an error if there was an issue with the request or response.
// For sites making use of the [Relationship Billing](https://chargify.zendesk.com/hc/en-us/articles/4407737494171) and [Customer Hierarchy](https://chargify.zendesk.com/hc/en-us/articles/4407746683291) features, it is possible to add existing subscriptions to subscription groups.
// Passing `group` parameters with a `target` containing a `type` and optional `id` is all that's needed. When the `target` parameter specifies a `"customer"` or `"subscription"` that is already part of a hierarchy, the subscription will become a member of the customer's subscription group.  If the target customer or subscription is not part of a subscription group, a new group will be created and the subscription will become part of the group with the specified target customer set as the responsible payer for the group's subscriptions.
// **Please Note:** In order to add an existing subscription to a subscription group, it must belong to either the same customer record as the target, or be within the same customer hierarchy.
// Rather than specifying a customer, the `target` parameter could instead simply have a value of
// * `"self"` which indicates the subscription will be paid for not by some other customer, but by the subscribing customer,
// * `"parent"` which indicates the subscription will be paid for by the subscribing customer's parent within a customer hierarchy, or
// * `"eldest"` which indicates the subscription will be paid for by the root-level customer in the subscribing customer's hierarchy.
// To create a new subscription into a subscription group, please reference the following:
// [Create Subscription in a Subscription Group](https://developers.chargify.com/docs/api-docs/d571659cf0f24-create-subscription#subscription-in-a-subscription-group)
func (s *SubscriptionGroupsController) AddSubscriptionToGroup(
    ctx context.Context,
    subscriptionId int,
    body *models.AddSubscriptionToAGroup) (
    models.ApiResponse[models.SubscriptionGroupResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/group.json", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.SubscriptionGroupResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionGroupResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    return models.NewApiResponse(result, resp), err
}

// RemoveSubscriptionFromGroup takes context, subscriptionId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// For sites making use of the [Relationship Billing](https://chargify.zendesk.com/hc/en-us/articles/4407737494171) and [Customer Hierarchy](https://chargify.zendesk.com/hc/en-us/articles/4407746683291) features, it is possible to remove existing subscription from subscription group.
func (s *SubscriptionGroupsController) RemoveSubscriptionFromGroup(
    ctx context.Context,
    subscriptionId int) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/subscriptions/%v/group.json", subscriptionId),
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
    if context.Response.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if context.Response.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return context.Response, err
}
