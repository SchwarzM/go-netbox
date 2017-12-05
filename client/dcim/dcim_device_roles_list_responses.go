// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimDeviceRolesListReader is a Reader for the DcimDeviceRolesList structure.
type DcimDeviceRolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimDeviceRolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDeviceRolesListOK creates a DcimDeviceRolesListOK with default headers values
func NewDcimDeviceRolesListOK() *DcimDeviceRolesListOK {
	return &DcimDeviceRolesListOK{}
}

/*DcimDeviceRolesListOK handles this case with default header values.

DcimDeviceRolesListOK dcim device roles list o k
*/
type DcimDeviceRolesListOK struct {
}

func (o *DcimDeviceRolesListOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/device-roles/][%d] dcimDeviceRolesListOK ", 200)
}

func (o *DcimDeviceRolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
