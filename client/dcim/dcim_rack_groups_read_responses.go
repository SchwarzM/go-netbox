// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimRackGroupsReadReader is a Reader for the DcimRackGroupsRead structure.
type DcimRackGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRackGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRackGroupsReadOK creates a DcimRackGroupsReadOK with default headers values
func NewDcimRackGroupsReadOK() *DcimRackGroupsReadOK {
	return &DcimRackGroupsReadOK{}
}

/*DcimRackGroupsReadOK handles this case with default header values.

DcimRackGroupsReadOK dcim rack groups read o k
*/
type DcimRackGroupsReadOK struct {
}

func (o *DcimRackGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/rack-groups/{id}/][%d] dcimRackGroupsReadOK ", 200)
}

func (o *DcimRackGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
