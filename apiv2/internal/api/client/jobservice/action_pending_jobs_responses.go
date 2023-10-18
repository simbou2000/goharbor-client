// Code generated by go-swagger; DO NOT EDIT.

package jobservice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/simbou2000/goharbor-client/v5/apiv2/model"
)

// ActionPendingJobsReader is a Reader for the ActionPendingJobs structure.
type ActionPendingJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionPendingJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionPendingJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewActionPendingJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewActionPendingJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActionPendingJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewActionPendingJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewActionPendingJobsOK creates a ActionPendingJobsOK with default headers values
func NewActionPendingJobsOK() *ActionPendingJobsOK {
	return &ActionPendingJobsOK{}
}

/*ActionPendingJobsOK handles this case with default header values.

take action to the jobs in the queue successfully.
*/
type ActionPendingJobsOK struct {
}

func (o *ActionPendingJobsOK) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsOK ", 200)
}

func (o *ActionPendingJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionPendingJobsUnauthorized creates a ActionPendingJobsUnauthorized with default headers values
func NewActionPendingJobsUnauthorized() *ActionPendingJobsUnauthorized {
	return &ActionPendingJobsUnauthorized{}
}

/*ActionPendingJobsUnauthorized handles this case with default header values.

Unauthorized
*/
type ActionPendingJobsUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ActionPendingJobsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *ActionPendingJobsUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ActionPendingJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionPendingJobsForbidden creates a ActionPendingJobsForbidden with default headers values
func NewActionPendingJobsForbidden() *ActionPendingJobsForbidden {
	return &ActionPendingJobsForbidden{}
}

/*ActionPendingJobsForbidden handles this case with default header values.

Forbidden
*/
type ActionPendingJobsForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ActionPendingJobsForbidden) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsForbidden  %+v", 403, o.Payload)
}

func (o *ActionPendingJobsForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ActionPendingJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionPendingJobsNotFound creates a ActionPendingJobsNotFound with default headers values
func NewActionPendingJobsNotFound() *ActionPendingJobsNotFound {
	return &ActionPendingJobsNotFound{}
}

/*ActionPendingJobsNotFound handles this case with default header values.

Not found
*/
type ActionPendingJobsNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ActionPendingJobsNotFound) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsNotFound  %+v", 404, o.Payload)
}

func (o *ActionPendingJobsNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ActionPendingJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionPendingJobsInternalServerError creates a ActionPendingJobsInternalServerError with default headers values
func NewActionPendingJobsInternalServerError() *ActionPendingJobsInternalServerError {
	return &ActionPendingJobsInternalServerError{}
}

/*ActionPendingJobsInternalServerError handles this case with default header values.

Internal server error
*/
type ActionPendingJobsInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ActionPendingJobsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /jobservice/queues/{job_type}][%d] actionPendingJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionPendingJobsInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ActionPendingJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
