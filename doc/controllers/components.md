# Components

```go
componentsController := client.ComponentsController()
```

## Class Name

`ComponentsController`

## Methods

* [Create Metered Component](../../doc/controllers/components.md#create-metered-component)
* [Create Quantity Based Component](../../doc/controllers/components.md#create-quantity-based-component)
* [Create on Off Component](../../doc/controllers/components.md#create-on-off-component)
* [Create Prepaid Usage Component](../../doc/controllers/components.md#create-prepaid-usage-component)
* [Create Event Based Component](../../doc/controllers/components.md#create-event-based-component)
* [Find Component](../../doc/controllers/components.md#find-component)
* [Read Component](../../doc/controllers/components.md#read-component)
* [Update Product Family Component](../../doc/controllers/components.md#update-product-family-component)
* [Archive Component](../../doc/controllers/components.md#archive-component)
* [List Components](../../doc/controllers/components.md#list-components)
* [Update Component](../../doc/controllers/components.md#update-component)
* [Promote Component Price Point to Default](../../doc/controllers/components.md#promote-component-price-point-to-default)
* [List Components for Product Family](../../doc/controllers/components.md#list-components-for-product-family)
* [Create Component Price Point](../../doc/controllers/components.md#create-component-price-point)
* [List Component Price Points](../../doc/controllers/components.md#list-component-price-points)
* [Bulk Create Component Price Points](../../doc/controllers/components.md#bulk-create-component-price-points)
* [Update Component Price Point](../../doc/controllers/components.md#update-component-price-point)
* [Archive Component Price Point](../../doc/controllers/components.md#archive-component-price-point)
* [Unarchive Component Price Point](../../doc/controllers/components.md#unarchive-component-price-point)
* [Create Currency Prices](../../doc/controllers/components.md#create-currency-prices)
* [Update Currency Prices](../../doc/controllers/components.md#update-currency-prices)
* [List All Component Price Points](../../doc/controllers/components.md#list-all-component-price-points)


# Create Metered Component

This request will create a component definition of kind **metered_component** under the specified product family. Metered component can then be added and “allocated” for a subscription.

Metered components are used to bill for any type of unit that resets to 0 at the end of the billing period (think daily Google Adwords clicks or monthly cell phone minutes). This is most commonly associated with usage-based billing and many other pricing schemes.

Note that this is different from recurring quantity-based components, which DO NOT reset to zero at the start of every billing period. If you want to bill for a quantity of something that does not change unless you change it, then you want quantity components, instead.

For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).

```go
CreateMeteredComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateMeteredComponent) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the component belongs |
| `body` | [`*models.CreateMeteredComponent`](../../doc/models/create-metered-component.md) | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

body := models.CreateMeteredComponent{
    MeteredComponent: models.MeteredComponent{
        Name:                      "Text messages",
        UnitName:                  "text message",
        Taxable:                   models.ToPointer(false),
        PricingScheme:             models.PricingScheme("per_unit"),
        Prices:                    []models.Price{
            models.Price{
                StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(1),
                UnitPrice:        models.PriceUnitPriceContainer.FromPrecision(float64(1)),
            },
        },
    },
}

apiResponse, err := componentsController.CreateMeteredComponent(ctx, productFamilyId, &body)
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
    "handle": "text-messages",
    "pricing_scheme": "per_unit",
    "unit_name": "unit",
    "unit_price": "10.0",
    "product_family_id": 528484,
    "product_family_name": "Cloud Compute Servers",
    "price_per_unit_in_cents": null,
    "kind": "metered_component",
    "archived": false,
    "taxable": false,
    "description": null,
    "default_price_point_id": 2944263,
    "prices": [
      {
        "id": 55423,
        "component_id": 30002,
        "starting_quantity": 1,
        "ending_quantity": null,
        "unit_price": "10.0",
        "price_point_id": 2944263,
        "formatted_unit_price": "$10.00",
        "segment_id": null
      }
    ],
    "price_point_count": 1,
    "price_points_url": "https://demo-3238403362.chargify.com/components/30002/price_points",
    "default_price_point_name": "Original",
    "tax_code": null,
    "recurring": false,
    "upgrade_charge": null,
    "downgrade_credit": null,
    "created_at": "2024-01-23T06:08:05-05:00",
    "updated_at": "2024-01-23T06:08:05-05:00",
    "archived_at": null,
    "hide_date_range_on_invoice": false,
    "allow_fractional_quantities": false,
    "use_site_exchange_rate": true,
    "item_category": null,
    "accounting_code": null
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Create Quantity Based Component

This request will create a component definition of kind **quantity_based_component** under the specified product family. Quantity Based component can then be added and “allocated” for a subscription.

When defining Quantity Based component, You can choose one of 2 types:

#### Recurring

Recurring quantity-based components are used to bill for the number of some unit (think monthly software user licenses or the number of pairs of socks in a box-a-month club). This is most commonly associated with billing for user licenses, number of users, number of employees, etc.

#### One-time

One-time quantity-based components are used to create ad hoc usage charges that do not recur. For example, at the time of signup, you might want to charge your customer a one-time fee for onboarding or other services.

The allocated quantity for one-time quantity-based components immediately gets reset back to zero after the allocation is made.

For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).

```go
CreateQuantityBasedComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateQuantityBasedComponent) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the component belongs |
| `body` | [`*models.CreateQuantityBasedComponent`](../../doc/models/create-quantity-based-component.md) | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

body := models.CreateQuantityBasedComponent{
    QuantityBasedComponent: models.QuantityBasedComponent{
        Name:                      "Quantity Based Component",
        UnitName:                  "Component",
        Description:               models.ToPointer("Example of JSON per-unit component example"),
        Taxable:                   models.ToPointer(true),
        PricingScheme:             models.PricingScheme("per_unit"),
        UnitPrice:                 models.ToPointer(models.QuantityBasedComponentUnitPriceContainer.FromString("10")),
        DisplayOnHostedPage:       models.ToPointer(true),
        AllowFractionalQuantities: models.ToPointer(true),
        PublicSignupPageIds:       []int{
            323397,
        },
    },
}

apiResponse, err := componentsController.CreateQuantityBasedComponent(ctx, productFamilyId, &body)
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
    "handle": "text-messages",
    "pricing_scheme": "per_unit",
    "unit_name": "unit",
    "unit_price": "10.0",
    "product_family_id": 528484,
    "product_family_name": "Cloud Compute Servers",
    "price_per_unit_in_cents": null,
    "kind": "quantity_based_component",
    "archived": false,
    "taxable": false,
    "description": null,
    "default_price_point_id": 2944263,
    "prices": [
      {
        "id": 55423,
        "component_id": 30002,
        "starting_quantity": 1,
        "ending_quantity": null,
        "unit_price": "10.0",
        "price_point_id": 2944263,
        "formatted_unit_price": "$10.00",
        "segment_id": null
      }
    ],
    "price_point_count": 1,
    "price_points_url": "https://demo-3238403362.chargify.com/components/30002/price_points",
    "default_price_point_name": "Original",
    "tax_code": null,
    "recurring": false,
    "upgrade_charge": null,
    "downgrade_credit": null,
    "created_at": "2024-01-23T06:08:05-05:00",
    "updated_at": "2024-01-23T06:08:05-05:00",
    "archived_at": null,
    "hide_date_range_on_invoice": false,
    "allow_fractional_quantities": false,
    "use_site_exchange_rate": true,
    "item_category": null,
    "accounting_code": null
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Create on Off Component

This request will create a component definition of kind **on_off_component** under the specified product family. On/Off component can then be added and “allocated” for a subscription.

On/off components are used for any flat fee, recurring add on (think $99/month for tech support or a flat add on shipping fee).

For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).

```go
CreateOnOffComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateOnOffComponent) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the component belongs |
| `body` | [`*models.CreateOnOffComponent`](../../doc/models/create-on-off-component.md) | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

body := models.CreateOnOffComponent{
    OnOffComponent: models.OnOffComponent{
        Name:                      "Annual Support Services",
        Description:               models.ToPointer("Prepay for support services"),
        Taxable:                   models.ToPointer(true),
        Prices:                    []models.Price{
            models.Price{
                StartingQuantity: models.PriceStartingQuantityContainer.FromString("0"),
                UnitPrice:        models.PriceUnitPriceContainer.FromString("100.00"),
            },
        },
        DisplayOnHostedPage:       models.ToPointer(true),
        PublicSignupPageIds:       []int{
            320495,
        },
    },
}

apiResponse, err := componentsController.CreateOnOffComponent(ctx, productFamilyId, &body)
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
    "name": "Test On-Off Component 46124",
    "handle": "test-on-off-component-4612422802",
    "pricing_scheme": null,
    "unit_name": "on/off",
    "unit_price": "10.0",
    "product_family_id": 528484,
    "product_family_name": "Cloud Compute Servers",
    "price_per_unit_in_cents": null,
    "kind": "on_off_component",
    "archived": false,
    "taxable": false,
    "description": null,
    "default_price_point_id": 2944263,
    "price_point_count": 1,
    "price_points_url": "https://demo-3238403362.chargify.com/components/30002/price_points",
    "default_price_point_name": "Original",
    "tax_code": null,
    "recurring": true,
    "upgrade_charge": null,
    "downgrade_credit": null,
    "created_at": "2024-01-23T06:08:05-05:00",
    "updated_at": "2024-01-23T06:08:05-05:00",
    "archived_at": null,
    "hide_date_range_on_invoice": false,
    "allow_fractional_quantities": false,
    "use_site_exchange_rate": true,
    "item_category": null,
    "accounting_code": null
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Create Prepaid Usage Component

This request will create a component definition of kind **prepaid_usage_component** under the specified product family. Prepaid component can then be added and “allocated” for a subscription.

Prepaid components allow customers to pre-purchase units that can be used up over time on their subscription. In a sense, they are the mirror image of metered components; while metered components charge at the end of the period for the amount of units used, prepaid components are charged for at the time of purchase, and we subsequently keep track of the usage against the amount purchased.

For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).

```go
CreatePrepaidUsageComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreatePrepaidComponent) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the component belongs |
| `body` | [`*models.CreatePrepaidComponent`](../../doc/models/create-prepaid-component.md) | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

body := models.CreatePrepaidComponent{
    PrepaidUsageComponent: models.PrepaidUsageComponent{
        Name:                      "Minutes",
        UnitName:                  models.ToPointer("minutes"),
        PricingScheme:             models.ToPointer(models.PricingScheme("per_unit")),
        UnitPrice:                 models.ToPointer(models.PrepaidUsageComponentUnitPriceContainer.FromPrecision(float64(2))),
        OveragePricing:            models.ToPointer(models.OveragePricing{
            PricingScheme: models.PricingScheme("stairstep"),
            Prices:        []models.Price{
                models.Price{
                    StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(1),
                    EndingQuantity:   models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(100))),
                    UnitPrice:        models.PriceUnitPriceContainer.FromPrecision(float64(3)),
                },
                models.Price{
                    StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(101),
                    UnitPrice:        models.PriceUnitPriceContainer.FromPrecision(float64(5)),
                },
            },
        }),
        RolloverPrepaidRemainder:  models.ToPointer(true),
        RenewPrepaidAllocation:    models.ToPointer(true),
        ExpirationInterval:        models.ToPointer(float64(15)),
        ExpirationIntervalUnit:    models.ToPointer(models.IntervalUnit("day")),
    },
}

apiResponse, err := componentsController.CreatePrepaidUsageComponent(ctx, productFamilyId, &body)
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
    "name": "Test Prepaid Component 98505",
    "handle": "test-prepaid-component-9850584842",
    "pricing_scheme": "per_unit",
    "unit_name": "unit",
    "unit_price": "10.0",
    "product_family_id": 528484,
    "product_family_name": "Test Product Family 27791",
    "price_per_unit_in_cents": null,
    "kind": "prepaid_usage_component",
    "archived": false,
    "taxable": false,
    "description": "Description for: Test Prepaid Component 98505",
    "default_price_point_id": 2944263,
    "overage_prices": [
      {
        "id": 55964,
        "component_id": 30427,
        "starting_quantity": 1,
        "ending_quantity": null,
        "unit_price": "1.0",
        "price_point_id": 2944756,
        "formatted_unit_price": "$1.00",
        "segment_id": null
      }
    ],
    "prices": [
      {
        "id": 55963,
        "component_id": 30427,
        "starting_quantity": 1,
        "ending_quantity": null,
        "unit_price": "1.0",
        "price_point_id": 2944756,
        "formatted_unit_price": "$1.00",
        "segment_id": null
      }
    ],
    "price_point_count": 1,
    "price_points_url": "https://demo-3238403362.chargify.com/components/30002/price_points",
    "default_price_point_name": "Original",
    "tax_code": null,
    "recurring": true,
    "upgrade_charge": null,
    "downgrade_credit": null,
    "created_at": "2024-01-23T06:08:05-05:00",
    "updated_at": "2024-01-23T06:08:05-05:00",
    "archived_at": null,
    "hide_date_range_on_invoice": false,
    "allow_fractional_quantities": false,
    "use_site_exchange_rate": true,
    "item_category": null,
    "accounting_code": null
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Create Event Based Component

This request will create a component definition of kind **event_based_component** under the specified product family. Event-based component can then be added and “allocated” for a subscription.

Event-based components are similar to other component types, in that you define the component parameters (such as name and taxability) and the pricing. A key difference for the event-based component is that it must be attached to a metric. This is because the metric provides the component with the actual quantity used in computing what and how much will be billed each period for each subscription.

So, instead of reporting usage directly for each component (as you would with metered components), the usage is derived from analysis of your events.

For more information on components, please see our documentation [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405020625677).

```go
CreateEventBasedComponent(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateEBBComponent) (
    models.ApiResponse[models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the component belongs |
| `body` | [`*models.CreateEBBComponent`](../../doc/models/create-ebb-component.md) | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

body := models.CreateEBBComponent{
    EventBasedComponent: models.EBBComponent{
        Name:                      "Component Name",
        UnitName:                  "string",
        Description:               models.ToPointer("string"),
        Handle:                    models.ToPointer("some_handle"),
        Taxable:                   models.ToPointer(true),
        PricingScheme:             models.PricingScheme("per_unit"),
        Prices:                    []models.Price{
            models.Price{
                StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(1),
                UnitPrice:        models.PriceUnitPriceContainer.FromString("0.49"),
            },
        },
        EventBasedBillingMetricId: 123,
    },
}

apiResponse, err := componentsController.CreateEventBasedComponent(ctx, productFamilyId, &body)
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
    "id": 1489581,
    "name": "stripeCharges",
    "handle": null,
    "pricing_scheme": null,
    "unit_name": "charge",
    "unit_price": null,
    "product_family_id": 1517093,
    "product_family_name": "Billing Plans",
    "price_per_unit_in_cents": null,
    "kind": "event_based_component",
    "archived": false,
    "taxable": false,
    "description": null,
    "default_price_point_id": null,
    "price_point_count": 0,
    "price_points_url": "https://staging.chargify.com/components/1489581/price_points",
    "default_price_point_name": "Original",
    "tax_code": null,
    "recurring": false,
    "upgrade_charge": null,
    "downgrade_credit": null,
    "created_at": "2021-10-12T07:33:24-05:00",
    "updated_at": "2021-10-12T07:33:24-05:00",
    "archived_at": null,
    "hide_date_range_on_invoice": false,
    "allow_fractional_quantities": false,
    "use_site_exchange_rate": null,
    "item_category": null,
    "accounting_code": null,
    "event_based_billing_metric_id": 1163
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Find Component

This request will return information regarding a component having the handle you provide. You can identify your components with a handle so you don't have to save or reference the IDs we generate.

```go
FindComponent(
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

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

handle := "handle6"

apiResponse, err := componentsController.FindComponent(ctx, handle)
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


# Read Component

This request will return information regarding a component from a specific product family.

You may read the component by either the component's id or handle. When using the handle, it must be prefixed with `handle:`.

```go
ReadComponent(
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
| `componentId` | `string` | Template, Required | Either the Chargify id of the component or the handle for the component prefixed with `handle:` |

## Response Type

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

componentId := "component_id8"

apiResponse, err := componentsController.ReadComponent(ctx, productFamilyId, componentId)
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
| `componentId` | `string` | Template, Required | Either the Chargify id of the component or the handle for the component prefixed with `handle:` |
| `body` | [`*models.UpdateComponentRequest`](../../doc/models/update-component-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

componentId := "component_id8"

body := models.UpdateComponentRequest{
    Component: models.UpdateComponent{
        ItemCategory:        models.NewOptional(models.ToPointer(models.ItemCategory("Business Software"))),
    },
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
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


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
| `componentId` | `string` | Template, Required | Either the Chargify id of the component or the handle for the component prefixed with `handle:` |

## Response Type

[`models.Component`](../../doc/models/component.md)

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
  "created_at": "2016-11-08T16:22:26-05:00",
  "default_price_point_name": "cupidatat Lorem non aliqua",
  "product_family_name": "do elit",
  "hide_date_range_on_invoice": false
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Components

This request will return a list of components for a site.

```go
ListComponents(
    ctx context.Context,
    input ListComponentsInput) (
    models.ApiResponse[[]models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `dateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Query, Optional | The type of filter you would like to apply to your search. |
| `startDate` | `*string` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `endDate` | `*string` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `startDatetime` | `*string` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `endDatetime` | `*string` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.  optional |
| `includeArchived` | `*bool` | Query, Optional | Include archived items |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `filter` | [`*models.ListComponentsFilter`](../../doc/models/list-components-filter.md) | Query, Optional | Filter to use for List Components operations |

## Response Type

[`[]models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListComponentsInput{
    DateField:       models.ToPointer(models.BasicDateField("updated_at")),
    Page:            models.ToPointer(2),
    PerPage:         models.ToPointer(50),
    Filter:          models.ToPointer(models.ListComponentsFilter{
        Ids:                 []int{
            1,
            2,
            3,
        },
    }),
}

apiResponse, err := componentsController.ListComponents(ctx, collectedInput)
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
| `body` | [`*models.UpdateComponentRequest`](../../doc/models/update-component-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

componentId := "component_id8"

body := models.UpdateComponentRequest{
    Component: models.UpdateComponent{
        ItemCategory:        models.NewOptional(models.ToPointer(models.ItemCategory("Business Software"))),
    },
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
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Promote Component Price Point to Default

Sets a new default price point for the component. This new default will apply to all new subscriptions going forward - existing subscriptions will remain on their current price point.

See [Price Points Documentation](https://chargify.zendesk.com/hc/en-us/articles/4407755865883#price-points) for more information on price points and moving subscriptions between price points.

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
| `componentId` | `int` | Template, Required | The Chargify id of the component to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |

## Response Type

[`models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

componentId := 222

pricePointId := 10

apiResponse, err := componentsController.PromoteComponentPricePointToDefault(ctx, componentId, pricePointId)
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
    ctx context.Context,
    input ListComponentsForProductFamilyInput) (
    models.ApiResponse[[]models.ComponentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family |
| `includeArchived` | `*bool` | Query, Optional | Include archived items. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `filter` | [`*models.ListComponentsFilter`](../../doc/models/list-components-filter.md) | Query, Optional | Filter to use for List Components operations |
| `dateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Query, Optional | The type of filter you would like to apply to your search. Use in query `date_field=created_at`. |
| `endDate` | `*string` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `endDatetime` | `*string` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. optional. |
| `startDate` | `*string` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `startDatetime` | `*string` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |

## Response Type

[`[]models.ComponentResponse`](../../doc/models/component-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListComponentsForProductFamilyInput{
    ProductFamilyId: 140,
    Page:            models.ToPointer(2),
    PerPage:         models.ToPointer(50),
    Filter:          models.ToPointer(models.ListComponentsFilter{
        Ids:                 []int{
            1,
            2,
            3,
        },
    }),
    DateField:       models.ToPointer(models.BasicDateField("updated_at")),
}

apiResponse, err := componentsController.ListComponentsForProductFamily(ctx, collectedInput)
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
| `body` | [`*models.CreateComponentPricePointRequest`](../../doc/models/create-component-price-point-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentPricePointResponse`](../../doc/models/component-price-point-response.md)

## Example Usage

```go
ctx := context.Background()

componentId := 222

body := models.CreateComponentPricePointRequest{
    PricePoint: models.CreateComponentPricePointRequestPricePointContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
        Name:                "Wholesale",
        Handle:              models.ToPointer("wholesale-handle"),
        PricingScheme:       models.PricingScheme("stairstep"),
        Prices:              []models.Price{
            models.Price{
                StartingQuantity: models.PriceStartingQuantityContainer.FromString("1"),
                EndingQuantity:   models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromString("100"))),
                UnitPrice:        models.PriceUnitPriceContainer.FromString("5.00"),
            },
            models.Price{
                StartingQuantity: models.PriceStartingQuantityContainer.FromString("101"),
                UnitPrice:        models.PriceUnitPriceContainer.FromString("4.00"),
            },
        },
        UseSiteExchangeRate: models.ToPointer(false),
    }),
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
    ctx context.Context,
    input ListComponentPricePointsInput) (
    models.ApiResponse[models.ComponentPricePointsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `int` | Template, Required | The Chargify id of the component |
| `currencyPrices` | `*bool` | Query, Optional | Include an array of currency price data |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `filterType` | [`[]models.PricePointType`](../../doc/models/price-point-type.md) | Query, Optional | Use in query: `filter[type]=catalog,default`. |

## Response Type

[`models.ComponentPricePointsResponse`](../../doc/models/component-price-points-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListComponentPricePointsInput{
    ComponentId:    222,
    Page:           models.ToPointer(2),
    PerPage:        models.ToPointer(50),
Liquid error: Value cannot be null. (Parameter 'key')}

apiResponse, err := componentsController.ListComponentPricePoints(ctx, collectedInput)
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
| `componentId` | `string` | Template, Required | The Chargify id of the component for which you want to fetch price points. |
| `body` | [`*models.CreateComponentPricePointsRequest`](../../doc/models/create-component-price-points-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentPricePointsResponse`](../../doc/models/component-price-points-response.md)

## Example Usage

```go
ctx := context.Background()

componentId := "component_id8"

body := models.CreateComponentPricePointsRequest{
    PricePoints: []models.CreateComponentPricePointsRequestPricePoints{
        models.CreateComponentPricePointsRequestPricePointsContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
            Name:                "Wholesale",
            Handle:              models.ToPointer("wholesale"),
            PricingScheme:       models.PricingScheme("per_unit"),
            Prices:              []models.Price{
                models.Price{
                    StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(1),
                    UnitPrice:        models.PriceUnitPriceContainer.FromPrecision(float64(5)),
                },
            },
        }),
        models.CreateComponentPricePointsRequestPricePointsContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
            Name:                "MSRP",
            Handle:              models.ToPointer("msrp"),
            PricingScheme:       models.PricingScheme("per_unit"),
            Prices:              []models.Price{
                models.Price{
                    StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(1),
                    UnitPrice:        models.PriceUnitPriceContainer.FromPrecision(float64(4)),
                },
            },
        }),
        models.CreateComponentPricePointsRequestPricePointsContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
            Name:                "Special Pricing",
            Handle:              models.ToPointer("special"),
            PricingScheme:       models.PricingScheme("per_unit"),
            Prices:              []models.Price{
                models.Price{
                    StartingQuantity: models.PriceStartingQuantityContainer.FromNumber(1),
                    UnitPrice:        models.PriceUnitPriceContainer.FromPrecision(float64(5)),
                },
            },
        }),
    },
}

apiResponse, err := componentsController.BulkCreateComponentPricePoints(ctx, componentId, &body)
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
| `body` | [`*models.UpdateComponentPricePointRequest`](../../doc/models/update-component-price-point-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentPricePointResponse`](../../doc/models/component-price-point-response.md)

## Example Usage

```go
ctx := context.Background()

componentId := 222

pricePointId := 10

body := models.UpdateComponentPricePointRequest{
    PricePoint: models.ToPointer(models.UpdateComponentPricePoint{
        Name:                models.ToPointer("Default"),
        Prices:              []models.UpdatePrice{
            models.UpdatePrice{
                Id:               models.ToPointer(1),
                EndingQuantity:   models.ToPointer(models.UpdatePriceEndingQuantityContainer.FromNumber(100)),
                UnitPrice:        models.ToPointer(models.UpdatePriceUnitPriceContainer.FromPrecision(float64(5))),
            },
            models.UpdatePrice{
                Id:               models.ToPointer(2),
                Destroy:          models.ToPointer(true),
            },
            models.UpdatePrice{
                UnitPrice:        models.ToPointer(models.UpdatePriceUnitPriceContainer.FromPrecision(float64(4))),
                StartingQuantity: models.ToPointer(models.UpdatePriceStartingQuantityContainer.FromNumber(101)),
            },
        },
    }),
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorArrayMapResponseException`](../../doc/models/error-array-map-response-exception.md) |


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

[`models.ComponentPricePointResponse`](../../doc/models/component-price-point-response.md)

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
| `componentId` | `int` | Template, Required | The Chargify id of the component to which the price point belongs |
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |

## Response Type

[`models.ComponentPricePointResponse`](../../doc/models/component-price-point-response.md)

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
    models.ApiResponse[models.ComponentCurrencyPricesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |
| `body` | [`*models.CreateCurrencyPricesRequest`](../../doc/models/create-currency-prices-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentCurrencyPricesResponse`](../../doc/models/component-currency-prices-response.md)

## Example Usage

```go
ctx := context.Background()

pricePointId := 10

body := models.CreateCurrencyPricesRequest{
    CurrencyPrices: []models.CreateCurrencyPrice{
        models.CreateCurrencyPrice{
            Currency: models.ToPointer("EUR"),
            Price:    models.ToPointer(float64(50)),
            PriceId:  models.ToPointer(20),
        },
        models.CreateCurrencyPrice{
            Currency: models.ToPointer("EUR"),
            Price:    models.ToPointer(float64(40)),
            PriceId:  models.ToPointer(21),
        },
    },
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
| `pricePointId` | `int` | Template, Required | The Chargify id of the price point |
| `body` | [`*models.UpdateCurrencyPricesRequest`](../../doc/models/update-currency-prices-request.md) | Body, Optional | - |

## Response Type

[`models.ComponentCurrencyPricesResponse`](../../doc/models/component-currency-prices-response.md)

## Example Usage

```go
ctx := context.Background()

pricePointId := 10

body := models.UpdateCurrencyPricesRequest{
    CurrencyPrices: []models.UpdateCurrencyPrice{
        models.UpdateCurrencyPrice{
            Id:    100,
            Price: 51,
        },
        models.UpdateCurrencyPrice{
            Id:    101,
            Price: 41,
        },
    },
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
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Query, Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |
| `filter` | [`*models.ListPricePointsFilter`](../../doc/models/list-price-points-filter.md) | Query, Optional | Filter to use for List PricePoints operations |

## Response Type

[`models.ListComponentsPricePointsResponse`](../../doc/models/list-components-price-points-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListAllComponentPricePointsInput{
    Include:   models.ToPointer(models.ListComponentsPricePointsInclude("currency_prices")),
    Page:      models.ToPointer(2),
    PerPage:   models.ToPointer(50),
    Filter:    models.ToPointer(models.ListPricePointsFilter{
        StartDate:     models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-17", func(err error) { log.Fatalln(err) })),
        EndDate:       models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-15", func(err error) { log.Fatalln(err) })),
        StartDatetime: models.ToPointer(parseTime(time.RFC3339, "12/19/2011 09:15:30", func(err error) { log.Fatalln(err) })),
        EndDatetime:   models.ToPointer(parseTime(time.RFC3339, "06/07/2019 17:20:06", func(err error) { log.Fatalln(err) })),
        Type:          []models.PricePointType{
            models.PricePointType("catalog"),
            models.PricePointType("default"),
            models.PricePointType("custom"),
        },
        Ids:           []int{
            1,
            2,
            3,
        },
    }),
}

apiResponse, err := componentsController.ListAllComponentPricePoints(ctx, collectedInput)
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

