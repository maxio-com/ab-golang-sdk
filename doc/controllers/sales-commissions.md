# Sales Commissions

```go
salesCommissionsController := client.SalesCommissionsController()
```

## Class Name

`SalesCommissionsController`

## Methods

* [List Sales Commission Settings](../../doc/controllers/sales-commissions.md#list-sales-commission-settings)
* [List Sales Reps](../../doc/controllers/sales-commissions.md#list-sales-reps)
* [Read Sales Rep](../../doc/controllers/sales-commissions.md#read-sales-rep)


# List Sales Commission Settings

Endpoint returns subscriptions with associated sales reps

## Modified Authentication Process

The Sales Commission API differs from other Chargify API endpoints. This resource is associated with the seller itself. Up to now all available resources were at the level of the site, therefore creating the API Key per site was a sufficient solution. To share resources at the seller level, a new authentication method was introduced, which is user authentication. Creating an API Key for a user is a required step to correctly use the Sales Commission API, more details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).

Access to the Sales Commission API endpoints is available to users with financial access, where the seller has the Advanced Analytics component enabled. For further information on getting access to Advanced Analytics please contact Maxio support.

> Note: The request is at seller level, it means `<<subdomain>>` variable will be replaced by `app`

```go
ListSalesCommissionSettings(
    ctx context.Context,
    input ListSalesCommissionSettingsInput) (
    models.ApiResponse[[]models.SaleRepSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `sellerId` | `string` | Template, Required | The Chargify id of your seller account |
| `authorization` | `*string` | Header, Optional | For authorization use user API key. See details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).<br>**Default**: `"Bearer <<apiKey>>"` |
| `liveMode` | `*bool` | Query, Optional | This parameter indicates if records should be fetched from live mode sites. Default value is true. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 100.<br>**Default**: `100` |

## Response Type

[`[]models.SaleRepSettings`](../../doc/models/sale-rep-settings.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListSalesCommissionSettingsInput{
    SellerId:      "seller_id8",
    Authorization: models.ToPointer("Bearer <<apiKey>>"),
    Page:          models.ToPointer(2),
    PerPage:       models.ToPointer(100),
}

apiResponse, err := salesCommissionsController.ListSalesCommissionSettings(ctx, collectedInput)
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
    "customer_name": "Ziomek Ziomeczek",
    "subscription_id": 81746,
    "site_link": "https://chargify9.staging-chargify.com/dashboard",
    "site_name": "Chargify",
    "subscription_mrr": "$200.00",
    "sales_rep_id": 48,
    "sales_rep_name": "John Candy"
  },
  {
    "customer_name": "Ziom Kom",
    "subscription_id": 83758,
    "site_link": "https://chargify9.staging-chargify.com/dashboard",
    "site_name": "Chargify",
    "subscription_mrr": "$200.00",
    "sales_rep_id": 49,
    "sales_rep_name": "Josh Acme"
  },
  {
    "customer_name": "George Bush",
    "subscription_id": 83790,
    "site_link": "https://chargify9.staging-chargify.com/dashboard",
    "site_name": "Chargify",
    "subscription_mrr": "$200.00",
    "sales_rep_id": 48,
    "sales_rep_name": "John Candy"
  }
]
```


# List Sales Reps

Endpoint returns sales rep list with details

## Modified Authentication Process

The Sales Commission API differs from other Chargify API endpoints. This resource is associated with the seller itself. Up to now all available resources were at the level of the site, therefore creating the API Key per site was a sufficient solution. To share resources at the seller level, a new authentication method was introduced, which is user authentication. Creating an API Key for a user is a required step to correctly use the Sales Commission API, more details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).

Access to the Sales Commission API endpoints is available to users with financial access, where the seller has the Advanced Analytics component enabled. For further information on getting access to Advanced Analytics please contact Maxio support.

> Note: The request is at seller level, it means `<<subdomain>>` variable will be replaced by `app`

```go
ListSalesReps(
    ctx context.Context,
    input ListSalesRepsInput) (
    models.ApiResponse[[]models.ListSaleRepItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `sellerId` | `string` | Template, Required | The Chargify id of your seller account |
| `authorization` | `*string` | Header, Optional | For authorization use user API key. See details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).<br>**Default**: `"Bearer <<apiKey>>"` |
| `liveMode` | `*bool` | Query, Optional | This parameter indicates if records should be fetched from live mode sites. Default value is true. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 100.<br>**Default**: `100` |

## Response Type

[`[]models.ListSaleRepItem`](../../doc/models/list-sale-rep-item.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListSalesRepsInput{
    SellerId:      "seller_id8",
    Authorization: models.ToPointer("Bearer <<apiKey>>"),
    Page:          models.ToPointer(2),
    PerPage:       models.ToPointer(100),
}

apiResponse, err := salesCommissionsController.ListSalesReps(ctx, collectedInput)
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
    "id": 48,
    "full_name": "John Candy",
    "subscriptions_count": 2,
    "mrr_data": {
      "november_2019": {
        "mrr": "$0.00",
        "usage": "$0.00",
        "recurring": "$0.00"
      },
      "december_2019": {
        "mrr": "$0.00",
        "usage": "$0.00",
        "recurring": "$0.00"
      },
      "january_2020": {
        "mrr": "$400.00",
        "usage": "$0.00",
        "recurring": "$400.00"
      },
      "february_2020": {
        "mrr": "$400.00",
        "usage": "$0.00",
        "recurring": "$400.00"
      },
      "march_2020": {
        "mrr": "$400.00",
        "usage": "$0.00",
        "recurring": "$400.00"
      },
      "april_2020": {
        "mrr": "$400.00",
        "usage": "$0.00",
        "recurring": "$400.00"
      }
    },
    "test_mode": true
  },
  {
    "id": 49,
    "full_name": "Josh Acme",
    "subscriptions_count": 1,
    "mrr_data": {
      "november_2019": {
        "mrr": "$0.00",
        "usage": "$0.00",
        "recurring": "$0.00"
      },
      "december_2019": {
        "mrr": "$0.00",
        "usage": "$0.00",
        "recurring": "$0.00"
      },
      "january_2020": {
        "mrr": "$200.00",
        "usage": "$0.00",
        "recurring": "$200.00"
      },
      "february_2020": {
        "mrr": "$200.00",
        "usage": "$0.00",
        "recurring": "$200.00"
      },
      "march_2020": {
        "mrr": "$200.00",
        "usage": "$0.00",
        "recurring": "$200.00"
      },
      "april_2020": {
        "mrr": "$200.00",
        "usage": "$0.00",
        "recurring": "$200.00"
      }
    },
    "test_mode": true
  }
]
```


# Read Sales Rep

Endpoint returns sales rep and attached subscriptions details.

## Modified Authentication Process

The Sales Commission API differs from other Chargify API endpoints. This resource is associated with the seller itself. Up to now all available resources were at the level of the site, therefore creating the API Key per site was a sufficient solution. To share resources at the seller level, a new authentication method was introduced, which is user authentication. Creating an API Key for a user is a required step to correctly use the Sales Commission API, more details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).

