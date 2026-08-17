package proteintranslation

import (
	"errors"
)

var (
	ErrStop        = errors.New("stop codon")
	ErrInvalidBase = errors.New("invalid codon")
)

var codonTable = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	result := make([]string, 0)
	for i := 0; i+2 < len(rna); i += 3 {
		codon := rna[i : i+3]
		amino, err := FromCodon(codon)
		if err != nil {
			if errors.Is(err, ErrStop) {
				return result, nil
			}
			return nil, err
		}
		result = append(result, amino)
	}
    if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	val, ok := codonTable[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if val == "STOP" {
		return "", ErrStop
	}
	return val, nil
}