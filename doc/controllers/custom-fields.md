# Custom Fields

```go
customFieldsController := client.CustomFieldsController()
```

## Class Name

`CustomFieldsController`

## Methods

* [Create Metafields](../../doc/controllers/custom-fields.md#create-metafields)
* [List Metafields](../../doc/controllers/custom-fields.md#list-metafields)
* [Update Metafield](../../doc/controllers/custom-fields.md#update-metafield)
* [Delete Metafield](../../doc/controllers/custom-fields.md#delete-metafield)
* [Create Metadata](../../doc/controllers/custom-fields.md#create-metadata)
* [List Metadata](../../doc/controllers/custom-fields.md#list-metadata)
* [Update Metadata](../../doc/controllers/custom-fields.md#update-metadata)
* [Delete Metadata](../../doc/controllers/custom-fields.md#delete-metadata)
* [List Metadata for Resource Type](../../doc/controllers/custom-fields.md#list-metadata-for-resource-type)


# Create Metafields

## Custom Fields: Metafield Intro

**Advanced Billing refers to Custom Fields in the API documentation as metafields and metadata.** Within the Advanced Billing UI, metadata and metafields are grouped together under the umbrella of "Custom Fields." All of our UI-based documentation that references custom fields will not cite the terminology metafields or metadata.

+ **Metafield is the custom field**
+ **Metadata is the data populating the custom field.**

Advanced Billing Metafields are used to add meaningful attributes to subscription and customer resources. Full documentation on how to create Custom Fields in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/sections/24266118312589-Custom-Fields). For additional documentation on how to record data within custom fields, please see our subscription-based documentation [here](https://maxio.zendesk.com/hc/en-us/articles/24251701302925-Subscription-Summary-Custom-Fields-Tab).

Metafield are the place where you will set up your resource to accept additional data. It is scoped to the site instead of a specific customer or subscription. Think of it as the key, and Metadata as the value on every record.

## Create Metafields

Use this endpoint to create metafields for your Site. Metafields can be populated with metadata after the fact.

Each site is limited to 100 unique Metafields (i.e. keys, or names) per resource. This means you can have 100 Metafields for Subscription and another 100 for Customer.

### Metafields "On-the-Fly"

It is possible to create Metafields “on the fly” when you create your Metadata – if a non-existant name is passed when creating Metadata, a Metafield for that key will be automatically created. The Metafield API, however, gives you more control over your “keys”.

### Metafield Scope Warning

If configuring metafields in the Admin UI or via the API, be careful sending updates to metafields with the scope attribute – **if a partial update is sent it will overwrite the current configuration**.

```go
CreateMetafields(
    ctx context.Context,
    resourceType models.ResourceType,
    body *models.CreateMetafieldsRequest) (
    models.ApiResponse[[]models.Metafield],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | the resource type to which the metafields belong |
| `body` | [`*models.CreateMetafieldsRequest`](../../doc/models/create-metafields-request.md) | Body, Optional | - |

## Response Type

[`[]models.Metafield`](../../doc/models/metafield.md)

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType("subscriptions")

body := models.CreateMetafieldsRequest{
    Metafields: models.CreateMetafieldsRequestMetafieldsContainer.FromCreateMetafield(models.CreateMetafield{
        Name:      models.ToPointer("Dropdown field"),
        Scope:     models.ToPointer(models.MetafieldScope{
            PublicShow: models.ToPointer(models.IncludeOption("1")),
            PublicEdit: models.ToPointer(models.IncludeOption("1")),
        }),
        InputType: models.ToPointer(models.MetafieldInput("dropdown")),
        Enum:      []string{
            "option 1",
            "option 2",
        },
    }),
}

apiResponse, err := customFieldsController.CreateMetafields(ctx, resourceType, &body)
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
    "name": "Color",
    "scope": {
      "csv": "0",
      "statements": "0",
      "invoices": "0",
      "portal": "0"
    },
    "data_count": 0,
    "input_type": "text",
    "enum": null
  },
  {
    "name": "Brand",
    "scope": {
      "csv": "0",
      "statements": "0",
      "invoices": "0",
      "portal": "0"
    },
    "data_count": 0,
    "input_type": "text",
    "enum": null
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`SingleErrorResponseException`](../../doc/models/single-error-response-exception.md) |


# List Metafields

This endpoint lists metafields associated with a site. The metafield description and usage is contained in the response.

```go
ListMetafields(
    ctx context.Context,
    input ListMetafieldsInput) (
    models.ApiResponse[models.ListMetafieldsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | the resource type to which the metafields belong |
| `name` | `*string` | Query, Optional | filter by the name of the metafield |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Query, Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |

## Response Type

[`models.ListMetafieldsResponse`](../../doc/models/list-metafields-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListMetafieldsInput{
    ResourceType: models.ResourceType("subscriptions"),
    Page:         models.ToPointer(2),
    PerPage:      models.ToPointer(50),
}

apiResponse, err := customFieldsController.ListMetafields(ctx, collectedInput)
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
  "total_count": 0,
  "current_page": 0,
  "total_pages": 0,
  "per_page": 0,
  "metafields": [
    {
      "id": 0,
      "name": "string",
      "scope": {
        "csv": "0",
        "statements": "0",
        "invoices": "0",
        "portal": "0",
        "public_show": "0",
        "public_edit": "0"
      },
      "data_count": 0,
      "input_type": "text",
      "enum": null
    }
  ]
}
```


# Update Metafield

Use the following method to update metafields for your Site. Metafields can be populated with metadata after the fact.

```go
UpdateMetafield(
    ctx context.Context,
    resourceType models.ResourceType,
    body *models.UpdateMetafieldsRequest) (
    models.ApiResponse[[]models.Metafield],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | the resource type to which the metafields belong |
| `body` | [`*models.UpdateMetafieldsRequest`](../../doc/models/update-metafields-request.md) | Body, Optional | - |

## Response Type

[`[]models.Metafield`](../../doc/models/metafield.md)

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType("subscriptions")



apiResponse, err := customFieldsController.UpdateMetafield(ctx, resourceType, nil)
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
| 422 | Unprocessable Entity (WebDAV) | [`SingleErrorResponseException`](../../doc/models/single-error-response-exception.md) |


# Delete Metafield

Use the following method to delete a metafield. This will remove the metafield from the Site.

Additionally, this will remove the metafield and associated metadata with all Subscriptions on the Site.

```go
DeleteMetafield(
    ctx context.Context,
    resourceType models.ResourceType,
    name *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | the resource type to which the metafields belong |
| `name` | `*string` | Query, Optional | The name of the metafield to be deleted |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType("subscriptions")



resp, err := customFieldsController.DeleteMetafield(ctx, resourceType, nil)
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


# Create Metadata

## Custom Fields: Metadata Intro

**Advanced Billing refers to Custom Fields in the API documentation as metafields and metadata.** Within the Advanced Billing UI, metadata and metafields are grouped together under the umbrella of "Custom Fields." All of our UI-based documentation that references custom fields will not cite the terminology metafields or metadata.

+ **Metafield is the custom field**
+ **Metadata is the data populating the custom field.**

Advanced Billing Metafields are used to add meaningful attributes to subscription and customer resources. Full documentation on how to create Custom Fields in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24266164865677-Custom-Fields-Overview). For additional documentation on how to record data within custom fields, please see our subscription-based documentation [here.](https://maxio.zendesk.com/hc/en-us/articles/24251701302925-Subscription-Summary-Custom-Fields-Tab)

Metadata is associated to a customer or subscription, and corresponds to a Metafield. When creating a new metadata object for a given record, **if the metafield is not present it will be created**.

## Metadata limits

Metadata values are limited to 2kB in size. Additonally, there are limits on the number of unique metafields available per resource.

## Create Metadata

This method will create a metafield for the site on the fly if it does not already exist, and populate the metadata value.

### Subscription or Customer Resource

Please pay special attention to the resource you use when creating metadata.

```go
CreateMetadata(
    ctx context.Context,
    resourceType models.ResourceType,
    resourceId int,
    body *models.CreateMetadataRequest) (
    models.ApiResponse[[]models.Metadata],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | the resource type to which the metafields belong |
| `resourceId` | `int` | Template, Required | The Advanced Billing id of the customer or the subscription for which the metadata applies |
| `body` | [`*models.CreateMetadataRequest`](../../doc/models/create-metadata-request.md) | Body, Optional | - |

## Response Type

[`[]models.Metadata`](../../doc/models/metadata.md)

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType("subscriptions")

resourceId := 60

body := models.CreateMetadataRequest{
    Metadata: []models.CreateMetadata{
        models.CreateMetadata{
            Name:  models.ToPointer("Color"),
            Value: models.ToPointer("Blue"),
        },
        models.CreateMetadata{
            Name:  models.ToPointer("Something"),
            Value: models.ToPointer("Useful"),
        },
    },
}

apiResponse, err := customFieldsController.CreateMetadata(ctx, resourceType, resourceId, &body)
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
| 422 | Unprocessable Entity (WebDAV) | [`SingleErrorResponseException`](../../doc/models/single-error-response-exception.md) |


# List Metadata

This request will list all of the metadata belonging to a particular resource (ie. subscription, customer) that is specified.

## Metadata Data

This endpoint will also display the current stats of your metadata to use as a tool for pagination.

```go
ListMetadata(
    ctx context.Context,
    input ListMetadataInput) (
    models.ApiResponse[models.PaginatedMetadata],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | the resource type to which the metafields belong |
| `resourceId` | `int` | Template, Required | The Advanced Billing id of the customer or the subscription for which the metadata applies |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |

## Response Type

[`models.PaginatedMetadata`](../../doc/models/paginated-metadata.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListMetadataInput{
    ResourceType: models.ResourceType("subscriptions"),
    ResourceId:   60,
    Page:         models.ToPointer(2),
    PerPage:      models.ToPointer(50),
}

apiResponse, err := customFieldsController.ListMetadata(ctx, collectedInput)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Metadata

This method allows you to update the existing metadata associated with a subscription or customer.

```go
UpdateMetadata(
    ctx context.Context,
    resourceType models.ResourceType,
    resourceId int,
    body *models.UpdateMetadataRequest) (
    models.ApiResponse[[]models.Metadata],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | the resource type to which the metafields belong |
| `resourceId` | `int` | Template, Required | The Advanced Billing id of the customer or the subscription for which the metadata applies |
| `body` | [`*models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Optional | - |

## Response Type

[`[]models.Metadata`](../../doc/models/metadata.md)

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType("subscriptions")

resourceId := 60



apiResponse, err := customFieldsController.UpdateMetadata(ctx, resourceType, resourceId, nil)
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
| 422 | Unprocessable Entity (WebDAV) | [`SingleErrorResponseException`](../../doc/models/single-error-response-exception.md) |


# Delete Metadata

This method removes the metadata from the subscriber/customer cited.

## Query String Usage

For instance if you wanted to delete the metadata for customer 99 named weight you would request:

```
https://acme.chargify.com/customers/99/metadata.json?name=weight
```

If you want to delete multiple metadata fields for a customer 99 named: `weight` and `age` you wrould request:

```
https://acme.chargify.com/customers/99/metadata.json?names[]=weight&names[]=age
```

## Successful Response

For a success, there will be a code `200` and the plain text response `true`.

## Unsuccessful Response

When a failed response is encountered, you will receive a `404` response and the plain text response of `true`.

```go
DeleteMetadata(
    ctx context.Context,
    resourceType models.ResourceType,
    resourceId int,
    name *string,
    names []string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | the resource type to which the metafields belong |
| `resourceId` | `int` | Template, Required | The Advanced Billing id of the customer or the subscription for which the metadata applies |
| `name` | `*string` | Query, Optional | Name of field to be removed. |
| `names` | `[]string` | Query, Optional | Names of fields to be removed. Use in query: `names[]=field1&names[]=my-field&names[]=another-field`. |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType("subscriptions")

resourceId := 60





resp, err := customFieldsController.DeleteMetadata(ctx, resourceType, resourceId, nil, nil)
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


# List Metadata for Resource Type

This method will provide you information on usage of metadata across your selected resource (ie. subscriptions, customers)

## Metadata Data

This endpoint will also display the current stats of your metadata to use as a tool for pagination.

### Metadata for multiple records

`https://acme.chargify.com/subscriptions/metadata.json?resource_ids[]=1&resource_ids[]=2`

## Read Metadata for a Site

This endpoint will list the number of pages of metadata information that are contained within a site.

```go
ListMetadataForResourceType(
    ctx context.Context,
    input ListMetadataForResourceTypeInput) (
    models.ApiResponse[models.PaginatedMetadata],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | the resource type to which the metafields belong |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `dateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Query, Optional | The type of filter you would like to apply to your search. |
| `startDate` | `*time.Time` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns metadata with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `endDate` | `*time.Time` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns metadata with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `startDatetime` | `*time.Time` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns metadata with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `endDatetime` | `*time.Time` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns metadata with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |
| `withDeleted` | `*bool` | Query, Optional | Allow to fetch deleted metadata. |
| `resourceIds` | `[]int` | Query, Optional | Allow to fetch metadata for multiple records based on provided ids. Use in query: `resource_ids[]=122&resource_ids[]=123&resource_ids[]=124`. |
| `direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Query, Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |

## Response Type

[`models.PaginatedMetadata`](../../doc/models/paginated-metadata.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListMetadataForResourceTypeInput{
    ResourceType:  models.ResourceType("subscriptions"),
    Page:          models.ToPointer(2),
    PerPage:       models.ToPointer(50),
    DateField:     models.ToPointer(models.BasicDateField("updated_at")),
}

apiResponse, err := customFieldsController.ListMetadataForResourceType(ctx, collectedInput)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

