# Events-Based Billing Segments

```go
eventsBasedBillingSegmentsController := client.EventsBasedBillingSegmentsController()
```

## Class Name

`EventsBasedBillingSegmentsController`

## Methods

* [Create Segment](../../doc/controllers/events-based-billing-segments.md#create-segment)
* [List Segments for Price Point](../../doc/controllers/events-based-billing-segments.md#list-segments-for-price-point)
* [Update Segment](../../doc/controllers/events-based-billing-segments.md#update-segment)
* [Delete Segment](../../doc/controllers/events-based-billing-segments.md#delete-segment)
* [Bulk Create Segments](../../doc/controllers/events-based-billing-segments.md#bulk-create-segments)
* [Bulk Update Segments](../../doc/controllers/events-based-billing-segments.md#bulk-update-segments)


# Create Segment

This endpoint creates a new Segment for a Component with segmented Metric. It allows you to specify properties to bill upon and prices for each Segment. You can only pass as many "property_values" as the related Metric has segmenting properties defined.

You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.

```go
CreateSegment(
    ctx context.Context,
    componentId string,
    pricePointId string,
    body *models.CreateSegmentRequest) (
    models.ApiResponse[models.SegmentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `string` | Template, Required | ID or Handle for the Component |
| `pricePointId` | `string` | Template, Required | ID or Handle for the Price Point belonging to the Component |
| `body` | [`*models.CreateSegmentRequest`](../../doc/models/create-segment-request.md) | Body, Optional | - |

## Response Type

[`models.SegmentResponse`](../../doc/models/segment-response.md)

## Example Usage

```go
ctx := context.Background()

componentId := "component_id8"

pricePointId := "price_point_id8"

body := models.CreateSegmentRequest{
    Segment: models.CreateSegment{
        SegmentProperty1Value: models.ToPointer(models.CreateSegmentSegmentProperty1ValueContainer.FromString("France")),
        SegmentProperty2Value: models.ToPointer(models.CreateSegmentSegmentProperty2ValueContainer.FromString("Spain")),
        PricingScheme:         models.PricingScheme("volume"),
        Prices:                []models.CreateOrUpdateSegmentPrice{
            models.CreateOrUpdateSegmentPrice{
                StartingQuantity: models.ToPointer(1),
                EndingQuantity:   models.ToPointer(10000),
                UnitPrice:        models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromPrecision(float64(0.19)),
            },
            models.CreateOrUpdateSegmentPrice{
                StartingQuantity: models.ToPointer(10001),
                UnitPrice:        models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromPrecision(float64(0.09)),
            },
        },
    },
}

apiResponse, err := eventsBasedBillingSegmentsController.CreateSegment(ctx, componentId, pricePointId, &body)
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
| 422 | Unprocessable Entity (WebDAV) | [`EventBasedBillingSegmentErrorsException`](../../doc/models/event-based-billing-segment-errors-exception.md) |


# List Segments for Price Point

This endpoint allows you to fetch Segments created for a given Price Point. They will be returned in the order of creation.

You can pass `page` and `per_page` parameters in order to access all of the segments. By default it will return `30` records. You can set `per_page` to `200` at most.

You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.

```go
ListSegmentsForPricePoint(
    ctx context.Context,
    input ListSegmentsForPricePointInput) (
    models.ApiResponse[models.ListSegmentsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `string` | Template, Required | ID or Handle for the Component |
| `pricePointId` | `string` | Template, Required | ID or Handle for the Price Point belonging to the Component |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br>**Default**: `1`<br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 30. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br>**Default**: `30`<br>**Constraints**: `<= 200` |
| `filter` | [`*models.ListSegmentsFilter`](../../doc/models/list-segments-filter.md) | Query, Optional | Filter to use for List Segments for a Price Point operation |

## Response Type

[`models.ListSegmentsResponse`](../../doc/models/list-segments-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListSegmentsForPricePointInput{
    ComponentId:  "component_id8",
    PricePointId: "price_point_id8",
    Page:         models.ToPointer(2),
    PerPage:      models.ToPointer(50),
    Filter:       models.ToPointer(models.ListSegmentsFilter{
        SegmentProperty1Value: models.ToPointer("EU"),
    }),
}

apiResponse, err := eventsBasedBillingSegmentsController.ListSegmentsForPricePoint(ctx, collectedInput)
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
| 422 | Unprocessable Entity (WebDAV) | [`EventBasedBillingListSegmentsErrorsException`](../../doc/models/event-based-billing-list-segments-errors-exception.md) |


# Update Segment

This endpoint updates a single Segment for a Component with a segmented Metric. It allows you to update the pricing for the segment.

You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.

```go
UpdateSegment(
    ctx context.Context,
    componentId string,
    pricePointId string,
    id float64,
    body *models.UpdateSegmentRequest) (
    models.ApiResponse[models.SegmentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `string` | Template, Required | ID or Handle of the Component |
| `pricePointId` | `string` | Template, Required | ID or Handle of the Price Point belonging to the Component |
| `id` | `float64` | Template, Required | The ID of the Segment |
| `body` | [`*models.UpdateSegmentRequest`](../../doc/models/update-segment-request.md) | Body, Optional | - |

## Response Type

[`models.SegmentResponse`](../../doc/models/segment-response.md)

## Example Usage

```go
ctx := context.Background()

componentId := "component_id8"

pricePointId := "price_point_id8"

id := float64(60)



apiResponse, err := eventsBasedBillingSegmentsController.UpdateSegment(ctx, componentId, pricePointId, id, nil)
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
| 422 | Unprocessable Entity (WebDAV) | [`EventBasedBillingSegmentErrorsException`](../../doc/models/event-based-billing-segment-errors-exception.md) |


# Delete Segment

This endpoint allows you to delete a Segment with specified ID.

You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.

```go
DeleteSegment(
    ctx context.Context,
    componentId string,
    pricePointId string,
    id float64) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `string` | Template, Required | ID or Handle of the Component |
| `pricePointId` | `string` | Template, Required | ID or Handle of the Price Point belonging to the Component |
| `id` | `float64` | Template, Required | The ID of the Segment |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

componentId := "component_id8"

pricePointId := "price_point_id8"

id := float64(60)

resp, err := eventsBasedBillingSegmentsController.DeleteSegment(ctx, componentId, pricePointId, id)
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
| 422 | Unprocessable Entity (WebDAV) | `ApiError` |


# Bulk Create Segments

This endpoint allows you to create multiple segments in one request. The array of segments can contain up to `2000` records.

If any of the records contain an error the whole request would fail and none of the requested segments get created. The error response contains a message for only the one segment that failed validation, with the corresponding index in the array.

You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.

```go
BulkCreateSegments(
    ctx context.Context,
    componentId string,
    pricePointId string,
    body *models.BulkCreateSegments) (
    models.ApiResponse[models.ListSegmentsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `string` | Template, Required | ID or Handle for the Component |
| `pricePointId` | `string` | Template, Required | ID or Handle for the Price Point belonging to the Component |
| `body` | [`*models.BulkCreateSegments`](../../doc/models/bulk-create-segments.md) | Body, Optional | - |

## Response Type

[`models.ListSegmentsResponse`](../../doc/models/list-segments-response.md)

## Example Usage

```go
ctx := context.Background()

componentId := "component_id8"

pricePointId := "price_point_id8"



apiResponse, err := eventsBasedBillingSegmentsController.BulkCreateSegments(ctx, componentId, pricePointId, nil)
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
| 422 | Unprocessable Entity (WebDAV) | [`EventBasedBillingSegmentException`](../../doc/models/event-based-billing-segment-exception.md) |


# Bulk Update Segments

This endpoint allows you to update multiple segments in one request. The array of segments can contain up to `1000` records.

If any of the records contain an error the whole request would fail and none of the requested segments get updated. The error response contains a message for only the one segment that failed validation, with the corresponding index in the array.

You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.

```go
BulkUpdateSegments(
    ctx context.Context,
    componentId string,
    pricePointId string,
    body *models.BulkUpdateSegments) (
    models.ApiResponse[models.ListSegmentsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `componentId` | `string` | Template, Required | ID or Handle for the Component |
| `pricePointId` | `string` | Template, Required | ID or Handle for the Price Point belonging to the Component |
| `body` | [`*models.BulkUpdateSegments`](../../doc/models/bulk-update-segments.md) | Body, Optional | - |

## Response Type

[`models.ListSegmentsResponse`](../../doc/models/list-segments-response.md)

## Example Usage

```go
ctx := context.Background()

componentId := "component_id8"

pricePointId := "price_point_id8"



apiResponse, err := eventsBasedBillingSegmentsController.BulkUpdateSegments(ctx, componentId, pricePointId, nil)
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
| 422 | Unprocessable Entity (WebDAV) | [`EventBasedBillingSegmentException`](../../doc/models/event-based-billing-segment-exception.md) |

