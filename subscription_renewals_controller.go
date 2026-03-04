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

// SubscriptionRenewalsController represents a controller struct.
type SubscriptionRenewalsController struct {
    baseController
}

// NewSubscriptionRenewalsController creates a new instance of SubscriptionRenewalsController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionRenewalsController.
func NewSubscriptionRenewalsController(baseController baseController) *SubscriptionRenewalsController {
    subscriptionRenewalsController := SubscriptionRenewalsController{baseController: baseController}
    return &subscriptionRenewalsController
}

// CreateScheduledRenewalConfiguration takes context, subscriptionId, body as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationResponse data and
// an error if there was an issue with the request or response.
// Creates a scheduled renewal configuration for a subscription. The scheduled renewal is based on the subscription’s current product and component setup.
func (s *SubscriptionRenewalsController) CreateScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    body *models.ScheduledRenewalConfigurationRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error) {
    req := s.prepareRequest(ctx, "POST", "/subscriptions/%v/scheduled_renewals.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ScheduledRenewalConfigurationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListScheduledRenewalConfigurations takes context, subscriptionId, status as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationsResponse data and
// an error if there was an issue with the request or response.
// Lists scheduled renewal configurations for the subscription and permits an optional status query filter.
func (s *SubscriptionRenewalsController) ListScheduledRenewalConfigurations(
    ctx context.Context,
    subscriptionId int,
    status *models.Status) (
    models.ApiResponse[models.ScheduledRenewalConfigurationsResponse],
    error) {
    req := s.prepareRequest(ctx, "GET", "/subscriptions/%v/scheduled_renewals.json")
    req.AppendTemplateParams(subscriptionId)
    req.Authenticate(NewAuth("BasicAuth"))
    if status != nil {
        req.QueryParam("status", *status)
    }
    
    var result models.ScheduledRenewalConfigurationsResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationsResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadScheduledRenewalConfiguration takes context, subscriptionId, id as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationResponse data and
// an error if there was an issue with the request or response.
// Retrieves the configuration settings for the scheduled renewal.
func (s *SubscriptionRenewalsController) ReadScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    id int) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      "/subscriptions/%v/scheduled_renewals/%v.json",
    )
    req.AppendTemplateParams(subscriptionId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.ScheduledRenewalConfigurationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateScheduledRenewalConfiguration takes context, subscriptionId, id, body as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationResponse data and
// an error if there was an issue with the request or response.
// Updates an existing configuration.
func (s *SubscriptionRenewalsController) UpdateScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    id int,
    body *models.ScheduledRenewalConfigurationRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      "/subscriptions/%v/scheduled_renewals/%v.json",
    )
    req.AppendTemplateParams(subscriptionId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ScheduledRenewalConfigurationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ScheduleScheduledRenewalLockIn takes context, subscriptionId, id, body as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationResponse data and
// an error if there was an issue with the request or response.
// Schedules a future lock-in date for the renewal.
func (s *SubscriptionRenewalsController) ScheduleScheduledRenewalLockIn(
    ctx context.Context,
    subscriptionId int,
    id int,
    body *models.ScheduledRenewalLockInRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      "/subscriptions/%v/scheduled_renewals/%v/schedule_lock_in.json",
    )
    req.AppendTemplateParams(subscriptionId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ScheduledRenewalConfigurationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// LockInScheduledRenewalImmediately takes context, subscriptionId, id as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationResponse data and
// an error if there was an issue with the request or response.
// Locks in the renewal immediately.
func (s *SubscriptionRenewalsController) LockInScheduledRenewalImmediately(
    ctx context.Context,
    subscriptionId int,
    id int) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      "/subscriptions/%v/scheduled_renewals/%v/immediate_lock_in.json",
    )
    req.AppendTemplateParams(subscriptionId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.ScheduledRenewalConfigurationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UnpublishScheduledRenewalConfiguration takes context, subscriptionId, id as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationResponse data and
// an error if there was an issue with the request or response.
// Returns a scheduled renewal configuration to an editable state.
func (s *SubscriptionRenewalsController) UnpublishScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    id int) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      "/subscriptions/%v/scheduled_renewals/%v/unpublish.json",
    )
    req.AppendTemplateParams(subscriptionId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.ScheduledRenewalConfigurationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CancelScheduledRenewalConfiguration takes context, subscriptionId, id as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationResponse data and
// an error if there was an issue with the request or response.
// Cancels a scheduled renewal configuration.
func (s *SubscriptionRenewalsController) CancelScheduledRenewalConfiguration(
    ctx context.Context,
    subscriptionId int,
    id int) (
    models.ApiResponse[models.ScheduledRenewalConfigurationResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      "/subscriptions/%v/scheduled_renewals/%v/cancel.json",
    )
    req.AppendTemplateParams(subscriptionId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.ScheduledRenewalConfigurationResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateScheduledRenewalConfigurationItem takes context, subscriptionId, scheduledRenewalsConfigurationId, body as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationItemResponse data and
// an error if there was an issue with the request or response.
// Adds product and component line items to the scheduled renewal.
func (s *SubscriptionRenewalsController) CreateScheduledRenewalConfigurationItem(
    ctx context.Context,
    subscriptionId int,
    scheduledRenewalsConfigurationId int,
    body *models.ScheduledRenewalConfigurationItemRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationItemResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      "/subscriptions/%v/scheduled_renewals/%v/configuration_items.json",
    )
    req.AppendTemplateParams(subscriptionId, scheduledRenewalsConfigurationId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ScheduledRenewalConfigurationItemResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationItemResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateScheduledRenewalConfigurationItem takes context, subscriptionId, scheduledRenewalsConfigurationId, id, body as parameters and
// returns an models.ApiResponse with models.ScheduledRenewalConfigurationItemResponse data and
// an error if there was an issue with the request or response.
// Updates an existing configuration item’s pricing and quantity.
func (s *SubscriptionRenewalsController) UpdateScheduledRenewalConfigurationItem(
    ctx context.Context,
    subscriptionId int,
    scheduledRenewalsConfigurationId int,
    id int,
    body *models.ScheduledRenewalUpdateRequest) (
    models.ApiResponse[models.ScheduledRenewalConfigurationItemResponse],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      "/subscriptions/%v/scheduled_renewals/%v/configuration_items/%v.json",
    )
    req.AppendTemplateParams(subscriptionId, scheduledRenewalsConfigurationId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorListResponse},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ScheduledRenewalConfigurationItemResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ScheduledRenewalConfigurationItemResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteScheduledRenewalConfigurationItem takes context, subscriptionId, scheduledRenewalsConfigurationId, id as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Removes an item from the pending renewal configuration.
func (s *SubscriptionRenewalsController) DeleteScheduledRenewalConfigurationItem(
    ctx context.Context,
    subscriptionId int,
    scheduledRenewalsConfigurationId int,
    id int) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      "/subscriptions/%v/scheduled_renewals/%v/configuration_items/%v.json",
    )
    req.AppendTemplateParams(subscriptionId, scheduledRenewalsConfigurationId, id)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {Message: "Unprocessable Entity (WebDAV)", Unmarshaller: errors.NewErrorListResponse},
    })
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
