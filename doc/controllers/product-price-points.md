# Product Price Points

```go
productPricePointsController := client.ProductPricePointsController()
```

## Class Name

`ProductPricePointsController`

## Methods

* [Create Product Price Point](../../doc/controllers/product-price-points.md#create-product-price-point)
* [List Product Price Points](../../doc/controllers/product-price-points.md#list-product-price-points)
* [Update Product Price Point](../../doc/controllers/product-price-points.md#update-product-price-point)
* [Read Product Price Point](../../doc/controllers/product-price-points.md#read-product-price-point)
* [Archive Product Price Point](../../doc/controllers/product-price-points.md#archive-product-price-point)
* [Unarchive Product Price Point](../../doc/controllers/product-price-points.md#unarchive-product-price-point)
* [Promote Product Price Point to Default](../../doc/controllers/product-price-points.md#promote-product-price-point-to-default)
* [Create Product Price Points](../../doc/controllers/product-price-points.md#create-product-price-points)
* [Create Product Currency Prices](../../doc/controllers/product-price-points.md#create-product-currency-prices)
* [Update Product Currency Prices](../../doc/controllers/product-price-points.md#update-product-currency-prices)
* [List All Product Price Points](../../doc/controllers/product-price-points.md#list-all-product-price-points)


# Create Product Price Point

[Product Price Point Documentation](https://chargify.zendesk.com/hc/en-us/articles/4407755824155)

```go
CreateProductPricePoint(
    ctx context.Context,
    productId interface{},
    body *models.CreateProductPricePointRequest) (
    models.ApiResponse[models.ProductPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `interface{}` | Template, Required | The id or handle of the product. When using the handle, it must be prefixed with `handle:` |
| `body` | [`*models.CreateProductPricePointRequest`](../../doc/models/create-product-price-point-request.md) | Body, Optional | - |

## Response Type

[`models.ProductPricePointResponse`](../../doc/models/product-price-point-response.md)

## Example Usage

```go
ctx := context.Background()
productId := interface{}("[key1, val1][key2, val2]")

bodyPricePoint := models.CreateProductPricePoint{
    Name:                    "Educational",
    Handle:                  models.ToPointer("educational"),
    PriceInCents:            int64(1000),
    Interval:                1,
    IntervalUnit:            models.IntervalUnit("month"),
    TrialPriceInCents:       models.ToPointer(int64(4900)),
    TrialInterval:           models.ToPointer(1),
    TrialIntervalUnit:       models.ToPointer(models.IntervalUnit("month")),
    TrialType:               models.ToPointer("payment_expected"),
    InitialChargeInCents:    models.ToPointer(int64(120000)),
    ExpirationInterval:      models.ToPointer(12),
    ExpirationIntervalUnit:  models.ToPointer(models.IntervalUnit("month")),
}

body := models.CreateProductPricePointRequest{
    PricePoint: bodyPricePoint,
}

apiResponse, err := productPricePointsController.CreateProductPricePoint(ctx, productId, &body)
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
  "price_point": {
    "id": 283,
    "name": "Educational",
    "handle": "educational",
    "price_in_cents": 1000,
    "interval": 1,
    "interval_unit": "month",
    "trial_price_in_cents": 4900,
    "trial_interval": 1,
    "trial_interval_unit": "month",
    "trial_type": "payment_expected",
    "initial_charge_in_cents": 120000,
    "initial_charge_after_trial": false,
    "expiration_interval": 12,
    "expiration_interval_unit": "month",
    "product_id": 901,
    "archived_at": "2023-11-30T06:37:20-05:00",
    "created_at": "2023-11-27T06:37:20-05:00",
    "updated_at": "2023-11-27T06:37:20-05:00"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ProductPricePointErrorResponseException`](../../doc/models/product-price-point-error-response-exception.md) |


# List Product Price Points

Use this endpoint to retrieve a list of product price points.

```go
ListProductPricePoints(
    ctx context.Context,input ListProductPricePointsInput) (
    models.ApiResponse[models.ListProductPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `interface{}` | Template, Required | The id or handle of the product. When using the handle, it must be prefixed with `handle:` |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 10. The maximum allowed values is 200; any per_page value over 200 will be changed to 200. |
| `currencyPrices` | `*bool` | Query, Optional | When fetching a product's price points, if you have defined multiple currencies at the site level, you can optionally pass the ?currency_prices=true query param to include an array of currency price data in the response. If the product price point is set to use_site_exchange_rate: true, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency. |
| `filterType` | [`[]models.PricePointType`](../../doc/models/price-point-type.md) | Query, Optional | Use in query: `filter[type]=catalog,default`. |

## Response Type

[`models.ListProductPricePointsResponse`](../../doc/models/list-product-price-points-response.md)

## Example Usage

```go
ctx := context.Background()
productId := interface{}("[key1, val1][key2, val2]")
page := 2
perPage := 10Liquid error: Value cannot be null. (Parameter 'key')

apiResponse, err := productPricePointsController.ListProductPricePoints(ctx, productId, &page, &perPage, nil, Liquid error: Value cannot be null. (Parameter 'key'))
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
  "price_points": [
    {
      "id": 283,
      "name": "Educational",
      "handle": "educational",
      "price_in_cents": 1000,
      "interval": 1,
      "interval_unit": "month",
      "trial_price_in_cents": 4900,
      "trial_interval": 1,
      "trial_interval_unit": "month",
      "trial_type": "payment_expected",
      "initial_charge_in_cents": 120000,
      "initial_charge_after_trial": false,
      "expiration_interval": 12,
      "expiration_interval_unit": "month",
      "product_id": 901,
      "archived_at": "2023-11-30T06:37:20-05:00",
      "created_at": "2023-11-27T06:37:20-05:00",
      "updated_at": "2023-11-27T06:37:20-05:00"
    }
  ]
}
```


# Update Product Price Point

Use this endpoint to update a product price point.

Note: Custom product price points are not able to be updated.

```go
UpdateProductPricePoint(
    ctx context.Context,
    productId interface{},
    pricePointId interface{},
    body *models.UpdateProductPricePointRequest) (
    models.ApiResponse[models.ProductPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `interface{}` | Template, Required | The id or handle of the product. When using the handle, it must be prefixed with `handle:` |
| `pricePointId` | `interface{}` | Template, Required | The id or handle of the price point. When using the handle, it must be prefixed with `handle:` |
| `body` | [`*models.UpdateProductPricePointRequest`](../../doc/models/update-product-price-point-request.md) | Body, Optional | - |

## Response Type

[`models.ProductPricePointResponse`](../../doc/models/product-price-point-response.md)

## Example Usage

```go
ctx := context.Background()
productId := interface{}("[key1, val1][key2, val2]")
pricePointId := interface{}("[key1, val1][key2, val2]")

bodyPricePoint := models.UpdateProductPricePoint{
    Handle:       models.ToPointer("educational"),
    PriceInCents: models.ToPointer(int64(1250)),
}

body := models.UpdateProductPricePointRequest{
    PricePoint: bodyPricePoint,
}

apiResponse, err := productPricePointsController.UpdateProductPricePoint(ctx, productId, pricePointId, &body)
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
  "price_point": {
    "id": 283,
    "name": "Educational",
    "handle": "educational",
    "price_in_cents": 1000,
    "interval": 1,
    "interval_unit": "month",
    "trial_price_in_cents": 4900,
    "trial_interval": 1,
    "trial_interval_unit": "month",
    "trial_type": "payment_expected",
    "initial_charge_in_cents": 120000,
    "initial_charge_after_trial": false,
    "expiration_interval": 12,
    "expiration_interval_unit": "month",
    "product_id": 901,
    "archived_at": "2023-11-30T06:37:20-05:00",
    "created_at": "2023-11-27T06:37:20-05:00",
    "updated_at": "2023-11-27T06:37:20-05:00"
  }
}
```


# Read Product Price Point

Use this endpoint to retrieve details for a specific product price point.

```go
ReadProductPricePoint(
    ctx context.Context,
    productId interface{},
    pricePointId interface{},
    currencyPrices *bool) (
    models.ApiResponse[models.ProductPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `interface{}` | Template, Required | The id or handle of the product. When using the handle, it must be prefixed with `handle:` |
| `pricePointId` | `interface{}` | Template, Required | The id or handle of the price point. When using the handle, it must be prefixed with `handle:` |
| `currencyPrices` | `*bool` | Query, Optional | When fetching a product's price points, if you have defined multiple currencies at the site level, you can optionally pass the ?currency_prices=true query param to include an array of currency price data in the response. If the product price point is set to use_site_exchange_rate: true, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency. |

## Response Type

[`models.ProductPricePointResponse`](../../doc/models/product-price-point-response.md)

## Example Usage

```go
ctx := context.Background()
productId := interface{}("[key1, val1][key2, val2]")
pricePointId := interface{}("[key1, val1][key2, val2]")

apiResponse, err := productPricePointsController.ReadProductPricePoint(ctx, productId, pricePointId, nil)
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
  "price_point": {
    "id": 283,
    "name": "Educational",
    "handle": "educational",
    "price_in_cents": 1000,
    "interval": 1,
    "interval_unit": "month",
    "trial_price_in_cents": 4900,
    "trial_interval": 1,
    "trial_interval_unit": "month",
    "trial_type": "payment_expected",
    "initial_charge_in_cents": 120000,
    "initial_charge_after_trial": false,
    "expiration_interval": 12,
    "expiration_interval_unit": "month",
    "product_id": 901,
    "archived_at": "2023-11-30T06:37:20-05:00",
    "created_at": "2023-11-27T06:37:20-05:00",
    "updated_at": "2023-11-27T06:37:20-05:00"
  }
}
```


# Archive Product Price Point

Use this endpoint to archive a product price point.

```go
ArchiveProductPricePoint(
    ctx context.Context,
    productId interface{},
    pricePointId interface{}) (
    models.ApiResponse[models.ProductPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `interface{}` | Template, Required | The id or handle of the product. When using the handle, it must be prefixed with `handle:` |
| `pricePointId` | `interface{}` | Template, Required | The id or handle of the price point. When using the handle, it must be prefixed with `handle:` |

## Response Type

[`models.ProductPricePointResponse`](../../doc/models/product-price-point-response.md)

## Example Usage

```go
ctx := context.Background()
productId := interface{}("[key1, val1][key2, val2]")
pricePointId := interface{}("[key1, val1][key2, val2]")

apiResponse, err := productPricePointsController.ArchiveProductPricePoint(ctx, productId, pricePointId)
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
  "price_point": {
    "id": 283,
    "name": "Educational",
    "handle": "educational",
    "price_in_cents": 1000,
    "interval": 1,
    "interval_unit": "month",
    "trial_price_in_cents": 4900,
    "trial_interval": 1,
    "trial_interval_unit": "month",
    "trial_type": "payment_expected",
    "initial_charge_in_cents": 120000,
    "initial_charge_after_trial": false,
    "expiration_interval": 12,
    "expiration_interval_unit": "month",
    "product_id": 901,
    "archived_at": "2023-11-30T06:37:20-05:00",
    "created_at": "2023-11-27T06:37:20-05:00",
    "updated_at": "2023-11-27T06:37:20-05:00"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Unarchive Product Price Point

Use this endpoint to unarchive an archived product price point.

```go
UnarchiveProductPricePoint(
    ctx context.Context,
    productId int,
    pricePointId int) (
    models.ApiResponse[models.ProductPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Chargify id of the product to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Chargify id of the product price point |

## Response Type

[`models.ProductPricePointResponse`](../../doc/models/product-price-point-response.md)

## Example Usage

```go
ctx := context.Background()
productId := 202
pricePointId := 10

apiResponse, err := productPricePointsController.UnarchiveProductPricePoint(ctx, productId, pricePointId)
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
  "price_point": {
    "id": 283,
    "name": "Educational",
    "handle": "educational",
    "price_in_cents": 1000,
    "interval": 1,
    "interval_unit": "month",
    "trial_price_in_cents": 4900,
    "trial_interval": 1,
    "trial_interval_unit": "month",
    "trial_type": "payment_expected",
    "initial_charge_in_cents": 120000,
    "initial_charge_after_trial": false,
    "expiration_interval": 12,
    "expiration_interval_unit": "month",
    "product_id": 901,
    "archived_at": "2023-11-30T06:37:20-05:00",
    "created_at": "2023-11-27T06:37:20-05:00",
    "updated_at": "2023-11-27T06:37:20-05:00"
  }
}
```


# Promote Product Price Point to Default

Use this endpoint to make a product price point the default for the product.

Note: Custom product price points are not able to be set as the default for a product.

```go
PromoteProductPricePointToDefault(
    ctx context.Context,
    productId int,
    pricePointId int) (
    models.ApiResponse[models.ProductResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Chargify id of the product to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Chargify id of the product price point |

## Response Type

[`models.ProductResponse`](../../doc/models/product-response.md)

## Example Usage

```go
ctx := context.Background()
productId := 202
pricePointId := 10

apiResponse, err := productPricePointsController.PromoteProductPricePointToDefault(ctx, productId, pricePointId)
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
  "product": {
    "id": 29778,
    "name": "Educational",
    "handle": "educational",
    "description": null,
    "accounting_code": null,
    "request_credit_card": true,
    "expiration_interval": 12,
    "expiration_interval_unit": "month",
    "created_at": "2023-12-01T06:56:12-05:00",
    "updated_at": "2023-12-01T06:56:26-05:00",
    "price_in_cents": 100,
    "interval": 2,
    "interval_unit": "month",
    "initial_charge_in_cents": 120000,
    "trial_price_in_cents": 4900,
    "trial_interval": 1,
    "trial_interval_unit": "month",
    "archived_at": null,
    "require_credit_card": true,
    "return_params": null,
    "taxable": false,
    "update_return_url": null,
    "tax_code": null,
    "initial_charge_after_trial": false,
    "version_number": 1,
    "update_return_params": null,
    "default_product_price_point_id": 32395,
    "request_billing_address": false,
    "require_billing_address": false,
    "require_shipping_address": false,
    "use_site_exchange_rate": true,
    "item_category": null,
    "product_price_point_id": 32395,
    "product_price_point_name": "Default",
    "product_price_point_handle": "uuid:8c878f50-726e-013c-c71b-0286551bb34f",
    "product_family": {
      "id": 933860,
      "name": "Acme Projects",
      "description": "Amazing project management tool",
      "handle": "acme-projects",
      "accounting_code": null,
      "created_at": "2023-12-01T06:56:12-05:00",
      "updated_at": "2023-12-01T06:56:12-05:00"
    },
    "public_signup_pages": []
  }
}
```


# Create Product Price Points

Use this endpoint to create multiple product price points in one request.

```go
CreateProductPricePoints(
    ctx context.Context,
    productId int,
    body *models.BulkCreateProductPricePointsRequest) (
    models.ApiResponse[models.BulkCreateProductPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Chargify id of the product to which the price points belong |
| `body` | [`*models.BulkCreateProductPricePointsRequest`](../../doc/models/bulk-create-product-price-points-request.md) | Body, Optional | - |

## Response Type

[`models.BulkCreateProductPricePointsResponse`](../../doc/models/bulk-create-product-price-points-response.md)

## Example Usage

```go
ctx := context.Background()
productId := 202

bodyPricePoints0 := models.CreateProductPricePoint{
    Name:                    "Educational",
    Handle:                  models.ToPointer("educational"),
    PriceInCents:            int64(1000),
    Interval:                1,
    IntervalUnit:            models.IntervalUnit("month"),
    TrialPriceInCents:       models.ToPointer(int64(4900)),
    TrialInterval:           models.ToPointer(1),
    TrialIntervalUnit:       models.ToPointer(models.IntervalUnit("month")),
    TrialType:               models.ToPointer("payment_expected"),
    InitialChargeInCents:    models.ToPointer(int64(120000)),
    ExpirationInterval:      models.ToPointer(12),
    ExpirationIntervalUnit:  models.ToPointer(models.IntervalUnit("month")),
}

bodyPricePoints1 := models.CreateProductPricePoint{
    Name:                    "More Educational",
    Handle:                  models.ToPointer("more-educational"),
    PriceInCents:            int64(2000),
    Interval:                1,
    IntervalUnit:            models.IntervalUnit("month"),
    TrialPriceInCents:       models.ToPointer(int64(4900)),
    TrialInterval:           models.ToPointer(1),
    TrialIntervalUnit:       models.ToPointer(models.IntervalUnit("month")),
    TrialType:               models.ToPointer("payment_expected"),
    InitialChargeInCents:    models.ToPointer(int64(120000)),
    ExpirationInterval:      models.ToPointer(12),
    ExpirationIntervalUnit:  models.ToPointer(models.IntervalUnit("month")),
}

bodyPricePoints := []models.CreateProductPricePoint{bodyPricePoints0, bodyPricePoints1}
body := models.BulkCreateProductPricePointsRequest{
    PricePoints: bodyPricePoints,
}

apiResponse, err := productPricePointsController.CreateProductPricePoints(ctx, productId, &body)
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
  "price_points": [
    {
      "id": 283,
      "name": "Educational",
      "handle": "educational",
      "price_in_cents": 1000,
      "interval": 1,
      "interval_unit": "month",
      "trial_price_in_cents": 4900,
      "trial_interval": 1,
      "trial_interval_unit": "month",
      "trial_type": "payment_expected",
      "initial_charge_in_cents": 120000,
      "initial_charge_after_trial": false,
      "expiration_interval": 12,
      "expiration_interval_unit": "month",
      "product_id": 901,
      "archived_at": "2023-11-30T06:37:20-05:00",
      "created_at": "2023-11-27T06:37:20-05:00",
      "updated_at": "2023-11-27T06:37:20-05:00"
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | `ApiError` |


# Create Product Currency Prices

This endpoint allows you to create currency prices for a given currency that has been defined on the site level in your settings.

When creating currency prices, they need to mirror the structure of your primary pricing. If the product price point defines a trial and/or setup fee, each currency must also define a trial and/or setup fee.

Note: Currency Prices are not able to be created for custom product price points.

```go
CreateProductCurrencyPrices(
    ctx context.Context,
    productPricePointId int,
    body *models.CreateProductCurrencyPricesRequest) (
    models.ApiResponse[models.ProductPricePointCurrencyPrice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productPricePointId` | `int` | Template, Required | The Chargify id of the product price point |
| `body` | [`*models.CreateProductCurrencyPricesRequest`](../../doc/models/create-product-currency-prices-request.md) | Body, Optional | - |

## Response Type

[`models.ProductPricePointCurrencyPrice`](../../doc/models/product-price-point-currency-price.md)

## Example Usage

```go
ctx := context.Background()
productPricePointId := 234

bodyCurrencyPrices0 := models.CreateProductCurrencyPrice{
    Currency: "EUR",
    Price:    60,
    Role:     models.CurrencyPriceRole("baseline"),
}

bodyCurrencyPrices1 := models.CreateProductCurrencyPrice{
    Currency: "EUR",
    Price:    30,
    Role:     models.CurrencyPriceRole("trial"),
}

bodyCurrencyPrices2 := models.CreateProductCurrencyPrice{
    Currency: "EUR",
    Price:    100,
    Role:     models.CurrencyPriceRole("initial"),
}

bodyCurrencyPrices := []models.CreateProductCurrencyPrice{bodyCurrencyPrices0, bodyCurrencyPrices1, bodyCurrencyPrices2}
body := models.CreateProductCurrencyPricesRequest{
    CurrencyPrices: bodyCurrencyPrices,
}

apiResponse, err := productPricePointsController.CreateProductCurrencyPrices(ctx, productPricePointId, &body)
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
  "currency_prices": [
    {
      "id": 123,
      "currency": "EUR",
      "price": 100,
      "formatted_price": "€123,00",
      "product_price_point_id": 32669,
      "role": "baseline"
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorMapResponseException`](../../doc/models/error-map-response-exception.md) |


# Update Product Currency Prices

This endpoint allows you to update the `price`s of currency prices for a given currency that exists on the product price point.

When updating the pricing, it needs to mirror the structure of your primary pricing. If the product price point defines a trial and/or setup fee, each currency must also define a trial and/or setup fee.

Note: Currency Prices are not able to be updated for custom product price points.

```go
UpdateProductCurrencyPrices(
    ctx context.Context,
    productPricePointId int,
    body *models.UpdateCurrencyPricesRequest) (
    models.ApiResponse[models.ProductPricePointCurrencyPrice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productPricePointId` | `int` | Template, Required | The Chargify id of the product price point |
| `body` | [`*models.UpdateCurrencyPricesRequest`](../../doc/models/update-currency-prices-request.md) | Body, Optional | - |

## Response Type

[`models.ProductPricePointCurrencyPrice`](../../doc/models/product-price-point-currency-price.md)

## Example Usage

```go
ctx := context.Background()
productPricePointId := 234

bodyCurrencyPrices0 := models.UpdateCurrencyPrice{
    Id:    200,
    Price: 15,
}

bodyCurrencyPrices1 := models.UpdateCurrencyPrice{
    Id:    201,
    Price: 5,
}

bodyCurrencyPrices := []models.UpdateCurrencyPrice{bodyCurrencyPrices0, bodyCurrencyPrices1}
body := models.UpdateCurrencyPricesRequest{
    CurrencyPrices: bodyCurrencyPrices,
}

apiResponse, err := productPricePointsController.UpdateProductCurrencyPrices(ctx, productPricePointId, &body)
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
  "currency_prices": [
    {
      "id": 123,
      "currency": "EUR",
      "price": 100,
      "formatted_price": "€123,00",
      "product_price_point_id": 32669,
      "role": "baseline"
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorMapResponseException`](../../doc/models/error-map-response-exception.md) |


# List All Product Price Points

This method allows retrieval of a list of Products Price Points belonging to a Site.

```go
ListAllProductPricePoints(
    ctx context.Context,input ListAllProductPricePointsInput) (
    models.ApiResponse[models.ListProductPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Query, Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |
| `filterArchivedAt` | [`*models.IncludeNotNull`](../../doc/models/include-not-null.md) | Query, Optional | Allows fetching price points only if archived_at is present or not. Use in query: `filter[archived_at]=not_null`. |
| `filterDateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Query, Optional | The type of filter you would like to apply to your search. Use in query: `filter[date_field]=created_at`. |
| `filterEndDate` | `*time.Time` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns price points with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `filterEndDatetime` | `*time.Time` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns price points with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |
| `filterIds` | `[]int` | Query, Optional | Allows fetching price points with matching id based on provided values. Use in query: `filter[ids]=1,2,3`. |
| `filterStartDate` | `*time.Time` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns price points with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `filterStartDatetime` | `*time.Time` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns price points with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `filterType` | [`[]models.PricePointType`](../../doc/models/price-point-type.md) | Query, Optional | Allows fetching price points with matching type. Use in query: `filter[type]=catalog,custom`. |
| `include` | [`*models.ListProductsPricePointsInclude`](../../doc/models/list-products-price-points-include.md) | Query, Optional | Allows including additional data in the response. Use in query: `include=currency_prices`. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |

## Response Type

[`models.ListProductPricePointsResponse`](../../doc/models/list-product-price-points-response.md)

## Example Usage

```go
ctx := context.Background()Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')
include := models.ListProductsPricePointsInclude("currency_prices")
page := 2
perPage := 50

apiResponse, err := productPricePointsController.ListAllProductPricePoints(ctx, nil, Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), &include, &page, &perPage)
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
  "price_points": [
    {
      "id": 0,
      "name": "My pricepoint",
      "handle": "handle",
      "price_in_cents": 10,
      "interval": 5,
      "interval_unit": "month",
      "trial_price_in_cents": 10,
      "trial_interval": 1,
      "trial_interval_unit": "month",
      "trial_type": "payment_expected",
      "introductory_offer": true,
      "initial_charge_in_cents": 0,
      "initial_charge_after_trial": true,
      "expiration_interval": 0,
      "expiration_interval_unit": "month",
      "product_id": 1230,
      "created_at": "2021-04-02T17:52:09-04:00",
      "updated_at": "2021-04-02T17:52:09-04:00",
      "use_site_exchange_rate": true
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

