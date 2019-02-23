package dna

import (
	"strings"
)

func DNAStrand(dna string) string {
	dnaMapping := strings.NewReplacer("A", "T", "C", "G", "T", "A", "G", "C")
	return dnaMapping.Replace(dna)
}
