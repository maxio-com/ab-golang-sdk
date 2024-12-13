/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "net/http"
)

// Client is an interface representing the main client for accessing configuration and controllers.
type ClientInterface interface {
    Configuration() Configuration
    CloneWithConfiguration(options ...ConfigurationOptions) ClientInterface
    APIExportsController() *APIExportsController
    AdvanceInvoiceController() *AdvanceInvoiceController
    BillingPortalController() *BillingPortalController
    CouponsController() *CouponsController
    ComponentsController() *ComponentsController
    ComponentPricePointsController() *ComponentPricePointsController
    CustomersController() *CustomersController
    CustomFieldsController() *CustomFieldsController
    EventsController() *EventsController
    EventsBasedBillingSegmentsController() *EventsBasedBillingSegmentsController
    InsightsController() *InsightsController
    InvoicesController() *InvoicesController
    OffersController() *OffersController
    PaymentProfilesController() *PaymentProfilesController
    ProductFamiliesController() *ProductFamiliesController
    ProductsController() *ProductsController
    ProductPricePointsController() *ProductPricePointsController
    ProformaInvoicesController() *ProformaInvoicesController
    ReasonCodesController() *ReasonCodesController
    ReferralCodesController() *ReferralCodesController
    SalesCommissionsController() *SalesCommissionsController
    SitesController() *SitesController
    SubscriptionsController() *SubscriptionsController
    SubscriptionComponentsController() *SubscriptionComponentsController
    SubscriptionGroupsController() *SubscriptionGroupsController
    SubscriptionGroupInvoiceAccountController() *SubscriptionGroupInvoiceAccountController
    SubscriptionGroupStatusController() *SubscriptionGroupStatusController
    SubscriptionInvoiceAccountController() *SubscriptionInvoiceAccountController
    SubscriptionNotesController() *SubscriptionNotesController
    SubscriptionProductsController() *SubscriptionProductsController
    SubscriptionStatusController() *SubscriptionStatusController
    WebhooksController() *WebhooksController
    UserAgent() *string
}

// client is an implementation of the Client interface.
type client struct {
    callBuilderFactory                        https.CallBuilderFactory
    configuration                             Configuration
    userAgent                                 string
    apiExportsController                      APIExportsController
    advanceInvoiceController                  AdvanceInvoiceController
    billingPortalController                   BillingPortalController
    couponsController                         CouponsController
    componentsController                      ComponentsController
    componentPricePointsController            ComponentPricePointsController
    customersController                       CustomersController
    customFieldsController                    CustomFieldsController
    eventsController                          EventsController
    eventsBasedBillingSegmentsController      EventsBasedBillingSegmentsController
    insightsController                        InsightsController
    invoicesController                        InvoicesController
    offersController                          OffersController
    paymentProfilesController                 PaymentProfilesController
    productFamiliesController                 ProductFamiliesController
    productsController                        ProductsController
    productPricePointsController              ProductPricePointsController
    proformaInvoicesController                ProformaInvoicesController
    reasonCodesController                     ReasonCodesController
    referralCodesController                   ReferralCodesController
    salesCommissionsController                SalesCommissionsController
    sitesController                           SitesController
    subscriptionsController                   SubscriptionsController
    subscriptionComponentsController          SubscriptionComponentsController
    subscriptionGroupsController              SubscriptionGroupsController
    subscriptionGroupInvoiceAccountController SubscriptionGroupInvoiceAccountController
    subscriptionGroupStatusController         SubscriptionGroupStatusController
    subscriptionInvoiceAccountController      SubscriptionInvoiceAccountController
    subscriptionNotesController               SubscriptionNotesController
    subscriptionProductsController            SubscriptionProductsController
    subscriptionStatusController              SubscriptionStatusController
    webhooksController                        WebhooksController
}

