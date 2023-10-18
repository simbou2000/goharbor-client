// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/simbou2000/goharbor-client/v5/apiv2/model"
)

// CopyArtifactReader is a Reader for the CopyArtifact structure.
type CopyArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCopyArtifactCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCopyArtifactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCopyArtifactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCopyArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCopyArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCopyArtifactMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCopyArtifactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCopyArtifactCreated creates a CopyArtifactCreated with default headers values
func NewCopyArtifactCreated() *CopyArtifactCreated {
	return &CopyArtifactCreated{}
}

/*CopyArtifactCreated handles this case with default header values.

Created
*/
type CopyArtifactCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CopyArtifactCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactCreated ", 201)
}

func (o *CopyArtifactCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewCopyArtifactBadRequest creates a CopyArtifactBadRequest with default headers values
func NewCopyArtifactBadRequest() *CopyArtifactBadRequest {
	return &CopyArtifactBadRequest{}
}

/*CopyArtifactBadRequest handles this case with default header values.

Bad request
*/
type CopyArtifactBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CopyArtifactBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactBadRequest  %+v", 400, o.Payload)
}

func (o *CopyArtifactBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CopyArtifactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyArtifactUnauthorized creates a CopyArtifactUnauthorized with default headers values
func NewCopyArtifactUnauthorized() *CopyArtifactUnauthorized {
	return &CopyArtifactUnauthorized{}
}

/*CopyArtifactUnauthorized handles this case with default header values.

Unauthorized
*/
type CopyArtifactUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CopyArtifactUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *CopyArtifactUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CopyArtifactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyArtifactForbidden creates a CopyArtifactForbidden with default headers values
func NewCopyArtifactForbidden() *CopyArtifactForbidden {
	return &CopyArtifactForbidden{}
}

/*CopyArtifactForbidden handles this case with default header values.

Forbidden
*/
type CopyArtifactForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CopyArtifactForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactForbidden  %+v", 403, o.Payload)
}

func (o *CopyArtifactForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CopyArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyArtifactNotFound creates a CopyArtifactNotFound with default headers values
func NewCopyArtifactNotFound() *CopyArtifactNotFound {
	return &CopyArtifactNotFound{}
}

/*CopyArtifactNotFound handles this case with default header values.

Not found
*/
type CopyArtifactNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CopyArtifactNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactNotFound  %+v", 404, o.Payload)
}

func (o *CopyArtifactNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CopyArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyArtifactMethodNotAllowed creates a CopyArtifactMethodNotAllowed with default headers values
func NewCopyArtifactMethodNotAllowed() *CopyArtifactMethodNotAllowed {
	return &CopyArtifactMethodNotAllowed{}
}

/*CopyArtifactMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type CopyArtifactMethodNotAllowed struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CopyArtifactMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CopyArtifactMethodNotAllowed) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CopyArtifactMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyArtifactInternalServerError creates a CopyArtifactInternalServerError with default headers values
func NewCopyArtifactInternalServerError() *CopyArtifactInternalServerError {
	return &CopyArtifactInternalServerError{}
}

/*CopyArtifactInternalServerError handles this case with default header values.

Internal server error
*/
type CopyArtifactInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *CopyArtifactInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *CopyArtifactInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *CopyArtifactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
