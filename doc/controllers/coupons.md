# Coupons

```go
couponsController := client.CouponsController()
```

## Class Name

`CouponsController`

## Methods

* [Create Coupon](../../doc/controllers/coupons.md#create-coupon)
* [List Coupons for Product Family](../../doc/controllers/coupons.md#list-coupons-for-product-family)
* [Find Coupon](../../doc/controllers/coupons.md#find-coupon)
* [Read Coupon](../../doc/controllers/coupons.md#read-coupon)
* [Update Coupon](../../doc/controllers/coupons.md#update-coupon)
* [Archive Coupon](../../doc/controllers/coupons.md#archive-coupon)
* [List Coupons](../../doc/controllers/coupons.md#list-coupons)
* [Read Coupon Usage](../../doc/controllers/coupons.md#read-coupon-usage)
* [Validate Coupon](../../doc/controllers/coupons.md#validate-coupon)
* [Create or Update Coupon Currency Prices](../../doc/controllers/coupons.md#create-or-update-coupon-currency-prices)
* [Create Coupon Subcodes](../../doc/controllers/coupons.md#create-coupon-subcodes)
* [List Coupon Subcodes](../../doc/controllers/coupons.md#list-coupon-subcodes)
* [Update Coupon Subcodes](../../doc/controllers/coupons.md#update-coupon-subcodes)
* [Delete Coupon Subcode](../../doc/controllers/coupons.md#delete-coupon-subcode)


# Create Coupon

## Coupons Documentation

Coupons can be administered in the Advanced Billing application or created via API. Please view our section on [creating coupons](https://maxio.zendesk.com/hc/en-us/articles/24261212433165-Creating-Editing-Deleting-Coupons) for more information.

Additionally, for documentation on how to apply a coupon to a subscription within the Advanced Billing UI, please see our documentation [here](https://maxio.zendesk.com/hc/en-us/articles/24261259337101-Coupons-and-Subscriptions).

## Create Coupon

This request will create a coupon, based on the provided information.

You can create either a flat amount coupon, by specyfing `amount_in_cents`, or percentage coupon by specyfing `percentage`.

You can restrict a coupon to only apply to specific products / components by optionally passing in `restricted_products` and/or `restricted_components` objects in the format:
`{ "<product_id/component_id>": boolean_value }`

```go
CreateCoupon(
    ctx context.Context,
    productFamilyId int,
    body *models.CouponRequest) (
    models.ApiResponse[models.CouponResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Advanced Billing id of the product family to which the coupon belongs |
| `body` | [`*models.CouponRequest`](../../doc/models/coupon-request.md) | Body, Optional | - |

## Response Type

[`models.CouponResponse`](../../doc/models/coupon-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

body := models.CouponRequest{
    Coupon:               models.ToPointer(models.CouponPayload{
        Name:                          models.ToPointer("15% off"),
        Code:                          models.ToPointer("15OFF"),
        Description:                   models.ToPointer("15% off for life"),
        Percentage:                    models.ToPointer(models.CouponPayloadPercentageContainer.FromPrecision(float64(15))),
        AllowNegativeBalance:          models.ToPointer(false),
        Recurring:                     models.ToPointer(false),
        EndDate:                       models.ToPointer(parseTime(models.DEFAULT_DATE, "2012-08-29", func(err error) { log.Fatalln(err) })),
        ProductFamilyId:               models.ToPointer("2"),
        Stackable:                     models.ToPointer(true),
        CompoundingStrategy:           models.ToPointer(models.CompoundingStrategy("compound")),
        ExcludeMidPeriodAllocations:   models.ToPointer(true),
        ApplyOnCancelAtEndOfPeriod:    models.ToPointer(true),
    }),
    RestrictedProducts:   map[string]bool{
        "1": true,
    },
    RestrictedComponents: map[string]bool{
        "1": true,
        "2": false,
    },
}

apiResponse, err := couponsController.CreateCoupon(ctx, productFamilyId, &body)
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
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Coupons for Product Family

List coupons for a specific Product Family in a Site.

If the coupon is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.

```go
ListCouponsForProductFamily(
    ctx context.Context,
    input ListCouponsForProductFamilyInput) (
    models.ApiResponse[[]models.CouponResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Advanced Billing id of the product family to which the coupon belongs |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 30. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `30`<br>**Constraints**: `<= 200` |
| `filter` | [`*models.ListCouponsFilter`](../../doc/models/list-coupons-filter.md) | Query, Optional | Filter to use for List Coupons operations |
| `currencyPrices` | `*bool` | Query, Optional | When fetching coupons, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response. Use in query `currency_prices=true`. |

## Response Type

[`[]models.CouponResponse`](../../doc/models/coupon-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListCouponsForProductFamilyInput{
    ProductFamilyId: 140,
    Page:            models.ToPointer(2),
    PerPage:         models.ToPointer(50),
    Filter:          models.ToPointer(models.ListCouponsFilter{
        StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-17", func(err error) { log.Fatalln(err) })),
        EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-15", func(err error) { log.Fatalln(err) })),
        StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "12/19/2011 09:15:30", func(err error) { log.Fatalln(err) })),
        EndDatetime:          models.ToPointer(parseTime(time.RFC3339, "06/07/2019 17:20:06", func(err error) { log.Fatalln(err) })),
        Ids:                  []int{
            1,
            2,
            3,
        },
        Codes:                []string{
            "free",
            "free_trial",
        },
    }),
    CurrencyPrices:  models.ToPointer(true),
}

apiResponse, err := couponsController.ListCouponsForProductFamily(ctx, collectedInput)
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
    "coupon": {
      "id": 999999,
      "name": "50% coupon",
      "code": "50PERCENT",
      "description": "50 PERCENT OFF",
      "amount_in_cents": null,
      "product_family_id": 527890,
      "created_at": "2016-10-21T17:02:08-04:00",
      "updated_at": "2016-10-21T17:06:11-04:00",
      "start_date": "2016-10-21T17:02:08-04:00",
      "end_date": null,
      "percentage": "50",
      "recurring": true,
      "duration_period_count": null,
      "duration_interval": 1,
      "duration_interval_unit": "day",
      "allow_negative_balance": true,
      "archived_at": null,
      "conversion_limit": "100",
      "stackable": false,
      "compounding_strategy": "compound",
      "use_site_exchange_rate": true
    }
  },
  {
    "coupon": {
      "id": 123456,
      "name": "100% coupon",
      "code": "100PERCENT",
      "description": "100 PERCENT OFF",
      "amount_in_cents": null,
      "product_family_id": 527890,
      "created_at": "2016-10-21T17:02:08-04:00",
      "updated_at": "2016-10-21T17:06:11-04:00",
      "start_date": "2016-10-21T17:02:08-04:00",
      "end_date": null,
      "percentage": "50",
      "recurring": true,
      "duration_period_count": null,
      "duration_interval": 1,
      "duration_interval_unit": "day",
      "allow_negative_balance": true,
      "archived_at": null,
      "conversion_limit": "100",
      "stackable": false,
      "compounding_strategy": "compound",
      "use_site_exchange_rate": true
    }
  },
  {
    "coupon": {
      "id": 888888,
      "name": "25% coupon",
      "code": "25PERCENT",
      "description": "25 PERCENT OFF",
      "amount_in_cents": null,
      "product_family_id": 527890,
      "created_at": "2016-10-21T17:02:08-04:00",
      "updated_at": "2016-10-21T17:06:11-04:00",
      "start_date": "2016-10-21T17:02:08-04:00",
      "end_date": null,
      "percentage": "25",
      "recurring": true,
      "duration_period_count": null,
      "duration_interval": 1,
      "duration_interval_unit": "day",
      "allow_negative_balance": true,
      "archived_at": null,
      "conversion_limit": "100",
      "stackable": false,
      "compounding_strategy": "compound",
      "coupon_restrictions": [
        {
          "id": 37,
          "item_type": "Component",
          "item_id": 519,
          "name": "test",
          "handle": null
        }
      ],
      "use_site_exchange_rate": true
    }
  }
]
```


# Find Coupon

You can search for a coupon via the API with the find method. By passing a code parameter, the find will attempt to locate a coupon that matches that code. If no coupon is found, a 404 is returned.

If you have more than one product family and if the coupon you are trying to find does not belong to the default product family in your site, then you will need to specify (either in the url or as a query string param) the product family id.

```go
FindCoupon(
    ctx context.Context,
    productFamilyId *int,
    code *string,
    currencyPrices *bool) (
    models.ApiResponse[models.CouponResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `*int` | Query, Optional | The Advanced Billing id of the product family to which the coupon belongs |
| `code` | `*string` | Query, Optional | The code of the coupon |
| `currencyPrices` | `*bool` | Query, Optional | When fetching coupons, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response. |

## Response Type

[`models.CouponResponse`](../../doc/models/coupon-response.md)

## Example Usage

```go
ctx := context.Background()





currencyPrices := true

apiResponse, err := couponsController.FindCoupon(ctx, nil, nil, &currencyPrices)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Read Coupon

You can retrieve the Coupon via the API with the Show method. You must identify the Coupon in this call by the ID parameter that Advanced Billing assigns.
If instead you would like to find a Coupon using a Coupon code, see the Coupon Find method.

When fetching a coupon, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response.

If the coupon is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.

```go
ReadCoupon(
    ctx context.Context,
    productFamilyId int,
    couponId int,
    currencyPrices *bool) (
    models.ApiResponse[models.CouponResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Advanced Billing id of the product family to which the coupon belongs |
| `couponId` | `int` | Template, Required | The Advanced Billing id of the coupon |
| `currencyPrices` | `*bool` | Query, Optional | When fetching coupons, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response. |

## Response Type

[`models.CouponResponse`](../../doc/models/coupon-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

couponId := 162

currencyPrices := true

apiResponse, err := couponsController.ReadCoupon(ctx, productFamilyId, couponId, &currencyPrices)
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
  "coupon": {
    "id": 67,
    "name": "Foo Bar",
    "code": "YEPPER99934",
    "description": "my cool coupon",
    "amount_in_cents": null,
    "product_family_id": 4,
    "product_family_name": "Billing Plans",
    "created_at": "2017-11-08T10:01:15-05:00",
    "updated_at": "2017-11-08T10:01:15-05:00",
    "start_date": "2017-11-08T10:01:15-05:00",
    "end_date": null,
    "percentage": "33.3333",
    "duration_period_count": null,
    "duration_interval": null,
    "duration_interval_unit": null,
    "allow_negative_balance": false,
    "archived_at": null,
    "conversion_limit": null,
    "stackable": true,
    "compounding_strategy": "compound"
  }
}
```


# Update Coupon

## Update Coupon

You can update a Coupon via the API with a PUT request to the resource endpoint.

You can restrict a coupon to only apply to specific products / components by optionally passing in hashes of `restricted_products` and/or `restricted_components` in the format:
`{ "<product/component_id>": boolean_value }`

```go
UpdateCoupon(
    ctx context.Context,
    productFamilyId int,
    couponId int,
    body *models.CouponRequest) (
    models.ApiResponse[models.CouponResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Advanced Billing id of the product family to which the coupon belongs |
| `couponId` | `int` | Template, Required | The Advanced Billing id of the coupon |
| `body` | [`*models.CouponRequest`](../../doc/models/coupon-request.md) | Body, Optional | - |

## Response Type

[`models.CouponResponse`](../../doc/models/coupon-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

couponId := 162

body := models.CouponRequest{
    Coupon:               models.ToPointer(models.CouponPayload{
        Name:                          models.ToPointer("15% off"),
        Code:                          models.ToPointer("15OFF"),
        Description:                   models.ToPointer("15% off for life"),
        Percentage:                    models.ToPointer(models.CouponPayloadPercentageContainer.FromPrecision(float64(15))),
        AllowNegativeBalance:          models.ToPointer(false),
        Recurring:                     models.ToPointer(false),
        EndDate:                       models.ToPointer(parseTime(models.DEFAULT_DATE, "2012-08-29", func(err error) { log.Fatalln(err) })),
        ProductFamilyId:               models.ToPointer("2"),
        Stackable:                     models.ToPointer(true),
        CompoundingStrategy:           models.ToPointer(models.CompoundingStrategy("compound")),
    }),
    RestrictedProducts:   map[string]bool{
        "1": true,
    },
    RestrictedComponents: map[string]bool{
        "1": true,
        "2": false,
    },
}

apiResponse, err := couponsController.UpdateCoupon(ctx, productFamilyId, couponId, &body)
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
  "coupon": {
    "id": 67,
    "name": "Foo Bar",
    "code": "YEPPER99934",
    "description": "my cool coupon",
    "amount_in_cents": 10000,
    "product_family_id": 4,
    "created_at": "2017-11-08T10:01:15-05:00",
    "updated_at": "2017-11-08T10:01:15-05:00",
    "start_date": "2017-11-08T10:01:15-05:00",
    "end_date": null,
    "percentage": null,
    "recurring": false,
    "duration_period_count": null,
    "duration_interval": null,
    "duration_interval_unit": null,
    "allow_negative_balance": false,
    "archived_at": null,
    "conversion_limit": null,
    "stackable": true,
    "compounding_strategy": "compound"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Archive Coupon

You can archive a Coupon via the API with the archive method.
Archiving makes that Coupon unavailable for future use, but allows it to remain attached and functional on existing Subscriptions that are using it.
The `archived_at` date and time will be assigned.

```go
ArchiveCoupon(
    ctx context.Context,
    productFamilyId int,
    couponId int) (
    models.ApiResponse[models.CouponResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Advanced Billing id of the product family to which the coupon belongs |
| `couponId` | `int` | Template, Required | The Advanced Billing id of the coupon |

## Response Type

[`models.CouponResponse`](../../doc/models/coupon-response.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

couponId := 162

apiResponse, err := couponsController.ArchiveCoupon(ctx, productFamilyId, couponId)
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
  "coupon": {
    "id": 67,
    "name": "Foo Bar",
    "code": "YEPPER99934",
    "description": "my cool coupon",
    "amount_in_cents": 10000,
    "product_family_id": 4,
    "created_at": "2017-11-08T10:01:15-05:00",
    "updated_at": "2017-11-08T10:01:15-05:00",
    "start_date": "2017-11-08T10:01:15-05:00",
    "end_date": null,
    "percentage": null,
    "recurring": false,
    "duration_period_count": null,
    "duration_interval": null,
    "duration_interval_unit": null,
    "allow_negative_balance": false,
    "archived_at": "2016-12-02T13:09:33-05:00",
    "conversion_limit": null,
    "stackable": true,
    "compounding_strategy": "compound"
  }
}
```


# List Coupons

You can retrieve a list of coupons.

If the coupon is set to `use_site_exchange_rate: true`, it will return pricing based on the current exchange rate. If the flag is set to false, it will return all of the defined prices for each currency.

```go
ListCoupons(
    ctx context.Context,
    input ListCouponsInput) (
    models.ApiResponse[[]models.CouponResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 30. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `30`<br>**Constraints**: `<= 200` |
| `filter` | [`*models.ListCouponsFilter`](../../doc/models/list-coupons-filter.md) | Query, Optional | Filter to use for List Coupons operations |
| `currencyPrices` | `*bool` | Query, Optional | When fetching coupons, if you have defined multiple currencies at the site level, you can optionally pass the `?currency_prices=true` query param to include an array of currency price data in the response. Use in query `currency_prices=true`. |

## Response Type

[`[]models.CouponResponse`](../../doc/models/coupon-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListCouponsInput{
    Page:           models.ToPointer(2),
    PerPage:        models.ToPointer(50),
    Filter:         models.ToPointer(models.ListCouponsFilter{
        StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-17", func(err error) { log.Fatalln(err) })),
        EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2011-12-15", func(err error) { log.Fatalln(err) })),
        StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "12/19/2011 09:15:30", func(err error) { log.Fatalln(err) })),
        EndDatetime:          models.ToPointer(parseTime(time.RFC3339, "06/07/2019 17:20:06", func(err error) { log.Fatalln(err) })),
        Ids:                  []int{
            1,
            2,
            3,
        },
        Codes:                []string{
            "free",
            "free_trial",
        },
    }),
    CurrencyPrices: models.ToPointer(true),
}

apiResponse, err := couponsController.ListCoupons(ctx, collectedInput)
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
    "coupon": {
      "id": 0,
      "name": "string",
      "code": "string",
      "description": "string",
      "amount": 0,
      "amount_in_cents": 0,
      "product_family_id": 0,
      "product_family_name": "string",
      "start_date": "2021-05-03T16:00:21-04:00",
      "end_date": "2023-05-05T16:00:21-04:00",
      "percentage": "10",
      "recurring": true,
      "recurring_scheme": "do_not_recur",
      "duration_period_count": 0,
      "duration_interval": 0,
      "duration_interval_unit": "string",
      "duration_interval_span": "string",
      "allow_negative_balance": true,
      "archived_at": null,
      "conversion_limit": "string",
      "stackable": true,
      "compounding_strategy": "compound",
      "use_site_exchange_rate": true,
      "created_at": "2021-05-05T16:00:21-04:00",
      "updated_at": "2021-05-05T16:00:21-04:00",
      "discount_type": "amount",
      "exclude_mid_period_allocations": true,
      "apply_on_cancel_at_end_of_period": true,
      "coupon_restrictions": [
        {
          "id": 0,
          "item_type": "Component",
          "item_id": 0,
          "name": "string",
          "handle": "string"
        }
      ]
    }
  }
]
```


# Read Coupon Usage

This request will provide details about the coupon usage as an array of data hashes, one per product.

```go
ReadCouponUsage(
    ctx context.Context,
    productFamilyId int,
    couponId int) (
    models.ApiResponse[[]models.CouponUsage],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Advanced Billing id of the product family to which the coupon belongs |
| `couponId` | `int` | Template, Required | The Advanced Billing id of the coupon |

## Response Type

[`[]models.CouponUsage`](../../doc/models/coupon-usage.md)

## Example Usage

```go
ctx := context.Background()

productFamilyId := 140

couponId := 162

apiResponse, err := couponsController.ReadCouponUsage(ctx, productFamilyId, couponId)
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
    "name": "No cost product",
    "id": 3903594,
    "signups": 0,
    "savings": 0,
    "savings_in_cents": 0,
    "revenue": 0,
    "revenue_in_cents": 0
  },
  {
    "name": "Product that expires",
    "id": 3853680,
    "signups": 0,
    "savings": 0,
    "savings_in_cents": 0,
    "revenue": 0,
    "revenue_in_cents": 0
  },
  {
    "name": "Trial Product",
    "id": 3861800,
    "signups": 1,
    "savings": 30,
    "savings_in_cents": 3000,
    "revenue": 20,
    "revenue_in_cents": 2000
  }
]
```


# Validate Coupon

You can verify if a specific coupon code is valid using the `validate` method. This method is useful for validating coupon codes that are entered by a customer. If the coupon is found and is valid, the coupon will be returned with a 200 status code.

If the coupon is invalid, the status code will be 404 and the response will say why it is invalid. If the coupon is valid, the status code will be 200 and the coupon will be returned. The following reasons for invalidity are supported:

+ Coupon not found
+ Coupon is invalid
+ Coupon expired

If you have more than one product family and if the coupon you are validating does not belong to the first product family in your site, then you will need to specify the product family, either in the url or as a query string param. This can be done by supplying the id or the handle in the `handle:my-family` format.

Eg.

```
https://<subdomain>.chargify.com/product_families/handle:<product_family_handle>/coupons/validate.<format>?code=<coupon_code>
```

Or:

```
https://<subdomain>.chargify.com/coupons/validate.<format>?code=<coupon_code>&product_family_id=<id>
```

```go
ValidateCoupon(
    ctx context.Context,
    code string,
    productFamilyId *int) (
    models.ApiResponse[models.CouponResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `code` | `string` | Query, Required | The code of the coupon |
| `productFamilyId` | `*int` | Query, Optional | The Advanced Billing id of the product family to which the coupon belongs |

## Response Type

[`models.CouponResponse`](../../doc/models/coupon-response.md)

## Example Usage

```go
ctx := context.Background()

code := "code8"



apiResponse, err := couponsController.ValidateCoupon(ctx, code, nil)
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
  "coupon": {
    "id": 66,
    "name": "Foo Bar",
    "code": "YEPPER9993",
    "description": "my cool coupon",
    "amount_in_cents": 10000,
    "product_family_id": 4,
    "created_at": "2017-11-07T14:51:52-05:00",
    "updated_at": "2017-11-07T15:14:24-05:00",
    "start_date": "2017-11-07T14:51:52-05:00",
    "end_date": null,
    "percentage": null,
    "recurring": false,
    "duration_period_count": null,
    "duration_interval": null,
    "duration_interval_unit": null,
    "allow_negative_balance": false,
    "archived_at": null,
    "conversion_limit": null,
    "stackable": true,
    "compounding_strategy": "full-price"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | [`SingleStringErrorResponseException`](../../doc/models/single-string-error-response-exception.md) |


# Create or Update Coupon Currency Prices

This endpoint allows you to create and/or update currency prices for an existing coupon. Multiple prices can be created or updated in a single request but each of the currencies must be defined on the site level already and the coupon must be an amount-based coupon, not percentage.

Currency pricing for coupons must mirror the setup of the primary coupon pricing - if the primary coupon is percentage based, you will not be able to define pricing in non-primary currencies.

```go
CreateOrUpdateCouponCurrencyPrices(
    ctx context.Context,
    couponId int,
    body *models.CouponCurrencyRequest) (
    models.ApiResponse[models.CouponCurrencyResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `couponId` | `int` | Template, Required | The Advanced Billing id of the coupon |
| `body` | [`*models.CouponCurrencyRequest`](../../doc/models/coupon-currency-request.md) | Body, Optional | - |

## Response Type

[`models.CouponCurrencyResponse`](../../doc/models/coupon-currency-response.md)

## Example Usage

```go
ctx := context.Background()

couponId := 162

body := models.CouponCurrencyRequest{
    CurrencyPrices:       []models.UpdateCouponCurrency{
        models.UpdateCouponCurrency{
            Currency:             "EUR",
            Price:                10,
        },
        models.UpdateCouponCurrency{
            Currency:             "GBP",
            Price:                9,
        },
    },
}

apiResponse, err := couponsController.CreateOrUpdateCouponCurrencyPrices(ctx, couponId, &body)
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
| 422 | Unprocessable Entity (WebDAV) | [`ErrorStringMapResponseException`](../../doc/models/error-string-map-response-exception.md) |


# Create Coupon Subcodes

## Coupon Subcodes Intro

Coupon Subcodes allow you to create a set of unique codes that allow you to expand the use of one coupon.

For example:

Master Coupon Code:

+ SPRING2020

Coupon Subcodes:

+ SPRING90210
+ DP80302
+ SPRINGBALTIMORE

Coupon subcodes can be administered in the Admin Interface or via the API.

When creating a coupon subcode, you must specify a coupon to attach it to using the coupon_id. Valid coupon subcodes are all capital letters, contain only letters and numbers, and do not have any spaces. Lowercase letters will be capitalized before the subcode is created.

## Coupon Subcodes Documentation

Full documentation on how to create coupon subcodes in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24261208729229-Coupon-Codes).

Additionally, for documentation on how to apply a coupon to a Subscription within the Advanced Billing UI, please see our documentation [here](https://maxio.zendesk.com/hc/en-us/articles/24261259337101-Coupons-and-Subscriptions).

## Create Coupon Subcode

This request allows you to create specific subcodes underneath an existing coupon code.

*Note*: If you are using any of the allowed special characters ("%", "@", "+", "-", "_", and "."), you must encode them for use in the URL.

    % to %25
    @ to %40
    + to %2B
    - to %2D
    _ to %5F
    . to %2E

So, if the coupon subcode is `20%OFF`, the URL to delete this coupon subcode would be: `https://<subdomain>.chargify.com/coupons/567/codes/20%25OFF.<format>`

```go
CreateCouponSubcodes(
    ctx context.Context,
    couponId int,
    body *models.CouponSubcodes) (
    models.ApiResponse[models.CouponSubcodesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `couponId` | `int` | Template, Required | The Advanced Billing id of the coupon |
| `body` | [`*models.CouponSubcodes`](../../doc/models/coupon-subcodes.md) | Body, Optional | - |

## Response Type

[`models.CouponSubcodesResponse`](../../doc/models/coupon-subcodes-response.md)

## Example Usage

```go
ctx := context.Background()

couponId := 162

body := models.CouponSubcodes{
    Codes:                []string{
        "BALTIMOREFALL",
        "ORLANDOFALL",
        "DETROITFALL",
    },
}

apiResponse, err := couponsController.CreateCouponSubcodes(ctx, couponId, &body)
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
  "created_codes": [
    "BALTIMOREFALL",
    "ORLANDOFALL",
    "DETROITFALL"
  ]
}
```


# List Coupon Subcodes

This request allows you to request the subcodes that are attached to a coupon.

```go
ListCouponSubcodes(
    ctx context.Context,
    input ListCouponSubcodesInput) (
    models.ApiResponse[models.CouponSubcodes],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `couponId` | `int` | Template, Required | The Advanced Billing id of the coupon |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `20`<br>**Constraints**: `<= 200` |

## Response Type

[`models.CouponSubcodes`](../../doc/models/coupon-subcodes.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListCouponSubcodesInput{
    CouponId: 162,
    Page:     models.ToPointer(2),
    PerPage:  models.ToPointer(50),
}

apiResponse, err := couponsController.ListCouponSubcodes(ctx, collectedInput)
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
  "codes": [
    "3JU6PR",
    "9RO6MP",
    "8OG1VV",
    "5FL7VV",
    "2SV8XK",
    "4LW8LH",
    "3VL4GZ",
    "9UI9XO",
    "0LZ0CC",
    "8XI9JV",
    "9UV5YE",
    "3UI4GX",
    "6SL5ST",
    "9WC8IJ",
    "2KA3PZ",
    "7WR1VR",
    "3VY7MN",
    "6KC3KB",
    "7DF7YT",
    "9FH1ED"
  ]
}
```


# Update Coupon Subcodes

You can update the subcodes for the given Coupon via the API with a PUT request to the resource endpoint.
Send an array of new coupon subcodes.

**Note**: All current subcodes for that Coupon will be deleted first, and replaced with the list of subcodes sent to this endpoint.
The response will contain:

+ The created subcodes,

+ Subcodes that were not created because they already exist,

+ Any subcodes not created because they are invalid.

```go
UpdateCouponSubcodes(
    ctx context.Context,
    couponId int,
    body *models.CouponSubcodes) (
    models.ApiResponse[models.CouponSubcodesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `couponId` | `int` | Template, Required | The Advanced Billing id of the coupon |
| `body` | [`*models.CouponSubcodes`](../../doc/models/coupon-subcodes.md) | Body, Optional | - |

## Response Type

[`models.CouponSubcodesResponse`](../../doc/models/coupon-subcodes-response.md)

## Example Usage

```go
ctx := context.Background()

couponId := 162

body := models.CouponSubcodes{
    Codes:                []string{
        "AAAA",
        "BBBB",
        "CCCC",
    },
}

apiResponse, err := couponsController.UpdateCouponSubcodes(ctx, couponId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Coupon Subcode

## Example

Given a coupon with an ID of 567, and a coupon subcode of 20OFF, the URL to `DELETE` this coupon subcode would be:

```
http://subdomain.chargify.com/coupons/567/codes/20OFF.<format>
```

Note: If you are using any of the allowed special characters (“%”, “@”, “+”, “-”, “_”, and “.”), you must encode them for use in the URL.

| Special character | Encoding |
|-------------------|----------|
| %                 | %25      |
| @                 | %40      |
| +                 | %2B      |
| –                 | %2D      |
| _                 | %5F      |
| .                 | %2E      |

## Percent Encoding Example

Or if the coupon subcode is 20%OFF, the URL to delete this coupon subcode would be: @https://<subdomain>.chargify.com/coupons/567/codes/20%25OFF.<format>

```go
DeleteCouponSubcode(
    ctx context.Context,
    couponId int,
    subcode string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `couponId` | `int` | Template, Required | The Advanced Billing id of the coupon to which the subcode belongs |
| `subcode` | `string` | Template, Required | The subcode of the coupon |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

couponId := 162

subcode := "subcode4"

resp, err := couponsController.DeleteCouponSubcode(ctx, couponId, subcode)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |

