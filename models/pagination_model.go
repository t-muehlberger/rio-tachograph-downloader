// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaginationModel pagination model
//
// swagger:model PaginationModel
type PaginationModel struct {

	// limit
	// Read Only: true
	Limit int32 `json:"limit,omitempty"`

	// next
	// Read Only: true
	Next string `json:"next,omitempty"`

	// offset
	// Read Only: true
	Offset int32 `json:"offset,omitempty"`

	// previous
	// Read Only: true
	Previous string `json:"previous,omitempty"`
}

// Validate validates this pagination model
func (m *PaginationModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this pagination model based on the context it is used
func (m *PaginationModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLimit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOffset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrevious(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginationModel) contextValidateLimit(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "limit", "body", int32(m.Limit)); err != nil {
		return err
	}

	return nil
}

func (m *PaginationModel) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "next", "body", string(m.Next)); err != nil {
		return err
	}

	return nil
}

func (m *PaginationModel) contextValidateOffset(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "offset", "body", int32(m.Offset)); err != nil {
		return err
	}

	return nil
}

func (m *PaginationModel) contextValidatePrevious(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "previous", "body", string(m.Previous)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginationModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationModel) UnmarshalBinary(b []byte) error {
	var res PaginationModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}