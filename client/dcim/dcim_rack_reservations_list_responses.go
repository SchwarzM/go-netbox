// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimRackReservationsListReader is a Reader for the DcimRackReservationsList structure.
type DcimRackReservationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRackReservationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRackReservationsListOK creates a DcimRackReservationsListOK with default headers values
func NewDcimRackReservationsListOK() *DcimRackReservationsListOK {
	return &DcimRackReservationsListOK{}
}

/*DcimRackReservationsListOK handles this case with default header values.

DcimRackReservationsListOK dcim rack reservations list o k
*/
type DcimRackReservationsListOK struct {
}

func (o *DcimRackReservationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/rack-reservations/][%d] dcimRackReservationsListOK ", 200)
}

func (o *DcimRackReservationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
