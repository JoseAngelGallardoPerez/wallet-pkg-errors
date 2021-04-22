package errors

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	CodeRequired            = "REQUIRED"
	CodeMax                 = "MAX"
	CodeMin                 = "MIN"
	CodeEmail               = "EMAIL"
	CodeLen                 = "LEN"
	CodeEqField             = "EQ_FIELD"
	CodeDecimalValidation   = "DECIMAL_VALIDATION"
	CodeDecimalGTValidation = "DECIMAL_GT_VALIDATION"
	CodeInvalid             = "INVALID_FIELD"
	CodeUnmarshalJSON       = "UNMARSHAL_JSON"
)

type TitleFunc func(e validator.FieldError, formattedField string) string
type MetaFunc func(e validator.FieldError) interface{}

type ValidationErrorFormatter struct {
	Code      string
	TitleFunc TitleFunc
	MetaFunc  MetaFunc
}

func SetFormatters(formatters map[string]*ValidationErrorFormatter) {
	customFormatters = formatters
}

var customFormatters map[string]*ValidationErrorFormatter

func defaultFormatters() map[string]*ValidationErrorFormatter {
	return map[string]*ValidationErrorFormatter{
		"required": {
			Code: CodeRequired,
			TitleFunc: func(_ validator.FieldError, formattedField string) string {
				return fmt.Sprintf("%s is required and cannot be empty", formattedField)
			},
		},
		"max": {
			Code: CodeMax,
			TitleFunc: func(e validator.FieldError, formattedField string) string {
				return fmt.Sprintf("%s cannot be longer than %s", formattedField, e.Param())
			},
			MetaFunc: func(e validator.FieldError) interface{} {
				return struct {
					Max string `json:"max"`
				}{e.Param()}
			},
		},
		"min": {
			Code: CodeMin,
			TitleFunc: func(e validator.FieldError, formattedField string) string {
				return fmt.Sprintf("%s must be longer than %s", formattedField, e.Param())
			},
			MetaFunc: func(e validator.FieldError) interface{} {
				return struct {
					Min string `json:"min"`
				}{e.Param()}
			},
		},
		"email": {
			Code: CodeEmail,
			TitleFunc: func(e validator.FieldError, _ string) string {
				return fmt.Sprintf("'%s' is not a valid email address", e.Value())
			},
		},
		"len": {
			Code: CodeLen,
			TitleFunc: func(e validator.FieldError, formattedField string) string {
				return fmt.Sprintf("%s must be %s characters long", formattedField, e.Param())
			},
			MetaFunc: func(e validator.FieldError) interface{} {
				return struct {
					Len string `json:"len"`
				}{e.Param()}
			},
		},
		"eqfield": {
			Code: CodeEqField,
			TitleFunc: func(e validator.FieldError, formattedField string) string {
				return fmt.Sprintf("%s must match the %s", formattedField, e.Param())
			},
			MetaFunc: func(e validator.FieldError) interface{} {
				return struct {
					Field string `json:"field"`
				}{e.Param()}
			},
		},
		"pwdpolicy": {
			Code: "PWD_POLICY",
			TitleFunc: func(e validator.FieldError, _ string) string {
				return fmt.Sprintf("%s must contain upper and lower case characters, and at least one digit and special character", e.Field())
			},
		},
		"decimal": {
			Code: CodeDecimalValidation,
			TitleFunc: func(_ validator.FieldError, formattedField string) string {
				return fmt.Sprintf("%s should be decimal", formattedField)
			},
		},
		"decimalGT": {
			Code: CodeDecimalGTValidation,
			TitleFunc: func(e validator.FieldError, formattedField string) string {
				return fmt.Sprintf("%s should be greater than %s", formattedField, e.Param())
			},
			MetaFunc: func(e validator.FieldError) interface{} {
				return struct {
					Value string `json:"value"`
				}{e.Param()}
			},
		},
	}
}

func getFormatter(tag string) *ValidationErrorFormatter {
	if customFormatters != nil {
		if formatter, ok := customFormatters[tag]; ok {
			return formatter
		}
	}
	return defaultFormatters()[tag]
}
