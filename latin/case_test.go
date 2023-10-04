package main

import (
	"testing"

	"github.com/bayashi/actually"
)

var genders = []Gender{Masculine, Feminine, Neuter}
var cases = []Case{Nominative, Genitive, Dative, Accusative, Ablative, Vocative}
var numbers = []Number{Singular, Plural}

func TestCaseFirstDeclension(t *testing.T) {
	{
		word := "amīca"
		noun := CreateNoun(Feminine, word, word+"e", 1)
		expected := [][]string{
			{"amīca", "amīcae", "amīcae", "amīcam", "amīcā", "amīca"},
			{"amīcae", "amīcārum", "amīcīs", "amīcās", "amīcīs", "amīcae"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}

	{
		noun := CreateNoun(Masculine, "Aenēās", "Aenēae", 1)
		expected := []string{"Aenēās", "Aenēae", "Aenēae", "Aenēam", "Aenēā", "Aenēās"}

		for j, case_ := range cases {
			actual := CaseFind(noun, Singular, case_)
			actually.Got(*actual).Expect(expected[j]).Same(t)
		}
	}
}

func TestCaseSecondDeclension(t *testing.T) {
	{
		noun := CreateNoun(Masculine, "dominus", "dominī", 2)
		expected := [][]string{
			{"dominus", "dominī", "dominō", "dominum", "dominō", "domine"},
			{"dominī", "dominōrum", "dominīs", "dominōs", "dominīs", "dominī"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}

	{
		noun := CreateNoun(Masculine, "puer", "puerī", 2)
		expected := [][]string{
			{"puer", "puerī", "puerō", "puerum", "puerō", "puer"},
			{"puerī", "puerōrum", "puerīs", "puerōs", "puerīs", "puerī"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}

	{
		noun := CreateNoun(Masculine, "magister", "magistrī", 2)
		expected := [][]string{
			{"magister", "magistrī", "magistrō", "magistrum", "magistrō", "magister"},
			{"magistrī", "magistrōrum", "magistrīs", "magistrōs", "magistrīs", "magistrī"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}
}

func TestCaseThirdDeclension(t *testing.T) {
	// non-neuter -ium
	{
		noun := CreateThirdDeclensionNoun(Feminine, "ars", "artis", "artium")
		expected := [][]string{
			{"ars", "artis", "artī", "artem", "arte", "ars"},
			{"artēs", "artium", "artibus", "artēs", "artibus", "artēs"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}
	// neuter -ium
	{
		noun := CreateThirdDeclensionNoun(Neuter, "animal", "animalis", "animalium")
		expected := [][]string{
			{"animal", "animalis", "animalī", "animal", "animalī", "animal"},
			{"animalia", "animalium", "animalibus", "animalia", "animalibus", "animalia"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}

	// non-neuter -um
	{
		noun := CreateThirdDeclensionNoun(Feminine, "salūs", "salūtis", "salūtum")
		expected := [][]string{
			{"salūs", "salūtis", "salūtī", "salūtem", "salūte", "salūs"},
			{"salūtēs", "salūtum", "salūtibus", "salūtēs", "salūtibus", "salūtēs"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}
	// neuter -um
	{
		noun := CreateThirdDeclensionNoun(Neuter, "genus", "generis", "generum")
		expected := [][]string{
			{"genus", "generis", "generī", "genus", "genere", "genus"},
			{"genera", "generum", "generibus", "genera", "generibus", "genera"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}
}

func TestCaseFourthDeclension(t *testing.T) {
	{
		noun := CreateNoun(Feminine, "manus", "manūs", 4)
		expected := [][]string{
			{"manus", "manūs", "manuī", "manum", "manū", "manus"},
			{"manūs", "manuum", "manibus", "manūs", "manibus", "manūs"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}
}

func TestCaseFifthDeclension(t *testing.T) {
	{
		noun := CreateNoun(Feminine, "speciēs", "speciēī", 5)
		expected := [][]string{
			{"speciēs", "speciēī", "speciēī", "speciem", "speciē", "speciēs"},
			{"speciēs", "speciērum", "speciēbus", "speciēs", "speciēbus", "speciēs"},
		}
		for i, number := range numbers {
			for j, case_ := range cases {
				actual := CaseFind(noun, number, case_)
				actually.Got(*actual).Expect(expected[i][j]).Same(t)
			}
		}
	}
}

func TestCaseAdjectiveFirstSecondDeclension(t *testing.T) {
	// -us, -a, -um
	{
		adjective := CreateAdjective("bonus", "bonī", "1+2")
		expected := [][][]string{
			{
				{"bonus", "bonī", "bonō", "bonum", "bonō", "bone"},
				{"bonī", "bonōrum", "bonīs", "bonōs", "bonīs", "bonī"},
			},
			{
				{"bona", "bonae", "bonae", "bonam", "bonā", "bona"},
				{"bonae", "bonārum", "bonīs", "bonās", "bonīs", "bonae"},
			},
			{
				{"bonum", "bonī", "bonō", "bonum", "bonō", "bonum"},
				{"bona", "bonōrum", "bonīs", "bona", "bonīs", "bona"},
			},
		}
		for i, gender := range genders {
			for j, number := range numbers {
				for k, case_ := range cases {
					actual := CaseFindAdjective(adjective, gender, number, case_)
					actually.Got(*actual).Expect(expected[i][j][k]).Same(t)
				}
			}
		}
	}
}
