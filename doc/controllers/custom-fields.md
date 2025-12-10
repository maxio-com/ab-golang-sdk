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

Creates metafields on a Site for either the Subscriptions or Customers resource.

Metafields and their metadata are created in the Custom Fields configuration page on your Site. Metafields can be populated with metadata when you create them or later with the [Update Metafield](../../doc/controllers/custom-fields.md#update-metafield), [Create Metadata](../../doc/controllers/custom-fields.md#create-metadata), or [Update Metadata](../../doc/controllers/custom-fields.md#update-metadata) endpoints. The Create Metadata and Update Metadata endpoints allow you to add metafields and metadata values to a specific subscription or customer.

Each site is limited to 100 unique metafields per resource. This means you can have 100 metafields for Subscriptions and another 100 for Customers.

> Note: After creating a metafield, the resource type cannot be modified.

In the UI and product documentation, metafields and metadata are called Custom Fields.

- Metafield is the custom field
- Metadata is the data populating the custom field.

See [Custom Fields Reference](https://docs.maxio.com/hc/en-us/articles/24266140850573-Custom-Fields-Reference) and [Custom Fields Tab](https://maxio.zendesk.com/hc/en-us/articles/24251701302925-Subscription-Summary-Custom-Fields-Tab) for information on using Custom Fields in the Advanced Billing UI.

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
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | The resource type to which the metafields belong. |
| `body` | [`*models.CreateMetafieldsRequest`](../../doc/models/create-metafields-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Metafield](../../doc/models/metafield.md).

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType_SUBSCRIPTIONS

body := models.CreateMetafieldsRequest{
    Metafields:           models.CreateMetafieldsRequestMetafieldsContainer.FromCreateMetafield(models.CreateMetafield{
        Name:                 models.ToPointer("Dropdown field"),
        Scope:                models.ToPointer(models.MetafieldScope{
            Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
            Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
            Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
            Portal:               models.ToPointer(models.IncludeOption_INCLUDE),
        }),
        InputType:            models.ToPointer(models.MetafieldInput_DROPDOWN),
        Enum:                 []string{
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

Lists the metafields and their associated details for a Site and resource type. You can filter the request to a specific metafield.

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
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | The resource type to which the metafields belong. |
| `name` | `*string` | Query, Optional | Filter by the name of the metafield. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Query, Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ListMetafieldsResponse](../../doc/models/list-metafields-response.md).

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListMetafieldsInput{
    ResourceType: models.ResourceType_SUBSCRIPTIONS,
    Page:         models.ToPointer(1),
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
  "total_count": 1,
  "current_page": 1,
  "total_pages": 0,
  "per_page": 50,
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

Updates metafields on your Site for a resource type.  Depending on the request structure, you can update or add metafields and metadata to the Subscriptions or Customers resource.

With this endpoint, you can:

- Add metafields. If the metafield specified in current_name does not exist, a new metafield is added.
  
  > Note: Each site is limited to 100 unique metafields per resource. This means you can have 100 metafields for Subscriptions and another 100 for Customers.

- Change the name of a metafield.
  
  > Note: To keep the metafield name the same and only update the metadata for the metafield, you must use the current metafield name in both the `current_name` and `name` parameters.

- Change the input type for the metafield. For example, you can change a metafield input type from text to a dropdown. If you change the input type from text to a dropdown or radio, you must update the specific subscriptions or customers where the metafield was used to reflect the updated metafield and metadata.

- Add metadata values to the existing metadata for a dropdown or radio metafield.
  
  > Note: Updates to metadata overwrite. To add one or more values, you must specify all metadata values including the new value you want to add.

- Add new metadata to a dropdown or radio for a metafield that was created without metadata.

- Remove  metadata for a dropdown or radio for a metafield.
  
  > Note: Updates to metadata overwrite existing values. To remove one or more values, specify all metadata values except those you want to remove.

- Add or update scope settings for a metafield.
  
  > Note: Scope changes overwrite existing settings. You must specify the complete scope, including the changes you want to make.

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
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | The resource type to which the metafields belong. |
| `body` | [`*models.UpdateMetafieldsRequest`](../../doc/models/update-metafields-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Metafield](../../doc/models/metafield.md).

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType_SUBSCRIPTIONS

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

Deletes a metafield from your Site. Removes the metafield and associated metadata from all Subscriptions or Customers resources on the Site.

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
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | The resource type to which the metafields belong. |
| `name` | `*string` | Query, Optional | The name of the metafield to be deleted |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType_SUBSCRIPTIONS

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

Creates metadata and metafields for a specific subscription or customer, or updates metadata values of existing metafields for a subscription or customer. Metadata values are limited to 2 KB in size.

If you create metadata on a subscription or customer with a metafield that does not already exist, the metafield is created with the metadata you specify and it is always added as a text field. You can update the input_type for the metafield with the [Update Metafield](../../doc/controllers/custom-fields.md#update-metafield) endpoint.

> Note: Each site is limited to 100 unique metafields per resource. This means you can have 100 metafields for Subscriptions and another 100 for Customers.

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
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | The resource type to which the metafields belong. |
| `resourceId` | `int` | Template, Required | The Advanced Billing id of the customer or the subscription for which the metadata applies |
| `body` | [`*models.CreateMetadataRequest`](../../doc/models/create-metadata-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Metadata](../../doc/models/metadata.md).

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType_SUBSCRIPTIONS

resourceId := 60

body := models.CreateMetadataRequest{
    Metadata:             []models.CreateMetadata{
        models.CreateMetadata{
            Name:                 models.ToPointer("Color"),
            Value:                models.ToPointer("Blue"),
        },
        models.CreateMetadata{
            Name:                 models.ToPointer("Something"),
            Value:                models.ToPointer("Useful"),
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

Lists metadata and metafields for a specific customer or subscription.

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
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | The resource type to which the metafields belong. |
| `resourceId` | `int` | Template, Required | The Advanced Billing id of the customer or the subscription for which the metadata applies |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.PaginatedMetadata](../../doc/models/paginated-metadata.md).

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListMetadataInput{
    ResourceType: models.ResourceType_SUBSCRIPTIONS,
    ResourceId:   60,
    Page:         models.ToPointer(1),
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

## Example Response *(as JSON)*

```json
{
  "total_count": 1,
  "current_page": 1,
  "total_pages": 1,
  "per_page": 50,
  "metadata": [
    {
      "id": 77889911,
      "value": "green",
      "resource_id": 1234567,
      "metafield_id": 112233,
      "deleted_at": null,
      "name": "Color"
    }
  ]
}
```


# Update Metadata

Updates metadata and metafields on the Site and the customer or subscription specified, and updates the metadata value on a subscription or customer.

If you update metadata on a subscription or customer with a metafield that does not already exist, the metafield is created with the metadata you specify and it is always added as a text field to the Site and to the subscription or customer you specify. You can update the input_type for the metafield with the Update Metafield endpoint.

Each site is limited to 100 unique metafields per resource. This means you can have 100 metafields for Subscription and another 100 for Customer.

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
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | The resource type to which the metafields belong. |
| `resourceId` | `int` | Template, Required | The Advanced Billing id of the customer or the subscription for which the metadata applies |
| `body` | [`*models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Metadata](../../doc/models/metadata.md).

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType_SUBSCRIPTIONS

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

Deletes one or more metafields (and associated metadata) from the specified subscription or customer.

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
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | The resource type to which the metafields belong. |
| `resourceId` | `int` | Template, Required | The Advanced Billing id of the customer or the subscription for which the metadata applies |
| `name` | `*string` | Query, Optional | Name of field to be removed. |
| `names` | `[]string` | Query, Optional | Names of fields to be removed. Use in query: `names[]=field1&names[]=my-field&names[]=another-field`. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

resourceType := models.ResourceType_SUBSCRIPTIONS

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

Lists  metadata for a specified array of subscriptions or customers.

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
| `resourceType` | [`models.ResourceType`](../../doc/models/resource-type.md) | Template, Required | The resource type to which the metafields belong. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`.<br><br>**Default**: `20`<br><br>**Constraints**: `<= 200` |
| `dateField` | [`*models.BasicDateField`](../../doc/models/basic-date-field.md) | Query, Optional | The type of filter you would like to apply to your search. |
| `startDate` | `*time.Time` | Query, Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns metadata with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `endDate` | `*time.Time` | Query, Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns metadata with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `startDatetime` | `*time.Time` | Query, Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns metadata with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. |
| `endDatetime` | `*time.Time` | Query, Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns metadata with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. |
| `withDeleted` | `*bool` | Query, Optional | Allow to fetch deleted metadata. |
| `resourceIds` | `[]int` | Query, Optional | Allow to fetch metadata for multiple records based on provided ids. Use in query: `resource_ids[]=122&resource_ids[]=123&resource_ids[]=124`.<br><br>**Constraints**: *Maximum Items*: `50` |
| `direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Query, Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.PaginatedMetadata](../../doc/models/paginated-metadata.md).

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListMetadataForResourceTypeInput{
    ResourceType:  models.ResourceType_SUBSCRIPTIONS,
    Page:          models.ToPointer(1),
    PerPage:       models.ToPointer(50),
    DateField:     models.ToPointer(models.BasicDateField_UPDATEDAT),
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

