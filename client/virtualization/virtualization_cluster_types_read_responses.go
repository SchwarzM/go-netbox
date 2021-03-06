// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VirtualizationClusterTypesReadReader is a Reader for the VirtualizationClusterTypesRead structure.
type VirtualizationClusterTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationClusterTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterTypesReadOK creates a VirtualizationClusterTypesReadOK with default headers values
func NewVirtualizationClusterTypesReadOK() *VirtualizationClusterTypesReadOK {
	return &VirtualizationClusterTypesReadOK{}
}

/*VirtualizationClusterTypesReadOK handles this case with default header values.

VirtualizationClusterTypesReadOK virtualization cluster types read o k
*/
type VirtualizationClusterTypesReadOK struct {
}

func (o *VirtualizationClusterTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesReadOK ", 200)
}

func (o *VirtualizationClusterTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
