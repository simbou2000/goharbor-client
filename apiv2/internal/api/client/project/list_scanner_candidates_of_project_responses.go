// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/simbou2000/goharbor-client/v5/apiv2/model"
)

// ListScannerCandidatesOfProjectReader is a Reader for the ListScannerCandidatesOfProject structure.
type ListScannerCandidatesOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListScannerCandidatesOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListScannerCandidatesOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListScannerCandidatesOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListScannerCandidatesOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListScannerCandidatesOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListScannerCandidatesOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListScannerCandidatesOfProjectOK creates a ListScannerCandidatesOfProjectOK with default headers values
func NewListScannerCandidatesOfProjectOK() *ListScannerCandidatesOfProjectOK {
	return &ListScannerCandidatesOfProjectOK{}
}

/*ListScannerCandidatesOfProjectOK handles this case with default header values.

A list of scanner registrations.
*/
type ListScannerCandidatesOfProjectOK struct {
	/*Link to previous page and next page
	 */
	Link string
	/*The total count of available items
	 */
	XTotalCount int64

	Payload []*model.ScannerRegistration
}

func (o *ListScannerCandidatesOfProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner/candidates][%d] listScannerCandidatesOfProjectOK  %+v", 200, o.Payload)
}

func (o *ListScannerCandidatesOfProjectOK) GetPayload() []*model.ScannerRegistration {
	return o.Payload
}

func (o *ListScannerCandidatesOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScannerCandidatesOfProjectBadRequest creates a ListScannerCandidatesOfProjectBadRequest with default headers values
func NewListScannerCandidatesOfProjectBadRequest() *ListScannerCandidatesOfProjectBadRequest {
	return &ListScannerCandidatesOfProjectBadRequest{}
}

/*ListScannerCandidatesOfProjectBadRequest handles this case with default header values.

Bad request
*/
type ListScannerCandidatesOfProjectBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListScannerCandidatesOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner/candidates][%d] listScannerCandidatesOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *ListScannerCandidatesOfProjectBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListScannerCandidatesOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScannerCandidatesOfProjectUnauthorized creates a ListScannerCandidatesOfProjectUnauthorized with default headers values
func NewListScannerCandidatesOfProjectUnauthorized() *ListScannerCandidatesOfProjectUnauthorized {
	return &ListScannerCandidatesOfProjectUnauthorized{}
}

/*ListScannerCandidatesOfProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type ListScannerCandidatesOfProjectUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListScannerCandidatesOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner/candidates][%d] listScannerCandidatesOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *ListScannerCandidatesOfProjectUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListScannerCandidatesOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScannerCandidatesOfProjectForbidden creates a ListScannerCandidatesOfProjectForbidden with default headers values
func NewListScannerCandidatesOfProjectForbidden() *ListScannerCandidatesOfProjectForbidden {
	return &ListScannerCandidatesOfProjectForbidden{}
}

/*ListScannerCandidatesOfProjectForbidden handles this case with default header values.

Forbidden
*/
type ListScannerCandidatesOfProjectForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListScannerCandidatesOfProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner/candidates][%d] listScannerCandidatesOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *ListScannerCandidatesOfProjectForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListScannerCandidatesOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScannerCandidatesOfProjectInternalServerError creates a ListScannerCandidatesOfProjectInternalServerError with default headers values
func NewListScannerCandidatesOfProjectInternalServerError() *ListScannerCandidatesOfProjectInternalServerError {
	return &ListScannerCandidatesOfProjectInternalServerError{}
}

/*ListScannerCandidatesOfProjectInternalServerError handles this case with default header values.

Internal server error
*/
type ListScannerCandidatesOfProjectInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListScannerCandidatesOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner/candidates][%d] listScannerCandidatesOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *ListScannerCandidatesOfProjectInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListScannerCandidatesOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
