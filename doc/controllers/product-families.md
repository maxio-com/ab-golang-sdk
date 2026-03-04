# Product Families

```go
productFamiliesController := client.ProductFamiliesController()
```

## Class Name

`ProductFamiliesController`

## Methods

* [List Products for Product Family](../../doc/controllers/product-families.md#list-products-for-product-family)
* [Create Product Family](../../doc/controllers/product-families.md#create-product-family)
* [List Product Families](../../doc/controllers/product-families.md#list-product-families)
* [Read Product Family](../../doc/controllers/product-families.md#read-product-family)


# List Products for Product Family

Retrieves a list of Products belonging to a Product Family.

```go
ListProductsForProductFamily(
    ctx context.Context,
    input ListProductsForProductFamilyInput) (
    models.ApiResponse[[]models.ProductResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `input` | [`models.ListProductsForProductFamilyInput`](../../doc/models/list-products-for-product-family-input.md) | Required | Input structure for the method ListProductsForProductFamily |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ProductResponse](../../doc/models/product-response.md).

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListProductsForProductFamilyInput{
    ProductFamilyId: "product_family_id4",
    Page:            models.ToPointer(1),
    PerPage:         models.ToPointer(50),
    DateField:       models.ToPointer(models.BasicDateField_UPDATEDAT),
    Filter:          models.ToPointer(models.ListProductsFilter{
        Ids:                      []int{
            1,
            2,
            3,
        },
    }),
    Include:         models.ToPointer(models.ListProductsInclude_PREPAIDPRODUCTPRICEPOINT),
}

apiResponse, err := productFamiliesController.ListProductsForProductFamily(ctx, collectedInput)
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

Creates a Product Family within your Advanced Billing site. Create a Product Family to act as a container for your products, components and coupons.

Full documentation on how Product Families operate within the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24261098936205-Product-Families).

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
| `body` | [`*models.CreateProductFamilyRequest`](../../doc/models/create-product-family-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ProductFamilyResponse](../../doc/models/product-family-response.md).

## Example Usage

```go
ctx := context.Background()

body := models.CreateProductFamilyRequest{
    ProductFamily:        models.CreateProductFamily{
        Name:                 "Acme Projects",
        Description:          models.NewOptional(models.ToPointer("Amazing project management tool")),
    },
}

apiResponse, err := productFamiliesController.CreateProductFamily(ctx, &body)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Product Families

Retrieve a list of Product Families for a site.

```go
ListProductFamilies(
    ctx context.Context,
    input ListProductFamiliesInput) (
    models.ApiResponse[[]models.ProductFamilyResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `input` | [`models.ListProductFamiliesInput`](../../doc/models/list-product-families-input.md) | Required | Input structure for the method ListProductFamilies |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ProductFamilyResponse](../../doc/models/product-family-response.md).

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListProductFamiliesInput{
    DateField:     models.ToPointer(models.BasicDateField_UPDATEDAT),
}

apiResponse, err := productFamiliesController.ListProductFamilies(ctx, collectedInput)
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
      "updated_at": "2013-02-20T15:05:51-07:00",
      "archived_at": null
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
      "updated_at": "2014-04-16T12:41:13-06:00",
      "archived_at": "2024-11-05T09:30:00-07:00"
    }
  }
]
```


# Read Product Family

Retrieves a Product Family via the `product_family_id`. The response will contain a Product Family object.

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
| `id` | `int` | Template, Required | The Advanced Billing id of the product family |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ProductFamilyResponse](../../doc/models/product-family-response.md).

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
    "accounting_code": null,
    "archived_at": null
  }
}
```

