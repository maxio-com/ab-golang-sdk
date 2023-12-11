package ab_golang_sdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"maxioadvancedbilling/errors"
	"maxioadvancedbilling/models"
)

// APIExportsController represents a controller struct.
type APIExportsController struct {
	baseController
}

// NewAPIExportsController creates a new instance of APIExportsController.
// It takes a baseController as a parameter and returns a pointer to the APIExportsController.
func NewAPIExportsController(baseController baseController) *APIExportsController {
	APIExportsController := APIExportsController{baseController: baseController}
	return &APIExportsController
}

// ListExportedProformaInvoices takes context, batchId, perPage, page as parameters and
// returns an models.ApiResponse with []models.ProformaInvoice data and
// an error if there was an issue with the request or response.
// This API returns an array of exported proforma invoices for a provided `batch_id`. Pay close attention to pagination in order to control responses from the server.
// Example: `GET https://{subdomain}.chargify.com/api_exports/proforma_invoices/123/rows?per_page=10000&page=1`.
func (a *APIExportsController) ListExportedProformaInvoices(
	ctx context.Context,
	batchId string,
	perPage *int,
	page *int) (
	models.ApiResponse[[]models.ProformaInvoice],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api_exports/proforma_invoices/%v/rows.json", batchId),
	)
	req.Authenticate(true)
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.ProformaInvoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ProformaInvoice](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}

// ListExportedInvoices takes context, batchId, perPage, page as parameters and
// returns an models.ApiResponse with []models.Invoice data and
// an error if there was an issue with the request or response.
// This API returns an array of exported invoices for a provided `batch_id`. Pay close attention to pagination in order to control responses from the server.
// Example: `GET https://{subdomain}.chargify.com/api_exports/invoices/123/rows?per_page=10000&page=1`.
func (a *APIExportsController) ListExportedInvoices(
	ctx context.Context,
	batchId string,
	perPage *int,
	page *int) (
	models.ApiResponse[[]models.Invoice],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api_exports/invoices/%v/rows.json", batchId),
	)
	req.Authenticate(true)
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Invoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Invoice](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}

// ListExportedSubscriptions takes context, batchId, perPage, page as parameters and
// returns an models.ApiResponse with []models.Subscription data and
// an error if there was an issue with the request or response.
// This API returns an array of exported subscriptions for a provided `batch_id`. Pay close attention to pagination in order to control responses from the server.
// Example: `GET https://{subdomain}.chargify.com/api_exports/subscriptions/123/rows?per_page=200&page=1`.
func (a *APIExportsController) ListExportedSubscriptions(
	ctx context.Context,
	batchId string,
	perPage *int,
	page *int) (
	models.ApiResponse[[]models.Subscription],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api_exports/subscriptions/%v/rows.json", batchId),
	)
	req.Authenticate(true)
	if perPage != nil {
		req.QueryParam("per_page", *perPage)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Subscription
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Subscription](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}

// ExportProformaInvoices takes context as parameters and
// returns an models.ApiResponse with models.BatchJobResponse data and
// an error if there was an issue with the request or response.
// This API creates a proforma invoices export and returns a batchjob object.
// It is only available for Relationship Invoicing architecture.
func (a *APIExportsController) ExportProformaInvoices(ctx context.Context) (
	models.ApiResponse[models.BatchJobResponse],
	error) {
	req := a.prepareRequest(ctx, "POST", "/api_exports/proforma_invoices.json")
	req.Authenticate(true)
	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	if resp.StatusCode == 409 {
		err = errors.NewSingleErrorResponseError(409, "Conflict")
	}
	return models.NewApiResponse(result, resp), err
}

// ExportInvoices takes context as parameters and
// returns an models.ApiResponse with models.BatchJobResponse data and
// an error if there was an issue with the request or response.
// This API creates an invoices export and returns a batchjob object.
func (a *APIExportsController) ExportInvoices(ctx context.Context) (
	models.ApiResponse[models.BatchJobResponse],
	error) {
	req := a.prepareRequest(ctx, "POST", "/api_exports/invoices.json")
	req.Authenticate(true)
	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	if resp.StatusCode == 409 {
		err = errors.NewSingleErrorResponseError(409, "Conflict")
	}
	return models.NewApiResponse(result, resp), err
}

// ExportSubscriptions takes context as parameters and
// returns an models.ApiResponse with models.BatchJobResponse data and
// an error if there was an issue with the request or response.
// This API creates a subscriptions export and returns a batchjob object.
func (a *APIExportsController) ExportSubscriptions(ctx context.Context) (
	models.ApiResponse[models.BatchJobResponse],
	error) {
	req := a.prepareRequest(ctx, "POST", "/api_exports/subscriptions.json")
	req.Authenticate(true)
	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 409 {
		err = errors.NewSingleErrorResponseError(409, "Conflict")
	}
	return models.NewApiResponse(result, resp), err
}

// ReadProformaInvoicesExport takes context, batchId as parameters and
// returns an models.ApiResponse with models.BatchJobResponse data and
// an error if there was an issue with the request or response.
// This API returns a batchjob object for proforma invoices export.
func (a *APIExportsController) ReadProformaInvoicesExport(
	ctx context.Context,
	batchId string) (
	models.ApiResponse[models.BatchJobResponse],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api_exports/proforma_invoices/%v.json", batchId),
	)
	req.Authenticate(true)

	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}

// ReadInvoicesExport takes context, batchId as parameters and
// returns an models.ApiResponse with models.BatchJobResponse data and
// an error if there was an issue with the request or response.
// This API returns a batchjob object for invoices export.
func (a *APIExportsController) ReadInvoicesExport(
	ctx context.Context,
	batchId string) (
	models.ApiResponse[models.BatchJobResponse],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api_exports/invoices/%v.json", batchId),
	)
	req.Authenticate(true)

	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}

// ReadSubscriptionsExport takes context, batchId as parameters and
// returns an models.ApiResponse with models.BatchJobResponse data and
// an error if there was an issue with the request or response.
// This API returns a batchjob object for subscriptions export.
func (a *APIExportsController) ReadSubscriptionsExport(
	ctx context.Context,
	batchId string) (
	models.ApiResponse[models.BatchJobResponse],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api_exports/subscriptions/%v.json", batchId),
	)
	req.Authenticate(true)

	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	if resp.StatusCode == 404 {
		err = errors.NewApiError(404, "Not Found")
	}
	return models.NewApiResponse(result, resp), err
}