Access to the Sales Commission API endpoints is available to users with financial access, where the seller has the Advanced Analytics component enabled. For further information on getting access to Advanced Analytics please contact Maxio support.

> Note: The request is at seller level, it means `<<subdomain>>` variable will be replaced by `app`

```go
ReadSalesRep(
    ctx context.Context,
    sellerId string,
    salesRepId string,
    authorization *string,
    liveMode *bool,
    page *int,
    perPage *int) (
    models.ApiResponse[models.SaleRep],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `sellerId` | `string` | Template, Required | The Chargify id of your seller account |
| `salesRepId` | `string` | Template, Required | The Advanced Billing id of sales rep. |
| `authorization` | `*string` | Header, Optional | For authorization use user API key. See details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).<br>**Default**: `"Bearer <<apiKey>>"` |
| `liveMode` | `*bool` | Query, Optional | This parameter indicates if records should be fetched from live mode sites. Default value is true. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 100.<br>**Default**: `100` |

## Response Type

[`models.SaleRep`](../../doc/models/sale-rep.md)

## Example Usage

```go
ctx := context.Background()

sellerId := "seller_id8"

salesRepId := "sales_rep_id4"

authorization := "Bearer <<apiKey>>"



page := 2

perPage := 100

apiResponse, err := salesCommissionsController.ReadSalesRep(ctx, sellerId, salesRepId, &authorization, nil, &page, &perPage)
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
  "id": 48,
  "full_name": "John Candy",
  "subscriptions_count": 2,
  "test_mode": true,
  "subscriptions": [
    {
      "id": 81746,
      "site_name": "Chargify",
      "subscription_url": "https://chargify9.staging-chargify.com/subscriptions/81746",
      "customer_name": "Ziomek Ziomeczek",
      "created_at": "2020-01-03T02:36:27-05:00",
      "mrr": "$200.00",
      "usage": "$0.00",
      "recurring": "$200.00",
      "last_payment": "2020-04-03T03:40:27-04:00",
      "churn_date": null
    },
    {
      "id": 83790,
      "site_name": "Chargify",
      "subscription_url": "https://chargify9.staging-chargify.com/subscriptions/83790",
      "customer_name": "George Bush",
      "created_at": "2020-01-17T07:34:32-05:00",
      "mrr": "$200.00",
      "usage": "$0.00",
      "recurring": "$200.00",
      "last_payment": "2020-04-17T08:41:03-04:00",
      "churn_date": null
    }
  ]
}
```

