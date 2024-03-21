package advancedbilling

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/maxio-com/ab-golang-sdk/errors"
	"github.com/maxio-com/ab-golang-sdk/models"
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

// ListExportedProformaInvoicesInput represents the input of the ListExportedProformaInvoices endpoint.
type ListExportedProformaInvoicesInput struct {
	// Id of a Batch Job.
	BatchId string
	// This parameter indicates how many records to fetch in each request.
	// Default value is 100.
	// The maximum allowed values is 10000; any per_page value over 10000 will be changed to 10000.
	PerPage *int
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
}

// ListExportedProformaInvoices takes context, batchId, perPage, page as parameters and
// returns an models.ApiResponse with []models.ProformaInvoice data and
// an error if there was an issue with the request or response.
// This API returns an array of exported proforma invoices for a provided `batch_id`. Pay close attention to pagination in order to control responses from the server.
// Example: `GET https://{subdomain}.chargify.com/api_exports/proforma_invoices/123/rows?per_page=10000&page=1`.
func (a *APIExportsController) ListExportedProformaInvoices(
	ctx context.Context,
	input ListExportedProformaInvoicesInput) (
	models.ApiResponse[[]models.ProformaInvoice],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api_exports/proforma_invoices/%v/rows.json", input.BatchId),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}

	var result []models.ProformaInvoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ProformaInvoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListExportedInvoicesInput represents the input of the ListExportedInvoices endpoint.
type ListExportedInvoicesInput struct {
	// Id of a Batch Job.
	BatchId string
	// This parameter indicates how many records to fetch in each request.
	// Default value is 100.
	// The maximum allowed values is 10000; any per_page value over 10000 will be changed to 10000.
	PerPage *int
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
}

// ListExportedInvoices takes context, batchId, perPage, page as parameters and
// returns an models.ApiResponse with []models.Invoice data and
// an error if there was an issue with the request or response.
// This API returns an array of exported invoices for a provided `batch_id`. Pay close attention to pagination in order to control responses from the server.
// Example: `GET https://{subdomain}.chargify.com/api_exports/invoices/123/rows?per_page=10000&page=1`.
func (a *APIExportsController) ListExportedInvoices(
	ctx context.Context,
	input ListExportedInvoicesInput) (
	models.ApiResponse[[]models.Invoice],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api_exports/invoices/%v/rows.json", input.BatchId),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}

	var result []models.Invoice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Invoice](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListExportedSubscriptionsInput represents the input of the ListExportedSubscriptions endpoint.
type ListExportedSubscriptionsInput struct {
	// Id of a Batch Job.
	BatchId string
	// This parameter indicates how many records to fetch in each request.
	// Default value is 100.
	// The maximum allowed values is 10000; any per_page value over 10000 will be changed to 10000.
	PerPage *int
	// Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.
	// Use in query `page=1`.
	Page *int
}

// ListExportedSubscriptions takes context, batchId, perPage, page as parameters and
// returns an models.ApiResponse with []models.Subscription data and
// an error if there was an issue with the request or response.
// This API returns an array of exported subscriptions for a provided `batch_id`. Pay close attention to pagination in order to control responses from the server.
// Example: `GET https://{subdomain}.chargify.com/api_exports/subscriptions/123/rows?per_page=200&page=1`.
func (a *APIExportsController) ListExportedSubscriptions(
	ctx context.Context,
	input ListExportedSubscriptionsInput) (
	models.ApiResponse[[]models.Subscription],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api_exports/subscriptions/%v/rows.json", input.BatchId),
	)
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})
	if input.PerPage != nil {
		req.QueryParam("per_page", *input.PerPage)
	}
	if input.Page != nil {
		req.QueryParam("page", *input.Page)
	}

	var result []models.Subscription
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Subscription](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"409": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSingleErrorResponse},
	})
	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
		"409": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSingleErrorResponse},
	})
	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"409": {TemplatedMessage: "HTTP Response Not OK. Status code: {$statusCode}. Response: '{$response.body}'.", Unmarshaller: errors.NewSingleErrorResponse},
	})
	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})

	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})

	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
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
	req.Authenticate(NewAuth("BasicAuth"))
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"404": {TemplatedMessage: "Not Found:'{$response.body}'"},
	})

	var result models.BatchJobResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.BatchJobResponse](decoder)
	return models.NewApiResponse(result, resp), err
}
