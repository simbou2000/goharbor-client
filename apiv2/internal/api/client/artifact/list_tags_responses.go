// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mittwald/goharbor-client/v3/apiv2/model"
)

// ListTagsReader is a Reader for the ListTags structure.
type ListTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTagsOK creates a ListTagsOK with default headers values
func NewListTagsOK() *ListTagsOK {
	return &ListTagsOK{}
}

/* ListTagsOK describes a response with status code 200, with default header values.

Success
*/
type ListTagsOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of tags
	 */
	XTotalCount int64

	Payload []*model.Tag
}

func (o *ListTagsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] listTagsOK  %+v", 200, o.Payload)
}
func (o *ListTagsOK) GetPayload() []*model.Tag {
	return o.Payload
}

func (o *ListTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTagsBadRequest creates a ListTagsBadRequest with default headers values
func NewListTagsBadRequest() *ListTagsBadRequest {
	return &ListTagsBadRequest{}
}

/* ListTagsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListTagsBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] listTagsBadRequest  %+v", 400, o.Payload)
}
func (o *ListTagsBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTagsUnauthorized creates a ListTagsUnauthorized with default headers values
func NewListTagsUnauthorized() *ListTagsUnauthorized {
	return &ListTagsUnauthorized{}
}

/* ListTagsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListTagsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] listTagsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListTagsUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTagsForbidden creates a ListTagsForbidden with default headers values
func NewListTagsForbidden() *ListTagsForbidden {
	return &ListTagsForbidden{}
}

/* ListTagsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListTagsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] listTagsForbidden  %+v", 403, o.Payload)
}
func (o *ListTagsForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTagsNotFound creates a ListTagsNotFound with default headers values
func NewListTagsNotFound() *ListTagsNotFound {
	return &ListTagsNotFound{}
}

/* ListTagsNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListTagsNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] listTagsNotFound  %+v", 404, o.Payload)
}
func (o *ListTagsNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListTagsInternalServerError creates a ListTagsInternalServerError with default headers values
func NewListTagsInternalServerError() *ListTagsInternalServerError {
	return &ListTagsInternalServerError{}
}

/* ListTagsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListTagsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags][%d] listTagsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListTagsInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
