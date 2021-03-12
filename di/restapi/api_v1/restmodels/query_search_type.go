// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// QuerySearchType query search type
// swagger:model QuerySearchType
type QuerySearchType string

const (

	// QuerySearchTypeTERM captures enum value "TERM"
	QuerySearchTypeTERM QuerySearchType = "TERM"

	// QuerySearchTypeNESTED captures enum value "NESTED"
	QuerySearchTypeNESTED QuerySearchType = "NESTED"

	// QuerySearchTypeMATCH captures enum value "MATCH"
	QuerySearchTypeMATCH QuerySearchType = "MATCH"

	// QuerySearchTypeALL captures enum value "ALL"
	QuerySearchTypeALL QuerySearchType = "ALL"
)

// for schema
var querySearchTypeEnum []interface{}

func init() {
	var res []QuerySearchType
	if err := json.Unmarshal([]byte(`["TERM","NESTED","MATCH","ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		querySearchTypeEnum = append(querySearchTypeEnum, v)
	}
}

func (m QuerySearchType) validateQuerySearchTypeEnum(path, location string, value QuerySearchType) error {
	if err := validate.Enum(path, location, value, querySearchTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this query search type
func (m QuerySearchType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateQuerySearchTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
