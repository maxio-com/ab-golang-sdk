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
// Migrates a subscription to a different product.
// To create a migration, you must pass the `product_id` or `product_handle` in the object when you send a POST request. You can also pass either a `product_price_point_id` or `product_price_point_handle` to choose which price point the subscription is moved to. If no price point identifier is passed, the subscription is moved to the product's default price point. The response is the updated subscription.
// ## Valid Subscriptions
// Subscriptions should be in the `active` or `trialing` state to be migrated.
// (For backwards compatibility reasons, it is possible to migrate a subscription that is in the `trial_ended` state via the API, however this is not recommended.  Since `trial_ended` is an end-of-life state, the subscription should be canceled, the product changed, and then the subscription can be reactivated.)
// For more information, see [Product Changes and Migrations](https://docs.maxio.com/hc/en-us/articles/24252069837581-Product-Changes-and-Migrations).
// ## Failed Migrations
// Important note: One of the most common ways that a migration can fail is when the attempt is made to migrate a subscription to its current product. 
// ## 3D Secure (3DS) Authentication post-authentication flow
// When a payment requires 3DS Authentication to adhere to Strong Customer Authentication (SCA), the request enters a post-authentication flow where a 422 Unprocessable Entity status is returned with an action_link that will direct the customer through 3DS Authentication. 
// See the [3D Secure Post-Authentication Flow](https://docs.maxio.com/hc/en-us/articles/44277749524365-3D-Secure-Post-Authentication-Flow) article in the product documentation to learn how to manage the redirect flow.
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
// Previews the charges resulting from migrating a subscription to a different product.
// ## Previewing a future date
// It is also possible to preview the migration for a date in the future, as long as it's still within the subscription's current billing period, by passing a `proration_date` along with the request (e.g., `"proration_date": "2020-12-18T18:25:43.511Z"`).
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
