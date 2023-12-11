package ab_golang_sdk

var customersController CustomersController

var eventsController EventsController

var insightsController InsightsController

var offersController OffersController

var paymentProfilesController PaymentProfilesController

var reasonCodesController ReasonCodesController

var sitesController SitesController

var subscriptionsController SubscriptionsController

var subscriptionGroupsController SubscriptionGroupsController

var webhooksController WebhooksController

// init is an initialization function that sets up the controllers.
// It creates a configuration from the environment with a specified HTTP configuration and initializes the client.
// Then, it assigns the different controllers from the client to the corresponding variables for further use.
func init() {
	config := CreateConfigurationFromEnvironment(
		WithHttpConfiguration(
			CreateHttpConfiguration(
				WithTimeout(30),
			),
		),
	)

	client := NewClient(config)

	customersController = *client.CustomersController()
	eventsController = *client.EventsController()
	insightsController = *client.InsightsController()
	offersController = *client.OffersController()
	paymentProfilesController = *client.PaymentProfilesController()
	reasonCodesController = *client.ReasonCodesController()
	sitesController = *client.SitesController()
	subscriptionsController = *client.SubscriptionsController()
	subscriptionGroupsController = *client.SubscriptionGroupsController()
	webhooksController = *client.WebhooksController()
}
