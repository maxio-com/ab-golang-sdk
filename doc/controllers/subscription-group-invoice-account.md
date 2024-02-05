# Subscription Group Invoice Account

```go
subscriptionGroupInvoiceAccountController := client.SubscriptionGroupInvoiceAccountController()
```

## Class Name

`SubscriptionGroupInvoiceAccountController`

## Methods

* [Create Subscription Group Prepayment](../../doc/controllers/subscription-group-invoice-account.md#create-subscription-group-prepayment)
* [List Prepayments for Subscription Group](../../doc/controllers/subscription-group-invoice-account.md#list-prepayments-for-subscription-group)
* [Issue Subscription Group Service Credit](../../doc/controllers/subscription-group-invoice-account.md#issue-subscription-group-service-credit)
* [Deduct Subscription Group Service Credit](../../doc/controllers/subscription-group-invoice-account.md#deduct-subscription-group-service-credit)


# Create Subscription Group Prepayment

A prepayment can be added for a subscription group identified by the group's `uid`. This endpoint requires a `amount`, `details`, `method`, and `memo`. On success, the prepayment will be added to the group's prepayment balance.

```go
CreateSubscriptionGroupPrepayment(
    ctx context.Context,
    uid string,
    body *models.SubscriptionGroupPrepaymentRequest) (
    models.ApiResponse[models.SubscriptionGroupPrepaymentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |
| `body` | [`*models.SubscriptionGroupPrepaymentRequest`](../../doc/models/subscription-group-prepayment-request.md) | Body, Optional | - |

## Response Type

[`models.SubscriptionGroupPrepaymentResponse`](../../doc/models/subscription-group-prepayment-response.md)

## Example Usage

```go
ctx := context.Background()
uid := "uid0"

apiResponse, err := subscriptionGroupInvoiceAccountController.CreateSubscriptionGroupPrepayment(ctx, uid, nil)
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
  "id": 6049554,
  "amount_in_cents": 10000,
  "ending_balance_in_cents": 5000,
  "entry_type": "Debit",
  "memo": "Debit from invoice account."
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Prepayments for Subscription Group

This request will list a subscription group's prepayments.

```go
ListPrepaymentsForSubscriptionGroup(
    ctx context.Context,input ListPrepaymentsForSubscriptionGroupInput) (
    models.ApiResponse[models.ListSubscriptionGroupPrepaymentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |
| `filterDateField` | [`*models.ListSubscriptionGroupPrepaymentDateField`](../../doc/models/list-subscription-group-prepayment-date-field.md) | Query, Optional | The type of filter you would like to apply to your search.<br>Use in query: `filter[date_field]=created_at`. |
| `filterEndDate` | `*time.Time` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field.<br>Returns prepayments with a timestamp up to and including 11:59:59PM in your site's time zone on the date specified.<br>Use in query: `filter[end_date]=2011-12-15`. |
| `filterStartDate` | `*time.Time` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field.<br>Returns prepayments with a timestamp at or after midnight (12:00:00 AM) in your site's time zone on the date specified.<br>Use in query: `filter[start_date]=2011-12-15`. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |

## Response Type

[`models.ListSubscriptionGroupPrepaymentResponse`](../../doc/models/list-subscription-group-prepayment-response.md)

## Example Usage

```go
ctx := context.Background()
uid := "uid0"Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')Liquid error: Value cannot be null. (Parameter 'key')
page := 2
perPage := 50

apiResponse, err := subscriptionGroupInvoiceAccountController.ListPrepaymentsForSubscriptionGroup(ctx, uid, Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), Liquid error: Value cannot be null. (Parameter 'key'), &page, &perPage)
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
  "prepayments": [
    {
      "prepayment": {
        "id": 142,
        "subscription_group_uid": "grp_b4qhx3bvx72t8",
        "amount_in_cents": 10000,
        "remaining_amount_in_cents": 10000,
        "details": "test",
        "external": true,
        "memo": "test",
        "payment_type": "cash",
        "created_at": "2023-06-21T04:37:02-04:00"
      }
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |


# Issue Subscription Group Service Credit

Credit can be issued for a subscription group identified by the group's `uid`. Credit will be added to the group in the amount specified in the request body. The credit will be applied to group member invoices as they are generated.

```go
IssueSubscriptionGroupServiceCredit(
    ctx context.Context,
    uid string,
    body *models.IssueServiceCreditRequest) (
    models.ApiResponse[models.ServiceCreditResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |
| `body` | [`*models.IssueServiceCreditRequest`](../../doc/models/issue-service-credit-request.md) | Body, Optional | - |

## Response Type

[`models.ServiceCreditResponse`](../../doc/models/service-credit-response.md)

## Example Usage

```go
ctx := context.Background()
uid := "uid0"

bodyServiceCredit := models.IssueServiceCredit{
    Amount: interface{}("[key1, val1][key2, val2]"),
    Memo:   "Credit the group account",
}

body := models.IssueServiceCreditRequest{
    ServiceCredit: bodyServiceCredit,
}

apiResponse, err := subscriptionGroupInvoiceAccountController.IssueSubscriptionGroupServiceCredit(ctx, uid, &body)
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
  "service_credit": {
    "id": 101,
    "amount_in_cents": 1000,
    "ending_balance_in_cents": 2000,
    "entry_type": "Credit",
    "memo": "Credit to group account"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Deduct Subscription Group Service Credit

Credit can be deducted for a subscription group identified by the group's `uid`. Credit will be deducted from the group in the amount specified in the request body.

```go
DeductSubscriptionGroupServiceCredit(
    ctx context.Context,
    uid string,
    body *models.DeductServiceCreditRequest) (
    models.ApiResponse[models.ServiceCredit],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |
| `body` | [`*models.DeductServiceCreditRequest`](../../doc/models/deduct-service-credit-request.md) | Body, Optional | - |

## Response Type

[`models.ServiceCredit`](../../doc/models/service-credit.md)

## Example Usage

```go
ctx := context.Background()
uid := "uid0"

bodyDeduction := models.DeductServiceCredit{
    Amount: interface{}("[key1, val1][key2, val2]"),
    Memo:   "Deduct from group account",
}

body := models.DeductServiceCreditRequest{
    Deduction: bodyDeduction,
}

apiResponse, err := subscriptionGroupInvoiceAccountController.DeductSubscriptionGroupServiceCredit(ctx, uid, &body)
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
  "id": 100,
  "amount_in_cents": 1000,
  "ending_balance_in_cents": 0,
  "entry_type": "Debit",
  "memo": "Debit from group account"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

