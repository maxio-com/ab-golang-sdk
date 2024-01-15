# Proforma Invoices

```go
proformaInvoicesController := client.ProformaInvoicesController()
```

## Class Name

`ProformaInvoicesController`

## Methods

* [Create Consolidated Proforma Invoice](../../doc/controllers/proforma-invoices.md#create-consolidated-proforma-invoice)
* [List Subscription Group Proforma Invoices](../../doc/controllers/proforma-invoices.md#list-subscription-group-proforma-invoices)
* [Read Proforma Invoice](../../doc/controllers/proforma-invoices.md#read-proforma-invoice)
* [Create Proforma Invoice](../../doc/controllers/proforma-invoices.md#create-proforma-invoice)
* [List Proforma Invoices](../../doc/controllers/proforma-invoices.md#list-proforma-invoices)
* [Void Proforma Invoice](../../doc/controllers/proforma-invoices.md#void-proforma-invoice)
* [Preview Proforma Invoice](../../doc/controllers/proforma-invoices.md#preview-proforma-invoice)
* [Create Signup Proforma Invoice](../../doc/controllers/proforma-invoices.md#create-signup-proforma-invoice)
* [Preview Signup Proforma Invoice](../../doc/controllers/proforma-invoices.md#preview-signup-proforma-invoice)


# Create Consolidated Proforma Invoice

This endpoint will trigger the creation of a consolidated proforma invoice asynchronously. It will return a 201 with no message, or a 422 with any errors. To find and view the new consolidated proforma invoice, you may poll the subscription group listing for proforma invoices; only one consolidated proforma invoice may be created per group at a time.

If the information becomes outdated, simply void the old consolidated proforma invoice and generate a new one.

## Restrictions

Proforma invoices are only available on Relationship Invoicing sites. To create a proforma invoice, the subscription must not be prepaid, and must be in a live state.

```go
CreateConsolidatedProformaInvoice(
    ctx context.Context,
    uid string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |

## Response Type

``

## Example Usage

```go
ctx := context.Background()
uid := "uid0"

resp, err := proformaInvoicesController.CreateConsolidatedProformaInvoice(ctx, uid)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Subscription Group Proforma Invoices

Only proforma invoices with a `consolidation_level` of parent are returned.

By default, proforma invoices returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, `custom_fields`. To include breakdowns, pass the specific field as a key in the query with a value set to true.

```go
ListSubscriptionGroupProformaInvoices(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.ProformaInvoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The uid of the subscription group |

## Response Type

[`models.ProformaInvoice`](../../doc/models/proforma-invoice.md)

## Example Usage

```go
ctx := context.Background()
uid := "uid0"

apiResponse, err := proformaInvoicesController.ListSubscriptionGroupProformaInvoices(ctx, uid)
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


# Read Proforma Invoice

Use this endpoint to read the details of an existing proforma invoice.

## Restrictions

Proforma invoices are only available on Relationship Invoicing sites.

```go
ReadProformaInvoice(
    ctx context.Context,
    proformaInvoiceUid int) (
    models.ApiResponse[models.ProformaInvoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `proformaInvoiceUid` | `int` | Template, Required | The uid of the proforma invoice |

## Response Type

[`models.ProformaInvoice`](../../doc/models/proforma-invoice.md)

## Example Usage

```go
ctx := context.Background()
proformaInvoiceUid := 242

apiResponse, err := proformaInvoicesController.ReadProformaInvoice(ctx, proformaInvoiceUid)
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


# Create Proforma Invoice

This endpoint will create a proforma invoice and return it as a response. If the information becomes outdated, simply void the old proforma invoice and generate a new one.

If you would like to preview the next billing amounts without generating a full proforma invoice, please use the renewal preview endpoint.

## Restrictions

Proforma invoices are only available on Relationship Invoicing sites. To create a proforma invoice, the subscription must not be in a group, must not be prepaid, and must be in a live state.

```go
CreateProformaInvoice(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.ProformaInvoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |

## Response Type

[`models.ProformaInvoice`](../../doc/models/proforma-invoice.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

apiResponse, err := proformaInvoicesController.CreateProformaInvoice(ctx, subscriptionId)
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


# List Proforma Invoices

By default, proforma invoices returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, or `custom_fields`. To include breakdowns, pass the specific field as a key in the query with a value set to `true`.

```go
ListProformaInvoices(
    ctx context.Context,input ListProformaInvoicesInput) (
    models.ApiResponse[[]models.ProformaInvoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `startDate` | `*string` | Query, Optional | The beginning date range for the invoice's Due Date, in the YYYY-MM-DD format. |
| `endDate` | `*string` | Query, Optional | The ending date range for the invoice's Due Date, in the YYYY-MM-DD format. |
| `status` | [`*models.InvoiceStatus`](../../doc/models/invoice-status.md) | Query, Optional | The current status of the invoice.  Allowed Values: draft, open, paid, pending, voided |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `direction` | [`*models.Direction`](../../doc/models/direction.md) | Query, Optional | The sort direction of the returned invoices. |
| `lineItems` | `*bool` | Query, Optional | Include line items data |
| `discounts` | `*bool` | Query, Optional | Include discounts data |
| `taxes` | `*bool` | Query, Optional | Include taxes data |
| `credits` | `*bool` | Query, Optional | Include credits data |
| `payments` | `*bool` | Query, Optional | Include payments data |
| `customFields` | `*bool` | Query, Optional | Include custom fields data |

## Response Type

[`[]models.ProformaInvoice`](../../doc/models/proforma-invoice.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222
page := 2
perPage := 50
direction := models.Direction("desc")
lineItems := false
discounts := false
taxes := false
credits := false
payments := false
customFields := false

apiResponse, err := proformaInvoicesController.ListProformaInvoices(ctx, subscriptionId, nil, nil, nil, &page, &perPage, &direction, &lineItems, &discounts, &taxes, &credits, &payments, &customFields)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Void Proforma Invoice

This endpoint will void a proforma invoice that has the status "draft".

## Restrictions

Proforma invoices are only available on Relationship Invoicing sites.

Only proforma invoices that have the appropriate status may be reopened. If the invoice identified by {uid} does not have the appropriate status, the response will have HTTP status code 422 and an error message.

A reason for the void operation is required to be included in the request body. If one is not provided, the response will have HTTP status code 422 and an error message.

```go
VoidProformaInvoice(
    ctx context.Context,
    proformaInvoiceUid string,
    body *models.VoidInvoiceRequest) (
    models.ApiResponse[models.ProformaInvoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `proformaInvoiceUid` | `string` | Template, Required | The uid of the proforma invoice |
| `body` | [`*models.VoidInvoiceRequest`](../../doc/models/void-invoice-request.md) | Body, Optional | - |

## Response Type

[`models.ProformaInvoice`](../../doc/models/proforma-invoice.md)

## Example Usage

```go
ctx := context.Background()
proformaInvoiceUid := "proforma_invoice_uid4"

apiResponse, err := proformaInvoicesController.VoidProformaInvoice(ctx, proformaInvoiceUid, nil)
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


# Preview Proforma Invoice

Return a preview of the data that will be included on a given subscription's proforma invoice if one were to be generated. It will have similar line items and totals as a renewal preview, but the response will be presented in the format of a proforma invoice. Consequently it will include additional information such as the name and addresses that will appear on the proforma invoice.

The preview endpoint is subject to all the same conditions as the proforma invoice endpoint. For example, previews are only available on the Relationship Invoicing architecture, and previews cannot be made for end-of-life subscriptions.

If all the data returned in the preview is as expected, you may then create a static proforma invoice and send it to your customer. The data within a preview will not be saved and will not be accessible after the call is made.

Alternatively, if you have some proforma invoices already, you may make a preview call to determine whether any billing information for the subscription's upcoming renewal has changed.

```go
PreviewProformaInvoice(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.ProformaInvoicePreview],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |

## Response Type

[`models.ProformaInvoicePreview`](../../doc/models/proforma-invoice-preview.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := 222

apiResponse, err := proformaInvoicesController.PreviewProformaInvoice(ctx, subscriptionId)
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


# Create Signup Proforma Invoice

This endpoint is only available for Relationship Invoicing sites. It cannot be used to create consolidated proforma invoices or preview prepaid subscriptions.

Create a proforma invoice to preview costs before a subscription's signup. Like other proforma invoices, it can be emailed to the customer, voided, and publicly viewed on the chargifypay domain.

Pass a payload that resembles a subscription create or signup preview request. For example, you can specify components, coupons/a referral, offers, custom pricing, and an existing customer or payment profile to populate a shipping or billing address.

A product and customer first name, last name, and email are the minimum requirements. We recommend associating the proforma invoice with a customer_id to easily find their proforma invoices, since the subscription_id will always be blank.

```go
CreateSignupProformaInvoice(
    ctx context.Context,
    body *models.CreateSubscriptionRequest) (
    models.ApiResponse[models.ProformaInvoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.CreateSubscriptionRequest`](../../doc/models/create-subscription-request.md) | Body, Optional | - |

## Response Type

[`models.ProformaInvoice`](../../doc/models/proforma-invoice.md)

## Example Usage

```go
ctx := context.Background()

bodySubscriptionCustomerAttributes := models.CustomerAttributes{
    FirstName:    models.ToPointer("Myra"),
    LastName:     models.ToPointer("Maisel"),
    Email:        models.ToPointer("mmaisel@example.com"),
}

bodySubscription := models.CreateSubscription{
    ProductHandle:                     models.ToPointer("gold-product"),
    CustomerAttributes:                models.ToPointer(bodySubscriptionCustomerAttributes),
}

body := models.CreateSubscriptionRequest{
    Subscription: bodySubscription,
}

apiResponse, err := proformaInvoicesController.CreateSignupProformaInvoice(ctx, &body)
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
| 400 | Bad Request | [`ProformaBadRequestErrorResponseException`](../../doc/models/proforma-bad-request-error-response-exception.md) |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorMapResponseException`](../../doc/models/error-map-response-exception.md) |


# Preview Signup Proforma Invoice

This endpoint is only available for Relationship Invoicing sites. It cannot be used to create consolidated proforma invoice previews or preview prepaid subscriptions.

Create a signup preview in the format of a proforma invoice to preview costs before a subscription's signup. You have the option of optionally previewing the first renewal's costs as well. The proforma invoice preview will not be persisted.

Pass a payload that resembles a subscription create or signup preview request. For example, you can specify components, coupons/a referral, offers, custom pricing, and an existing customer or payment profile to populate a shipping or billing address.

A product and customer first name, last name, and email are the minimum requirements.

```go
PreviewSignupProformaInvoice(
    ctx context.Context,
    includeNextProformaInvoice *string,
    body *models.CreateSubscriptionRequest) (
    models.ApiResponse[models.SignupProformaPreviewResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `includeNextProformaInvoice` | `*string` | Query, Optional | Choose to include a proforma invoice preview for the first renewal |
| `body` | [`*models.CreateSubscriptionRequest`](../../doc/models/create-subscription-request.md) | Body, Optional | - |

## Response Type

[`models.SignupProformaPreviewResponse`](../../doc/models/signup-proforma-preview-response.md)

## Example Usage

```go
ctx := context.Background()

bodySubscriptionCustomerAttributes := models.CustomerAttributes{
    FirstName:    models.ToPointer("first"),
    LastName:     models.ToPointer("last"),
    Email:        models.ToPointer("flast@example.com"),
}

bodySubscription := models.CreateSubscription{
    ProductHandle:                     models.ToPointer("gold-plan"),
    CustomerAttributes:                models.ToPointer(bodySubscriptionCustomerAttributes),
}

body := models.CreateSubscriptionRequest{
    Subscription: bodySubscription,
}

apiResponse, err := proformaInvoicesController.PreviewSignupProformaInvoice(ctx, nil, &body)
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
| 400 | Bad Request | [`ProformaBadRequestErrorResponseException`](../../doc/models/proforma-bad-request-error-response-exception.md) |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorMapResponseException`](../../doc/models/error-map-response-exception.md) |

