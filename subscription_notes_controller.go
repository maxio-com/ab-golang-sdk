package advancedbilling

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/utilities"
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
// Full documentation on how to use Notes in the Chargify UI can be located [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404434903181-Subscription-Summary#notes).
func (s *SubscriptionNotesController) CreateSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    body *models.UpdateSubscriptionNoteRequest) (
    models.ApiResponse[models.SubscriptionNoteResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/subscriptions/%v/notes.json", subscriptionId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.SubscriptionNoteResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SubscriptionNoteResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSubscriptionNotes takes context, subscriptionId, page, perPage as parameters and
// returns an models.ApiResponse with []models.SubscriptionNoteResponse data and
// an error if there was an issue with the request or response.
// Use this method to retrieve a list of Notes associated with a Subscription. The response will be an array of Notes.
func (s *SubscriptionNotesController) ListSubscriptionNotes(
    ctx context.Context,
    subscriptionId int,
    page *int,
    perPage *int) (
    models.ApiResponse[[]models.SubscriptionNoteResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/subscriptions/%v/notes.json", subscriptionId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
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
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/subscriptions/%v/notes/%v.json", subscriptionId, noteId),
    )
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
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/subscriptions/%v/notes/%v.json", subscriptionId, noteId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
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
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Use the following method to delete a note for a Subscription.
func (s *SubscriptionNotesController) DeleteSubscriptionNote(
    ctx context.Context,
    subscriptionId int,
    noteId int) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/subscriptions/%v/notes/%v.json", subscriptionId, noteId),
    )
    req.Authenticate(NewAuth("BasicAuth"))
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}
