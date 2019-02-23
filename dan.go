package dna

func DNAStrand(dna string) string {
	dnaMapping := map[string]string{"A": "T", "C": "G", "T": "A", "G": "C"}
	var dnaComplement string
	for _, monomeric := range dna {
		dnaComplement += dnaMapping[string(monomeric)]
	}
	return dnaComplement
}
