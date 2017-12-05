// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimPowerPortTemplatesReadReader is a Reader for the DcimPowerPortTemplatesRead structure.
type DcimPowerPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPowerPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesReadOK creates a DcimPowerPortTemplatesReadOK with default headers values
func NewDcimPowerPortTemplatesReadOK() *DcimPowerPortTemplatesReadOK {
	return &DcimPowerPortTemplatesReadOK{}
}

/*DcimPowerPortTemplatesReadOK handles this case with default header values.

DcimPowerPortTemplatesReadOK dcim power port templates read o k
*/
type DcimPowerPortTemplatesReadOK struct {
}

func (o *DcimPowerPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesReadOK ", 200)
}

func (o *DcimPowerPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
