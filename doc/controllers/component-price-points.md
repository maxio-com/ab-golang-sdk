# Component Price Points

```go
componentPricePointsController := client.ComponentPricePointsController()
```

## Class Name

`ComponentPricePointsController`

## Methods

* [Promote Component Price Point to Default](../../doc/controllers/component-price-points.md#promote-component-price-point-to-default)
* [Create Component Price Point](../../doc/controllers/component-price-points.md#create-component-price-point)
* [List Component Price Points](../../doc/controllers/component-price-points.md#list-component-price-points)
* [Bulk Create Component Price Points](../../doc/controllers/component-price-points.md#bulk-create-component-price-points)
* [Update Component Price Point](../../doc/controllers/component-price-points.md#update-component-price-point)
* [Read Component Price Point](../../doc/controllers/component-price-points.md#read-component-price-point)
* [Archive Component Price Point](../../doc/controllers/component-price-points.md#archive-component-price-point)
* [Unarchive Component Price Point](../../doc/controllers/component-price-points.md#unarchive-component-price-point)
* [Create Currency Prices](../../doc/controllers/component-price-points.md#create-currency-prices)
* [Update Currency Prices](../../doc/controllers/component-price-points.md#update-currency-prices)
* [List All Component Price Points](../../doc/controllers/component-price-points.md#list-all-component-price-points)


# Promote Component Price Point to Default

Sets a new default price point for the component. This new default will apply to all new subscriptions going forward - existing subscriptions will remain on their current price point.

See [Price Points Documentation](https://maxio.zendesk.com/hc/en-us/articles/24261191737101-Price-Points-Components) for more information on price points and moving subscriptions between price points.

Note: Custom price points are not able to be set as the default for a component.

```go
PromoteComponentPricePointToDefault(
    ctx context.Context,
    componentId int,
    pricePointId int) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `int` | Template, Required | The Advanced Billing id of the component to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Advanced Billing id of the price point |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentResponse](../../doc/models/component-response.md).

## Example Usage

```go
ctx := context.Background()

componentId := 222

pricePointId := 10

apiResponse, err := componentPricePointsController.PromoteComponentPricePointToDefault(ctx, componentId, pricePointId)
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
  "component": {
    "id": 292609,
    "name": "Text messages",
    "pricing_scheme": "stairstep",
    "unit_name": "text message",
    "unit_price": null,
    "product_family_id": 528484,
    "price_per_unit_in_cents": null,
    "kind": "metered_component",
    "archived": false,
    "taxable": false,
    "description": null,
    "created_at": "2019-08-02T05:54:53-04:00",
    "prices": [
      {
        "id": 47,
        "component_id": 292609,
        "starting_quantity": 1,
        "ending_quantity": null,
        "unit_price": "1.0",
        "price_point_id": 173,
        "formatted_unit_price": "$1.00"
      }
    ],
    "default_price_point_name": "Original"
  }
}
```


# Create Component Price Point

This endpoint can be used to create a new price point for an existing component.

```go
CreateComponentPricePoint(
    ctx context.Context,
    componentId int,
    body *models.CreateComponentPricePointRequest) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `int` | Template, Required | The Advanced Billing id of the component |
| `body` | [`*models.CreateComponentPricePointRequest`](../../doc/models/create-component-price-point-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentPricePointResponse](../../doc/models/component-price-point-response.md).

## Example Usage

```go
ctx := context.Background()

componentId := 222

body := models.CreateComponentPricePointRequest{
    PricePoint:           models.CreateComponentPricePointRequestPricePointContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
        Name:                 "Wholesale",
        Handle:               models.ToPointer("wholesale-handle"),
        PricingScheme:        models.PricingScheme_STAIRSTEP,
        Prices:               []models.Price{
            models.Price{
                StartingQuantity:     models.PriceStartingQuantityContainer.FromString("1"),
                EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromString("100"))),
                UnitPrice:            models.PriceUnitPriceContainer.FromString("5.00"),
            },
            models.Price{
                StartingQuantity:     models.PriceStartingQuantityContainer.FromString("101"),
                UnitPrice:            models.PriceUnitPriceContainer.FromString("4.00"),
            },
        },
        UseSiteExchangeRate:  models.ToPointer(false),
    }),
}

apiResponse, err := componentPricePointsController.CreateComponentPricePoint(ctx, componentId, &body)
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
| 422 | Unprocessable Entity (WebDAV) | [`ErrorArrayMapResponseException`](../../doc/models/error-array-map-response-exception.md) |


# List Component Price Points

Use this endpoint to read current price points that are associated with a component.

You may specify the component by using either the numeric id or the `handle:gold` syntax.

When fetching a component's price points, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response.

If the price point is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.

```go
ListComponentPricePoints(
    ctx context.Context,
    input ListComponentPricePointsInput) (
    models.ApiResponse[models.ComponentPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `int` | Template, Required | The Advanced Billing id of the component |
| `currencyPrices` | `*bool` | Query, Optional | Include an array of currency price data |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `filterType` | [`[]models.PricePointType`](../../doc/models/price-point-type.md) | Query, Optional | Use in query: `filter[type]=catalog,default`. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentPricePointsResponse](../../doc/models/component-price-points-response.md).

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListComponentPricePointsInput{
    ComponentId:    222,
    Page:           models.ToPointer(2),
    PerPage:        models.ToPointer(50),
Liquid error: Value cannot be null. (Parameter 'key')}

apiResponse, err := componentPricePointsController.ListComponentPricePoints(ctx, collectedInput)
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
      "id": 80,
      "default": false,
      "name": "Wholesale Two",
      "pricing_scheme": "per_unit",
      "component_id": 74,
      "handle": "wholesale-two",
      "archived_at": null,
      "created_at": "2017-07-05T13:55:40-04:00",
      "updated_at": "2017-07-05T13:55:40-04:00",
      "prices": [
        {
          "id": 121,
          "component_id": 74,
          "starting_quantity": 1,
          "ending_quantity": null,
          "unit_price": "5.0"
        }
      ]
    },
    {
      "id": 81,
      "default": false,
      "name": "MSRP",
      "pricing_scheme": "per_unit",
      "component_id": 74,
      "handle": "msrp",
      "archived_at": null,
      "created_at": "2017-07-05T13:55:40-04:00",
      "updated_at": "2017-07-05T13:55:40-04:00",
      "prices": [
        {
          "id": 122,
          "component_id": 74,
          "starting_quantity": 1,
          "ending_quantity": null,
          "unit_price": "4.0"
        }
      ]
    }
  ]
}
```


# Bulk Create Component Price Points

Use this endpoint to create multiple component price points in one request.

```go
BulkCreateComponentPricePoints(
    ctx context.Context,
    componentId string,
    body *models.CreateComponentPricePointsRequest) (
    models.ApiResponse[models.ComponentPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `string` | Template, Required | The Advanced Billing id of the component for which you want to fetch price points. |
| `body` | [`*models.CreateComponentPricePointsRequest`](../../doc/models/create-component-price-points-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentPricePointsResponse](../../doc/models/component-price-points-response.md).

## Example Usage

```go
ctx := context.Background()

componentId := "component_id8"

body := models.CreateComponentPricePointsRequest{
    PricePoints:          []models.CreateComponentPricePointsRequestPricePoints{
        models.CreateComponentPricePointsRequestPricePointsContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
            Name:                 "Wholesale",
            Handle:               models.ToPointer("wholesale"),
            PricingScheme:        models.PricingScheme_PERUNIT,
            Prices:               []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(1),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(5)),
                },
            },
        }),
        models.CreateComponentPricePointsRequestPricePointsContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
            Name:                 "MSRP",
            Handle:               models.ToPointer("msrp"),
            PricingScheme:        models.PricingScheme_PERUNIT,
            Prices:               []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(1),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(4)),
                },
            },
        }),
        models.CreateComponentPricePointsRequestPricePointsContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
            Name:                 "Special Pricing",
            Handle:               models.ToPointer("special"),
            PricingScheme:        models.PricingScheme_PERUNIT,
            Prices:               []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(1),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(5)),
                },
            },
        }),
    },
}

apiResponse, err := componentPricePointsController.BulkCreateComponentPricePoints(ctx, componentId, &body)
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
      "id": 80,
      "default": false,
      "name": "Wholesale Two",
      "pricing_scheme": "per_unit",
      "component_id": 74,
      "handle": "wholesale-two",
      "archived_at": null,
      "created_at": "2017-07-05T13:55:40-04:00",
      "updated_at": "2017-07-05T13:55:40-04:00",
      "prices": [
        {
          "id": 121,
          "component_id": 74,
          "starting_quantity": 1,
          "ending_quantity": null,
          "unit_price": "5.0"
        }
      ]
    },
    {
      "id": 81,
      "default": false,
      "name": "MSRP",
      "pricing_scheme": "per_unit",
      "component_id": 74,
      "handle": "msrp",
      "archived_at": null,
      "created_at": "2017-07-05T13:55:40-04:00",
      "updated_at": "2017-07-05T13:55:40-04:00",
      "prices": [
        {
          "id": 122,
          "component_id": 74,
          "starting_quantity": 1,
          "ending_quantity": null,
          "unit_price": "4.0"
        }
      ]
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Update Component Price Point

When updating a price point, it's prices can be updated as well by creating new prices or editing / removing existing ones.

Passing in a price bracket without an `id` will attempt to create a new price.

Including an `id` will update the corresponding price, and including the `_destroy` flag set to true along with the `id` will remove that price.

Note: Custom price points cannot be updated directly. They must be edited through the Subscription.

```go
UpdateComponentPricePoint(
    ctx context.Context,
    componentId models.UpdateComponentPricePointComponentId,
    pricePointId models.UpdateComponentPricePointPricePointId,
    body *models.UpdateComponentPricePointRequest) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | [`models.UpdateComponentPricePointComponentId`](../../doc/models/containers/update-component-price-point-component-id.md) | Template, Required | This is a container for one-of cases. |
| `pricePointId` | [`models.UpdateComponentPricePointPricePointId`](../../doc/models/containers/update-component-price-point-price-point-id.md) | Template, Required | This is a container for one-of cases. |
| `body` | [`*models.UpdateComponentPricePointRequest`](../../doc/models/update-component-price-point-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentPricePointResponse](../../doc/models/component-price-point-response.md).

## Example Usage

```go
ctx := context.Background()

componentId := models.UpdateComponentPricePointComponentIdContainer.FromNumber(144)

pricePointId := models.UpdateComponentPricePointPricePointIdContainer.FromNumber(188)

body := models.UpdateComponentPricePointRequest{
    PricePoint:           models.ToPointer(models.UpdateComponentPricePoint{
        Name:                 models.ToPointer("Default"),
        Prices:               []models.UpdatePrice{
            models.UpdatePrice{
                Id:                   models.ToPointer(1),
                EndingQuantity:       models.ToPointer(models.UpdatePriceEndingQuantityContainer.FromNumber(100)),
                UnitPrice:            models.ToPointer(models.UpdatePriceUnitPriceContainer.FromPrecision(float64(5))),
            },
            models.UpdatePrice{
                Id:                   models.ToPointer(2),
                Destroy:              models.ToPointer(true),
            },
            models.UpdatePrice{
                UnitPrice:            models.ToPointer(models.UpdatePriceUnitPriceContainer.FromPrecision(float64(4))),
                StartingQuantity:     models.ToPointer(models.UpdatePriceStartingQuantityContainer.FromNumber(101)),
            },
        },
    }),
}

apiResponse, err := componentPricePointsController.UpdateComponentPricePoint(ctx, componentId, pricePointId, &body)
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
| 422 | Unprocessable Entity (WebDAV) | [`ErrorArrayMapResponseException`](../../doc/models/error-array-map-response-exception.md) |


# Read Component Price Point

Use this endpoint to retrieve details for a specific component price point. You can achieve this by using either the component price point ID or handle.

```go
ReadComponentPricePoint(
    ctx context.Context,
    componentId models.ReadComponentPricePointComponentId,
    pricePointId models.ReadComponentPricePointPricePointId,
    currencyPrices *bool) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | [`models.ReadComponentPricePointComponentId`](../../doc/models/containers/read-component-price-point-component-id.md) | Template, Required | This is a container for one-of cases. |
| `pricePointId` | [`models.ReadComponentPricePointPricePointId`](../../doc/models/containers/read-component-price-point-price-point-id.md) | Template, Required | This is a container for one-of cases. |
| `currencyPrices` | `*bool` | Query, Optional | Include an array of currency price data |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentPricePointResponse](../../doc/models/component-price-point-response.md).

## Example Usage

```go
ctx := context.Background()

componentId := models.ReadComponentPricePointComponentIdContainer.FromNumber(144)

pricePointId := models.ReadComponentPricePointPricePointIdContainer.FromNumber(188)



apiResponse, err := componentPricePointsController.ReadComponentPricePoint(ctx, componentId, pricePointId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Archive Component Price Point

A price point can be archived at any time. Subscriptions using a price point that has been archived will continue using it until they're moved to another price point.

```go
ArchiveComponentPricePoint(
    ctx context.Context,
    componentId models.ArchiveComponentPricePointComponentId,
    pricePointId models.ArchiveComponentPricePointPricePointId) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | [`models.ArchiveComponentPricePointComponentId`](../../doc/models/containers/archive-component-price-point-component-id.md) | Template, Required | This is a container for one-of cases. |
| `pricePointId` | [`models.ArchiveComponentPricePointPricePointId`](../../doc/models/containers/archive-component-price-point-price-point-id.md) | Template, Required | This is a container for one-of cases. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentPricePointResponse](../../doc/models/component-price-point-response.md).

## Example Usage

```go
ctx := context.Background()

componentId := models.ArchiveComponentPricePointComponentIdContainer.FromNumber(144)

pricePointId := models.ArchiveComponentPricePointPricePointIdContainer.FromNumber(188)

apiResponse, err := componentPricePointsController.ArchiveComponentPricePoint(ctx, componentId, pricePointId)
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
    "id": 79,
    "default": false,
    "name": "Wholesale",
    "pricing_scheme": "stairstep",
    "component_id": 74,
    "handle": "wholesale-handle",
    "archived_at": "2017-07-06T15:04:00-04:00",
    "created_at": "2017-07-05T13:44:30-04:00",
    "updated_at": "2017-07-05T13:44:30-04:00",
    "prices": [
      {
        "id": 119,
        "component_id": 74,
        "starting_quantity": 1,
        "ending_quantity": 100,
        "unit_price": "5.0"
      },
      {
        "id": 120,
        "component_id": 74,
        "starting_quantity": 101,
        "ending_quantity": null,
        "unit_price": "4.0"
      }
    ]
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Unarchive Component Price Point

Use this endpoint to unarchive a component price point.

```go
UnarchiveComponentPricePoint(
    ctx context.Context,
    componentId int,
    pricePointId int) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `int` | Template, Required | The Advanced Billing id of the component to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Advanced Billing id of the price point |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentPricePointResponse](../../doc/models/component-price-point-response.md).

## Example Usage

```go
ctx := context.Background()

componentId := 222

pricePointId := 10

apiResponse, err := componentPricePointsController.UnarchiveComponentPricePoint(ctx, componentId, pricePointId)
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
    "id": 79,
    "default": false,
    "name": "Wholesale",
    "pricing_scheme": "stairstep",
    "component_id": 74,
    "handle": "wholesale-handle",
    "archived_at": null,
    "created_at": "2017-07-05T13:44:30-04:00",
    "updated_at": "2017-07-05T13:44:30-04:00",
    "prices": [
      {
        "id": 119,
        "component_id": 74,
        "starting_quantity": 1,
        "ending_quantity": 100,
        "unit_price": "5.0"
      },
      {
        "id": 120,
        "component_id": 74,
        "starting_quantity": 101,
        "ending_quantity": null,
        "unit_price": "4.0"
      }
    ]
  }
}
```


# Create Currency Prices

This endpoint allows you to create currency prices for a given currency that has been defined on the site level in your settings.

When creating currency prices, they need to mirror the structure of your primary pricing. For each price level defined on the component price point, there should be a matching price level created in the given currency.

Note: Currency Prices are not able to be created for custom price points.

```go
CreateCurrencyPrices(
    ctx context.Context,
    pricePointId int,
    body *models.CreateCurrencyPricesRequest) (
    models.ApiResponse[models.ComponentCurrencyPricesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `pricePointId` | `int` | Template, Required | The Advanced Billing id of the price point |
| `body` | [`*models.CreateCurrencyPricesRequest`](../../doc/models/create-currency-prices-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentCurrencyPricesResponse](../../doc/models/component-currency-prices-response.md).

## Example Usage

```go
ctx := context.Background()

pricePointId := 10

body := models.CreateCurrencyPricesRequest{
    CurrencyPrices:       []models.CreateCurrencyPrice{
        models.CreateCurrencyPrice{
            Currency:             models.ToPointer("EUR"),
            Price:                models.ToPointer(float64(50)),
            PriceId:              models.ToPointer(20),
        },
        models.CreateCurrencyPrice{
            Currency:             models.ToPointer("EUR"),
            Price:                models.ToPointer(float64(40)),
            PriceId:              models.ToPointer(21),
        },
    },
}

apiResponse, err := componentPricePointsController.CreateCurrencyPrices(ctx, pricePointId, &body)
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
      "id": 100,
      "currency": "EUR",
      "price": "123",
      "formatted_price": "€123,00",
      "price_id": 32669,
      "price_point_id": 25554
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorArrayMapResponseException`](../../doc/models/error-array-map-response-exception.md) |


# Update Currency Prices

This endpoint allows you to update currency prices for a given currency that has been defined on the site level in your settings.

Note: Currency Prices are not able to be updated for custom price points.

```go
UpdateCurrencyPrices(
    ctx context.Context,
    pricePointId int,
    body *models.UpdateCurrencyPricesRequest) (
    models.ApiResponse[models.ComponentCurrencyPricesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `pricePointId` | `int` | Template, Required | The Advanced Billing id of the price point |
| `body` | [`*models.UpdateCurrencyPricesRequest`](../../doc/models/update-currency-prices-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ComponentCurrencyPricesResponse](../../doc/models/component-currency-prices-response.md).

## Example Usage

```go
ctx := context.Background()

pricePointId := 10

body := models.UpdateCurrencyPricesRequest{
    CurrencyPrices:       []models.UpdateCurrencyPrice{
        models.UpdateCurrencyPrice{
            Id:                   100,
            Price:                float64(51),
        },
        models.UpdateCurrencyPrice{
            Id:                   101,
            Price:                float64(41),
        },
    },
}

apiResponse, err := componentPricePointsController.UpdateCurrencyPrices(ctx, pricePointId, &body)
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
      "id": 100,
      "currency": "EUR",
      "price": "123",
      "formatted_price": "€123,00",
      "price_id": 32669,
      "price_point_id": 25554
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorArrayMapResponseException`](../../doc/models/error-array-map-response-exception.md) |


# List All Component Price Points

This method allows to retrieve a list of Components Price Points belonging to a Site.

```go
ListAllComponentPricePoints(
    ctx context.Context,
    input ListAllComponentPricePointsInput) (
    models.ApiResponse[models.ListComponentsPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `include` | [`*models.ListComponentsPricePointsInclude`](../../doc/models/list-components-price-points-include.md) | Query, Optional | Allows including additional data in the response. Use in query: `include=currency_prices`. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Query, Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |
| `filter` | [`*models.ListPricePointsFilter`](../../doc/models/list-price-points-filter.md) | Query, Optional | Filter to use for List PricePoints operations |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ListComponentsPricePointsResponse](../../doc/models/list-components-price-points-response.md).

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListAllComponentPricePointsInput{
    Include:   models.ToPointer(models.ListComponentsPricePointsInclude_CURRENCYPRICES),
    Page:      models.ToPointer(2),
    PerPage:   models.ToPointer(50),
    Filter:    models.ToPointer(models.ListPricePointsFilter{
        StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-17", func(err error) { log.Fatalln(err) })),
        EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-15", func(err error) { log.Fatalln(err) })),
        StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "12/19/2011 09:15:30", func(err error) { log.Fatalln(err) })),
        EndDatetime:          models.ToPointer(parseTime(time.RFC3339, "06/07/2019 17:20:06", func(err error) { log.Fatalln(err) })),
        Type:                 []models.PricePointType{
            models.PricePointType_CATALOG,
            models.PricePointType_ENUMDEFAULT,
            models.PricePointType_CUSTOM,
        },
        Ids:                  []int{
            1,
            2,
            3,
        },
    }),
}

apiResponse, err := componentPricePointsController.ListAllComponentPricePoints(ctx, collectedInput)
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
      "id": 1,
      "name": "Auto-created",
      "type": "default",
      "pricing_scheme": "per_unit",
      "component_id": 2,
      "handle": "auto-created",
      "archived_at": null,
      "created_at": "2021-02-21T11:05:57-05:00",
      "updated_at": "2021-02-21T11:05:57-05:00",
      "prices": [
        {
          "id": 3,
          "component_id": 2,
          "starting_quantity": 0,
          "ending_quantity": null,
          "unit_price": "1.0",
          "price_point_id": 1,
          "formatted_unit_price": "$1.00",
          "segment_id": null
        }
      ],
      "tax_included": false
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

