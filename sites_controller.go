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
// Full documentation on Sites in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407870738587).
// Specifically, the [Clearing Site Data](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405428327309) section is extremely relevant to this endpoint documentation.
// #### Relationship invoicing enabled
// If site has RI enabled then you will see more settings like:
// "customer_hierarchy_enabled": true,
// "whopays_enabled": true,
// "whopays_default_payer": "self"
// You can read more about these settings here:
// [Who Pays & Customer Hierarchy](https://chargify.zendesk.com/hc/en-us/articles/4407746683291)
func (s *SitesController) ReadSite(ctx context.Context) (
	models.ApiResponse[models.SiteResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/site.json")
	req.Authenticate(true)
	var result models.SiteResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SiteResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// ClearSite takes context, cleanupScope as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This call is asynchronous and there may be a delay before the site data is fully deleted. If you are clearing site data for an automated test, you will need to build in a delay and/or check that there are no products, etc., in the site before proceeding.
// **This functionality will only work on sites in TEST mode. Attempts to perform this on sites in “live” mode will result in a response of 403 FORBIDDEN.**
func (s *SitesController) ClearSite(
	ctx context.Context,
	cleanupScope *models.CleanupScope) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/sites/clear_data.json")
	req.Authenticate(true)
	if cleanupScope != nil {
		req.QueryParam("cleanup_scope", *cleanupScope)
	}
	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// ListChargifyJsPublicKeys takes context, page, perPage as parameters and
// returns an models.ApiResponse with models.ListPublicKeysResponse data and
// an error if there was an issue with the request or response.
// This endpoint returns public keys used for Chargify.js.
func (s *SitesController) ListChargifyJsPublicKeys(
	ctx context.Context,
	page *int,
	perPage *int) (
	models.ApiResponse[models.ListPublicKeysResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/chargify_js_keys.json")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	var result models.ListPublicKeysResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListPublicKeysResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
