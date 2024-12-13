
# Client Class Documentation

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `site` | `string` | The subdomain for your Advanced Billing site.<br>*Default*: `"subdomain"` |
| `environment` | `Environment` | The API environment. <br> **Default: `Environment.US`** |
| `httpConfiguration` | [`HttpConfiguration`](http-configuration.md) | Configurable http client options like timeout and retries. |
| `basicAuthCredentials` | [`BasicAuthCredentials`](auth/basic-authentication.md) | The Credentials Setter for Basic Authentication |

The API client can be initialized as follows:

```go
client := advancedbilling.NewClient(
    advancedbilling.CreateConfiguration(
        advancedbilling.WithHttpConfiguration(
            advancedbilling.CreateHttpConfiguration(
                advancedbilling.WithTimeout(120),
            ),
        ),
        advancedbilling.WithEnvironment(advancedbilling.US),
        advancedbilling.WithBasicAuthCredentials(
            advancedbilling.NewBasicAuthCredentials(
                "BasicAuthUserName",
                "BasicAuthPassword",
            ),
        ),
        advancedbilling.WithSite("subdomain"),
    ),
)
```

## Maxio Advanced Billing Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name | Description |
|  --- | --- |
| APIExportsController() | Gets APIExportsController |
| AdvanceInvoiceController() | Gets AdvanceInvoiceController |
| BillingPortalController() | Gets BillingPortalController |
| CouponsController() | Gets CouponsController |
| ComponentsController() | Gets ComponentsController |
| ComponentPricePointsController() | Gets ComponentPricePointsController |
| CustomersController() | Gets CustomersController |
| CustomFieldsController() | Gets CustomFieldsController |
| EventsController() | Gets EventsController |
| EventsBasedBillingSegmentsController() | Gets EventsBasedBillingSegmentsController |
| InsightsController() | Gets InsightsController |
| InvoicesController() | Gets InvoicesController |
| OffersController() | Gets OffersController |
| PaymentProfilesController() | Gets PaymentProfilesController |
| ProductFamiliesController() | Gets ProductFamiliesController |
| ProductsController() | Gets ProductsController |
| ProductPricePointsController() | Gets ProductPricePointsController |
| ProformaInvoicesController() | Gets ProformaInvoicesController |
| ReasonCodesController() | Gets ReasonCodesController |
| ReferralCodesController() | Gets ReferralCodesController |
| SalesCommissionsController() | Gets SalesCommissionsController |
| SitesController() | Gets SitesController |
| SubscriptionsController() | Gets SubscriptionsController |
| SubscriptionComponentsController() | Gets SubscriptionComponentsController |
| SubscriptionGroupsController() | Gets SubscriptionGroupsController |
| SubscriptionGroupInvoiceAccountController() | Gets SubscriptionGroupInvoiceAccountController |
| SubscriptionGroupStatusController() | Gets SubscriptionGroupStatusController |
| SubscriptionInvoiceAccountController() | Gets SubscriptionInvoiceAccountController |
| SubscriptionNotesController() | Gets SubscriptionNotesController |
| SubscriptionProductsController() | Gets SubscriptionProductsController |
| SubscriptionStatusController() | Gets SubscriptionStatusController |
| WebhooksController() | Gets WebhooksController |

