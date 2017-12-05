// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimInterfaceTemplatesListReader is a Reader for the DcimInterfaceTemplatesList structure.
type DcimInterfaceTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimInterfaceTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesListOK creates a DcimInterfaceTemplatesListOK with default headers values
func NewDcimInterfaceTemplatesListOK() *DcimInterfaceTemplatesListOK {
	return &DcimInterfaceTemplatesListOK{}
}

/*DcimInterfaceTemplatesListOK handles this case with default header values.

DcimInterfaceTemplatesListOK dcim interface templates list o k
*/
type DcimInterfaceTemplatesListOK struct {
}

func (o *DcimInterfaceTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/interface-templates/][%d] dcimInterfaceTemplatesListOK ", 200)
}

func (o *DcimInterfaceTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
