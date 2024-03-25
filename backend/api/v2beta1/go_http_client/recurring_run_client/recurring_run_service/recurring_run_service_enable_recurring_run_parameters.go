// Code generated by go-swagger; DO NOT EDIT.

package recurring_run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRecurringRunServiceEnableRecurringRunParams creates a new RecurringRunServiceEnableRecurringRunParams object
// with the default values initialized.
func NewRecurringRunServiceEnableRecurringRunParams() *RecurringRunServiceEnableRecurringRunParams {
	var ()
	return &RecurringRunServiceEnableRecurringRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRecurringRunServiceEnableRecurringRunParamsWithTimeout creates a new RecurringRunServiceEnableRecurringRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRecurringRunServiceEnableRecurringRunParamsWithTimeout(timeout time.Duration) *RecurringRunServiceEnableRecurringRunParams {
	var ()
	return &RecurringRunServiceEnableRecurringRunParams{

		timeout: timeout,
	}
}

// NewRecurringRunServiceEnableRecurringRunParamsWithContext creates a new RecurringRunServiceEnableRecurringRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewRecurringRunServiceEnableRecurringRunParamsWithContext(ctx context.Context) *RecurringRunServiceEnableRecurringRunParams {
	var ()
	return &RecurringRunServiceEnableRecurringRunParams{

		Context: ctx,
	}
}

// NewRecurringRunServiceEnableRecurringRunParamsWithHTTPClient creates a new RecurringRunServiceEnableRecurringRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRecurringRunServiceEnableRecurringRunParamsWithHTTPClient(client *http.Client) *RecurringRunServiceEnableRecurringRunParams {
	var ()
	return &RecurringRunServiceEnableRecurringRunParams{
		HTTPClient: client,
	}
}

/*RecurringRunServiceEnableRecurringRunParams contains all the parameters to send to the API endpoint
for the recurring run service enable recurring run operation typically these are written to a http.Request
*/
type RecurringRunServiceEnableRecurringRunParams struct {

	/*RecurringRunID
	  The ID of the recurring runs to be enabled.

	*/
	RecurringRunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the recurring run service enable recurring run params
func (o *RecurringRunServiceEnableRecurringRunParams) WithTimeout(timeout time.Duration) *RecurringRunServiceEnableRecurringRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recurring run service enable recurring run params
func (o *RecurringRunServiceEnableRecurringRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recurring run service enable recurring run params
func (o *RecurringRunServiceEnableRecurringRunParams) WithContext(ctx context.Context) *RecurringRunServiceEnableRecurringRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recurring run service enable recurring run params
func (o *RecurringRunServiceEnableRecurringRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recurring run service enable recurring run params
func (o *RecurringRunServiceEnableRecurringRunParams) WithHTTPClient(client *http.Client) *RecurringRunServiceEnableRecurringRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recurring run service enable recurring run params
func (o *RecurringRunServiceEnableRecurringRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecurringRunID adds the recurringRunID to the recurring run service enable recurring run params
func (o *RecurringRunServiceEnableRecurringRunParams) WithRecurringRunID(recurringRunID string) *RecurringRunServiceEnableRecurringRunParams {
	o.SetRecurringRunID(recurringRunID)
	return o
}

// SetRecurringRunID adds the recurringRunId to the recurring run service enable recurring run params
func (o *RecurringRunServiceEnableRecurringRunParams) SetRecurringRunID(recurringRunID string) {
	o.RecurringRunID = recurringRunID
}

// WriteToRequest writes these params to a swagger request
func (o *RecurringRunServiceEnableRecurringRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param recurring_run_id
	if err := r.SetPathParam("recurring_run_id", o.RecurringRunID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}