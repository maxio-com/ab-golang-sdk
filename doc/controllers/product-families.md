# Product Families

```go
productFamiliesController := client.ProductFamiliesController()
```

## Class Name

`ProductFamiliesController`

## Methods

* [List Products for Product Family](product-families.md#list-products-for-product-family)
* [Create Product Family](product-families.md#create-product-family)
* [List Product Families](product-families.md#list-product-families)
* [Read Product Family](product-families.md#read-product-family)


# List Products for Product Family

This method allows to retrieve a list of Products belonging to a Product Family.

```go
ListProductsForProductFamily(
    ctx context.Context,input ListProductsForProductFamilyInput) (
    models.ApiResponse[[]models.ProductResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the product belongs |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `20`<br>**Constraints**: `<= 200` |
| `dateField` | [`*models.BasicDateFieldEnum`](../models/basic-date-field-enum.md) | Query, Optional | The type of filter you would like to apply to your search.<br>Use in query: `date_field=created_at`. |
| `startDate` | `*string` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `endDate` | `*string` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `startDatetime` | `*string` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `endDatetime` | `*string` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |
| `includeArchived` | `*bool` | Query, Optional | Include archived products |
| `include` | [`*models.ListProductsIncludeEnum`](../models/list-products-include-enum.md) | Query, Optional | Allows including additional data in the response. Use in query `include=prepaid_product_price_point`. |
| `filterPrepaidProductPricePointProductPricePointId` | [`*models.IncludeNotNullEnum`](../models/include-not-null-enum.md) | Query, Optional | Allows fetching products only if a prepaid product price point is present or not. To use this filter you also have to include the following param in the request `include=prepaid_product_price_point`. Use in query `filter[prepaid_product_price_point][product_price_point_id]=not_null`. |
| `filterUseSiteExchangeRate` | `*bool` | Query, Optional | Allows fetching products with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`. |

## Response Type

[`[]models.ProductResponse`](../models/product-response.md)

## Example Usage

```go
ctx := context.Background()
productFamilyId := 140
page := 2
perPage := 50
dateField := models.BasicDateFieldEnum("updated_at")
include := models.ListProductsIncludeEnum("prepaid_product_price_point")Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')

apiResponse, err := productFamiliesController.ListProductsForProductFamily(ctx, productFamilyId, &page, &perPage, &dateField, nil, nil, nil, nil, nil, &include, Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'))
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
    "product": {
      "id": 3801242,
      "name": "Free product",
      "handle": "zero-dollar-product",
      "description": "",
      "accounting_code": "",
      "request_credit_card": true,
      "expiration_interval": null,
      "expiration_interval_unit": "never",
      "created_at": "2016-04-21T16:08:39-04:00",
      "updated_at": "2016-08-03T11:27:53-04:00",
      "price_in_cents": 10000,
      "interval": 1,
      "interval_unit": "month",
      "initial_charge_in_cents": 0,
      "trial_price_in_cents": null,
      "trial_interval": null,
      "trial_interval_unit": "month",
      "archived_at": null,
      "require_credit_card": false,
      "return_params": "",
      "taxable": false,
      "update_return_url": "",
      "initial_charge_after_trial": false,
      "version_number": 4,
      "update_return_params": "",
      "product_family": {
        "id": 527890,
        "name": "Acme Projects",
        "description": "",
        "handle": "billing-plans",
        "accounting_code": null
      },
      "public_signup_pages": [
        {
          "id": 283460,
          "return_url": null,
          "return_params": "",
          "url": "https://general-goods.chargify.com/subscribe/smcc4j3d2w6h/zero-dollar-product"
        }
      ],
      "product_price_point_name": "Default",
      "use_site_exchange_rate": true
    }
  },
  {
    "product": {
      "id": 3858146,
      "name": "Calendar Billing Product",
      "handle": "calendar-billing-product",
      "description": "",
      "accounting_code": "",
      "request_credit_card": true,
      "expiration_interval": null,
      "expiration_interval_unit": "never",
      "created_at": "2016-07-05T13:07:38-04:00",
      "updated_at": "2016-07-05T13:07:38-04:00",
      "price_in_cents": 10000,
      "interval": 1,
      "interval_unit": "month",
      "initial_charge_in_cents": null,
      "trial_price_in_cents": null,
      "trial_interval": null,
      "trial_interval_unit": "month",
      "archived_at": null,
      "require_credit_card": true,
      "return_params": "",
      "taxable": false,
      "update_return_url": "",
      "initial_charge_after_trial": false,
      "version_number": 1,
      "update_return_params": "",
      "product_family": {
        "id": 527890,
        "name": "Acme Projects",
        "description": "",
        "handle": "billing-plans",
        "accounting_code": null
      },
      "public_signup_pages": [
        {
          "id": 289193,
          "return_url": "",
          "return_params": "",
          "url": "https://general-goods.chargify.com/subscribe/gxdbfxzxhcjq/calendar-billing-product"
        }
      ],
      "product_price_point_name": "Default",
      "use_site_exchange_rate": true
    }
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |


# Create Product Family

This method will create a Product Family within your Chargify site. Create a Product Family to act as a container for your products, components and coupons.

Full documentation on how Product Families operate within the Chargify UI can be located [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405369633421).

```go
CreateProductFamily(
    ctx context.Context,
    body *models.CreateProductFamilyRequest) (
    models.ApiResponse[models.ProductFamilyResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.CreateProductFamilyRequest`](../models/create-product-family-request.md) | Body, Optional | - |

## Response Type

[`models.ProductFamilyResponse`](../models/product-family-response.md)

## Example Usage

```go
ctx := context.Background()

bodyProductFamily := models.CreateProductFamily{
    Name:        models.ToPointer("Acme Projects"),
    Description: models.NewOptional(models.ToPointer("Amazing project management tool")),
}

body := models.CreateProductFamilyRequest{
    ProductFamily: bodyProductFamily,
}

apiResponse, err := productFamiliesController.CreateProductFamily(ctx, &body)
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
  "product_family": {
    "id": 933860,
    "name": "Acme Projects",
    "description": "Amazing project management tool",
    "handle": "acme-projects",
    "accounting_code": null
  }
}
```


# List Product Families

This method allows to retrieve a list of Product Families for a site.

```go
ListProductFamilies(
    ctx context.Context,input ListProductFamiliesInput) (
    models.ApiResponse[[]models.ProductFamilyResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `dateField` | [`*models.BasicDateFieldEnum`](../models/basic-date-field-enum.md) | Query, Optional | The type of filter you would like to apply to your search.<br>Use in query: `date_field=created_at`. |
| `startDate` | `*string` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `endDate` | `*string` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `startDatetime` | `*string` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `endDatetime` | `*string` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |

## Response Type

[`[]models.ProductFamilyResponse`](../models/product-family-response.md)

## Example Usage

```go
ctx := context.Background()
dateField := models.BasicDateFieldEnum("updated_at")

apiResponse, err := productFamiliesController.ListProductFamilies(ctx, &dateField, nil, nil, nil, nil)
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
    "product_family": {
      "id": 37,
      "name": "Acme Projects",
      "description": null,
      "handle": "acme-projects",
      "accounting_code": null,
      "created_at": "2013-02-20T15:05:51-07:00",
      "updated_at": "2013-02-20T15:05:51-07:00"
    }
  },
  {
    "product_family": {
      "id": 155,
      "name": "Bat Family",
      "description": "Another family.",
      "handle": "bat-family",
      "accounting_code": null,
      "created_at": "2014-04-16T12:41:13-06:00",
      "updated_at": "2014-04-16T12:41:13-06:00"
    }
  }
]
```


# Read Product Family

This method allows to retrieve a Product Family via the `product_family_id`. The response will contain a Product Family object.

The product family can be specified either with the id number, or with the `handle:my-family` format.

```go
ReadProductFamily(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.ProductFamilyResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `int` | Template, Required | The Chargify id of the product family |

## Response Type

[`models.ProductFamilyResponse`](../models/product-family-response.md)

## Example Usage

```go
ctx := context.Background()
id := 112

apiResponse, err := productFamiliesController.ReadProductFamily(ctx, id)
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
  "product_family": {
    "id": 527890,
    "name": "Acme Projects",
    "description": "",
    "handle": "billing-plans",
    "accounting_code": null
  }
}
```

