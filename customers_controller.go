package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
)

// CustomersController represents a controller struct.
type CustomersController struct {
    baseController
}

// NewCustomersController creates a new instance of CustomersController.
// It takes a baseController as a parameter and returns a pointer to the CustomersController.
func NewCustomersController(baseController baseController) *CustomersController {
    customersController := CustomersController{baseController: baseController}
    return &customersController
}

// CreateCustomer takes context, body as parameters and
// returns an models.ApiResponse with models.CustomerResponse data and
// an error if there was an issue with the request or response.
// You may create a new Customer at any time, or you may create a Customer at the same time you create a Subscription. The only validation restriction is that you may only create one customer for a given reference value.
// If provided, the `reference` value must be unique. It represents a unique identifier for the customer from your own app, i.e. the customer’s ID. This allows you to retrieve a given customer via a piece of shared information. Alternatively, you may choose to leave `reference` blank, and store Chargify’s unique ID for the customer, which is in the `id` attribute.
// Full documentation on how to locate, create and edit Customers in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407659914267).
// ## Required Country Format
// Chargify requires that you use the ISO Standard Country codes when formatting country attribute of the customer.
// Countries should be formatted as 2 characters. For more information, please see the following wikipedia article on [ISO_3166-1.](http://en.wikipedia.org/wiki/ISO_3166-1#Current_codes)
// ## Required State Format
// Chargify requires that you use the ISO Standard State codes when formatting state attribute of the customer.
// + US States (2 characters): [ISO_3166-2](https://en.wikipedia.org/wiki/ISO_3166-2:US)
// + States Outside the US (2-3 characters): To find the correct state codes outside of the US, please go to [ISO_3166-1](http://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) and click on the link in the “ISO 3166-2 codes” column next to country you wish to populate.
// ## Locale
// Chargify allows you to attribute a language/region to your customer to deliver invoices in any required language.
// For more: [Customer Locale](https://chargify.zendesk.com/hc/en-us/articles/4407870384283#customer-locale)
func (c *CustomersController) CreateCustomer(
    ctx context.Context,
    body *models.CreateCustomerRequest) (
    models.ApiResponse[models.CustomerResponse],
    error) {
    req := c.prepareRequest(ctx, "POST", "/customers.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewCustomerErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    var result models.CustomerResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CustomerResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListCustomers takes context, direction, page, perPage, dateField, startDate, endDate, startDatetime, endDatetime, q as parameters and
// returns an models.ApiResponse with []models.CustomerResponse data and
// an error if there was an issue with the request or response.
// This request will by default list all customers associated with your Site.
// ## Find Customer
// Use the search feature with the `q` query parameter to retrieve an array of customers that matches the search query.
// Common use cases are:
// + Search by an email
// + Search by a Chargify ID
// + Search by an organization
// + Search by a reference value from your application
// + Search by a first or last name
// To retrieve a single, exact match by reference, please use the [lookup endpoint](https://developers.chargify.com/docs/api-docs/b710d8fbef104-read-customer-by-reference).
func (c *CustomersController) ListCustomers(
    ctx context.Context,
    direction *models.SortingDirection,
    page *int,
    perPage *int,
    dateField *models.BasicDateField,
    startDate *string,
    endDate *string,
    startDatetime *string,
    endDatetime *string,
    q *string) (
    models.ApiResponse[[]models.CustomerResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/customers.json")
    req.Authenticate(NewAuth("BasicAuth"))
    if direction != nil {
        req.QueryParam("direction", *direction)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if dateField != nil {
        req.QueryParam("date_field", *dateField)
    }
    if startDate != nil {
        req.QueryParam("start_date", *startDate)
    }
    if endDate != nil {
        req.QueryParam("end_date", *endDate)
    }
    if startDatetime != nil {
        req.QueryParam("start_datetime", *startDatetime)
    }
    if endDatetime != nil {
        req.QueryParam("end_datetime", *endDatetime)
    }
    if q != nil {
        req.QueryParam("q", *q)
    }
    var result []models.CustomerResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.CustomerResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadCustomer takes context, id as parameters and
// returns an models.ApiResponse with models.CustomerResponse data and
// an error if there was an issue with the request or response.
// This method allows to retrieve the Customer properties by Chargify-generated Customer ID.
func (c *CustomersController) ReadCustomer(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.CustomerResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", fmt.Sprintf("/customers/%v.json", id))
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.CustomerResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CustomerResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateCustomer takes context, id, body as parameters and
// returns an models.ApiResponse with models.CustomerResponse data and
// an error if there was an issue with the request or response.
// This method allows to update the Customer.
func (c *CustomersController) UpdateCustomer(
    ctx context.Context,
    id int,
    body *models.UpdateCustomerRequest) (
    models.ApiResponse[models.CustomerResponse],
    error) {
    req := c.prepareRequest(ctx, "PUT", fmt.Sprintf("/customers/%v.json", id))
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewCustomerErrorResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.CustomerResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CustomerResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteCustomer takes context, id as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This method allows you to delete the Customer.
func (c *CustomersController) DeleteCustomer(
    ctx context.Context,
    id int) (
    *http.Response,
    error) {
    req := c.prepareRequest(ctx, "DELETE", fmt.Sprintf("/customers/%v.json", id))
    req.Authenticate(NewAuth("BasicAuth"))
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// ReadCustomerByReference takes context, reference as parameters and
// returns an models.ApiResponse with models.CustomerResponse data and
// an error if there was an issue with the request or response.
// Use this method to return the customer object if you have the unique **Reference ID (Your App)** value handy. It will return a single match.
func (c *CustomersController) ReadCustomerByReference(
    ctx context.Context,
    reference string) (
    models.ApiResponse[models.CustomerResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/customers/lookup.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.QueryParam("reference", reference)
    var result models.CustomerResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CustomerResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListCustomerSubscriptions takes context, customerId as parameters and
// returns an models.ApiResponse with []models.SubscriptionResponse data and
// an error if there was an issue with the request or response.
// This method lists all subscriptions that belong to a customer.
func (c *CustomersController) ListCustomerSubscriptions(
    ctx context.Context,
    customerId int) (
    models.ApiResponse[[]models.SubscriptionResponse],
    error) {
    req := c.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/customers/%v/subscriptions.json", customerId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result []models.SubscriptionResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SubscriptionResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
