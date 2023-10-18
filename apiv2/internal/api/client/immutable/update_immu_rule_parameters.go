// Code generated by go-swagger; DO NOT EDIT.

package immutable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/simbou2000/goharbor-client/v5/apiv2/model"
)

// NewUpdateImmuRuleParams creates a new UpdateImmuRuleParams object
// with the default values initialized.
func NewUpdateImmuRuleParams() *UpdateImmuRuleParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &UpdateImmuRuleParams{
		XIsResourceName: &xIsResourceNameDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateImmuRuleParamsWithTimeout creates a new UpdateImmuRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateImmuRuleParamsWithTimeout(timeout time.Duration) *UpdateImmuRuleParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &UpdateImmuRuleParams{
		XIsResourceName: &xIsResourceNameDefault,

		timeout: timeout,
	}
}

// NewUpdateImmuRuleParamsWithContext creates a new UpdateImmuRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateImmuRuleParamsWithContext(ctx context.Context) *UpdateImmuRuleParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &UpdateImmuRuleParams{
		XIsResourceName: &xIsResourceNameDefault,

		Context: ctx,
	}
}

// NewUpdateImmuRuleParamsWithHTTPClient creates a new UpdateImmuRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateImmuRuleParamsWithHTTPClient(client *http.Client) *UpdateImmuRuleParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &UpdateImmuRuleParams{
		XIsResourceName: &xIsResourceNameDefault,
		HTTPClient:      client,
	}
}

/*UpdateImmuRuleParams contains all the parameters to send to the API endpoint
for the update immu rule operation typically these are written to a http.Request
*/
type UpdateImmuRuleParams struct {

	/*ImmutableRule*/
	ImmutableRule *model.ImmutableRule
	/*XIsResourceName
	  The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.

	*/
	XIsResourceName *bool
	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ImmutableRuleID
	  The ID of the immutable rule

	*/
	ImmutableRuleID int64
	/*ProjectNameOrID
	  The name or id of the project

	*/
	ProjectNameOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update immu rule params
func (o *UpdateImmuRuleParams) WithTimeout(timeout time.Duration) *UpdateImmuRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update immu rule params
func (o *UpdateImmuRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update immu rule params
func (o *UpdateImmuRuleParams) WithContext(ctx context.Context) *UpdateImmuRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update immu rule params
func (o *UpdateImmuRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update immu rule params
func (o *UpdateImmuRuleParams) WithHTTPClient(client *http.Client) *UpdateImmuRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update immu rule params
func (o *UpdateImmuRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImmutableRule adds the immutableRule to the update immu rule params
func (o *UpdateImmuRuleParams) WithImmutableRule(immutableRule *model.ImmutableRule) *UpdateImmuRuleParams {
	o.SetImmutableRule(immutableRule)
	return o
}

// SetImmutableRule adds the immutableRule to the update immu rule params
func (o *UpdateImmuRuleParams) SetImmutableRule(immutableRule *model.ImmutableRule) {
	o.ImmutableRule = immutableRule
}

// WithXIsResourceName adds the xIsResourceName to the update immu rule params
func (o *UpdateImmuRuleParams) WithXIsResourceName(xIsResourceName *bool) *UpdateImmuRuleParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the update immu rule params
func (o *UpdateImmuRuleParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the update immu rule params
func (o *UpdateImmuRuleParams) WithXRequestID(xRequestID *string) *UpdateImmuRuleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update immu rule params
func (o *UpdateImmuRuleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithImmutableRuleID adds the immutableRuleID to the update immu rule params
func (o *UpdateImmuRuleParams) WithImmutableRuleID(immutableRuleID int64) *UpdateImmuRuleParams {
	o.SetImmutableRuleID(immutableRuleID)
	return o
}

// SetImmutableRuleID adds the immutableRuleId to the update immu rule params
func (o *UpdateImmuRuleParams) SetImmutableRuleID(immutableRuleID int64) {
	o.ImmutableRuleID = immutableRuleID
}

// WithProjectNameOrID adds the projectNameOrID to the update immu rule params
func (o *UpdateImmuRuleParams) WithProjectNameOrID(projectNameOrID string) *UpdateImmuRuleParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the update immu rule params
func (o *UpdateImmuRuleParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateImmuRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImmutableRule != nil {
		if err := r.SetBodyParam(o.ImmutableRule); err != nil {
			return err
		}
	}

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	// path param immutable_rule_id
	if err := r.SetPathParam("immutable_rule_id", swag.FormatInt64(o.ImmutableRuleID)); err != nil {
		return err
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
