# Subscription Products

```go
subscriptionProductsController := client.SubscriptionProductsController()
```

## Class Name

`SubscriptionProductsController`

## Methods

* [Migrate Subscription Product](../../doc/controllers/subscription-products.md#migrate-subscription-product)
* [Preview Subscription Product Migration](../../doc/controllers/subscription-products.md#preview-subscription-product-migration)


# Migrate Subscription Product

In order to create a migration, you must pass the `product_id` or `product_handle` in the object when you send a POST request. You may also pass either a `product_price_point_id` or `product_price_point_handle` to choose which price point the subscription is moved to. If no price point identifier is passed the subscription will be moved to the products default price point. The response will be the updated subscription.

## Valid Subscriptions

Subscriptions should be in the `active` or `trialing` state in order to be migrated.

(For backwards compatibility reasons, it is possible to migrate a subscription that is in the `trial_ended` state via the API, however this is not recommended.  Since `trial_ended` is an end-of-life state, the subscription should be canceled, the product changed, and then the subscription can be reactivated.)

## Migrations Documentation

Full documentation on how to record Migrations in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407898373531).

## Failed Migrations

One of the most common ways that a migration can fail is when the attempt is made to migrate a subscription to it's current product. Please be aware of this issue!

## Migration 3D Secure - Stripe

It may happen that a payment needs 3D Secure Authentication when the subscription is migrated to a new product; this is referred to in our help docs as a [post-authentication flow](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405177432077#psd2-flows-pre-authentication-and-post-authentication). The server returns `422 Unprocessable Entity` in this case with the following response:

```json
{
  "errors": [
    "Your card was declined. This transaction requires 3D secure authentication."
  ],
  "gateway_payment_id": "pi_1F0aGoJ2UDb3Q4av7zU3sHPh",
  "description": "This card requires 3D secure authentication. Redirect the customer to the URL from the action_link attribute to authenticate. Attach callback_url param to this URL if you want to be notified about the result of 3D Secure authentication. Attach redirect_url param to this URL if you want to redirect a customer back to your page after 3D Secure authentication. Example: https://mysite.chargify.com/3d-secure/pi_1FCm4RKDeye4C0XfbqquXRYm?one_time_token_id=128&callback_url=https://localhost:4000&redirect_url=https://yourpage.com will do a POST request to https://localhost:4000 after payment is authenticated and will redirect a customer to https://yourpage.com after 3DS authentication.",
  "action_link": "http://acme.chargify.com/3d-secure/pi_1F0aGoJ2UDb3Q4av7zU3sHPh?one_time_token_id=242"
}
```

To let the customer go through 3D Secure Authentication, they need to be redirected to the URL specified in `action_link`.
Optionally, you can specify `callback_url` parameter in the `action_link` URL if you’d like to be notified about the result of 3D Secure Authentication. The `callback_url` will return the following information:

- whether the authentication was successful (`success`)
- the gateway ID for the payment (`gateway_payment_id`)
- the subscription ID (`subscription_id`)

Lastly, you can also specify a `redirect_url` within the `action_link` URL if you’d like to redirect a customer back to your site.

It is not possible to use `action_link` in an iframe inside a custom application. You have to redirect the customer directly to the `action_link`, then, to be notified about the result, use `redirect_url` or `callback_url`.

The final URL that you send a customer to to complete 3D Secure may resemble the following, where the first half is the `action_link` and the second half contains a `redirect_url` and `callback_url`: `https://mysite.chargify.com/3d-secure/pi_1FCm4RKDeye4C0XfbqquXRYm?one_time_token_id=128&callback_url=https://localhost:4000&redirect_url=https://yourpage.com`

### Example Redirect Flow

You may wish to redirect customers to different pages depending on whether their SCA was performed successfully. Here's an example flow to use as a reference:

1. Create a migration via API; it requires 3DS
2. You receive a `gateway_payment_id` in the `action_link` along other params in the response.
3. Use this `gateway_payment_id` to, for example, connect with your internal resources or generate a session_id
4. Include 1 of those attributes inside the `callback_url` and `redirect_url` to be aware which “session” this applies to
5. Redirect the customer to the `action_link` with `callback_url` and `redirect_url` applied
6. After the customer finishes 3DS authentication, we let you know the result by making a request to applied `callback_url`.
7. After that, we redirect the customer to the `redirect_url`; at this point the result of authentication is known
8. Optionally, you can use the applied "msg" param in the `redirect_url` to determine whether it was successful or not.

```go
MigrateSubscriptionProduct(
    ctx context.Context,
    subscriptionId int,
    body *models.SubscriptionProductMigrationRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.SubscriptionProductMigrationRequest`](../../doc/models/subscription-product-migration-request.md) | Body, Optional | - |

## Response Type

[`models.SubscriptionResponse`](../../doc/models/subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

bodyMigration := models.SubscriptionProductMigration{
    ProductId:               models.ToPointer(3801242),
    IncludeCoupons:          models.ToPointer(true),
    PreservePeriod:          models.ToPointer(true),
}

body := models.SubscriptionProductMigrationRequest{
    Migration: bodyMigration,
}

apiResponse, err := subscriptionProductsController.MigrateSubscriptionProduct(ctx, subscriptionId, &body)
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
    "id": 15054201,
    "state": "trialing",
    "trial_started_at": "2016-11-03T13:43:36-04:00",
    "trial_ended_at": "2016-11-10T12:43:36-05:00",
    "activated_at": "2016-11-02T10:20:57-04:00",
    "created_at": "2016-11-02T10:20:55-04:00",
    "updated_at": "2016-11-03T13:43:36-04:00",
    "expires_at": null,
    "balance_in_cents": -13989,
    "current_period_ends_at": "2016-11-10T12:43:36-05:00",
    "next_assessment_at": "2016-11-10T12:43:36-05:00",
    "canceled_at": null,
    "cancellation_message": null,
    "next_product_id": null,
    "cancel_at_end_of_period": false,
    "payment_collection_method": "automatic",
    "snap_day": null,
    "cancellation_method": null,
    "current_period_started_at": "2016-11-03T13:43:35-04:00",
    "previous_state": "active",
    "signup_payment_id": 160680121,
    "signup_revenue": "0.00",
    "delayed_cancel_at": null,
    "coupon_code": null,
    "total_revenue_in_cents": 14000,
    "product_price_in_cents": 1000,
    "product_version_number": 6,
    "payment_type": "credit_card",
    "referral_code": "ghnhvy",
    "coupon_use_count": null,
    "coupon_uses_allowed": null,
    "customer": {
      "id": 14543792,
      "first_name": "Frankie",
      "last_name": "Test",
      "organization": null,
      "email": "testfrankie111@test.com",
      "created_at": "2016-11-02T10:20:55-04:00",
      "updated_at": "2016-11-02T10:20:58-04:00",
      "reference": null,
      "address": null,
      "address_2": null,
      "city": null,
      "state": null,
      "zip": null,
      "country": null,
      "phone": "5555551212",
      "portal_invite_last_sent_at": "2016-11-02T10:20:58-04:00",
      "portal_invite_last_accepted_at": null,
      "verified": false,
      "portal_customer_created_at": "2016-11-02T10:20:58-04:00",
      "cc_emails": null
    },
    "product": {
      "id": 3861800,
      "name": "Trial Product",
      "handle": "trial-product",
      "description": "Trial period with payment expected at end of trial.",
      "accounting_code": "",
      "request_credit_card": true,
      "expiration_interval": null,
      "expiration_interval_unit": "never",
      "created_at": "2016-07-08T09:53:55-04:00",
      "updated_at": "2016-09-05T13:00:36-04:00",
      "price_in_cents": 1000,
      "interval": 1,
      "interval_unit": "month",
      "initial_charge_in_cents": null,
      "trial_price_in_cents": 0,
      "trial_interval": 7,
      "trial_interval_unit": "day",
      "archived_at": null,
      "require_credit_card": true,
      "return_params": "",
      "taxable": false,
      "update_return_url": "",
      "initial_charge_after_trial": false,
      "version_number": 6,
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
          "id": 294791,
          "return_url": "",
          "return_params": "",
          "url": "https://general-goods.chargify.com/subscribe/xv52yrcc3byx/trial-product"
        }
      ]
    },
    "credit_card": {
      "id": 10088716,
      "first_name": "F",
      "last_name": "NB",
      "masked_card_number": "XXXX-XXXX-XXXX-1",
      "card_type": "bogus",
      "expiration_month": 1,
      "expiration_year": 2017,
      "customer_id": 14543792,
      "current_vault": "bogus",
      "vault_token": "1",
      "billing_address": "123 Montana Way",
      "billing_city": "Billings",
      "billing_state": "MT",
      "billing_zip": "59101",
      "billing_country": "US",
      "customer_vault_token": null,
      "billing_address_2": "Apt. 10",
      "payment_type": "credit_card"
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Preview Subscription Product Migration

## Previewing a future date

It is also possible to preview the migration for a date in the future, as long as it's still within the subscription's current billing period, by passing a `proration_date` along with the request (eg: `"proration_date": "2020-12-18T18:25:43.511Z"`).

This will calculate the prorated adjustment, charge, payment and credit applied values assuming the migration is done at that date in the future as opposed to right now.

```go
PreviewSubscriptionProductMigration(
    ctx context.Context,
    subscriptionId int,
    body *models.SubscriptionMigrationPreviewRequest) (
    models.ApiResponse[models.SubscriptionMigrationPreviewResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.SubscriptionMigrationPreviewRequest`](../../doc/models/subscription-migration-preview-request.md) | Body, Optional | - |

## Response Type

[`models.SubscriptionMigrationPreviewResponse`](../../doc/models/subscription-migration-preview-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

bodyMigration := models.SubscriptionMigrationPreviewOptions{
    IncludeCoupons:          models.ToPointer(true),
}

body := models.SubscriptionMigrationPreviewRequest{
    Migration: bodyMigration,
}

apiResponse, err := subscriptionProductsController.PreviewSubscriptionProductMigration(ctx, subscriptionId, &body)
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
  "migration": {
    "prorated_adjustment_in_cents": 0,
    "charge_in_cents": 5000,
    "payment_due_in_cents": 0,
    "credit_applied_in_cents": 0
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

