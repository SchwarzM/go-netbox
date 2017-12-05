// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimRacksListReader is a Reader for the DcimRacksList structure.
type DcimRacksListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRacksListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRacksListOK creates a DcimRacksListOK with default headers values
func NewDcimRacksListOK() *DcimRacksListOK {
	return &DcimRacksListOK{}
}

/*DcimRacksListOK handles this case with default header values.

DcimRacksListOK dcim racks list o k
*/
type DcimRacksListOK struct {
}

func (o *DcimRacksListOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/racks/][%d] dcimRacksListOK ", 200)
}

func (o *DcimRacksListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
