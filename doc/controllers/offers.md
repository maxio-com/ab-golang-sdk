# Offers

```go
offersController := client.OffersController()
```

## Class Name

`OffersController`

## Methods

* [Create Offer](../../doc/controllers/offers.md#create-offer)
* [List Offers](../../doc/controllers/offers.md#list-offers)
* [Read Offer](../../doc/controllers/offers.md#read-offer)
* [Archive Offer](../../doc/controllers/offers.md#archive-offer)
* [Unarchive Offer](../../doc/controllers/offers.md#unarchive-offer)


# Create Offer

Create an offer within your Chargify site by sending a POST request.

## Documentation

Offers allow you to package complicated combinations of products, components and coupons into a convenient package which can then be subscribed to just like products.

Once an offer is defined it can be used as an alternative to the product when creating subscriptions.

Full documentation on how to use offers in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407753852059).

## Using a Product Price Point

You can optionally pass in a `product_price_point_id` that corresponds with the `product_id` and the offer will use that price point. If a `product_price_point_id` is not passed in, the product's default price point will be used.

```go
CreateOffer(
    ctx context.Context,
    body *models.CreateOfferRequest) (
    models.ApiResponse[models.OfferResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.CreateOfferRequest`](../../doc/models/create-offer-request.md) | Body, Optional | - |

## Response Type

[`models.OfferResponse`](../../doc/models/offer-response.md)

## Example Usage

```go
ctx := context.Background()

body := models.CreateOfferRequest{
    Offer: models.CreateOffer{
        Name:                "Solo",
        Handle:              "han_shot_first",
        Description:         models.ToPointer("A Star Wars Story"),
        ProductId:           31,
        ProductPricePointId: models.ToPointer(102),
        Components:          []models.CreateOfferComponent{
            models.CreateOfferComponent{
                ComponentId:      models.ToPointer(24),
                StartingQuantity: models.ToPointer(1),
            },
        },
        Coupons:             []string{
            "DEF456",
        },
    },
}

apiResponse, err := offersController.CreateOffer(ctx, &body)
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
  "offer": {
    "id": 3,
    "site_id": 2,
    "product_family_id": 4,
    "product_family_name": "Chargify",
    "product_id": 31,
    "product_name": "30-Day Square Trial",
    "product_price_in_cents": 2000,
    "product_revisable_number": 0,
    "name": "Solo",
    "handle": "han_shot_first",
    "description": "A Star Wars Story",
    "created_at": "2018-06-08T14:51:52-04:00",
    "updated_at": "2018-06-08T14:51:52-04:00",
    "archived_at": null,
    "product_price_point_name": "Default",
    "offer_items": [
      {
        "component_id": 24,
        "component_name": "Invoices",
        "component_unit_price": "3.0",
        "price_point_id": 104,
        "price_point_name": "Original",
        "starting_quantity": "1.0",
        "editable": false
      }
    ],
    "offer_discounts": [
      {
        "coupon_id": 3,
        "coupon_code": "DEF456",
        "coupon_name": "IB Loyalty"
      }
    ]
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorArrayMapResponseException`](../../doc/models/error-array-map-response-exception.md) |


# List Offers

This endpoint will list offers for a site.

```go
ListOffers(
    ctx context.Context,
    input ListOffersInput) (
    models.ApiResponse[models.ListOffersResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `includeArchived` | `*bool` | Query, Optional | Include archived products. Use in query: `include_archived=true`. |

## Response Type

[`models.ListOffersResponse`](../../doc/models/list-offers-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListOffersInput{
    Page:            models.ToPointer(2),
    PerPage:         models.ToPointer(50),
    IncludeArchived: models.ToPointer(true),
}

apiResponse, err := offersController.ListOffers(ctx, collectedInput)
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
  "offers": [
    {
      "id": 239,
      "site_id": 48110,
      "product_family_id": 1025627,
      "product_family_name": "Gold",
      "product_id": 110,
      "product_name": "Pro",
      "product_price_in_cents": 1000,
      "product_revisable_number": 0,
      "product_price_point_id": 138,
      "product_price_point_name": "Default",
      "name": "Third Offer",
      "handle": "third",
      "description": "",
      "created_at": "2018-08-03T09:56:11-05:00",
      "updated_at": "2018-08-03T09:56:11-05:00",
      "archived_at": null,
      "offer_items": [
        {
          "component_id": 426665,
          "component_name": "Database Size (GB)",
          "component_unit_price": "1.0",
          "price_point_id": 149438,
          "price_point_name": "Auto-created",
          "starting_quantity": "0.0",
          "editable": false
        }
      ],
      "offer_discounts": [
        {
          "coupon_id": 234,
          "coupon_code": "GR8_CUSTOMER",
          "coupon_name": "Multi-service Discount"
        }
      ],
      "offer_signup_pages": [
        {
          "id": 356482,
          "nickname": "ggoods",
          "enabled": true,
          "return_url": "",
          "return_params": "",
          "url": "https://general-goods.chargifypay.com/subscribe/hjpvhnw63tzy"
        }
      ]
    }
  ]
}
```


# Read Offer

This method allows you to list a specific offer's attributes. This is different than list all offers for a site, as it requires an `offer_id`.

```go
ReadOffer(
    ctx context.Context,
    offerId int) (
    models.ApiResponse[models.OfferResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `offerId` | `int` | Template, Required | The Chargify id of the offer |

## Response Type

[`models.OfferResponse`](../../doc/models/offer-response.md)

## Example Usage

```go
ctx := context.Background()

offerId := 130

apiResponse, err := offersController.ReadOffer(ctx, offerId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Archive Offer

Archive an existing offer. Please provide an `offer_id` in order to archive the correct item.

```go
ArchiveOffer(
    ctx context.Context,
    offerId int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `offerId` | `int` | Template, Required | The Chargify id of the offer |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

offerId := 130

resp, err := offersController.ArchiveOffer(ctx, offerId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Unarchive Offer

Unarchive a previously archived offer. Please provide an `offer_id` in order to un-archive the correct item.

```go
UnarchiveOffer(
    ctx context.Context,
    offerId int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `offerId` | `int` | Template, Required | The Chargify id of the offer |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

offerId := 130

resp, err := offersController.UnarchiveOffer(ctx, offerId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

