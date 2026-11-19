# Entitlements

Entitlements report what a subscription is allowed to do, and how much of it, based on the features granted by the subscription's products and components.

When more than one product or component grants the same feature, the values are aggregated into one entry per feature key and periodicity. For more information, see [Read Subscription Entitlements with the API](https://docs.maxio.com/billing/catalog/entitlements/read-subscription-entitlements-with-the-api).

```go
entitlementsController := client.EntitlementsController()
```

## Class Name

`EntitlementsController`


# Read Subscription Entitlements

Returns every feature a subscription is entitled to, collapsed into one entry per feature key and periodicity window across all products and components on the subscription. A `usage_limit` feature granted with two different periodicities comes back as two entries sharing one `feature_key`, each identified by its own `periodicity_key`.

When more than one product or component grants the same feature key and periodicity, the values are combined:

- **`access_right`** features are combined with a boolean OR. If any contributor grants access, the aggregate is `true`. `source_products` only lists the contributors that granted `true`.
- **`usage_limit`** features are summed across every contributor sharing the same periodicity window. `source_products` lists every contributor. Grants with different periodicities are not summed together. Each periodicity is returned as a separate entry.
- **`service_right`** features are not combined: one contributor's value wins. Do not rely on which one when several grant the same feature key.

`enabled` reflects both the aggregated value and the subscription's state. The field is `false` whenever the subscription is not in a live state (`active`, `trialing`, `assessing`, `past_due`, `soft_failure`), regardless of the aggregated value. Entitlements deliberately stay enabled through dunning.

```go
ReadSubscriptionEntitlements(
    ctx context.Context,
    subscriptionId int) (
    models.ApiResponse[models.AggregatedEntitlementsResponse],
    error)
```

## Authentication

This endpoint requires [BasicAuth](../../doc/auth/basic-authentication.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `int` | Template, Required | The Chargify id of the subscription. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.AggregatedEntitlementsResponse](../../doc/models/aggregated-entitlements-response.md).

## Example Usage

```go
ctx := context.Background()

subscriptionId := 222

apiResponse, err := entitlementsController.ReadSubscriptionEntitlements(ctx, subscriptionId)
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
  "subscription_id": 12345,
  "customer_id": 678,
  "status": "active",
  "entitlements": [
    {
      "feature_key": "feature.sso",
      "periodicity_key": "feature.sso",
      "name": "SSO",
      "type": "access_right",
      "value": true,
      "enabled": true,
      "periodicity": null,
      "source_products": [
        "Gold Plan"
      ]
    },
    {
      "feature_key": "usage.api_calls",
      "periodicity_key": "usage.api_calls:1:month",
      "name": "API Calls",
      "type": "usage_limit",
      "value": 50000,
      "enabled": true,
      "periodicity": {
        "interval": 1,
        "unit": "month"
      },
      "source_products": [
        "Gold Plan"
      ]
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 403 | Forbidden | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |
| 404 | Not Found | `ApiError` |

