// このスクリプトからの相対パス XXXX/words.toml を見て、
// XXXX/YYYY-FILENAME.md にファイルを配置する。
// 一応再帰的にファイルを見て、1 段下以外でも words.toml があれば反応する。

package main

import (
	"log"
	"os"
	"path/filepath"

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
	directories := []string{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		basename := filepath.Base(path)
		if !info.IsDir() && basename == "words.toml" {
			log.Println(path)
			dirname := filepath.Dir(path)
			directories = append(directories, dirname)
		}
		return nil
	})
	check(err)

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
