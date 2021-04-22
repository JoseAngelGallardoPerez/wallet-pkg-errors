package errors

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Confialink/wallet-pkg-utils"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func validationErrorToPublic(e validator.FieldError) *PublicError {
	tagFormatter := getFormatter(e.Tag())
	formattedField := utils.Text.ToDelimiter(e.Field(), " ")
	if tagFormatter == nil {
		return &PublicError{
			Code:  CodeInvalid,
			Title: fmt.Sprintf("%s is not valid", formattedField),
		}
	}
	var title string
	var meta interface{}
	if tagFormatter.TitleFunc != nil {
		title = tagFormatter.TitleFunc(e, formattedField)
	}
	if tagFormatter.MetaFunc != nil {
		meta = tagFormatter.MetaFunc(e)

	}
	return &PublicError{
		Code:  tagFormatter.Code,
		Title: title,
		Meta:  meta,
	}
}

func ShouldBindToTyped(err error) TypedError {
	if result := ValidationErrorsToTyped(err); result != nil {
		return result
	}

	cause := errors.Cause(err)
	if unmarshalError, ok := cause.(*json.UnmarshalTypeError); ok {
		return &UnmarshalError{
			Err: unmarshalError,
		}
	}

	return &PrivateError{OriginalError: err}
}

func ValidationErrorsToTyped(err error) *ValidationErrors {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		result := ValidationErrors{}
		var source string
		for _, validationError := range validationErrors {
			publicErr := validationErrorToPublic(validationError)
			if len(validationError.Namespace()) > 0 {
				// First part of the namespace is always name of a structure.
				// We should cut this name.
				// Dynamic structures must be wrapped by an anonymous structure.
				ns := validationError.Namespace()
				if strings.Contains(validationError.Namespace(), ".") {
					r := strings.Split(ns, ".")
					ns = strings.Join(r[1:], ".")
				}
				source = toLowerCamel(ns)

			} else {
				source = toLowerCamel(validationError.Field())
			}
			result.Errors = append(result.Errors, ValidationError{
				Code:   publicErr.Code,
				Source: source,
				Title:  publicErr.Title,
				Meta:   publicErr.Meta,
			})
		}

		return &result
	}

	return nil
}
