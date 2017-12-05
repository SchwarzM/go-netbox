// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimInterfaceConnectionsReadReader is a Reader for the DcimInterfaceConnectionsRead structure.
type DcimInterfaceConnectionsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceConnectionsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimInterfaceConnectionsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfaceConnectionsReadOK creates a DcimInterfaceConnectionsReadOK with default headers values
func NewDcimInterfaceConnectionsReadOK() *DcimInterfaceConnectionsReadOK {
	return &DcimInterfaceConnectionsReadOK{}
}

/*DcimInterfaceConnectionsReadOK handles this case with default header values.

DcimInterfaceConnectionsReadOK dcim interface connections read o k
*/
type DcimInterfaceConnectionsReadOK struct {
}

func (o *DcimInterfaceConnectionsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/interface-connections/{id}/][%d] dcimInterfaceConnectionsReadOK ", 200)
}

func (o *DcimInterfaceConnectionsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
