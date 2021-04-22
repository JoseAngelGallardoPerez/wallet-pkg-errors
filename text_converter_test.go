package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactoryInitForms(t *testing.T) {

	cases := []struct {
		input    string
		expected string
	}{
		{"BeneficiaryBankDetails.Address", "beneficiaryBankDetails.address"},
		{"BeneficiaryBankDetails.address", "beneficiaryBankDetails.address"},
		{"Beneficiary_Bank_Details.Address", "beneficiaryBankDetails.address"},
		{"AdditionalInstructions", "additionalInstructions"},
		{"additional_instructions", "additionalInstructions"},
		{"additional_instructions.pt", "additionalInstructions.pt"},
		{"PhysicalAddresses[1].City", "physicalAddresses.1.city"},
	}

	for i, oneCase := range cases {
		assert.Equal(t, oneCase.expected, toLowerCamel(oneCase.input), fmt.Sprintf("case %d is invalid", i))
	}
}
