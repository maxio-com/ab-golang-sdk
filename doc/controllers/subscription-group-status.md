# Subscription Group Status

```go
subscriptionGroupStatusController := client.SubscriptionGroupStatusController()
```

## Class Name

`SubscriptionGroupStatusController`

## Methods

* [Cancel Subscriptions in Group](../../doc/controllers/subscription-group-status.md#cancel-subscriptions-in-group)
* [Initiate Delayed Cancellation for Group](../../doc/controllers/subscription-group-status.md#initiate-delayed-cancellation-for-group)
* [Cancel Delayed Cancellation for Group](../../doc/controllers/subscription-group-status.md#cancel-delayed-cancellation-for-group)
* [Reactivate Subscription Group](../../doc/controllers/subscription-group-status.md#reactivate-subscription-group)


# Cancel Subscriptions in Group

Cancels all subscriptions within the specified group immediately. The group is identified by the `uid` that is passed in the URL. To successfully cancel the group, the primary subscription must be on automatic billing. The group members must be on automatic billing or prepaid.

To cancel a subscription group while also charging for any unbilled usage on metered or prepaid components, the `charge_unbilled_usage=true` parameter must be included in the request.

```go
CancelSubscriptionsInGroup(
    ctx context.Context,
    uid string,
    body *models.CancelGroupedSubscriptionsRequest) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |
| `body` | [`*models.CancelGroupedSubscriptionsRequest`](../../doc/models/cancel-grouped-subscriptions-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

body := models.CancelGroupedSubscriptionsRequest{
    ChargeUnbilledUsage:  models.ToPointer(true),
}

resp, err := subscriptionGroupStatusController.CancelSubscriptionsInGroup(ctx, uid, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Initiate Delayed Cancellation for Group

This endpoint will schedule all subscriptions within the specified group to be canceled at the end of their billing period. The group is identified by its uid passed in the URL.

All subscriptions in the group must be on automatic billing in order to successfully cancel them, and the group must not be in a "past_due" state.

```go
InitiateDelayedCancellationForGroup(
    ctx context.Context,
    uid string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

resp, err := subscriptionGroupStatusController.InitiateDelayedCancellationForGroup(ctx, uid)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Cancel Delayed Cancellation for Group

Removing the delayed cancellation on a subscription group will ensure that the subscriptions do not get canceled at the end of the period. The request will reset the `cancel_at_end_of_period` flag to false on each member in the group.

```go
CancelDelayedCancellationForGroup(
    ctx context.Context,
    uid string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

resp, err := subscriptionGroupStatusController.CancelDelayedCancellationForGroup(ctx, uid)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Reactivate Subscription Group

This endpoint will attempt to reactivate or resume a cancelled subscription group. Upon reactivation, any canceled invoices created after the beginning of the primary subscription's billing period will be reopened and payment will be attempted on them. If the subscription group is being reactivated (as opposed to resumed), new charges will also be assessed for the new billing period.

Whether a subscription group is reactivated (a new billing period is created) or resumed (the current billing period is respected) will depend on the parameters that are sent with the request as well as the date of the request relative to the primary subscription's period.

## Reactivating within the current period

If a subscription group is cancelled and reactivated within the primary subscription's current period, we can choose to either start a new billing period or maintain the existing one. If we want to maintain the existing billing period the `resume=true` option must be passed in request parameters.

An exception to the above are subscriptions that are on calendar billing. These subscriptions cannot be reactivated within the current period. If the `resume=true` option is not passed the request will return an error.

The `resume_members` option is ignored in this case. All eligible group members will be automatically resumed.

## Reactivating beyond the current period

In this case, a subscription group can only be reactivated with a new billing period. If the `resume=true` option is passed it will be ignored.

Member subscriptions can have billing periods that are longer than the primary (e.g. a monthly primary with annual group members). If the primary subscription in a group cannot be reactivated within the current period, but other group members can be, passing `resume_members=true` will resume the existing billing period for eligible group members. The primary subscription will begin a new billing period.

For calendar billing subscriptions, the new billing period created will be a partial one, spanning from the date of reactivation to the next corresponding calendar renewal date.

```go
ReactivateSubscriptionGroup(
    ctx context.Context,
    uid string,
    body *models.ReactivateSubscriptionGroupRequest) (
    models.ApiResponse[models.ReactivateSubscriptionGroupResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |
| `body` | [`*models.ReactivateSubscriptionGroupRequest`](../../doc/models/reactivate-subscription-group-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ReactivateSubscriptionGroupResponse](../../doc/models/reactivate-subscription-group-response.md).

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

body := models.ReactivateSubscriptionGroupRequest{
    Resume:               models.ToPointer(true),
}

apiResponse, err := subscriptionGroupStatusController.ReactivateSubscriptionGroup(ctx, uid, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "uid": "grp_93wgm89cbjkw6",
  "scheme": 1,
  "customer_id": 1,
  "payment_profile_id": 1,
  "subscription_ids": [
    1,
    2
  ],
  "primary_subscription_id": 1,
  "next_assessment_at": "2020-06-18T12:00:00-04:00",
  "state": "active",
  "cancel_at_end_of_period": false
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

