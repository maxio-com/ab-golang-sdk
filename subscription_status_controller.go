package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// SubscriptionStatusController represents a controller struct.
type SubscriptionStatusController struct {
    baseController
}

// NewSubscriptionStatusController creates a new instance of SubscriptionStatusController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionStatusController.
func NewSubscriptionStatusController(baseController baseController) *SubscriptionStatusController {
    subscriptionStatusController := SubscriptionStatusController{baseController: baseController}
    return &subscriptionStatusController
}

// RetrySubscription takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Chargify offers the ability to retry collecting the balance due on a past due Subscription without waiting for the next scheduled attempt.
// ## Successful Reactivation
// The response will be `200 OK` with the updated Subscription.
// ## Failed Reactivation
// The response will be `422 "Unprocessable Entity`.
func (s *SubscriptionStatusController) RetrySubscription(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/subscriptions/%v/retry.json", subscriptionId),
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
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// CancelSubscription takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// The DELETE action causes the cancellation of the Subscription. This means, the method sets the Subscription state to "canceled".
func (s *SubscriptionStatusController) CancelSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.CancellationRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/subscriptions/%v.json", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
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
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewApiError(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ResumeSubscription takes context, subscriptionId, calendarBillingResumptionCharge as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Resume a paused (on-hold) subscription. If the normal next renewal date has not passed, the subscription will return to active and will renew on that date.  Otherwise, it will behave like a reactivation, setting the billing date to 'now' and charging the subscriber.
func (s *SubscriptionStatusController) ResumeSubscription(
    ctx context.Context,
    subscriptionId int,
    calendarBillingResumptionCharge *models.ResumptionCharge) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/resume.json", subscriptionId),
    )
    req.Authenticate(true)
    if calendarBillingResumptionCharge != nil {
        req.QueryParam("calendar_billing['resumption_charge']", *calendarBillingResumptionCharge)
    }
    
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
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// PauseSubscription takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// This will place the subscription in the on_hold state and it will not renew.
// ## Limitations
// You may not place a subscription on hold if the `next_billing` date is within 24 hours.
func (s *SubscriptionStatusController) PauseSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.PauseRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/hold.json", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
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
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// UpdateAutomaticSubscriptionResumption takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Once a subscription has been paused / put on hold, you can update the date which was specified to automatically resume the subscription.
// To update a subscription's resume date, use this method to change or update the `automatically_resume_at` date.
// ### Remove the resume date
// Alternately, you can change the `automatically_resume_at` to `null` if you would like the subscription to not have a resume date.
func (s *SubscriptionStatusController) UpdateAutomaticSubscriptionResumption(
    ctx context.Context,
    subscriptionId int,
    body *models.PauseRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/subscriptions/%v/hold.json", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
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
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ReactivateSubscription takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// Chargify offers the ability to reactivate a previously canceled subscription. For details on how the reactivation works, and how to reactivate subscriptions through the application, see [reactivation](https://chargify.zendesk.com/hc/en-us/articles/4407898737691).
// **Please note: The term
// "resume" is used also during another process in Chargify. This occurs when an on-hold subscription is "resumed". This returns the subscription to an active state.**
// + The response returns the subscription object in the `active` or `trialing` state.
// + The `canceled_at` and `cancellation_message` fields do not have values.
// + The method works for "Canceled" or "Trial Ended" subscriptions.
// + It will not work for items not marked as "Canceled", "Unpaid", or "Trial Ended".
// ## Resume the current billing period for a subscription
// A subscription is considered "resumable" if you are attempting to reactivate within the billing period the subscription was canceled in.
// A resumed subscription's billing date remains the same as before it was canceled. In other words, it does not start a new billing period. Payment may or may not be collected for a resumed subscription, depending on whether or not the subscription had a balance when it was canceled (for example, if it was canceled because of dunning).
// Consider a subscription which was created on June 1st, and would renew on July 1st. The subscription is then canceled on June 15.
// If a reactivation with `resume: true` were attempted _before_ what would have been the next billing date of July 1st, then Chargify would resume the subscription.
// If a reactivation with `resume: true` were attempted _after_ what would have been the next billing date of July 1st, then Chargify would not resume the subscription, and instead it would be reactivated with a new billing period.
// | Canceled | Reactivation | Resumable? |
// |---|---|---|
// | Jun 15 | June 28 | Yes |
// | Jun 15 | July 2 | No |
// ## Reactivation Scenarios
// ### Reactivating Canceled Subscription While Preserving Balance
// + Given you have a product that costs $20
// + Given you have a canceled subscription to the $20 product
// + 1 charge should exist for $20
// + 1 payment should exist for $20
// + When the subscription has canceled due to dunning, it retained a negative balance of $20
// #### Results
// The resulting charges upon reactivation will be:
// + 1 charge for $20 for the new product
// + 1 charge for $20 for the balance due
// + Total charges = $40
// + The subscription will transition to active
// + The subscription balance will be zero
// ### Reactivating a Canceled Subscription With Coupon
// + Given you have a canceled subscription
// + It has no current period defined
// + You have a coupon code "EARLYBIRD"
// + The coupon is set to recur for 6 periods
// PUT request sent to:
// `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?coupon_code=EARLYBIRD`
// #### Results
// + The subscription will transition to active
// + The subscription should have applied a coupon with code "EARLYBIRD"
// ### Reactivating Canceled Subscription With a Trial, Without the include_trial Flag
// + Given you have a canceled subscription
// + The product associated with the subscription has a trial
// + PUT request to
// `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json`
// #### Results
// + The subscription will transition to active
// ### Reactivating Canceled Subscription With Trial, With the include_trial Flag
// + Given you have a canceled subscription
// + The product associated with the subscription has a trial
// + Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?include_trial=1`
// #### Results
// + The subscription will transition to trialing
// ### Reactivating Trial Ended Subscription
// + Given you have a trial_ended subscription
// + Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json`
// #### Results
// + The subscription will transition to active
// ### Resuming a Canceled Subscription
// + Given you have a `canceled` subscription and it is resumable
// + Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume=true`
// #### Results
// + The subscription will transition to active
// + The next billing date should not have changed
// ### Attempting to resume a subscription which is not resumable
// + Given you have a `canceled` subscription, and it is not resumable
// + Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume=true`
// #### Results
// + The subscription will transition to active, with a new billing period.
// ### Attempting to resume but not reactivate a subscription which is not resumable
// + Given you have a `canceled` subscription, and it is not resumable
// + Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume[require_resume]=true`
// + The response status should be "422 UNPROCESSABLE ENTITY"
// + The subscription should be canceled with the following response
// ```
// {
// "errors": ["Request was 'resume only', but this subscription cannot be resumed."]
// }
// ```
// #### Results
// + The subscription should remain `canceled`
// + The next billing date should not have changed
// ### Resuming Subscription Which Was Trialing
// + Given you have a `trial_ended` subscription, and it is resumable
// + And the subscription was canceled in the middle of a trial
// + And there is still time left on the trial
// + Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume=true`
// #### Results
// + The subscription will transition to trialing
// + The next billing date should not have changed
// ### Resuming Subscription Which Was trial_ended
// + Given you have a `trial_ended` subscription, and it is resumable
// + Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume=true`
// #### Results
// + The subscription will transition to active
// + The next billing date should not have changed
// + Any product-related charges should have been collected
func (s *SubscriptionStatusController) ReactivateSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.ReactivateSubscriptionRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/subscriptions/%v/reactivate.json", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
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
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// InitiateDelayedCancellation takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.DelayedCancellationResponse data and
// an error if there was an issue with the request or response.
// Chargify offers the ability to cancel a subscription at the end of the current billing period. This period is set by its current product.
// Requesting to cancel the subscription at the end of the period sets the `cancel_at_end_of_period` flag to true.
// Note that you cannot set `cancel_at_end_of_period` at subscription creation, or if the subscription is past due.
func (s *SubscriptionStatusController) InitiateDelayedCancellation(
    ctx context.Context,
    subscriptionId int,
    body *models.CancellationRequest) (
    models.ApiResponse[models.DelayedCancellationResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/delayed_cancel.json", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.DelayedCancellationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.DelayedCancellationResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    return models.NewApiResponse(result, resp), err
}

// CancelDelayedCancellation takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.DelayedCancellationResponse data and
// an error if there was an issue with the request or response.
// Removing the delayed cancellation on a subscription will ensure that it doesn't get canceled at the end of the period that it is in. The request will reset the `cancel_at_end_of_period` flag to `false`.
// This endpoint is idempotent. If the subscription was not set to cancel in the future, removing the delayed cancellation has no effect and the call will be successful.
func (s *SubscriptionStatusController) CancelDelayedCancellation(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.DelayedCancellationResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/subscriptions/%v/delayed_cancel.json", subscriptionId),
    )
    req.Authenticate(true)
    
    var result models.DelayedCancellationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.DelayedCancellationResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    return models.NewApiResponse(result, resp), err
}

// CancelDunning takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// If a subscription is currently in dunning, the subscription will be set to active and the active Dunner will be resolved.
func (s *SubscriptionStatusController) CancelDunning(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/cancel_dunning.json", subscriptionId),
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

// PreviewRenewal takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.RenewalPreviewResponse data and
// an error if there was an issue with the request or response.
// The Chargify API allows you to preview a renewal by posting to the renewals endpoint. Renewal Preview is an object representing a subscription’s next assessment. You can retrieve it to see a snapshot of how much your customer will be charged on their next renewal.
// The "Next Billing" amount and "Next Billing" date are already represented in the UI on each Subscriber's Summary. For more information, please see our documentation [here](https://chargify.zendesk.com/hc/en-us/articles/4407884887835#next-billing).
// ## Optional Component Fields
// This endpoint is particularly useful due to the fact that it will return the computed billing amount for the base product and the components which are in use by a subscriber.
// By default, the preview will include billing details for all components _at their **current** quantities_. This means:
// * Current `allocated_quantity` for quantity-based components
// * Current enabled/disabled status for on/off components
// * Current metered usage `unit_balance` for metered components
// * Current metric quantity value for events recorded thus far for events-based components
// In the above statements, "current" means the quantity or value as of the call to the renewal preview endpoint. We do not predict end-of-period values for components, so metered or events-based usage may be less than it will eventually be at the end of the period.
// Optionally, **you may provide your own custom quantities** for any component to see a billing preview for non-current quantities. This is accomplished by sending a request body with data under the `components` key. See the request body documentation below.
// ## Subscription Side Effects
// You can request a `POST` to obtain this data from the endpoint without any side effects. Plain and simple, this will preview data, not log any changes against a subscription.
func (s *SubscriptionStatusController) PreviewRenewal(
    ctx context.Context,
    subscriptionId int,
    body *models.RenewalPreviewRequest) (
    models.ApiResponse[models.RenewalPreviewResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/renewals/preview.json", subscriptionId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.RenewalPreviewResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RenewalPreviewResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 422 {
        err = errors.NewErrorListResponse(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}