// NewClient is the constructor for creating a new client instance.
// It takes a Configuration object as a parameter and returns the Client interface.
func NewClient(configuration Configuration) ClientInterface {
    client := &client{
        configuration: configuration,
    }
    
    client.userAgent = utilities.UpdateUserAgent("AB SDK Go:5.0.0 on OS {os-info}")
    client.callBuilderFactory = callBuilderHandler(
    	func(server string) string {
    		if server == "" {
    			server = "production"
    		}
    		return getBaseUri(Server(server), client.configuration)
    	},
    	createAuthenticationFromConfig(configuration),
    	https.NewHttpClient(configuration.HttpConfiguration()),
    	configuration.httpConfiguration.RetryConfiguration(),
    	https.Csv,
        withUserAgent(client.userAgent),
        withGlobalErrors(),
    )
    
    baseController := NewBaseController(client)
    client.apiExportsController = *NewAPIExportsController(*baseController)
    client.advanceInvoiceController = *NewAdvanceInvoiceController(*baseController)
    client.billingPortalController = *NewBillingPortalController(*baseController)
    client.couponsController = *NewCouponsController(*baseController)
    client.componentsController = *NewComponentsController(*baseController)
    client.componentPricePointsController = *NewComponentPricePointsController(*baseController)
    client.customersController = *NewCustomersController(*baseController)
    client.customFieldsController = *NewCustomFieldsController(*baseController)
    client.eventsController = *NewEventsController(*baseController)
    client.eventsBasedBillingSegmentsController = *NewEventsBasedBillingSegmentsController(*baseController)
    client.insightsController = *NewInsightsController(*baseController)
    client.invoicesController = *NewInvoicesController(*baseController)
    client.offersController = *NewOffersController(*baseController)
    client.paymentProfilesController = *NewPaymentProfilesController(*baseController)
    client.productFamiliesController = *NewProductFamiliesController(*baseController)
    client.productsController = *NewProductsController(*baseController)
    client.productPricePointsController = *NewProductPricePointsController(*baseController)
    client.proformaInvoicesController = *NewProformaInvoicesController(*baseController)
    client.reasonCodesController = *NewReasonCodesController(*baseController)
    client.referralCodesController = *NewReferralCodesController(*baseController)
    client.salesCommissionsController = *NewSalesCommissionsController(*baseController)
    client.sitesController = *NewSitesController(*baseController)
    client.subscriptionsController = *NewSubscriptionsController(*baseController)
    client.subscriptionComponentsController = *NewSubscriptionComponentsController(*baseController)
    client.subscriptionGroupsController = *NewSubscriptionGroupsController(*baseController)
    client.subscriptionGroupInvoiceAccountController = *NewSubscriptionGroupInvoiceAccountController(*baseController)
    client.subscriptionGroupStatusController = *NewSubscriptionGroupStatusController(*baseController)
    client.subscriptionInvoiceAccountController = *NewSubscriptionInvoiceAccountController(*baseController)
    client.subscriptionNotesController = *NewSubscriptionNotesController(*baseController)
    client.subscriptionProductsController = *NewSubscriptionProductsController(*baseController)
    client.subscriptionStatusController = *NewSubscriptionStatusController(*baseController)
    client.webhooksController = *NewWebhooksController(*baseController)
    return client
}

// Configuration returns the configuration instance of the client.
func (c *client) Configuration() Configuration {
    return c.configuration
}

// CloneWithConfiguration returns a new copy with the provided options of the configuration instance of the client.
func (c *client) CloneWithConfiguration(options ...ConfigurationOptions) ClientInterface {
    return NewClient(c.configuration.cloneWithOptions(options...))
}

// APIExportsController returns the apiExportsController instance of the client.
func (c *client) APIExportsController() *APIExportsController {
    return &c.apiExportsController
}

// AdvanceInvoiceController returns the advanceInvoiceController instance of the client.
func (c *client) AdvanceInvoiceController() *AdvanceInvoiceController {
    return &c.advanceInvoiceController
}

// BillingPortalController returns the billingPortalController instance of the client.
func (c *client) BillingPortalController() *BillingPortalController {
    return &c.billingPortalController
}

