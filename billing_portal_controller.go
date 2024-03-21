package advancedbilling

import (
	"context"
	"fmt"
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
// ## Billing Portal Documentation
// Full documentation on how the Billing Portal operates within the Chargify UI can be located [here](https://chargify.zendesk.com/hc/en-us/articles/4407648972443).
// This documentation is focused on how the to configure the Billing Portal Settings, as well as Subscriber Interaction and Merchant Management of the Billing Portal.
// You can use this endpoint to enable Billing Portal access for a Customer, with the option of sending the Customer an Invitation email at the same time.
// ## Billing Portal Security
// If your customer has been invited to the Billing Portal, then they will receive a link to manage their subscription (the “Management URL”) automatically at the bottom of their statements, invoices, and receipts. **This link changes periodically for security and is only valid for 65 days.**
// If you need to provide your customer their Management URL through other means, you can retrieve it via the API. Because the URL is cryptographically signed with a timestamp, it is not possible for merchants to generate the URL without requesting it from Chargify.
// In order to prevent abuse & overuse, we ask that you request a new URL only when absolutely necessary. Management URLs are good for 65 days, so you should re-use a previously generated one as much as possible. If you use the URL frequently (such as to display on your website), please **do not** make an API request to Chargify every time.
func (b *BillingPortalController) EnableBillingPortalForCustomer(
	ctx context.Context,
	customerId int,
	autoInvite *models.AutoInvite) (
	models.ApiResponse[models.CustomerResponse],
	error) {
	req := b.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/portal/customers/%v/enable.json", customerId),
	)
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
// This method will provide to the API user the exact URL required for a subscriber to access the Billing Portal.
// ## Rules for Management Link API
// + When retrieving a management URL, multiple requests for the same customer in a short period will return the **same** URL
// + We will not generate a new URL for 15 days
// + You must cache and remember this URL if you are going to need it again within 15 days
// + Only request a new URL after the `new_link_available_at` date
// + You are limited to 15 requests for the same URL. If you make more than 15 requests before `new_link_available_at`, you will be blocked from further Management URL requests (with a response code `429`)
func (b *BillingPortalController) ReadBillingPortalLink(
	ctx context.Context,
	customerId int) (
	models.ApiResponse[models.PortalManagementLink],
	error) {
	req := b.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/portal/customers/%v/management_link.json", customerId),
	)
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
// You can resend a customer's Billing Portal invitation.
// If you attempt to resend an invitation 5 times within 30 minutes, you will receive a `422` response with `error` message in the body.
// If you attempt to resend an invitation when the Billing Portal is already disabled for a Customer, you will receive a `422` error response.
// If you attempt to resend an invitation when the Billing Portal is already disabled for a Customer, you will receive a `422` error response.
// If you attempt to resend an invitation when the Customer does not exist a Customer, you will receive a `404` error response.
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
		fmt.Sprintf("/portal/customers/%v/invitations/invite.json", customerId),
	)
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
// You can revoke a customer's Billing Portal invitation.
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
		fmt.Sprintf("/portal/customers/%v/invitations/revoke.json", customerId),
	)
	req.Authenticate(NewAuth("BasicAuth"))

	var result models.RevokedInvitation
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RevokedInvitation](decoder)
	return models.NewApiResponse(result, resp), err
}
