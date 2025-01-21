/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
)

// SubscriptionNotesController represents a controller struct.
type SubscriptionNotesController struct {
    baseController
}

// NewSubscriptionNotesController creates a new instance of SubscriptionNotesController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionNotesController.
func NewSubscriptionNotesController(baseController baseController) *SubscriptionNotesController {
    subscriptionNotesController := SubscriptionNotesController{baseController: baseController}
    return &subscriptionNotesController
}

// CreateSubscriptionNote takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionNoteResponse data and
// an error if there was an issue with the request or response.
// Use the following method to create a note for a subscription.
// ## How to Use Subscription Notes
// Notes allow you to record information about a particular Subscription in a free text format.
// If you have structured data such as birth date, color, etc., consider using Metadata instead.
// Full documentation on how to use Notes in the Advanced Billing UI can be located [here](https://maxio.zendesk.com/hc/en-us/articles/24251712214413-Subscription-Summary-Overview).
func (s *SubscriptionNotesController) CreateSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    body *models.UpdateSubscriptionNoteRequest) (
    models.ApiResponse[models.SubscriptionNoteResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscriptions/%v/notes.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SubscriptionNoteResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionNoteResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSubscriptionNotesInput represents the input of the ListSubscriptionNotes endpoint.
type ListSubscriptionNotesInput struct {
    // The Chargify id of the subscription
    SubscriptionId int  
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page           *int 
    // This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage        *int 
}

// ListSubscriptionNotes takes context, subscriptionId, page, perPage as parameters and
// returns an models.ApiResponse with []models.SubscriptionNoteResponse data and
// an error if there was an issue with the request or response.
// Use this method to retrieve a list of Notes associated with a Subscription. The response will be an array of Notes.
func (s *SubscriptionNotesController) ListSubscriptionNotes(
    ctx context.Context,
    input ListSubscriptionNotesInput) (
    models.ApiResponse[[]models.SubscriptionNoteResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/subscriptions/%v/notes.json")
    req.AppendTemplateParams(input.SubscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    
    var result []models.SubscriptionNoteResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SubscriptionNoteResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadSubscriptionNote takes context, subscriptionId, noteId as parameters and
// returns an models.ApiResponse with models.SubscriptionNoteResponse data and
// an error if there was an issue with the request or response.
// Once you have obtained the ID of the note you wish to read, use this method to show a particular note attached to a subscription.
func (s *SubscriptionNotesController) ReadSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    noteId int) (
    models.ApiResponse[models.SubscriptionNoteResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/subscriptions/%v/notes/%v.json")
    req.AppendTemplateParams(subscriptionId, noteId)
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.SubscriptionNoteResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionNoteResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionNote takes context, subscriptionId, noteId, body as parameters and
// returns an models.ApiResponse with models.SubscriptionNoteResponse data and
// an error if there was an issue with the request or response.
// Use the following method to update a note for a Subscription.
func (s *SubscriptionNotesController) UpdateSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    noteId int,
    body *models.UpdateSubscriptionNoteRequest) (
    models.ApiResponse[models.SubscriptionNoteResponse],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/subscriptions/%v/notes/%v.json")
    req.AppendTemplateParams(subscriptionId, noteId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SubscriptionNoteResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionNoteResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSubscriptionNote takes context, subscriptionId, noteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Use the following method to delete a note for a Subscription.
func (s *SubscriptionNotesController) DeleteSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    noteId int) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/subscriptions/%v/notes/%v.json")
    req.AppendTemplateParams(subscriptionId, noteId)
    req.Authenticate(NewAuth("BasicAuth"))
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
