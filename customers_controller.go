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
// If provided, the `reference` value must be unique. It represents a unique identifier for the customer from your own app, i.e. the customer’s ID. This allows you to retrieve a given customer via a piece of shared information. Alternatively, you may choose to leave `reference` blank, and store Advanced Billing’s unique ID for the customer, which is in the `id` attribute.
// Full documentation on how to locate, create and edit Customers in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24252190590093-Customer-Details).
// ## Required Country Format
// Advanced Billing requires that you use the ISO Standard Country codes when formatting country attribute of the customer.
// Countries should be formatted as 2 characters. For more information, please see the following wikipedia article on [ISO_3166-1.](http://en.wikipedia.org/wiki/ISO_3166-1#Current_codes)
// ## Required State Format
// Advanced Billing requires that you use the ISO Standard State codes when formatting state attribute of the customer.
// + US States (2 characters): [ISO_3166-2](https://en.wikipedia.org/wiki/ISO_3166-2:US)
// + States Outside the US (2-3 characters): To find the correct state codes outside of the US, please go to [ISO_3166-1](http://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) and click on the link in the “ISO 3166-2 codes” column next to country you wish to populate.
// ## Locale
// Advanced Billing allows you to attribute a language/region to your customer to deliver invoices in any required language.
// For more: [Customer Locale](https://maxio.zendesk.com/hc/en-us/articles/24286672013709-Customer-Locale)
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
        req.Json(body)
    }
    var result models.CustomerResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CustomerResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListCustomersInput represents the input of the ListCustomers endpoint.
type ListCustomersInput struct {
    // Direction to sort customers by time of creation
    Direction     *models.SortingDirection 
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page          *int                     
    // This parameter indicates how many records to fetch in each request. Default value is 50. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage       *int                     
    // The type of filter you would like to apply to your search.
    // Use in query: `date_field=created_at`.
    DateField     *models.BasicDateField   
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
    StartDate     *string                  
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns subscriptions with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
    EndDate       *string                  
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
    StartDatetime *string                  
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns subscriptions with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.
    EndDatetime   *string                  
    // A search query by which to filter customers (can be an email, an ID, a reference, organization)
    Q             *string                  
}

// ListCustomers takes context, direction, page, perPage, dateField, startDate, endDate, startDatetime, endDatetime, q as parameters and
// returns an models.ApiResponse with []models.CustomerResponse data and
// an error if there was an issue with the request or response.
// This request will by default list all customers associated with your Site.
// ## Find Customer
// Use the search feature with the `q` query parameter to retrieve an array of customers that matches the search query.
// Common use cases are:
// + Search by an email
// + Search by an Advanced Billing ID
// + Search by an organization
// + Search by a reference value from your application
// + Search by a first or last name
// To retrieve a single, exact match by reference, please use the [lookup endpoint](https://developers.chargify.com/docs/api-docs/b710d8fbef104-read-customer-by-reference).
func (c *CustomersController) ListCustomers(
    ctx context.Context,
    input ListCustomersInput) (
    models.ApiResponse[[]models.CustomerResponse],
    error) {
    req := c.prepareRequest(ctx, "GET", "/customers.json")
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Direction != nil {
        req.QueryParam("direction", *input.Direction)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.DateField != nil {
        req.QueryParam("date_field", *input.DateField)
    }
    if input.StartDate != nil {
        req.QueryParam("start_date", *input.StartDate)
    }
    if input.EndDate != nil {
        req.QueryParam("end_date", *input.EndDate)
    }
    if input.StartDatetime != nil {
        req.QueryParam("start_datetime", *input.StartDatetime)
    }
    if input.EndDatetime != nil {
        req.QueryParam("end_datetime", *input.EndDatetime)
    }
    if input.Q != nil {
        req.QueryParam("q", *input.Q)
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
// This method allows to retrieve the Customer properties by Advanced Billing-generated Customer ID.
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
        req.Json(body)
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
