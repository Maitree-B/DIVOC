// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Enrollment enrollment
//
// swagger:model enrollment
type Enrollment struct {

	// address
	Address *Address `json:"address,omitempty"`

	// appointments
	Appointments []*EnrollmentAppointmentsItems0 `json:"appointments"`

	// beneficiary phone
	BeneficiaryPhone string `json:"beneficiaryPhone,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// comorbidities
	Comorbidities []string `json:"comorbidities"`

	// dob
	// Format: date
	Dob strfmt.Date `json:"dob,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// gender
	// Enum: [Male Female Other]
	Gender string `json:"gender,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// national Id
	// Required: true
	NationalID *string `json:"nationalId"`

	// phone
	Phone string `json:"phone,omitempty"`

	// yob
	Yob int64 `json:"yob,omitempty"`
}

// Validate validates this enrollment
func (m *Enrollment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppointments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNationalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Enrollment) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *Enrollment) validateAppointments(formats strfmt.Registry) error {

	if swag.IsZero(m.Appointments) { // not required
		return nil
	}

	for i := 0; i < len(m.Appointments); i++ {
		if swag.IsZero(m.Appointments[i]) { // not required
			continue
		}

		if m.Appointments[i] != nil {
			if err := m.Appointments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appointments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Enrollment) validateDob(formats strfmt.Registry) error {

	if swag.IsZero(m.Dob) { // not required
		return nil
	}

	if err := validate.FormatOf("dob", "body", "date", m.Dob.String(), formats); err != nil {
		return err
	}

	return nil
}

var enrollmentTypeGenderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Male","Female","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enrollmentTypeGenderPropEnum = append(enrollmentTypeGenderPropEnum, v)
	}
}

const (

	// EnrollmentGenderMale captures enum value "Male"
	EnrollmentGenderMale string = "Male"

	// EnrollmentGenderFemale captures enum value "Female"
	EnrollmentGenderFemale string = "Female"

	// EnrollmentGenderOther captures enum value "Other"
	EnrollmentGenderOther string = "Other"
)

// prop value enum
func (m *Enrollment) validateGenderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, enrollmentTypeGenderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Enrollment) validateGender(formats strfmt.Registry) error {

	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", m.Gender); err != nil {
		return err
	}

	return nil
}

func (m *Enrollment) validateNationalID(formats strfmt.Registry) error {

	if err := validate.Required("nationalId", "body", m.NationalID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Enrollment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Enrollment) UnmarshalBinary(b []byte) error {
	var res Enrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EnrollmentAppointmentsItems0 enrollment appointments items0
//
// swagger:model EnrollmentAppointmentsItems0
type EnrollmentAppointmentsItems0 struct {

	// appointment date
	// Format: date
	AppointmentDate strfmt.Date `json:"appointmentDate,omitempty"`

	// appointment slot
	AppointmentSlot string `json:"appointmentSlot,omitempty"`

	// certified
	Certified *bool `json:"certified,omitempty"`

	// dose
	Dose string `json:"dose,omitempty"`

	// enrollment scope Id
	EnrollmentScopeID string `json:"enrollmentScopeId,omitempty"`

	// program Id
	ProgramID string `json:"programId,omitempty"`
}

// Validate validates this enrollment appointments items0
func (m *EnrollmentAppointmentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppointmentDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnrollmentAppointmentsItems0) validateAppointmentDate(formats strfmt.Registry) error {

	if swag.IsZero(m.AppointmentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("appointmentDate", "body", "date", m.AppointmentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnrollmentAppointmentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnrollmentAppointmentsItems0) UnmarshalBinary(b []byte) error {
	var res EnrollmentAppointmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
