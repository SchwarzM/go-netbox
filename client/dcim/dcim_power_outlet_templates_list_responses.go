// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimPowerOutletTemplatesListReader is a Reader for the DcimPowerOutletTemplatesList structure.
type DcimPowerOutletTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPowerOutletTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesListOK creates a DcimPowerOutletTemplatesListOK with default headers values
func NewDcimPowerOutletTemplatesListOK() *DcimPowerOutletTemplatesListOK {
	return &DcimPowerOutletTemplatesListOK{}
}

/*DcimPowerOutletTemplatesListOK handles this case with default header values.

DcimPowerOutletTemplatesListOK dcim power outlet templates list o k
*/
type DcimPowerOutletTemplatesListOK struct {
}

func (o *DcimPowerOutletTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesListOK ", 200)
}

func (o *DcimPowerOutletTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
