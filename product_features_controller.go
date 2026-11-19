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

// ProductFeaturesController represents a controller struct.
type ProductFeaturesController struct {
    baseController
}

// NewProductFeaturesController creates a new instance of ProductFeaturesController.
// It takes a baseController as a parameter and returns a pointer to the ProductFeaturesController.
func NewProductFeaturesController(baseController baseController) *ProductFeaturesController {
    productFeaturesController := ProductFeaturesController{baseController: baseController}
    return &productFeaturesController
}

// ListProductFeatures takes context, productId as parameters and
// returns an models.ApiResponse with models.FeatureCatalogItemsListResponse data and
// an error if there was an issue with the request or response.
// Lists the feature catalog items attached to this product, including price-point-specific overrides.
func (p *ProductFeaturesController) ListProductFeatures(
    ctx context.Context,
    productId int) (
    models.ApiResponse[models.FeatureCatalogItemsListResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/products/%v/features.json")
    req.AppendTemplateParams(productId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
    })
    
    var result models.FeatureCatalogItemsListResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureCatalogItemsListResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateProductFeature takes context, productId, body as parameters and
// returns an models.ApiResponse with models.FeatureCatalogItemResponse data and
// an error if there was an issue with the request or response.
// Attaches a feature template to this product. Pass `price_point_type: "ProductPricePoint"` and `price_point_id` to scope the feature catalog item to a single product price point instead of the whole product. `value`, the periodicity and the grant expiration fall back to the feature template's defaults when omitted.
// **Product- or component-level items are a template.** A feature catalog item without `price_point_id` is copied onto price points created after it (except bulk-created price points; [Clone Component Price Point]($e/Component%3A%20Price%20Points/cloneComponentPricePoint) applies only the `feature_catalog_items` you send with it, and the automatic clone described below copies the source's features), but it is not added to price points that already exist, so their subscribers get nothing from it, even with `propagate_to_subscriptions: true`. To grant a feature to existing subscribers, create the item on each existing price point.
// **Restoring a template item.** Restoring an archived product- or component-level item, or creating one again for the same feature template, also restores the price-point copies that were archived together with it (for example, when the feature template was archived). If one of those price points has subscribers without the feature, the call returns `422` naming that price point, unless you send `propagate_to_subscriptions: true`, which restores the copies and grants the feature to those subscribers.
// **Price points with subscribers.** When the feature catalog item is scoped to a price point that already has subscribers and you don't send `propagate_to_subscriptions: true`, the new feature applies to new subscriptions only. The price point is cloned with its current features, the new feature is applied to the clone, and the response carries the clone's `price_point_id`. The feature catalog item in the response is also a different record, with a new `id`: later calls for the clone must use that `id`, not the original one. Existing subscribers stay on the original price point, unchanged; sign new subscriptions up on the clone to give them the change. A product price point has subscribers when any live subscription (including one in `awaiting_signup`) is on it. Price points that can't be cloned (archived or custom ones, or a component price point with an interval on a site without multi-frequency) make the request return `422`, with a message naming `propagate_to_subscriptions: true` as the way to apply the change to existing subscribers instead.
func (p *ProductFeaturesController) CreateProductFeature(
    ctx context.Context,
    productId int,
    body *models.CreateFeatureCatalogItemRequest) (
    models.ApiResponse[models.FeatureCatalogItemResponse],
    error) {
    req := p.prepareRequest(ctx, "POST", "/products/%v/features.json")
    req.AppendTemplateParams(productId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.FeatureCatalogItemResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureCatalogItemResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadProductFeature takes context, productId, id as parameters and
// returns an models.ApiResponse with models.FeatureCatalogItemResponse data and
// an error if there was an issue with the request or response.
// Returns a single feature catalog item attached to this product.
func (p *ProductFeaturesController) ReadProductFeature(
    ctx context.Context,
    productId int,
    id int) (
    models.ApiResponse[models.FeatureCatalogItemResponse],
    error) {
    req := p.prepareRequest(ctx, "GET", "/products/%v/features/%v.json")
    req.AppendTemplateParams(productId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
    })
    
    var result models.FeatureCatalogItemResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureCatalogItemResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateProductFeature takes context, productId, id, body as parameters and
// returns an models.ApiResponse with models.FeatureCatalogItemResponse data and
// an error if there was an issue with the request or response.
// Updates the value or periodicity of a feature catalog item attached to this product. A request that changes nothing returns the item as it is.
// **Price points with subscribers.** When the feature catalog item is scoped to a price point that already has subscribers and you don't send `propagate_to_subscriptions: true`, the update applies to new subscriptions only. The price point is cloned with its current features, the update is applied to the clone, and the response carries the clone's `price_point_id`. The feature catalog item in the response is also a different record, with a new `id`: later calls for the clone must use that `id`, not the original one. Existing subscribers stay on the original price point, unchanged; sign new subscriptions up on the clone to give them the change. A product price point has subscribers when any live subscription (including one in `awaiting_signup`) is on it. Price points that can't be cloned (archived or custom ones, or a component price point with an interval on a site without multi-frequency) make the request return `422`, with a message naming `propagate_to_subscriptions: true` as the way to apply the change to existing subscribers instead.
func (p *ProductFeaturesController) UpdateProductFeature(
    ctx context.Context,
    productId int,
    id int,
    body *models.UpdateFeatureCatalogItemRequest) (
    models.ApiResponse[models.FeatureCatalogItemResponse],
    error) {
    req := p.prepareRequest(ctx, "PUT", "/products/%v/features/%v.json")
    req.AppendTemplateParams(productId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.FeatureCatalogItemResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureCatalogItemResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// RemoveProductFeature takes context, productId, id, destroyEntitlements as parameters and
// returns an models.ApiResponse with models.FeatureCatalogItemResponse data and
// an error if there was an issue with the request or response.
// Removes a feature catalog item from this product. Returns `204` when the item is removed in place.
// **Price points with subscribers.** When the feature catalog item is scoped to a price point that already has subscribers and you don't send `destroy_entitlements=true`, the removal applies to new subscriptions only. The price point is cloned with its current features, the removal is applied to the clone, and the call returns `200` with the archived feature catalog item on the clone (carrying the clone's `price_point_id` and a new `id`) instead of `204`. Existing subscribers stay on the original price point, unchanged; sign new subscriptions up on the clone to give them the change. A product price point has subscribers when any live subscription (including one in `awaiting_signup`) is on it. Price points that can't be cloned (archived or custom ones, or a component price point with an interval on a site without multi-frequency) make the request return `422`, with a message naming `destroy_entitlements=true` as the way to apply the change to existing subscribers instead.
func (p *ProductFeaturesController) RemoveProductFeature(
    ctx context.Context,
    productId int,
    id int,
    destroyEntitlements *bool) (
    models.ApiResponse[models.FeatureCatalogItemResponse],
    error) {
    req := p.prepareRequest(ctx, "DELETE", "/products/%v/features/%v.json")
    req.AppendTemplateParams(productId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
        "422": {Message: "The price point has subscribers but can't be cloned (it is archived or custom, or it is a component price point with an interval on a site without multi-frequency). Send `destroy_entitlements=true` to remove the feature from its existing subscribers instead.", Unmarshaller: errors.NewErrorListResponse},
    })
    if destroyEntitlements != nil {
        req.QueryParam("destroy_entitlements", *destroyEntitlements)
    }
    
    var result models.FeatureCatalogItemResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureCatalogItemResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// RestoreProductFeature takes context, productId, id, propagateToSubscriptions as parameters and
// returns an models.ApiResponse with models.FeatureCatalogItemResponse data and
// an error if there was an issue with the request or response.
// Clears the archived state of a feature catalog item attached to this product. Returns `422` if the parent feature template is still archived. Restore the feature template first. Otherwise, restoring an item that isn't archived returns it unchanged.
// **Price points with subscribers.** When the feature catalog item is scoped to a price point that already has subscribers and you don't send `propagate_to_subscriptions=true`, the restore applies to new subscriptions only. The price point is cloned with its current features, the restore is applied to the clone, and the response carries the clone's `price_point_id`. The feature catalog item in the response is also a different record, with a new `id`: later calls for the clone must use that `id`, not the original one. Existing subscribers stay on the original price point, unchanged; sign new subscriptions up on the clone to give them the change. A product price point has subscribers when any live subscription (including one in `awaiting_signup`) is on it. Price points that can't be cloned (archived or custom ones, or a component price point with an interval on a site without multi-frequency) make the request return `422`, with a message naming `propagate_to_subscriptions: true` as the way to apply the change to existing subscribers instead.
func (p *ProductFeaturesController) RestoreProductFeature(
    ctx context.Context,
    productId int,
    id int,
    propagateToSubscriptions *bool) (
    models.ApiResponse[models.FeatureCatalogItemResponse],
    error) {
    req := p.prepareRequest(ctx, "POST", "/products/%v/features/%v/restore.json")
    req.AppendTemplateParams(productId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    if propagateToSubscriptions != nil {
        req.QueryParam("propagate_to_subscriptions", *propagateToSubscriptions)
    }
    
    var result models.FeatureCatalogItemResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureCatalogItemResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
