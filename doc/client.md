
# Client Class Documentation

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `subdomain` | `string` | The subdomain for your Chargify site.<br>*Default*: `"subdomain"` |
| `domain` | `string` | The Chargify server domain.<br>*Default*: `"chargify.com"` |
| `environment` | Environment | The API environment. <br> **Default: `Environment.PRODUCTION`** |
| `httpConfiguration` | [`HttpConfiguration`](http-configuration.md) | Configurable http client options like timeout and retries. |
| `basicAuthUserName` | `string` | The username to use with basic authentication |
| `basicAuthPassword` | `string` | The password to use with basic authentication |

The API client can be initialized as follows:

```go
config := advancedbilling.CreateConfiguration(
    advancedbilling.WithHttpConfiguration(
        advancedbilling.CreateHttpConfiguration(
            advancedbilling.WithTimeout(30),
            advancedbilling.WithTransport(http.DefaultTransport),
            advancedbilling.WithRetryConfiguration(
                advancedbilling.CreateRetryConfiguration(
                    advancedbilling.WithMaxRetryAttempts(0),
                    advancedbilling.WithRetryOnTimeout(true),
                    advancedbilling.WithRetryInterval(1),
                    advancedbilling.WithMaximumRetryWaitTime(0),
                    advancedbilling.WithBackoffFactor(2),
                    advancedbilling.WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524, 408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
                    advancedbilling.WithHttpMethodsToRetry([]string{"GET", "PUT", "GET", "PUT"}),
                ),
            ),
        ),
    ),
    advancedbilling.WithEnvironment(advancedbilling.PRODUCTION),
    advancedbilling.WithBasicAuthUserName("BasicAuthUserName"),
    advancedbilling.WithBasicAuthPassword("BasicAuthPassword"),
)
client := advancedbilling.NewClient(config)
```

## Maxio Advanced Billing Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name | Description |
|  --- | --- |
| aPIExports | Gets APIExportsController |
| advanceInvoice | Gets AdvanceInvoiceController |
| billingPortal | Gets BillingPortalController |
| coupons | Gets CouponsController |
| components | Gets ComponentsController |
| customers | Gets CustomersController |
| customFields | Gets CustomFieldsController |
| events | Gets EventsController |
| eventsBasedBillingSegments | Gets EventsBasedBillingSegmentsController |
| insights | Gets InsightsController |
| invoices | Gets InvoicesController |
| offers | Gets OffersController |
| paymentProfiles | Gets PaymentProfilesController |
| productFamilies | Gets ProductFamiliesController |
| products | Gets ProductsController |
| productPricePoints | Gets ProductPricePointsController |
| proformaInvoices | Gets ProformaInvoicesController |
| reasonCodes | Gets ReasonCodesController |
| referralCodes | Gets ReferralCodesController |
| salesCommissions | Gets SalesCommissionsController |
| sites | Gets SitesController |
| subscriptions | Gets SubscriptionsController |
| subscriptionComponents | Gets SubscriptionComponentsController |
| subscriptionGroups | Gets SubscriptionGroupsController |
| subscriptionGroupInvoiceAccount | Gets SubscriptionGroupInvoiceAccountController |
| subscriptionGroupStatus | Gets SubscriptionGroupStatusController |
| subscriptionInvoiceAccount | Gets SubscriptionInvoiceAccountController |
| subscriptionNotes | Gets SubscriptionNotesController |
| subscriptionProducts | Gets SubscriptionProductsController |
| subscriptionStatus | Gets SubscriptionStatusController |
| webhooks | Gets WebhooksController |

