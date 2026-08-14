// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/maxio-com/ab-golang-sdk/errors"
    "github.com/maxio-com/ab-golang-sdk/models"
)

// BillingPortalController represents a controller struct.
type BillingPortalController struct {
    baseController
}

// NewBillingPortalController creates a new instance of BillingPortalController.
// It takes a baseController as a parameter and returns a pointer to the BillingPortalController.
func NewBillingPortalController(baseController baseController) *BillingPortalController {
    billingPortalController := BillingPortalController{baseController: baseController}
    return &billingPortalController
}

// EnableBillingPortalForCustomer takes context, customerId, autoInvite as parameters and
// returns an models.ApiResponse with models.CustomerResponse data and
// an error if there was an issue with the request or response.
// Enables Billing Portal access for a customer, with an option to send an invitation email at the same time.
// ## Billing Portal Security
// If your customer has been invited to the Billing Portal, they receive a link to manage their subscription (the “Management URL”) automatically at the bottom of their statements, invoices, and receipts. **This link changes periodically for security and is only valid for 65 days.**
// If you need to provide your customer their Management URL through other means, you can retrieve it [via the API]($e/Billing%20Portal/readBillingPortalLink). Because the URL is cryptographically signed with a timestamp, merchants cannot generate the URL without requesting it through the API.
// To prevent abuse and overuse, request a new URL only when absolutely necessary. Management URLs are good for 65 days, so you should re-use a previously generated one as much as possible. If you use the URL frequently (such as to display on your website), **do not** make an API request every time.
// For more information configuring the Billing Portal, see [Billing Portal Overview](https://maxio.zendesk.com/hc/en-us/articles/24252412965133-Billing-Portal-Overview).
func (b *BillingPortalController) EnableBillingPortalForCustomer(
    ctx context.Context,
    customerId int,
    autoInvite *models.AutoInvite) (
    models.ApiResponse[models.CustomerResponse],
    error) {
    req := b.prepareRequest(ctx, "POST", "/portal/customers/%v/enable.json")
    req.AppendTemplateParams(customerId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    if autoInvite != nil {
        req.QueryParam("auto_invite", *autoInvite)
    }
    
    var result models.CustomerResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.CustomerResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadBillingPortalLink takes context, customerId as parameters and
// returns an models.ApiResponse with models.PortalManagementLink data and
// an error if there was an issue with the request or response.
// Returns the exact URL required for a subscriber to access the Billing Portal.
// ## Management Link Request Rules
// + When retrieving a management URL, multiple requests for the same customer in a short period return the **same** URL
// + A new URL is not generated for 15 days
// + You must cache and remember this URL if you are going to need it again within 15 days
// + Only request a new URL after the `new_link_available_at` date
// + You are limited to 15 requests for the same URL. If you make more than 15 requests before `new_link_available_at`, you are blocked from further Management URL requests (with a response code `429`).
func (b *BillingPortalController) ReadBillingPortalLink(
    ctx context.Context,
    customerId int) (
    models.ApiResponse[models.PortalManagementLink],
    error) {
    req := b.prepareRequest(ctx, "GET", "/portal/customers/%v/management_link.json")
    req.AppendTemplateParams(customerId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
        "429": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewTooManyManagementLinkRequestsError},
    })
    
    var result models.PortalManagementLink
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PortalManagementLink](decoder)
    return models.NewApiResponse(result, resp), err
}

// ResendBillingPortalInvitation takes context, customerId as parameters and
// returns an models.ApiResponse with models.ResentInvitation data and
// an error if there was an issue with the request or response.
// Resends a customer's Billing Portal invitation.
// If you attempt to resend an invitation 5 times within 30 minutes, you will receive a `422` response with an `error` message in the body.
// If you attempt to resend an invitation when the Billing Portal is already disabled for a Customer, you will receive a `422` error response.
// If you attempt to resend an invitation when the Customer does not exist, you will receive a `404` error response.
// ## Limitations
// This endpoint will only return a JSON response.
func (b *BillingPortalController) ResendBillingPortalInvitation(
    ctx context.Context,
    customerId int) (
    models.ApiResponse[models.ResentInvitation],
    error) {
    req := b.prepareRequest(
      ctx,
      "POST",
      "/portal/customers/%v/invitations/invite.json",
    )
    req.AppendTemplateParams(customerId)
    req.Authenticate(NewAuth("BasicAuth"))
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "404": {TemplatedMessage: "Not Found:'{$response.body}'"},
        "422": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewErrorListResponse},
    })
    
    var result models.ResentInvitation
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResentInvitation](decoder)
    return models.NewApiResponse(result, resp), err
}

// RevokeBillingPortalAccess takes context, customerId as parameters and
// returns an models.ApiResponse with models.RevokedInvitation data and
// an error if there was an issue with the request or response.
// Revokes a customer's Billing Portal invitation.
// If you attempt to revoke an invitation when the Billing Portal is already disabled for a Customer, you will receive a 422 error response.
// ## Limitations
// This endpoint will only return a JSON response.
func (b *BillingPortalController) RevokeBillingPortalAccess(
    ctx context.Context,
    customerId int) (
    models.ApiResponse[models.RevokedInvitation],
    error) {
    req := b.prepareRequest(
      ctx,
      "DELETE",
      "/portal/customers/%v/invitations/revoke.json",
    )
    req.AppendTemplateParams(customerId)
    req.Authenticate(NewAuth("BasicAuth"))
    
    var result models.RevokedInvitation
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RevokedInvitation](decoder)
    return models.NewApiResponse(result, resp), err
}
