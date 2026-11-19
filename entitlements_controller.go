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

// EntitlementsController represents a controller struct.
type EntitlementsController struct {
    baseController
}

// NewEntitlementsController creates a new instance of EntitlementsController.
// It takes a baseController as a parameter and returns a pointer to the EntitlementsController.
func NewEntitlementsController(baseController baseController) *EntitlementsController {
    entitlementsController := EntitlementsController{baseController: baseController}
    return &entitlementsController
}

// ReadSubscriptionEntitlements takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.AggregatedEntitlementsResponse data and
// an error if there was an issue with the request or response.
// Returns every feature a subscription is entitled to, collapsed into one entry per feature key and periodicity window across all products and components on the subscription. A `usage_limit` feature granted with two different periodicities comes back as two entries sharing one `feature_key`, each identified by its own `periodicity_key`.
// When more than one product or component grants the same feature key and periodicity, the values are combined:
// - **`access_right`** features are combined with a boolean OR. If any contributor grants access, the aggregate is `true`. `source_products` only lists the contributors that granted `true`.
// - **`usage_limit`** features are summed across every contributor sharing the same periodicity window. `source_products` lists every contributor. Grants with different periodicities are not summed together. Each periodicity is returned as a separate entry.
// - **`service_right`** features are not combined: one contributor's value wins. Do not rely on which one when several grant the same feature key.
// `enabled` reflects both the aggregated value and the subscription's state. The field is `false` whenever the subscription is not in a live state (`active`, `trialing`, `assessing`, `past_due`, `soft_failure`), regardless of the aggregated value. Entitlements deliberately stay enabled through dunning.
func (e *EntitlementsController) ReadSubscriptionEntitlements(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.AggregatedEntitlementsResponse],
    error) {
    req := e.prepareRequest(ctx, "GET", "/subscriptions/%v/entitlements.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
    })
    
    var result models.AggregatedEntitlementsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AggregatedEntitlementsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
