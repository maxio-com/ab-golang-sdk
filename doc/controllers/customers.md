# Customers

```go
customersController := client.CustomersController()
```

## Class Name

`CustomersController`

## Methods

* [Create Customer](../../doc/controllers/customers.md#create-customer)
* [List Customers](../../doc/controllers/customers.md#list-customers)
* [Read Customer](../../doc/controllers/customers.md#read-customer)
* [Update Customer](../../doc/controllers/customers.md#update-customer)
* [Delete Customer](../../doc/controllers/customers.md#delete-customer)
* [Read Customer by Reference](../../doc/controllers/customers.md#read-customer-by-reference)
* [List Customer Subscriptions](../../doc/controllers/customers.md#list-customer-subscriptions)


# Create Customer

You may create a new Customer at any time, or you may create a Customer at the same time you create a Subscription. The only validation restriction is that you may only create one customer for a given reference value.

If provided, the `reference` value must be unique. It represents a unique identifier for the customer from your own app, i.e. the customer’s ID. This allows you to retrieve a given customer via a piece of shared information. Alternatively, you may choose to leave `reference` blank, and store Chargify’s unique ID for the customer, which is in the `id` attribute.

Full documentation on how to locate, create and edit Customers in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407659914267).

## Required Country Format

Chargify requires that you use the ISO Standard Country codes when formatting country attribute of the customer.

Countries should be formatted as 2 characters. For more information, please see the following wikipedia article on [ISO_3166-1.](http://en.wikipedia.org/wiki/ISO_3166-1#Current_codes)

## Required State Format

Chargify requires that you use the ISO Standard State codes when formatting state attribute of the customer.

+ US States (2 characters): [ISO_3166-2](https://en.wikipedia.org/wiki/ISO_3166-2:US)

+ States Outside the US (2-3 characters): To find the correct state codes outside of the US, please go to [ISO_3166-1](http://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) and click on the link in the “ISO 3166-2 codes” column next to country you wish to populate.

## Locale

Chargify allows you to attribute a language/region to your customer to deliver invoices in any required language.
For more: [Customer Locale](https://chargify.zendesk.com/hc/en-us/articles/4407870384283#customer-locale)

```go
CreateCustomer(
    ctx context.Context,
    body *models.CreateCustomerRequest) (
    models.ApiResponse[models.CustomerResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.CreateCustomerRequest`](../../doc/models/create-customer-request.md) | Body, Optional | - |

## Response Type

[`models.CustomerResponse`](../../doc/models/customer-response.md)

## Example Usage

```go
ctx := context.Background()

body := models.CreateCustomerRequest{
    Customer: models.CreateCustomer{
        FirstName:       "Martha",
        LastName:        "Washington",
        Email:           "martha@example.com",
        CcEmails:        models.ToPointer("george@example.com"),
        Organization:    models.ToPointer("ABC, Inc."),
        Reference:       models.ToPointer("1234567890"),
        Address:         models.ToPointer("123 Main Street"),
        Address2:        models.ToPointer("Unit 10"),
        City:            models.ToPointer("Anytown"),
        State:           models.ToPointer("MA"),
        Zip:             models.ToPointer("02120"),
        Country:         models.ToPointer("US"),
        Phone:           models.ToPointer("555-555-1212"),
        Locale:          models.ToPointer("es-MX"),
    },
}

apiResponse, err := customersController.CreateCustomer(ctx, &body)
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
  "customer": {
    "first_name": "Cathryn",
    "last_name": "Parisian",
    "email": "Stella.McLaughlin6@example.net",
    "cc_emails": null,
    "organization": "Greenholt - Oberbrunner",
    "reference": null,
    "id": 76,
    "created_at": "2021-03-29T07:47:00-04:00",
    "updated_at": "2021-03-29T07:47:00-04:00",
    "address": "739 Stephon Bypass",
    "address_2": "Apt. 386",
    "city": "Sedrickchester",
    "state": "KY",
    "state_name": "Kentucky",
    "zip": "46979-7719",
    "country": "US",
    "country_name": "United States",
    "phone": "230-934-3685",
    "verified": false,
    "portal_customer_created_at": null,
    "portal_invite_last_sent_at": null,
    "portal_invite_last_accepted_at": null,
    "tax_exempt": false,
    "vat_number": null,
    "parent_id": null,
    "locale": "en-US"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`CustomerErrorResponseException`](../../doc/models/customer-error-response-exception.md) |


# List Customers

This request will by default list all customers associated with your Site.

## Find Customer

Use the search feature with the `q` query parameter to retrieve an array of customers that matches the search query.

Common use cases are:

+ Search by an email
+ Search by a Chargify ID
+ Search by an organization
+ Search by a reference value from your application
+ Search by a first or last name

To retrieve a single, exact match by reference, please use the [lookup endpoint](https://developers.chargify.com/docs/api-docs/b710d8fbef104-read-customer-by-reference).

```go
ListCustomers(
    ctx context.Context,
    input ListCustomersInput) (
    models.ApiResponse[[]models.CustomerResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Query, Optional | Direction to sort customers by time of creation |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 50. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `dateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Query, Optional | The type of filter you would like to apply to your search.<br>Use in query: `date_field=created_at`. |
| `startDate` | `*string` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `endDate` | `*string` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `startDatetime` | `*string` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `endDatetime` | `*string` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |
| `q` | `*string` | Query, Optional | A search query by which to filter customers (can be an email, an ID, a reference, organization) |

## Response Type

[`[]models.CustomerResponse`](../../doc/models/customer-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListCustomersInput{
    Page:          models.ToPointer(2),
    PerPage:       models.ToPointer(30),
    DateField:     models.ToPointer(models.BasicDateField("updated_at")),
}

apiResponse, err := customersController.ListCustomers(ctx, collectedInput)
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
    "customer": {
      "first_name": "Kayla",
      "last_name": "Test",
      "email": "kayla@example.com",
      "cc_emails": "john@example.com, sue@example.com",
      "organization": "",
      "reference": null,
      "id": 14126091,
      "created_at": "2016-10-04T15:22:27-04:00",
      "updated_at": "2016-10-04T15:22:30-04:00",
      "address": "",
      "address_2": "",
      "city": "",
      "state": "",
      "zip": "",
      "country": "",
      "phone": "",
      "verified": null,
      "portal_customer_created_at": "2016-10-04T15:22:29-04:00",
      "portal_invite_last_sent_at": "2016-10-04T15:22:30-04:00",
      "portal_invite_last_accepted_at": null,
      "tax_exempt": false
    }
  },
  {
    "customer": {
      "first_name": "Nick ",
      "last_name": "Test",
      "email": "nick@example.com",
      "cc_emails": "john@example.com, sue@example.com",
      "organization": "",
      "reference": null,
      "id": 14254093,
      "created_at": "2016-10-13T16:52:51-04:00",
      "updated_at": "2016-10-13T16:52:54-04:00",
      "address": "",
      "address_2": "",
      "city": "",
      "state": "",
      "zip": "",
      "country": "",
      "phone": "",
      "verified": null,
      "portal_customer_created_at": "2016-10-13T16:52:54-04:00",
      "portal_invite_last_sent_at": "2016-10-13T16:52:54-04:00",
      "portal_invite_last_accepted_at": null,
      "tax_exempt": false,
      "parent_id": 123
    }
  },
  {
    "customer": {
      "first_name": "Don",
      "last_name": "Test",
      "email": "don@example.com",
      "cc_emails": "john@example.com, sue@example.com",
      "organization": "",
      "reference": null,
      "id": 14332342,
      "created_at": "2016-10-19T10:49:13-04:00",
      "updated_at": "2016-10-19T10:49:19-04:00",
      "address": "1737 15th St",
      "address_2": "",
      "city": "Boulder",
      "state": "CO",
      "zip": "80302",
      "country": "US",
      "phone": "",
      "verified": null,
      "portal_customer_created_at": "2016-10-19T10:49:19-04:00",
      "portal_invite_last_sent_at": "2016-10-19T10:49:19-04:00",
      "portal_invite_last_accepted_at": null,
      "tax_exempt": false,
      "parent_id": null
    }
  }
]
```


# Read Customer

This method allows to retrieve the Customer properties by Chargify-generated Customer ID.

```go
ReadCustomer(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.CustomerResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `int` | Template, Required | The Chargify id of the customer |

## Response Type

[`models.CustomerResponse`](../../doc/models/customer-response.md)

## Example Usage

```go
ctx := context.Background()

id := 112

apiResponse, err := customersController.ReadCustomer(ctx, id)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Customer

This method allows to update the Customer.

```go
UpdateCustomer(
    ctx context.Context,
    id int,
    body *models.UpdateCustomerRequest) (
    models.ApiResponse[models.CustomerResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `int` | Template, Required | The Chargify id of the customer |
| `body` | [`*models.UpdateCustomerRequest`](../../doc/models/update-customer-request.md) | Body, Optional | - |

## Response Type

[`models.CustomerResponse`](../../doc/models/customer-response.md)

## Example Usage

```go
ctx := context.Background()

id := 112

body := models.UpdateCustomerRequest{
    Customer: models.UpdateCustomer{
        FirstName:       models.ToPointer("Martha"),
        LastName:        models.ToPointer("Washington"),
        Email:           models.ToPointer("martha.washington@example.com"),
    },
}

apiResponse, err := customersController.UpdateCustomer(ctx, id, &body)
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
  "customer": {
    "first_name": "Martha",
    "last_name": "Washington",
    "email": "martha.washington@example.com",
    "cc_emails": "george.washington@example.com",
    "organization": null,
    "reference": null,
    "id": 14967442,
    "created_at": "2016-12-05T10:33:07-05:00",
    "updated_at": "2016-12-05T10:38:00-05:00",
    "address": null,
    "address_2": null,
    "city": null,
    "state": null,
    "zip": null,
    "country": null,
    "phone": null,
    "verified": false,
    "portal_customer_created_at": null,
    "portal_invite_last_sent_at": null,
    "portal_invite_last_accepted_at": null,
    "tax_exempt": false,
    "vat_number": "012345678"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`CustomerErrorResponseException`](../../doc/models/customer-error-response-exception.md) |


# Delete Customer

This method allows you to delete the Customer.

```go
DeleteCustomer(
    ctx context.Context,
    id int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `int` | Template, Required | The Chargify id of the customer |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

id := 112

resp, err := customersController.DeleteCustomer(ctx, id)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Read Customer by Reference

Use this method to return the customer object if you have the unique **Reference ID (Your App)** value handy. It will return a single match.

```go
ReadCustomerByReference(
    ctx context.Context,
    reference string) (
    models.ApiResponse[models.CustomerResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `reference` | `string` | Query, Required | Customer reference |

## Response Type

[`models.CustomerResponse`](../../doc/models/customer-response.md)

## Example Usage

```go
ctx := context.Background()

reference := "reference4"

apiResponse, err := customersController.ReadCustomerByReference(ctx, reference)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# List Customer Subscriptions

This method lists all subscriptions that belong to a customer.

```go
ListCustomerSubscriptions(
    ctx context.Context,
    customerId int) (
    models.ApiResponse[[]models.SubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `int` | Template, Required | The Chargify id of the customer |

## Response Type

[`[]models.SubscriptionResponse`](../../doc/models/subscription-response.md)

## Example Usage

```go
ctx := context.Background()

customerId := 150

apiResponse, err := customersController.ListCustomerSubscriptions(ctx, customerId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

