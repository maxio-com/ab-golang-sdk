package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// SalesCommissionsController represents a controller struct.
type SalesCommissionsController struct {
    baseController
}

// NewSalesCommissionsController creates a new instance of SalesCommissionsController.
// It takes a baseController as a parameter and returns a pointer to the SalesCommissionsController.
func NewSalesCommissionsController(baseController baseController) *SalesCommissionsController {
    salesCommissionsController := SalesCommissionsController{baseController: baseController}
    return &salesCommissionsController
}

// ListSalesCommissionSettingsInput represents the input of the ListSalesCommissionSettings endpoint.
type ListSalesCommissionSettingsInput struct {
    // The Chargify id of your seller account
    SellerId      string  
    // For authorization use user API key. See details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).
    Authorization *string 
    // This parameter indicates if records should be fetched from live mode sites. Default value is true.
    LiveMode      *bool   
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page          *int    
    // This parameter indicates how many records to fetch in each request. Default value is 100.
    PerPage       *int    
}

// ListSalesCommissionSettings takes context, sellerId, authorization, liveMode, page, perPage as parameters and
// returns an models.ApiResponse with []models.SaleRepSettings data and
// an error if there was an issue with the request or response.
// Endpoint returns subscriptions with associated sales reps
// ## Modified Authentication Process
// The Sales Commission API differs from other Chargify API endpoints. This resource is associated with the seller itself. Up to now all available resources were at the level of the site, therefore creating the API Key per site was a sufficient solution. To share resources at the seller level, a new authentication method was introduced, which is user authentication. Creating an API Key for a user is a required step to correctly use the Sales Commission API, more details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).
// Access to the Sales Commission API endpoints is available to users with financial access, where the seller has the Advanced Analytics component enabled. For further information on getting access to Advanced Analytics please contact Chargify support.
// > Note: The request is at seller level, it means `<<subdomain>>` variable will be replaced by `app`
func (s *SalesCommissionsController) ListSalesCommissionSettings(
    ctx context.Context,
    input ListSalesCommissionSettingsInput) (
    models.ApiResponse[[]models.SaleRepSettings],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/sellers/%v/sales_commission_settings.json", input.SellerId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Authorization != nil {
        req.Header("Authorization", *input.Authorization)
    }
    if input.LiveMode != nil {
        req.QueryParam("live_mode", *input.LiveMode)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    
    var result []models.SaleRepSettings
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SaleRepSettings](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSalesRepsInput represents the input of the ListSalesReps endpoint.
type ListSalesRepsInput struct {
    // The Chargify id of your seller account
    SellerId      string  
    // For authorization use user API key. See details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).
    Authorization *string 
    // This parameter indicates if records should be fetched from live mode sites. Default value is true.
    LiveMode      *bool   
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page          *int    
    // This parameter indicates how many records to fetch in each request. Default value is 100.
    PerPage       *int    
}

// ListSalesReps takes context, sellerId, authorization, liveMode, page, perPage as parameters and
// returns an models.ApiResponse with []models.ListSaleRepItem data and
// an error if there was an issue with the request or response.
// Endpoint returns sales rep list with details
// ## Modified Authentication Process
// The Sales Commission API differs from other Chargify API endpoints. This resource is associated with the seller itself. Up to now all available resources were at the level of the site, therefore creating the API Key per site was a sufficient solution. To share resources at the seller level, a new authentication method was introduced, which is user authentication. Creating an API Key for a user is a required step to correctly use the Sales Commission API, more details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).
// Access to the Sales Commission API endpoints is available to users with financial access, where the seller has the Advanced Analytics component enabled. For further information on getting access to Advanced Analytics please contact Chargify support.
// > Note: The request is at seller level, it means `<<subdomain>>` variable will be replaced by `app`
func (s *SalesCommissionsController) ListSalesReps(
    ctx context.Context,
    input ListSalesRepsInput) (
    models.ApiResponse[[]models.ListSaleRepItem],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/sellers/%v/sales_reps.json", input.SellerId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Authorization != nil {
        req.Header("Authorization", *input.Authorization)
    }
    if input.LiveMode != nil {
        req.QueryParam("live_mode", *input.LiveMode)
    }
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    
    var result []models.ListSaleRepItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ListSaleRepItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadSalesRep takes context, sellerId, salesRepId, authorization, liveMode, page, perPage as parameters and
// returns an models.ApiResponse with models.SaleRep data and
// an error if there was an issue with the request or response.
// Endpoint returns sales rep and attached subscriptions details.
// ## Modified Authentication Process
// The Sales Commission API differs from other Chargify API endpoints. This resource is associated with the seller itself. Up to now all available resources were at the level of the site, therefore creating the API Key per site was a sufficient solution. To share resources at the seller level, a new authentication method was introduced, which is user authentication. Creating an API Key for a user is a required step to correctly use the Sales Commission API, more details [here](https://developers.chargify.com/docs/developer-docs/ZG9jOjMyNzk5NTg0-2020-04-20-new-api-authentication).
// Access to the Sales Commission API endpoints is available to users with financial access, where the seller has the Advanced Analytics component enabled. For further information on getting access to Advanced Analytics please contact Chargify support.
// > Note: The request is at seller level, it means `<<subdomain>>` variable will be replaced by `app`
func (s *SalesCommissionsController) ReadSalesRep(
    ctx context.Context,
    sellerId string,
    salesRepId string,
    authorization *string,
    liveMode *bool,
    page *int,
    perPage *int) (
    models.ApiResponse[models.SaleRep],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/sellers/%v/sales_reps/%v.json", sellerId, salesRepId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if authorization != nil {
        req.Header("Authorization", *authorization)
    }
    if liveMode != nil {
        req.QueryParam("live_mode", *liveMode)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    
    var result models.SaleRep
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SaleRep](decoder)
    return models.NewApiResponse(result, resp), err
}
