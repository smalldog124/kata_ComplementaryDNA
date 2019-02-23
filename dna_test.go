package dna_test

import (
	dna "kata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DNAStrand_Input_ATTGC_Should_Be_TAACG(t *testing.T) {
	expected := "TAACG"
	dnaStrand := "ATTGC"

	actual := dna.DNAStrand(dnaStrand)

	assert.Equal(t, expected, actual)
}
func Test_DNAStrand_Input_GTAT_Should_Be_CATA(t *testing.T) {
	expected := "GTAT"
	dnaStrand := "CATA"

	actual := dna.DNAStrand(dnaStrand)

	assert.Equal(t, expected, actual)
}
