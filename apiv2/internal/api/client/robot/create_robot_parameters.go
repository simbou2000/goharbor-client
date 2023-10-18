// Code generated by go-swagger; DO NOT EDIT.

package robot

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

	"github.com/simbou2000/goharbor-client/v5/apiv2/model"
)

// NewCreateRobotParams creates a new CreateRobotParams object
// with the default values initialized.
func NewCreateRobotParams() *CreateRobotParams {
	var ()
	return &CreateRobotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRobotParamsWithTimeout creates a new CreateRobotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRobotParamsWithTimeout(timeout time.Duration) *CreateRobotParams {
	var ()
	return &CreateRobotParams{

		timeout: timeout,
	}
}

// NewCreateRobotParamsWithContext creates a new CreateRobotParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRobotParamsWithContext(ctx context.Context) *CreateRobotParams {
	var ()
	return &CreateRobotParams{

		Context: ctx,
	}
}

// NewCreateRobotParamsWithHTTPClient creates a new CreateRobotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRobotParamsWithHTTPClient(client *http.Client) *CreateRobotParams {
	var ()
	return &CreateRobotParams{
		HTTPClient: client,
	}
}

/*CreateRobotParams contains all the parameters to send to the API endpoint
for the create robot operation typically these are written to a http.Request
*/
type CreateRobotParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Robot
	  The JSON object of a robot account.

	*/
	Robot *model.RobotCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create robot params
func (o *CreateRobotParams) WithTimeout(timeout time.Duration) *CreateRobotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create robot params
func (o *CreateRobotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create robot params
func (o *CreateRobotParams) WithContext(ctx context.Context) *CreateRobotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create robot params
func (o *CreateRobotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create robot params
func (o *CreateRobotParams) WithHTTPClient(client *http.Client) *CreateRobotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create robot params
func (o *CreateRobotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create robot params
func (o *CreateRobotParams) WithXRequestID(xRequestID *string) *CreateRobotParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create robot params
func (o *CreateRobotParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRobot adds the robot to the create robot params
func (o *CreateRobotParams) WithRobot(robot *model.RobotCreate) *CreateRobotParams {
	o.SetRobot(robot)
	return o
}

// SetRobot adds the robot to the create robot params
func (o *CreateRobotParams) SetRobot(robot *model.RobotCreate) {
	o.Robot = robot
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRobotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Robot != nil {
		if err := r.SetBodyParam(o.Robot); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
