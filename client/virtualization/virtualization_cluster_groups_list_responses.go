// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VirtualizationClusterGroupsListReader is a Reader for the VirtualizationClusterGroupsList structure.
type VirtualizationClusterGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationClusterGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsListOK creates a VirtualizationClusterGroupsListOK with default headers values
func NewVirtualizationClusterGroupsListOK() *VirtualizationClusterGroupsListOK {
	return &VirtualizationClusterGroupsListOK{}
}

/*VirtualizationClusterGroupsListOK handles this case with default header values.

VirtualizationClusterGroupsListOK virtualization cluster groups list o k
*/
type VirtualizationClusterGroupsListOK struct {
}

func (o *VirtualizationClusterGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /api/virtualization/cluster-groups/][%d] virtualizationClusterGroupsListOK ", 200)
}

func (o *VirtualizationClusterGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
