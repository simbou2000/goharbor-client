// Code generated by go-swagger; DO NOT EDIT.

package products

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
)

// NewGetProjectsProjectIDRobotsRobotIDParams creates a new GetProjectsProjectIDRobotsRobotIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectsProjectIDRobotsRobotIDParams() *GetProjectsProjectIDRobotsRobotIDParams {
	return &GetProjectsProjectIDRobotsRobotIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsProjectIDRobotsRobotIDParamsWithTimeout creates a new GetProjectsProjectIDRobotsRobotIDParams object
// with the ability to set a timeout on a request.
func NewGetProjectsProjectIDRobotsRobotIDParamsWithTimeout(timeout time.Duration) *GetProjectsProjectIDRobotsRobotIDParams {
	return &GetProjectsProjectIDRobotsRobotIDParams{
		timeout: timeout,
	}
}

// NewGetProjectsProjectIDRobotsRobotIDParamsWithContext creates a new GetProjectsProjectIDRobotsRobotIDParams object
// with the ability to set a context for a request.
func NewGetProjectsProjectIDRobotsRobotIDParamsWithContext(ctx context.Context) *GetProjectsProjectIDRobotsRobotIDParams {
	return &GetProjectsProjectIDRobotsRobotIDParams{
		Context: ctx,
	}
}

// NewGetProjectsProjectIDRobotsRobotIDParamsWithHTTPClient creates a new GetProjectsProjectIDRobotsRobotIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectsProjectIDRobotsRobotIDParamsWithHTTPClient(client *http.Client) *GetProjectsProjectIDRobotsRobotIDParams {
	return &GetProjectsProjectIDRobotsRobotIDParams{
		HTTPClient: client,
	}
}

/* GetProjectsProjectIDRobotsRobotIDParams contains all the parameters to send to the API endpoint
   for the get projects project ID robots robot ID operation.

   Typically these are written to a http.Request.
*/
type GetProjectsProjectIDRobotsRobotIDParams struct {

	/* ProjectID.

	   Relevant project ID.

	   Format: int64
	*/
	ProjectID int64

	/* RobotID.

	   The ID of robot account.

	   Format: int64
	*/
	RobotID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get projects project ID robots robot ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsProjectIDRobotsRobotIDParams) WithDefaults() *GetProjectsProjectIDRobotsRobotIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get projects project ID robots robot ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsProjectIDRobotsRobotIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) WithTimeout(timeout time.Duration) *GetProjectsProjectIDRobotsRobotIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) WithContext(ctx context.Context) *GetProjectsProjectIDRobotsRobotIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) WithHTTPClient(client *http.Client) *GetProjectsProjectIDRobotsRobotIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) WithProjectID(projectID int64) *GetProjectsProjectIDRobotsRobotIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithRobotID adds the robotID to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) WithRobotID(robotID int64) *GetProjectsProjectIDRobotsRobotIDParams {
	o.SetRobotID(robotID)
	return o
}

// SetRobotID adds the robotId to the get projects project ID robots robot ID params
func (o *GetProjectsProjectIDRobotsRobotIDParams) SetRobotID(robotID int64) {
	o.RobotID = robotID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsProjectIDRobotsRobotIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	// path param robot_id
	if err := r.SetPathParam("robot_id", swag.FormatInt64(o.RobotID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
