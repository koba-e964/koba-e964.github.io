package main

type NounEntry struct {
	Noun        Noun
	Filename    string
	Translation string
	Examples    []struct {
		Latin    string `toml:"latin"`
		Japanese string `toml:"japanese"`
	}
}

type NounConfig struct {
	Title          string `toml:"title"`
	Synopsis       string `toml:"synopsis"`
	Genitive       string `toml:"genitive"`
	PluralGenitive string `toml:"plural_genitive"`
	Declension     int    `toml:"declension"`
	Gender         string `toml:"gender"`
	Translation    string `toml:"translation"`
	// example sentence
	Examples []struct {
		Latin    string `toml:"latin"`
		Japanese string `toml:"japanese"`
	} `toml:"examples"`
}

func (conf NounConfig) ToNounEntry() NounEntry {
	var noun Noun
	var gender Gender
	if conf.Gender == "masculine" {
		gender = Masculine
	} else if conf.Gender == "feminine" {
		gender = Feminine
	} else if conf.Gender == "neuter" {
		gender = Neuter
	} else {
		panic(conf.Gender)
	}
	if conf.Declension != 3 {
		noun = CreateNoun(gender, conf.Synopsis, conf.Genitive, conf.Declension)
	} else {
		noun = CreateThirdDeclensionNoun(gender, conf.Synopsis, conf.Genitive, conf.PluralGenitive)
	}
	return NounEntry{
		Noun:        noun,
		Filename:    conf.Title,
		Translation: conf.Translation,
		Examples:    conf.Examples,
	}
}
