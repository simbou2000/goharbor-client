// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/simbou2000/goharbor-client/v5/apiv2/model"
)

// GetRegistryReader is a Reader for the GetRegistry structure.
type GetRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegistryOK creates a GetRegistryOK with default headers values
func NewGetRegistryOK() *GetRegistryOK {
	return &GetRegistryOK{}
}

/*GetRegistryOK handles this case with default header values.

Success
*/
type GetRegistryOK struct {
	Payload *model.Registry
}

func (o *GetRegistryOK) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] getRegistryOK  %+v", 200, o.Payload)
}

func (o *GetRegistryOK) GetPayload() *model.Registry {
	return o.Payload
}

func (o *GetRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Registry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryUnauthorized creates a GetRegistryUnauthorized with default headers values
func NewGetRegistryUnauthorized() *GetRegistryUnauthorized {
	return &GetRegistryUnauthorized{}
}

/*GetRegistryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetRegistryUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] getRegistryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRegistryUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryForbidden creates a GetRegistryForbidden with default headers values
func NewGetRegistryForbidden() *GetRegistryForbidden {
	return &GetRegistryForbidden{}
}

/*GetRegistryForbidden handles this case with default header values.

Forbidden
*/
type GetRegistryForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRegistryForbidden) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] getRegistryForbidden  %+v", 403, o.Payload)
}

func (o *GetRegistryForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryNotFound creates a GetRegistryNotFound with default headers values
func NewGetRegistryNotFound() *GetRegistryNotFound {
	return &GetRegistryNotFound{}
}

/*GetRegistryNotFound handles this case with default header values.

Not found
*/
type GetRegistryNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRegistryNotFound) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] getRegistryNotFound  %+v", 404, o.Payload)
}

func (o *GetRegistryNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryInternalServerError creates a GetRegistryInternalServerError with default headers values
func NewGetRegistryInternalServerError() *GetRegistryInternalServerError {
	return &GetRegistryInternalServerError{}
}

/*GetRegistryInternalServerError handles this case with default header values.

Internal server error
*/
type GetRegistryInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRegistryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registries/{id}][%d] getRegistryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegistryInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRegistryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
