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
)

// PaymentProfilesController represents a controller struct.
type PaymentProfilesController struct {
    baseController
}

// NewPaymentProfilesController creates a new instance of PaymentProfilesController.
// It takes a baseController as a parameter and returns a pointer to the PaymentProfilesController.
func NewPaymentProfilesController(baseController baseController) *PaymentProfilesController {
    paymentProfilesController := PaymentProfilesController{baseController: baseController}
    return &paymentProfilesController
}

// CreatePaymentProfile takes context, body as parameters and
// returns an models.ApiResponse with models.PaymentProfileResponse data and
// an error if there was an issue with the request or response.
// Creates a payment profile for a customer.
// When you create a new payment profile for a customer via the API, it does not automatically make the profile current for any of the customer’s subscriptions. To use the payment profile as the default, you must set it explicitly for the subscription or subscription group.
// Select an option from the **Request Examples** drop-down on the right side of the portal to see examples of common scenarios for creating payment profiles. 
// Do not use real card information for testing. See the Sites articles that cover [testing your site setup](https://docs.maxio.com/hc/en-us/articles/24250712113165-Testing-Overview#testing-overview-0-0) for more details on testing in your sandbox.
// Note that collecting and sending raw card details in production requires [PCI compliance](https://docs.maxio.com/hc/en-us/articles/24183956938381-PCI-Compliance#pci-compliance-0-0) on your end. If your business is not PCI compliant, use [Chargify.js](https://docs.maxio.com/hc/en-us/articles/38163190843789-Chargify-js-Overview#chargify-js-overview-0-0) to collect credit card or bank account information.
// See the following articles to learn more about subscriptions and payments:
// + [Subscriber Payment Details](https://maxio.zendesk.com/hc/en-us/articles/24251599929613-Subscription-Summary-Payment-Details-Tab)
// + [Self Service Pages](https://maxio.zendesk.com/hc/en-us/articles/24261425318541-Self-Service-Pages) (Allows credit card updates by Subscriber)
// + [Public Signup Pages payment settings](https://maxio.zendesk.com/hc/en-us/articles/24261368332557-Individual-Page-Settings)
// + [Taxes](https://developers.chargify.com/docs/developer-docs/d2e9e34db740e-signups#taxes)
// + [Chargify.js](https://docs.maxio.com/hc/en-us/articles/38163190843789-Chargify-js-Overview)
// + [Chargify.js with GoCardless - minimal example](https://docs.maxio.com/hc/en-us/articles/38206331271693-Examples#h_01K0PJ15QQZKCER8CFK40MR6XJ)
// + [Chargify.js with GoCardless - full example](https://docs.maxio.com/hc/en-us/articles/38206331271693-Examples#h_01K0PJ15QR09JVHWW0MCA7HVJV)
// + [Chargify.js with Stripe Direct Debit - minimal example](https://docs.maxio.com/hc/en-us/articles/38206331271693-Examples#h_01K0PJ15QQFKKN8Z7B7DZ9AJS5)
// + [Chargify.js with Stripe Direct Debit - full example](https://docs.maxio.com/hc/en-us/articles/38206331271693-Examples#h_01K0PJ15QRECQQ4ECS3ZA55GY7)
// + [Chargify.js with Stripe BECS Direct Debit - minimal example](https://developers.chargify.com/docs/developer-docs/ZG9jOjE0NjAzNDIy-examples#minimal-example-with-sepa-or-becs-direct-debit-stripe-gateway)
// + [Chargify.js with Stripe BECS Direct Debit - full example](https://developers.chargify.com/docs/developer-docs/ZG9jOjE0NjAzNDIy-examples#full-example-with-sepa-direct-debit-stripe-gateway)
// + [Full documentation on GoCardless](https://maxio.zendesk.com/hc/en-us/articles/24176159136909-GoCardless)
// + [Full documentation on Stripe SEPA Direct Debit](https://maxio.zendesk.com/hc/en-us/articles/24176170430093-Stripe-SEPA-and-BECS-Direct-Debit)
// + [Full documentation on Stripe BECS Direct Debit](https://maxio.zendesk.com/hc/en-us/articles/24176170430093-Stripe-SEPA-and-BECS-Direct-Debit)
// + [Full documentation on Stripe BACS Direct Debit](https://maxio.zendesk.com/hc/en-us/articles/24176170430093-Stripe-SEPA-and-BECS-Direct-Debit)
// ## 3D Secure Authentication during payment profile creation.
// When a payment requires 3D Secure Authentication to adhear to Strong Customer Authentication (SCA) during payment profile creation, the request enters a [post-authentication flow](https://maxio.zendesk.com/hc/en-us/articles/24176278996493-Testing-Implementing-3D-Secure#psd2-flows-pre-authentication-and-post-authentication). In this case, a 422 Unprocessable Entity status is returned with the following response:
// ```json
// {
// "jsonapi": {
// "version": "1.0"
// },
// "errors": [
// {
// "title": "This card requires 3DSecure verification.",
// "detail": "This card requires 3D secure authentication. Redirect the customer to the URL from the action_link attribute to authenticate. Attach callback_url param to this URL if you want to be notified about the result of 3D Secure authentication. Attach redirect_url param to this URL if you want to redirect a customer back to your page after 3D Secure authentication. Example: https://checkout-test.chargifypay.test/3d-secure/checkout/pay_uerzhsxd5uhkbodx5jhvkg6yeu?one_time_token_id=93&callback_url=http://localhost:4000&redirect_url=https://yourpage.com will do a POST request to https://localhost:4000 after credit card is authenticated and will redirect a customer to https://yourpage.com after 3DS authentication.",
// "links": {
// "action_link": "https://checkout-test.chargifypay.test/3d-secure/checkout/pay_uerzhsxd5uhkbodx5jhvkg6yeu?one_time_token_id=93"
// }
// }
// ]
// }
// ```
// To let the customer go through 3D Secure Authentication, they need to be redirected to the URL specified in `action_link`.
// Optionally, you can specify the `callback_url` parameter in the `action_link` URL to receive notification about the result of 3D Secure Authentication.
// The `callback_url` will return the following information:
// - whether the authentication was successful (`success`)
// - the payment profile ID (`payment_profile_id`)
// You can also specify a `redirect_url` parameter in the `action_link` URL to redirect the customer back to your site.
// You cannot use action_link in an iframe inside a custom application. You must redirect the customer directly to the `action_link` and use the `redirect_url` or `callback_url` to be notified of the result.
// The final URL that you send a customer to complete 3D Secure may resemble the following, where the first half is the `action_link` and the second half contains a `redirect_url` and `callback_url`:
// `https://checkout-test.chargifypay.test/3d-secure/checkout/pay_uerzhsxd5uhkbodx5jhvkg6yeu?one_time_token_id=93&callback_url=http://localhost:4000&redirect_url=https://yourpage.com`
// ### Example Redirect Flow
// Here's an example flow to redirect customers to different pages depending on whether SCA was performed successfully:
// 1. Create a payment profile via the API; it requires 3DS.
// 2. You receive an `action_link` in the response.
// 3. Use this `action_link` to, for example, connect with your internal resources or generate a `session_id`.
// 4. Include one of those attributes inside the `callback_url` and `redirect_url` to be aware which “session” this applies to.
// 5. Redirect the customer to the `action_link` with `callback_url` and `redirect_url` applied
// 6. After the customer completes 3DS authentication, we notify you of the result via the applied `callback_url`.
// 7. After that, we redirect the customer to the `redirect_url`; at this point the result of authentication is known.
// 8. Optionally, you can use the applied "msg" param in the `redirect_url` to determine if the redirect was successful.
func (p *PaymentProfilesController) CreatePaymentProfile(
    ctx context.Context,
    body *models.CreatePaymentProfileRequest) (
    models.ApiResponse[models.PaymentProfileResponse],
    error) {
    req := p.prepareRequest(ctx, "POST", "/payment_profiles.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.PaymentProfileResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PaymentProfileResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListPaymentProfilesInput represents the input of the ListPaymentProfiles endpoint.
type ListPaymentProfilesInput struct {
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page       *int 
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage    *int 
    // The ID of the customer for which you wish to list payment profiles
    CustomerId *int 
}

// ListPaymentProfiles takes context, page, perPage, customerId as parameters and
// returns an models.ApiResponse with []models.PaymentProfileResponse data and
// an error if there was an issue with the request or response.
// This method will return all of the active `payment_profiles` for a Site, or for one Customer within a site.  If no payment profiles are found, this endpoint will return an empty array, not a 404.
func (p *PaymentProfilesController) ListPaymentProfiles(
    ctx context.Context,
    input ListPaymentProfilesInput) (
    models.ApiResponse[[]models.PaymentProfileResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/payment_profiles.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.CustomerId != nil {
        req.QueryParam("customer_id", *input.CustomerId)
    }
    var result []models.PaymentProfileResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.PaymentProfileResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadPaymentProfile takes context, paymentProfileId as parameters and
// returns an models.ApiResponse with models.PaymentProfileResponse data and
// an error if there was an issue with the request or response.
// Using the GET method you can retrieve a Payment Profile identified by its unique ID.
// Note that a different JSON object will be returned if the card method on file is a bank account.
// ### Response for Bank Account
// Example response for Bank Account:
// ```
// {
// "payment_profile": {
// "id": 10089892,
// "first_name": "Chester",
// "last_name": "Tester",
// "created_at": "2025-01-01T00:00:00-05:00",
// "updated_at": "2025-01-01T00:00:00-05:00",
// "customer_id": 14543792,
// "current_vault": "bogus",
// "vault_token": "0011223344",
// "billing_address": "456 Juniper Court",
// "billing_city": "Boulder",
// "billing_state": "CO",
// "billing_zip": "80302",
// "billing_country": "US",
// "customer_vault_token": null,
// "billing_address_2": "",
// "bank_name": "Bank of Kansas City",
// "masked_bank_routing_number": "XXXX6789",
// "masked_bank_account_number": "XXXX3344",
// "bank_account_type": "checking",
// "bank_account_holder_type": "personal",
// "payment_type": "bank_account",
// "site_gateway_setting_id": 1,
// "gateway_handle": null
// }
// }
// ```
func (p *PaymentProfilesController) ReadPaymentProfile(
    ctx context.Context,
    paymentProfileId int) (
    models.ApiResponse[models.PaymentProfileResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/payment_profiles/%v.json")
    req.AppendTemplateParams(paymentProfileId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {Message: "Not Found"},
    })
    
    var result models.PaymentProfileResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PaymentProfileResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdatePaymentProfile takes context, paymentProfileId, body as parameters and
// returns an models.ApiResponse with models.PaymentProfileResponse data and
// an error if there was an issue with the request or response.
// ## Partial Card Updates
// In the event that you are using the Authorize.net, Stripe, Cybersource, Forte or Braintree Blue payment gateways, you can update just the billing and contact information for a payment method. Note the lack of credit-card related data contained in the JSON payload.
// In this case, the following JSON is acceptable:
// ```
// {
// "payment_profile": {
// "first_name": "Kelly",
// "last_name": "Test",
// "billing_address": "789 Juniper Court",
// "billing_city": "Boulder",
// "billing_state": "CO",
// "billing_zip": "80302",
// "billing_country": "US",
// "billing_address_2": null
// }
// }
// ```
// The result will be that you have updated the billing information for the card, yet retained the original card number data.
// ## Specific notes on updating payment profiles
// - Merchants with **Authorize.net**, **Cybersource**, **Forte**, **Braintree Blue** or **Stripe** as their payment gateway can update their Customer’s credit cards without passing in the full credit card number and CVV.
// - If you are using **Authorize.net**, **Cybersource**, **Forte**, **Braintree Blue** or **Stripe**, Advanced Billing will ignore the credit card number and CVV when processing an update via the API, and attempt a partial update instead. If you wish to change the card number on a payment profile, you will need to create a new payment profile for the given customer.
// - A Payment Profile cannot be updated with the attributes of another type of Payment Profile. For example, if the payment profile you are attempting to update is a credit card, you cannot pass in bank account attributes (like `bank_account_number`), and vice versa.
// - Updating a payment profile directly will not trigger an attempt to capture a past-due balance. If this is the intent, update the card details via the Subscription instead.
// - If you are using Authorize.net or Stripe, you may elect to manually trigger a retry for a past due subscription after a partial update.
func (p *PaymentProfilesController) UpdatePaymentProfile(
    ctx context.Context,
    paymentProfileId int,
    body *models.UpdatePaymentProfileRequest) (
    models.ApiResponse[models.PaymentProfileResponse],
    error) {
    req := p.prepareRequest(ctx, "PUT", "/payment_profiles/%v.json")
    req.AppendTemplateParams(paymentProfileId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {Message: "Not Found"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorStringMapResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.PaymentProfileResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PaymentProfileResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteUnusedPaymentProfile takes context, paymentProfileId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Deletes an unused payment profile.
// If the payment profile is in use by one or more subscriptions or groups, a 422 and error message will be returned.
func (p *PaymentProfilesController) DeleteUnusedPaymentProfile(
    ctx context.Context,
    paymentProfileId int) (
    *http.Response,
    error) {
    req := p.prepareRequest(ctx, "DELETE", "/payment_profiles/%v.json")
    req.AppendTemplateParams(paymentProfileId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {Message: "Not Found"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// DeleteSubscriptionsPaymentProfile takes context, subscriptionId, paymentProfileId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This will delete a payment profile belonging to the customer on the subscription.
// + If the customer has multiple subscriptions, the payment profile will be removed from all of them.
// + If you delete the default payment profile for a subscription, you will need to specify another payment profile to be the default through the api, or either prompt the user to enter a card in the billing portal or on the self-service page, or visit the Payment Details tab on the subscription in the Admin UI and use the “Add New Credit Card” or “Make Active Payment Method” link, (depending on whether there are other cards present).
func (p *PaymentProfilesController) DeleteSubscriptionsPaymentProfile(
    ctx context.Context,
    subscriptionId int,
    paymentProfileId int) (
    *http.Response,
    error) {
    req := p.prepareRequest(
      ctx,
      "DELETE",
      "/subscriptions/%v/payment_profiles/%v.json",
    )
    req.AppendTemplateParams(subscriptionId, paymentProfileId)
    req.Authenticate(NewAuth("BasicAuth"))
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// VerifyBankAccount takes context, bankAccountId, body as parameters and
// returns an models.ApiResponse with models.BankAccountResponse data and
// an error if there was an issue with the request or response.
// Submit the two small deposit amounts the customer received in their bank account in order to verify the bank account. (Stripe only)
func (p *PaymentProfilesController) VerifyBankAccount(
    ctx context.Context,
    bankAccountId int,
    body *models.BankAccountVerificationRequest) (
    models.ApiResponse[models.BankAccountResponse],
    error) {
    req := p.prepareRequest(ctx, "PUT", "/bank_accounts/%v/verification.json")
    req.AppendTemplateParams(bankAccountId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.BankAccountResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.BankAccountResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSubscriptionGroupPaymentProfile takes context, uid, paymentProfileId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This will delete a Payment Profile belonging to a Subscription Group.
// **Note**: If the Payment Profile belongs to multiple Subscription Groups and/or Subscriptions, it will be removed from all of them.
func (p *PaymentProfilesController) DeleteSubscriptionGroupPaymentProfile(
    ctx context.Context,
    uid string,
    paymentProfileId int) (
    *http.Response,
    error) {
    req := p.prepareRequest(
      ctx,
      "DELETE",
      "/subscription_groups/%v/payment_profiles/%v.json",
    )
    req.AppendTemplateParams(uid, paymentProfileId)
    req.Authenticate(NewAuth("BasicAuth"))
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// ChangeSubscriptionDefaultPaymentProfile takes context, subscriptionId, paymentProfileId as parameters and
// returns an models.ApiResponse with models.PaymentProfileResponse data and
// an error if there was an issue with the request or response.
// This will change the default payment profile on the subscription to the existing payment profile with the id specified.
// You must elect to change the existing payment profile to a new payment profile ID in order to receive a satisfactory response from this endpoint.
func (p *PaymentProfilesController) ChangeSubscriptionDefaultPaymentProfile(
    ctx context.Context,
    subscriptionId int,
    paymentProfileId int) (
    models.ApiResponse[models.PaymentProfileResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "POST",
      "/subscriptions/%v/payment_profiles/%v/change_payment_profile.json",
    )
    req.AppendTemplateParams(subscriptionId, paymentProfileId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {Message: "Not Found"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.PaymentProfileResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PaymentProfileResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ChangeSubscriptionGroupDefaultPaymentProfile takes context, uid, paymentProfileId as parameters and
// returns an models.ApiResponse with models.PaymentProfileResponse data and
// an error if there was an issue with the request or response.
// This will change the default payment profile on the subscription group to the existing payment profile with the id specified.
// You must elect to change the existing payment profile to a new payment profile ID in order to receive a satisfactory response from this endpoint.
// The new payment profile must belong to the subscription group's customer, otherwise you will receive an error.
func (p *PaymentProfilesController) ChangeSubscriptionGroupDefaultPaymentProfile(
    ctx context.Context,
    uid string,
    paymentProfileId int) (
    models.ApiResponse[models.PaymentProfileResponse],
    error) {
    req := p.prepareRequest(
      ctx,
      "POST",
      "/subscription_groups/%v/payment_profiles/%v/change_payment_profile.json",
    )
    req.AppendTemplateParams(uid, paymentProfileId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.PaymentProfileResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PaymentProfileResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadOneTimeToken takes context, chargifyToken as parameters and
// returns an models.ApiResponse with models.GetOneTimeTokenRequest data and
// an error if there was an issue with the request or response.
// One Time Tokens aka Advanced Billing Tokens house the credit card or ACH (Authorize.Net or Stripe only) data for a customer.
// You can use One Time Tokens while creating a subscription or payment profile instead of passing all bank account or credit card data directly to a given API endpoint.
// To obtain a One Time Token you have to use [Chargify.js](https://docs.maxio.com/hc/en-us/articles/38163190843789-Chargify-js-Overview#chargify-js-overview-0-0).
func (p *PaymentProfilesController) ReadOneTimeToken(
    ctx context.Context,
    chargifyToken string) (
    models.ApiResponse[models.GetOneTimeTokenRequest],
    error) {
    req := p.prepareRequest(ctx, "GET", "/one_time_tokens/%v.json")
    req.AppendTemplateParams(chargifyToken)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.GetOneTimeTokenRequest
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.GetOneTimeTokenRequest](decoder)
    return models.NewApiResponse(result, resp), err
}

// SendRequestUpdatePaymentEmail takes context, subscriptionId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// You can send a "request payment update" email to the customer associated with the subscription.
// If you attempt to send a "request payment update" email more than five times within a 30-minute period, you will receive a `422` response with an error message in the body. This error message will indicate that the request has been rejected due to excessive attempts, and will provide instructions on how to resubmit the request.
// Additionally, if you attempt to send a "request payment update" email for a subscription that does not exist, you will receive a `404` error response. This error message will indicate that the subscription could not be found, and will provide instructions on how to correct the error and resubmit the request.
// These error responses are designed to prevent excessive or invalid requests, and to provide clear and helpful information to users who encounter errors during the request process.
func (p *PaymentProfilesController) SendRequestUpdatePaymentEmail(
    ctx context.Context,
    subscriptionId int) (
    *http.Response,
    error) {
    req := p.prepareRequest(
      ctx,
      "POST",
      "/subscriptions/%v/request_payment_profiles_update.json",
    )
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
