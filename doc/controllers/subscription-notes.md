# Subscription Notes

```go
subscriptionNotesController := client.SubscriptionNotesController()
```

## Class Name

`SubscriptionNotesController`

## Methods

* [Create Subscription Note](../../doc/controllers/subscription-notes.md#create-subscription-note)
* [List Subscription Notes](../../doc/controllers/subscription-notes.md#list-subscription-notes)
* [Read Subscription Note](../../doc/controllers/subscription-notes.md#read-subscription-note)
* [Update Subscription Note](../../doc/controllers/subscription-notes.md#update-subscription-note)
* [Delete Subscription Note](../../doc/controllers/subscription-notes.md#delete-subscription-note)


# Create Subscription Note

Creates a note for a subscription.

Notes allow you to record information about a particular Subscription in a free text format.

If you have structured data such as birth date, color, etc., consider using [Metadata](../../doc/controllers/custom-fields.md#create-metadata) instead.

For more information, see [Adding Notes](https://docs.maxio.com/hc/en-us/articles/24251654953997-Understanding-the-Subscription-Summary-Page#billing-portal-status:~:text=documentation%20for%20more.-,Adding%20Notes,-Notes%20are%20optional) in the product documentation.

```go
CreateSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    body *models.UpdateSubscriptionNoteRequest) (
    models.ApiResponse[models.SubscriptionNoteResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `body` | [`*models.UpdateSubscriptionNoteRequest`](../../doc/models/update-subscription-note-request.md) | Body, Optional | Updatable fields for Subscription Note |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SubscriptionNoteResponse](../../doc/models/subscription-note-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

body := models.UpdateSubscriptionNoteRequest{
    Note:                 models.UpdateSubscriptionNote{
        Body:                 "New test note.",
        Sticky:               true,
    },
}

apiResponse, err := subscriptionNotesController.CreateSubscriptionNote(ctx, subscriptionId, &body)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Subscription Notes

Retrieves a list of notes associated with a subscription. The response will be an array of Notes.

```go
ListSubscriptionNotes(
    ctx context.Context,
    input ListSubscriptionNotesInput) (
    models.ApiResponse[[]models.SubscriptionNoteResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `input` | [`models.ListSubscriptionNotesInput`](../../doc/models/list-subscription-notes-input.md) | Required | Input structure for the method ListSubscriptionNotes |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.SubscriptionNoteResponse](../../doc/models/subscription-note-response.md).

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListSubscriptionNotesInput{
    SubscriptionId: 222,
    Page:           models.ToPointer(1),
    PerPage:        models.ToPointer(50),
}

apiResponse, err := subscriptionNotesController.ListSubscriptionNotes(ctx, collectedInput)
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
[
  {
    "note": {
      "body": "Test note.",
      "created_at": "2015-06-15T13:26:47-04:00",
      "id": 5,
      "sticky": false,
      "subscription_id": 100046,
      "updated_at": "2015-06-15T13:28:12-04:00"
    }
  },
  {
    "note": {
      "body": "Another test note.",
      "created_at": "2015-06-15T12:04:46-04:00",
      "id": 4,
      "sticky": false,
      "subscription_id": 100046,
      "updated_at": "2015-06-15T13:26:33-04:00"
    }
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Read Subscription Note

Retrieves a specific note attached to a subscription.

```go
ReadSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    noteId int) (
    models.ApiResponse[models.SubscriptionNoteResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `noteId` | `int` | Template, Required | The Advanced Billing id of the note |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SubscriptionNoteResponse](../../doc/models/subscription-note-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

noteId := 66

apiResponse, err := subscriptionNotesController.ReadSubscriptionNote(ctx, subscriptionId, noteId)
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
  "note": {
    "body": "Test note.",
    "created_at": "2015-06-15T13:26:47-04:00",
    "id": 5,
    "sticky": false,
    "subscription_id": 100046,
    "updated_at": "2015-06-15T13:28:12-04:00"
  }
}
```


# Update Subscription Note

Updates a note for a subscription.

```go
UpdateSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    noteId int,
    body *models.UpdateSubscriptionNoteRequest) (
    models.ApiResponse[models.SubscriptionNoteResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `noteId` | `int` | Template, Required | The Advanced Billing id of the note |
| `body` | [`*models.UpdateSubscriptionNoteRequest`](../../doc/models/update-subscription-note-request.md) | Body, Optional | Updatable fields for Subscription Note |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SubscriptionNoteResponse](../../doc/models/subscription-note-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

noteId := 66

body := models.UpdateSubscriptionNoteRequest{
    Note:                 models.UpdateSubscriptionNote{
        Body:                 "Modified test note.",
        Sticky:               true,
    },
}

apiResponse, err := subscriptionNotesController.UpdateSubscriptionNote(ctx, subscriptionId, noteId, &body)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Delete Subscription Note

Deletes a note for a Subscription.

```go
DeleteSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    noteId int) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `noteId` | `int` | Template, Required | The Advanced Billing id of the note |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

noteId := 66

resp, err := subscriptionNotesController.DeleteSubscriptionNote(ctx, subscriptionId, noteId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

