// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Access access
//
// swagger:model Access
type Access struct {

	// The action of the access. Possible actions are *, pull, push, create, read, update, delete, list, operate, scanner-pull and stop.
	Action string `json:"action,omitempty"`

	// The effect of the access
	Effect string `json:"effect,omitempty"`

	// The resource of the access. Possible resources are *, artifact, artifact-addition, artifact-label, audit-log, catalog, configuration, distribution, garbage-collection, helm-chart, helm-chart-version, helm-chart-version-label, immutable-tag, label, ldap-user, log, member, metadata, notification-policy, preheat-instance, preheat-policy, project, quota, registry, replication, replication-adapter, replication-policy, repository, robot, scan, scan-all, scanner, system-volumes, tag, tag-retention, user, user-group or "" (for self-reference).
	Resource string `json:"resource,omitempty"`
}

// Validate validates this access
func (m *Access) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Access) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Access) UnmarshalBinary(b []byte) error {
	var res Access
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
