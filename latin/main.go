package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateMarkdownFileForNoun(noun NounEntry) {
	filename := "number-theory/noun-" + noun.Filename + ".md"
	noun.Noun.Fill()
	err := os.WriteFile(filename, []byte(MakeMarkdownForNoun(noun)), os.FileMode(0666))
	check(err)
}

func main() {
	noun := CreateNoun(Masculine, "integer", "integrī", 2)
	entry := NounEntry{
		Noun:        noun,
		Filename:    "integer",
		Translation: "整数",
	}
	generateMarkdownFileForNoun(entry)
}
