# Subscription Status

```go
subscriptionStatusController := client.SubscriptionStatusController()
```

## Class Name

`SubscriptionStatusController`

## Methods

* [Retry Subscription](../../doc/controllers/subscription-status.md#retry-subscription)
* [Cancel Subscription](../../doc/controllers/subscription-status.md#cancel-subscription)
* [Resume Subscription](../../doc/controllers/subscription-status.md#resume-subscription)
* [Pause Subscription](../../doc/controllers/subscription-status.md#pause-subscription)
* [Update Automatic Subscription Resumption](../../doc/controllers/subscription-status.md#update-automatic-subscription-resumption)
* [Reactivate Subscription](../../doc/controllers/subscription-status.md#reactivate-subscription)
* [Initiate Delayed Cancellation](../../doc/controllers/subscription-status.md#initiate-delayed-cancellation)
* [Cancel Delayed Cancellation](../../doc/controllers/subscription-status.md#cancel-delayed-cancellation)
* [Cancel Dunning](../../doc/controllers/subscription-status.md#cancel-dunning)
* [Preview Renewal](../../doc/controllers/subscription-status.md#preview-renewal)


# Retry Subscription

Chargify offers the ability to retry collecting the balance due on a past due Subscription without waiting for the next scheduled attempt.

## Successful Reactivation

The response will be `200 OK` with the updated Subscription.

## Failed Reactivation

The response will be `422 "Unprocessable Entity`.

```go
RetrySubscription(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.SubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |

## Response Type

[`models.SubscriptionResponse`](../../doc/models/subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

apiResponse, err := subscriptionStatusController.RetrySubscription(ctx, subscriptionId)
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
  "subscription": {
    "id": 46330,
    "state": "active",
    "trial_started_at": null,
    "trial_ended_at": null,
    "activated_at": "2018-10-22T13:10:46-06:00",
    "created_at": "2018-10-22T13:10:46-06:00",
    "updated_at": "2021-06-10T09:23:43-06:00",
    "expires_at": null,
    "balance_in_cents": 18600,
    "current_period_ends_at": "2021-06-22T13:10:46-06:00",
    "next_assessment_at": "2021-06-22T13:10:46-06:00",
    "canceled_at": null,
    "cancellation_message": null,
    "next_product_id": null,
    "cancel_at_end_of_period": null,
    "payment_collection_method": "automatic",
    "snap_day": null,
    "cancellation_method": null,
    "product_price_point_id": 3464,
    "next_product_price_point_id": null,
    "receives_invoice_emails": null,
    "net_terms": null,
    "locale": null,
    "currency": "USD",
    "reference": null,
    "scheduled_cancellation_at": null,
    "current_period_started_at": "2021-05-22T13:10:46-06:00",
    "previous_state": "past_due",
    "signup_payment_id": 651268,
    "signup_revenue": "6.00",
    "delayed_cancel_at": null,
    "coupon_code": null,
    "total_revenue_in_cents": 600,
    "product_price_in_cents": 600,
    "product_version_number": 501,
    "payment_type": null,
    "referral_code": "rzqvrx",
    "coupon_use_count": null,
    "coupon_uses_allowed": null,
    "reason_code": null,
    "automatically_resume_at": null,
    "coupon_codes": [],
    "offer_id": null,
    "credit_balance_in_cents": 0,
    "prepayment_balance_in_cents": 0,
    "payer_id": 142365,
    "stored_credential_transaction_id": null,
    "next_product_handle": null,
    "on_hold_at": null,
    "prepaid_dunning": false,
    "customer": {
      "id": 142365,
      "first_name": "Lavern",
      "last_name": "Fahey",
      "organization": null,
      "email": "millie2@example.com",
      "created_at": "2018-10-22T13:10:46-06:00",
      "updated_at": "2018-10-22T13:10:46-06:00",
      "reference": null,
      "address": null,
      "address_2": null,
      "city": null,
      "state": null,
      "zip": null,
      "country": null,
      "phone": null,
      "portal_invite_last_sent_at": null,
      "portal_invite_last_accepted_at": null,
      "verified": false,
      "portal_customer_created_at": "2018-10-22T13:10:46-06:00",
      "vat_number": null,
      "cc_emails": "john@example.com, sue@example.com",
      "tax_exempt": false,
      "parent_id": null,
      "locale": null
    },
    "product": {
      "id": 8080,
      "name": "Pro Versions",
      "handle": null,
      "description": "",
      "accounting_code": "",
      "request_credit_card": true,
      "expiration_interval": null,
      "expiration_interval_unit": "month",
      "created_at": "2019-02-15T10:15:00-07:00",
      "updated_at": "2019-02-15T10:30:34-07:00",
      "price_in_cents": 600,
      "interval": 1,
      "interval_unit": "month",
      "initial_charge_in_cents": null,
      "trial_price_in_cents": null,
      "trial_interval": null,
      "trial_interval_unit": "month",
      "archived_at": null,
      "require_credit_card": true,
      "return_params": "",
      "require_shipping_address": false,
      "request_billing_address": false,
      "require_billing_address": false,
      "taxable": false,
      "update_return_url": "",
      "tax_code": "",
      "initial_charge_after_trial": false,
      "default_product_price_point_id": 3464,
      "version_number": 501,
      "update_return_params": "",
      "product_price_point_id": 3464,
      "product_price_point_name": "Default",
      "product_price_point_handle": "uuid:5305c3f0-1375-0137-5619-065dfbfdc636",
      "product_family": {
        "id": 37,
        "name": "Acme Projects",
        "description": null,
        "handle": "acme-projects",
        "accounting_code": null,
        "created_at": "2013-02-20T15:05:51-07:00",
        "updated_at": "2013-02-20T15:05:51-07:00"
      },
      "public_signup_pages": [
        {
          "id": 1540,
          "return_url": null,
          "return_params": "",
          "url": "https://acme-test.staging-chargifypay.com/subscribe/2f6y53rrqgsf"
        }
      ]
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Cancel Subscription

The DELETE action causes the cancellation of the Subscription. This means, the method sets the Subscription state to "canceled".

```go
CancelSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.CancellationRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.CancellationRequest`](../../doc/models/cancellation-request.md) | Body, Optional | - |

## Response Type

[`models.SubscriptionResponse`](../../doc/models/subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

apiResponse, err := subscriptionStatusController.CancelSubscription(ctx, subscriptionId, nil)
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
  "subscription": {
    "id": 15254809,
    "state": "canceled",
    "trial_started_at": null,
    "trial_ended_at": null,
    "activated_at": "2016-11-15T15:33:44-05:00",
    "created_at": "2016-11-15T15:33:44-05:00",
    "updated_at": "2016-11-15T17:13:06-05:00",
    "expires_at": null,
    "balance_in_cents": 0,
    "current_period_ends_at": "2017-08-29T12:00:00-04:00",
    "next_assessment_at": "2017-08-29T12:00:00-04:00",
    "canceled_at": "2016-11-15T17:13:06-05:00",
    "cancellation_message": "Canceling the subscription via the API",
    "next_product_id": null,
    "cancel_at_end_of_period": false,
    "payment_collection_method": "automatic",
    "snap_day": null,
    "cancellation_method": "merchant_api",
    "current_period_started_at": "2016-11-15T15:33:44-05:00",
    "previous_state": "active",
    "signup_payment_id": 0,
    "signup_revenue": "0.00",
    "delayed_cancel_at": null,
    "coupon_code": null,
    "total_revenue_in_cents": 0,
    "product_price_in_cents": 1000,
    "product_version_number": 7,
    "payment_type": "credit_card",
    "referral_code": "tg8qbq",
    "coupon_use_count": null,
    "coupon_uses_allowed": null,
    "customer": {
      "id": 14731081,
      "first_name": "John",
      "last_name": "Doe",
      "organization": "Acme Widgets",
      "email": "john.doe@example.com",
      "created_at": "2016-11-15T15:33:44-05:00",
      "updated_at": "2016-11-15T15:33:45-05:00",
      "reference": "123",
      "address": null,
      "address_2": null,
      "city": null,
      "state": null,
      "zip": null,
      "country": null,
      "phone": null,
      "portal_invite_last_sent_at": "2016-11-15T15:33:45-05:00",
      "portal_invite_last_accepted_at": null,
      "verified": false,
      "portal_customer_created_at": "2016-11-15T15:33:45-05:00",
      "cc_emails": null
    },
    "product": {
      "id": 3792003,
      "name": "$10 Basic Plan",
      "handle": "basic",
      "description": "lorem ipsum",
      "accounting_code": "basic",
      "request_credit_card": false,
      "expiration_interval": null,
      "expiration_interval_unit": "never",
      "created_at": "2016-03-24T13:38:39-04:00",
      "updated_at": "2016-11-03T13:03:05-04:00",
      "price_in_cents": 1000,
      "interval": 1,
      "interval_unit": "day",
      "initial_charge_in_cents": null,
      "trial_price_in_cents": null,
      "trial_interval": null,
      "trial_interval_unit": "month",
      "archived_at": null,
      "require_credit_card": false,
      "return_params": "",
      "taxable": false,
      "update_return_url": "",
      "initial_charge_after_trial": false,
      "version_number": 7,
      "update_return_params": "",
      "product_family": {
        "id": 527890,
        "name": "Acme Projects",
        "description": "",
        "handle": "billing-plans",
        "accounting_code": null
      },
      "public_signup_pages": [
        {
          "id": 281054,
          "return_url": "http://www.example.com?successfulsignup",
          "return_params": "",
          "url": "https://general-goods.chargify.com/subscribe/kqvmfrbgd89q/basic"
        },
        {
          "id": 281240,
          "return_url": "",
          "return_params": "",
          "url": "https://general-goods.chargify.com/subscribe/dkffht5dxfd8/basic"
        },
        {
          "id": 282694,
          "return_url": "",
          "return_params": "",
          "url": "https://general-goods.chargify.com/subscribe/jwffwgdd95s8/basic"
        }
      ]
    },
    "credit_card": {
      "id": 10202898,
      "first_name": "John",
      "last_name": "Doe",
      "masked_card_number": "XXXX-XXXX-XXXX-1111",
      "card_type": "visa",
      "expiration_month": 12,
      "expiration_year": 2020,
      "customer_id": 14731081,
      "current_vault": "authorizenet",
      "vault_token": "12345",
      "billing_address": null,
      "billing_city": null,
      "billing_state": null,
      "billing_zip": null,
      "billing_country": null,
      "customer_vault_token": "67890",
      "billing_address_2": null,
      "payment_type": "credit_card"
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | `ApiError` |


# Resume Subscription

Resume a paused (on-hold) subscription. If the normal next renewal date has not passed, the subscription will return to active and will renew on that date.  Otherwise, it will behave like a reactivation, setting the billing date to 'now' and charging the subscriber.

```go
ResumeSubscription(
    ctx context.Context,
    subscriptionId int,
    calendarBillingResumptionCharge *models.ResumptionCharge) (
    models.ApiResponse[models.SubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `calendarBillingResumptionCharge` | [`*models.ResumptionCharge`](../../doc/models/resumption-charge.md) | Query, Optional | (For calendar billing subscriptions only) The way that the resumed subscription's charge should be handled |

## Response Type

[`models.SubscriptionResponse`](../../doc/models/subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222Liquid error: Value cannot be null. (Parameter 'key')

apiResponse, err := subscriptionStatusController.ResumeSubscription(ctx, subscriptionId, Liquid error: Value cannot be null. (Parameter 'key'))
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
  "subscription": {
    "id": 18220670,
    "state": "active",
    "trial_started_at": null,
    "trial_ended_at": null,
    "activated_at": "2017-06-27T13:45:15-05:00",
    "created_at": "2017-06-27T13:45:13-05:00",
    "updated_at": "2017-06-30T09:26:50-05:00",
    "expires_at": null,
    "balance_in_cents": 10000,
    "current_period_ends_at": "2017-06-30T12:00:00-05:00",
    "next_assessment_at": "2017-06-30T12:00:00-05:00",
    "canceled_at": null,
    "cancellation_message": null,
    "next_product_id": null,
    "cancel_at_end_of_period": false,
    "payment_collection_method": "automatic",
    "snap_day": "end",
    "cancellation_method": null,
    "current_period_started_at": "2017-06-27T13:45:13-05:00",
    "previous_state": "active",
    "signup_payment_id": 191819284,
    "signup_revenue": "0.00",
    "delayed_cancel_at": null,
    "coupon_code": null,
    "total_revenue_in_cents": 0,
    "product_price_in_cents": 0,
    "product_version_number": 1,
    "payment_type": null,
    "referral_code": "d3pw7f",
    "coupon_use_count": null,
    "coupon_uses_allowed": null,
    "reason_code": null,
    "automatically_resume_at": null,
    "current_billing_amount_in_cents": 10000,
    "customer": {
      "id": 17780587,
      "first_name": "Catie",
      "last_name": "Test",
      "organization": "Acme, Inc.",
      "email": "catie@example.com",
      "created_at": "2017-06-27T13:01:05-05:00",
      "updated_at": "2017-06-30T09:23:10-05:00",
      "reference": "123ABC",
      "address": "123 Anywhere Street",
      "address_2": "Apartment #10",
      "city": "Los Angeles",
      "state": "CA",
      "zip": "90210",
      "country": "US",
      "phone": "555-555-5555",
      "portal_invite_last_sent_at": "2017-06-27T13:45:16-05:00",
      "portal_invite_last_accepted_at": null,
      "verified": true,
      "portal_customer_created_at": "2017-06-27T13:01:08-05:00",
      "cc_emails": "support@example.com",
      "tax_exempt": true
    },
    "product": {
      "id": 4470347,
      "name": "Zero Dollar Product",
      "handle": "zero-dollar-product",
      "description": "",
      "accounting_code": "",
      "request_credit_card": true,
      "expiration_interval": null,
      "expiration_interval_unit": "never",
      "created_at": "2017-03-23T10:54:12-05:00",
      "updated_at": "2017-04-20T15:18:46-05:00",
      "price_in_cents": 0,
      "interval": 1,
      "interval_unit": "month",
      "initial_charge_in_cents": null,
      "trial_price_in_cents": null,
      "trial_interval": null,
      "trial_interval_unit": "month",
      "archived_at": null,
      "require_credit_card": false,
      "return_params": "",
      "taxable": false,
      "update_return_url": "",
      "tax_code": "",
      "initial_charge_after_trial": false,
      "version_number": 1,
      "update_return_params": "",
      "product_family": {
        "id": 997233,
        "name": "Acme Products",
        "description": "",
        "handle": "acme-products",
        "accounting_code": null
      },
      "public_signup_pages": [
        {
          "id": 316810,
          "return_url": "",
          "return_params": "",
          "url": "https://general-goods.chargify.com/subscribe/69x825m78v3d/zero-dollar-product"
        }
      ]
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Pause Subscription

This will place the subscription in the on_hold state and it will not renew.

## Limitations

You may not place a subscription on hold if the `next_billing` date is within 24 hours.

```go
PauseSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.PauseRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.PauseRequest`](../../doc/models/pause-request.md) | Body, Optional | Allows to pause a Subscription |

## Response Type

[`models.SubscriptionResponse`](../../doc/models/subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

bodyHoldAutomaticallyResumeAt, err := time.Parse(time.RFC3339, "2017-05-25T11:25:00Z")
if err != nil {
    log.Fatalln(err)
}
bodyHold := models.AutoResume{
    AutomaticallyResumeAt: models.NewOptional(models.ToPointer(bodyHoldAutomaticallyResumeAt)),
}

body := models.PauseRequest{
    Hold: models.ToPointer(bodyHold),
}

apiResponse, err := subscriptionStatusController.PauseSubscription(ctx, subscriptionId, &body)
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
  "subscription": {
    "id": 18220670,
    "state": "on_hold",
    "trial_started_at": null,
    "trial_ended_at": null,
    "activated_at": "2017-06-27T13:45:15-05:00",
    "created_at": "2017-06-27T13:45:13-05:00",
    "updated_at": "2017-06-30T09:26:50-05:00",
    "expires_at": null,
    "balance_in_cents": 10000,
    "current_period_ends_at": "2017-06-30T12:00:00-05:00",
    "next_assessment_at": "2017-06-30T12:00:00-05:00",
    "canceled_at": null,
    "cancellation_message": null,
    "next_product_id": null,
    "cancel_at_end_of_period": false,
    "payment_collection_method": "automatic",
    "snap_day": "end",
    "cancellation_method": null,
    "current_period_started_at": "2017-06-27T13:45:13-05:00",
    "previous_state": "active",
    "signup_payment_id": 191819284,
    "signup_revenue": "0.00",
    "delayed_cancel_at": null,
    "coupon_code": null,
    "total_revenue_in_cents": 0,
    "product_price_in_cents": 0,
    "product_version_number": 1,
    "payment_type": null,
    "referral_code": "d3pw7f",
    "coupon_use_count": null,
    "coupon_uses_allowed": null,
    "reason_code": null,
    "automatically_resume_at": null,
    "current_billing_amount_in_cents": 10000,
    "customer": {
      "id": 17780587,
      "first_name": "Catie",
      "last_name": "Test",
      "organization": "Acme, Inc.",
      "email": "catie@example.com",
      "created_at": "2017-06-27T13:01:05-05:00",
      "updated_at": "2017-06-30T09:23:10-05:00",
      "reference": "123ABC",
      "address": "123 Anywhere Street",
      "address_2": "Apartment #10",
      "city": "Los Angeles",
      "state": "CA",
      "zip": "90210",
      "country": "US",
      "phone": "555-555-5555",
      "portal_invite_last_sent_at": "2017-06-27T13:45:16-05:00",
      "portal_invite_last_accepted_at": null,
      "verified": true,
      "portal_customer_created_at": "2017-06-27T13:01:08-05:00",
      "cc_emails": "support@example.com",
      "tax_exempt": true
    },
    "product": {
      "id": 4470347,
      "name": "Zero Dollar Product",
      "handle": "zero-dollar-product",
      "description": "",
      "accounting_code": "",
      "request_credit_card": true,
      "expiration_interval": null,
      "expiration_interval_unit": "never",
      "created_at": "2017-03-23T10:54:12-05:00",
      "updated_at": "2017-04-20T15:18:46-05:00",
      "price_in_cents": 0,
      "interval": 1,
      "interval_unit": "month",
      "initial_charge_in_cents": null,
      "trial_price_in_cents": null,
      "trial_interval": null,
      "trial_interval_unit": "month",
      "archived_at": null,
      "require_credit_card": false,
      "return_params": "",
      "taxable": false,
      "update_return_url": "",
      "tax_code": "",
      "initial_charge_after_trial": false,
      "version_number": 1,
      "update_return_params": "",
      "product_family": {
        "id": 997233,
        "name": "Acme Products",
        "description": "",
        "handle": "acme-products",
        "accounting_code": null
      },
      "public_signup_pages": [
        {
          "id": 316810,
          "return_url": "",
          "return_params": "",
          "url": "https://general-goods.chargify.com/subscribe/69x825m78v3d/zero-dollar-product"
        }
      ]
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Update Automatic Subscription Resumption

Once a subscription has been paused / put on hold, you can update the date which was specified to automatically resume the subscription.

To update a subscription's resume date, use this method to change or update the `automatically_resume_at` date.

### Remove the resume date

Alternately, you can change the `automatically_resume_at` to `null` if you would like the subscription to not have a resume date.

```go
UpdateAutomaticSubscriptionResumption(
    ctx context.Context,
    subscriptionId int,
    body *models.PauseRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.PauseRequest`](../../doc/models/pause-request.md) | Body, Optional | Allows to pause a Subscription |

## Response Type

[`models.SubscriptionResponse`](../../doc/models/subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

bodyHoldAutomaticallyResumeAt, err := time.Parse(time.RFC3339, "2019-01-20T00:00:00")
if err != nil {
    log.Fatalln(err)
}
bodyHold := models.AutoResume{
    AutomaticallyResumeAt: models.NewOptional(models.ToPointer(bodyHoldAutomaticallyResumeAt)),
}

body := models.PauseRequest{
    Hold: models.ToPointer(bodyHold),
}

apiResponse, err := subscriptionStatusController.UpdateAutomaticSubscriptionResumption(ctx, subscriptionId, &body)
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
  "subscription": {
    "id": 20359140,
    "state": "on_hold",
    "trial_started_at": null,
    "trial_ended_at": null,
    "activated_at": "2018-01-05T17:15:50-06:00",
    "created_at": "2018-01-05T17:15:49-06:00",
    "updated_at": "2018-01-09T10:26:14-06:00",
    "expires_at": null,
    "balance_in_cents": 0,
    "current_period_ends_at": "2023-01-05T17:15:00-06:00",
    "next_assessment_at": "2023-01-05T17:15:00-06:00",
    "canceled_at": null,
    "cancellation_message": null,
    "next_product_id": null,
    "cancel_at_end_of_period": false,
    "payment_collection_method": "automatic",
    "snap_day": null,
    "cancellation_method": null,
    "current_period_started_at": "2018-01-05T17:15:49-06:00",
    "previous_state": "active",
    "signup_payment_id": 219829722,
    "signup_revenue": "100.00",
    "delayed_cancel_at": null,
    "coupon_code": null,
    "total_revenue_in_cents": 10009991,
    "product_price_in_cents": 10000,
    "product_version_number": 1,
    "payment_type": "credit_card",
    "referral_code": "8y7jqr",
    "coupon_use_count": null,
    "coupon_uses_allowed": null,
    "reason_code": null,
    "automatically_resume_at": "2019-01-20T00:00:00-06:00",
    "coupon_codes": [],
    "customer": {
      "id": 19948683,
      "first_name": "Vanessa",
      "last_name": "Test",
      "organization": "",
      "email": "vanessa@example.com",
      "created_at": "2018-01-05T17:15:49-06:00",
      "updated_at": "2018-01-05T17:15:51-06:00",
      "reference": null,
      "address": "123 Anywhere Ln",
      "address_2": "",
      "city": "Boston",
      "state": "MA",
      "zip": "02120",
      "country": "US",
      "phone": "555-555-1212",
      "portal_invite_last_sent_at": "2018-01-05T17:15:51-06:00",
      "portal_invite_last_accepted_at": null,
      "verified": null,
      "portal_customer_created_at": "2018-01-05T17:15:51-06:00",
      "cc_emails": null,
      "tax_exempt": false
    },
    "product": {
      "id": 4535643,
      "name": "Annual Product",
      "handle": "annual-product",
      "description": "",
      "accounting_code": "",
      "request_credit_card": true,
      "expiration_interval": null,
      "expiration_interval_unit": "never",
      "created_at": "2017-08-25T10:25:31-05:00",
      "updated_at": "2017-08-25T10:25:31-05:00",
      "price_in_cents": 10000,
      "interval": 12,
      "interval_unit": "month",
      "initial_charge_in_cents": null,
      "trial_price_in_cents": null,
      "trial_interval": null,
      "trial_interval_unit": "month",
      "archived_at": null,
      "require_credit_card": true,
      "return_params": "",
      "taxable": false,
      "update_return_url": "",
      "tax_code": "",
      "initial_charge_after_trial": false,
      "version_number": 1,
      "update_return_params": "",
      "product_family": {
        "id": 1025627,
        "name": "Acme Products",
        "description": "",
        "handle": "acme-products",
        "accounting_code": null
      },
      "public_signup_pages": []
    },
    "credit_card": {
      "id": 13826563,
      "first_name": "Bomb 3",
      "last_name": "Test",
      "masked_card_number": "XXXX-XXXX-XXXX-1",
      "card_type": "bogus",
      "expiration_month": 1,
      "expiration_year": 2028,
      "customer_id": 19948683,
      "current_vault": "bogus",
      "vault_token": "1",
      "billing_address": "123 Anywhere Lane",
      "billing_city": "Boston",
      "billing_state": "Ma",
      "billing_zip": "02120",
      "billing_country": "US",
      "customer_vault_token": null,
      "billing_address_2": "",
      "payment_type": "credit_card"
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Reactivate Subscription

Chargify offers the ability to reactivate a previously canceled subscription. For details on how the reactivation works, and how to reactivate subscriptions through the application, see [reactivation](https://chargify.zendesk.com/hc/en-us/articles/4407898737691).

**Please note: The term
"resume" is used also during another process in Chargify. This occurs when an on-hold subscription is "resumed". This returns the subscription to an active state.**

+ The response returns the subscription object in the `active` or `trialing` state.
+ The `canceled_at` and `cancellation_message` fields do not have values.
+ The method works for "Canceled" or "Trial Ended" subscriptions.
+ It will not work for items not marked as "Canceled", "Unpaid", or "Trial Ended".

## Resume the current billing period for a subscription

A subscription is considered "resumable" if you are attempting to reactivate within the billing period the subscription was canceled in.

A resumed subscription's billing date remains the same as before it was canceled. In other words, it does not start a new billing period. Payment may or may not be collected for a resumed subscription, depending on whether or not the subscription had a balance when it was canceled (for example, if it was canceled because of dunning).

Consider a subscription which was created on June 1st, and would renew on July 1st. The subscription is then canceled on June 15.

If a reactivation with `resume: true` were attempted _before_ what would have been the next billing date of July 1st, then Chargify would resume the subscription.

If a reactivation with `resume: true` were attempted _after_ what would have been the next billing date of July 1st, then Chargify would not resume the subscription, and instead it would be reactivated with a new billing period.

| Canceled | Reactivation | Resumable? |
|---|---|---|
| Jun 15 | June 28 | Yes |
| Jun 15 | July 2 | No |

## Reactivation Scenarios

### Reactivating Canceled Subscription While Preserving Balance

+ Given you have a product that costs $20
+ Given you have a canceled subscription to the $20 product
  + 1 charge should exist for $20
  + 1 payment should exist for $20
+ When the subscription has canceled due to dunning, it retained a negative balance of $20

#### Results

The resulting charges upon reactivation will be:

+ 1 charge for $20 for the new product

+ 1 charge for $20 for the balance due

+ Total charges = $40

+ The subscription will transition to active

+ The subscription balance will be zero

### Reactivating a Canceled Subscription With Coupon

+ Given you have a canceled subscription
+ It has no current period defined
+ You have a coupon code "EARLYBIRD"
+ The coupon is set to recur for 6 periods

PUT request sent to:
`https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?coupon_code=EARLYBIRD`

#### Results

+ The subscription will transition to active
+ The subscription should have applied a coupon with code "EARLYBIRD"

### Reactivating Canceled Subscription With a Trial, Without the include_trial Flag

+ Given you have a canceled subscription

+ The product associated with the subscription has a trial

+ PUT request to
  `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json`

#### Results

+ The subscription will transition to active

### Reactivating Canceled Subscription With Trial, With the include_trial Flag

+ Given you have a canceled subscription

+ The product associated with the subscription has a trial

+ Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?include_trial=1`

#### Results

+ The subscription will transition to trialing

### Reactivating Trial Ended Subscription

+ Given you have a trial_ended subscription
+ Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json`

#### Results

+ The subscription will transition to active

### Resuming a Canceled Subscription

+ Given you have a `canceled` subscription and it is resumable
+ Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume=true`

#### Results

+ The subscription will transition to active
+ The next billing date should not have changed

### Attempting to resume a subscription which is not resumable

+ Given you have a `canceled` subscription, and it is not resumable
+ Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume=true`

#### Results

+ The subscription will transition to active, with a new billing period.

### Attempting to resume but not reactivate a subscription which is not resumable

+ Given you have a `canceled` subscription, and it is not resumable
+ Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume[require_resume]=true`
+ The response status should be "422 UNPROCESSABLE ENTITY"
+ The subscription should be canceled with the following response

```
  {
    "errors": ["Request was 'resume only', but this subscription cannot be resumed."]
  }
```

#### Results

+ The subscription should remain `canceled`
+ The next billing date should not have changed

### Resuming Subscription Which Was Trialing

+ Given you have a `trial_ended` subscription, and it is resumable
+ And the subscription was canceled in the middle of a trial
+ And there is still time left on the trial
+ Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume=true`

#### Results

+ The subscription will transition to trialing
+ The next billing date should not have changed

### Resuming Subscription Which Was trial_ended

+ Given you have a `trial_ended` subscription, and it is resumable
+ Send a PUT request to `https://acme.chargify.com/subscriptions/{subscription_id}/reactivate.json?resume=true`

#### Results

+ The subscription will transition to active
+ The next billing date should not have changed
+ Any product-related charges should have been collected

```go
ReactivateSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.ReactivateSubscriptionRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.ReactivateSubscriptionRequest`](../../doc/models/reactivate-subscription-request.md) | Body, Optional | - |

## Response Type

[`models.SubscriptionResponse`](../../doc/models/subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

bodyCalendarBilling := models.ReactivationBilling{
    ReactivationCharge: models.ToPointer(models.ReactivationCharge("prorated")),
}

bodyResume := models.ReactivateSubscriptionRequestResumeContainer.FromBoolean(true)

body := models.ReactivateSubscriptionRequest{
    IncludeTrial:             models.ToPointer(true),
    PreserveBalance:          models.ToPointer(true),
    CouponCode:               models.ToPointer("10OFF"),
    UseCreditsAndPrepayments: models.ToPointer(true),
    CalendarBilling:          models.ToPointer(bodyCalendarBilling),
    Resume:                   models.ToPointer(bodyResume),
}

apiResponse, err := subscriptionStatusController.ReactivateSubscription(ctx, subscriptionId, &body)
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
  "subscription": {
    "id": 18220670,
    "state": "active",
    "trial_started_at": null,
    "trial_ended_at": null,
    "activated_at": "2017-06-27T13:45:15-05:00",
    "created_at": "2017-06-27T13:45:13-05:00",
    "updated_at": "2017-06-30T09:26:50-05:00",
    "expires_at": null,
    "balance_in_cents": 10000,
    "current_period_ends_at": "2017-06-30T12:00:00-05:00",
    "next_assessment_at": "2017-06-30T12:00:00-05:00",
    "canceled_at": null,
    "cancellation_message": null,
    "next_product_id": null,
    "cancel_at_end_of_period": false,
    "payment_collection_method": "automatic",
    "snap_day": "end",
    "cancellation_method": null,
    "current_period_started_at": "2017-06-27T13:45:13-05:00",
    "previous_state": "active",
    "signup_payment_id": 191819284,
    "signup_revenue": "0.00",
    "delayed_cancel_at": null,
    "coupon_code": null,
    "total_revenue_in_cents": 0,
    "product_price_in_cents": 0,
    "product_version_number": 1,
    "payment_type": null,
    "referral_code": "d3pw7f",
    "coupon_use_count": null,
    "coupon_uses_allowed": null,
    "reason_code": null,
    "automatically_resume_at": null,
    "current_billing_amount_in_cents": 10000,
    "customer": {
      "id": 17780587,
      "first_name": "Catie",
      "last_name": "Test",
      "organization": "Acme, Inc.",
      "email": "catie@example.com",
      "created_at": "2017-06-27T13:01:05-05:00",
      "updated_at": "2017-06-30T09:23:10-05:00",
      "reference": "123ABC",
      "address": "123 Anywhere Street",
      "address_2": "Apartment #10",
      "city": "Los Angeles",
      "state": "CA",
      "zip": "90210",
      "country": "US",
      "phone": "555-555-5555",
      "portal_invite_last_sent_at": "2017-06-27T13:45:16-05:00",
      "portal_invite_last_accepted_at": null,
      "verified": true,
      "portal_customer_created_at": "2017-06-27T13:01:08-05:00",
      "cc_emails": "support@example.com",
      "tax_exempt": true,
      "vat_number": "012345678"
    },
    "product": {
      "id": 4470347,
      "name": "Zero Dollar Product",
      "handle": "zero-dollar-product",
      "description": "",
      "accounting_code": "",
      "request_credit_card": true,
      "expiration_interval": null,
      "expiration_interval_unit": "never",
      "created_at": "2017-03-23T10:54:12-05:00",
      "updated_at": "2017-04-20T15:18:46-05:00",
      "price_in_cents": 0,
      "interval": 1,
      "interval_unit": "month",
      "initial_charge_in_cents": null,
      "trial_price_in_cents": null,
      "trial_interval": null,
      "trial_interval_unit": "month",
      "archived_at": null,
      "require_credit_card": false,
      "return_params": "",
      "taxable": false,
      "update_return_url": "",
      "tax_code": "",
      "initial_charge_after_trial": false,
      "version_number": 1,
      "update_return_params": "",
      "product_family": {
        "id": 997233,
        "name": "Acme Products",
        "description": "",
        "handle": "acme-products",
        "accounting_code": null
      },
      "public_signup_pages": [
        {
          "id": 316810,
          "return_url": "",
          "return_params": "",
          "url": "https://general-goods.chargify.com/subscribe/69x825m78v3d/zero-dollar-product"
        }
      ]
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Initiate Delayed Cancellation

Chargify offers the ability to cancel a subscription at the end of the current billing period. This period is set by its current product.

Requesting to cancel the subscription at the end of the period sets the `cancel_at_end_of_period` flag to true.

Note that you cannot set `cancel_at_end_of_period` at subscription creation, or if the subscription is past due.

```go
InitiateDelayedCancellation(
    ctx context.Context,
    subscriptionId int,
    body *models.CancellationRequest) (
    models.ApiResponse[models.DelayedCancellationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.CancellationRequest`](../../doc/models/cancellation-request.md) | Body, Optional | - |

## Response Type

[`models.DelayedCancellationResponse`](../../doc/models/delayed-cancellation-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

apiResponse, err := subscriptionStatusController.InitiateDelayedCancellation(ctx, subscriptionId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |


# Cancel Delayed Cancellation

Removing the delayed cancellation on a subscription will ensure that it doesn't get canceled at the end of the period that it is in. The request will reset the `cancel_at_end_of_period` flag to `false`.

This endpoint is idempotent. If the subscription was not set to cancel in the future, removing the delayed cancellation has no effect and the call will be successful.

```go
CancelDelayedCancellation(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.DelayedCancellationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |

## Response Type

[`models.DelayedCancellationResponse`](../../doc/models/delayed-cancellation-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

apiResponse, err := subscriptionStatusController.CancelDelayedCancellation(ctx, subscriptionId)
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
  "message": "This subscription will no longer be canceled"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |


# Cancel Dunning

If a subscription is currently in dunning, the subscription will be set to active and the active Dunner will be resolved.

```go
CancelDunning(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.SubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |

## Response Type

[`models.SubscriptionResponse`](../../doc/models/subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

apiResponse, err := subscriptionStatusController.CancelDunning(ctx, subscriptionId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Preview Renewal

The Chargify API allows you to preview a renewal by posting to the renewals endpoint. Renewal Preview is an object representing a subscription’s next assessment. You can retrieve it to see a snapshot of how much your customer will be charged on their next renewal.

The "Next Billing" amount and "Next Billing" date are already represented in the UI on each Subscriber's Summary. For more information, please see our documentation [here](https://chargify.zendesk.com/hc/en-us/articles/4407884887835#next-billing).

## Optional Component Fields

This endpoint is particularly useful due to the fact that it will return the computed billing amount for the base product and the components which are in use by a subscriber.

By default, the preview will include billing details for all components _at their **current** quantities_. This means:

* Current `allocated_quantity` for quantity-based components
* Current enabled/disabled status for on/off components
* Current metered usage `unit_balance` for metered components
* Current metric quantity value for events recorded thus far for events-based components

In the above statements, "current" means the quantity or value as of the call to the renewal preview endpoint. We do not predict end-of-period values for components, so metered or events-based usage may be less than it will eventually be at the end of the period.

Optionally, **you may provide your own custom quantities** for any component to see a billing preview for non-current quantities. This is accomplished by sending a request body with data under the `components` key. See the request body documentation below.

## Subscription Side Effects

You can request a `POST` to obtain this data from the endpoint without any side effects. Plain and simple, this will preview data, not log any changes against a subscription.

```go
PreviewRenewal(
    ctx context.Context,
    subscriptionId int,
    body *models.RenewalPreviewRequest) (
    models.ApiResponse[models.RenewalPreviewResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.RenewalPreviewRequest`](../../doc/models/renewal-preview-request.md) | Body, Optional | - |

## Response Type

[`models.RenewalPreviewResponse`](../../doc/models/renewal-preview-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

bodyComponents0ComponentId := models.RenewalPreviewComponentComponentIdContainer.FromNumber(10708)

bodyComponents0 := models.RenewalPreviewComponent{
    Quantity:     models.ToPointer(10000),
    ComponentId:  models.ToPointer(bodyComponents0ComponentId),
}

bodyComponents1ComponentId := models.RenewalPreviewComponentComponentIdContainer.FromString("handle:small-instance-hours")

bodyComponents1PricePointId := models.RenewalPreviewComponentPricePointIdContainer.FromNumber(8712)

bodyComponents1 := models.RenewalPreviewComponent{
    Quantity:     models.ToPointer(10000),
    ComponentId:  models.ToPointer(bodyComponents1ComponentId),
    PricePointId: models.ToPointer(bodyComponents1PricePointId),
}

bodyComponents2ComponentId := models.RenewalPreviewComponentComponentIdContainer.FromString("handle:large-instance-hours")

bodyComponents2PricePointId := models.RenewalPreviewComponentPricePointIdContainer.FromString("handle:startup-pricing")

bodyComponents2 := models.RenewalPreviewComponent{
    Quantity:     models.ToPointer(100),
    ComponentId:  models.ToPointer(bodyComponents2ComponentId),
    PricePointId: models.ToPointer(bodyComponents2PricePointId),
}

bodyComponents := []models.RenewalPreviewComponent{bodyComponents0, bodyComponents1, bodyComponents2}
body := models.RenewalPreviewRequest{
    Components: bodyComponents,
}

apiResponse, err := subscriptionStatusController.PreviewRenewal(ctx, subscriptionId, &body)
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
  "renewal_preview": {
    "next_assessment_at": "2017-03-13T12:50:55-04:00",
    "subtotal_in_cents": 6000,
    "total_tax_in_cents": 0,
    "total_discount_in_cents": 0,
    "total_in_cents": 6000,
    "existing_balance_in_cents": 0,
    "total_amount_due_in_cents": 6000,
    "uncalculated_taxes": false,
    "line_items": [
      {
        "transaction_type": "charge",
        "kind": "baseline",
        "amount_in_cents": 5000,
        "memo": "Gold Product (03/13/2017 - 04/13/2017)",
        "discount_amount_in_cents": 0,
        "taxable_amount_in_cents": 0,
        "product_id": 1,
        "product_handle": "gold-product",
        "product_name": "Gold Product",
        "period_range_start": "01/10/2024",
        "period_range_end": "02/10/2024"
      },
      {
        "transaction_type": "charge",
        "kind": "quantity_based_component",
        "amount_in_cents": 1000,
        "memo": "Quantity Component: 10 Quantity Components",
        "discount_amount_in_cents": 0,
        "taxable_amount_in_cents": 0,
        "component_id": 104,
        "component_handle": "quantity-component",
        "component_name": "Quantity Component",
        "period_range_start": "01/10/2024",
        "period_range_end": "02/10/2024"
      }
    ]
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

