package main

import (
	"fmt"
	"strings"
)

type Case int

// Don't know how to say it correctly
type Number int

type Gender int

const (
	Singular Number = iota
	Plural
)

const (
	Nominative Case = iota
	Genitive
	Dative
	Accusative
	Ablative
	Vocative
)

const (
	Masculine Gender = iota
	Feminine
	Neuter
)

type Noun struct {
	Gender         Gender
	DeclensionType int

	// nil: not yet filled
	// &"": actually missing
	Declensions [2][6]*string
}

func CreateNoun(gender Gender, nom string, gen string, decltype int) Noun {
	return Noun{
		Gender:         gender,
		DeclensionType: decltype,
		Declensions:    [2][6]*string{{&nom, &gen}, {}},
	}
}

func CreateThirdDeclensionNoun(gender Gender, nom string, gen string, pluralGen string) Noun {
	return Noun{
		Gender:         gender,
		DeclensionType: 3,
		Declensions:    [2][6]*string{{&nom, &gen}, {nil, &pluralGen}},
	}
}

// CaseFind finds the declension for a noun and a case.
// noun.Declensions[0][0] and noun.Declensions[0][1] should be given.
// If something wrong happens, this function returns nil.
func CaseFind(noun Noun, number Number, case_ Case) *string {
	nom := *noun.Declensions[0][0]
	gen := *noun.Declensions[0][1]

	if noun.DeclensionType == 1 {
		// -a, -ae
		prefix, found := strings.CutSuffix(*noun.Declensions[0][0], "a")
		if !found {
			// Like Aenēās, nominative can be not regular.
			prefix, found = strings.CutSuffix(*noun.Declensions[0][1], "ae")
			if !found {
				return nil
			}
		}
		if number == Singular && (case_ == Nominative || case_ == Vocative) {
			return noun.Declensions[0][0]
		}
		value := prefix + [][]string{{"a", "ae", "ae", "am", "ā", "a"}, {"ae", "ārum", "īs", "ās", "īs", "ae"}}[number][case_]
		return &value
	}
	if noun.DeclensionType == 2 {
		// -us, -ī / -um, -ī
		if noun.Gender == Neuter {
			// -um, -ī
			prefix, found := strings.CutSuffix(nom, "um")
			if !found {
				return nil
			}
			value := prefix + [][]string{{"um", "ī", "ō", "um", "ō", "um"}, {"a", "ōrum", "īs", "a", "īs", "a"}}[number][case_]
			return &value
		}
		if number == Singular && case_ == Nominative {
			return noun.Declensions[0][0]
		}
		if number == Singular && case_ == Vocative {
			if a, found := strings.CutSuffix(nom, "us"); found {
				b := a + "e"
				return &b
			}
			return noun.Declensions[0][0]
		}
		var prefix string
		if p, found := strings.CutSuffix(nom, "us"); found {
			prefix = p
		} else {
			// e.g. magister, puer, vir, liber
			p, found = strings.CutSuffix(gen, "ī") // magistr-, puer-, vir-, libr-
			if !found {
				return nil
			}
			prefix = p
		}
		value := prefix + [][]string{{"us", "ī", "ō", "um", "ō", "e"}, {"ī", "ōrum", "īs", "ōs", "īs", "ī"}}[number][case_]
		return &value
	}
	if noun.DeclensionType == 4 {
		// -us, -ūs
		// TODO: support -ū
		prefix, found := strings.CutSuffix(*noun.Declensions[0][0], "us")
		if !found {
			return nil
		}
		value := prefix + [][]string{{"us", "ūs", "uī", "um", "ū", "us"}, {"ūs", "uum", "ibus", "ūs", "ibus", "ūs"}}[number][case_]
		return &value
	}
	if noun.DeclensionType == 5 {
		// -ēs, ēī
		prefix, found := strings.CutSuffix(*noun.Declensions[0][0], "ēs")
		if !found {
			return nil
		}
		value := prefix + [][]string{{"ēs", "ēī", "ēī", "em", "ē", "ēs"}, {"ēs", "ērum", "ēbus", "ēs", "ēbus", "ēs"}}[number][case_]
		return &value
	}
	if noun.DeclensionType != 3 {
		return nil
	}
	// 3rd declension: the hardest part
	prefix, found := strings.CutSuffix(*noun.Declensions[0][1], "is")
	if !found {
		return nil
	}
	if number == Singular && (case_ == Nominative || case_ == Vocative) {
		return noun.Declensions[0][0]
	}
	if number == Plural && case_ == Genitive {
		// plural genitive can be one of -um and -ium
		// since there's no telling which is which, we require
		// the plural genitive to be given.
		return noun.Declensions[Plural][Genitive]
	}
	if noun.Gender != Neuter {
		value := prefix + [][]string{{"", "is", "ī", "em", "e", ""}, {"ēs", "um", "ibus", "ēs", "ibus", "ēs"}}[number][case_]
		return &value
	}
	// neuter; accusative = nominative
	if number == Singular && case_ == Accusative {
		return noun.Declensions[0][0]
	}
	value := prefix + [][]string{{"", "is", "ī", "", "e", ""}, {"a", "um", "ibus", "a", "ibus", "a"}}[number][case_]
	return &value
}

func main() {
	fmt.Println("2")
}
