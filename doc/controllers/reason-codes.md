# Reason Codes

```go
reasonCodesController := client.ReasonCodesController()
```

## Class Name

`ReasonCodesController`

## Methods

* [Create Reason Code](../../doc/controllers/reason-codes.md#create-reason-code)
* [List Reason Codes](../../doc/controllers/reason-codes.md#list-reason-codes)
* [Read Reason Code](../../doc/controllers/reason-codes.md#read-reason-code)
* [Update Reason Code](../../doc/controllers/reason-codes.md#update-reason-code)
* [Delete Reason Code](../../doc/controllers/reason-codes.md#delete-reason-code)


# Create Reason Code

# Reason Codes Intro

ReasonCodes are a way to gain a high level view of why your customers are cancelling the subcription to your product or service.

Add a set of churn reason codes to be displayed in-app and/or the Maxio Billing Portal. As your subscribers decide to cancel their subscription, learn why they decided to cancel.

## Reason Code Documentation

Full documentation on how Reason Codes operate within Advanced Billing can be located under the following links.

[Churn Reason Codes](https://maxio.zendesk.com/hc/en-us/articles/24286647554701-Churn-Reason-Codes)

## Create Reason Code

This method gives a merchant the option to create a reason codes for a given Site.

```go
CreateReasonCode(
    ctx context.Context,
    body *models.CreateReasonCodeRequest) (
    models.ApiResponse[models.ReasonCodeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.CreateReasonCodeRequest`](../../doc/models/create-reason-code-request.md) | Body, Optional | - |

## Response Type

[`models.ReasonCodeResponse`](../../doc/models/reason-code-response.md)

## Example Usage

```go
ctx := context.Background()

body := models.CreateReasonCodeRequest{
    ReasonCode: models.CreateReasonCode{
        Code:        "NOTHANKYOU",
        Description: "No thank you!",
        Position:    models.ToPointer(5),
    },
}

apiResponse, err := reasonCodesController.CreateReasonCode(ctx, &body)
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


# List Reason Codes

This method gives a merchant the option to retrieve a list of all of the current churn codes for a given site.

```go
ListReasonCodes(
    ctx context.Context,
    input ListReasonCodesInput) (
    models.ApiResponse[[]models.ReasonCodeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `20`<br>**Constraints**: `<= 200` |

## Response Type

[`[]models.ReasonCodeResponse`](../../doc/models/reason-code-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListReasonCodesInput{
    Page:    models.ToPointer(2),
    PerPage: models.ToPointer(50),
}

apiResponse, err := reasonCodesController.ListReasonCodes(ctx, collectedInput)
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
    "reason_code": {
      "id": 2,
      "site_id": 2,
      "code": "LARGE",
      "description": "This is too complicated",
      "position": 1,
      "created_at": "2017-02-16T16:49:07-05:00",
      "updated_at": "2017-02-17T16:29:51-05:00"
    }
  },
  {
    "reason_code": {
      "id": 1,
      "site_id": 2,
      "code": "CH1",
      "description": "This doesnt meet my needs",
      "position": 2,
      "created_at": "2017-02-16T16:48:45-05:00",
      "updated_at": "2017-02-17T16:29:59-05:00"
    }
  },
  {
    "reason_code": {
      "id": 5,
      "site_id": 2,
      "code": "HAN99",
      "description": "Hard to setup",
      "position": 3,
      "created_at": "2017-02-17T16:29:42-05:00",
      "updated_at": "2017-02-17T16:29:59-05:00"
    }
  }
]
```


# Read Reason Code

This method gives a merchant the option to retrieve a list of a particular code for a given Site by providing the unique numerical ID of the code.

```go
ReadReasonCode(
    ctx context.Context,
    reasonCodeId int) (
    models.ApiResponse[models.ReasonCodeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `reasonCodeId` | `int` | Template, Required | The Advanced Billing id of the reason code |

## Response Type

[`models.ReasonCodeResponse`](../../doc/models/reason-code-response.md)

## Example Usage

```go
ctx := context.Background()

reasonCodeId := 32

apiResponse, err := reasonCodesController.ReadReasonCode(ctx, reasonCodeId)
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
| 404 | Not Found | `ApiError` |


# Update Reason Code

This method gives a merchant the option to update an existing reason code for a given site.

```go
UpdateReasonCode(
    ctx context.Context,
    reasonCodeId int,
    body *models.UpdateReasonCodeRequest) (
    models.ApiResponse[models.ReasonCodeResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `reasonCodeId` | `int` | Template, Required | The Advanced Billing id of the reason code |
| `body` | [`*models.UpdateReasonCodeRequest`](../../doc/models/update-reason-code-request.md) | Body, Optional | - |

## Response Type

[`models.ReasonCodeResponse`](../../doc/models/reason-code-response.md)

## Example Usage

```go
ctx := context.Background()

reasonCodeId := 32



apiResponse, err := reasonCodesController.UpdateReasonCode(ctx, reasonCodeId, nil)
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
| 404 | Not Found | `ApiError` |


# Delete Reason Code

This method gives a merchant the option to delete one reason code from the Churn Reason Codes. This code will be immediately removed. This action is not reversable.

```go
DeleteReasonCode(
    ctx context.Context,
    reasonCodeId int) (
    models.ApiResponse[models.ReasonCodesJsonResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `reasonCodeId` | `int` | Template, Required | The Advanced Billing id of the reason code |

## Response Type

[`models.ReasonCodesJsonResponse`](../../doc/models/reason-codes-json-response.md)

## Example Usage

```go
ctx := context.Background()

reasonCodeId := 32

apiResponse, err := reasonCodesController.DeleteReasonCode(ctx, reasonCodeId)
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
  "ok": "ok"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |

