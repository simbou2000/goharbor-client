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

// GetArtifactReader is a Reader for the GetArtifact structure.
type GetArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArtifactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArtifactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArtifactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArtifactOK creates a GetArtifactOK with default headers values
func NewGetArtifactOK() *GetArtifactOK {
	return &GetArtifactOK{}
}

/*GetArtifactOK handles this case with default header values.

Success
*/
type GetArtifactOK struct {
	Payload *model.Artifact
}

func (o *GetArtifactOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactOK  %+v", 200, o.Payload)
}

func (o *GetArtifactOK) GetPayload() *model.Artifact {
	return o.Payload
}

func (o *GetArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Artifact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactBadRequest creates a GetArtifactBadRequest with default headers values
func NewGetArtifactBadRequest() *GetArtifactBadRequest {
	return &GetArtifactBadRequest{}
}

/*GetArtifactBadRequest handles this case with default header values.

Bad request
*/
type GetArtifactBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetArtifactBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactBadRequest  %+v", 400, o.Payload)
}

func (o *GetArtifactBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetArtifactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactUnauthorized creates a GetArtifactUnauthorized with default headers values
func NewGetArtifactUnauthorized() *GetArtifactUnauthorized {
	return &GetArtifactUnauthorized{}
}

/*GetArtifactUnauthorized handles this case with default header values.

Unauthorized
*/
type GetArtifactUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetArtifactUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArtifactUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetArtifactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactForbidden creates a GetArtifactForbidden with default headers values
func NewGetArtifactForbidden() *GetArtifactForbidden {
	return &GetArtifactForbidden{}
}

/*GetArtifactForbidden handles this case with default header values.

Forbidden
*/
type GetArtifactForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetArtifactForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactForbidden  %+v", 403, o.Payload)
}

func (o *GetArtifactForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactNotFound creates a GetArtifactNotFound with default headers values
func NewGetArtifactNotFound() *GetArtifactNotFound {
	return &GetArtifactNotFound{}
}

/*GetArtifactNotFound handles this case with default header values.

Not found
*/
type GetArtifactNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetArtifactNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactNotFound  %+v", 404, o.Payload)
}

func (o *GetArtifactNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactInternalServerError creates a GetArtifactInternalServerError with default headers values
func NewGetArtifactInternalServerError() *GetArtifactInternalServerError {
	return &GetArtifactInternalServerError{}
}

/*GetArtifactInternalServerError handles this case with default header values.

Internal server error
*/
type GetArtifactInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetArtifactInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}][%d] getArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArtifactInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetArtifactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
