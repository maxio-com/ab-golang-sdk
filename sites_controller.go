// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
)

// SitesController represents a controller struct.
type SitesController struct {
    baseController
}

// NewSitesController creates a new instance of SitesController.
// It takes a baseController as a parameter and returns a pointer to the SitesController.
func NewSitesController(baseController baseController) *SitesController {
    sitesController := SitesController{baseController: baseController}
    return &sitesController
}

// ReadSite takes context as parameters and
// returns an models.ApiResponse with models.SiteResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to fetch some site data.
// Full documentation on Sites in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/sections/24250550707085-Sites).
// Specifically, the [Clearing Site Data](https://maxio.zendesk.com/hc/en-us/articles/24250617028365-Clearing-Site-Data) section is extremely relevant to this endpoint documentation.
// #### Relationship invoicing enabled
// If site has RI enabled then you will see more settings like:
// "customer_hierarchy_enabled": true,
// "whopays_enabled": true,
// "whopays_default_payer": "self"
// You can read more about these settings here:
// [Who Pays & Customer Hierarchy](https://maxio.zendesk.com/hc/en-us/articles/24252185211533-Customer-Hierarchies-WhoPays)
func (s *SitesController) ReadSite(ctx context.Context) (
    models.ApiResponse[models.SiteResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/site.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    var result models.SiteResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ClearSite takes context, cleanupScope as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This call is asynchronous and there may be a delay before the site data is fully deleted. If you are clearing site data for an automated test, you will need to build in a delay and/or check that there are no products, etc., in the site before proceeding.
// **This functionality will only work on sites in TEST mode. Attempts to perform this on sites in “live” mode will result in a response of 403 FORBIDDEN.**
func (s *SitesController) ClearSite(
    ctx context.Context,
    cleanupScope *models.CleanupScope) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/sites/clear_data.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    if cleanupScope != nil {
        req.QueryParam("cleanup_scope", *cleanupScope)
    }
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// ListChargifyJsPublicKeysInput represents the input of the ListChargifyJsPublicKeys endpoint.
type ListChargifyJsPublicKeysInput struct {
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page    *int 
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage *int 
}

// ListChargifyJsPublicKeys takes context, page, perPage as parameters and
// returns an models.ApiResponse with models.ListPublicKeysResponse data and
// an error if there was an issue with the request or response.
// This endpoint returns public keys used for Chargify.js.
func (s *SitesController) ListChargifyJsPublicKeys(
    ctx context.Context,
    input ListChargifyJsPublicKeysInput) (
    models.ApiResponse[models.ListPublicKeysResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/chargify_js_keys.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    var result models.ListPublicKeysResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListPublicKeysResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
