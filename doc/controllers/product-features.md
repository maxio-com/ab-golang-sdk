# Product Features

```go
productFeaturesController := client.ProductFeaturesController()
```

## Class Name

`ProductFeaturesController`

## Methods

* [List Product Features](../../doc/controllers/product-features.md#list-product-features)
* [Create Product Feature](../../doc/controllers/product-features.md#create-product-feature)
* [Read Product Feature](../../doc/controllers/product-features.md#read-product-feature)
* [Update Product Feature](../../doc/controllers/product-features.md#update-product-feature)
* [Remove Product Feature](../../doc/controllers/product-features.md#remove-product-feature)
* [Restore Product Feature](../../doc/controllers/product-features.md#restore-product-feature)


# List Product Features

Lists the feature catalog items attached to this product, including price-point-specific overrides.

```go
ListProductFeatures(
    ctx context.Context,
    productId int) (
    models.ApiResponse[models.FeatureCatalogItemsListResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Advanced Billing id of the product. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.FeatureCatalogItemsListResponse](../../doc/models/feature-catalog-items-list-response.md).

## Example Usage

```go
ctx := context.Background()

productId := 202

apiResponse, err := productFeaturesController.ListProductFeatures(ctx, productId)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |
| 404 | Not Found | `ApiError` |


# Create Product Feature

Attaches a feature template to this product. Pass `price_point_type: "ProductPricePoint"` and `price_point_id` to scope the feature catalog item to a single product price point instead of the whole product. `value`, the periodicity and the grant expiration fall back to the feature template's defaults when omitted.

**Product- or component-level items are a template.** A feature catalog item without `price_point_id` is copied onto price points created after it (except bulk-created price points; [Clone Component Price Point](../../doc/controllers/component-price-points.md#clone-component-price-point) applies only the `feature_catalog_items` you send with it, and the automatic clone described below copies the source's features), but it is not added to price points that already exist, so their subscribers get nothing from it, even with `propagate_to_subscriptions: true`. To grant a feature to existing subscribers, create the item on each existing price point.

**Restoring a template item.** Restoring an archived product- or component-level item, or creating one again for the same feature template, also restores the price-point copies that were archived together with it (for example, when the feature template was archived). If one of those price points has subscribers without the feature, the call returns `422` naming that price point, unless you send `propagate_to_subscriptions: true`, which restores the copies and grants the feature to those subscribers.

**Price points with subscribers.** When the feature catalog item is scoped to a price point that already has subscribers and you don't send `propagate_to_subscriptions: true`, the new feature applies to new subscriptions only. The price point is cloned with its current features, the new feature is applied to the clone, and the response carries the clone's `price_point_id`. The feature catalog item in the response is also a different record, with a new `id`: later calls for the clone must use that `id`, not the original one. Existing subscribers stay on the original price point, unchanged; sign new subscriptions up on the clone to give them the change. A product price point has subscribers when any live subscription (including one in `awaiting_signup`) is on it. Price points that can't be cloned (archived or custom ones, or a component price point with an interval on a site without multi-frequency) make the request return `422`, with a message naming `propagate_to_subscriptions: true` as the way to apply the change to existing subscribers instead.

```go
CreateProductFeature(
    ctx context.Context,
    productId int,
    body *models.CreateFeatureCatalogItemRequest) (
    models.ApiResponse[models.FeatureCatalogItemResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Advanced Billing id of the product. |
| `body` | [`*models.CreateFeatureCatalogItemRequest`](../../doc/models/create-feature-catalog-item-request.md) | Body, Optional | The owning product or component is taken from the URL and must not be included in the request body. |

## Response Type

**201**: Created

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.FeatureCatalogItemResponse](../../doc/models/feature-catalog-item-response.md).

## Example Usage

```go
ctx := context.Background()

productId := 202

body := models.CreateFeatureCatalogItemRequest{
    Feature:              models.Feature2{
        FeatureTemplateId:        196,
        PropagateToSubscriptions: models.ToPointer(false),
    },
}

apiResponse, err := productFeaturesController.CreateProductFeature(ctx, productId, &body)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Read Product Feature

Returns a single feature catalog item attached to this product.

```go
ReadProductFeature(
    ctx context.Context,
    productId int,
    id int) (
    models.ApiResponse[models.FeatureCatalogItemResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Advanced Billing id of the product. |
| `id` | `int` | Template, Required | The Advanced Billing id of the feature catalog item. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.FeatureCatalogItemResponse](../../doc/models/feature-catalog-item-response.md).

## Example Usage

```go
ctx := context.Background()

productId := 202

id := 112

apiResponse, err := productFeaturesController.ReadProductFeature(ctx, productId, id)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |
| 404 | Not Found | `ApiError` |


# Update Product Feature

Updates the value or periodicity of a feature catalog item attached to this product. A request that changes nothing returns the item as it is.

**Price points with subscribers.** When the feature catalog item is scoped to a price point that already has subscribers and you don't send `propagate_to_subscriptions: true`, the update applies to new subscriptions only. The price point is cloned with its current features, the update is applied to the clone, and the response carries the clone's `price_point_id`. The feature catalog item in the response is also a different record, with a new `id`: later calls for the clone must use that `id`, not the original one. Existing subscribers stay on the original price point, unchanged; sign new subscriptions up on the clone to give them the change. A product price point has subscribers when any live subscription (including one in `awaiting_signup`) is on it. Price points that can't be cloned (archived or custom ones, or a component price point with an interval on a site without multi-frequency) make the request return `422`, with a message naming `propagate_to_subscriptions: true` as the way to apply the change to existing subscribers instead.

```go
UpdateProductFeature(
    ctx context.Context,
    productId int,
    id int,
    body *models.UpdateFeatureCatalogItemRequest) (
    models.ApiResponse[models.FeatureCatalogItemResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Advanced Billing id of the product. |
| `id` | `int` | Template, Required | The Advanced Billing id of the feature catalog item. |
| `body` | [`*models.UpdateFeatureCatalogItemRequest`](../../doc/models/update-feature-catalog-item-request.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.FeatureCatalogItemResponse](../../doc/models/feature-catalog-item-response.md).

## Example Usage

```go
ctx := context.Background()

productId := 202

id := 112

body := models.UpdateFeatureCatalogItemRequest{
    Feature:              models.Feature3{
        PropagateToSubscriptions: models.ToPointer(false),
    },
}

apiResponse, err := productFeaturesController.UpdateProductFeature(ctx, productId, id, &body)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Remove Product Feature

Removes a feature catalog item from this product. Returns `204` when the item is removed in place.

**Price points with subscribers.** When the feature catalog item is scoped to a price point that already has subscribers and you don't send `destroy_entitlements=true`, the removal applies to new subscriptions only. The price point is cloned with its current features, the removal is applied to the clone, and the call returns `200` with the archived feature catalog item on the clone (carrying the clone's `price_point_id` and a new `id`) instead of `204`. Existing subscribers stay on the original price point, unchanged; sign new subscriptions up on the clone to give them the change. A product price point has subscribers when any live subscription (including one in `awaiting_signup`) is on it. Price points that can't be cloned (archived or custom ones, or a component price point with an interval on a site without multi-frequency) make the request return `422`, with a message naming `destroy_entitlements=true` as the way to apply the change to existing subscribers instead.

```go
RemoveProductFeature(
    ctx context.Context,
    productId int,
    id int,
    destroyEntitlements *bool) (
    models.ApiResponse[*models.FeatureCatalogItemResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Advanced Billing id of the product. |
| `id` | `int` | Template, Required | The Advanced Billing id of the feature catalog item. |
| `destroyEntitlements` | `*bool` | Query, Optional | When `true`, permanently deletes this feature catalog item and every entitlement it created, revoking subscriber access immediately. When `false` (default), the feature catalog item is archived and existing entitlements are preserved.<br><br>**Default**: `false` |

## Response Type

**200**: The price point had subscribers, so the removal was made on a clone of it. Returns the archived feature catalog item on the clone.

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [*models.FeatureCatalogItemResponse](../../doc/models/feature-catalog-item-response.md).

## Example Usage

```go
ctx := context.Background()

productId := 202

id := 112

destroyEntitlements := false

apiResponse, err := productFeaturesController.RemoveProductFeature(ctx, productId, id, &destroyEntitlements)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |
| 404 | Not Found | `ApiError` |
| 422 | The price point has subscribers but can't be cloned (it is archived or custom, or it is a component price point with an interval on a site without multi-frequency). Send `destroy_entitlements=true` to remove the feature from its existing subscribers instead. | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Restore Product Feature

Clears the archived state of a feature catalog item attached to this product. Returns `422` if the parent feature template is still archived. Restore the feature template first. Otherwise, restoring an item that isn't archived returns it unchanged.

**Price points with subscribers.** When the feature catalog item is scoped to a price point that already has subscribers and you don't send `propagate_to_subscriptions=true`, the restore applies to new subscriptions only. The price point is cloned with its current features, the restore is applied to the clone, and the response carries the clone's `price_point_id`. The feature catalog item in the response is also a different record, with a new `id`: later calls for the clone must use that `id`, not the original one. Existing subscribers stay on the original price point, unchanged; sign new subscriptions up on the clone to give them the change. A product price point has subscribers when any live subscription (including one in `awaiting_signup`) is on it. Price points that can't be cloned (archived or custom ones, or a component price point with an interval on a site without multi-frequency) make the request return `422`, with a message naming `propagate_to_subscriptions: true` as the way to apply the change to existing subscribers instead.

```go
RestoreProductFeature(
    ctx context.Context,
    productId int,
    id int,
    propagateToSubscriptions *bool) (
    models.ApiResponse[models.FeatureCatalogItemResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Advanced Billing id of the product. |
| `id` | `int` | Template, Required | The Advanced Billing id of the feature catalog item. |
| `propagateToSubscriptions` | `*bool` | Query, Optional | When `true`, the restored feature is applied to the existing subscriptions on its price point. When `false` (default), an item scoped to a price point with subscribers is restored on a clone of that price point instead; a product- or component-level item is never cloned, and returns `422` if a price-point copy it restores has subscribers without the feature. See the description.<br><br>**Default**: `false` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.FeatureCatalogItemResponse](../../doc/models/feature-catalog-item-response.md).

## Example Usage

```go
ctx := context.Background()

productId := 202

id := 112

propagateToSubscriptions := false

apiResponse, err := productFeaturesController.RestoreProductFeature(ctx, productId, id, &propagateToSubscriptions)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

