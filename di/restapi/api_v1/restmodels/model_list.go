// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ModelList model list
// swagger:model ModelList
type ModelList struct {

	// models
	Models []*Model `json:"models"`

	// pages
	Pages int64 `json:"pages,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this model list
func (m *ModelList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelList) validateModels(formats strfmt.Registry) error {

	if swag.IsZero(m.Models) { // not required
		return nil
	}

	for i := 0; i < len(m.Models); i++ {
		if swag.IsZero(m.Models[i]) { // not required
			continue
		}

		if m.Models[i] != nil {
			if err := m.Models[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("models" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelList) UnmarshalBinary(b []byte) error {
	var res ModelList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
