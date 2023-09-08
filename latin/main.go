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

func generateMarkdownFileForNoun(directory string, noun NounEntry) {
	filename := directory + "/noun-" + noun.Filename + ".md"
	noun.Noun.Fill()
	err := os.WriteFile(filename, []byte(MakeMarkdownForNoun(noun)), os.FileMode(0666))
	check(err)
}

func main() {
	directories := []string{"number-theory", "imperative-intro"}
	for _, directory := range directories {
		dat, err := os.ReadFile(directory + "/words.toml")
		if err != nil {
			log.Fatal(err)
		}
		var config Config
		str := string(dat)
		toml.Decode(str, &config)
		for _, nounConfig := range config.Nouns {
			entry := nounConfig.ToNounEntry()
			generateMarkdownFileForNoun(directory, entry)
		}
	}
}
