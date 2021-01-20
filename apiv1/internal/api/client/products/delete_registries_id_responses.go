// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRegistriesIDReader is a Reader for the DeleteRegistriesID structure.
type DeleteRegistriesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegistriesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRegistriesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRegistriesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRegistriesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRegistriesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRegistriesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRegistriesIDOK creates a DeleteRegistriesIDOK with default headers values
func NewDeleteRegistriesIDOK() *DeleteRegistriesIDOK {
	return &DeleteRegistriesIDOK{}
}

/* DeleteRegistriesIDOK describes a response with status code 200, with default header values.

Registry deleted successfully.
*/
type DeleteRegistriesIDOK struct {
}

func (o *DeleteRegistriesIDOK) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistriesIdOK ", 200)
}

func (o *DeleteRegistriesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRegistriesIDBadRequest creates a DeleteRegistriesIDBadRequest with default headers values
func NewDeleteRegistriesIDBadRequest() *DeleteRegistriesIDBadRequest {
	return &DeleteRegistriesIDBadRequest{}
}

/* DeleteRegistriesIDBadRequest describes a response with status code 400, with default header values.

Registry's ID is invalid or the registry is used by policies.
*/
type DeleteRegistriesIDBadRequest struct {
}

func (o *DeleteRegistriesIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistriesIdBadRequest ", 400)
}

func (o *DeleteRegistriesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRegistriesIDUnauthorized creates a DeleteRegistriesIDUnauthorized with default headers values
func NewDeleteRegistriesIDUnauthorized() *DeleteRegistriesIDUnauthorized {
	return &DeleteRegistriesIDUnauthorized{}
}

/* DeleteRegistriesIDUnauthorized describes a response with status code 401, with default header values.

Only admin has this authority.
*/
type DeleteRegistriesIDUnauthorized struct {
}

func (o *DeleteRegistriesIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistriesIdUnauthorized ", 401)
}

func (o *DeleteRegistriesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRegistriesIDNotFound creates a DeleteRegistriesIDNotFound with default headers values
func NewDeleteRegistriesIDNotFound() *DeleteRegistriesIDNotFound {
	return &DeleteRegistriesIDNotFound{}
}

/* DeleteRegistriesIDNotFound describes a response with status code 404, with default header values.

Registry does not exist.
*/
type DeleteRegistriesIDNotFound struct {
}

func (o *DeleteRegistriesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistriesIdNotFound ", 404)
}

func (o *DeleteRegistriesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRegistriesIDInternalServerError creates a DeleteRegistriesIDInternalServerError with default headers values
func NewDeleteRegistriesIDInternalServerError() *DeleteRegistriesIDInternalServerError {
	return &DeleteRegistriesIDInternalServerError{}
}

/* DeleteRegistriesIDInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type DeleteRegistriesIDInternalServerError struct {
}

func (o *DeleteRegistriesIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistriesIdInternalServerError ", 500)
}

func (o *DeleteRegistriesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
