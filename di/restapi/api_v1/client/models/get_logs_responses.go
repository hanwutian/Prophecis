// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"webank/DI/restapi/api_v1/restmodels"
)

// GetLogsReader is a Reader for the GetLogs structure.
type GetLogsReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLogsOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogsOK creates a GetLogsOK with default headers values
func NewGetLogsOK(writer io.Writer) *GetLogsOK {
	return &GetLogsOK{
		Payload: writer,
	}
}

/*GetLogsOK handles this case with default header values.

Dump of the training log to-date
*/
type GetLogsOK struct {
	Payload io.Writer
}

func (o *GetLogsOK) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/logs][%d] getLogsOK  %+v", 200, o.Payload)
}

func (o *GetLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsUnauthorized creates a GetLogsUnauthorized with default headers values
func NewGetLogsUnauthorized() *GetLogsUnauthorized {
	return &GetLogsUnauthorized{}
}

/*GetLogsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetLogsUnauthorized struct {
	Payload *restmodels.Error
}

func (o *GetLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/logs][%d] getLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsNotFound creates a GetLogsNotFound with default headers values
func NewGetLogsNotFound() *GetLogsNotFound {
	return &GetLogsNotFound{}
}

/*GetLogsNotFound handles this case with default header values.

The model cannot be found.
*/
type GetLogsNotFound struct {
	Payload *restmodels.Error
}

func (o *GetLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/logs][%d] getLogsNotFound  %+v", 404, o.Payload)
}

func (o *GetLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
