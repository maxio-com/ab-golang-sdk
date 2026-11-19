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
    "time"
)

// FeatureTemplatesController represents a controller struct.
type FeatureTemplatesController struct {
    baseController
}

// NewFeatureTemplatesController creates a new instance of FeatureTemplatesController.
// It takes a baseController as a parameter and returns a pointer to the FeatureTemplatesController.
func NewFeatureTemplatesController(baseController baseController) *FeatureTemplatesController {
    featureTemplatesController := FeatureTemplatesController{baseController: baseController}
    return &featureTemplatesController
}

// ListFeatureTemplates takes context, page, perPage, status, q, kind, updatedFrom, updatedTo, sortBy, sortDirection as parameters and
// returns an models.ApiResponse with models.FeatureTemplatesListResponse data and
// an error if there was an issue with the request or response.
// Lists the feature templates defined for your site, active (non-archived) ones by default. Pass `status=archived` or `status=all` to widen the result set.
// Supply `page` or `per_page` to paginate. Without either parameter, the response includes the full result set.
func (f *FeatureTemplatesController) ListFeatureTemplates(
    ctx context.Context,
    page *int,
    perPage *int,
    status *models.Status1,
    q *string,
    kind *models.Kind,
    updatedFrom *time.Time,
    updatedTo *time.Time,
    sortBy *models.SortBy,
    sortDirection *models.SortDirection) (
    models.ApiResponse[models.FeatureTemplatesListResponse],
    error) {
    req := f.prepareRequest(ctx, "GET", "/features.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
    })
    if page != nil {
        req.QueryParam("page", *page)
    }
    if perPage != nil {
        req.QueryParam("per_page", *perPage)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    if q != nil {
        req.QueryParam("q", *q)
    }
    if kind != nil {
        req.QueryParam("kind", *kind)
    }
    if updatedFrom != nil {
        req.QueryParam("updated_from", updatedFrom.Format(models.DEFAULT_DATE))
    }
    if updatedTo != nil {
        req.QueryParam("updated_to", updatedTo.Format(models.DEFAULT_DATE))
    }
    if sortBy != nil {
        req.QueryParam("sort_by", *sortBy)
    }
    if sortDirection != nil {
        req.QueryParam("sort_direction", *sortDirection)
    }
    var result models.FeatureTemplatesListResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureTemplatesListResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateFeatureTemplate takes context, body as parameters and
// returns an models.ApiResponse with models.FeatureTemplateResponse data and
// an error if there was an issue with the request or response.
// Defines a new feature at the site level. Feature templates aren't billable on their own. Attach a template to products or components to grant the feature to subscribers.
func (f *FeatureTemplatesController) CreateFeatureTemplate(
    ctx context.Context,
    body *models.CreateFeatureTemplateRequest) (
    models.ApiResponse[models.FeatureTemplateResponse],
    error) {
    req := f.prepareRequest(ctx, "POST", "/features.json")
    
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.FeatureTemplateResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureTemplateResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadFeatureTemplate takes context, id as parameters and
// returns an models.ApiResponse with models.FeatureTemplateResponse data and
// an error if there was an issue with the request or response.
// Returns a single feature template. Archived feature templates are not addressable here and return `404`. Restore a template first to read or update it.
func (f *FeatureTemplatesController) ReadFeatureTemplate(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.FeatureTemplateResponse],
    error) {
    req := f.prepareRequest(ctx, "GET", "/features/%v.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
    })
    
    var result models.FeatureTemplateResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureTemplateResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateFeatureTemplate takes context, id, body as parameters and
// returns an models.ApiResponse with models.FeatureTemplateResponse data and
// an error if there was an issue with the request or response.
// Updates the name, description, unit, value type, default value, or default periodicity of a feature template. `key` is rejected on every update. `kind` is rejected once any feature catalog item has been created from this template.
// Archived feature templates are not addressable here and return `404`. Restore a template first to update it.
func (f *FeatureTemplatesController) UpdateFeatureTemplate(
    ctx context.Context,
    id int,
    body *models.UpdateFeatureTemplateRequest) (
    models.ApiResponse[models.FeatureTemplateResponse],
    error) {
    req := f.prepareRequest(ctx, "PUT", "/features/%v.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.FeatureTemplateResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureTemplateResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ArchiveFeatureTemplate takes context, id, removeFromCatalog as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Archives a feature template. Archived feature templates are not addressable via [Read Feature Template]($e/Feature%20Templates/readFeatureTemplate) or [Update Feature Template]($e/Feature%20Templates/updateFeatureTemplate). Both endpoints return `404` until the template is restored.
// The feature template record itself is never hard-deleted, and can always be restored with [Restore Feature Template]($e/Feature%20Templates/restoreFeatureTemplate). Reversibility does not extend to `remove_from_catalog=true`: the feature catalog items and entitlements that parameter destroys are gone permanently, and restoring the template will not bring subscriber access back.
func (f *FeatureTemplatesController) ArchiveFeatureTemplate(
    ctx context.Context,
    id int,
    removeFromCatalog *bool) (
    *http.Response,
    error) {
    req := f.prepareRequest(ctx, "DELETE", "/features/%v.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
    })
    if removeFromCatalog != nil {
        req.QueryParam("remove_from_catalog", *removeFromCatalog)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// RestoreFeatureTemplate takes context, id as parameters and
// returns an models.ApiResponse with models.FeatureTemplateResponse data and
// an error if there was an issue with the request or response.
// Clears the feature template's archived state. Feature catalog items created from this template are not automatically restored. Restore each one individually.
func (f *FeatureTemplatesController) RestoreFeatureTemplate(
    ctx context.Context,
    id int) (
    models.ApiResponse[models.FeatureTemplateResponse],
    error) {
    req := f.prepareRequest(ctx, "POST", "/features/%v/restore.json")
    req.AppendTemplateParams(id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "403": {Message: "Forbidden", Unmarshaller: errors.NewErrorListResponse},
        "404": {Message: "Not Found"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.FeatureTemplateResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FeatureTemplateResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
