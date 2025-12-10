# Payment Profiles

```go
paymentProfilesController := client.PaymentProfilesController()
```

## Class Name

`PaymentProfilesController`

## Methods

* [Create Payment Profile](../../doc/controllers/payment-profiles.md#create-payment-profile)
* [List Payment Profiles](../../doc/controllers/payment-profiles.md#list-payment-profiles)
* [Read Payment Profile](../../doc/controllers/payment-profiles.md#read-payment-profile)
* [Update Payment Profile](../../doc/controllers/payment-profiles.md#update-payment-profile)
* [Delete Unused Payment Profile](../../doc/controllers/payment-profiles.md#delete-unused-payment-profile)
* [Delete Subscriptions Payment Profile](../../doc/controllers/payment-profiles.md#delete-subscriptions-payment-profile)
* [Verify Bank Account](../../doc/controllers/payment-profiles.md#verify-bank-account)
* [Delete Subscription Group Payment Profile](../../doc/controllers/payment-profiles.md#delete-subscription-group-payment-profile)
* [Change Subscription Default Payment Profile](../../doc/controllers/payment-profiles.md#change-subscription-default-payment-profile)
* [Change Subscription Group Default Payment Profile](../../doc/controllers/payment-profiles.md#change-subscription-group-default-payment-profile)
* [Read One Time Token](../../doc/controllers/payment-profiles.md#read-one-time-token)
* [Send Request Update Payment Email](../../doc/controllers/payment-profiles.md#send-request-update-payment-email)


# Create Payment Profile

Creates a payment profile for a customer.

When you create a new payment profile for a customer via the API, it does not automatically make the profile current for any of the customer’s subscriptions. To use the payment profile as the default, you must set it explicitly for the subscription or subscription group.

Select an option from the **Request Examples** drop-down on the right side of the portal to see examples of common scenarios for creating payment profiles.

Do not use real card information for testing. See the Sites articles that cover [testing your site setup](https://docs.maxio.com/hc/en-us/articles/24250712113165-Testing-Overview#testing-overview-0-0) for more details on testing in your sandbox.

Note that collecting and sending raw card details in production requires [PCI compliance](https://docs.maxio.com/hc/en-us/articles/24183956938381-PCI-Compliance#pci-compliance-0-0) on your end. If your business is not PCI compliant, use [Chargify.js](https://docs.maxio.com/hc/en-us/articles/38163190843789-Chargify-js-Overview#chargify-js-overview-0-0) to collect credit card or bank account information.

See the following articles to learn more about subscriptions and payments:

+ [Subscriber Payment Details](https://maxio.zendesk.com/hc/en-us/articles/24251599929613-Subscription-Summary-Payment-Details-Tab)
+ [Self Service Pages](https://maxio.zendesk.com/hc/en-us/articles/24261425318541-Self-Service-Pages) (Allows credit card updates by Subscriber)
+ [Public Signup Pages payment settings](https://maxio.zendesk.com/hc/en-us/articles/24261368332557-Individual-Page-Settings)
+ [Taxes](https://developers.chargify.com/docs/developer-docs/d2e9e34db740e-signups#taxes)
+ [Chargify.js](https://docs.maxio.com/hc/en-us/articles/38163190843789-Chargify-js-Overview)
  + [Chargify.js with GoCardless - minimal example](https://docs.maxio.com/hc/en-us/articles/38206331271693-Examples#h_01K0PJ15QQZKCER8CFK40MR6XJ)
  + [Chargify.js with GoCardless - full example](https://docs.maxio.com/hc/en-us/articles/38206331271693-Examples#h_01K0PJ15QR09JVHWW0MCA7HVJV)
  + [Chargify.js with Stripe Direct Debit - minimal example](https://docs.maxio.com/hc/en-us/articles/38206331271693-Examples#h_01K0PJ15QQFKKN8Z7B7DZ9AJS5)
  + [Chargify.js with Stripe Direct Debit - full example](https://docs.maxio.com/hc/en-us/articles/38206331271693-Examples#h_01K0PJ15QRECQQ4ECS3ZA55GY7)
  + [Chargify.js with Stripe BECS Direct Debit - minimal example](https://developers.chargify.com/docs/developer-docs/ZG9jOjE0NjAzNDIy-examples#minimal-example-with-sepa-or-becs-direct-debit-stripe-gateway)
  + [Chargify.js with Stripe BECS Direct Debit - full example](https://developers.chargify.com/docs/developer-docs/ZG9jOjE0NjAzNDIy-examples#full-example-with-sepa-direct-debit-stripe-gateway)
+ [Full documentation on GoCardless](https://maxio.zendesk.com/hc/en-us/articles/24176159136909-GoCardless)
+ [Full documentation on Stripe SEPA Direct Debit](https://maxio.zendesk.com/hc/en-us/articles/24176170430093-Stripe-SEPA-and-BECS-Direct-Debit)
+ [Full documentation on Stripe BECS Direct Debit](https://maxio.zendesk.com/hc/en-us/articles/24176170430093-Stripe-SEPA-and-BECS-Direct-Debit)
+ [Full documentation on Stripe BACS Direct Debit](https://maxio.zendesk.com/hc/en-us/articles/24176170430093-Stripe-SEPA-and-BECS-Direct-Debit)

## 3D Secure Authentication during payment profile creation.

When a payment requires 3D Secure Authentication to adhear to Strong Customer Authentication (SCA) during payment profile creation, the request enters a [post-authentication flow](https://maxio.zendesk.com/hc/en-us/articles/24176278996493-Testing-Implementing-3D-Secure#psd2-flows-pre-authentication-and-post-authentication). In this case, a 422 Unprocessable Entity status is returned with the following response:

```json
{
    "jsonapi": {
        "version": "1.0"
    },
    "errors": [
        {
            "title": "This card requires 3DSecure verification.",
            "detail": "This card requires 3D secure authentication. Redirect the customer to the URL from the action_link attribute to authenticate. Attach callback_url param to this URL if you want to be notified about the result of 3D Secure authentication. Attach redirect_url param to this URL if you want to redirect a customer back to your page after 3D Secure authentication. Example: https://checkout-test.chargifypay.test/3d-secure/checkout/pay_uerzhsxd5uhkbodx5jhvkg6yeu?one_time_token_id=93&callback_url=http://localhost:4000&redirect_url=https://yourpage.com will do a POST request to https://localhost:4000 after credit card is authenticated and will redirect a customer to https://yourpage.com after 3DS authentication.",
            "links": {
                "action_link": "https://checkout-test.chargifypay.test/3d-secure/checkout/pay_uerzhsxd5uhkbodx5jhvkg6yeu?one_time_token_id=93"
            }
        }
    ]
}
```

To let the customer go through 3D Secure Authentication, they need to be redirected to the URL specified in `action_link`.

Optionally, you can specify the `callback_url` parameter in the `action_link` URL to receive notification about the result of 3D Secure Authentication.

The `callback_url` will return the following information:

- whether the authentication was successful (`success`)
- the payment profile ID (`payment_profile_id`)

You can also specify a `redirect_url` parameter in the `action_link` URL to redirect the customer back to your site.

You cannot use action_link in an iframe inside a custom application. You must redirect the customer directly to the `action_link` and use the `redirect_url` or `callback_url` to be notified of the result.

The final URL that you send a customer to complete 3D Secure may resemble the following, where the first half is the `action_link` and the second half contains a `redirect_url` and `callback_url`:

`https://checkout-test.chargifypay.test/3d-secure/checkout/pay_uerzhsxd5uhkbodx5jhvkg6yeu?one_time_token_id=93&callback_url=http://localhost:4000&redirect_url=https://yourpage.com`

### Example Redirect Flow

Here's an example flow to redirect customers to different pages depending on whether SCA was performed successfully:

1. Create a payment profile via the API; it requires 3DS.
2. You receive an `action_link` in the response.
3. Use this `action_link` to, for example, connect with your internal resources or generate a `session_id`.
4. Include one of those attributes inside the `callback_url` and `redirect_url` to be aware which “session” this applies to.
5. Redirect the customer to the `action_link` with `callback_url` and `redirect_url` applied
6. After the customer completes 3DS authentication, we notify you of the result via the applied `callback_url`.
7. After that, we redirect the customer to the `redirect_url`; at this point the result of authentication is known.
8. Optionally, you can use the applied "msg" param in the `redirect_url` to determine if the redirect was successful.

```go
CreatePaymentProfile(
    ctx context.Context,
    body *models.CreatePaymentProfileRequest) (
    models.ApiResponse[models.PaymentProfileResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.CreatePaymentProfileRequest`](../../doc/models/create-payment-profile-request.md) | Body, Optional | When following the IBAN or the Local Bank details examples, a customer, bank account and mandate will be created in your current vault. If the customer, bank account, and mandate already exist in your vault, follow the Import example to link the payment profile into Advanced Billing. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.PaymentProfileResponse](../../doc/models/payment-profile-response.md).

## Example Usage

```go
ctx := context.Background()

body := models.CreatePaymentProfileRequest{
    PaymentProfile:       models.CreatePaymentProfile{
        ChargifyToken:         models.ToPointer("tok_w68qcpnftyv53jk33jv6wk3w"),
        CustomerId:            models.ToPointer(1036),
    },
}

apiResponse, err := paymentProfilesController.CreatePaymentProfile(ctx, &body)
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
  "payment_profile": {
    "first_name": "Jessica",
    "last_name": "Test",
    "card_type": "visa",
    "masked_card_number": "XXXX-XXXX-XXXX-1111",
    "expiration_month": 10,
    "expiration_year": 2018,
    "customer_id": 19195410,
    "current_vault": "bogus",
    "vault_token": "1",
    "billing_address": "123 Main St.",
    "billing_city": "Boston",
    "billing_state": "MA",
    "billing_zip": "02120",
    "billing_country": "US",
    "customer_vault_token": null,
    "billing_address_2": null,
    "payment_type": "credit_card",
    "site_gateway_setting_id": 1,
    "gateway_handle": "handle",
    "disabled": false
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Payment Profiles

This method will return all of the active `payment_profiles` for a Site, or for one Customer within a site.  If no payment profiles are found, this endpoint will return an empty array, not a 404.

```go
ListPaymentProfiles(
    ctx context.Context,
    input ListPaymentProfilesInput) (
    models.ApiResponse[[]models.PaymentProfileResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `customerId` | `*int` | Query, Optional | The ID of the customer for which you wish to list payment profiles |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.PaymentProfileResponse](../../doc/models/payment-profile-response.md).

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListPaymentProfilesInput{
    Page:       models.ToPointer(1),
    PerPage:    models.ToPointer(50),
}

apiResponse, err := paymentProfilesController.ListPaymentProfiles(ctx, collectedInput)
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
[
  {
    "payment_profile": {
      "id": 10089892,
      "first_name": "Chester",
      "last_name": "Tester",
      "created_at": "2025-01-01T00:00:00-05:00",
      "updated_at": "2025-01-01T00:00:00-05:00",
      "customer_id": 14543792,
      "current_vault": "bogus",
      "vault_token": "0011223344",
      "billing_address": "456 Juniper Court",
      "billing_city": "Boulder",
      "billing_state": "CO",
      "billing_zip": "80302",
      "billing_country": "US",
      "customer_vault_token": null,
      "billing_address_2": "",
      "bank_name": "Bank of Kansas City",
      "masked_bank_routing_number": "XXXX6789",
      "masked_bank_account_number": "XXXX3344",
      "bank_account_type": "checking",
      "bank_account_holder_type": "personal",
      "payment_type": "bank_account",
      "verified": true,
      "site_gateway_setting_id": 1,
      "gateway_handle": "handle"
    }
  },
  {
    "payment_profile": {
      "id": 10188522,
      "first_name": "Frankie",
      "last_name": "Tester",
      "created_at": "2025-01-01T00:00:00-05:00",
      "updated_at": "2025-01-01T00:00:00-05:00",
      "customer_id": 14543712,
      "current_vault": "bogus",
      "vault_token": "123456789",
      "billing_address": "123 Montana Way",
      "billing_city": "Los Angeles",
      "billing_state": "CA",
      "billing_zip": "90210",
      "billing_country": "US",
      "customer_vault_token": null,
      "billing_address_2": "",
      "bank_name": "Bank of Kansas City",
      "masked_bank_routing_number": "XXXX6789",
      "masked_bank_account_number": "XXXX6789",
      "bank_account_type": "checking",
      "bank_account_holder_type": "personal",
      "payment_type": "bank_account",
      "verified": true,
      "site_gateway_setting_id": 1,
      "gateway_handle": "handle"
    }
  }
]
```


# Read Payment Profile

Using the GET method you can retrieve a Payment Profile identified by its unique ID.

Note that a different JSON object will be returned if the card method on file is a bank account.

### Response for Bank Account

Example response for Bank Account:

```
{
  "payment_profile": {
    "id": 10089892,
    "first_name": "Chester",
    "last_name": "Tester",
    "created_at": "2025-01-01T00:00:00-05:00",
    "updated_at": "2025-01-01T00:00:00-05:00",
    "customer_id": 14543792,
    "current_vault": "bogus",
    "vault_token": "0011223344",
    "billing_address": "456 Juniper Court",
    "billing_city": "Boulder",
    "billing_state": "CO",
    "billing_zip": "80302",
    "billing_country": "US",
    "customer_vault_token": null,
    "billing_address_2": "",
    "bank_name": "Bank of Kansas City",
    "masked_bank_routing_number": "XXXX6789",
    "masked_bank_account_number": "XXXX3344",
    "bank_account_type": "checking",
    "bank_account_holder_type": "personal",
    "payment_type": "bank_account",
    "site_gateway_setting_id": 1,
    "gateway_handle": null
  }
}
```

```go
ReadPaymentProfile(
    ctx context.Context,
    paymentProfileId int) (
    models.ApiResponse[models.PaymentProfileResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `paymentProfileId` | `int` | Template, Required | The Chargify id of the payment profile |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.PaymentProfileResponse](../../doc/models/payment-profile-response.md).

## Example Usage

```go
ctx := context.Background()

paymentProfileId := 198

apiResponse, err := paymentProfilesController.ReadPaymentProfile(ctx, paymentProfileId)
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
  "payment_profile": {
    "id": 10088716,
    "first_name": "Test",
    "last_name": "Subscription",
    "masked_card_number": "XXXX-XXXX-XXXX-1",
    "card_type": "bogus",
    "expiration_month": 1,
    "expiration_year": 2022,
    "created_at": "2025-01-01T00:00:00-05:00",
    "updated_at": "2025-01-01T00:00:00-05:00",
    "customer_id": 14543792,
    "current_vault": "bogus",
    "vault_token": "1",
    "billing_address": "123 Montana Way",
    "billing_city": "Billings",
    "billing_state": "MT",
    "billing_zip": "59101",
    "billing_country": "US",
    "customer_vault_token": null,
    "billing_address_2": "",
    "payment_type": "credit_card",
    "site_gateway_setting_id": 1,
    "gateway_handle": null
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |


# Update Payment Profile

## Partial Card Updates

In the event that you are using the Authorize.net, Stripe, Cybersource, Forte or Braintree Blue payment gateways, you can update just the billing and contact information for a payment method. Note the lack of credit-card related data contained in the JSON payload.

In this case, the following JSON is acceptable:

```
{
  "payment_profile": {
    "first_name": "Kelly",
    "last_name": "Test",
    "billing_address": "789 Juniper Court",
    "billing_city": "Boulder",
    "billing_state": "CO",
    "billing_zip": "80302",
    "billing_country": "US",
    "billing_address_2": null
  }
}
```

The result will be that you have updated the billing information for the card, yet retained the original card number data.

## Specific notes on updating payment profiles

- Merchants with **Authorize.net**, **Cybersource**, **Forte**, **Braintree Blue** or **Stripe** as their payment gateway can update their Customer’s credit cards without passing in the full credit card number and CVV.

- If you are using **Authorize.net**, **Cybersource**, **Forte**, **Braintree Blue** or **Stripe**, Advanced Billing will ignore the credit card number and CVV when processing an update via the API, and attempt a partial update instead. If you wish to change the card number on a payment profile, you will need to create a new payment profile for the given customer.

- A Payment Profile cannot be updated with the attributes of another type of Payment Profile. For example, if the payment profile you are attempting to update is a credit card, you cannot pass in bank account attributes (like `bank_account_number`), and vice versa.

- Updating a payment profile directly will not trigger an attempt to capture a past-due balance. If this is the intent, update the card details via the Subscription instead.

- If you are using Authorize.net or Stripe, you may elect to manually trigger a retry for a past due subscription after a partial update.

```go
UpdatePaymentProfile(
    ctx context.Context,
    paymentProfileId int,
    body *models.UpdatePaymentProfileRequest) (
    models.ApiResponse[models.PaymentProfileResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `paymentProfileId` | `int` | Template, Required | The Chargify id of the payment profile |
| `body` | [`*models.UpdatePaymentProfileRequest`](../../doc/models/update-payment-profile-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.PaymentProfileResponse](../../doc/models/payment-profile-response.md).

## Example Usage

```go
ctx := context.Background()

paymentProfileId := 198

body := models.UpdatePaymentProfileRequest{
    PaymentProfile:       models.UpdatePaymentProfile{
        FirstName:            models.ToPointer("Graham"),
        LastName:             models.ToPointer("Test"),
        BillingAddress:       models.ToPointer("456 Juniper Court"),
        BillingCity:          models.ToPointer("Boulder"),
        BillingState:         models.ToPointer("CO"),
        BillingZip:           models.ToPointer("80302"),
        BillingCountry:       models.ToPointer("US"),
        BillingAddress2:      models.NewOptional(models.ToPointer("billing_address_22")),
    },
}

apiResponse, err := paymentProfilesController.UpdatePaymentProfile(ctx, paymentProfileId, &body)
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
  "payment_profile": {
    "id": 10088716,
    "first_name": "Test",
    "last_name": "Subscription",
    "billing_address": "123 Montana Way",
    "billing_city": "Billings",
    "billing_state": "MT",
    "billing_zip": "59101",
    "billing_country": "US",
    "billing_address_2": "",
    "payment_type": "bank_account"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorStringMapResponseException`](../../doc/models/error-string-map-response-exception.md) |


# Delete Unused Payment Profile

Deletes an unused payment profile.

If the payment profile is in use by one or more subscriptions or groups, a 422 and error message will be returned.

```go
DeleteUnusedPaymentProfile(
    ctx context.Context,
    paymentProfileId int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `paymentProfileId` | `int` | Template, Required | The Chargify id of the payment profile |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

paymentProfileId := 198

resp, err := paymentProfilesController.DeleteUnusedPaymentProfile(ctx, paymentProfileId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Delete Subscriptions Payment Profile

This will delete a payment profile belonging to the customer on the subscription.

+ If the customer has multiple subscriptions, the payment profile will be removed from all of them.

+ If you delete the default payment profile for a subscription, you will need to specify another payment profile to be the default through the api, or either prompt the user to enter a card in the billing portal or on the self-service page, or visit the Payment Details tab on the subscription in the Admin UI and use the “Add New Credit Card” or “Make Active Payment Method” link, (depending on whether there are other cards present).

```go
DeleteSubscriptionsPaymentProfile(
    ctx context.Context,
    subscriptionId int,
    paymentProfileId int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `paymentProfileId` | `int` | Template, Required | The Chargify id of the payment profile |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

paymentProfileId := 198

resp, err := paymentProfilesController.DeleteSubscriptionsPaymentProfile(ctx, subscriptionId, paymentProfileId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Verify Bank Account

Submit the two small deposit amounts the customer received in their bank account in order to verify the bank account. (Stripe only)

```go
VerifyBankAccount(
    ctx context.Context,
    bankAccountId int,
    body *models.BankAccountVerificationRequest) (
    models.ApiResponse[models.BankAccountResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `bankAccountId` | `int` | Template, Required | Identifier of the bank account in the system. |
| `body` | [`*models.BankAccountVerificationRequest`](../../doc/models/bank-account-verification-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.BankAccountResponse](../../doc/models/bank-account-response.md).

## Example Usage

```go
ctx := context.Background()

bankAccountId := 252

body := models.BankAccountVerificationRequest{
    BankAccountVerification: models.BankAccountVerification{
        Deposit1InCents:      models.ToPointer(int64(32)),
        Deposit2InCents:      models.ToPointer(int64(45)),
    },
}

apiResponse, err := paymentProfilesController.VerifyBankAccount(ctx, bankAccountId, &body)
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
  "payment_profile": {
    "id": 10089892,
    "first_name": "John",
    "last_name": "Doe",
    "customer_id": 14543792,
    "current_vault": "stripe_connect",
    "vault_token": "cus_0123abc456def",
    "billing_address": "456 Juniper Court",
    "billing_city": "Boulder",
    "billing_state": "CO",
    "billing_zip": "80302",
    "billing_country": "US",
    "customer_vault_token": null,
    "billing_address_2": "",
    "bank_name": "Bank of Kansas City",
    "masked_bank_routing_number": "XXXX6789",
    "masked_bank_account_number": "XXXX3344",
    "bank_account_type": "checking",
    "bank_account_holder_type": "personal",
    "payment_type": "bank_account"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Delete Subscription Group Payment Profile

This will delete a Payment Profile belonging to a Subscription Group.

**Note**: If the Payment Profile belongs to multiple Subscription Groups and/or Subscriptions, it will be removed from all of them.

```go
DeleteSubscriptionGroupPaymentProfile(
    ctx context.Context,
    uid string,
    paymentProfileId int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |
| `paymentProfileId` | `int` | Template, Required | The Chargify id of the payment profile |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

paymentProfileId := 198

resp, err := paymentProfilesController.DeleteSubscriptionGroupPaymentProfile(ctx, uid, paymentProfileId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Change Subscription Default Payment Profile

This will change the default payment profile on the subscription to the existing payment profile with the id specified.

You must elect to change the existing payment profile to a new payment profile ID in order to receive a satisfactory response from this endpoint.

```go
ChangeSubscriptionDefaultPaymentProfile(
    ctx context.Context,
    subscriptionId int,
    paymentProfileId int) (
    models.ApiResponse[models.PaymentProfileResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `paymentProfileId` | `int` | Template, Required | The Chargify id of the payment profile |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.PaymentProfileResponse](../../doc/models/payment-profile-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

paymentProfileId := 198

apiResponse, err := paymentProfilesController.ChangeSubscriptionDefaultPaymentProfile(ctx, subscriptionId, paymentProfileId)
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
  "payment_profile": {
    "id": 10211899,
    "first_name": "Amelia",
    "last_name": "Example",
    "masked_card_number": "XXXX-XXXX-XXXX-1",
    "card_type": "bogus",
    "expiration_month": 2,
    "expiration_year": 2018,
    "customer_id": 14399371,
    "current_vault": "bogus",
    "vault_token": "1",
    "billing_address": "",
    "billing_city": "",
    "billing_state": "",
    "billing_zip": "",
    "billing_country": "",
    "customer_vault_token": null,
    "billing_address_2": "",
    "payment_type": "credit_card",
    "site_gateway_setting_id": 1,
    "gateway_handle": null
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Change Subscription Group Default Payment Profile

This will change the default payment profile on the subscription group to the existing payment profile with the id specified.

You must elect to change the existing payment profile to a new payment profile ID in order to receive a satisfactory response from this endpoint.

The new payment profile must belong to the subscription group's customer, otherwise you will receive an error.

```go
ChangeSubscriptionGroupDefaultPaymentProfile(
    ctx context.Context,
    uid string,
    paymentProfileId int) (
    models.ApiResponse[models.PaymentProfileResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |
| `paymentProfileId` | `int` | Template, Required | The Chargify id of the payment profile |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.PaymentProfileResponse](../../doc/models/payment-profile-response.md).

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

paymentProfileId := 198

apiResponse, err := paymentProfilesController.ChangeSubscriptionGroupDefaultPaymentProfile(ctx, uid, paymentProfileId)
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
  "payment_profile": {
    "id": 10211899,
    "first_name": "Amelia",
    "last_name": "Example",
    "masked_card_number": "XXXX-XXXX-XXXX-1",
    "card_type": "bogus",
    "expiration_month": 2,
    "expiration_year": 2018,
    "customer_id": 14399371,
    "current_vault": "bogus",
    "vault_token": "1",
    "billing_address": "",
    "billing_city": "",
    "billing_state": "",
    "billing_zip": "",
    "billing_country": "",
    "customer_vault_token": null,
    "billing_address_2": "",
    "payment_type": "credit_card",
    "site_gateway_setting_id": 1,
    "gateway_handle": null
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Read One Time Token

One Time Tokens aka Advanced Billing Tokens house the credit card or ACH (Authorize.Net or Stripe only) data for a customer.

You can use One Time Tokens while creating a subscription or payment profile instead of passing all bank account or credit card data directly to a given API endpoint.

To obtain a One Time Token you have to use [Chargify.js](https://docs.maxio.com/hc/en-us/articles/38163190843789-Chargify-js-Overview#chargify-js-overview-0-0).

```go
ReadOneTimeToken(
    ctx context.Context,
    chargifyToken string) (
    models.ApiResponse[models.GetOneTimeTokenRequest],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `chargifyToken` | `string` | Template, Required | Advanced Billing Token |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.GetOneTimeTokenRequest](../../doc/models/get-one-time-token-request.md).

## Example Usage

```go
ctx := context.Background()

chargifyToken := "chargify_token8"

apiResponse, err := paymentProfilesController.ReadOneTimeToken(ctx, chargifyToken)
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
| 404 | Not Found | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Send Request Update Payment Email

You can send a "request payment update" email to the customer associated with the subscription.

If you attempt to send a "request payment update" email more than five times within a 30-minute period, you will receive a `422` response with an error message in the body. This error message will indicate that the request has been rejected due to excessive attempts, and will provide instructions on how to resubmit the request.

Additionally, if you attempt to send a "request payment update" email for a subscription that does not exist, you will receive a `404` error response. This error message will indicate that the subscription could not be found, and will provide instructions on how to correct the error and resubmit the request.

These error responses are designed to prevent excessive or invalid requests, and to provide clear and helpful information to users who encounter errors during the request process.

```go
SendRequestUpdatePaymentEmail(
    ctx context.Context,
    subscriptionId int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

resp, err := paymentProfilesController.SendRequestUpdatePaymentEmail(ctx, subscriptionId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

