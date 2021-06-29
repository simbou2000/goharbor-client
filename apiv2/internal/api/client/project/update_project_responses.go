// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v3/apiv2/model"
)

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/* UpdateProjectOK describes a response with status code 200, with default header values.

Success
*/
type UpdateProjectOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}][%d] updateProjectOK ", 200)
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateProjectBadRequest creates a UpdateProjectBadRequest with default headers values
func NewUpdateProjectBadRequest() *UpdateProjectBadRequest {
	return &UpdateProjectBadRequest{}
}

/* UpdateProjectBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateProjectBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}][%d] updateProjectBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateProjectBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectUnauthorized creates a UpdateProjectUnauthorized with default headers values
func NewUpdateProjectUnauthorized() *UpdateProjectUnauthorized {
	return &UpdateProjectUnauthorized{}
}

/* UpdateProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateProjectUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}][%d] updateProjectUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateProjectUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectForbidden creates a UpdateProjectForbidden with default headers values
func NewUpdateProjectForbidden() *UpdateProjectForbidden {
	return &UpdateProjectForbidden{}
}

/* UpdateProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateProjectForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}][%d] updateProjectForbidden  %+v", 403, o.Payload)
}
func (o *UpdateProjectForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectNotFound creates a UpdateProjectNotFound with default headers values
func NewUpdateProjectNotFound() *UpdateProjectNotFound {
	return &UpdateProjectNotFound{}
}

/* UpdateProjectNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateProjectNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}][%d] updateProjectNotFound  %+v", 404, o.Payload)
}
func (o *UpdateProjectNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectInternalServerError creates a UpdateProjectInternalServerError with default headers values
func NewUpdateProjectInternalServerError() *UpdateProjectInternalServerError {
	return &UpdateProjectInternalServerError{}
}

/* UpdateProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateProjectInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *UpdateProjectInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}][%d] updateProjectInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateProjectInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *UpdateProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
