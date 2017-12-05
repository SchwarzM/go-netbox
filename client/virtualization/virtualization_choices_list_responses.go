// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VirtualizationChoicesListReader is a Reader for the VirtualizationChoicesList structure.
type VirtualizationChoicesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationChoicesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationChoicesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationChoicesListOK creates a VirtualizationChoicesListOK with default headers values
func NewVirtualizationChoicesListOK() *VirtualizationChoicesListOK {
	return &VirtualizationChoicesListOK{}
}

/*VirtualizationChoicesListOK handles this case with default header values.

VirtualizationChoicesListOK virtualization choices list o k
*/
type VirtualizationChoicesListOK struct {
}

func (o *VirtualizationChoicesListOK) Error() string {
	return fmt.Sprintf("[GET /api/virtualization/_choices/][%d] virtualizationChoicesListOK ", 200)
}

func (o *VirtualizationChoicesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
