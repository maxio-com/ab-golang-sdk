
# Client Class Documentation

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `subdomain` | `string` | The subdomain for your Chargify site.<br>*Default*: `"subdomain"` |
| `domain` | `string` | The Chargify server domain.<br>*Default*: `"chargify.com"` |
| `environment` | `Environment` | The API environment. <br> **Default: `Environment.PRODUCTION`** |
| `httpConfiguration` | [`HttpConfiguration`](http-configuration.md) | Configurable http client options like timeout and retries. |
| `basicAuthCredentials` | [`BasicAuthCredentials`](auth/basic-authentication.md) | The Credentials Setter for Basic Authentication |

The API client can be initialized as follows:

```go
client := advancedbilling.NewClient(
    advancedbilling.CreateConfiguration(
        advancedbilling.WithHttpConfiguration(
            advancedbilling.CreateHttpConfiguration(
                advancedbilling.WithTimeout(30),
            ),
        ),
        advancedbilling.WithEnvironment(advancedbilling.PRODUCTION),
        advancedbilling.WithBasicAuthCredentials(
            advancedbilling.NewBasicAuthCredentials(
                "BasicAuthUserName",
                "BasicAuthPassword",
            ),
        ),
        advancedbilling.WithSubdomain("subdomain"),
        advancedbilling.WithDomain("chargify.com"),
    ),
)
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
| componentPricePoints | Gets ComponentPricePointsController |
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

