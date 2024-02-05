package advancedbilling

import (
    "context"
    "fmt"
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
      fmt.Sprintf("/components/%v/price_points/%v/segments.json", componentId, pricePointId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.SegmentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SegmentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewEventBasedBillingSegmentErrors(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// ListSegmentsForPricePoint takes context, componentId, pricePointId, page, perPage, filterSegmentProperty1Value, filterSegmentProperty2Value, filterSegmentProperty3Value, filterSegmentProperty4Value as parameters and
// returns an models.ApiResponse with models.ListSegmentsResponse data and
// an error if there was an issue with the request or response.
// This endpoint allows you to fetch Segments created for a given Price Point. They will be returned in the order of creation.
// You can pass `page` and `per_page` parameters in order to access all of the segments. By default it will return `30` records. You can set `per_page` to `200` at most.
// You may specify component and/or price point by using either the numeric ID or the `handle:gold` syntax.
func (e *EventsBasedBillingSegmentsController) ListSegmentsForPricePoint(
    ctx context.Context,
    componentId string,
    pricePointId string,
    page *int,
    perPage *int,
    filterSegmentProperty1Value *string,
    filterSegmentProperty2Value *string,
    filterSegmentProperty3Value *string,
    filterSegmentProperty4Value *string) (
    models.ApiResponse[models.ListSegmentsResponse],
    error) {
    req := e.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/components/%v/price_points/%v/segments.json", componentId, pricePointId),
    )
    req.Authenticate(true)
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if filterSegmentProperty1Value != nil {
        req.QueryParam("filter[segment_property_1_value]", *filterSegmentProperty1Value)
    }
    if filterSegmentProperty2Value != nil {
        req.QueryParam("filter[segment_property_2_value]", *filterSegmentProperty2Value)
    }
    if filterSegmentProperty3Value != nil {
        req.QueryParam("filter[segment_property_3_value]", *filterSegmentProperty3Value)
    }
    if filterSegmentProperty4Value != nil {
        req.QueryParam("filter[segment_property_4_value]", *filterSegmentProperty4Value)
    }
    
    var result models.ListSegmentsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListSegmentsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewEventBasedBillingListSegmentsErrors(422, "Unprocessable Entity (WebDAV)")
    }
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
      fmt.Sprintf("/components/%v/price_points/%v/segments/%v.json", componentId, pricePointId, id),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.SegmentResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SegmentResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewEventBasedBillingSegmentErrors(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}

// DeleteSegment takes context, componentId, pricePointId, id as parameters and
// returns an models.ApiResponse with  data and
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
      fmt.Sprintf("/components/%v/price_points/%v/segments/%v.json", componentId, pricePointId, id),
    )
    req.Authenticate(true)
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    if context.Response.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if context.Response.StatusCode == 422 {
        err = errors.NewApiError(422, "Unprocessable Entity (WebDAV)")
    }
    return context.Response, err
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
      fmt.Sprintf("/components/%v/price_points/%v/segments/bulk.json", componentId, pricePointId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ListSegmentsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListSegmentsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewEventBasedBillingSegment(422, "Unprocessable Entity (WebDAV)")
    }
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
      fmt.Sprintf("/components/%v/price_points/%v/segments/bulk.json", componentId, pricePointId),
    )
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    
    var result models.ListSegmentsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    err = validateResponse(*resp)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ListSegmentsResponse](decoder)
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    if resp.StatusCode == 404 {
        err = errors.NewApiError(404, "Not Found")
    }
    if resp.StatusCode == 422 {
        err = errors.NewEventBasedBillingSegment(422, "Unprocessable Entity (WebDAV)")
    }
    return models.NewApiResponse(result, resp), err
}
