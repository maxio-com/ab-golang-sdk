# Advance Invoice

```go
advanceInvoiceController := client.AdvanceInvoiceController()
```

## Class Name

`AdvanceInvoiceController`

## Methods

* [Issue Advance Invoice](../../doc/controllers/advance-invoice.md#issue-advance-invoice)
* [Read Advance Invoice](../../doc/controllers/advance-invoice.md#read-advance-invoice)
* [Void Advance Invoice](../../doc/controllers/advance-invoice.md#void-advance-invoice)


# Issue Advance Invoice

Generate an invoice in advance for a subscription's next renewal date. [Please see our docs](../../doc/models/invoice.md) for more information on advance invoices, including eligibility on generating one; for the most part, they function like any other invoice, except they are issued early and have special behavior upon being voided.
A subscription may only have one advance invoice per billing period. Attempting to issue an advance invoice when one already exists will return an error.
That said, regeneration of the invoice may be forced with the params `force: true`, which will void an advance invoice if one exists and generate a new one. If no advance invoice exists, a new one will be generated.
We recommend using either the create or preview endpoints for proforma invoices to preview this advance invoice before using this endpoint to generate it.

```go
IssueAdvanceInvoice(
    ctx context.Context,
    subscriptionId int,
    body *models.IssueAdvanceInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.IssueAdvanceInvoiceRequest`](../../doc/models/issue-advance-invoice-request.md) | Body, Optional | - |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

body := models.IssueAdvanceInvoiceRequest{
    Force: models.ToPointer(true),
}

apiResponse, err := advanceInvoiceController.IssueAdvanceInvoice(ctx, subscriptionId, &body)
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
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Read Advance Invoice

Once an advance invoice has been generated for a subscription's upcoming renewal, it can be viewed through this endpoint. There can only be one advance invoice per subscription per billing cycle.

```go
ReadAdvanceInvoice(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

apiResponse, err := advanceInvoiceController.ReadAdvanceInvoice(ctx, subscriptionId)
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


# Void Advance Invoice

Void a subscription's existing advance invoice. Once voided, it can later be regenerated if desired.
A `reason` is required in order to void, and the invoice must have an open status. Voiding will cause any prepayments and credits that were applied to the invoice to be returned to the subscription. For a full overview of the impact of voiding, please [see our help docs](../../doc/models/invoice.md).

```go
VoidAdvanceInvoice(
    ctx context.Context,
    subscriptionId int,
    body *models.VoidInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.VoidInvoiceRequest`](../../doc/models/void-invoice-request.md) | Body, Optional | - |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

apiResponse, err := advanceInvoiceController.VoidAdvanceInvoice(ctx, subscriptionId, nil)
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

