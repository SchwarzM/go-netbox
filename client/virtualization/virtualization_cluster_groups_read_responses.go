// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VirtualizationClusterGroupsReadReader is a Reader for the VirtualizationClusterGroupsRead structure.
type VirtualizationClusterGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationClusterGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsReadOK creates a VirtualizationClusterGroupsReadOK with default headers values
func NewVirtualizationClusterGroupsReadOK() *VirtualizationClusterGroupsReadOK {
	return &VirtualizationClusterGroupsReadOK{}
}

/*VirtualizationClusterGroupsReadOK handles this case with default header values.

VirtualizationClusterGroupsReadOK virtualization cluster groups read o k
*/
type VirtualizationClusterGroupsReadOK struct {
}

func (o *VirtualizationClusterGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadOK ", 200)
}

func (o *VirtualizationClusterGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
