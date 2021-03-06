// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimRackRolesListReader is a Reader for the DcimRackRolesList structure.
type DcimRackRolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRackRolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRackRolesListOK creates a DcimRackRolesListOK with default headers values
func NewDcimRackRolesListOK() *DcimRackRolesListOK {
	return &DcimRackRolesListOK{}
}

/*DcimRackRolesListOK handles this case with default header values.

DcimRackRolesListOK dcim rack roles list o k
*/
type DcimRackRolesListOK struct {
}

func (o *DcimRackRolesListOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/rack-roles/][%d] dcimRackRolesListOK ", 200)
}

func (o *DcimRackRolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