// CouponsController returns the couponsController instance of the client.
func (c *client) CouponsController() *CouponsController {
    return &c.couponsController
}

// ComponentsController returns the componentsController instance of the client.
func (c *client) ComponentsController() *ComponentsController {
    return &c.componentsController
}

// ComponentPricePointsController returns the componentPricePointsController instance of the client.
func (c *client) ComponentPricePointsController() *ComponentPricePointsController {
    return &c.componentPricePointsController
}

// CustomersController returns the customersController instance of the client.
func (c *client) CustomersController() *CustomersController {
    return &c.customersController
}

// CustomFieldsController returns the customFieldsController instance of the client.
func (c *client) CustomFieldsController() *CustomFieldsController {
    return &c.customFieldsController
}

// EventsController returns the eventsController instance of the client.
func (c *client) EventsController() *EventsController {
    return &c.eventsController
}

// EventsBasedBillingSegmentsController returns the eventsBasedBillingSegmentsController instance of the client.
func (c *client) EventsBasedBillingSegmentsController() *EventsBasedBillingSegmentsController {
    return &c.eventsBasedBillingSegmentsController
}

// InsightsController returns the insightsController instance of the client.
func (c *client) InsightsController() *InsightsController {
    return &c.insightsController
}

// InvoicesController returns the invoicesController instance of the client.
func (c *client) InvoicesController() *InvoicesController {
    return &c.invoicesController
}

// OffersController returns the offersController instance of the client.
func (c *client) OffersController() *OffersController {
    return &c.offersController
}

// PaymentProfilesController returns the paymentProfilesController instance of the client.
func (c *client) PaymentProfilesController() *PaymentProfilesController {
    return &c.paymentProfilesController
}

// ProductFamiliesController returns the productFamiliesController instance of the client.
func (c *client) ProductFamiliesController() *ProductFamiliesController {
    return &c.productFamiliesController
}

// ProductsController returns the productsController instance of the client.
func (c *client) ProductsController() *ProductsController {
    return &c.productsController
}

// ProductPricePointsController returns the productPricePointsController instance of the client.
func (c *client) ProductPricePointsController() *ProductPricePointsController {
    return &c.productPricePointsController
}

// ProformaInvoicesController returns the proformaInvoicesController instance of the client.
func (c *client) ProformaInvoicesController() *ProformaInvoicesController {
    return &c.proformaInvoicesController
}

// ReasonCodesController returns the reasonCodesController instance of the client.
func (c *client) ReasonCodesController() *ReasonCodesController {
    return &c.reasonCodesController
}

// ReferralCodesController returns the referralCodesController instance of the client.
func (c *client) ReferralCodesController() *ReferralCodesController {
    return &c.referralCodesController
}

// SalesCommissionsController returns the salesCommissionsController instance of the client.
func (c *client) SalesCommissionsController() *SalesCommissionsController {
    return &c.salesCommissionsController
}

// SitesController returns the sitesController instance of the client.
func (c *client) SitesController() *SitesController {
    return &c.sitesController
}

// SubscriptionsController returns the subscriptionsController instance of the client.
func (c *client) SubscriptionsController() *SubscriptionsController {
    return &c.subscriptionsController
}

// SubscriptionComponentsController returns the subscriptionComponentsController instance of the client.
func (c *client) SubscriptionComponentsController() *SubscriptionComponentsController {
    return &c.subscriptionComponentsController
}

// SubscriptionGroupsController returns the subscriptionGroupsController instance of the client.
func (c *client) SubscriptionGroupsController() *SubscriptionGroupsController {
    return &c.subscriptionGroupsController
}

// SubscriptionGroupInvoiceAccountController returns the subscriptionGroupInvoiceAccountController instance of the client.
func (c *client) SubscriptionGroupInvoiceAccountController() *SubscriptionGroupInvoiceAccountController {
    return &c.subscriptionGroupInvoiceAccountController
}

