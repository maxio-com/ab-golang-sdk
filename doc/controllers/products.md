# Products

```go
productsController := client.ProductsController()
```

## Class Name

`ProductsController`

## Methods

* [Create Product](../../doc/controllers/products.md#create-product)
* [Read Product](../../doc/controllers/products.md#read-product)
* [Update Product](../../doc/controllers/products.md#update-product)
* [Archive Product](../../doc/controllers/products.md#archive-product)
* [Read Product by Handle](../../doc/controllers/products.md#read-product-by-handle)
* [List Products](../../doc/controllers/products.md#list-products)


# Create Product

Use this method to create a product within your Chargify site.

+ [Products Documentation](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405561405709)
+ [Changing a Subscription's Product](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404225334669-Product-Changes-Migrations)

```go
CreateProduct(
    ctx context.Context,
    productFamilyId int,
    body *models.CreateOrUpdateProductRequest) (
    models.ApiResponse[models.ProductResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productFamilyId` | `int` | Template, Required | The Chargify id of the product family to which the product belongs |
| `body` | [`*models.CreateOrUpdateProductRequest`](../../doc/models/create-or-update-product-request.md) | Body, Optional | - |

## Response Type

[`models.ProductResponse`](../../doc/models/product-response.md)

## Example Usage

```go
ctx := context.Background()
productFamilyId := 140

bodyProduct := models.CreateOrUpdateProduct{
    Name:                   "Gold Plan",
    Handle:                 models.ToPointer("gold"),
    Description:            "This is our gold plan.",
    AccountingCode:         models.ToPointer("123"),
    RequireCreditCard:      models.ToPointer(true),
    PriceInCents:           int64(1000),
    Interval:               1,
    IntervalUnit:           models.IntervalUnit("month"),
    AutoCreateSignupPage:   models.ToPointer(true),
    TaxCode:                models.ToPointer("D0000000"),
}

body := models.CreateOrUpdateProductRequest{
    Product: bodyProduct,
}

apiResponse, err := productsController.CreateProduct(ctx, productFamilyId, &body)
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
    "id": 4364984,
    "name": "Gold Plan",
    "handle": "gold",
    "description": "This is our gold plan.",
    "accounting_code": "123",
    "request_credit_card": true,
    "created_at": "2016-11-04T16:31:15-04:00",
    "updated_at": "2016-11-04T16:31:15-04:00",
    "price_in_cents": 1000,
    "interval": 1,
    "interval_unit": "month",
    "expiration_interval_unit": null,
    "initial_charge_in_cents": null,
    "trial_price_in_cents": null,
    "trial_interval": null,
    "trial_interval_unit": null,
    "archived_at": null,
    "require_credit_card": true,
    "return_params": null,
    "taxable": false,
    "update_return_url": null,
    "initial_charge_after_trial": false,
    "version_number": 1,
    "update_return_params": null,
    "product_family": {
      "id": 527890,
      "name": "Acme Projects",
      "description": "",
      "handle": "billing-plans",
      "accounting_code": null
    },
    "public_signup_pages": [
      {
        "id": 301078,
        "return_url": null,
        "return_params": null,
        "url": "https://general-goods.chargify.com/subscribe/ftgbpq7f5qpr/gold"
      }
    ],
    "product_price_point_name": "Default"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Read Product

This endpoint allows you to read the current details of a product that you've created in Chargify.

```go
ReadProduct(
    ctx context.Context,
    productId int) (
    models.ApiResponse[models.ProductResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Chargify id of the product |

## Response Type

[`models.ProductResponse`](../../doc/models/product-response.md)

## Example Usage

```go
ctx := context.Background()
productId := 202

apiResponse, err := productsController.ReadProduct(ctx, productId)
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
    "id": 4535635,
    "name": "Paid Annual Seats",
    "handle": "paid-annual-seats",
    "description": "Paid annual seats for our commercial enterprise product",
    "accounting_code": "paid-annual-seats",
    "request_credit_card": true,
    "expiration_interval": 1,
    "expiration_interval_unit": "day",
    "created_at": "2017-08-25T10:25:31-05:00",
    "updated_at": "2018-01-16T12:58:04-06:00",
    "price_in_cents": 10000,
    "interval": 12,
    "interval_unit": "month",
    "initial_charge_in_cents": 4900,
    "trial_price_in_cents": 1000,
    "trial_interval": 14,
    "trial_interval_unit": "day",
    "archived_at": null,
    "require_credit_card": true,
    "return_params": "id={subscription_id}&ref={customer_reference}",
    "taxable": true,
    "update_return_url": "http://www.example.com",
    "tax_code": "D0000000",
    "initial_charge_after_trial": false,
    "version_number": 4,
    "update_return_params": "id={subscription_id}&ref={customer_reference}",
    "product_family": {
      "id": 1025627,
      "name": "Acme Products",
      "description": "",
      "handle": "acme-products",
      "accounting_code": null
    },
    "public_signup_pages": [],
    "product_price_point_name": "Default"
  }
}
```


# Update Product

Use this method to change aspects of an existing product.

### Input Attributes Update Notes

+ `update_return_params` The parameters we will append to your `update_return_url`. See Return URLs and Parameters

### Product Price Point

Updating a product using this endpoint will create a new price point and set it as the default price point for this product. If you should like to update an existing product price point, that must be done separately.

```go
UpdateProduct(
    ctx context.Context,
    productId int,
    body *models.CreateOrUpdateProductRequest) (
    models.ApiResponse[models.ProductResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Chargify id of the product |
| `body` | [`*models.CreateOrUpdateProductRequest`](../../doc/models/create-or-update-product-request.md) | Body, Optional | - |

## Response Type

[`models.ProductResponse`](../../doc/models/product-response.md)

## Example Usage

```go
ctx := context.Background()
productId := 202

apiResponse, err := productsController.UpdateProduct(ctx, productId, nil)
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
    "id": 4365034,
    "name": "Platinum Plan",
    "handle": "platinum",
    "description": "This is our platinum plan.",
    "accounting_code": "123",
    "request_credit_card": true,
    "created_at": "2016-11-04T16:34:29-04:00",
    "updated_at": "2016-11-04T16:37:11-04:00",
    "price_in_cents": 1000,
    "interval": 1,
    "interval_unit": "month",
    "initial_charge_in_cents": null,
    "trial_price_in_cents": null,
    "trial_interval": null,
    "trial_interval_unit": null,
    "archived_at": null,
    "require_credit_card": true,
    "return_params": null,
    "taxable": false,
    "update_return_url": null,
    "initial_charge_after_trial": false,
    "version_number": 1,
    "update_return_params": null,
    "product_family": {
      "id": 527890,
      "name": "Acme Projects",
      "description": "",
      "handle": "billing-plans",
      "accounting_code": null
    },
    "public_signup_pages": [
      {
        "id": 301079,
        "return_url": null,
        "return_params": null,
        "url": "https://general-goods.chargify.com/subscribe/wgyd96tb5pj9/platinum"
      }
    ],
    "product_price_point_name": "Original"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Archive Product

Sending a DELETE request to this endpoint will archive the product. All current subscribers will be unffected; their subscription/purchase will continue to be charged monthly.

This will restrict the option to chose the product for purchase via the Billing Portal, as well as disable Public Signup Pages for the product.

```go
ArchiveProduct(
    ctx context.Context,
    productId int) (
    models.ApiResponse[models.ProductResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `productId` | `int` | Template, Required | The Chargify id of the product |

## Response Type

[`models.ProductResponse`](../../doc/models/product-response.md)

## Example Usage

```go
ctx := context.Background()
productId := 202

apiResponse, err := productsController.ArchiveProduct(ctx, productId)
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
    "id": 4535638,
    "name": "Business Monthly",
    "handle": null,
    "description": "Business Monthly",
    "accounting_code": "",
    "request_credit_card": true,
    "expiration_interval": null,
    "expiration_interval_unit": "never",
    "created_at": "2017-08-25T10:25:31-05:00",
    "updated_at": "2018-01-16T13:02:44-06:00",
    "price_in_cents": 4900,
    "interval": 1,
    "interval_unit": "month",
    "initial_charge_in_cents": null,
    "trial_price_in_cents": 0,
    "trial_interval": 1,
    "trial_interval_unit": "day",
    "archived_at": "2018-01-16T13:02:44-06:00",
    "require_credit_card": false,
    "return_params": "",
    "taxable": false,
    "update_return_url": "",
    "tax_code": "",
    "initial_charge_after_trial": false,
    "version_number": 1,
    "update_return_params": "",
    "product_family": {
      "id": 1025627,
      "name": "Acme Products",
      "description": "",
      "handle": "acme-products",
      "accounting_code": null
    },
    "public_signup_pages": [],
    "product_price_point_name": "Default"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Read Product by Handle

This method allows to retrieve a Product object by its `api_handle`.

```go
ReadProductByHandle(
    ctx context.Context,
    apiHandle string) (
    models.ApiResponse[models.ProductResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `apiHandle` | `string` | Template, Required | The handle of the product |

## Response Type

[`models.ProductResponse`](../../doc/models/product-response.md)

## Example Usage

```go
ctx := context.Background()
apiHandle := "api_handle6"

apiResponse, err := productsController.ReadProductByHandle(ctx, apiHandle)
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
    "id": 3903594,
    "name": "No cost product",
    "handle": "no-cost-product",
    "description": "",
    "accounting_code": "",
    "request_credit_card": true,
    "expiration_interval": null,
    "expiration_interval_unit": "never",
    "created_at": "2016-09-02T17:11:29-04:00",
    "updated_at": "2016-11-30T11:46:13-05:00",
    "price_in_cents": 0,
    "interval": 1,
    "interval_unit": "month",
    "initial_charge_in_cents": null,
    "trial_price_in_cents": 5,
    "trial_interval": 1,
    "trial_interval_unit": "month",
    "archived_at": null,
    "require_credit_card": false,
    "return_params": "reference=5678",
    "taxable": false,
    "update_return_url": "",
    "initial_charge_after_trial": false,
    "version_number": 1,
    "update_return_params": "reference=5678",
    "product_family": {
      "id": 527890,
      "name": "Acme Projects",
      "description": "",
      "handle": "billing-plans",
      "accounting_code": null
    },
    "public_signup_pages": [
      {
        "id": 281174,
        "return_url": "",
        "return_params": "",
        "url": "https://general-goods.chargify.com/subscribe/xgdxtk4vhtbz/no-cost-product"
      },
      {
        "id": 282270,
        "return_url": "",
        "return_params": "",
        "url": "https://general-goods.chargify.com/subscribe/xxqmrgtsbd9k/no-cost-product"
      },
      {
        "id": 291587,
        "return_url": "",
        "return_params": "",
        "url": "https://general-goods.chargify.com/subscribe/pvhwss7zjjnh/no-cost-product"
      },
      {
        "id": 294832,
        "return_url": "http://www.example.com/",
        "return_params": "engine=md7a",
        "url": "https://general-goods.chargify.com/subscribe/m6tbcq4mcgpw/no-cost-product"
      }
    ],
    "product_price_point_name": "Default"
  }
}
```


# List Products

This method allows to retrieve a list of Products belonging to a Site.

```go
ListProducts(
    ctx context.Context,input ListProductsInput) (
    models.ApiResponse[[]models.ProductResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `dateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Query, Optional | The type of filter you would like to apply to your search.<br>Use in query: `date_field=created_at`. |
| `endDate` | `*time.Time` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `endDatetime` | `*time.Time` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of end_date. |
| `startDate` | `*time.Time` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns products with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `startDatetime` | `*time.Time` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns products with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of start_date. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `includeArchived` | `*bool` | Query, Optional | Include archived products. Use in query: `include_archived=true`. |
| `include` | [`*models.ListProductsInclude`](../../doc/models/list-products-include.md) | Query, Optional | Allows including additional data in the response. Use in query `include=prepaid_product_price_point`. |
| `filterPrepaidProductPricePointProductPricePointId` | [`*models.IncludeNotNull`](../../doc/models/include-not-null.md) | Query, Optional | Allows fetching products only if a prepaid product price point is present or not. To use this filter you also have to include the following param in the request `include=prepaid_product_price_point`. Use in query `filter[prepaid_product_price_point][product_price_point_id]=not_null`. |
| `filterUseSiteExchangeRate` | `*bool` | Query, Optional | Allows fetching products with matching use_site_exchange_rate based on provided value (refers to default price point). Use in query `filter[use_site_exchange_rate]=true`. |

## Response Type

[`[]models.ProductResponse`](../../doc/models/product-response.md)

## Example Usage

```go
ctx := context.Background()
dateField := models.BasicDateField("updated_at")
page := 2
perPage := 50
includeArchived := true
include := models.ListProductsInclude("prepaid_product_price_point")Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')

apiResponse, err := productsController.ListProducts(ctx, &dateField, nil, nil, nil, nil, &page, &perPage, &includeArchived, &include, Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'))
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
      "id": 0,
      "name": "string",
      "handle": "string",
      "description": "string",
      "accounting_code": "string",
      "request_credit_card": true,
      "expiration_interval": 0,
      "expiration_interval_unit": "month",
      "created_at": "2023-11-23T10:28:34-05:00",
      "updated_at": "2023-11-23T10:28:34-05:00",
      "price_in_cents": 0,
      "interval": 0,
      "interval_unit": "month",
      "initial_charge_in_cents": 0,
      "trial_price_in_cents": 0,
      "trial_interval": 0,
      "trial_interval_unit": "month",
      "archived_at": null,
      "require_credit_card": true,
      "return_params": "string",
      "taxable": true,
      "update_return_url": "string",
      "initial_charge_after_trial": true,
      "version_number": 0,
      "update_return_params": "string",
      "product_family": {
        "id": 0,
        "name": "string",
        "handle": "string",
        "accounting_code": null,
        "description": "string",
        "created_at": "2021-05-05T16:00:21-04:00",
        "updated_at": "2021-05-05T16:00:21-04:00"
      },
      "public_signup_pages": [
        {
          "id": 0,
          "return_url": "string",
          "return_params": "string",
          "url": "string"
        }
      ],
      "product_price_point_name": "string",
      "request_billing_address": true,
      "require_billing_address": true,
      "require_shipping_address": true,
      "use_site_exchange_rate": true,
      "tax_code": "string",
      "default_product_price_point_id": 0
    }
  }
]
```

