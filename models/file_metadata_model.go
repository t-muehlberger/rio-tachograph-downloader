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

// FileMetadataModel Tachograph File Metadata Response
//
// swagger:model FileMetadataModel
type FileMetadataModel struct {

	// Created by
	CreatedBy string `json:"created_by,omitempty"`

	// Customer id
	CustomerID string `json:"customer_id,omitempty"`

	// File identification. Globally unique.
	FileID int32 `json:"file_id,omitempty"`

	// File name
	FileName string `json:"file_name,omitempty"`

	// File type
	FileType string `json:"file_type,omitempty"`

	// Has sections valid
	HasSectionsValid bool `json:"has_sections_valid,omitempty"`

	// Is corrupted
	IsCorrupted bool `json:"is_corrupted,omitempty"`

	// Last modified by
	ModifiedBy string `json:"modified_by,omitempty"`

	// In case of file_type driver this is the driverId, when file_type vehicle this is the vehicle id / asset  id, for other cases this is empty
	// Format: uuid
	ReferenceID strfmt.UUID `json:"reference_id,omitempty"`

	// File size
	Size int32 `json:"size,omitempty"`

	// UTC value formatted as date-time (see: RFC 3339, section 5.6)
	// Format: date-time
	TimeCreated strfmt.DateTime `json:"time_created,omitempty"`

	// UTC value formatted as date-time (see: RFC 3339, section 5.6)
	// Format: date-time
	TimeModified strfmt.DateTime `json:"time_modified,omitempty"`
}

// Validate validates this file metadata model
func (m *FileMetadataModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeModified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileMetadataModel) validateReferenceID(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferenceID) { // not required
		return nil
	}

	if err := validate.FormatOf("reference_id", "body", "uuid", m.ReferenceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FileMetadataModel) validateTimeCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("time_created", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FileMetadataModel) validateTimeModified(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeModified) { // not required
		return nil
	}

	if err := validate.FormatOf("time_modified", "body", "date-time", m.TimeModified.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this file metadata model based on context it is used
func (m *FileMetadataModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileMetadataModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileMetadataModel) UnmarshalBinary(b []byte) error {
	var res FileMetadataModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
