// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimInventoryItemsListReader is a Reader for the DcimInventoryItemsList structure.
type DcimInventoryItemsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimInventoryItemsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInventoryItemsListOK creates a DcimInventoryItemsListOK with default headers values
func NewDcimInventoryItemsListOK() *DcimInventoryItemsListOK {
	return &DcimInventoryItemsListOK{}
}

/*DcimInventoryItemsListOK handles this case with default header values.

DcimInventoryItemsListOK dcim inventory items list o k
*/
type DcimInventoryItemsListOK struct {
}

func (o *DcimInventoryItemsListOK) Error() string {
	return fmt.Sprintf("[GET /api/dcim/inventory-items/][%d] dcimInventoryItemsListOK ", 200)
}

func (o *DcimInventoryItemsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
