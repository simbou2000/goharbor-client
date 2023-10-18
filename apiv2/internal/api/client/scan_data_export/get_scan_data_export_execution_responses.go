// Code generated by go-swagger; DO NOT EDIT.

package scan_data_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/simbou2000/goharbor-client/v5/apiv2/model"
)

// GetScanDataExportExecutionReader is a Reader for the GetScanDataExportExecution structure.
type GetScanDataExportExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanDataExportExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScanDataExportExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScanDataExportExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScanDataExportExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScanDataExportExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScanDataExportExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScanDataExportExecutionOK creates a GetScanDataExportExecutionOK with default headers values
func NewGetScanDataExportExecutionOK() *GetScanDataExportExecutionOK {
	return &GetScanDataExportExecutionOK{}
}

/*GetScanDataExportExecutionOK handles this case with default header values.

Success
*/
type GetScanDataExportExecutionOK struct {
	Payload *model.ScanDataExportExecution
}

func (o *GetScanDataExportExecutionOK) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionOK  %+v", 200, o.Payload)
}

func (o *GetScanDataExportExecutionOK) GetPayload() *model.ScanDataExportExecution {
	return o.Payload
}

func (o *GetScanDataExportExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ScanDataExportExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanDataExportExecutionUnauthorized creates a GetScanDataExportExecutionUnauthorized with default headers values
func NewGetScanDataExportExecutionUnauthorized() *GetScanDataExportExecutionUnauthorized {
	return &GetScanDataExportExecutionUnauthorized{}
}

/*GetScanDataExportExecutionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetScanDataExportExecutionUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScanDataExportExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanDataExportExecutionUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanDataExportExecutionForbidden creates a GetScanDataExportExecutionForbidden with default headers values
func NewGetScanDataExportExecutionForbidden() *GetScanDataExportExecutionForbidden {
	return &GetScanDataExportExecutionForbidden{}
}

/*GetScanDataExportExecutionForbidden handles this case with default header values.

Forbidden
*/
type GetScanDataExportExecutionForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScanDataExportExecutionForbidden) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionForbidden  %+v", 403, o.Payload)
}

func (o *GetScanDataExportExecutionForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanDataExportExecutionNotFound creates a GetScanDataExportExecutionNotFound with default headers values
func NewGetScanDataExportExecutionNotFound() *GetScanDataExportExecutionNotFound {
	return &GetScanDataExportExecutionNotFound{}
}

/*GetScanDataExportExecutionNotFound handles this case with default header values.

Not found
*/
type GetScanDataExportExecutionNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScanDataExportExecutionNotFound) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionNotFound  %+v", 404, o.Payload)
}

func (o *GetScanDataExportExecutionNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanDataExportExecutionInternalServerError creates a GetScanDataExportExecutionInternalServerError with default headers values
func NewGetScanDataExportExecutionInternalServerError() *GetScanDataExportExecutionInternalServerError {
	return &GetScanDataExportExecutionInternalServerError{}
}

/*GetScanDataExportExecutionInternalServerError handles this case with default header values.

Internal server error
*/
type GetScanDataExportExecutionInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetScanDataExportExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScanDataExportExecutionInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
