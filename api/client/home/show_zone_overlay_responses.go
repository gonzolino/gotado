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

// ShowZoneOverlayReader is a Reader for the ShowZoneOverlay structure.
type ShowZoneOverlayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowZoneOverlayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowZoneOverlayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewShowZoneOverlayUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowZoneOverlayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowZoneOverlayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowZoneOverlayOK creates a ShowZoneOverlayOK with default headers values
func NewShowZoneOverlayOK() *ShowZoneOverlayOK {
	return &ShowZoneOverlayOK{}
}

/* ShowZoneOverlayOK describes a response with status code 200, with default header values.

Zone overlay.
*/
type ShowZoneOverlayOK struct {
	Payload *models.Overlay
}

func (o *ShowZoneOverlayOK) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/overlay][%d] showZoneOverlayOK  %+v", 200, o.Payload)
}
func (o *ShowZoneOverlayOK) GetPayload() *models.Overlay {
	return o.Payload
}

func (o *ShowZoneOverlayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, _ strfmt.Registry) error {

	o.Payload = new(models.Overlay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowZoneOverlayUnauthorized creates a ShowZoneOverlayUnauthorized with default headers values
func NewShowZoneOverlayUnauthorized() *ShowZoneOverlayUnauthorized {
	return &ShowZoneOverlayUnauthorized{}
}

/* ShowZoneOverlayUnauthorized describes a response with status code 401, with default header values.

User authentication failed.
*/
type ShowZoneOverlayUnauthorized struct {
	Payload *models.ClientErrorModel
}

func (o *ShowZoneOverlayUnauthorized) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/overlay][%d] showZoneOverlayUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowZoneOverlayUnauthorized) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *ShowZoneOverlayUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, _ strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowZoneOverlayForbidden creates a ShowZoneOverlayForbidden with default headers values
func NewShowZoneOverlayForbidden() *ShowZoneOverlayForbidden {
	return &ShowZoneOverlayForbidden{}
}

/* ShowZoneOverlayForbidden describes a response with status code 403, with default header values.

Authenticated user has no access rights to the requested entity.
*/
type ShowZoneOverlayForbidden struct {
	Payload *models.ClientErrorModel
}

func (o *ShowZoneOverlayForbidden) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/overlay][%d] showZoneOverlayForbidden  %+v", 403, o.Payload)
}
func (o *ShowZoneOverlayForbidden) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *ShowZoneOverlayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, _ strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowZoneOverlayNotFound creates a ShowZoneOverlayNotFound with default headers values
func NewShowZoneOverlayNotFound() *ShowZoneOverlayNotFound {
	return &ShowZoneOverlayNotFound{}
}

/* ShowZoneOverlayNotFound describes a response with status code 404, with default header values.

Requested entity not found.
*/
type ShowZoneOverlayNotFound struct {
	Payload *models.ClientErrorModel
}

func (o *ShowZoneOverlayNotFound) Error() string {
	return fmt.Sprintf("[GET /homes/{home_id}/zones/{zone_id}/overlay][%d] showZoneOverlayNotFound  %+v", 404, o.Payload)
}
func (o *ShowZoneOverlayNotFound) GetPayload() *models.ClientErrorModel {
	return o.Payload
}

func (o *ShowZoneOverlayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, _ strfmt.Registry) error {

	o.Payload = new(models.ClientErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
