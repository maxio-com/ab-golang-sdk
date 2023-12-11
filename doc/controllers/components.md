# Components

```go
componentsController := client.ComponentsController()
```

## Class Name

`ComponentsController`

## Methods

* [Create Component](components.md#create-component)
* [Read Component by Handle](components.md#read-component-by-handle)
* [Read Component by Id](components.md#read-component-by-id)
* [Update Product Family Component](components.md#update-product-family-component)
* [Archive Component](components.md#archive-component)
* [List Components](components.md#list-components)
* [Update Component](components.md#update-component)
* [Update Default Price Point for Component](components.md#update-default-price-point-for-component)
* [List Components for Product Family](components.md#list-components-for-product-family)
* [Create Component Price Point](components.md#create-component-price-point)
* [List Component Price Points](components.md#list-component-price-points)
* [Create Component Price Points](components.md#create-component-price-points)
* [Update Component Price Point](components.md#update-component-price-point)
* [Archive Component Price Point](components.md#archive-component-price-point)
* [Unarchive Component Price Point](components.md#unarchive-component-price-point)
* [Create Currency Prices](components.md#create-currency-prices)
* [Update Currency Prices](components.md#update-currency-prices)
* [List All Component Price Points](components.md#list-all-component-price-points)


# Create Component

This request will create a component definition under the specified product family. These component definitions determine what components are named, how they are measured, and how much they cost.

Components can then be added and “allocated” for each subscription to a product in the product family. These component line-items affect how much a subscription will be charged, depending on the current allocations (i.e. 4 IP Addresses, or SSL “enabled”)

This documentation covers both component definitions and component line-items. Please understand the difference.

Please note that you may not edit components via API. To do so, please log into the application.

### Component Documentation

For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).

For information on how to record component usage against a subscription, please see the following resources:

+ [Proration and Component Allocations](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677#applying-proration-and-recording-components)
+ [Recording component usage against a subscription](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404606587917#recording-component-usage)

```go
CreateComponent(
    ctx context.Context,
    productFamilyId int,
    componentKind models.ComponentKindPathEnum,
    body *interface{}) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the component belongs |
| `componentKind` | [`models.ComponentKindPathEnum`](../models/component-kind-path-enum.md) | Template, Required | The component kind |
| `body` | `*interface{}` | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../models/component-response.md)

## Example Usage

```go
ctx := context.Background()
productFamilyId := 140
componentKind := models.ComponentKindPathEnum("on_off_components")
body := interface{}("[metered_component, DotLiquid.Hash]")

apiResponse, err := componentsController.CreateComponent(ctx, productFamilyId, componentKind, &body)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../models/error-list-response-exception.md) |


# Read Component by Handle

This request will return information regarding a component having the handle you provide. You can identify your components with a handle so you don't have to save or reference the IDs we generate.

```go
ReadComponentByHandle(
    ctx context.Context,
    handle string) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `handle` | `string` | Query, Required | The handle of the component to find |

## Response Type

[`models.ComponentResponse`](../models/component-response.md)

## Example Usage

```go
ctx := context.Background()
handle := "handle6"

apiResponse, err := componentsController.ReadComponentByHandle(ctx, handle)
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
    "id": 399853,
    "name": "Annual Support Services",
    "pricing_scheme": null,
    "unit_name": "on/off",
    "unit_price": "100.0",
    "product_family_id": 997233,
    "price_per_unit_in_cents": null,
    "kind": "on_off_component",
    "archived": false,
    "taxable": true,
    "description": "Prepay for support services",
    "default_price_point_id": 121003,
    "price_point_count": 4,
    "price_points_url": "https://general-goods.chargify.com/components/399853/price_points",
    "tax_code": "D0000000",
    "recurring": true,
    "upgrade_charge": null,
    "downgrade_credit": null,
    "created_at": "2019-08-02T05:54:53-04:00",
    "default_price_point_name": "Original",
    "product_family_name": "Chargify"
  }
}
```


# Read Component by Id

This request will return information regarding a component from a specific product family.

You may read the component by either the component's id or handle. When using the handle, it must be prefixed with `handle:`.

```go
ReadComponentById(
    ctx context.Context,
    productFamilyId int,
    componentId string) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the component belongs |
| `componentId` | `string` | Template, Required | Either the Chargify id of the component or the handle for the component prefixed with `handle:`<br>**Constraints**: *Pattern*: `/\A(?:\d+\|handle:(?:uuid:\|[a-z])(?:\w\|-)+)\z/` |

## Response Type

[`models.ComponentResponse`](../models/component-response.md)

## Example Usage

```go
ctx := context.Background()
productFamilyId := 140
componentId := "component_id8"

apiResponse, err := componentsController.ReadComponentById(ctx, productFamilyId, componentId)
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
    "id": 399853,
    "name": "Annual Support Services",
    "pricing_scheme": null,
    "unit_name": "on/off",
    "unit_price": "100.0",
    "product_family_id": 997233,
    "price_per_unit_in_cents": null,
    "kind": "on_off_component",
    "archived": false,
    "taxable": true,
    "description": "Prepay for support services",
    "default_price_point_id": 121003,
    "price_point_count": 4,
    "price_points_url": "https://general-goods.chargify.com/components/399853/price_points",
    "tax_code": "D0000000",
    "recurring": true,
    "upgrade_charge": null,
    "downgrade_credit": null,
    "created_at": "2019-08-02T05:54:53-04:00",
    "default_price_point_name": "Original",
    "product_family_name": "Chargify"
  }
}
```


# Update Product Family Component

This request will update a component from a specific product family.

You may read the component by either the component's id or handle. When using the handle, it must be prefixed with `handle:`.

```go
UpdateProductFamilyComponent(
    ctx context.Context,
    productFamilyId int,
    componentId string,
    body *models.UpdateComponentRequest) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the component belongs |
| `componentId` | `string` | Template, Required | Either the Chargify id of the component or the handle for the component prefixed with `handle:`<br>**Constraints**: *Pattern*: `/\A(?:\d+\|handle:(?:uuid:\|[a-z])(?:\w\|-)+)\z/` |
| `body` | [`*models.UpdateComponentRequest`](../models/update-component-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../models/component-response.md)

## Example Usage

```go
ctx := context.Background()
productFamilyId := 140
componentId := "component_id8"

bodyComponent := models.UpdateComponent{
    ItemCategory:        models.NewOptional(models.ToPointer(models.ItemCategoryEnum("Business Software"))),
}

body := models.UpdateComponentRequest{
    Component: bodyComponent,
}

apiResponse, err := componentsController.UpdateProductFamilyComponent(ctx, productFamilyId, componentId, &body)
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
    "id": 399853,
    "name": "Annual Support Services",
    "pricing_scheme": null,
    "unit_name": "on/off",
    "unit_price": "100.0",
    "product_family_id": 997233,
    "price_per_unit_in_cents": null,
    "kind": "on_off_component",
    "archived": false,
    "taxable": true,
    "description": "Prepay for support services",
    "default_price_point_id": 121003,
    "price_point_count": 4,
    "price_points_url": "https://general-goods.chargify.com/components/399853/price_points",
    "tax_code": "D0000000",
    "recurring": true,
    "upgrade_charge": null,
    "downgrade_credit": null,
    "created_at": "2019-08-02T05:54:53-04:00",
    "default_price_point_name": "Original",
    "product_family_name": "Chargify"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../models/error-list-response-exception.md) |


# Archive Component

Sending a DELETE request to this endpoint will archive the component. All current subscribers will be unffected; their subscription/purchase will continue to be charged as usual.

```go
ArchiveComponent(
    ctx context.Context,
    productFamilyId int,
    componentId string) (
    models.ApiResponse[models.Component],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the component belongs |
| `componentId` | `string` | Template, Required | Either the Chargify id of the component or the handle for the component prefixed with `handle:`<br>**Constraints**: *Pattern*: `/\A(?:\d+\|handle:(?:uuid:\|[a-z])(?:\w\|-)+)\z/` |

## Response Type

[`models.Component`](../models/component.md)

## Example Usage

```go
ctx := context.Background()
productFamilyId := 140
componentId := "component_id8"

apiResponse, err := componentsController.ArchiveComponent(ctx, productFamilyId, componentId)
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
  "id": 25407138,
  "name": "cillum aute",
  "pricing_scheme": "stairstep",
  "unit_name": "nulla in",
  "unit_price": "Excepteur veniam",
  "product_family_id": -56705047,
  "kind": "prepaid_usage_component",
  "archived": true,
  "taxable": false,
  "description": "reprehenderit laborum qui fugiat",
  "default_price_point_id": -64328176,
  "price_point_count": 15252407,
  "price_points_url": "dolor mollit consequat",
  "tax_code": "ea nisi",
  "recurring": false,
  "created_at": "dolor qui deserunt tempor",
  "default_price_point_name": "cupidatat Lorem non aliqua",
  "product_family_name": "do elit",
  "hide_date_range_on_invoice": false
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../models/error-list-response-exception.md) |


# List Components

This request will return a list of components for a site.

```go
ListComponents(
    ctx context.Context,input ListComponentsInput) (
    models.ApiResponse[[]models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `dateField` | [`*models.BasicDateFieldEnum`](../models/basic-date-field-enum.md) | Query, Optional | The type of filter you would like to apply to your search. |
| `startDate` | `*string` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `endDate` | `*string` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `startDatetime` | `*string` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `endDatetime` | `*string` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.  optional |
| `includeArchived` | `*bool` | Query, Optional | Include archived items |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `20`<br>**Constraints**: `<= 200` |
| `filterIds` | `[]string` | Query, Optional | Allows fetching components with matching id based on provided value. Use in query `filter[ids]=1,2,3`. |
| `filterUseSiteExchangeRate` | `*bool` | Query, Optional | Allows fetching components with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`. |

## Response Type

[`[]models.ComponentResponse`](../models/component-response.md)

## Example Usage

```go
ctx := context.Background()
dateField := models.BasicDateFieldEnum("updated_at")
page := 2
perPage := 50Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')

apiResponse, err := componentsController.ListComponents(ctx, &dateField, nil, nil, nil, nil, nil, &page, &perPage, Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'))
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
    "component": {
      "id": 399850,
      "name": "$1.00 component",
      "pricing_scheme": "per_unit",
      "unit_name": "Component",
      "unit_price": "1.0",
      "product_family_id": 997233,
      "price_per_unit_in_cents": null,
      "kind": "quantity_based_component",
      "archived": false,
      "taxable": false,
      "description": "Component",
      "default_price_point_id": 121000,
      "prices": [
        {
          "id": 630687,
          "component_id": 399850,
          "starting_quantity": 1,
          "ending_quantity": null,
          "unit_price": "1.0",
          "price_point_id": 121000,
          "formatted_unit_price": "$1.00"
        }
      ],
      "price_point_count": 2,
      "price_points_url": "https://general-goods.chargify.com/components/399850/price_points",
      "tax_code": null,
      "recurring": true,
      "upgrade_charge": null,
      "downgrade_credit": null,
      "created_at": "2019-08-01T09:35:38-04:00",
      "default_price_point_name": "Original",
      "product_family_name": "Chargify",
      "use_site_exchange_rate": true
    }
  },
  {
    "component": {
      "id": 399853,
      "name": "Annual Support Services",
      "pricing_scheme": null,
      "unit_name": "on/off",
      "unit_price": "100.0",
      "product_family_id": 997233,
      "price_per_unit_in_cents": null,
      "kind": "on_off_component",
      "archived": false,
      "taxable": true,
      "description": "Prepay for support services",
      "default_price_point_id": 121003,
      "price_point_count": 4,
      "price_points_url": "https://general-goods.chargify.com/components/399853/price_points",
      "tax_code": "D0000000",
      "recurring": true,
      "upgrade_charge": null,
      "downgrade_credit": null,
      "created_at": "2019-08-01T09:35:37-04:00",
      "default_price_point_name": "Original",
      "product_family_name": "Chargify",
      "use_site_exchange_rate": true
    }
  },
  {
    "component": {
      "id": 386937,
      "name": "Cancellation fee",
      "pricing_scheme": null,
      "unit_name": "on/off",
      "unit_price": "35.0",
      "product_family_id": 997233,
      "price_per_unit_in_cents": null,
      "kind": "on_off_component",
      "archived": false,
      "taxable": false,
      "description": "",
      "default_price_point_id": 108307,
      "price_point_count": 1,
      "price_points_url": "https://general-goods.chargify.com/components/386937/price_points",
      "tax_code": null,
      "recurring": true,
      "upgrade_charge": null,
      "downgrade_credit": null,
      "created_at": "2019-08-01T09:35:38-04:00",
      "default_price_point_name": "Original",
      "product_family_name": "Chargify",
      "use_site_exchange_rate": true
    }
  }
]
```


# Update Component

This request will update a component.

You may read the component by either the component's id or handle. When using the handle, it must be prefixed with `handle:`.

```go
UpdateComponent(
    ctx context.Context,
    componentId string,
    body *models.UpdateComponentRequest) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `string` | Template, Required | The id or handle of the component |
| `body` | [`*models.UpdateComponentRequest`](../models/update-component-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../models/component-response.md)

## Example Usage

```go
ctx := context.Background()
componentId := "component_id8"

bodyComponent := models.UpdateComponent{
    ItemCategory:        models.NewOptional(models.ToPointer(models.ItemCategoryEnum("Business Software"))),
}

body := models.UpdateComponentRequest{
    Component: bodyComponent,
}

apiResponse, err := componentsController.UpdateComponent(ctx, componentId, &body)
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
    "id": 399853,
    "name": "Annual Support Services",
    "pricing_scheme": null,
    "unit_name": "on/off",
    "unit_price": "100.0",
    "product_family_id": 997233,
    "price_per_unit_in_cents": null,
    "kind": "on_off_component",
    "archived": false,
    "taxable": true,
    "description": "Prepay for support services",
    "default_price_point_id": 121003,
    "price_point_count": 4,
    "price_points_url": "https://general-goods.chargify.com/components/399853/price_points",
    "tax_code": "D0000000",
    "recurring": true,
    "upgrade_charge": null,
    "downgrade_credit": null,
    "created_at": "2019-08-02T05:54:53-04:00",
    "default_price_point_name": "Original",
    "product_family_name": "Chargify"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../models/error-list-response-exception.md) |


# Update Default Price Point for Component

Sets a new default price point for the component. This new default will apply to all new subscriptions going forward - existing subscriptions will remain on their current price point.

See [Price Points Documentation](https://chargify.zendesk.com/hc/en-us/articles/4407755865883#price-points) for more information on price points and moving subscriptions between price points.

Note: Custom price points are not able to be set as the default for a component.

```go
UpdateDefaultPricePointForComponent(
    ctx context.Context,
    componentId int,
    pricePointId int) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `int` | Template, Required | The Chargify id of the component to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |

## Response Type

[`models.ComponentResponse`](../models/component-response.md)

## Example Usage

```go
ctx := context.Background()
componentId := 222
pricePointId := 10

apiResponse, err := componentsController.UpdateDefaultPricePointForComponent(ctx, componentId, pricePointId)
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


# List Components for Product Family

This request will return a list of components for a particular product family.

```go
ListComponentsForProductFamily(
    ctx context.Context,input ListComponentsForProductFamilyInput) (
    models.ApiResponse[[]models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family |
| `includeArchived` | `*bool` | Query, Optional | Include archived items. |
| `filterIds` | `[]int` | Query, Optional | Allows fetching components with matching id based on provided value. Use in query `filter[ids]=1,2`. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `20`<br>**Constraints**: `<= 200` |
| `dateField` | [`*models.BasicDateFieldEnum`](../models/basic-date-field-enum.md) | Query, Optional | The type of filter you would like to apply to your search. Use in query `date_field=created_at`. |
| `endDate` | `*string` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `endDatetime` | `*string` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. optional. |
| `startDate` | `*string` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `startDatetime` | `*string` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `filterUseSiteExchangeRate` | `*bool` | Query, Optional | Allows fetching components with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`. |

## Response Type

[`[]models.ComponentResponse`](../models/component-response.md)

## Example Usage

```go
ctx := context.Background()
productFamilyId := 140Liquid error: Value cannot be null. (Parameter 'key')
page := 2
perPage := 50
dateField := models.BasicDateFieldEnum("updated_at")Liquid error: Value cannot be null. (Parameter 'key')

apiResponse, err := componentsController.ListComponentsForProductFamily(ctx, productFamilyId, nil, Liquid error: Value cannot be null. (Parameter 'key'), &page, &perPage, &dateField, nil, nil, nil, nil, Liquid error: Value cannot be null. (Parameter 'key'))
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
    "component": {
      "id": 399850,
      "name": "$1.00 component",
      "pricing_scheme": "per_unit",
      "unit_name": "Component",
      "unit_price": "1.0",
      "product_family_id": 997233,
      "price_per_unit_in_cents": null,
      "kind": "quantity_based_component",
      "archived": false,
      "taxable": false,
      "description": "Component",
      "default_price_point_id": 121000,
      "prices": [
        {
          "id": 630687,
          "component_id": 399850,
          "starting_quantity": 1,
          "ending_quantity": null,
          "unit_price": "1.0",
          "price_point_id": 121000,
          "formatted_unit_price": "$1.00"
        }
      ],
      "price_point_count": 2,
      "price_points_url": "https://general-goods.chargify.com/components/399850/price_points",
      "tax_code": null,
      "recurring": true,
      "upgrade_charge": null,
      "downgrade_credit": null,
      "created_at": "2019-08-01T09:35:38-04:00",
      "default_price_point_name": "Original",
      "product_family_name": "Chargify",
      "use_site_exchange_rate": true
    }
  },
  {
    "component": {
      "id": 399853,
      "name": "Annual Support Services",
      "pricing_scheme": null,
      "unit_name": "on/off",
      "unit_price": "100.0",
      "product_family_id": 997233,
      "price_per_unit_in_cents": null,
      "kind": "on_off_component",
      "archived": false,
      "taxable": true,
      "description": "Prepay for support services",
      "default_price_point_id": 121003,
      "price_point_count": 4,
      "price_points_url": "https://general-goods.chargify.com/components/399853/price_points",
      "tax_code": "D0000000",
      "recurring": true,
      "upgrade_charge": null,
      "downgrade_credit": null,
      "created_at": "2019-08-01T09:35:37-04:00",
      "default_price_point_name": "Original",
      "product_family_name": "Chargify",
      "use_site_exchange_rate": true
    }
  },
  {
    "component": {
      "id": 386937,
      "name": "Cancellation fee",
      "pricing_scheme": null,
      "unit_name": "on/off",
      "unit_price": "35.0",
      "product_family_id": 997233,
      "price_per_unit_in_cents": null,
      "kind": "on_off_component",
      "archived": false,
      "taxable": false,
      "description": "",
      "default_price_point_id": 108307,
      "price_point_count": 1,
      "price_points_url": "https://general-goods.chargify.com/components/386937/price_points",
      "tax_code": null,
      "recurring": true,
      "upgrade_charge": null,
      "downgrade_credit": null,
      "created_at": "2019-08-01T09:35:38-04:00",
      "default_price_point_name": "Original",
      "product_family_name": "Chargify",
      "use_site_exchange_rate": true
    }
  }
]
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
| `componentId` | `int` | Template, Required | The Chargify id of the component |
| `body` | [`*models.CreateComponentPricePointRequest`](../models/create-component-price-point-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentPricePointResponse`](../models/component-price-point-response.md)

## Example Usage

```go
ctx := context.Background()
componentId := 222

bodyPricePointPrices0 := models.Price{
    StartingQuantity: interface{}("[key1, val1][key2, val2]"),
    EndingQuantity:   models.NewOptional(models.ToPointer(interface{}("[key1, val1][key2, val2]"))),
    UnitPrice:        interface{}("[key1, val1][key2, val2]"),
}

bodyPricePointPrices1 := models.Price{
    StartingQuantity: interface{}("[key1, val1][key2, val2]"),
    UnitPrice:        interface{}("[key1, val1][key2, val2]"),
}

bodyPricePointPrices := []models.Price{bodyPricePointPrices0, bodyPricePointPrices1}
bodyPricePoint := models.PricePoint{
    Name:                     models.ToPointer("Wholesale"),
    Handle:                   models.ToPointer("wholesale-handle"),
    PricingScheme:            models.ToPointer(models.PricingSchemeEnum("stairstep")),
    Prices:                   bodyPricePointPrices,
}

body := models.CreateComponentPricePointRequest{
    PricePoint: bodyPricePoint,
}

apiResponse, err := componentsController.CreateComponentPricePoint(ctx, componentId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# List Component Price Points

Use this endpoint to read current price points that are associated with a component.

You may specify the component by using either the numeric id or the `handle:gold` syntax.

When fetching a component's price points, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response.

If the price point is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.

```go
ListComponentPricePoints(
    ctx context.Context,input ListComponentPricePointsInput) (
    models.ApiResponse[models.ComponentPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `int` | Template, Required | The Chargify id of the component |
| `currencyPrices` | `*bool` | Query, Optional | Include an array of currency price data |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `20`<br>**Constraints**: `<= 200` |
| `filterType` | [`[]models.PricePointTypeEnum`](../models/price-point-type-enum.md) | Query, Optional | Use in query: `filter[type]=catalog,default`. |

## Response Type

[`models.ComponentPricePointsResponse`](../models/component-price-points-response.md)

## Example Usage

```go
ctx := context.Background()
componentId := 222
page := 2
perPage := 50Liquid error: Value cannot be null. (Parameter 'key')

apiResponse, err := componentsController.ListComponentPricePoints(ctx, componentId, nil, &page, &perPage, Liquid error: Value cannot be null. (Parameter 'key'))
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


# Create Component Price Points

Use this endpoint to create multiple component price points in one request.

```go
CreateComponentPricePoints(
    ctx context.Context,
    componentId string,
    body *models.CreateComponentPricePointsRequest) (
    models.ApiResponse[models.ComponentPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `string` | Template, Required | The Chargify id of the component for which you want to fetch price points. |
| `body` | [`*models.CreateComponentPricePointsRequest`](../models/create-component-price-points-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentPricePointsResponse`](../models/component-price-points-response.md)

## Example Usage

```go
ctx := context.Background()
componentId := "component_id8"

bodyPricePoints0Prices0 := models.Price{
    StartingQuantity: interface{}("[key1, val1][key2, val2]"),
    UnitPrice:        interface{}("[key1, val1][key2, val2]"),
}

bodyPricePoints0Prices := []models.Price{bodyPricePoints0Prices0}
bodyPricePoints0 := models.PricePoint{
    Name:                     models.ToPointer("Wholesale"),
    Handle:                   models.ToPointer("wholesale"),
    PricingScheme:            models.ToPointer(models.PricingSchemeEnum("per_unit")),
    Prices:                   bodyPricePoints0Prices,
}

bodyPricePoints1Prices0 := models.Price{
    StartingQuantity: interface{}("[key1, val1][key2, val2]"),
    UnitPrice:        interface{}("[key1, val1][key2, val2]"),
}

bodyPricePoints1Prices := []models.Price{bodyPricePoints1Prices0}
bodyPricePoints1 := models.PricePoint{
    Name:                     models.ToPointer("MSRP"),
    Handle:                   models.ToPointer("msrp"),
    PricingScheme:            models.ToPointer(models.PricingSchemeEnum("per_unit")),
    Prices:                   bodyPricePoints1Prices,
}

bodyPricePoints2Prices0 := models.Price{
    StartingQuantity: interface{}("[key1, val1][key2, val2]"),
    UnitPrice:        interface{}("[key1, val1][key2, val2]"),
}

bodyPricePoints2Prices := []models.Price{bodyPricePoints2Prices0}
bodyPricePoints2 := models.PricePoint{
    Name:                     models.ToPointer("Special Pricing"),
    Handle:                   models.ToPointer("special"),
    PricingScheme:            models.ToPointer(models.PricingSchemeEnum("per_unit")),
    Prices:                   bodyPricePoints2Prices,
}

bodyPricePoints := []models.PricePoint{bodyPricePoints0, bodyPricePoints1, bodyPricePoints2}
body := models.CreateComponentPricePointsRequest{
    PricePoints: bodyPricePoints,
}

apiResponse, err := componentsController.CreateComponentPricePoints(ctx, componentId, &body)
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


# Update Component Price Point

When updating a price point, it's prices can be updated as well by creating new prices or editing / removing existing ones.

Passing in a price bracket without an `id` will attempt to create a new price.

Including an `id` will update the corresponding price, and including the `_destroy` flag set to true along with the `id` will remove that price.

Note: Custom price points cannot be updated directly. They must be edited through the Subscription.

```go
UpdateComponentPricePoint(
    ctx context.Context,
    componentId int,
    pricePointId int,
    body *models.UpdateComponentPricePointRequest) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `int` | Template, Required | The Chargify id of the component to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |
| `body` | [`*models.UpdateComponentPricePointRequest`](../models/update-component-price-point-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentPricePointResponse`](../models/component-price-point-response.md)

## Example Usage

```go
ctx := context.Background()
componentId := 222
pricePointId := 10

bodyPricePointPrices0 := models.UpdatePrice{
    Id:               models.ToPointer(1),
    EndingQuantity:   models.ToPointer(100),
    UnitPrice:        models.ToPointer(5),
}

bodyPricePointPrices1 := models.UpdatePrice{
    Id:               models.ToPointer(2),
    Destroy:          models.ToPointer("true"),
}

bodyPricePointPrices2 := models.UpdatePrice{
    UnitPrice:        models.ToPointer(4),
    StartingQuantity: models.ToPointer(101),
}

bodyPricePointPrices := []models.UpdatePrice{bodyPricePointPrices0, bodyPricePointPrices1, bodyPricePointPrices2}
bodyPricePoint := models.UpdateComponentPricePoint{
    Name:   models.ToPointer("Default"),
    Prices: bodyPricePointPrices,
}

body := models.UpdateComponentPricePointRequest{
    PricePoint: models.ToPointer(bodyPricePoint),
}

apiResponse, err := componentsController.UpdateComponentPricePoint(ctx, componentId, pricePointId, &body)
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
    componentId int,
    pricePointId int) (
    models.ApiResponse[models.ComponentPricePointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `int` | Template, Required | The Chargify id of the component to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |

## Response Type

[`models.ComponentPricePointResponse`](../models/component-price-point-response.md)

## Example Usage

```go
ctx := context.Background()
componentId := 222
pricePointId := 10

apiResponse, err := componentsController.ArchiveComponentPricePoint(ctx, componentId, pricePointId)
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
| `componentId` | `int` | Template, Required | The Chargify id of the component to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |

## Response Type

[`models.ComponentPricePointResponse`](../models/component-price-point-response.md)

## Example Usage

```go
ctx := context.Background()
componentId := 222
pricePointId := 10

apiResponse, err := componentsController.UnarchiveComponentPricePoint(ctx, componentId, pricePointId)
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
    models.ApiResponse[[]models.CurrencyPrice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |
| `body` | [`*models.CreateCurrencyPricesRequest`](../models/create-currency-prices-request.md) | Body, Optional | - |

## Response Type

[`[]models.CurrencyPrice`](../models/currency-price.md)

## Example Usage

```go
ctx := context.Background()
pricePointId := 10

bodyCurrencyPrices0 := models.CreateCurrencyPrice{
    Currency: models.ToPointer("EUR"),
    Price:    models.ToPointer(50),
    PriceId:  models.ToPointer(20),
}

bodyCurrencyPrices1 := models.CreateCurrencyPrice{
    Currency: models.ToPointer("EUR"),
    Price:    models.ToPointer(40),
    PriceId:  models.ToPointer(21),
}

bodyCurrencyPrices := []models.CreateCurrencyPrice{bodyCurrencyPrices0, bodyCurrencyPrices1}
body := models.CreateCurrencyPricesRequest{
    CurrencyPrices: bodyCurrencyPrices,
}

apiResponse, err := componentsController.CreateCurrencyPrices(ctx, pricePointId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Currency Prices

This endpoint allows you to update currency prices for a given currency that has been defined on the site level in your settings.

Note: Currency Prices are not able to be updated for custom price points.

```go
UpdateCurrencyPrices(
    ctx context.Context,
    pricePointId int,
    body *models.UpdateCurrencyPricesRequest) (
    models.ApiResponse[[]models.CurrencyPrice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |
| `body` | [`*models.UpdateCurrencyPricesRequest`](../models/update-currency-prices-request.md) | Body, Optional | - |

## Response Type

[`[]models.CurrencyPrice`](../models/currency-price.md)

## Example Usage

```go
ctx := context.Background()
pricePointId := 10

bodyCurrencyPrices0 := models.UpdateCurrencyPrice{
    Id:    100,
    Price: 51,
}

bodyCurrencyPrices1 := models.UpdateCurrencyPrice{
    Id:    101,
    Price: 41,
}

bodyCurrencyPrices := []models.UpdateCurrencyPrice{bodyCurrencyPrices0, bodyCurrencyPrices1}
body := models.UpdateCurrencyPricesRequest{
    CurrencyPrices: bodyCurrencyPrices,
}

apiResponse, err := componentsController.UpdateCurrencyPrices(ctx, pricePointId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# List All Component Price Points

This method allows to retrieve a list of Components Price Points belonging to a Site.

```go
ListAllComponentPricePoints(
    ctx context.Context,input ListAllComponentPricePointsInput) (
    models.ApiResponse[models.ListComponentsPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `filterDateField` | [`*models.BasicDateFieldEnum`](../models/basic-date-field-enum.md) | Query, Optional | The type of filter you would like to apply to your search. Use in query: `filter[date_field]=created_at`. |
| `filterEndDate` | `*time.Time` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns price points with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `filterEndDatetime` | `*time.Time` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns price points with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |
| `include` | [`*models.ListComponentsPricePointsIncludeEnum`](../models/list-components-price-points-include-enum.md) | Query, Optional | Allows including additional data in the response. Use in query: `include=currency_prices`. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `20`<br>**Constraints**: `<= 200` |
| `filterStartDate` | `*time.Time` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns price points with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `filterStartDatetime` | `*time.Time` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns price points with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `filterType` | [`[]models.PricePointTypeEnum`](../models/price-point-type-enum.md) | Query, Optional | Allows fetching price points with matching type. Use in query: `filter[type]=custom,catalog`. |
| `direction` | [`*models.SortingDirectionEnum`](../models/sorting-direction-enum.md) | Query, Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |
| `filterIds` | `[]int` | Query, Optional | Allows fetching price points with matching id based on provided values. Use in query: `filter[ids]=1,2,3`. |
| `filterArchivedAt` | [`*models.IncludeNotNullEnum`](../models/include-not-null-enum.md) | Query, Optional | Allows fetching price points only if archived_at is present or not. Use in query: `filter[archived_at]=not_null`. |

## Response Type

[`models.ListComponentsPricePointsResponse`](../models/list-components-price-points-response.md)

## Example Usage

```go
ctx := context.Background()Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')
include := models.ListComponentsPricePointsIncludeEnum("currency_prices")
page := 2
perPage := 50Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')

apiResponse, err := componentsController.ListAllComponentPricePoints(ctx, Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), &include, &page, &perPage, Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), nil, Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'))
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
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../models/error-list-response-exception.md) |

