// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
)

// GetLabelsReader is a Reader for the GetLabels structure.
type GetLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLabelsOK creates a GetLabelsOK with default headers values
func NewGetLabelsOK() *GetLabelsOK {
	return &GetLabelsOK{}
}

/* GetLabelsOK describes a response with status code 200, with default header values.

Get successfully.
*/
type GetLabelsOK struct {

	/* Link to previous page and next page
	 */
	Link string

	/* The total count of available items
	 */
	XTotalCount int64

	Payload []*legacy.Label
}

func (o *GetLabelsOK) Error() string {
	return fmt.Sprintf("[GET /labels][%d] getLabelsOK  %+v", 200, o.Payload)
}
func (o *GetLabelsOK) GetPayload() []*legacy.Label {
	return o.Payload
}

func (o *GetLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLabelsBadRequest creates a GetLabelsBadRequest with default headers values
func NewGetLabelsBadRequest() *GetLabelsBadRequest {
	return &GetLabelsBadRequest{}
}

/* GetLabelsBadRequest describes a response with status code 400, with default header values.

Invalid parameters.
*/
type GetLabelsBadRequest struct {
}

func (o *GetLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /labels][%d] getLabelsBadRequest ", 400)
}

func (o *GetLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelsUnauthorized creates a GetLabelsUnauthorized with default headers values
func NewGetLabelsUnauthorized() *GetLabelsUnauthorized {
	return &GetLabelsUnauthorized{}
}

/* GetLabelsUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type GetLabelsUnauthorized struct {
}

func (o *GetLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /labels][%d] getLabelsUnauthorized ", 401)
}

func (o *GetLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelsInternalServerError creates a GetLabelsInternalServerError with default headers values
func NewGetLabelsInternalServerError() *GetLabelsInternalServerError {
	return &GetLabelsInternalServerError{}
}

/* GetLabelsInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetLabelsInternalServerError struct {
}

func (o *GetLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /labels][%d] getLabelsInternalServerError ", 500)
}

func (o *GetLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
