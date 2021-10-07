// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gonzolino/gotado/api/models"
)

// ShowZoneCapabilitiesReader is a Reader for the ShowZoneCapabilities structure.
type ShowZoneCapabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowZoneCapabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowZoneCapabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewShowZoneCapabilitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowZoneCapabilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowZoneCapabilitiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowZoneCapabilitiesOK creates a ShowZoneCapabilitiesOK with default headers values
func NewShowZoneCapabilitiesOK() *ShowZoneCapabilitiesOK {
	return &ShowZoneCapabilitiesOK{}
}

/* ShowZoneCapabilitiesOK describes a response with status code 200, with default header values.

Zone capabilities.
*/
type ShowZoneCapabilitiesOK struct {
	Payload models.GenericZoneCapabilities
}

func (o *ShowZoneCapabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/capabilities][%d] showZoneCapabilitiesOK  %+v", 200, o.Payload)
}
func (o *ShowZoneCapabilitiesOK) GetPayload() models.GenericZoneCapabilities {
	return o.Payload
}

func (o *ShowZoneCapabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, _ strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGenericZoneCapabilities(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewShowZoneCapabilitiesUnauthorized creates a ShowZoneCapabilitiesUnauthorized with default headers values
func NewShowZoneCapabilitiesUnauthorized() *ShowZoneCapabilitiesUnauthorized {
	return &ShowZoneCapabilitiesUnauthorized{}
}

/* ShowZoneCapabilitiesUnauthorized describes a response with status code 401, with default header values.

User authentication failed.
*/
type ShowZoneCapabilitiesUnauthorized struct {
	Payload *models.ClientErrorModel
}

func (o *ShowZoneCapabilitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/capabilities][%d] showZoneCapabilitiesUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowZoneCapabilitiesUnauthorized) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *ShowZoneCapabilitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, _ strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowZoneCapabilitiesForbidden creates a ShowZoneCapabilitiesForbidden with default headers values
func NewShowZoneCapabilitiesForbidden() *ShowZoneCapabilitiesForbidden {
	return &ShowZoneCapabilitiesForbidden{}
}

/* ShowZoneCapabilitiesForbidden describes a response with status code 403, with default header values.

Authenticated user has no access rights to the requested entity.
*/
type ShowZoneCapabilitiesForbidden struct {
	Payload *models.ClientErrorModel
}

func (o *ShowZoneCapabilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/capabilities][%d] showZoneCapabilitiesForbidden  %+v", 403, o.Payload)
}
func (o *ShowZoneCapabilitiesForbidden) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *ShowZoneCapabilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, _ strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowZoneCapabilitiesNotFound creates a ShowZoneCapabilitiesNotFound with default headers values
func NewShowZoneCapabilitiesNotFound() *ShowZoneCapabilitiesNotFound {
	return &ShowZoneCapabilitiesNotFound{}
}

/* ShowZoneCapabilitiesNotFound describes a response with status code 404, with default header values.

Requested entity not found.
*/
type ShowZoneCapabilitiesNotFound struct {
	Payload *models.ClientErrorModel
}

func (o *ShowZoneCapabilitiesNotFound) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/capabilities][%d] showZoneCapabilitiesNotFound  %+v", 404, o.Payload)
}
func (o *ShowZoneCapabilitiesNotFound) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *ShowZoneCapabilitiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, _ strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
