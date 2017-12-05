// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimConsoleServerPortsReadReader is a Reader for the DcimConsoleServerPortsRead structure.
type DcimConsoleServerPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimConsoleServerPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimConsoleServerPortsReadOK creates a DcimConsoleServerPortsReadOK with default headers values
func NewDcimConsoleServerPortsReadOK() *DcimConsoleServerPortsReadOK {
	return &DcimConsoleServerPortsReadOK{}
}

/*DcimConsoleServerPortsReadOK handles this case with default header values.

DcimConsoleServerPortsReadOK dcim console server ports read o k
*/
type DcimConsoleServerPortsReadOK struct {
}

func (o *DcimConsoleServerPortsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsReadOK ", 200)
}

func (o *DcimConsoleServerPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
