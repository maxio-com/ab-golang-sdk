# Subscription Renewals

```go
subscriptionRenewalsController := client.SubscriptionRenewalsController()
```

## Class Name

`SubscriptionRenewalsController`

## Methods

* [Create Scheduled Renewal Configuration](../../doc/controllers/subscription-renewals.md#create-scheduled-renewal-configuration)
* [List Scheduled Renewal Configurations](../../doc/controllers/subscription-renewals.md#list-scheduled-renewal-configurations)
* [Read Scheduled Renewal Configuration](../../doc/controllers/subscription-renewals.md#read-scheduled-renewal-configuration)
* [Update Scheduled Renewal Configuration](../../doc/controllers/subscription-renewals.md#update-scheduled-renewal-configuration)
* [Schedule Scheduled Renewal Lock In](../../doc/controllers/subscription-renewals.md#schedule-scheduled-renewal-lock-in)
* [Lock in Scheduled Renewal Immediately](../../doc/controllers/subscription-renewals.md#lock-in-scheduled-renewal-immediately)
* [Unpublish Scheduled Renewal Configuration](../../doc/controllers/subscription-renewals.md#unpublish-scheduled-renewal-configuration)
* [Cancel Scheduled Renewal Configuration](../../doc/controllers/subscription-renewals.md#cancel-scheduled-renewal-configuration)
* [Create Scheduled Renewal Configuration Item](../../doc/controllers/subscription-renewals.md#create-scheduled-renewal-configuration-item)
* [Update Scheduled Renewal Configuration Item](../../doc/controllers/subscription-renewals.md#update-scheduled-renewal-configuration-item)
* [Delete Scheduled Renewal Configuration Item](../../doc/controllers/subscription-renewals.md#delete-scheduled-renewal-configuration-item)


# Create Scheduled Renewal Configuration

Creates a scheduled renewal configuration for a subscription. The scheduled renewal is based on the subscription’s current product and component setup.

```go
CreateScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    body *models.ScheduledRenewalConfigurationRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `body` | [`*models.ScheduledRenewalConfigurationRequest`](../../doc/models/scheduled-renewal-configuration-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationResponse](../../doc/models/scheduled-renewal-configuration-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

body := models.ScheduledRenewalConfigurationRequest{
    RenewalConfiguration: models.ScheduledRenewalConfigurationRequestBody{
        StartsAt:             models.ToPointer(parseTime(time.RFC3339, "2024-12-01T00:00:00Z", func(err error) { log.Fatalln(err) })),
        EndsAt:               models.ToPointer(parseTime(time.RFC3339, "2025-12-01T00:00:00Z", func(err error) { log.Fatalln(err) })),
        LockInAt:             models.ToPointer(parseTime(time.RFC3339, "2024-11-15T00:00:00Z", func(err error) { log.Fatalln(err) })),
        ContractId:           models.ToPointer(222),
    },
}

apiResponse, err := subscriptionRenewalsController.CreateScheduledRenewalConfiguration(ctx, subscriptionId, &body)
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
{
  "scheduled_renewal_configuration": {
    "id": 123,
    "site_id": 456,
    "subscription_id": 12345,
    "starts_at": "2024-12-01T00:00:00Z",
    "ends_at": "2025-12-01T00:00:00Z",
    "lock_in_at": "2024-11-15T00:00:00Z",
    "created_at": "2024-09-01T12:00:00Z",
    "status": "scheduled",
    "scheduled_renewal_configuration_items": [
      {
        "id": 789,
        "subscription_id": 12345,
        "subscription_renewal_configuration_id": 123,
        "item_id": 4,
        "item_type": "Product",
        "item_subclass": "Product",
        "price_point_id": 7,
        "price_point_type": "ProductPricePoint",
        "quantity": 1,
        "decimal_quantity": "1.0",
        "created_at": "2024-09-01T12:00:00Z"
      }
    ],
    "contract": {
      "id": 107,
      "maxio_id": "maxio-id",
      "number": null,
      "register": {
        "id": 12,
        "maxio_id": "maxio_id-id",
        "name": "Register",
        "currency_code": "USD"
      }
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Scheduled Renewal Configurations

Lists scheduled renewal configurations for the subscription and permits an optional status query filter.

```go
ListScheduledRenewalConfigurations(
    ctx context.Context,
    subscriptionId int,
    status *models.Status) (
    models.ApiResponse[models.ScheduledRenewalConfigurationsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `status` | [`*models.Status`](../../doc/models/status.md) | Query, Optional | (Optional) Status filter for scheduled renewal configurations. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationsResponse](../../doc/models/scheduled-renewal-configurations-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

apiResponse, err := subscriptionRenewalsController.ListScheduledRenewalConfigurations(ctx, subscriptionId, nil)
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
  "scheduled_renewal_configurations": [
    {
      "id": 123,
      "site_id": 456,
      "subscription_id": 12345,
      "starts_at": "2024-12-01T00:00:00Z",
      "ends_at": "2025-12-01T00:00:00Z",
      "lock_in_at": "2024-11-15T00:00:00Z",
      "created_at": "2024-09-01T12:00:00Z",
      "status": "scheduled",
      "scheduled_renewal_configuration_items": [
        {
          "id": 789,
          "subscription_id": 12345,
          "subscription_renewal_configuration_id": 123,
          "item_id": 4,
          "item_type": "Product",
          "item_subclass": "Product",
          "price_point_id": 7,
          "price_point_type": "ProductPricePoint",
          "quantity": 1,
          "decimal_quantity": "1.0",
          "created_at": "2024-09-01T12:00:00Z"
        }
      ],
      "contract": {
        "id": 107,
        "maxio_id": "maxio-id",
        "number": null,
        "register": {
          "id": 12,
          "maxio_id": "maxio-id",
          "name": "Register",
          "currency_code": "USD"
        }
      }
    }
  ]
}
```


# Read Scheduled Renewal Configuration

Retrieves the configuration settings for the scheduled renewal.

```go
ReadScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    id int) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `id` | `int` | Template, Required | The renewal id. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationResponse](../../doc/models/scheduled-renewal-configuration-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

id := 112

apiResponse, err := subscriptionRenewalsController.ReadScheduledRenewalConfiguration(ctx, subscriptionId, id)
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
  "scheduled_renewal_configuration": {
    "id": 123,
    "site_id": 456,
    "subscription_id": 12345,
    "starts_at": "2024-12-01T00:00:00Z",
    "ends_at": "2025-12-01T00:00:00Z",
    "lock_in_at": "2024-11-15T00:00:00Z",
    "created_at": "2024-09-01T12:00:00Z",
    "status": "scheduled",
    "scheduled_renewal_configuration_items": [
      {
        "id": 789,
        "subscription_id": 12345,
        "subscription_renewal_configuration_id": 123,
        "item_id": 4,
        "item_type": "Product",
        "item_subclass": "Product",
        "price_point_id": 7,
        "price_point_type": "ProductPricePoint",
        "quantity": 1,
        "decimal_quantity": "1.0",
        "created_at": "2024-09-01T12:00:00Z"
      }
    ],
    "contract": {
      "id": 107,
      "maxio_id": "maxio-id",
      "number": null,
      "register": {
        "id": 12,
        "maxio_id": "maxio-id",
        "name": "Register",
        "currency_code": "USD"
      }
    }
  }
}
```


# Update Scheduled Renewal Configuration

Updates an existing configuration.

```go
UpdateScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    id int,
    body *models.ScheduledRenewalConfigurationRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `id` | `int` | Template, Required | The renewal id. |
| `body` | [`*models.ScheduledRenewalConfigurationRequest`](../../doc/models/scheduled-renewal-configuration-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationResponse](../../doc/models/scheduled-renewal-configuration-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

id := 112

body := models.ScheduledRenewalConfigurationRequest{
    RenewalConfiguration: models.ScheduledRenewalConfigurationRequestBody{
        StartsAt:             models.ToPointer(parseTime(time.RFC3339, "2025-12-01T00:00:00Z", func(err error) { log.Fatalln(err) })),
        EndsAt:               models.ToPointer(parseTime(time.RFC3339, "2026-12-01T00:00:00Z", func(err error) { log.Fatalln(err) })),
        LockInAt:             models.ToPointer(parseTime(time.RFC3339, "2025-11-15T00:00:00Z", func(err error) { log.Fatalln(err) })),
    },
}

apiResponse, err := subscriptionRenewalsController.UpdateScheduledRenewalConfiguration(ctx, subscriptionId, id, &body)
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
{
  "scheduled_renewal_configuration": {
    "id": 123,
    "site_id": 456,
    "subscription_id": 12345,
    "starts_at": "2025-12-01T00:00:00Z",
    "ends_at": "2026-12-01T00:00:00Z",
    "lock_in_at": "2025-11-15T00:00:00Z",
    "created_at": "2025-09-01T12:00:00Z",
    "status": "scheduled",
    "scheduled_renewal_configuration_items": [
      {
        "id": 789,
        "subscription_id": 12345,
        "subscription_renewal_configuration_id": 123,
        "item_id": 4,
        "item_type": "Product",
        "item_subclass": "Product",
        "price_point_id": 7,
        "price_point_type": "ProductPricePoint",
        "quantity": 1,
        "decimal_quantity": "1.0",
        "created_at": "2025-09-01T12:00:00Z"
      }
    ],
    "contract": {
      "id": 107,
      "maxio_id": "maxio-id",
      "number": null,
      "register": {
        "id": 12,
        "maxio_id": "maxio-id",
        "name": "Register",
        "currency_code": "USD"
      }
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Schedule Scheduled Renewal Lock In

Schedules a future lock-in date for the renewal.

```go
ScheduleScheduledRenewalLockIn(
    ctx context.Context,
    subscriptionId int,
    id int,
    body *models.ScheduledRenewalLockInRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `id` | `int` | Template, Required | The renewal id. |
| `body` | [`*models.ScheduledRenewalLockInRequest`](../../doc/models/scheduled-renewal-lock-in-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationResponse](../../doc/models/scheduled-renewal-configuration-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

id := 112

body := models.ScheduledRenewalLockInRequest{
    LockInAt:             parseTime(models.DEFAULT_DATE, "2025-11-15", func(err error) { log.Fatalln(err) }),
}

apiResponse, err := subscriptionRenewalsController.ScheduleScheduledRenewalLockIn(ctx, subscriptionId, id, &body)
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
{
  "scheduled_renewal_configuration": {
    "id": 123,
    "site_id": 456,
    "subscription_id": 12345,
    "starts_at": "2025-12-01T00:00:00Z",
    "ends_at": "2026-12-01T00:00:00Z",
    "lock_in_at": "2025-11-15T00:00:00Z",
    "created_at": "2025-09-01T12:00:00Z",
    "status": "scheduled",
    "scheduled_renewal_configuration_items": [
      {
        "id": 789,
        "subscription_id": 12345,
        "subscription_renewal_configuration_id": 123,
        "item_id": 4,
        "item_type": "Product",
        "item_subclass": "Product",
        "price_point_id": 7,
        "price_point_type": "ProductPricePoint",
        "quantity": 1,
        "decimal_quantity": "1.0",
        "created_at": "2025-09-01T12:00:00Z"
      }
    ],
    "contract": {
      "id": 107,
      "maxio_id": "maxio-id",
      "number": null,
      "register": {
        "id": 12,
        "maxio_id": "maxio-id",
        "name": "Register",
        "currency_code": "USD"
      }
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Lock in Scheduled Renewal Immediately

Locks in the renewal immediately.

```go
LockInScheduledRenewalImmediately(
    ctx context.Context,
    subscriptionId int,
    id int) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `id` | `int` | Template, Required | The renewal id. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationResponse](../../doc/models/scheduled-renewal-configuration-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

id := 112

apiResponse, err := subscriptionRenewalsController.LockInScheduledRenewalImmediately(ctx, subscriptionId, id)
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
{
  "scheduled_renewal_configuration": {
    "id": 987,
    "site_id": 321,
    "subscription_id": 12345,
    "starts_at": "2025-12-01T00:00:00Z",
    "ends_at": "2026-12-01T00:00:00Z",
    "lock_in_at": "2025-11-15T00:00:00Z",
    "created_at": "2025-09-01T12:00:00Z",
    "status": "scheduled",
    "scheduled_renewal_configuration_items": [
      {
        "id": 555,
        "subscription_id": 12345,
        "subscription_renewal_configuration_id": 987,
        "item_id": 42,
        "item_type": "Product",
        "price_point_id": 73,
        "price_point_type": "ProductPricePoint",
        "quantity": 1,
        "decimal_quantity": "1.0",
        "created_at": "2025-09-01T12:00:00Z"
      }
    ],
    "contract": {
      "id": 222,
      "maxio_id": "maxio-id",
      "number": null,
      "register": {
        "id": 12,
        "maxio_id": "maxio-id",
        "name": "Register",
        "currency_code": "USD"
      }
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Unpublish Scheduled Renewal Configuration

Returns a scheduled renewal configuration to an editable state.

```go
UnpublishScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    id int) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `id` | `int` | Template, Required | The renewal id. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationResponse](../../doc/models/scheduled-renewal-configuration-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

id := 112

apiResponse, err := subscriptionRenewalsController.UnpublishScheduledRenewalConfiguration(ctx, subscriptionId, id)
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
{
  "scheduled_renewal_configuration": {
    "id": 987,
    "site_id": 321,
    "subscription_id": 12345,
    "starts_at": "2025-12-01T00:00:00Z",
    "ends_at": "2026-12-01T00:00:00Z",
    "lock_in_at": "2025-11-15T00:00:00Z",
    "created_at": "2025-09-01T12:00:00Z",
    "status": "draft",
    "scheduled_renewal_configuration_items": [
      {
        "id": 555,
        "subscription_id": 12345,
        "subscription_renewal_configuration_id": 987,
        "item_id": 42,
        "item_type": "Product",
        "price_point_id": 73,
        "price_point_type": "ProductPricePoint",
        "quantity": 1,
        "decimal_quantity": "1.0",
        "created_at": "2025-09-01T12:00:00Z"
      }
    ],
    "contract": {
      "id": 222
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Cancel Scheduled Renewal Configuration

Cancels a scheduled renewal configuration.

```go
CancelScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    id int) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `id` | `int` | Template, Required | The renewal id. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationResponse](../../doc/models/scheduled-renewal-configuration-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

id := 112

apiResponse, err := subscriptionRenewalsController.CancelScheduledRenewalConfiguration(ctx, subscriptionId, id)
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
{
  "scheduled_renewal_configuration": {
    "id": 987,
    "site_id": 321,
    "subscription_id": 12345,
    "starts_at": "2025-12-01T00:00:00Z",
    "ends_at": "2026-12-01T00:00:00Z",
    "lock_in_at": "2025-11-15T00:00:00Z",
    "created_at": "2025-09-01T12:00:00Z",
    "status": "canceled",
    "scheduled_renewal_configuration_items": [
      {
        "id": 555,
        "subscription_id": 12345,
        "subscription_renewal_configuration_id": 987,
        "item_id": 42,
        "item_type": "Product",
        "price_point_id": 73,
        "price_point_type": "ProductPricePoint",
        "quantity": 1,
        "decimal_quantity": "1.0",
        "created_at": "2025-09-01T12:00:00Z"
      }
    ],
    "contract": {
      "id": 222
    }
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Create Scheduled Renewal Configuration Item

Adds product and component line items to the scheduled renewal.

```go
CreateScheduledRenewalConfigurationItem(
    ctx context.Context,
    subscriptionId int,
    scheduledRenewalsConfigurationId int,
    body *models.ScheduledRenewalConfigurationItemRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `scheduledRenewalsConfigurationId` | `int` | Template, Required | The scheduled renewal configuration id. |
| `body` | [`*models.ScheduledRenewalConfigurationItemRequest`](../../doc/models/scheduled-renewal-configuration-item-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationItemResponse](../../doc/models/scheduled-renewal-configuration-item-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

scheduledRenewalsConfigurationId := 250

body := models.ScheduledRenewalConfigurationItemRequest{
    RenewalConfigurationItem: models.ScheduledRenewalConfigurationItemRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyComponent(models.ScheduledRenewalItemRequestBodyComponent{
        ItemType:             "Component",
        ItemId:               57,
        Quantity:             models.ToPointer(1),
        CustomPrice:          models.ToPointer(models.ScheduledRenewalComponentCustomPrice{
            PricingScheme:        models.PricingScheme_STAIRSTEP,
            Prices:               []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(1),
                    EndingQuantity:       models.NewOptional[models.PriceEndingQuantity](nil),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(5)),
                },
            },
        }),
    }),
}

apiResponse, err := subscriptionRenewalsController.CreateScheduledRenewalConfigurationItem(ctx, subscriptionId, scheduledRenewalsConfigurationId, &body)
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
{
  "scheduled_renewal_configuration_item": {
    "id": 555,
    "subscription_id": 12345,
    "subscription_renewal_configuration_id": 987,
    "item_id": 42,
    "item_type": "Product",
    "item_subclass": "SubscriptionProduct",
    "price_point_id": 73,
    "price_point_type": "ProductPricePoint",
    "quantity": 1,
    "decimal_quantity": "1.0",
    "created_at": "2025-09-01T12:00:00Z"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Update Scheduled Renewal Configuration Item

Updates an existing configuration item’s pricing and quantity.

```go
UpdateScheduledRenewalConfigurationItem(
    ctx context.Context,
    subscriptionId int,
    scheduledRenewalsConfigurationId int,
    id int,
    body *models.ScheduledRenewalUpdateRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `scheduledRenewalsConfigurationId` | `int` | Template, Required | The scheduled renewal configuration id. |
| `id` | `int` | Template, Required | The scheduled renewal configuration item id. |
| `body` | [`*models.ScheduledRenewalUpdateRequest`](../../doc/models/scheduled-renewal-update-request.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ScheduledRenewalConfigurationItemResponse](../../doc/models/scheduled-renewal-configuration-item-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

scheduledRenewalsConfigurationId := 250

id := 112

body := models.ScheduledRenewalUpdateRequest{
    RenewalConfigurationItem: models.ScheduledRenewalUpdateRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyComponent(models.ScheduledRenewalItemRequestBodyComponent{
        ItemType:             "Component",
        ItemId:               57,
        Quantity:             models.ToPointer(2),
        CustomPrice:          models.ToPointer(models.ScheduledRenewalComponentCustomPrice{
            PricingScheme:        models.PricingScheme_STAIRSTEP,
            Prices:               []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(1),
                    EndingQuantity:       models.NewOptional[models.PriceEndingQuantity](nil),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(5)),
                },
            },
        }),
    }),
}

apiResponse, err := subscriptionRenewalsController.UpdateScheduledRenewalConfigurationItem(ctx, subscriptionId, scheduledRenewalsConfigurationId, id, &body)
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
{
  "scheduled_renewal_configuration_item": {
    "id": 555,
    "subscription_id": 12345,
    "subscription_renewal_configuration_id": 987,
    "item_id": 42,
    "item_type": "Component",
    "item_subclass": "SubscriptionComponent",
    "price_point_id": 73,
    "price_point_type": "ComponentPricePoint",
    "quantity": 3,
    "decimal_quantity": "3.0",
    "created_at": "2025-09-01T12:00:00Z"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# Delete Scheduled Renewal Configuration Item

Removes an item from the pending renewal configuration.

```go
DeleteScheduledRenewalConfigurationItem(
    ctx context.Context,
    subscriptionId int,
    scheduledRenewalsConfigurationId int,
    id int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |
| `scheduledRenewalsConfigurationId` | `int` | Template, Required | The scheduled renewal configuration id. |
| `id` | `int` | Template, Required | The scheduled renewal configuration item id. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

scheduledRenewalsConfigurationId := 250

id := 112

resp, err := subscriptionRenewalsController.DeleteScheduledRenewalConfigurationItem(ctx, subscriptionId, scheduledRenewalsConfigurationId, id)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ErrorListResponse:
            log.Fatalln("ErrorListResponseException: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

