package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateMarkdownFileForNoun(filename string, noun Noun) {
	noun.Fill()
	err := os.WriteFile(filename, []byte(MakeMarkdownForNoun(noun)), os.FileMode(0666))
	check(err)
}

func main() {
	noun := CreateNoun(Masculine, "integer", "integrī", 2)
	generateMarkdownFileForNoun("number-theory/integer.md", noun)
}
