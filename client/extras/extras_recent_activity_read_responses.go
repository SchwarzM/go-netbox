// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ExtrasRecentActivityReadReader is a Reader for the ExtrasRecentActivityRead structure.
type ExtrasRecentActivityReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRecentActivityReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExtrasRecentActivityReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasRecentActivityReadOK creates a ExtrasRecentActivityReadOK with default headers values
func NewExtrasRecentActivityReadOK() *ExtrasRecentActivityReadOK {
	return &ExtrasRecentActivityReadOK{}
}

/*ExtrasRecentActivityReadOK handles this case with default header values.

ExtrasRecentActivityReadOK extras recent activity read o k
*/
type ExtrasRecentActivityReadOK struct {
}

func (o *ExtrasRecentActivityReadOK) Error() string {
	return fmt.Sprintf("[GET /api/extras/recent-activity/{id}/][%d] extrasRecentActivityReadOK ", 200)
}

func (o *ExtrasRecentActivityReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
