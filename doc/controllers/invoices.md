# Invoices

```go
invoicesController := client.InvoicesController()
```

## Class Name

`InvoicesController`

## Methods

* [Refund Invoice](../../doc/controllers/invoices.md#refund-invoice)
* [List Invoices](../../doc/controllers/invoices.md#list-invoices)
* [Read Invoice](../../doc/controllers/invoices.md#read-invoice)
* [List Invoice Events](../../doc/controllers/invoices.md#list-invoice-events)
* [Record Payment for Invoice](../../doc/controllers/invoices.md#record-payment-for-invoice)
* [Record Payment for Multiple Invoices](../../doc/controllers/invoices.md#record-payment-for-multiple-invoices)
* [List Credit Notes](../../doc/controllers/invoices.md#list-credit-notes)
* [Read Credit Note](../../doc/controllers/invoices.md#read-credit-note)
* [Record Payment for Subscription](../../doc/controllers/invoices.md#record-payment-for-subscription)
* [Reopen Invoice](../../doc/controllers/invoices.md#reopen-invoice)
* [Void Invoice](../../doc/controllers/invoices.md#void-invoice)
* [List Consolidated Invoice Segments](../../doc/controllers/invoices.md#list-consolidated-invoice-segments)
* [Create Invoice](../../doc/controllers/invoices.md#create-invoice)
* [Send Invoice](../../doc/controllers/invoices.md#send-invoice)
* [Preview Customer Information Changes](../../doc/controllers/invoices.md#preview-customer-information-changes)
* [Update Customer Information](../../doc/controllers/invoices.md#update-customer-information)
* [Issue Invoice](../../doc/controllers/invoices.md#issue-invoice)


# Refund Invoice

Refund an invoice, segment, or consolidated invoice.

## Partial Refund for Consolidated Invoice

A refund less than the total of a consolidated invoice will be split across its segments.

A $50.00 refund on a $100.00 consolidated invoice with one $60.00 and one $40.00 segment, the refunded amount will be applied as 50% of each ($30.00 and $20.00 respectively).

```go
RefundInvoice(
    ctx context.Context,
    uid string,
    body *models.RefundInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier for the invoice, this does not refer to the public facing invoice number. |
| `body` | [`*models.RefundInvoiceRequest`](../../doc/models/refund-invoice-request.md) | Body, Optional | - |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

body := models.RefundInvoiceRequest{
    Refund: models.RefundInvoiceRequestRefundContainer.FromRefundInvoice(models.RefundInvoice{
        Amount:      "100.00",
        Memo:        "Refund for Basic Plan renewal",
        PaymentId:   12345,
        External:    models.ToPointer(false),
        ApplyCredit: models.ToPointer(false),
        VoidInvoice: models.ToPointer(true),
    }),
}

apiResponse, err := invoicesController.RefundInvoice(ctx, uid, &body)
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


# List Invoices

By default, invoices returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, `custom_fields`, or `refunds`. To include breakdowns, pass the specific field as a key in the query with a value set to `true`.

```go
ListInvoices(
    ctx context.Context,
    input ListInvoicesInput) (
    models.ApiResponse[models.ListInvoicesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `startDate` | `*string` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns invoices with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `endDate` | `*string` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns invoices with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `status` | [`*models.InvoiceStatus`](../../doc/models/invoice-status.md) | Query, Optional | The current status of the invoice.  Allowed Values: draft, open, paid, pending, voided |
| `subscriptionId` | `*int` | Query, Optional | The subscription's ID. |
| `subscriptionGroupUid` | `*string` | Query, Optional | The UID of the subscription group you want to fetch consolidated invoices for. This will return a paginated list of consolidated invoices for the specified group. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `direction` | [`*models.Direction`](../../doc/models/direction.md) | Query, Optional | The sort direction of the returned invoices. |
| `lineItems` | `*bool` | Query, Optional | Include line items data |
| `discounts` | `*bool` | Query, Optional | Include discounts data |
| `taxes` | `*bool` | Query, Optional | Include taxes data |
| `credits` | `*bool` | Query, Optional | Include credits data |
| `payments` | `*bool` | Query, Optional | Include payments data |
| `customFields` | `*bool` | Query, Optional | Include custom fields data |
| `refunds` | `*bool` | Query, Optional | Include refunds data |
| `dateField` | [`*models.InvoiceDateField`](../../doc/models/invoice-date-field.md) | Query, Optional | The type of filter you would like to apply to your search. Use in query `date_field=issue_date`. |
| `startDatetime` | `*string` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns invoices with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. Allowed to be used only along with date_field set to created_at or updated_at. |
| `endDatetime` | `*string` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns invoices with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. Allowed to be used only along with date_field set to created_at or updated_at. |
| `customerIds` | `[]int` | Query, Optional | Allows fetching invoices with matching customer id based on provided values. Use in query `customer_ids=1,2,3`. |
| `number` | `[]string` | Query, Optional | Allows fetching invoices with matching invoice number based on provided values. Use in query `number=1234,1235`. |
| `productIds` | `[]int` | Query, Optional | Allows fetching invoices with matching line items product ids based on provided values. Use in query `product_ids=23,34`. |
| `sort` | [`*models.InvoiceSortField`](../../doc/models/invoice-sort-field.md) | Query, Optional | Allows specification of the order of the returned list. Use in query `sort=total_amount`. |

## Response Type

[`models.ListInvoicesResponse`](../../doc/models/list-invoices-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListInvoicesInput{
    Page:                 models.ToPointer(2),
    PerPage:              models.ToPointer(50),
    Direction:            models.ToPointer(models.Direction("desc")),
    LineItems:            models.ToPointer(false),
    Discounts:            models.ToPointer(false),
    Taxes:                models.ToPointer(false),
    Credits:              models.ToPointer(false),
    Payments:             models.ToPointer(false),
    CustomFields:         models.ToPointer(false),
    Refunds:              models.ToPointer(false),
    DateField:            models.ToPointer(models.InvoiceDateField("issue_date")),
    CustomerIds:          []int{
        1,
        2,
        3,
    },
    Number:               []string{
        "1234",
        "1235",
    },
    ProductIds:           []int{
        23,
        34,
    },
    Sort:                 models.ToPointer(models.InvoiceSortField("total_amount")),
}

apiResponse, err := invoicesController.ListInvoices(ctx, collectedInput)
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
  "invoices": [
    {
      "uid": "inv_8htcd29wcq3q6",
      "site_id": 51288,
      "customer_id": 20153415,
      "subscription_id": 23277588,
      "number": "125",
      "sequence_number": 125,
      "issue_date": "2018-09-20",
      "due_date": "2018-09-20",
      "paid_date": "2018-09-20",
      "status": "paid",
      "collection_method": "automatic",
      "payment_instructions": "Make checks payable to Acme, Inc.",
      "currency": "USD",
      "consolidation_level": "parent",
      "parent_invoice_uid": null,
      "parent_invoice_number": null,
      "group_primary_subscription_id": 23277588,
      "product_name": "Trial and setup fee",
      "product_family_name": "Billing Plans",
      "seller": {
        "name": "General Goods",
        "address": {
          "street": "123 General Goods Way",
          "line2": "Apt. 10",
          "city": "Boston",
          "state": "MA",
          "zip": "02120",
          "country": "US"
        },
        "phone": "555-555-1212"
      },
      "customer": {
        "chargify_id": 20153415,
        "first_name": "Meg",
        "last_name": "Example",
        "organization": "",
        "email": "meg@example.com"
      },
      "memo": "Please pay within 15 days.",
      "billing_address": {
        "street": "123 I Love Cats Way",
        "line2": "",
        "city": "Boston",
        "state": "MA",
        "zip": "90210",
        "country": "US"
      },
      "shipping_address": {
        "street": "123 I Love Cats Way",
        "line2": "",
        "city": "Boston",
        "state": "MA",
        "zip": "90210",
        "country": "US"
      },
      "subtotal_amount": "100.0",
      "discount_amount": "0.0",
      "tax_amount": "0.0",
      "total_amount": "100.0",
      "credit_amount": "0.0",
      "paid_amount": "100.0",
      "refund_amount": "0.0",
      "due_amount": "0.0",
      "public_url": "https://www.chargifypay.com/invoice/inv_8htcd29wcq3q6?token=n9fr5fxff5v74c7h9srg3cwd"
    },
    {
      "uid": "inv_8hr3546xp4h8n",
      "site_id": 51288,
      "customer_id": 21687686,
      "subscription_id": 22007644,
      "number": "124",
      "sequence_number": 124,
      "issue_date": "2018-09-18",
      "due_date": "2018-09-18",
      "paid_date": null,
      "status": "open",
      "collection_method": "remittance",
      "payment_instructions": "Make checks payable to Acme, Inc.",
      "currency": "USD",
      "consolidation_level": "none",
      "parent_invoice_uid": null,
      "parent_invoice_number": null,
      "group_primary_subscription_id": null,
      "product_name": "Trial and setup fee",
      "product_family_name": "Billing Plans",
      "seller": {
        "name": "General Goods",
        "address": {
          "street": "123 General Goods Way",
          "line2": "Apt. 10",
          "city": "Boston",
          "state": "MA",
          "zip": "02120",
          "country": "US"
        },
        "phone": "555-555-1212"
      },
      "customer": {
        "chargify_id": 21687686,
        "first_name": "Charlene",
        "last_name": "Tester",
        "organization": "",
        "email": "food@example.com"
      },
      "memo": "Please pay within 15 days.",
      "billing_address": {
        "street": "",
        "line2": "",
        "city": "",
        "state": "",
        "zip": "",
        "country": ""
      },
      "shipping_address": {
        "street": "",
        "line2": "",
        "city": "",
        "state": "",
        "zip": "",
        "country": ""
      },
      "subtotal_amount": "100.0",
      "discount_amount": "0.0",
      "tax_amount": "0.0",
      "total_amount": "100.0",
      "credit_amount": "0.0",
      "paid_amount": "0.0",
      "refund_amount": "0.0",
      "due_amount": "100.0",
      "public_url": "https://www.chargifypay.com/invoice/inv_8hr3546xp4h8n?token=n9fr5fxff5v74c7h9srg3cwd"
    },
    {
      "uid": "inv_8hr3546wdwxkr",
      "site_id": 51288,
      "customer_id": 21687670,
      "subscription_id": 22007627,
      "number": "123",
      "sequence_number": 123,
      "issue_date": "2018-09-18",
      "due_date": "2018-09-18",
      "paid_date": "2018-09-18",
      "status": "paid",
      "collection_method": "automatic",
      "payment_instructions": "Make checks payable to Acme, Inc.",
      "currency": "USD",
      "consolidation_level": "none",
      "parent_invoice_uid": null,
      "parent_invoice_number": null,
      "group_primary_subscription_id": null,
      "product_name": "Trial End - Free",
      "product_family_name": "Billing Plans",
      "seller": {
        "name": "General Goods",
        "address": {
          "street": "123 General Goods Way",
          "line2": "Apt. 10",
          "city": "Boston",
          "state": "MA",
          "zip": "02120",
          "country": "US"
        },
        "phone": "555-555-1212"
      },
      "customer": {
        "chargify_id": 21687670,
        "first_name": "Hello",
        "last_name": "World",
        "organization": "123",
        "email": "example@example.com"
      },
      "memo": "Please pay within 15 days.",
      "billing_address": {
        "street": "123 Anywhere Street",
        "line2": "",
        "city": "Boston",
        "state": "MA",
        "zip": "02120",
        "country": "US"
      },
      "shipping_address": {
        "street": "",
        "line2": "",
        "city": "Boston",
        "state": "AL",
        "zip": "02120",
        "country": "US"
      },
      "subtotal_amount": "0.0",
      "discount_amount": "0.0",
      "tax_amount": "0.0",
      "total_amount": "0.0",
      "credit_amount": "0.0",
      "paid_amount": "0.0",
      "refund_amount": "0.0",
      "due_amount": "0.0",
      "public_url": "https://www.chargifypay.com/invoice/inv_8hr3546wdwxkr?token=n9fr5fxff5v74c7h9srg3cwd"
    },
    {
      "uid": "inv_8hjtk8bz56bbp",
      "site_id": 51288,
      "customer_id": 20137757,
      "subscription_id": 20541100,
      "number": "122",
      "sequence_number": 122,
      "issue_date": "2018-09-10",
      "due_date": "2018-09-10",
      "paid_date": "2018-09-10",
      "status": "paid",
      "collection_method": "automatic",
      "payment_instructions": "Make checks payable to Acme, Inc.",
      "currency": "USD",
      "consolidation_level": "none",
      "parent_invoice_uid": null,
      "parent_invoice_number": null,
      "group_primary_subscription_id": null,
      "product_name": "$0 Product",
      "product_family_name": "Billing Plans",
      "seller": {
        "name": "General Goods",
        "address": {
          "street": "123 General Goods Way",
          "line2": "Apt. 10",
          "city": "Boston",
          "state": "MA",
          "zip": "02120",
          "country": "US"
        },
        "phone": "555-555-1212"
      },
      "customer": {
        "chargify_id": 20137757,
        "first_name": "Sasha",
        "last_name": "Example",
        "organization": "",
        "email": "example@example.com"
      },
      "memo": "Please pay within 15 days.",
      "billing_address": {
        "street": "123 I Love Cats Way",
        "line2": "",
        "city": "Catville",
        "state": "MA",
        "zip": "90210",
        "country": "US"
      },
      "shipping_address": {
        "street": "123 I Love Cats Way",
        "line2": "",
        "city": "Catville",
        "state": "AL",
        "zip": "90210",
        "country": "US"
      },
      "subtotal_amount": "0.0",
      "discount_amount": "0.0",
      "tax_amount": "0.0",
      "total_amount": "0.0",
      "credit_amount": "0.0",
      "paid_amount": "0.0",
      "refund_amount": "0.0",
      "due_amount": "0.0",
      "public_url": "https://www.chargifypay.com/invoice/inv_8hjtk8bz56bbp?token=fb6kpjz5rcr2vttyjs4rcv6y"
    }
  ]
}
```


# Read Invoice

Use this endpoint to retrieve the details for an invoice.

```go
ReadInvoice(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier for the invoice, this does not refer to the public facing invoice number. |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

apiResponse, err := invoicesController.ReadInvoice(ctx, uid)
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
  "uid": "inv_8gd8tdhtd3hgr",
  "site_id": 51288,
  "customer_id": 20194505,
  "subscription_id": 20597774,
  "number": "117",
  "sequence_number": 117,
  "issue_date": "2018-07-26",
  "due_date": "2018-07-26",
  "paid_date": "2018-07-26",
  "status": "paid",
  "collection_method": "automatic",
  "payment_instructions": "Make checks payable to Acme, Inc.",
  "currency": "USD",
  "consolidation_level": "none",
  "parent_invoice_uid": null,
  "parent_invoice_number": null,
  "group_primary_subscription_id": null,
  "product_name": "Monthly Product",
  "product_family_name": "Billing Plans",
  "seller": {
    "name": "General Goods",
    "address": {
      "street": "123 General Goods Way",
      "line2": "Apt. 10",
      "city": "Boston",
      "state": "MA",
      "zip": "02120",
      "country": "US"
    },
    "phone": "555-555-1212"
  },
  "customer": {
    "chargify_id": 20194505,
    "first_name": "Joe",
    "last_name": "Example",
    "organization": null,
    "email": "joe@example.com"
  },
  "memo": "Please pay within 15 days.",
  "billing_address": {
    "street": null,
    "line2": null,
    "city": null,
    "state": null,
    "zip": null,
    "country": null
  },
  "shipping_address": {
    "street": null,
    "line2": null,
    "city": null,
    "state": null,
    "zip": null,
    "country": null
  },
  "subtotal_amount": "100.0",
  "discount_amount": "0.0",
  "tax_amount": "0.0",
  "total_amount": "100.0",
  "credit_amount": "0.0",
  "paid_amount": "100.0",
  "refund_amount": "0.0",
  "due_amount": "0.0",
  "line_items": [
    {
      "uid": "li_8gd8tdhhgk55k",
      "title": "Monthly Product",
      "description": "Jul 26, 2018 - Aug 26, 2018",
      "quantity": "1.0",
      "unit_price": "100.0",
      "subtotal_amount": "100.0",
      "discount_amount": "0.0",
      "tax_amount": "0.0",
      "total_amount": "100.0",
      "tiered_unit_price": false,
      "period_range_start": "2018-07-26",
      "period_range_end": "2018-08-26",
      "product_id": 4607632,
      "product_version": 1,
      "component_id": null,
      "price_point_id": null
    }
  ],
  "payments": [
    {
      "transaction_time": "2018-07-26T15:22:02Z",
      "memo": "Joe Example - Monthly Product: Renewal payment",
      "original_amount": "100.0",
      "applied_amount": "100.0",
      "payment_method": {
        "card_brand": "bogus",
        "card_expiration": "10/2020",
        "last_four": null,
        "masked_card_number": "XXXX-XXXX-XXXX-1",
        "type": "credit_card"
      },
      "transaction_id": 253028955,
      "prepayment": false
    }
  ],
  "public_url": "https://www.chargifypay.com/invoice/inv_8jzrw74xq8kxr?token=fb6kpjz5rcr2vttyjs4rcv6y"
}
```


# List Invoice Events

This endpoint returns a list of invoice events. Each event contains event "data" (such as an applied payment) as well as a snapshot of the `invoice` at the time of event completion.

Exposed event types are:

+ issue_invoice
+ apply_credit_note
+ apply_payment
+ refund_invoice
+ void_invoice
+ void_remainder
+ backport_invoice
+ change_invoice_status
+ change_invoice_collection_method
+ remove_payment
+ failed_payment
+ apply_debit_note
+ create_debit_note
+ change_chargeback_status

Invoice events are returned in ascending order.

If both a `since_date` and `since_id` are provided in request parameters, the `since_date` will be used.

Note - invoice events that occurred prior to 09/05/2018 __will not__ contain an `invoice` snapshot.

```go
ListInvoiceEvents(
    ctx context.Context,
    input ListInvoiceEventsInput) (
    models.ApiResponse[models.ListInvoiceEventsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `sinceDate` | `*string` | Query, Optional | The timestamp in a format `YYYY-MM-DD T HH:MM:SS Z`, or `YYYY-MM-DD`(in this case, it returns data from the beginning of the day). of the event from which you want to start the search. All the events before the `since_date` timestamp are not returned in the response. |
| `sinceId` | `*int` | Query, Optional | The ID of the event from which you want to start the search(ID is not included. e.g. if ID is set to 2, then all events with ID 3 and more will be shown) This parameter is not used if since_date is defined. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 100. The maximum allowed values is 200; any per_page value over 200 will be changed to 200. |
| `invoiceUid` | `*string` | Query, Optional | Providing an invoice_uid allows for scoping of the invoice events to a single invoice or credit note. |
| `withChangeInvoiceStatus` | `*string` | Query, Optional | Use this parameter if you want to fetch also invoice events with change_invoice_status type. |
| `eventTypes` | [`[]models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Query, Optional | Filter results by event_type. Supply a comma separated list of event types (listed above). Use in query: `event_types=void_invoice,void_remainder`. |

## Response Type

[`models.ListInvoiceEventsResponse`](../../doc/models/list-invoice-events-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListInvoiceEventsInput{
    Page:                    models.ToPointer(2),
    PerPage:                 models.ToPointer(100),
}

apiResponse, err := invoicesController.ListInvoiceEvents(ctx, collectedInput)
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
  "events": [
    {
      "id": 83,
      "event_type": "apply_payment",
      "event_data": {
        "memo": "Non-Resumable Canceled On Purpose - Standard Plan: Renewal payment",
        "original_amount": "168.61",
        "applied_amount": "168.61",
        "transaction_time": "2018-08-01T16:00:00Z",
        "payment_method": {
          "card_brand": "visa",
          "card_expiration": "12/2022",
          "last_four": null,
          "masked_card_number": "XXXX-XXXX-XXXX-1111",
          "type": "credit_card"
        },
        "consolidation_level": "none"
      },
      "timestamp": "2018-08-01T16:00:00Z",
      "invoice": {
        "id": 614942008934401500,
        "uid": "inv_8gk5bwkct3gqt",
        "site_id": 20,
        "customer_id": 6,
        "subscription_id": 10,
        "number": "25",
        "sequence_number": 25,
        "transaction_time": "2018-08-01T16:00:00Z",
        "created_at": "2018-08-01T16:00:00Z",
        "updated_at": "2018-08-01T16:00:00Z",
        "issue_date": "2018-08-01",
        "due_date": "2018-08-01",
        "paid_date": "2018-08-01",
        "status": "paid",
        "role": "renewal",
        "collection_method": "automatic",
        "payment_instructions": "Please make checks payable to \"Acme, Inc.\"",
        "currency": "USD",
        "consolidation_level": "none",
        "parent_invoice_id": null,
        "subscription_group_id": null,
        "parent_invoice_number": null,
        "product_name": "Standard Plan",
        "product_family_name": "Cloud Compute Servers",
        "seller": {
          "name": "Acme, Inc.",
          "address": {
            "street": null,
            "line2": null,
            "city": null,
            "state": null,
            "zip": null,
            "country": null
          },
          "phone": "555-555-1234 x137",
          "logo_url": null
        },
        "customer": {
          "chargify_id": 6,
          "first_name": "Non-Resumable",
          "last_name": "Canceled On Purpose",
          "organization": null,
          "email": "evan4@example.com"
        },
        "payer": {
          "chargify_id": 6,
          "first_name": "Non-Resumable",
          "last_name": "Canceled On Purpose",
          "organization": null,
          "email": "evan4@example.com"
        },
        "net_terms": 0,
        "memo": "Thanks for your business! If you have any questions, please contact your account manager.",
        "billing_address": {
          "street": "200 Billing Rd.",
          "line2": "Suite 100",
          "city": "Needham",
          "state": "MA",
          "zip": "02494",
          "country": "US"
        },
        "shipping_address": {
          "street": "100 Shipping St.",
          "line2": "Apt 200",
          "city": "Pleasantville",
          "state": "NC",
          "zip": "12345",
          "country": "US"
        },
        "line_items": [
          {
            "uid": "li_8gk5bwkct3gqk",
            "title": "Standard Plan",
            "description": "08/01/2018 - 09/01/2018",
            "quantity": "1.0",
            "unit_price": "99.0",
            "subtotal_amount": "99.0",
            "discount_amount": "9.9",
            "tax_amount": "6.01425",
            "total_amount": "95.11425",
            "tiered_unit_price": false,
            "period_range_start": "2018-08-01",
            "period_range_end": "2018-09-01",
            "transaction_id": 120,
            "product_id": 84,
            "product_version": 1,
            "component_id": null,
            "price_point_id": null,
            "hide": false
          },
          {
            "uid": "li_8gk5bwkct3gqm",
            "title": "Small Instance (Hourly)",
            "description": "07/22/2018 - 08/01/2018",
            "quantity": "162.0",
            "unit_price": "0.09567901",
            "subtotal_amount": "15.5",
            "discount_amount": "1.55",
            "tax_amount": "0.941625",
            "total_amount": "14.891625",
            "tiered_unit_price": true,
            "period_range_start": "2018-07-22",
            "period_range_end": "2018-08-01",
            "transaction_id": 121,
            "product_id": 84,
            "product_version": 1,
            "component_id": 76,
            "price_point_id": null,
            "hide": false,
            "component_cost_data": {
              "rates": [
                {
                  "component_code_id": null,
                  "price_point_id": 160,
                  "product_id": 84,
                  "quantity": "162.0",
                  "amount": "15.5",
                  "pricing_scheme": "tiered",
                  "tiers": [
                    {
                      "starting_quantity": 1,
                      "ending_quantity": 100,
                      "quantity": "100.0",
                      "unit_price": "0.0",
                      "amount": "0.0"
                    },
                    {
                      "starting_quantity": 101,
                      "ending_quantity": null,
                      "quantity": "62.0",
                      "unit_price": "0.25",
                      "amount": "15.5"
                    }
                  ]
                }
              ]
            }
          },
          {
            "uid": "li_8gk5bwkct3gqn",
            "title": "Large Instance (Hourly)",
            "description": "07/22/2018 - 08/01/2018",
            "quantity": "194.0",
            "unit_price": "0.24226804",
            "subtotal_amount": "47.0",
            "discount_amount": "4.7",
            "tax_amount": "2.85525",
            "total_amount": "45.15525",
            "tiered_unit_price": true,
            "period_range_start": "2018-07-22",
            "period_range_end": "2018-08-01",
            "transaction_id": 122,
            "product_id": 84,
            "product_version": 1,
            "component_id": 77,
            "price_point_id": null,
            "hide": false,
            "component_cost_data": {
              "rates": [
                {
                  "component_code_id": null,
                  "price_point_id": 161,
                  "product_id": 84,
                  "quantity": "194.0",
                  "amount": "47.0",
                  "pricing_scheme": "tiered",
                  "tiers": [
                    {
                      "starting_quantity": 1,
                      "ending_quantity": 100,
                      "quantity": "100.0",
                      "unit_price": "0.0",
                      "amount": "0.0"
                    },
                    {
                      "starting_quantity": 101,
                      "ending_quantity": null,
                      "quantity": "94.0",
                      "unit_price": "0.5",
                      "amount": "47.0"
                    }
                  ]
                }
              ]
            }
          },
          {
            "uid": "li_8gk5bwkct3gqp",
            "title": "IP Addresses",
            "description": "08/01/2018 - 09/01/2018",
            "quantity": "7.0",
            "unit_price": "2.0",
            "subtotal_amount": "14.0",
            "discount_amount": "1.4",
            "tax_amount": "0.8505",
            "total_amount": "13.4505",
            "tiered_unit_price": false,
            "period_range_start": "2018-08-01",
            "period_range_end": "2018-09-01",
            "transaction_id": 123,
            "product_id": 84,
            "product_version": 1,
            "component_id": 79,
            "price_point_id": 163,
            "hide": false,
            "component_cost_data": {
              "rates": [
                {
                  "component_code_id": null,
                  "price_point_id": 163,
                  "product_id": 84,
                  "quantity": "7.0",
                  "amount": "14.0",
                  "pricing_scheme": "per_unit",
                  "tiers": [
                    {
                      "starting_quantity": 1,
                      "ending_quantity": null,
                      "quantity": "7.0",
                      "unit_price": "2.0",
                      "amount": "14.0"
                    }
                  ]
                }
              ]
            }
          }
        ],
        "subtotal_amount": "175.5",
        "discount_amount": "17.55",
        "discounts": [
          {
            "uid": "dli_8gk5bwkct3gqq",
            "title": "Multi-service discount (10%)",
            "description": null,
            "code": "MULTI3",
            "source_type": "Coupon",
            "source_id": 40,
            "discount_type": "percentage",
            "percentage": "10.0",
            "eligible_amount": "175.5",
            "discount_amount": "17.55",
            "transaction_id": 124,
            "line_item_breakouts": [
              {
                "uid": "li_8gk5bwkct3gqk",
                "eligible_amount": "99.0",
                "discount_amount": "9.9"
              },
              {
                "uid": "li_8gk5bwkct3gqm",
                "eligible_amount": "15.5",
                "discount_amount": "1.55"
              },
              {
                "uid": "li_8gk5bwkct3gqn",
                "eligible_amount": "47.0",
                "discount_amount": "4.7"
              },
              {
                "uid": "li_8gk5bwkct3gqp",
                "eligible_amount": "14.0",
                "discount_amount": "1.4"
              }
            ]
          }
        ],
        "tax_amount": "10.66",
        "taxes": [
          {
            "uid": "tli_8gk5bwkct3gqr",
            "title": "NC Sales Tax",
            "description": null,
            "source_type": "Tax",
            "source_id": 1,
            "percentage": "6.75",
            "taxable_amount": "157.95",
            "tax_amount": "10.66",
            "transaction_id": 125,
            "line_item_breakouts": [
              {
                "uid": "li_8gk5bwkct3gqk",
                "taxable_amount": "89.1",
                "tax_amount": "6.01425"
              },
              {
                "uid": "li_8gk5bwkct3gqm",
                "taxable_amount": "13.95",
                "tax_amount": "0.941625"
              },
              {
                "uid": "li_8gk5bwkct3gqn",
                "taxable_amount": "42.3",
                "tax_amount": "2.85525"
              },
              {
                "uid": "li_8gk5bwkct3gqp",
                "taxable_amount": "12.6",
                "tax_amount": "0.8505"
              }
            ],
            "tax_component_breakouts": [
              {
                "tax_rule_id": 1,
                "percentage": "6.75",
                "country_code": "US",
                "subdivision_code": "NC"
              }
            ]
          }
        ],
        "credit_amount": "0.0",
        "refund_amount": "0.0",
        "total_amount": "168.61",
        "paid_amount": "168.61",
        "due_amount": "0.0",
        "payments": [
          {
            "memo": "Non-Resumable Canceled On Purpose - Standard Plan: Renewal payment",
            "original_amount": "168.61",
            "applied_amount": "168.61",
            "transaction_time": "2018-08-01T16:00:00Z",
            "payment_method": {
              "card_brand": "visa",
              "card_expiration": "12/2022",
              "last_four": null,
              "masked_card_number": "XXXX-XXXX-XXXX-1111",
              "type": "credit_card"
            },
            "transaction_id": 126,
            "prepayment": false
          }
        ],
        "display_settings": {
          "hide_zero_subtotal_lines": false,
          "include_discounts_on_lines": false
        }
      }
    }
  ],
  "page": 48,
  "per_page": 1,
  "total_pages": 102
}
```


# Record Payment for Invoice

This API call should be used when you want to record a payment of a given type against a specific invoice. If you would like to apply a payment across multiple invoices, you can use the Bulk Payment endpoint.

## Create a Payment from the existing payment profile

In order to apply a payment to an invoice using an existing payment profile, specify `type` as `payment`, the amount less than the invoice total, and the customer's `payment_profile_id`. The ID of a payment profile might be retrieved via the Payment Profiles API endpoint.

```
{
  "type": "payment",
  "payment": {
    "amount": 10.00,
    "payment_profile_id": 123
  }
}
```

## Create a Payment from the Subscription's Prepayment Account

In order apply a prepayment to an invoice, specify the `type` as `prepayment`, and also the `amount`.

```
{
  "type": "prepayment",
  "payment": {
    "amount": 10.00
  }
}
```

Note that the `amount` must be less than or equal to the Subscription's Prepayment account balance.

## Create a Payment from the Subscription's Service Credit Account

In order to apply a service credit to an invoice, specify the `type` as `service_credit`, and also the `amount`:

```
{
  "type": "service_credit",
  "payment": {
    "amount": 10.00
  }
}
```

Note that Chargify will attempt to fully pay the invoice's `due_amount` from the Subscription's Service Credit account. At this time, partial payments from a Service Credit Account are only allowed for consolidated invoices (subscription groups). Therefore, for normal invoices the Service Credit account balance must be greater than or equal to the invoice's `due_amount`.

```go
RecordPaymentForInvoice(
    ctx context.Context,
    uid string,
    body *models.CreateInvoicePaymentRequest) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier for the invoice, this does not refer to the public facing invoice number. |
| `body` | [`*models.CreateInvoicePaymentRequest`](../../doc/models/create-invoice-payment-request.md) | Body, Optional | - |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

body := models.CreateInvoicePaymentRequest{
    Payment: models.CreateInvoicePayment{
        Amount:           models.ToPointer(models.CreateInvoicePaymentAmountContainer.FromPrecision(float64(124.33))),
        Memo:             models.ToPointer("for John Smith"),
        Method:           models.ToPointer(models.InvoicePaymentMethodType("check")),
        Details:          models.ToPointer("#0102"),
    },
}

apiResponse, err := invoicesController.RecordPaymentForInvoice(ctx, uid, &body)
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
| 422 | Unprocessable Entity | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Record Payment for Multiple Invoices

This API call should be used when you want to record an external payment against multiple invoices.

In order apply a payment to multiple invoices, at minimum, specify the `amount` and `applications` (i.e., `invoice_uid` and `amount`) details.

```
{
  "payment": {
    "memo": "to pay the bills",
    "details": "check number 8675309",
    "method": "check",
    "amount": "250.00",
    "applications": [
      {
        "invoice_uid": "inv_8gk5bwkct3gqt",
        "amount": "100.00"
      },
      {
        "invoice_uid": "inv_7bc6bwkct3lyt",
        "amount": "150.00"
      }
    ]
  }
}
```

Note that the invoice payment amounts must be greater than 0. Total amount must be greater or equal to invoices payment amount sum.

```go
RecordPaymentForMultipleInvoices(
    ctx context.Context,
    body *models.CreateMultiInvoicePaymentRequest) (
    models.ApiResponse[models.MultiInvoicePaymentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.CreateMultiInvoicePaymentRequest`](../../doc/models/create-multi-invoice-payment-request.md) | Body, Optional | - |

## Response Type

[`models.MultiInvoicePaymentResponse`](../../doc/models/multi-invoice-payment-response.md)

## Example Usage

```go
ctx := context.Background()

body := models.CreateMultiInvoicePaymentRequest{
    Payment: models.CreateMultiInvoicePayment{
        Memo:         models.ToPointer("to pay the bills"),
        Details:      models.ToPointer("check number 8675309"),
        Method:       models.ToPointer(models.InvoicePaymentMethodType("check")),
        Amount:       models.CreateMultiInvoicePaymentAmountContainer.FromString("100.00"),
        Applications: []models.CreateInvoicePaymentApplication{
            models.CreateInvoicePaymentApplication{
                InvoiceUid: "inv_8gk5bwkct3gqt",
                Amount:     "50.00",
            },
            models.CreateInvoicePaymentApplication{
                InvoiceUid: "inv_7bc6bwkct3lyt",
                Amount:     "50.00",
            },
        },
    },
}

apiResponse, err := invoicesController.RecordPaymentForMultipleInvoices(ctx, &body)
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
  "payment": {
    "transaction_id": 1,
    "total_amount": "100.00",
    "currency_code": "USD",
    "applications": [
      {
        "invoice_uid": "inv_8gk5bwkct3gqt",
        "application_uid": "pmt_1tr0hgsct3ybx",
        "applied_amount": "50.00"
      },
      {
        "invoice_uid": "inv_7bc6bwkct3lyt",
        "application_uid": "pmt_2",
        "applied_amount": "50.00"
      }
    ]
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Credit Notes

Credit Notes are like inverse invoices. They reduce the amount a customer owes.

By default, the credit notes returned by this endpoint will exclude the arrays of `line_items`, `discounts`, `taxes`, `applications`, or `refunds`. To include these arrays, pass the specific field as a key in the query with a value set to `true`.

```go
ListCreditNotes(
    ctx context.Context,
    input ListCreditNotesInput) (
    models.ApiResponse[models.ListCreditNotesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `*int` | Query, Optional | The subscription's Chargify id |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `lineItems` | `*bool` | Query, Optional | Include line items data |
| `discounts` | `*bool` | Query, Optional | Include discounts data |
| `taxes` | `*bool` | Query, Optional | Include taxes data |
| `refunds` | `*bool` | Query, Optional | Include refunds data |
| `applications` | `*bool` | Query, Optional | Include applications data |

## Response Type

[`models.ListCreditNotesResponse`](../../doc/models/list-credit-notes-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListCreditNotesInput{
    Page:           models.ToPointer(2),
    PerPage:        models.ToPointer(50),
    LineItems:      models.ToPointer(false),
    Discounts:      models.ToPointer(false),
    Taxes:          models.ToPointer(false),
    Refunds:        models.ToPointer(false),
    Applications:   models.ToPointer(false),
}

apiResponse, err := invoicesController.ListCreditNotes(ctx, collectedInput)
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
  "credit_notes": [
    {
      "uid": "cn_8m9vbd5kkv7kr",
      "site_id": 20,
      "customer_id": 3,
      "subscription_id": 2,
      "number": "77",
      "sequence_number": 78,
      "issue_date": "2018-12-31",
      "applied_date": "2018-12-31",
      "status": "applied",
      "currency": "USD",
      "memo": "Refund for overpayment",
      "seller": {
        "name": "Acme, Inc.",
        "address": {
          "street": "122 E Houston St",
          "line2": "Suite 105",
          "city": "San Antonio",
          "state": "TX",
          "zip": "78205",
          "country": "US"
        },
        "phone": "555-555-1234 x137"
      },
      "customer": {
        "chargify_id": 3,
        "first_name": "Marty",
        "last_name": "McFly",
        "organization": "Time Travellers, Inc.",
        "email": "timetraveller1985@example.com",
        "reference": null
      },
      "billing_address": {
        "street": "200 Billing Rd.",
        "line2": "Suite 100",
        "city": "Needham",
        "state": "MA",
        "zip": "02494",
        "country": "US"
      },
      "shipping_address": {
        "street": "100 Shipping St.",
        "line2": "Apt 200",
        "city": "Pleasantville",
        "state": "NC",
        "zip": "12345",
        "country": "US"
      },
      "subtotal_amount": "208.69341779",
      "discount_amount": "20.87125167",
      "tax_amount": "12.67783387",
      "total_amount": "200.5",
      "applied_amount": "200.5",
      "remaining_amount": "0.0",
      "line_items": [
        {
          "uid": "cnli_8k5jvdzct4h9x",
          "title": "IP Addresses: 5 to 10 addresses",
          "description": "38.2% credit",
          "quantity": "0.9855",
          "unit_price": "2.0",
          "subtotal_amount": "1.971004",
          "discount_amount": "0.19862831",
          "tax_amount": "0.11963536",
          "total_amount": "1.89201105",
          "tiered_unit_price": false,
          "period_range_start": "2018-11-30",
          "period_range_end": "2018-11-30",
          "product_id": 85,
          "product_version": 1,
          "component_id": 81,
          "price_point_id": 165
        },
        {
          "uid": "cnli_8kjttvjcjx8b4",
          "title": "Professional Plan",
          "description": "38.2% credit",
          "quantity": "0.382",
          "unit_price": "299.0",
          "subtotal_amount": "114.21127834",
          "discount_amount": "11.42112783",
          "tax_amount": "6.93833516",
          "total_amount": "109.72848567",
          "tiered_unit_price": false,
          "period_range_start": "2018-12-30",
          "period_range_end": "2018-12-30",
          "product_id": 85,
          "product_version": 1,
          "component_id": null,
          "price_point_id": null
        },
        {
          "uid": "cnli_8kjttvjknzhx7",
          "title": "Small Instance (Hourly)",
          "description": "38.2% credit",
          "quantity": "74.8676",
          "unit_price": "0.12244898",
          "subtotal_amount": "9.16746047",
          "discount_amount": "0.91674605",
          "tax_amount": "0.55692322",
          "total_amount": "8.80763764",
          "tiered_unit_price": true,
          "period_range_start": "2018-11-30",
          "period_range_end": "2018-11-30",
          "product_id": 85,
          "product_version": 1,
          "component_id": 78,
          "price_point_id": null
        },
        {
          "uid": "cnli_8kjttvjnmh25w",
          "title": "Large Instance (Hourly)",
          "description": "38.2% credit",
          "quantity": "183.3492",
          "unit_price": "0.39583333",
          "subtotal_amount": "72.57572871",
          "discount_amount": "7.25757287",
          "tax_amount": "4.40897552",
          "total_amount": "69.72713136",
          "tiered_unit_price": true,
          "period_range_start": "2018-11-30",
          "period_range_end": "2018-11-30",
          "product_id": 85,
          "product_version": 1,
          "component_id": 79,
          "price_point_id": null
        },
        {
          "uid": "cnli_8kjttvjqn86kc",
          "title": "Email Messages",
          "description": "38.2% credit",
          "quantity": "10076.9489",
          "unit_price": "0.00031045",
          "subtotal_amount": "3.12839588",
          "discount_amount": "0.31322157",
          "tax_amount": "0.19002427",
          "total_amount": "3.00519858",
          "tiered_unit_price": true,
          "period_range_start": "2018-11-30",
          "period_range_end": "2018-11-30",
          "product_id": 85,
          "product_version": 1,
          "component_id": 80,
          "price_point_id": null
        },
        {
          "uid": "cnli_8kjttvjtxxbdd",
          "title": "IP Addresses",
          "description": "38.2% credit",
          "quantity": "3.8198",
          "unit_price": "2.0",
          "subtotal_amount": "7.63955039",
          "discount_amount": "0.76395504",
          "tax_amount": "0.46410269",
          "total_amount": "7.33969804",
          "tiered_unit_price": false,
          "period_range_start": "2018-12-30",
          "period_range_end": "2018-12-30",
          "product_id": 85,
          "product_version": 1,
          "component_id": 81,
          "price_point_id": 165
        }
      ],
      "discounts": [
        {
          "uid": "cndli_8k5jvdzct4h9y",
          "title": "Multi-service discount (10%)",
          "code": "MULTI3",
          "source_type": "Coupon",
          "source_id": 40,
          "discount_type": "percentage",
          "percentage": "10.0",
          "eligible_amount": "208.69341779",
          "discount_amount": "20.87125167",
          "line_item_breakouts": [
            {
              "uid": "cnli_8k5jvdzct4h9x",
              "eligible_amount": "1.971004",
              "discount_amount": "0.19862831"
            },
            {
              "uid": "cnli_8kjttvjcjx8b4",
              "eligible_amount": "114.21127834",
              "discount_amount": "11.42112783"
            },
            {
              "uid": "cnli_8kjttvjknzhx7",
              "eligible_amount": "9.16746047",
              "discount_amount": "0.91674605"
            },
            {
              "uid": "cnli_8kjttvjnmh25w",
              "eligible_amount": "72.57572871",
              "discount_amount": "7.25757287"
            },
            {
              "uid": "cnli_8kjttvjqn86kc",
              "eligible_amount": "3.12839588",
              "discount_amount": "0.31322157"
            },
            {
              "uid": "cnli_8kjttvjtxxbdd",
              "eligible_amount": "7.63955039",
              "discount_amount": "0.76395504"
            }
          ]
        }
      ],
      "taxes": [
        {
          "uid": "cntli_8k5jvdzct4h9z",
          "title": "NC Sales Tax",
          "source_type": "Tax",
          "source_id": 1,
          "percentage": "6.75",
          "taxable_amount": "187.82216613",
          "tax_amount": "12.67783387",
          "line_item_breakouts": [
            {
              "uid": "cnli_8k5jvdzct4h9x",
              "taxable_amount": "1.77237569",
              "tax_amount": "0.11963536"
            },
            {
              "uid": "cnli_8kjttvjcjx8b4",
              "taxable_amount": "102.7901505",
              "tax_amount": "6.93833516"
            },
            {
              "uid": "cnli_8kjttvjknzhx7",
              "taxable_amount": "8.25071442",
              "tax_amount": "0.55692322"
            },
            {
              "uid": "cnli_8kjttvjnmh25w",
              "taxable_amount": "65.31815584",
              "tax_amount": "4.40897552"
            },
            {
              "uid": "cnli_8kjttvjqn86kc",
              "taxable_amount": "2.81517432",
              "tax_amount": "0.19002427"
            },
            {
              "uid": "cnli_8kjttvjtxxbdd",
              "taxable_amount": "6.87559535",
              "tax_amount": "0.46410269"
            }
          ]
        }
      ],
      "applications": [
        {
          "uid": "cdt_8m9vbdbdwd28n",
          "transaction_time": "2018-12-31T21:19:28Z",
          "invoice_uid": "inv_8k5jvdzct4hb2",
          "memo": "Refund for overpayment",
          "applied_amount": "200.5"
        }
      ],
      "refunds": [
        {
          "transaction_id": 329,
          "payment_id": 39,
          "memo": "Refund for overpayment",
          "original_amount": "524.9",
          "applied_amount": "200.5"
        }
      ]
    }
  ]
}
```


# Read Credit Note

Use this endpoint to retrieve the details for a credit note.

```go
ReadCreditNote(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.CreditNote],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier of the credit note |

## Response Type

[`models.CreditNote`](../../doc/models/credit-note.md)

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

apiResponse, err := invoicesController.ReadCreditNote(ctx, uid)
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
  "uid": "cn_8m9vbd5kkv7kr",
  "site_id": 20,
  "customer_id": 3,
  "subscription_id": 2,
  "number": "77",
  "sequence_number": 78,
  "issue_date": "2018-12-31",
  "applied_date": "2018-12-31",
  "status": "applied",
  "currency": "USD",
  "memo": "Refund for overpayment",
  "seller": {
    "name": "Acme, Inc.",
    "address": {
      "street": "122 E Houston St",
      "line2": "Suite 105",
      "city": "San Antonio",
      "state": "TX",
      "zip": "78205",
      "country": "US"
    },
    "phone": "555-555-1234 x137"
  },
  "customer": {
    "chargify_id": 3,
    "first_name": "Marty",
    "last_name": "McFly",
    "organization": "Time Travellers, Inc.",
    "email": "timetraveller1985@example.com",
    "reference": null
  },
  "billing_address": {
    "street": "200 Billing Rd.",
    "line2": "Suite 100",
    "city": "Needham",
    "state": "MA",
    "zip": "02494",
    "country": "US"
  },
  "shipping_address": {
    "street": "100 Shipping St.",
    "line2": "Apt 200",
    "city": "Pleasantville",
    "state": "NC",
    "zip": "12345",
    "country": "US"
  },
  "subtotal_amount": "208.69341779",
  "discount_amount": "20.87125167",
  "tax_amount": "12.67783387",
  "total_amount": "200.5",
  "applied_amount": "200.5",
  "remaining_amount": "0.0",
  "line_items": [
    {
      "uid": "cnli_8k5jvdzct4h9x",
      "title": "IP Addresses: 5 to 10 addresses",
      "description": "38.2% credit",
      "quantity": "0.9855",
      "unit_price": "2.0",
      "subtotal_amount": "1.971004",
      "discount_amount": "0.19862831",
      "tax_amount": "0.11963536",
      "total_amount": "1.89201105",
      "tiered_unit_price": false,
      "period_range_start": "2018-11-30",
      "period_range_end": "2018-11-30",
      "product_id": 85,
      "product_version": 1,
      "component_id": 81,
      "price_point_id": 165,
      "billing_schedule_item_id": null,
      "custom_item": false
    },
    {
      "uid": "cnli_8kjttvjcjx8b4",
      "title": "Professional Plan",
      "description": "38.2% credit",
      "quantity": "0.382",
      "unit_price": "299.0",
      "subtotal_amount": "114.21127834",
      "discount_amount": "11.42112783",
      "tax_amount": "6.93833516",
      "total_amount": "109.72848567",
      "tiered_unit_price": false,
      "period_range_start": "2018-12-30",
      "period_range_end": "2018-12-30",
      "product_id": 85,
      "product_version": 1,
      "component_id": null,
      "price_point_id": null,
      "billing_schedule_item_id": null,
      "custom_item": false
    },
    {
      "uid": "cnli_8kjttvjknzhx7",
      "title": "Small Instance (Hourly)",
      "description": "38.2% credit",
      "quantity": "74.8676",
      "unit_price": "0.12244898",
      "subtotal_amount": "9.16746047",
      "discount_amount": "0.91674605",
      "tax_amount": "0.55692322",
      "total_amount": "8.80763764",
      "tiered_unit_price": true,
      "period_range_start": "2018-11-30",
      "period_range_end": "2018-11-30",
      "product_id": 85,
      "product_version": 1,
      "component_id": 78,
      "price_point_id": null,
      "billing_schedule_item_id": null,
      "custom_item": false
    },
    {
      "uid": "cnli_8kjttvjnmh25w",
      "title": "Large Instance (Hourly)",
      "description": "38.2% credit",
      "quantity": "183.3492",
      "unit_price": "0.39583333",
      "subtotal_amount": "72.57572871",
      "discount_amount": "7.25757287",
      "tax_amount": "4.40897552",
      "total_amount": "69.72713136",
      "tiered_unit_price": true,
      "period_range_start": "2018-11-30",
      "period_range_end": "2018-11-30",
      "product_id": 85,
      "product_version": 1,
      "component_id": 79,
      "price_point_id": null,
      "billing_schedule_item_id": null,
      "custom_item": false
    },
    {
      "uid": "cnli_8kjttvjqn86kc",
      "title": "Email Messages",
      "description": "38.2% credit",
      "quantity": "10076.9489",
      "unit_price": "0.00031045",
      "subtotal_amount": "3.12839588",
      "discount_amount": "0.31322157",
      "tax_amount": "0.19002427",
      "total_amount": "3.00519858",
      "tiered_unit_price": true,
      "period_range_start": "2018-11-30",
      "period_range_end": "2018-11-30",
      "product_id": 85,
      "product_version": 1,
      "component_id": 80,
      "price_point_id": null,
      "billing_schedule_item_id": null,
      "custom_item": false
    },
    {
      "uid": "cnli_8kjttvjtxxbdd",
      "title": "IP Addresses",
      "description": "38.2% credit",
      "quantity": "3.8198",
      "unit_price": "2.0",
      "subtotal_amount": "7.63955039",
      "discount_amount": "0.76395504",
      "tax_amount": "0.46410269",
      "total_amount": "7.33969804",
      "tiered_unit_price": false,
      "period_range_start": "2018-12-30",
      "period_range_end": "2018-12-30",
      "product_id": 85,
      "product_version": 1,
      "component_id": 81,
      "price_point_id": 165,
      "billing_schedule_item_id": null,
      "custom_item": false
    }
  ],
  "discounts": [
    {
      "uid": "cndli_8k5jvdzct4h9y",
      "title": "Multi-service discount (10%)",
      "code": "MULTI3",
      "source_type": "Coupon",
      "source_id": 40,
      "discount_type": "percentage",
      "percentage": "10.0",
      "eligible_amount": "208.69341779",
      "discount_amount": "20.87125167",
      "line_item_breakouts": [
        {
          "uid": "cnli_8k5jvdzct4h9x",
          "eligible_amount": "1.971004",
          "discount_amount": "0.19862831"
        },
        {
          "uid": "cnli_8kjttvjcjx8b4",
          "eligible_amount": "114.21127834",
          "discount_amount": "11.42112783"
        },
        {
          "uid": "cnli_8kjttvjknzhx7",
          "eligible_amount": "9.16746047",
          "discount_amount": "0.91674605"
        },
        {
          "uid": "cnli_8kjttvjnmh25w",
          "eligible_amount": "72.57572871",
          "discount_amount": "7.25757287"
        },
        {
          "uid": "cnli_8kjttvjqn86kc",
          "eligible_amount": "3.12839588",
          "discount_amount": "0.31322157"
        },
        {
          "uid": "cnli_8kjttvjtxxbdd",
          "eligible_amount": "7.63955039",
          "discount_amount": "0.76395504"
        }
      ]
    }
  ],
  "taxes": [
    {
      "uid": "cntli_8k5jvdzct4h9z",
      "title": "NC Sales Tax",
      "source_type": "Tax",
      "source_id": 1,
      "percentage": "6.75",
      "taxable_amount": "187.82216613",
      "tax_amount": "12.67783387",
      "line_item_breakouts": [
        {
          "uid": "cnli_8k5jvdzct4h9x",
          "taxable_amount": "1.77237569",
          "tax_amount": "0.11963536"
        },
        {
          "uid": "cnli_8kjttvjcjx8b4",
          "taxable_amount": "102.7901505",
          "tax_amount": "6.93833516"
        },
        {
          "uid": "cnli_8kjttvjknzhx7",
          "taxable_amount": "8.25071442",
          "tax_amount": "0.55692322"
        },
        {
          "uid": "cnli_8kjttvjnmh25w",
          "taxable_amount": "65.31815584",
          "tax_amount": "4.40897552"
        },
        {
          "uid": "cnli_8kjttvjqn86kc",
          "taxable_amount": "2.81517432",
          "tax_amount": "0.19002427"
        },
        {
          "uid": "cnli_8kjttvjtxxbdd",
          "taxable_amount": "6.87559535",
          "tax_amount": "0.46410269"
        }
      ]
    }
  ],
  "applications": [
    {
      "uid": "cdt_8m9vbdbdwd28n",
      "transaction_time": "2018-12-31T21:19:28Z",
      "invoice_uid": "inv_8k5jvdzct4hb2",
      "memo": "Refund for overpayment",
      "applied_amount": "200.5"
    }
  ],
  "refunds": [
    {
      "transaction_id": 329,
      "payment_id": 39,
      "memo": "Refund for overpayment",
      "original_amount": "524.9",
      "applied_amount": "200.5"
    }
  ]
}
```


# Record Payment for Subscription

Record an external payment made against a subscription that will pay partially or in full one or more invoices.

Payment will be applied starting with the oldest open invoice and then next oldest, and so on until the amount of the payment is fully consumed.

Excess payment will result in the creation of a prepayment on the Invoice Account.

Only ungrouped or primary subscriptions may be paid using the "bulk" payment request.

```go
RecordPaymentForSubscription(
    ctx context.Context,
    subscriptionId int,
    body *models.RecordPaymentRequest) (
    models.ApiResponse[models.RecordPaymentResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.RecordPaymentRequest`](../../doc/models/record-payment-request.md) | Body, Optional | - |

## Response Type

[`models.RecordPaymentResponse`](../../doc/models/record-payment-response.md)

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

body := models.RecordPaymentRequest{
    Payment: models.CreatePayment{
        Amount:         "10.0",
        Memo:           "to pay the bills",
        PaymentDetails: "check number 8675309",
        PaymentMethod:  models.InvoicePaymentMethodType("check"),
    },
}

apiResponse, err := invoicesController.RecordPaymentForSubscription(ctx, subscriptionId, &body)
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
  "paid_invoices": [
    {
      "invoice_id": "inv_bchyhr6z5grby",
      "status": "paid",
      "due_amount": "0.0",
      "paid_amount": "50.0"
    },
    {
      "invoice_id": "inv_bchyhrgvyb6vm",
      "status": "paid",
      "due_amount": "0.0",
      "paid_amount": "50.0"
    }
  ],
  "prepayment": null
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Reopen Invoice

This endpoint allows you to reopen any invoice with the "canceled" status. Invoices enter "canceled" status if they were open at the time the subscription was canceled (whether through dunning or an intentional cancellation).

Invoices with "canceled" status are no longer considered to be due. Once reopened, they are considered due for payment. Payment may then be captured in one of the following ways:

- Reactivating the subscription, which will capture all open invoices (See note below about automatic reopening of invoices.)
- Recording a payment directly against the invoice

A note about reactivations: any canceled invoices from the most recent active period are automatically opened as a part of the reactivation process. Reactivating via this endpoint prior to reactivation is only necessary when you wish to capture older invoices from previous periods during the reactivation.

### Reopening Consolidated Invoices

When reopening a consolidated invoice, all of its canceled segments will also be reopened.

```go
ReopenInvoice(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier for the invoice, this does not refer to the public facing invoice number. |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

apiResponse, err := invoicesController.ReopenInvoice(ctx, uid)
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


# Void Invoice

This endpoint allows you to void any invoice with the "open" or "canceled" status.  It will also allow voiding of an invoice with the "pending" status if it is not a consolidated invoice.

```go
VoidInvoice(
    ctx context.Context,
    uid string,
    body *models.VoidInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier for the invoice, this does not refer to the public facing invoice number. |
| `body` | [`*models.VoidInvoiceRequest`](../../doc/models/void-invoice-request.md) | Body, Optional | - |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

body := models.VoidInvoiceRequest{
    Void: models.VoidInvoice{
        Reason: "Duplicate invoice",
    },
}

apiResponse, err := invoicesController.VoidInvoice(ctx, uid, &body)
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


# List Consolidated Invoice Segments

Invoice segments returned on the index will only include totals, not detailed breakdowns for `line_items`, `discounts`, `taxes`, `credits`, `payments`, or `custom_fields`.

```go
ListConsolidatedInvoiceSegments(
    ctx context.Context,
    input ListConsolidatedInvoiceSegmentsInput) (
    models.ApiResponse[models.ConsolidatedInvoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `invoiceUid` | `string` | Template, Required | The unique identifier of the consolidated invoice |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `direction` | [`*models.Direction`](../../doc/models/direction.md) | Query, Optional | Sort direction of the returned segments. |

## Response Type

[`models.ConsolidatedInvoice`](../../doc/models/consolidated-invoice.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListConsolidatedInvoiceSegmentsInput{
    InvoiceUid: "invoice_uid0",
    Page:       models.ToPointer(2),
    PerPage:    models.ToPointer(50),
    Direction:  models.ToPointer(models.Direction("asc")),
}

apiResponse, err := invoicesController.ListConsolidatedInvoiceSegments(ctx, collectedInput)
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
  "invoices": [
    {
      "uid": "inv_8htcd29wcq3q6",
      "site_id": 51288,
      "customer_id": 20153415,
      "subscription_id": 23277588,
      "number": "125",
      "sequence_number": 125,
      "issue_date": "2018-09-20",
      "due_date": "2018-09-20",
      "paid_date": "2018-09-20",
      "status": "paid",
      "collection_method": "automatic",
      "payment_instructions": "Make checks payable to Acme, Inc.",
      "currency": "USD",
      "consolidation_level": "parent",
      "parent_invoice_uid": null,
      "parent_invoice_number": null,
      "group_primary_subscription_id": 23277588,
      "product_name": "Trial and setup fee",
      "product_family_name": "Billing Plans",
      "seller": {
        "name": "General Goods",
        "address": {
          "street": "123 General Goods Way",
          "line2": "Apt. 10",
          "city": "Boston",
          "state": "MA",
          "zip": "02120",
          "country": "US"
        },
        "phone": "555-555-1212"
      },
      "customer": {
        "chargify_id": 20153415,
        "first_name": "Meg",
        "last_name": "Example",
        "organization": "",
        "email": "meg@example.com"
      },
      "memo": "Please pay within 15 days.",
      "billing_address": {
        "street": "123 I Love Cats Way",
        "line2": "",
        "city": "Boston",
        "state": "MA",
        "zip": "90210",
        "country": "US"
      },
      "shipping_address": {
        "street": "123 I Love Cats Way",
        "line2": "",
        "city": "Boston",
        "state": "MA",
        "zip": "90210",
        "country": "US"
      },
      "subtotal_amount": "100.0",
      "discount_amount": "0.0",
      "tax_amount": "0.0",
      "total_amount": "100.0",
      "credit_amount": "0.0",
      "paid_amount": "100.0",
      "refund_amount": "0.0",
      "due_amount": "0.0",
      "public_url": "https://www.chargifypay.com/invoice/inv_8htcd29wcq3q6?token=fb6kpjz5rcr2vttyjs4rcv6y"
    },
    {
      "uid": "inv_8hr3546xp4h8n",
      "site_id": 51288,
      "customer_id": 21687686,
      "subscription_id": 22007644,
      "number": "124",
      "sequence_number": 124,
      "issue_date": "2018-09-18",
      "due_date": "2018-09-18",
      "paid_date": null,
      "status": "open",
      "collection_method": "remittance",
      "payment_instructions": "Make checks payable to Acme, Inc.",
      "currency": "USD",
      "consolidation_level": "none",
      "parent_invoice_uid": null,
      "parent_invoice_number": null,
      "group_primary_subscription_id": null,
      "product_name": "Trial and setup fee",
      "product_family_name": "Billing Plans",
      "seller": {
        "name": "General Goods",
        "address": {
          "street": "123 General Goods Way",
          "line2": "Apt. 10",
          "city": "Boston",
          "state": "MA",
          "zip": "02120",
          "country": "US"
        },
        "phone": "555-555-1212"
      },
      "customer": {
        "chargify_id": 21687686,
        "first_name": "Charlene",
        "last_name": "Tester",
        "organization": "",
        "email": "food@example.com"
      },
      "memo": "Please pay within 15 days.",
      "billing_address": {
        "street": "",
        "line2": "",
        "city": "",
        "state": "",
        "zip": "",
        "country": ""
      },
      "shipping_address": {
        "street": "",
        "line2": "",
        "city": "",
        "state": "",
        "zip": "",
        "country": ""
      },
      "subtotal_amount": "100.0",
      "discount_amount": "0.0",
      "tax_amount": "0.0",
      "total_amount": "100.0",
      "credit_amount": "0.0",
      "paid_amount": "0.0",
      "refund_amount": "0.0",
      "due_amount": "100.0",
      "public_url": "https://www.chargifypay.com/invoice/inv_8hr3546xp4h8n?token=fb6kpjz5rcr2vttyjs4rcv6y"
    },
    {
      "uid": "inv_8hr3546wdwxkr",
      "site_id": 51288,
      "customer_id": 21687670,
      "subscription_id": 22007627,
      "number": "123",
      "sequence_number": 123,
      "issue_date": "2018-09-18",
      "due_date": "2018-09-18",
      "paid_date": "2018-09-18",
      "status": "paid",
      "collection_method": "automatic",
      "payment_instructions": "Make checks payable to Acme, Inc.",
      "currency": "USD",
      "consolidation_level": "none",
      "parent_invoice_uid": null,
      "parent_invoice_number": null,
      "group_primary_subscription_id": null,
      "product_name": "Trial End - Free",
      "product_family_name": "Billing Plans",
      "seller": {
        "name": "General Goods",
        "address": {
          "street": "123 General Goods Way",
          "line2": "Apt. 10",
          "city": "Boston",
          "state": "MA",
          "zip": "02120",
          "country": "US"
        },
        "phone": "555-555-1212"
      },
      "customer": {
        "chargify_id": 21687670,
        "first_name": "Hello",
        "last_name": "World",
        "organization": "123",
        "email": "example@example.com"
      },
      "memo": "Please pay within 15 days.",
      "billing_address": {
        "street": "123 Anywhere Street",
        "line2": "",
        "city": "Boston",
        "state": "MA",
        "zip": "02120",
        "country": "US"
      },
      "shipping_address": {
        "street": "",
        "line2": "",
        "city": "Boston",
        "state": "AL",
        "zip": "02120",
        "country": "US"
      },
      "subtotal_amount": "0.0",
      "discount_amount": "0.0",
      "tax_amount": "0.0",
      "total_amount": "0.0",
      "credit_amount": "0.0",
      "paid_amount": "0.0",
      "refund_amount": "0.0",
      "due_amount": "0.0",
      "public_url": "https://www.chargifypay.com/invoice/inv_8hr3546wdwxkr?token=fb6kpjz5rcr2vttyjs4rcv6y"
    },
    {
      "uid": "inv_8hjtk8bz56bbp",
      "site_id": 51288,
      "customer_id": 20137757,
      "subscription_id": 20541100,
      "number": "122",
      "sequence_number": 122,
      "issue_date": "2018-09-10",
      "due_date": "2018-09-10",
      "paid_date": "2018-09-10",
      "status": "paid",
      "collection_method": "automatic",
      "payment_instructions": "Make checks payable to Acme, Inc.",
      "currency": "USD",
      "consolidation_level": "none",
      "parent_invoice_uid": null,
      "parent_invoice_number": null,
      "group_primary_subscription_id": null,
      "product_name": "$0 Product",
      "product_family_name": "Billing Plans",
      "seller": {
        "name": "General Goods",
        "address": {
          "street": "123 General Goods Way",
          "line2": "Apt. 10",
          "city": "Boston",
          "state": "MA",
          "zip": "02120",
          "country": "US"
        },
        "phone": "555-555-1212"
      },
      "customer": {
        "chargify_id": 20137757,
        "first_name": "Sasha",
        "last_name": "Example",
        "organization": "",
        "email": "example@example.com"
      },
      "memo": "Please pay within 15 days.",
      "billing_address": {
        "street": "123 I Love Cats Way",
        "line2": "",
        "city": "Catville",
        "state": "MA",
        "zip": "90210",
        "country": "US"
      },
      "shipping_address": {
        "street": "123 I Love Cats Way",
        "line2": "",
        "city": "Catville",
        "state": "AL",
        "zip": "90210",
        "country": "US"
      },
      "subtotal_amount": "0.0",
      "discount_amount": "0.0",
      "tax_amount": "0.0",
      "total_amount": "0.0",
      "credit_amount": "0.0",
      "paid_amount": "0.0",
      "refund_amount": "0.0",
      "due_amount": "0.0",
      "public_url": "https://www.chargifypay.com/invoice/inv_8jzrw74xq8kxr?token=fb6kpjz5rcr2vttyjs4rcv6y"
    }
  ]
}
```


# Create Invoice

This endpoint will allow you to create an ad hoc invoice.

### Basic Behavior

You can create a basic invoice by sending an array of line items to this endpoint. Each line item, at a minimum, must include a title, a quantity and a unit price. Example:

```json
{
  "invoice": {
    "line_items": [
      {
        "title": "A Product",
        "quantity": 12,
        "unit_price": "150.00"
      }
    ]
  }
}
```

### Catalog items

Instead of creating custom products like in above example, You can pass existing items like products, components.

```json
{
  "invoice": {
    "line_items": [
      {
        "product_id": "handle:gold-product",
        "quantity": 2,
      }
    ]
  }
}
```

The price for each line item will be calculated as well as a total due amount for the invoice. Multiple line items can be sent.

### Line items types

When defining line item, You can choose one of 3 types for one line item:

#### Custom item

Like in basic behavior example above, You can pass `title` and `unit_price` for custom item.

#### Product id

Product handle (with handle: prefix) or id from the scope of current subscription's site can be provided with `product_id`. By default `unit_price` is taken from product's default price point, but can be overwritten by passing `unit_price` or `product_price_point_id`. If `product_id` is used, following fields cannot be used: `title`, `component_id`.

#### Component id

Component handle (with handle: prefix) or id from the scope of current subscription's site can be provided with `component_id`. If `component_id` is used, following fields cannot be used: `title`, `product_id`. By default `unit_price` is taken from product's default price point, but can be overwritten by passing `unit_price` or `price_point_id`. At this moment price points are supportted only for quantity based, on/off and metered components. For prepaid and event based billing components `unit_price` is required.

### Coupons

When creating ad hoc invoice, new discounts can be applied in following way:

```json
{
  "invoice": {
    "line_items": [
      {
        "product_id": "handle:gold-product",
        "quantity": 1
      }
    ],
    "coupons": [
      {
        "code": "COUPONCODE",
        "percentage": 50.0
      }
    ]
  }
}
```

If You want to use existing coupon for discount creation, only `code` and optional `product_family_id` is needed

```json
...
 "coupons": [
      {
        "code": "FREESETUP",
        "product_family_id": 1
      }
  ]
...
```

### Coupon options

#### Code

Coupon `code` will be displayed on invoice discount section.
Coupon code can only contain uppercase letters, numbers, and allowed special characters.
Lowercase letters will be converted to uppercase. It can be used to select an existing coupon from the catalog, or as an ad hoc coupon when passed with `percentage` or `amount`.

#### Percentage

Coupon `percentage` can take values from 0 to 100 and up to 4 decimal places. It cannot be used with `amount`. Only for ad hoc coupons, will be ignored if `code` is used to select an existing coupon from the catalog.

#### Amount

Coupon `amount` takes number value. It cannot be used with `percentage`. Used only when not matching existing coupon by `code`.

#### Description

Optional `description` will be displayed with coupon `code`. Used only when not matching existing coupon by `code`.

#### Product Family id

Optional `product_family_id` handle (with handle: prefix) or id is used to match existing coupon within site, when codes are not unique.

#### Compounding Strategy

Optional `compounding_strategy` for percentage coupons, can take values `compound` or `full-price`.

For amount coupons, discounts will be always calculated against the original item price, before other discounts are applied.

`compound` strategy:
Percentage-based discounts will be calculated against the remaining price, after prior discounts have been calculated. It is set by default.

`full-price` strategy:
Percentage-based discounts will always be calculated against the original item price, before other discounts are applied.

### Line Item Options

#### Period Date Range

A custom period date range can be defined for each line item with the `period_range_start` and `period_range_end` parameters. Dates must be sent in the `YYYY-MM-DD` format.
`period_range_end` must be greater or equal `period_range_start`.

#### Taxes

The `taxable` parameter can be sent as `true` if taxes should be calculated for a specific line item. For this to work, the site should be configured to use and calculate taxes. Further, if the site uses Avalara for tax calculations, a `tax_code` parameter should also be sent. For existing catalog items: products/components taxes cannot be overwritten.

#### Price Point

Price point handle (with handle: prefix) or id from the scope of current subscription's site can be provided with `price_point_id` for components with `component_id` or `product_price_point_id` for products with `product_id` parameter. If price point is passed `unit_price` cannot be used. It can be used only with catalog items products and components.

#### Description

Optional `description` parameter, it will overwrite default generated description for line item.

### Invoice Options

#### Issue Date

By default, invoices will be created with a issue date set to today. `issue_date` parameter can be send to alter that. Only dates in the past can be send. `issue_date` should be send in `YYYY-MM-DD` format.

#### Net Terms

By default, invoices will be created with a due date matching the date of invoice creation. If a different due date is desired, the `net_terms` parameter can be sent indicating the number of days in advance the due date should be.

#### Addresses

The seller, shipping and billing addresses can be sent to override the site's defaults. Each address requires to send a `first_name` at a minimum in order to work. Please see below for the details on which parameters can be sent for each address object.

#### Memo and Payment Instructions

A custom memo can be sent with the `memo` parameter to override the site's default. Likewise, custom payment instructions can be sent with the `payment_instrucions` parameter.

#### Status

By default, invoices will be created with open status. Possible alternative is `draft`.

```go
CreateInvoice(
    ctx context.Context,
    subscriptionId int,
    body *models.CreateInvoiceRequest) (
    models.ApiResponse[models.InvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription |
| `body` | [`*models.CreateInvoiceRequest`](../../doc/models/create-invoice-request.md) | Body, Optional | - |

## Response Type

[`models.InvoiceResponse`](../../doc/models/invoice-response.md)

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

body := models.CreateInvoiceRequest{
    Invoice: models.CreateInvoice{
        LineItems:           []models.CreateInvoiceItem{
            models.CreateInvoiceItem{
                Title:               models.ToPointer("A Product"),
                Quantity:            models.ToPointer(models.CreateInvoiceItemQuantityContainer.FromPrecision(float64(12))),
                UnitPrice:           models.ToPointer(models.CreateInvoiceItemUnitPriceContainer.FromString("150.00")),
            },
        },
    },
}

apiResponse, err := invoicesController.CreateInvoice(ctx, subscriptionId, &body)
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
  "invoice": {
    "uid": "inv_98nbmb93gxjz8",
    "site_id": 5,
    "customer_id": 13,
    "subscription_id": 17,
    "number": "1",
    "sequence_number": 1,
    "issue_date": "2020-12-02",
    "due_date": "2020-12-02",
    "paid_date": null,
    "status": "open",
    "collection_method": "remittance",
    "payment_instructions": "",
    "currency": "USD",
    "consolidation_level": "none",
    "parent_invoice_uid": null,
    "parent_invoice_number": null,
    "group_primary_subscription_id": null,
    "product_name": "Digitized discrete initiative",
    "product_family_name": "Networked bottom-line orchestration",
    "role": "adhoc",
    "seller": {
      "name": "Steuber, West and Hegmann 1",
      "address": {
        "street": null,
        "line2": null,
        "city": null,
        "state": null,
        "zip": null,
        "country": null
      },
      "phone": "(766) 316-0492"
    },
    "customer": {
      "chargify_id": 13,
      "first_name": "Dean",
      "last_name": "Adams",
      "organization": null,
      "email": "brandi1@example.com",
      "vat_number": null,
      "reference": null
    },
    "memo": "",
    "billing_address": {
      "street": null,
      "line2": null,
      "city": null,
      "state": null,
      "zip": null,
      "country": null
    },
    "shipping_address": {
      "street": null,
      "line2": null,
      "city": null,
      "state": null,
      "zip": null,
      "country": null
    },
    "subtotal_amount": "1800.0",
    "discount_amount": "0.0",
    "tax_amount": "0.0",
    "total_amount": "1800.0",
    "credit_amount": "0.0",
    "paid_amount": "0.0",
    "refund_amount": "0.0",
    "due_amount": "1800.0",
    "line_items": [
      {
        "uid": "li_98nbmb9jgz84c",
        "title": "Some product",
        "description": "12/02/2020 - 12/02/2020",
        "quantity": "12.0",
        "unit_price": "150.0",
        "subtotal_amount": "1800.0",
        "discount_amount": "0.0",
        "tax_amount": "0.0",
        "total_amount": "1800.0",
        "tiered_unit_price": false,
        "period_range_start": "2020-12-02",
        "period_range_end": "2020-12-02",
        "product_id": null,
        "product_version": null,
        "product_price_point_id": null,
        "component_id": null,
        "price_point_id": null
      }
    ],
    "public_url": "https://www.test-chargifypay.com/invoice/inv_98nbmb93gxjz8?token=rmfmwvbdy4xmyw5f29j5gc6x"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorArrayMapResponseException`](../../doc/models/error-array-map-response-exception.md) |


# Send Invoice

This endpoint allows for invoices to be programmatically delivered via email. This endpoint supports the delivery of both ad-hoc and automatically generated invoices. Additionally, this endpoint supports email delivery to direct recipients, carbon-copy (cc) recipients, and blind carbon-copy (bcc) recipients.

Please note that if no recipient email addresses are specified in the request, then the subscription's default email configuration will be used. For example, if `recipient_emails` is left blank, then the invoice will be delivered to the subscription's customer email address.

On success, a 204 no-content response will be returned. Please note that this does not indicate that email(s) have been delivered, but instead indicates that emails have been successfully queued for delivery. If _any_ invalid or malformed email address is found in the request body, the entire request will be rejected and a 422 response will be returned.

```go
SendInvoice(
    ctx context.Context,
    uid string,
    body *models.SendInvoiceRequest) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier for the invoice, this does not refer to the public facing invoice number. |
| `body` | [`*models.SendInvoiceRequest`](../../doc/models/send-invoice-request.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

body := models.SendInvoiceRequest{
    RecipientEmails:    []string{
        "user0@example.com",
    },
    CcRecipientEmails:  []string{
        "user1@example.com",
    },
    BccRecipientEmails: []string{
        "user2@example.com",
    },
}

resp, err := invoicesController.SendInvoice(ctx, uid, &body)
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


# Preview Customer Information Changes

Customer information may change after an invoice is issued which may lead to a mismatch between customer information that are present on an open invoice and actual customer information. This endpoint allows to preview these differences, if any.

The endpoint doesn't accept a request body. Customer information differences are calculated on the application side.

```go
PreviewCustomerInformationChanges(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.CustomerChangesPreviewResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier for the invoice, this does not refer to the public facing invoice number. |

## Response Type

[`models.CustomerChangesPreviewResponse`](../../doc/models/customer-changes-preview-response.md)

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

apiResponse, err := invoicesController.PreviewCustomerInformationChanges(ctx, uid)
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
  "changes": {
    "payer": {
      "before": {
        "last_name": "Beatty"
      },
      "after": {
        "last_name": "Doe"
      }
    },
    "shipping_address": {
      "before": {
        "line2": "Suite 703"
      },
      "after": {
        "line2": "Suite 702"
      }
    },
    "billing_address": {
      "before": {
        "line2": "Suite 703"
      },
      "after": {
        "line2": "Suite 702"
      }
    },
    "custom_fields": {
      "before": [
        {
          "owner_id": 1002,
          "owner_type": "Customer",
          "name": "Color",
          "value": "blue",
          "metadatum_id": 20
        }
      ],
      "after": [
        {
          "owner_id": 1002,
          "owner_type": "Customer",
          "name": "Color",
          "value": "green",
          "metadatum_id": 20
        }
      ]
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Update Customer Information

This endpoint updates customer information on an open invoice and returns the updated invoice. If you would like to preview changes that will be applied, use the `/invoices/{uid}/customer_information/preview.json` endpoint before.

The endpoint doesn't accept a request body. Customer information differences are calculated on the application side.

```go
UpdateCustomerInformation(
    ctx context.Context,
    uid string) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier for the invoice, this does not refer to the public facing invoice number. |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

apiResponse, err := invoicesController.UpdateCustomerInformation(ctx, uid)
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
  "uid": "elit Ut",
  "site_id": 46283786,
  "customer_id": -62349460,
  "subscription_id": 12801726,
  "number": "dolore et ut",
  "sequence_number": -84210096,
  "issue_date": "2017-01-01",
  "due_date": "2017-01-30",
  "paid_date": "2017-01-28",
  "status": "open",
  "collection_method": "automatic",
  "payment_instructions": "enim officia",
  "currency": "dolore",
  "consolidation_level": "none",
  "product_name": "occaecat veniam culpa",
  "product_family_name": "qui commodo ea dolore cillum",
  "seller": {
    "name": "co",
    "phone": "ullamco in officia"
  },
  "customer": {
    "chargify_id": -55826334,
    "first_name": "deserunt",
    "last_name": "velit dolore",
    "email": "aliquip sed velit Lorem"
  },
  "memo": "ea cupidatat deserunt",
  "billing_address": {
    "street": "qui commodo cupidatat sunt",
    "line2": "ut officia enim",
    "city": "velit minim dolore sint nulla",
    "state": "velit",
    "zip": "ullamco",
    "country": "irure est laborum deserun"
  },
  "shipping_address": {
    "street": "do fugiat dolore deserunt officia",
    "line2": "ipsum cillum",
    "city": "aliqua laboris incididunt ut",
    "state": "et fugiat sit",
    "zip": "dolore do",
    "country": "Excepteur consequat cillum"
  },
  "subtotal_amount": "dolore mollit",
  "discount_amount": "aute",
  "tax_amount": "eu aliqua est velit ea",
  "total_amount": "ut non",
  "credit_amount": "sit",
  "refund_amount": "et eiusmod qui sed",
  "paid_amount": "amet nulla s",
  "due_amount": "non esse ullamco",
  "line_items": [
    {
      "description": "qui",
      "price_point_id": 123,
      "tax_amount": "occaecat deserunt veniam",
      "subtotal_amount": "commodo consequat tempor et Duis"
    },
    {
      "uid": "",
      "subtotal_amount": "ven"
    },
    {
      "price_point_id": 94750853,
      "product_id": 79058036,
      "tax_amount": "1.0",
      "subtotal_amount": "128.5"
    },
    {
      "unit_price": "eiusmod consequat ut nostrud",
      "tax_amount": "quis nulla proident"
    },
    {
      "period_range_end": "2022-02-02",
      "product_id": 57352537,
      "description": "minim in dolore Ut Excepteur",
      "uid": "sit qui in ullamco anim"
    }
  ],
  "discounts": [
    {
      "title": "nostrud"
    }
  ],
  "taxes": [
    {
      "source_type": "Tax",
      "line_item_breakouts": [
        {
          "uid": "in ipsum",
          "tax_amount": "velit",
          "taxable_amount": "quis sint"
        },
        {
          "uid": "co"
        }
      ]
    },
    {
      "uid": "enim irure in",
      "title": "incididunt est mollit irure"
    }
  ],
  "credits": [
    {
      "uid": "exercitation eiusmod",
      "transaction_time": "2024-01-23T13:51:27Z",
      "credit_note_number": "qui fugiat labore laborum",
      "credit_note_uid": "ipsum sunt"
    },
    {
      "memo": "dolor"
    }
  ],
  "refunds": [
    {
      "memo": "deserunt elit"
    },
    {
      "original_amount": "Duis nulla"
    }
  ],
  "payments": [
    {
      "prepayment": false,
      "memo": "enim Excepteur Lorem magna sit"
    },
    {
      "transaction_time": "2024-01-23T13:51:27Z",
      "prepayment": false,
      "payment_method": {
        "details": "labore ut et",
        "kind": "dolor qui",
        "memo": "ea commodo",
        "type": "fugiat veniam",
        "card_brand": "consequat",
        "card_expiration": "aliqua a",
        "last_four": "ut in consectetur sed",
        "masked_card_number": "minim ea ullamco nostrud tempor"
      }
    },
    {
      "prepayment": true,
      "transaction_id": 67527234
    },
    {
      "original_amount": "c",
      "memo": "dolore fugiat labore"
    }
  ],
  "custom_fields": [
    {
      "name": "CustomerStatus",
      "value": "Gold",
      "owner_type": "Customer",
      "owner_id": 18482224,
      "metadatum_id": 13924
    },
    {
      "name": "SubscriptionTag",
      "value": "Special Subscriber",
      "owner_type": "Subscription",
      "owner_id": 21344,
      "metadatum_id": 139245
    }
  ],
  "public_url": "dolo",
  "previous_balance_data": {
    "captured_at": "2024-01-09T11:22:23-05:00",
    "invoices": [
      {
        "number": "veniam dolore labore ipsum cupidatat",
        "uid": "tempor",
        "outstanding_amount": "Excepteur nostrud irur"
      },
      {
        "outstanding_amount": "id"
      }
    ]
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Issue Invoice

This endpoint allows you to issue an invoice that is in "pending" status. For example, you can issue an invoice that was created when allocating new quantity on a component and using "accrue charges" option.

You cannot issue a pending child invoice that was created for a member subscription in a group.

For Remittance subscriptions, the invoice will go into "open" status and payment won't be attempted. The value for `on_failed_payment` would be rejected if sent. Any prepayments or service credits that exist on subscription will be automatically applied. Additionally, if setting is on, an email will be sent for issued invoice.

For Automatic subscriptions, prepayments and service credits will apply to the invoice and before payment is attempted. On successful payment, the invoice will go into "paid" status and email will be sent to the customer (if setting applies). When payment fails, the next event depends on the `on_failed_payment` value:

- `leave_open_invoice` - prepayments and credits applied to invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history. This is the default option.
- `rollback_to_pending` - prepayments and credits not applied; invoice remains in "pending" status; no email sent to the customer; payment failure recorded in the invoice history.
- `initiate_dunning` - prepayments and credits applied to the invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history; subscription will  most likely go into "past_due" or "canceled" state (depending upon net terms and dunning settings).

```go
IssueInvoice(
    ctx context.Context,
    uid string,
    body *models.IssueInvoiceRequest) (
    models.ApiResponse[models.Invoice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `uid` | `string` | Template, Required | The unique identifier for the invoice, this does not refer to the public facing invoice number. |
| `body` | [`*models.IssueInvoiceRequest`](../../doc/models/issue-invoice-request.md) | Body, Optional | - |

## Response Type

[`models.Invoice`](../../doc/models/invoice.md)

## Example Usage

```go
ctx := context.Background()

uid := "uid0"

body := models.IssueInvoiceRequest{
    OnFailedPayment: models.ToPointer(models.FailedPaymentAction("leave_open_invoice")),
}

apiResponse, err := invoicesController.IssueInvoice(ctx, uid, &body)
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

