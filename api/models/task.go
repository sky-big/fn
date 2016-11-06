package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

const (
	// TypeNone ...
	TypeNone = ""
	// TypeSync ...
	TypeSync = "sync"
	// TypeAsync ...
	TypeAsync = "async"
)

/*Task task

swagger:model Task
*/
type Task struct {
	NewTask

	IDStatus

	/* Time when task completed, whether it was successul or failed. Always in UTC.
	 */
	CompletedAt strfmt.DateTime `json:"completed_at,omitempty"`

	/* Time when task was submitted. Always in UTC.

	Read Only: true
	*/
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	/* Env vars for the task. Comes from the ones set on the Route.
	 */
	EnvVars map[string]string `json:"env_vars,omitempty"`

	/* The error message, if status is 'error'. This is errors due to things outside the task itself. Errors from user code will be found in the log.
	 */
	Error string `json:"error,omitempty"`

	/* App this task belongs to.

	Read Only: true
	*/
	AppName string `json:"app_name,omitempty"`

	Path string `json:"path"`

	/* Machine usable reason for task being in this state.
	Valid values for error status are `timeout | killed | bad_exit`.
	Valid values for cancelled status are `client_request`.
	For everything else, this is undefined.

	*/
	Reason string `json:"reason,omitempty"`

	/* If this field is set, then this task was retried by the task referenced in this field.

	Read Only: true
	*/
	RetryAt string `json:"retry_at,omitempty"`

	/* If this field is set, then this task is a retry of the ID in this field.

	Read Only: true
	*/
	RetryOf string `json:"retry_of,omitempty"`

	/* Time when task started execution. Always in UTC.
	 */
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.NewTask.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.IDStatus.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvVars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateEnvVars(formats strfmt.Registry) error {

	if err := validate.Required("env_vars", "body", m.EnvVars); err != nil {
		return err
	}

	return nil
}

var taskTypeReasonPropEnum []interface{}

// property enum
func (m *Task) validateReasonEnum(path, location string, value string) error {
	if taskTypeReasonPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["timeout","killed","bad_exit","client_request"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			taskTypeReasonPropEnum = append(taskTypeReasonPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, taskTypeReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Task) validateReason(formats strfmt.Registry) error {

	// value enum
	if err := m.validateReasonEnum("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}
