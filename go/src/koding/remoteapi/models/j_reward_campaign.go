package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JRewardCampaign j reward campaign
// swagger:model JRewardCampaign
type JRewardCampaign struct {

	// id
	ID string `json:"_id,omitempty"`

	// created at
	CreatedAt strfmt.Date `json:"createdAt,omitempty"`

	// end date
	EndDate strfmt.Date `json:"endDate,omitempty"`

	// given amount
	GivenAmount float64 `json:"givenAmount,omitempty"`

	// initial amount
	InitialAmount float64 `json:"initialAmount,omitempty"`

	// is active
	IsActive bool `json:"isActive,omitempty"`

	// max amount
	MaxAmount float64 `json:"maxAmount,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// per event amount
	PerEventAmount float64 `json:"perEventAmount,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// start date
	StartDate strfmt.Date `json:"startDate,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// unit
	Unit string `json:"unit,omitempty"`
}

// Validate validates this j reward campaign
func (m *JRewardCampaign) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JRewardCampaign) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
