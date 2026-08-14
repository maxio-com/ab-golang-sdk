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

Migrates a subscription to a different product.

To create a migration, you must pass the `product_id` or `product_handle` in the object when you send a POST request. You can also pass either a `product_price_point_id` or `product_price_point_handle` to choose which price point the subscription is moved to. If no price point identifier is passed, the subscription is moved to the product's default price point. The response is the updated subscription.

## Valid Subscriptions

Subscriptions should be in the `active` or `trialing` state to be migrated.

(For backwards compatibility reasons, it is possible to migrate a subscription that is in the `trial_ended` state via the API, however this is not recommended.  Since `trial_ended` is an end-of-life state, the subscription should be canceled, the product changed, and then the subscription can be reactivated.)

For more information, see [Product Changes and Migrations](https://docs.maxio.com/hc/en-us/articles/24252069837581-Product-Changes-and-Migrations).

## Failed Migrations

Important note: One of the most common ways that a migration can fail is when the attempt is made to migrate a subscription to its current product.

## 3D Secure (3DS) Authentication post-authentication flow

When a payment requires 3DS Authentication to adhere to Strong Customer Authentication (SCA), the request enters a post-authentication flow where a 422 Unprocessable Entity status is returned with an action_link that will direct the customer through 3DS Authentication.

See the [3D Secure Post-Authentication Flow](https://docs.maxio.com/hc/en-us/articles/44277749524365-3D-Secure-Post-Authentication-Flow) article in the product documentation to learn how to manage the redirect flow.

```go
MigrateSubscriptionProduct(
    ctx context.Context,
    subscriptionId int,
    body *models.SubscriptionProductMigrationRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `body` | [`*models.SubscriptionProductMigrationRequest`](../../doc/models/subscription-product-migration-request.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SubscriptionResponse](../../doc/models/subscription-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

body := models.SubscriptionProductMigrationRequest{
    Migration:            models.SubscriptionProductMigration{
        ProductId:               models.ToPointer(3801242),
        IncludeTrial:            models.ToPointer(false),
        IncludeInitialCharge:    models.ToPointer(false),
        IncludeCoupons:          models.ToPointer(true),
        PreservePeriod:          models.ToPointer(true),
    },
}

apiResponse, err := subscriptionProductsController.MigrateSubscriptionProduct(ctx, subscriptionId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ErrorListResponse:
            log.Fatalln("ErrorListResponseException: ", typedErr)
        default:
            log.Fatalln(err)
    }
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

Previews the charges resulting from migrating a subscription to a different product.

## Previewing a future date

It is also possible to preview the migration for a date in the future, as long as it's still within the subscription's current billing period, by passing a `proration_date` along with the request (e.g., `"proration_date": "2020-12-18T18:25:43.511Z"`).

This will calculate the prorated adjustment, charge, payment and credit applied values assuming the migration is done at that date in the future as opposed to right now.

```go
PreviewSubscriptionProductMigration(
    ctx context.Context,
    subscriptionId int,
    body *models.SubscriptionMigrationPreviewRequest) (
    models.ApiResponse[models.SubscriptionMigrationPreviewResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `body` | [`*models.SubscriptionMigrationPreviewRequest`](../../doc/models/subscription-migration-preview-request.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SubscriptionMigrationPreviewResponse](../../doc/models/subscription-migration-preview-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

body := models.SubscriptionMigrationPreviewRequest{
    Migration:            models.SubscriptionMigrationPreviewOptions{
        IncludeTrial:            models.ToPointer(false),
        IncludeInitialCharge:    models.ToPointer(false),
        IncludeCoupons:          models.ToPointer(true),
        PreservePeriod:          models.ToPointer(false),
    },
}

apiResponse, err := subscriptionProductsController.PreviewSubscriptionProductMigration(ctx, subscriptionId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ErrorListResponse:
            log.Fatalln("ErrorListResponseException: ", typedErr)
        default:
            log.Fatalln(err)
    }
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

