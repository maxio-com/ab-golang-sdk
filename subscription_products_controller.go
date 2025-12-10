// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// SubscriptionProductsController represents a controller struct.
type SubscriptionProductsController struct {
    baseController
}

// NewSubscriptionProductsController creates a new instance of SubscriptionProductsController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionProductsController.
func NewSubscriptionProductsController(baseController baseController) *SubscriptionProductsController {
    subscriptionProductsController := SubscriptionProductsController{baseController: baseController}
    return &subscriptionProductsController
}

// MigrateSubscriptionProduct takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// In order to create a migration, you must pass the `product_id` or `product_handle` in the object when you send a POST request. You may also pass either a `product_price_point_id` or `product_price_point_handle` to choose which price point the subscription is moved to. If no price point identifier is passed the subscription will be moved to the products default price point. The response will be the updated subscription.
// ## Valid Subscriptions
// Subscriptions should be in the `active` or `trialing` state in order to be migrated.
// (For backwards compatibility reasons, it is possible to migrate a subscription that is in the `trial_ended` state via the API, however this is not recommended.  Since `trial_ended` is an end-of-life state, the subscription should be canceled, the product changed, and then the subscription can be reactivated.)
// ## Migrations Documentation
// Full documentation on how to record Migrations in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24181589372429-Data-Migration-to-Advanced-Billing).
// ## Failed Migrations
// Importaint note: One of the most common ways that a migration can fail is when the attempt is made to migrate a subscription to its current product.
// ## Migration 3D Secure - Stripe
// When a payment requires 3D Secure Authentication to adhear to Strong Customer Authentication (SCA) when the subscription is migrated to a new product, the request enters a [post-authentication flow](https://maxio.zendesk.com/hc/en-us/articles/24176278996493-Testing-Implementing-3D-Secure#psd2-flows-pre-authentication-and-post-authentication). The server returns `422 Unprocessable Entity` in this case with the following response:
// ```json
// {
// "errors": [
// "Your card was declined. This transaction requires 3D secure authentication."
// ],
// "gateway_payment_id": "pi_1F0aGoJ2UDb3Q4av7zU3sHPh",
// "description": "This card requires 3D secure authentication. Redirect the customer to the URL from the action_link attribute to authenticate. Attach callback_url param to this URL if you want to be notified about the result of 3D Secure authentication. Attach redirect_url param to this URL if you want to redirect a customer back to your page after 3D Secure authentication. Example: https://mysite.chargify.com/3d-secure/pi_1FCm4RKDeye4C0XfbqquXRYm?one_time_token_id=128&callback_url=https://localhost:4000&redirect_url=https://yourpage.com will do a POST request to https://localhost:4000 after payment is authenticated and will redirect a customer to https://yourpage.com after 3DS authentication.",
// "action_link": "http://acme.chargify.com/3d-secure/pi_1F0aGoJ2UDb3Q4av7zU3sHPh?one_time_token_id=242"
// }
// ```
// To let the customer go through 3D Secure Authentication, they need to be redirected to the URL specified in `action_link`.
// Optionally, you can specify `callback_url` parameter in the `action_link` URL if you’d like to be notified about the result of 3D Secure Authentication. The `callback_url` will return the following information:
// - whether the authentication was successful (`success`)
// - the gateway ID for the payment (`gateway_payment_id`)
// - the subscription ID (`subscription_id`)
// Lastly, you can also specify a `redirect_url` within the `action_link` URL if you’d like to redirect a customer back to your site.
// It is not possible to use `action_link` in an iframe inside a custom application. You have to redirect the customer directly to the `action_link`, then, to be notified about the result, use `redirect_url` or `callback_url`.
// The final URL that you send a customer to to complete 3D Secure may resemble the following, where the first half is the `action_link` and the second half contains a `redirect_url` and `callback_url`: `https://mysite.chargify.com/3d-secure/pi_1FCm4RKDeye4C0XfbqquXRYm?one_time_token_id=128&callback_url=https://localhost:4000&redirect_url=https://yourpage.com`
// ### Example Redirect Flow
// You may wish to redirect customers to different pages depending on whether SCA was performed successfully. Here's an example flow to use as a reference:
// 1. Create a migration via API; it requires 3DS
// 2. You receive a `gateway_payment_id` in the `action_link` along other params in the response.
// 3. Use this `gateway_payment_id` to, for example, connect with your internal resources or generate a session_id
// 4. Include 1 of those attributes inside the `callback_url` and `redirect_url` to be aware which “session” this applies to
// 5. Redirect the customer to the `action_link` with `callback_url` and `redirect_url` applied
// 6. After the customer finishes 3DS authentication, we let you know the result by making a request to applied `callback_url`.
// 7. After that, we redirect the customer to the `redirect_url`; at this point the result of authentication is known
// 8. Optionally, you can use the applied "msg" param in the `redirect_url` to determine whether it was successful or not.
func (s *SubscriptionProductsController) MigrateSubscriptionProduct(
    ctx context.Context,
    subscriptionId int,
    body *models.SubscriptionProductMigrationRequest) (
    models.ApiResponse[models.SubscriptionResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscriptions/%v/migrations.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// PreviewSubscriptionProductMigration takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionMigrationPreviewResponse data and
// an error if there was an issue with the request or response.
// ## Previewing a future date
// It is also possible to preview the migration for a date in the future, as long as it's still within the subscription's current billing period, by passing a `proration_date` along with the request (eg: `"proration_date": "2020-12-18T18:25:43.511Z"`).
// This will calculate the prorated adjustment, charge, payment and credit applied values assuming the migration is done at that date in the future as opposed to right now.
func (s *SubscriptionProductsController) PreviewSubscriptionProductMigration(
    ctx context.Context,
    subscriptionId int,
    body *models.SubscriptionMigrationPreviewRequest) (
    models.ApiResponse[models.SubscriptionMigrationPreviewResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscriptions/%v/migrations/preview.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SubscriptionMigrationPreviewResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionMigrationPreviewResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
