package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
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
	dat, err := os.ReadFile("number-theory/words.toml")
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	str := string(dat)
	toml.Decode(str, &config)
	for _, nounConfig := range config.Nouns {
		entry := nounConfig.ToNounEntry()
		generateMarkdownFileForNoun(entry)
	}
}
