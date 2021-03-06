// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimInterfaceConnectionsListReader is a Reader for the DcimInterfaceConnectionsList structure.
type DcimInterfaceConnectionsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceConnectionsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimInterfaceConnectionsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfaceConnectionsListOK creates a DcimInterfaceConnectionsListOK with default headers values
func NewDcimInterfaceConnectionsListOK() *DcimInterfaceConnectionsListOK {
	return &DcimInterfaceConnectionsListOK{}
}

/*DcimInterfaceConnectionsListOK handles this case with default header values.

DcimInterfaceConnectionsListOK dcim interface connections list o k
*/
type DcimInterfaceConnectionsListOK struct {
}

func (o *DcimInterfaceConnectionsListOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/interface-connections/][%d] dcimInterfaceConnectionsListOK ", 200)
}

func (o *DcimInterfaceConnectionsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
