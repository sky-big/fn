// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StatMetrics stat metrics
// swagger:model statMetrics
type StatMetrics struct {

	// cpu kernel
	CPUKernel int64 `json:"cpu_kernel,omitempty"`

	// cpu total
	CPUTotal int64 `json:"cpu_total,omitempty"`

	// cpu user
	CPUUser int64 `json:"cpu_user,omitempty"`

	// disk read
	DiskRead int64 `json:"disk_read,omitempty"`

	// disk write
	DiskWrite int64 `json:"disk_write,omitempty"`

	// mem limit
	MemLimit int64 `json:"mem_limit,omitempty"`

	// mem usage
	MemUsage int64 `json:"mem_usage,omitempty"`

	// net rx
	NetRx int64 `json:"net_rx,omitempty"`

	// net tx
	NetTx int64 `json:"net_tx,omitempty"`
}

// Validate validates this stat metrics
func (m *StatMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StatMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatMetrics) UnmarshalBinary(b []byte) error {
	var res StatMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}