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

	"github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
)

// NewPutProjectsProjectIDRobotsRobotIDParams creates a new PutProjectsProjectIDRobotsRobotIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutProjectsProjectIDRobotsRobotIDParams() *PutProjectsProjectIDRobotsRobotIDParams {
	return &PutProjectsProjectIDRobotsRobotIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectsProjectIDRobotsRobotIDParamsWithTimeout creates a new PutProjectsProjectIDRobotsRobotIDParams object
// with the ability to set a timeout on a request.
func NewPutProjectsProjectIDRobotsRobotIDParamsWithTimeout(timeout time.Duration) *PutProjectsProjectIDRobotsRobotIDParams {
	return &PutProjectsProjectIDRobotsRobotIDParams{
		timeout: timeout,
	}
}

// NewPutProjectsProjectIDRobotsRobotIDParamsWithContext creates a new PutProjectsProjectIDRobotsRobotIDParams object
// with the ability to set a context for a request.
func NewPutProjectsProjectIDRobotsRobotIDParamsWithContext(ctx context.Context) *PutProjectsProjectIDRobotsRobotIDParams {
	return &PutProjectsProjectIDRobotsRobotIDParams{
		Context: ctx,
	}
}

// NewPutProjectsProjectIDRobotsRobotIDParamsWithHTTPClient creates a new PutProjectsProjectIDRobotsRobotIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutProjectsProjectIDRobotsRobotIDParamsWithHTTPClient(client *http.Client) *PutProjectsProjectIDRobotsRobotIDParams {
	return &PutProjectsProjectIDRobotsRobotIDParams{
		HTTPClient: client,
	}
}

/* PutProjectsProjectIDRobotsRobotIDParams contains all the parameters to send to the API endpoint
   for the put projects project ID robots robot ID operation.

   Typically these are written to a http.Request.
*/
type PutProjectsProjectIDRobotsRobotIDParams struct {

	/* ProjectID.

	   Relevant project ID.

	   Format: int64
	*/
	ProjectID int64

	/* Robot.

	   Request body of enable/disable a robot account.
	*/
	Robot *legacy.RobotAccountUpdate

	/* RobotID.

	   The ID of robot account.

	   Format: int64
	*/
	RobotID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put projects project ID robots robot ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectsProjectIDRobotsRobotIDParams) WithDefaults() *PutProjectsProjectIDRobotsRobotIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put projects project ID robots robot ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectsProjectIDRobotsRobotIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) WithTimeout(timeout time.Duration) *PutProjectsProjectIDRobotsRobotIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) WithContext(ctx context.Context) *PutProjectsProjectIDRobotsRobotIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) WithHTTPClient(client *http.Client) *PutProjectsProjectIDRobotsRobotIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) WithProjectID(projectID int64) *PutProjectsProjectIDRobotsRobotIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithRobot adds the robot to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) WithRobot(robot *legacy.RobotAccountUpdate) *PutProjectsProjectIDRobotsRobotIDParams {
	o.SetRobot(robot)
	return o
}

// SetRobot adds the robot to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) SetRobot(robot *legacy.RobotAccountUpdate) {
	o.Robot = robot
}

// WithRobotID adds the robotID to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) WithRobotID(robotID int64) *PutProjectsProjectIDRobotsRobotIDParams {
	o.SetRobotID(robotID)
	return o
}

// SetRobotID adds the robotId to the put projects project ID robots robot ID params
func (o *PutProjectsProjectIDRobotsRobotIDParams) SetRobotID(robotID int64) {
	o.RobotID = robotID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectsProjectIDRobotsRobotIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}
	if o.Robot != nil {
		if err := r.SetBodyParam(o.Robot); err != nil {
			return err
		}
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
