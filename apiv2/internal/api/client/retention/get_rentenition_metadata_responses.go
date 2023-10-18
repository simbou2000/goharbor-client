// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/simbou2000/goharbor-client/v5/apiv2/model"
)

// GetRentenitionMetadataReader is a Reader for the GetRentenitionMetadata structure.
type GetRentenitionMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRentenitionMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRentenitionMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRentenitionMetadataOK creates a GetRentenitionMetadataOK with default headers values
func NewGetRentenitionMetadataOK() *GetRentenitionMetadataOK {
	return &GetRentenitionMetadataOK{}
}

/*GetRentenitionMetadataOK handles this case with default header values.

Get Retention Metadatas successfully.
*/
type GetRentenitionMetadataOK struct {
	Payload *model.RetentionMetadata
}

func (o *GetRentenitionMetadataOK) Error() string {
	return fmt.Sprintf("[GET /retentions/metadatas][%d] getRentenitionMetadataOK  %+v", 200, o.Payload)
}

func (o *GetRentenitionMetadataOK) GetPayload() *model.RetentionMetadata {
	return o.Payload
}

func (o *GetRentenitionMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RetentionMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
