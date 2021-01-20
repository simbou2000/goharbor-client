// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmailServerSetting email server setting
//
// swagger:model EmailServerSetting
type EmailServerSetting struct {

	// The host of email server.
	EmailHost string `json:"email_host,omitempty"`

	// The dentity of email server.
	EmailIdentity string `json:"email_identity,omitempty"`

	// The password of email server.
	EmailPassword string `json:"email_password,omitempty"`

	// The port of email server.
	EmailPort int64 `json:"email_port,omitempty"`

	// Use ssl/tls or not.
	EmailSsl bool `json:"email_ssl,omitempty"`

	// The username of email server.
	EmailUsername string `json:"email_username,omitempty"`
}

// Validate validates this email server setting
func (m *EmailServerSetting) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this email server setting based on context it is used
func (m *EmailServerSetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailServerSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailServerSetting) UnmarshalBinary(b []byte) error {
	var res EmailServerSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
