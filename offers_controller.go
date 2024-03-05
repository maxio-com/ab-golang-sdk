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

// OffersController represents a controller struct.
type OffersController struct {
    baseController
}

// NewOffersController creates a new instance of OffersController.
// It takes a baseController as a parameter and returns a pointer to the OffersController.
func NewOffersController(baseController baseController) *OffersController {
    offersController := OffersController{baseController: baseController}
    return &offersController
}

// CreateOffer takes context, body as parameters and
// returns an models.ApiResponse with models.OfferResponse data and
// an error if there was an issue with the request or response.
// Create an offer within your Chargify site by sending a POST request.
// ## Documentation
// Offers allow you to package complicated combinations of products, components and coupons into a convenient package which can then be subscribed to just like products.
// Once an offer is defined it can be used as an alternative to the product when creating subscriptions.
// Full documentation on how to use offers in the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407753852059).
// ## Using a Product Price Point
// You can optionally pass in a `product_price_point_id` that corresponds with the `product_id` and the offer will use that price point. If a `product_price_point_id` is not passed in, the product's default price point will be used.
func (o *OffersController) CreateOffer(
    ctx context.Context,
    body *models.CreateOfferRequest) (
    models.ApiResponse[models.OfferResponse],
    error) {
    req := o.prepareRequest(ctx, "POST", "/offers.json")
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorArrayMapResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    var result models.OfferResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OfferResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListOffers takes context, page, perPage, includeArchived as parameters and
// returns an models.ApiResponse with models.ListOffersResponse data and
// an error if there was an issue with the request or response.
// This endpoint will list offers for a site.
func (o *OffersController) ListOffers(
    ctx context.Context,
    page *int,
    perPage *int,
    includeArchived *bool) (
    models.ApiResponse[models.ListOffersResponse],
    error) {
    req := o.prepareRequest(ctx, "GET", "/offers.json")
    req.Authenticate(NewAuth("BasicAuth"))
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if includeArchived != nil {
        req.QueryParam("include_archived", *includeArchived)
    }
    var result models.ListOffersResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListOffersResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadOffer takes context, offerId as parameters and
// returns an models.ApiResponse with models.OfferResponse data and
// an error if there was an issue with the request or response.
// This method allows you to list a specific offer's attributes. This is different than list all offers for a site, as it requires an `offer_id`.
func (o *OffersController) ReadOffer(
    ctx context.Context,
    offerId int) (
    models.ApiResponse[models.OfferResponse],
    error) {
    req := o.prepareRequest(ctx, "GET", fmt.Sprintf("/offers/%v.json", offerId))
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.OfferResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OfferResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ArchiveOffer takes context, offerId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Archive an existing offer. Please provide an `offer_id` in order to archive the correct item.
func (o *OffersController) ArchiveOffer(
    ctx context.Context,
    offerId int) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/offers/%v/archive.json", offerId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// UnarchiveOffer takes context, offerId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Unarchive a previously archived offer. Please provide an `offer_id` in order to un-archive the correct item.
func (o *OffersController) UnarchiveOffer(
    ctx context.Context,
    offerId int) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/offers/%v/unarchive.json", offerId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}