// SubscriptionGroupStatusController returns the subscriptionGroupStatusController instance of the client.
func (c *client) SubscriptionGroupStatusController() *SubscriptionGroupStatusController {
    return &c.subscriptionGroupStatusController
}

// SubscriptionInvoiceAccountController returns the subscriptionInvoiceAccountController instance of the client.
func (c *client) SubscriptionInvoiceAccountController() *SubscriptionInvoiceAccountController {
    return &c.subscriptionInvoiceAccountController
}

// SubscriptionNotesController returns the subscriptionNotesController instance of the client.
func (c *client) SubscriptionNotesController() *SubscriptionNotesController {
    return &c.subscriptionNotesController
}

// SubscriptionProductsController returns the subscriptionProductsController instance of the client.
func (c *client) SubscriptionProductsController() *SubscriptionProductsController {
    return &c.subscriptionProductsController
}

// SubscriptionStatusController returns the subscriptionStatusController instance of the client.
func (c *client) SubscriptionStatusController() *SubscriptionStatusController {
    return &c.subscriptionStatusController
}

// WebhooksController returns the webhooksController instance of the client.
func (c *client) WebhooksController() *WebhooksController {
    return &c.webhooksController
}

// UserAgent returns the userAgent instance of the client.
func (c *client) UserAgent() *string {
    return &c.userAgent
}

// GetCallBuilder returns the CallBuilderFactory used by the client.
func (c *client) GetCallBuilder() https.CallBuilderFactory {
    return c.callBuilderFactory
}

// getBaseUri returns the base URI based on the server and configuration.
func getBaseUri(
    server Server,
    configuration Configuration) string {
    if configuration.Environment() == Environment(US) {
        if server == Server(PRODUCTION) {
            return fmt.Sprintf("https://%v.chargify.com", configuration.Site())
        }
        if server == Server(EBB) {
            return fmt.Sprintf("https://events.chargify.com/%v", configuration.Site())
        }
    }
    if configuration.Environment() == Environment(EU) {
        if server == Server(PRODUCTION) {
            return fmt.Sprintf("https://%v.ebilling.maxio.com", configuration.Site())
        }
        if server == Server(EBB) {
            return fmt.Sprintf("https://events.chargify.com/%v", configuration.Site())
        }
    }
    return "TODO: Select a valid server."
}

// clientOptions is a function type representing options for the client.
type clientOptions func(cb https.CallBuilder)

// callBuilderHandler creates the call builder factory with various options.
func callBuilderHandler(
    baseUrlProvider func(server string) string,
    auth map[string]https.AuthInterface,
    httpClient https.HttpClient,
    retryConfig RetryConfiguration,
    arraySerializationOption https.ArraySerializationOption,
    opts ...clientOptions) https.CallBuilderFactory {
    callBuilderFactory := https.CreateCallBuilderFactory(baseUrlProvider, auth, httpClient, retryConfig, arraySerializationOption)
    return tap(callBuilderFactory, opts...)
}

// tap is a utility function to apply client options to the call builder factory.
func tap(
    callBuilderFactory https.CallBuilderFactory,
    opts ...clientOptions) https.CallBuilderFactory {
    return func(ctx context.Context, httpMethod, path string) https.CallBuilder {
    	callBuilder := callBuilderFactory(ctx, httpMethod, path)
    	for _, opt := range opts {
    		opt(callBuilder)
    	}
    	return callBuilder
    }
}

// withUserAgent is an option to add a user agent header to the HTTP request.
func withUserAgent(userAgent string) clientOptions {
    f := func(request *http.Request) *http.Request {
    	request.Header.Set("user-agent", userAgent)
    	return request
    }
    return func(cb https.CallBuilder) {
    	cb.InterceptRequest(f)
    }
}

// withGlobalErrors will add all globally defined errors to callBuilder.
func withGlobalErrors() clientOptions {
    return func(cb https.CallBuilder) {
        cb.AppendErrors(map[string]https.ErrorBuilder[error]{
            "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
            "0": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'."},
        })
    }
}
