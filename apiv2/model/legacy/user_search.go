// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserSearch user search
//
// swagger:model UserSearch
type UserSearch struct {

	// The ID of the user.
	UserID int64 `json:"user_id,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this user search
func (m *UserSearch) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user search based on context it is used
func (m *UserSearch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSearch) UnmarshalBinary(b []byte) error {
	var res UserSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
