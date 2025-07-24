// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
    "net/http"
)

// EventsBasedBillingSegmentsController represents a controller struct.
type EventsBasedBillingSegmentsController struct {
    baseController
}

// NewEventsBasedBillingSegmentsController creates a new instance of EventsBasedBillingSegmentsController.
// It takes a baseController as a parameter and returns a pointer to the EventsBasedBillingSegmentsController.
func NewEventsBasedBillingSegmentsController(baseController baseController) *EventsBasedBillingSegmentsController {
    eventsBasedBillingSegmentsController := EventsBasedBillingSegmentsController{baseController: baseController}
    return &eventsBasedBillingSegmentsController
}

// CreateSegment takes context, componentId, pricePointId, body as parameters and
// returns an models.ApiResponse with models.SegmentResponse data and
// an error if there was an issue with the request or response.
// This endpoint creates a new Segment for a Component with segmented Metric. It allows you to specify properties to bill upon and prices for each Segment. You can only pass as many "property_values" as the related Metric has segmenting properties defined.
// You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.
func (e *EventsBasedBillingSegmentsController) CreateSegment(
    ctx context.Context,
    componentId string,
    pricePointId string,
    body *models.CreateSegmentRequest) (
    models.ApiResponse[models.SegmentResponse],
    error) {
    req := e.prepareRequest(
      ctx,
      "POST",
      "/components/%v/price_points/%v/segments.json",
    )
    req.AppendTemplateParams(componentId, pricePointId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewEventBasedBillingSegmentErrors},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SegmentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SegmentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSegmentsForPricePointInput represents the input of the ListSegmentsForPricePoint endpoint.
type ListSegmentsForPricePointInput struct {
    // ID or Handle for the Component
    ComponentId  string                     
    // ID or Handle for the Price Point belonging to the Component
    PricePointId string                     
    // Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
    // Use in query `page=1`.
    Page         *int                       
    // This parameter indicates how many records to fetch in each request. Default value is 30. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.
    // Use in query `per_page=200`.
    PerPage      *int                       
    // Filter to use for List Segments for a Price Point operation
    Filter       *models.ListSegmentsFilter 
}

// ListSegmentsForPricePoint takes context, componentId, pricePointId, page, perPage, filter as parameters and
// returns an models.ApiResponse with models.ListSegmentsResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to fetch Segments created for a given Price Point. They will be returned in the order of creation.
// You can pass `page` and `per_page` parameters in order to access all of the segments. By default it will return `30` records. You can set `per_page` to `200` at most.
// You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.
func (e *EventsBasedBillingSegmentsController) ListSegmentsForPricePoint(
    ctx context.Context,
    input ListSegmentsForPricePointInput) (
    models.ApiResponse[models.ListSegmentsResponse],
    error) {
    req := e.prepareRequest(
      ctx,
      "GET",
      "/components/%v/price_points/%v/segments.json",
    )
    req.AppendTemplateParams(input.ComponentId, input.PricePointId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewEventBasedBillingListSegmentsErrors},
    })
    if input.Page != nil {
        req.QueryParam("page", *input.Page)
    }
    if input.PerPage != nil {
        req.QueryParam("per_page", *input.PerPage)
    }
    if input.Filter != nil {
        req.QueryParam("filter", *input.Filter)
    }
    
    var result models.ListSegmentsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListSegmentsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSegment takes context, componentId, pricePointId, id, body as parameters and
// returns an models.ApiResponse with models.SegmentResponse data and
// an error if there was an issue with the request or response.
// This endpoint updates a single Segment for a Component with a segmented Metric. It allows you to update the pricing for the segment.
// You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.
func (e *EventsBasedBillingSegmentsController) UpdateSegment(
    ctx context.Context,
    componentId string,
    pricePointId string,
    id float64,
    body *models.UpdateSegmentRequest) (
    models.ApiResponse[models.SegmentResponse],
    error) {
    req := e.prepareRequest(
      ctx,
      "PUT",
      "/components/%v/price_points/%v/segments/%v.json",
    )
    req.AppendTemplateParams(componentId, pricePointId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewEventBasedBillingSegmentErrors},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SegmentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SegmentResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSegment takes context, componentId, pricePointId, id as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This endpoint allows you to delete a Segment with specified ID.
// You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.
func (e *EventsBasedBillingSegmentsController) DeleteSegment(
    ctx context.Context,
    componentId string,
    pricePointId string,
    id float64) (
    *http.Response,
    error) {
    req := e.prepareRequest(
      ctx,
      "DELETE",
      "/components/%v/price_points/%v/segments/%v.json",
    )
    req.AppendTemplateParams(componentId, pricePointId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'."},
    })
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// BulkCreateSegments takes context, componentId, pricePointId, body as parameters and
// returns an models.ApiResponse with models.ListSegmentsResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to create multiple segments in one request. The array of segments can contain up to `2000` records.
// If any of the records contain an error the whole request would fail and none of the requested segments get created. The error response contains a message for only the one segment that failed validation, with the corresponding index in the array.
// You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.
func (e *EventsBasedBillingSegmentsController) BulkCreateSegments(
    ctx context.Context,
    componentId string,
    pricePointId string,
    body *models.BulkCreateSegments) (
    models.ApiResponse[models.ListSegmentsResponse],
    error) {
    req := e.prepareRequest(
      ctx,
      "POST",
      "/components/%v/price_points/%v/segments/bulk.json",
    )
    req.AppendTemplateParams(componentId, pricePointId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewEventBasedBillingSegment},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ListSegmentsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListSegmentsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// BulkUpdateSegments takes context, componentId, pricePointId, body as parameters and
// returns an models.ApiResponse with models.ListSegmentsResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to update multiple segments in one request. The array of segments can contain up to `1000` records.
// If any of the records contain an error the whole request would fail and none of the requested segments get updated. The error response contains a message for only the one segment that failed validation, with the corresponding index in the array.
// You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.
func (e *EventsBasedBillingSegmentsController) BulkUpdateSegments(
    ctx context.Context,
    componentId string,
    pricePointId string,
    body *models.BulkUpdateSegments) (
    models.ApiResponse[models.ListSegmentsResponse],
    error) {
    req := e.prepareRequest(
      ctx,
      "PUT",
      "/components/%v/price_points/%v/segments/bulk.json",
    )
    req.AppendTemplateParams(componentId, pricePointId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewEventBasedBillingSegment},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ListSegmentsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListSegmentsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
