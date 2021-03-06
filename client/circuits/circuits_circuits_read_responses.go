// Code generated by go-swagger; DO NOT EDIT.

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CircuitsCircuitsReadReader is a Reader for the CircuitsCircuitsRead structure.
type CircuitsCircuitsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCircuitsCircuitsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCircuitsCircuitsReadOK creates a CircuitsCircuitsReadOK with default headers values
func NewCircuitsCircuitsReadOK() *CircuitsCircuitsReadOK {
	return &CircuitsCircuitsReadOK{}
}

/*CircuitsCircuitsReadOK handles this case with default header values.

CircuitsCircuitsReadOK circuits circuits read o k
*/
type CircuitsCircuitsReadOK struct {
}

func (o *CircuitsCircuitsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/circuits/circuits/{id}/][%d] circuitsCircuitsReadOK ", 200)
}

func (o *CircuitsCircuitsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
