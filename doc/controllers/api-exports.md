# API Exports

```go
APIExportsController := client.APIExportsController()
```

## Class Name

`APIExportsController`

## Methods

* [List Exported Proforma Invoices](../../doc/controllers/api-exports.md#list-exported-proforma-invoices)
* [List Exported Invoices](../../doc/controllers/api-exports.md#list-exported-invoices)
* [List Exported Subscriptions](../../doc/controllers/api-exports.md#list-exported-subscriptions)
* [Export Proforma Invoices](../../doc/controllers/api-exports.md#export-proforma-invoices)
* [Export Invoices](../../doc/controllers/api-exports.md#export-invoices)
* [Export Subscriptions](../../doc/controllers/api-exports.md#export-subscriptions)
* [Read Proforma Invoices Export](../../doc/controllers/api-exports.md#read-proforma-invoices-export)
* [Read Invoices Export](../../doc/controllers/api-exports.md#read-invoices-export)
* [Read Subscriptions Export](../../doc/controllers/api-exports.md#read-subscriptions-export)


# List Exported Proforma Invoices

This API returns an array of exported proforma invoices for a provided `batch_id`. Pay close attention to pagination in order to control responses from the server.

Example: `GET https://{subdomain}.chargify.com/api_exports/proforma_invoices/123/rows?per_page=10000&page=1`.

```go
ListExportedProformaInvoices(
    ctx context.Context,input ListExportedProformaInvoicesInput) (
    models.ApiResponse[[]models.ProformaInvoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `batchId` | `string` | Template, Required | Id of a Batch Job. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request.<br>Default value is 100.<br>The maximum allowed values is 10000; any per_page value over 10000 will be changed to 10000. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |

## Response Type

[`[]models.ProformaInvoice`](../../doc/models/proforma-invoice.md)

## Example Usage

```go
ctx := context.Background()
batchId := "batch_id8"
perPage := 100
page := 2

apiResponse, err := APIExportsController.ListExportedProformaInvoices(ctx, batchId, &perPage, &page)
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


# List Exported Invoices

This API returns an array of exported invoices for a provided `batch_id`. Pay close attention to pagination in order to control responses from the server.

Example: `GET https://{subdomain}.chargify.com/api_exports/invoices/123/rows?per_page=10000&page=1`.

```go
ListExportedInvoices(
    ctx context.Context,input ListExportedInvoicesInput) (
    models.ApiResponse[[]models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `batchId` | `string` | Template, Required | Id of a Batch Job. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request.<br>Default value is 100.<br>The maximum allowed values is 10000; any per_page value over 10000 will be changed to 10000. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |

## Response Type

[`[]models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()
batchId := "batch_id8"
perPage := 100
page := 2

apiResponse, err := APIExportsController.ListExportedInvoices(ctx, batchId, &perPage, &page)
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


# List Exported Subscriptions

This API returns an array of exported subscriptions for a provided `batch_id`. Pay close attention to pagination in order to control responses from the server.

Example: `GET https://{subdomain}.chargify.com/api_exports/subscriptions/123/rows?per_page=200&page=1`.

```go
ListExportedSubscriptions(
    ctx context.Context,input ListExportedSubscriptionsInput) (
    models.ApiResponse[[]models.Subscription],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `batchId` | `string` | Template, Required | Id of a Batch Job. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request.<br>Default value is 100.<br>The maximum allowed values is 10000; any per_page value over 10000 will be changed to 10000. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |

## Response Type

[`[]models.Subscription`](../../doc/models/subscription.md)

## Example Usage

```go
ctx := context.Background()
batchId := "batch_id8"
perPage := 100
page := 2

apiResponse, err := APIExportsController.ListExportedSubscriptions(ctx, batchId, &perPage, &page)
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


# Export Proforma Invoices

This API creates a proforma invoices export and returns a batchjob object.

It is only available for Relationship Invoicing architecture.

```go
ExportProformaInvoices(
    ctx context.Context) (
    models.ApiResponse[models.BatchJobResponse],
    error)
```

## Response Type

[`models.BatchJobResponse`](../../doc/models/batch-job-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := APIExportsController.ExportProformaInvoices(ctx)
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
| 409 | Conflict | [`SingleErrorResponseException`](../../doc/models/single-error-response-exception.md) |


# Export Invoices

This API creates an invoices export and returns a batchjob object.

```go
ExportInvoices(
    ctx context.Context) (
    models.ApiResponse[models.BatchJobResponse],
    error)
```

## Response Type

[`models.BatchJobResponse`](../../doc/models/batch-job-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := APIExportsController.ExportInvoices(ctx)
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
| 409 | Conflict | [`SingleErrorResponseException`](../../doc/models/single-error-response-exception.md) |


# Export Subscriptions

This API creates a subscriptions export and returns a batchjob object.

```go
ExportSubscriptions(
    ctx context.Context) (
    models.ApiResponse[models.BatchJobResponse],
    error)
```

## Response Type

[`models.BatchJobResponse`](../../doc/models/batch-job-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := APIExportsController.ExportSubscriptions(ctx)
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
| 409 | Conflict | [`SingleErrorResponseException`](../../doc/models/single-error-response-exception.md) |


# Read Proforma Invoices Export

This API returns a batchjob object for proforma invoices export.

```go
ReadProformaInvoicesExport(
    ctx context.Context,
    batchId string) (
    models.ApiResponse[models.BatchJobResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `batchId` | `string` | Template, Required | Id of a Batch Job. |

## Response Type

[`models.BatchJobResponse`](../../doc/models/batch-job-response.md)

## Example Usage

```go
ctx := context.Background()
batchId := "batch_id8"

apiResponse, err := APIExportsController.ReadProformaInvoicesExport(ctx, batchId)
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


# Read Invoices Export

This API returns a batchjob object for invoices export.

```go
ReadInvoicesExport(
    ctx context.Context,
    batchId string) (
    models.ApiResponse[models.BatchJobResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `batchId` | `string` | Template, Required | Id of a Batch Job. |

## Response Type

[`models.BatchJobResponse`](../../doc/models/batch-job-response.md)

## Example Usage

```go
ctx := context.Background()
batchId := "batch_id8"

apiResponse, err := APIExportsController.ReadInvoicesExport(ctx, batchId)
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


# Read Subscriptions Export

This API returns a batchjob object for subscriptions export.

```go
ReadSubscriptionsExport(
    ctx context.Context,
    batchId string) (
    models.ApiResponse[models.BatchJobResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `batchId` | `string` | Template, Required | Id of a Batch Job. |

## Response Type

[`models.BatchJobResponse`](../../doc/models/batch-job-response.md)

## Example Usage

```go
ctx := context.Background()
batchId := "batch_id8"

apiResponse, err := APIExportsController.ReadSubscriptionsExport(ctx, batchId)
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

