// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// IPAMRolesListReader is a Reader for the IPAMRolesList structure.
type IPAMRolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMRolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMRolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMRolesListOK creates a IPAMRolesListOK with default headers values
func NewIPAMRolesListOK() *IPAMRolesListOK {
	return &IPAMRolesListOK{}
}

/*IPAMRolesListOK handles this case with default header values.

IPAMRolesListOK ipam roles list o k
*/
type IPAMRolesListOK struct {
}

func (o *IPAMRolesListOK) Error() string {
	return fmt.Sprintf("[GET /api/ipam/roles/][%d] ipamRolesListOK ", 200)
}

func (o *IPAMRolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
