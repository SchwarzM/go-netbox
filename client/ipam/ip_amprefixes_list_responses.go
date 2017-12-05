// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// IPAMPrefixesListReader is a Reader for the IPAMPrefixesList structure.
type IPAMPrefixesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMPrefixesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMPrefixesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMPrefixesListOK creates a IPAMPrefixesListOK with default headers values
func NewIPAMPrefixesListOK() *IPAMPrefixesListOK {
	return &IPAMPrefixesListOK{}
}

/*IPAMPrefixesListOK handles this case with default header values.

IPAMPrefixesListOK ipam prefixes list o k
*/
type IPAMPrefixesListOK struct {
}

func (o *IPAMPrefixesListOK) Error() string {
	return fmt.Sprintf("[GET /api/ipam/prefixes/][%d] ipamPrefixesListOK ", 200)
}

func (o *IPAMPrefixesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
