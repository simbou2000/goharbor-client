// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostUsersReader is a Reader for the PostUsers structure.
type PostUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostUsersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUsersCreated creates a PostUsersCreated with default headers values
func NewPostUsersCreated() *PostUsersCreated {
	return &PostUsersCreated{}
}

/* PostUsersCreated describes a response with status code 201, with default header values.

User created successfully.
*/
type PostUsersCreated struct {
}

func (o *PostUsersCreated) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersCreated ", 201)
}

func (o *PostUsersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUsersBadRequest creates a PostUsersBadRequest with default headers values
func NewPostUsersBadRequest() *PostUsersBadRequest {
	return &PostUsersBadRequest{}
}

/* PostUsersBadRequest describes a response with status code 400, with default header values.

Unsatisfied with constraints of the user creation.
*/
type PostUsersBadRequest struct {
}

func (o *PostUsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersBadRequest ", 400)
}

func (o *PostUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUsersForbidden creates a PostUsersForbidden with default headers values
func NewPostUsersForbidden() *PostUsersForbidden {
	return &PostUsersForbidden{}
}

/* PostUsersForbidden describes a response with status code 403, with default header values.

User registration can only be used by admin role user when self-registration is off.
*/
type PostUsersForbidden struct {
}

func (o *PostUsersForbidden) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersForbidden ", 403)
}

func (o *PostUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUsersUnsupportedMediaType creates a PostUsersUnsupportedMediaType with default headers values
func NewPostUsersUnsupportedMediaType() *PostUsersUnsupportedMediaType {
	return &PostUsersUnsupportedMediaType{}
}

/* PostUsersUnsupportedMediaType describes a response with status code 415, with default header values.

The Media Type of the request is not supported, it has to be "application/json"
*/
type PostUsersUnsupportedMediaType struct {
}

func (o *PostUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersUnsupportedMediaType ", 415)
}

func (o *PostUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUsersInternalServerError creates a PostUsersInternalServerError with default headers values
func NewPostUsersInternalServerError() *PostUsersInternalServerError {
	return &PostUsersInternalServerError{}
}

/* PostUsersInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type PostUsersInternalServerError struct {
}

func (o *PostUsersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users][%d] postUsersInternalServerError ", 500)
}

func (o *PostUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
