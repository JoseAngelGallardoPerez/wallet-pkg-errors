package errors

import (
	"regexp"
	"strings"
)

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

// toLowerCamel converts a string to lowerCamelCase
// Examples:
// BeneficiaryBankDetails.Address -> beneficiaryBankDetails.address
// AdditionalInstructions -> additionalInstructions
// additional_instructions -> additionalInstructions
// PhysicalAddresses[1].City -> physicalAddresses.1.city
func toLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	nextUppertoLower := false
	for _, v := range s {
		switch true {
		case v >= 'A' && v <= 'Z':
			if !nextUppertoLower {
				n += string(v)
			} else {
				n += strings.ToLower(string(v))
			}
		case v >= '0' && v <= '9':
			n += string(v)
		case v >= 'a' && v <= 'z':
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}

		switch v {
		case '.':
			n += string(v)
			nextUppertoLower = true
		case '[':
			n += "."
			nextUppertoLower = false
		default:
			nextUppertoLower = false
		}

		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}
